#!/usr/bin/env bash
#
# Generate one SDK module from its OpenAPI specification in a staging directory,
# then sync only the files that actually changed back into the module.
#
# Usage: ./scripts/generate-module.sh <git-user-id> <git-repo-id> <module>
#
# Environment:
#   DRY_RUN=1       Report what would change without writing anything.
#   CHECK=1         Implies DRY_RUN and exits non-zero if anything would change.
#   KEEP_STAGING=1  Keep the staging directory on failure for debugging.
#
# Prerequisites: docker (with the pinned openapi-generator image below), go,
# goimports in $(go env GOPATH)/bin or on PATH.

set -euo pipefail

GIT_USER_ID="${1:-}"
GIT_REPO_ID="${2:-}"
MODULE="${3:-}"

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
MODULE_DIR="$REPO_ROOT/$MODULE"

OPENAPI_GENERATOR_IMAGE="openapitools/openapi-generator-cli:v7.0.1"
STAGE_ROOT="$REPO_ROOT/.generate"
STAGE_DIR="$STAGE_ROOT/$MODULE"

DRY_RUN="${DRY_RUN:-0}"
CHECK="${CHECK:-0}"
KEEP_STAGING="${KEEP_STAGING:-0}"
if [ "$CHECK" = "1" ]; then
    DRY_RUN=1
fi

# Hand-maintained or ignore-listed files the sync must never touch.
PRESERVE_PATTERN='go.mod|go.sum|.version|CHANGELOG.md|GNUmakefile|client.go|openapitools.json|generate/*|test/*'

# Files that belong to the SDK even though the generator no longer emits them.
NON_GENERATED_FILES='api_hal_ext.go|api_utils_pagination_ext.go|model_entity_array_ext.go|model_paged_cursor_ext.go|client_ext.go|docs/EntityArrayPagedIterator.md|docs/PagedCursor.md'

log()  { echo "==> $*"; }
warn() { echo "==> WARNING: $*" >&2; }
die()  { echo "==> ERROR: $*" >&2; exit 1; }

[ -n "$GIT_USER_ID" ] || die "usage: generate-module.sh <git-user-id> <git-repo-id> <module>"
[ -n "$GIT_REPO_ID" ] || die "usage: generate-module.sh <git-user-id> <git-repo-id> <module>"
[ -n "$MODULE" ] || die "usage: generate-module.sh <git-user-id> <git-repo-id> <module>"
[ -d "$MODULE_DIR" ] || die "Module $MODULE not found at $MODULE_DIR"
[ -f "$MODULE_DIR/generate/pingone-$MODULE.yml" ] || die "Spec missing: $MODULE/generate/pingone-$MODULE.yml"
command -v docker >/dev/null 2>&1 || die "docker not found on PATH (required to run $OPENAPI_GENERATOR_IMAGE)"
command -v go >/dev/null 2>&1 || die "go not found on PATH"

VERSION="$(head -n 1 "$MODULE_DIR/.version")"

TMP_FILES="$(mktemp -t genmod-addded.XXXXXX) $(mktemp -t genmod-modified.XXXXXX) $(mktemp -t genmod-stagefiles.XXXXXX)"
cleanup() {
    rm -f $TMP_FILES
    if [ "$KEEP_STAGING" = "1" ]; then
        log "Staging kept at $STAGE_DIR"
    else
        rm -rf "$STAGE_DIR"
        rmdir "$STAGE_ROOT" 2>/dev/null || true
    fi
}
trap cleanup EXIT

# --- generate into staging ----------------------------------------------------

rm -rf "$STAGE_DIR"
mkdir -p "$STAGE_DIR"

# Copy the module's ignore list into the stage root: the generator honours an
# .openapi-generator-ignore found in the output directory. Do NOT pass
# --ignore-file-override - with it, the generator writes ignored files anyway.
# go.mod/go.sum are restored AFTER the docker run because the generator writes
# its own (stale) go.mod; the committed one is needed for the build gate below.
cp "$MODULE_DIR/.openapi-generator-ignore" "$STAGE_DIR/"

