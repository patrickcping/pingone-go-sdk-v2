#!/usr/bin/env bash

if [ ! -d "../$3" ]; then
    echo "Module $3 not found. Run this file from the module directory itself. Skipping..."

else
    echo "Module $3 found.  Running generate.."

    # Generate from OpenAPI
    version=$(head -n 1 .version)

    if [[ -f "generate/pingone-$3.yml" ]]; then \
        echo "==> Running codegen-$3..."
        openapi-generator generate -i generate/pingone-$3.yml -g go --additional-properties=packageName=$3,packageVersion=$version,isGoSubmodule=true,enumClassPrefix=true,apiNameSuffix=Api -o . --git-repo-id $2 --git-user-id $1 --http-user-agent "PingOne-GOLANG-SDK/$3/$version/go"; \
        go get -u ./...
        go mod tidy
        go mod vendor

        echo "==> Copying custom templated files..."
        template=$(cat ../scripts/client_ext.go.tmpl)
        template=${template//PACKAGENAME/$3}
        echo "$template" > "client_ext.go"

        echo "==> Applying common postprocessing..."
        go run ../scripts/generate-replace-regex.go .

        echo "==> Applying module specific postprocessing..."
        go run generate/postprocessing/generate-replace-regex.go .

    else \
        echo "pingone-$3.yml missing.  Skipping"; \
    fi
fi
