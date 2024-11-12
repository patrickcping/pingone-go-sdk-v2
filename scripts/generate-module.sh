#!/usr/bin/env bash

if [ ! -d "../$3" ]; then
    echo "Module $3 not found. Run this file from the module directory itself. Skipping..."

else
    echo "Module $3 found.  Running generate.."

    # Generate from OpenAPI
    version=$(head -n 1 .version)

    if [[ -f "generate/pingone-$3.yml" ]]; then \
        echo "==> Running codegen-$3..."
        openapi-generator-cli version-manager set 7.0.1
        openapi-generator-cli generate -i generate/pingone-$3.yml -g go --additional-properties=packageName=$3,packageVersion=$version,isGoSubmodule=true,enumClassPrefix=true,apiNameSuffix=Api -o . --git-repo-id $2 --git-user-id $1 --http-user-agent \"pingtools PingOne-GOLANG-SDK-$3/$version\"; \

        echo "==> Copying custom templated files..."
        template=$(cat ../scripts/templates/client_ext.go.tmpl)
        template=${template//PACKAGENAME/$3}
        echo "$template" > "client_ext.go"

        template=$(cat ../scripts/templates/model_entity_array_ext.go.tmpl)
        template=${template//PACKAGENAME/$3}
        echo "$template" > "model_entity_array_ext.go"

        template=$(cat ../scripts/templates/api_hal_ext.go.tmpl)
        template=${template//PACKAGENAME/$3}
        echo "$template" > "api_hal_ext.go"

        template=$(cat ../scripts/templates/model_paged_cursor_ext.go.tmpl)
        template=${template//PACKAGENAME/$3}
        echo "$template" > "model_paged_cursor_ext.go"

        template=$(cat ../scripts/templates/api_utils_pagination_ext.go.tmpl)
        template=${template//PACKAGENAME/$3}
        echo "$template" > "api_utils_pagination_ext.go"

        echo "==> Applying common postprocessing..."
        go run ../scripts/generate-replace-regex.go .

        echo "==> Applying module specific postprocessing..."
        go run generate/postprocessing/generate-replace-regex.go .

        go get -u ./...
        go mod tidy
        go work vendor

    else \
        echo "pingone-$3.yml missing.  Skipping"; \
    fi
fi
