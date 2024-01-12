TEST?=$$(go list ./...)
OWNER=patrickcping
REPO=pingone-go-sdk-v2
MODULE=agreementmanagement

default: build

build: fmtcheck
	@go mod tidy
	@go mod vendor
	@go build ./...

test: fmtcheck
	go test $(TEST) $(TESTARGS) -timeout=5m

testacc: fmtcheck
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

vet:
	@echo "==> Running go vet..."
	@go vet ./... ; if [ $$? -ne 0 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

depscheck:
	@echo "==> Checking source code with go mod tidy..."
	@go mod tidy
	@git diff --exit-code -- go.mod go.sum || \
		(echo; echo "Unexpected difference in go.mod/go.sum files. Run 'go mod tidy' command or revert any go.mod/go.sum changes and commit."; exit 1)

lint: golangci-lint

golangci-lint:
	@echo "==> Checking source code with golangci-lint..."
	@golangci-lint run ./...

gosec:
	@gosec -exclude-generated ./...

generate:
	@echo "==> Running generate for $(MODULE)..."
	@./../scripts/generate-module.sh $(OWNER) $(REPO) $(MODULE)

devcheck: build vet lint gosec test testacc

.PHONY: build test testacc vet fmtcheck depscheck lint gosec golangci-lint generate