log "Generating $MODULE v$VERSION from OpenAPI (staged in .generate/$MODULE)..."

docker run --rm \
    -v "$REPO_ROOT:/local" \
    --user "$(id -u):$(id -g)" \
    "$OPENAPI_GENERATOR_IMAGE" generate \
    -t /local/scripts/templates/generator \
    -i "/local/$MODULE/generate/pingone-$MODULE.yml" \
    -g go \
    --additional-properties=packageName="$MODULE",packageVersion="$VERSION",isGoSubmodule=true,enumClassPrefix=true,apiNameSuffix=Api \
    -o "/local/.generate/$MODULE" \
    --git-repo-id "$GIT_REPO_ID" \
    --git-user-id "$GIT_USER_ID" \
    --http-user-agent "pingtools PingOne-GOLANG-SDK-$MODULE/$VERSION"

# --- post-process the stage (same order as the previous in-place flow) --------

log "Applying file post-processing..."

# The generator writes its own (stale) go.mod; restore the committed one so
# formatting and the build gate see a real module. Ignore-listed module files
# (client.go etc.) are restored AFTER the regex passes below - restoring earlier
# would let the regexes re-wrap already-processed code.
cp "$MODULE_DIR/go.mod" "$MODULE_DIR/go.sum" "$STAGE_DIR/"

GOFMT="$(command -v gofmt || echo "$(go env GOROOT)/bin/gofmt")"
[ -x "$GOFMT" ] || die "gofmt not found at $GOFMT"
# Build the repo-pinned goimports (declared in the root go.mod tool list) once
# into a temp binary. `go tool goimports` is not usable directly - the tool
# only exists after `go work vendor`, and vendoring state varies per checkout.
GOIMPORTS_BIN="$(mktemp -t goimports-pinned.XXXXXX)"
TMP_FILES="$TMP_FILES $GOIMPORTS_BIN"
if ! ( cd "$REPO_ROOT" && GOWORK=off go build -o "$GOIMPORTS_BIN" golang.org/x/tools/cmd/goimports ) >/dev/null 2>&1; then
    die "could not build the pinned goimports. Run 'go mod tidy' at the repo root, or install goimports in $(go env GOPATH)/bin"
fi

find "$STAGE_DIR" -name "*.go" -exec "$GOFMT" -w -s {} \;
find "$STAGE_DIR" -name "*.go" -exec "$GOIMPORTS_BIN" -w {} \;

log "Copying custom templated files..."
# $(cat) strips the template's trailing newlines and printf '%s\n' adds exactly
# one back - matching the old flow's `echo "$template" > file` byte-for-byte.
for tmpl in client_ext.go model_entity_array_ext.go api_hal_ext.go model_paged_cursor_ext.go api_utils_pagination_ext.go; do
    printf '%s\n' "$(sed "s/PACKAGENAME/$MODULE/g" "$SCRIPT_DIR/templates/$tmpl.tmpl")" > "$STAGE_DIR/$tmpl"
done
printf '%s\n' "$(sed "s/PACKAGENAME/$MODULE/g" "$SCRIPT_DIR/templates/PagedCursor.md.tmpl")" > "$STAGE_DIR/docs/PagedCursor.md"
printf '%s\n' "$(sed "s/PACKAGENAME/$MODULE/g" "$SCRIPT_DIR/templates/EntityArrayPagedIterator.md.tmpl")" > "$STAGE_DIR/docs/EntityArrayPagedIterator.md"

log "Applying common and module postprocessing..."
(
    cd "$STAGE_DIR" || exit 1
    GOWORK=off go run "$SCRIPT_DIR/generate-replace-regex.go" .
    GOWORK=off go run "$SCRIPT_DIR/generate-replace-regex.go" ./docs
    GOWORK=off go run "$MODULE_DIR/generate/postprocessing/generate-replace-regex.go" .
    GOWORK=off go run "$MODULE_DIR/generate/postprocessing/generate-replace-regex.go" ./docs
)

