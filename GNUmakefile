

OWNER=patrickcping
REPO=pingone-go-sdk-v2

default: build

tools:
	go generate -tags tools tools/tools.go

build: fmtcheck
	@./scripts/build-all.sh

lint:
	@./scripts/lint-all.sh

generate:
	@./scripts/generate-all.sh $(OWNER) $(REPO)
	
.PHONY: tools build codecheck lint codegen fmtcheck