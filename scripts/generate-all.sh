#!/usr/bin/env bash
#
# Regenerate SDK modules whose OpenAPI specification (or generation tooling)
# changed, skipping everything else.
#
# Usage: ./scripts/generate-all.sh
#
# Environment:
#   MODULES="mfa management"  Regenerate only these modules (space-separated).
#   ALL=1                     Regenerate every module.
#   BASE_REF=origin/main      Git ref to diff against for change detection
#                             (default: origin/main, falling back to main).
#   DRY_RUN=1 / CHECK=1       Passed through to generate-module.sh; CHECK
#                             additionally exits non-zero if anything would change.
#
# Change detection: a module is selected when its generate/**, .version or
# .openapi-generator-ignore differs from BASE_REF (or is untracked). Changes to
# scripts/templates/**, generate-module.sh or generate-replace-regex.go select
# all modules. Generated output files do not trigger regeneration.

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

OWNER="${OWNER:-patrickcping}"
REPO="${REPO:-pingone-go-sdk-v2}"

log()  { echo "==> $*"; }
die()  { echo "==> ERROR: $*" >&2; exit 1; }

# --- collect modules -----------------------------------------------------------

modules=()
for spec in "$REPO_ROOT"/*/generate/pingone-*.yml; do
    [ -f "$spec" ] || continue
    dir="$(basename "$(dirname "$(dirname "$spec")")")"
    modules+=("$dir")
done
if [ "${#modules[@]}" -eq 0 ]; then
    die "No modules found (expected <module>/generate/pingone-*.yml under $REPO_ROOT)"
fi

# --- decide which modules to run -----------------------------------------------

select_all=0
requested="${MODULES:-}"
if [ "${ALL:-0}" = "1" ]; then
    select_all=1
fi

if [ "$select_all" = "0" ] && [ -z "$requested" ]; then
    BASE_REF="${BASE_REF:-}"
    if [ -z "$BASE_REF" ]; then
        if git -C "$REPO_ROOT" rev-parse --verify --quiet origin/main >/dev/null; then
            BASE_REF="origin/main"
        else
            BASE_REF="main"
        fi
    fi
    merge_base="$(git -C "$REPO_ROOT" merge-base HEAD "$BASE_REF" 2>/dev/null || echo "")"
    if [ -z "$merge_base" ]; then
        log "Cannot resolve merge base against $BASE_REF; regenerating all modules. Set BASE_REF to override."
        select_all=1
    else
        changed_paths_file="$(mktemp -t genall-changed.XXXXXX)"
        trap 'rm -f "$changed_paths_file"' EXIT
        {
            git -C "$REPO_ROOT" diff --name-only "$merge_base" -- . || true
            git -C "$REPO_ROOT" ls-files --others --exclude-standard || true
        } | sort -u > "$changed_paths_file"

        if grep -qE '^scripts/(templates/|generate-module\.sh$|generate-replace-regex\.go$)' "$changed_paths_file"; then
            log "Generation tooling changed; regenerating all modules."
            select_all=1
        else
            for m in "${modules[@]}"; do
                if grep -qE "^$m/(generate/|\.version$|\.openapi-generator-ignore$)" "$changed_paths_file"; then
                    requested="$requested $m"
                fi
            done
            requested="${requested# }"
        fi
    fi
fi

if [ "$select_all" = "1" ]; then
    selected=("${modules[@]}")
else
    if [ -z "$requested" ]; then
        log "No modules changed. Nothing to generate."
        exit 0
    fi
    read -r -a selected <<< "$requested"
    for m in "${selected[@]}"; do
        found=0
        for known in "${modules[@]}"; do
            [ "$m" = "$known" ] && found=1
        done
        [ "$found" = "1" ] || die "Unknown module '$m' (known: ${modules[*]})"
    done
fi

# --- run -----------------------------------------------------------------------

rc=0
for m in "${selected[@]}"; do
    log "Generating module $m..."
    if ! "$SCRIPT_DIR/generate-module.sh" "$OWNER" "$REPO" "$m"; then
        log "Module $m FAILED."
        rc=1
        break
    fi
done

exit "$rc"