# Restore committed state for every ignore-listed path. The generator honours
# the ignore list (pre-copied above) by skipping those files, but FILES still
# records them and hand-maintained files like client.go must exist for the
# build gate. Must run AFTER the regex passes - restoring earlier would let
# the regexes re-wrap already-processed code.
ignore_file="$MODULE_DIR/.openapi-generator-ignore"
while IFS= read -r pattern || [ -n "$pattern" ]; do
    [ -n "$pattern" ] || continue
    case "$pattern" in \#*) continue ;; esac
    pattern="${pattern%$'\r'}"          # tolerate CRLF
    [ -n "${pattern#/}" ] || continue   # "/" alone would clobber everything
    case "$pattern" in
        test/*|generate/*) continue ;;  # handled by sync PRESERVE_PATTERN
    esac
    if [[ "$pattern" == */* ]]; then
        # Directory glob such as test/*: restore every committed file under it.
        dir="${pattern%%/*}"
        if [ -d "$MODULE_DIR/$dir" ]; then
            ( cd "$MODULE_DIR/$dir" && find . -type f | sed 's|^\./||' ) | while IFS= read -r f; do
                mkdir -p "$STAGE_DIR/$dir/$(dirname "$f")"
                cp "$MODULE_DIR/$dir/$f" "$STAGE_DIR/$dir/$f"
            done
        fi
    elif [ -f "$MODULE_DIR/$pattern" ]; then
        cp "$MODULE_DIR/$pattern" "$STAGE_DIR/$pattern"
    fi
done < "$ignore_file"

# --- validate the stage --------------------------------------------------------

log "Validating generated output..."

leaks="$(grep -rlE 'PACKAGENAME|GIT_USER_ID|GIT_REPO_ID' "$STAGE_DIR" --include='*.go' --include='*.md' --include='*.yaml' --include='*.yml' || true)"
if [ -n "$leaks" ]; then
    echo "$leaks" | sed "s|^$STAGE_DIR/|    leak: |"
    die "Generator placeholders leaked into output (files above). Post-processing did not run correctly."
fi

# Case-only filename collisions cannot be detected from the stage itself on a
# case-insensitive filesystem (macOS collapses them), so look for two distinct
# casings of the same basename in the generator's case-sensitive FILES manifest.
collision_report="$(
    awk -F/ '{print $NF}' "$STAGE_DIR/.openapi-generator/FILES" 2>/dev/null \
        | sort -f | uniq -di -f0 | sort -u
)"
if [ -n "$collision_report" ]; then
    echo "$collision_report" | sed 's/^/    collision: /'
    die "Case-only filename collisions detected (above): the generator emits names that differ only by case. Rename the tracked file with git mv to the generator's casing, then re-run."
fi

bad_packages="$(find "$STAGE_DIR" -maxdepth 1 -name '*.go' -exec grep -L "^package $MODULE\$" {} + || true)"
if [ -n "$bad_packages" ]; then
    echo "$bad_packages" | sed "s|^$STAGE_DIR/|    wrong package: |"
    die "Generated files with wrong package name (files above)."
fi

build_log="$(mktemp -t genmod-build.XXXXXX)"
TMP_FILES="$TMP_FILES $build_log"
if ! ( cd "$STAGE_DIR" && GOWORK=off go build ./... ) >"$build_log" 2>&1; then
    sed 's/^/    /' "$build_log" >&2
    die "Generated code failed to build (output above). Re-run with KEEP_STAGING=1 to inspect .generate/$MODULE."
fi

gen_version="$(cat "$STAGE_DIR/.openapi-generator/VERSION" 2>/dev/null || echo missing)"
pinned_version="${OPENAPI_GENERATOR_IMAGE##*:}"
pinned_version="${pinned_version#v}"
[ "$gen_version" = "$pinned_version" ] || die "Generator version mismatch: stage has $gen_version, expected $pinned_version"

# --- determine what would change ----------------------------------------------

