#!/usr/bin/env bash

## declare an array variable
declare -a arr=("agreementmanagement" "authorize" "credentials" "davinci" "management" "mfa" "risk" "verify")
##declare -a arr=("credentials")
## now loop through the above array
for i in "${arr[@]}"
do
   echo "$i ...."
   
    if [ ! -d "$i" ]; then
        echo "Module $i not found. Skipping..."

    else
        echo "Module $i found.  Running generate.."

        # Generate from OpenAPI
        pushd $i
            echo "==> Running codegen-$i..."

            version=$(head -n 1 .version)

            if [[ -f "../pingone-$i.yml" ]]; then \
                openapi-generator generate -i ../pingone-$i.yml -g go --additional-properties=packageName=$i,packageVersion=$version,isGoSubmodule=true,enumClassPrefix=true -o . --git-repo-id $2 --git-user-id $1; \
                go get -u ./...
                go mod tidy
                go mod vendor

                # Generate the retry code
                template=$(cat ../scripts/client_ext.go.tmpl)
                template=${template//PACKAGENAME/$i}
                echo "$template" > "client_ext.go"

                go run ../scripts/generate-replace-regex.go .
                
            else \
                echo "pingone-$i.yml missing.  Skipping"; \
            fi
        popd

    fi

done