(cd "$STAGE_DIR" && find . -type f | sed 's|^\./||') | while IFS= read -r rel; do
    [ -n "$rel" ] || continue
    case "$rel" in
        $PRESERVE_PATTERN) continue ;;
    esac
    if [ ! -f "$MODULE_DIR/$rel" ]; then
        echo "$rel" >> "$(echo "$TMP_FILES" | cut -d' ' -f1)"
    elif ! cmp -s "$STAGE_DIR/$rel" "$MODULE_DIR/$rel"; then
        echo "$rel" >> "$(echo "$TMP_FILES" | cut -d' ' -f2)"
    fi
done
(cd "$STAGE_DIR" && find . -type f | sed 's|^\./||') | sort > "$(echo "$TMP_FILES" | cut -d' ' -f3)"

ADDED_LIST="$(echo "$TMP_FILES" | cut -d' ' -f1)"
MODIFIED_LIST="$(echo "$TMP_FILES" | cut -d' ' -f2)"
STAGE_FILES="$(echo "$TMP_FILES" | cut -d' ' -f3)"

# Stale files: previously generated (old FILES list or tracked) but absent from
# the new output. Report only — deleting exported types breaks SDK consumers.
stale_files="$(
    {
        grep -E '^(api_.*\.go|model_.*\.go|docs/.*\.md)$' "$MODULE_DIR/.openapi-generator/FILES" 2>/dev/null || true
        git -C "$MODULE_DIR" ls-files -- 'api_*.go' 'model_*.go' 'docs/*.md' || true
    } | sort -u | comm -23 - "$STAGE_FILES" | grep -vE "^($NON_GENERATED_FILES)$" || true
)"

added_count="$(wc -l < "$ADDED_LIST" | tr -d ' ')"
modified_count="$(wc -l < "$MODIFIED_LIST" | tr -d ' ')"
would_change=0
if [ "$added_count" != "0" ] || [ "$modified_count" != "0" ]; then
    would_change=1
fi

# --- reporting -----------------------------------------------------------------

if [ "$would_change" = "1" ]; then
    while IFS= read -r rel; do [ -n "$rel" ] && echo "  A $rel"; done < "$ADDED_LIST"
    while IFS= read -r rel; do [ -n "$rel" ] && echo "  M $rel"; done < "$MODIFIED_LIST"
fi
total_stage="$(wc -l < "$STAGE_FILES" | tr -d ' ')"
echo "  $((total_stage - added_count - modified_count)) file(s) unchanged"

if [ -n "$stale_files" ]; then
    warn "stale generated files that the generator no longer produces (report only, not deleted):"
    echo "$stale_files" | sed 's/^/    stale: /'
fi

if [ "$would_change" = "1" ]; then
    committed_version="$(git -C "$REPO_ROOT" show "HEAD:$MODULE/.version" 2>/dev/null || echo "$VERSION")"
    if [ "$VERSION" = "$committed_version" ]; then
        warn "generated output changed but $MODULE/.version is unchanged ($VERSION). Bump .version before generating."
    fi
fi

if [ "$DRY_RUN" = "1" ]; then
    if [ "$would_change" = "1" ]; then
        log "DRY RUN: $MODULE would be updated ($added_count added, $modified_count modified). No files written."
        if [ "$CHECK" = "1" ]; then
            exit 1
        fi
        exit 0
    fi
    log "DRY RUN: $MODULE is up to date."
    exit 0
fi

if [ "$would_change" = "0" ]; then
    log "$MODULE is already up to date. Nothing to do."
    exit 0
fi

# --- sync ----------------------------------------------------------------------

log "Syncing changes into $MODULE/..."
while IFS= read -r rel; do
    [ -n "$rel" ] || continue
    mkdir -p "$MODULE_DIR/$(dirname "$rel")"
    cp "$STAGE_DIR/$rel" "$MODULE_DIR/$rel"
    echo "  A $rel"
done < "$ADDED_LIST"
while IFS= read -r rel; do
    [ -n "$rel" ] || continue
    cp "$STAGE_DIR/$rel" "$MODULE_DIR/$rel"
    echo "  M $rel"
done < "$MODIFIED_LIST"

log "$MODULE generation complete: $added_count added, $modified_count modified."
