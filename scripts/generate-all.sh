#!/usr/bin/env bash

## declare an array variable
declare -a arr=("authorize" "credentials" "davinci" "fraud" "management" "mfa" "risk" "verify")

## now loop through the above array
for i in "${arr[@]}"
do
   echo "$i ...."
   
    if [ ! -d "$i" ]; then
        echo "Module $i not found. Skipping..."

    else
        echo "Module $i found.  Running generate.."

        pushd $i
            echo "==> Running codegen-$i..."

            version=$(head -n 1 .version)

            if [[ -f "../pingone-$i.yml" ]]; then \
                openapi-generator generate -i ../pingone-$i.yml -g go --additional-properties=packageName=$i,packageVersion=$version,isGoSubmodule=true -o . --git-repo-id $2 --git-user-id $1; \
                go mod tidy
                go mod vendor
            else \
                echo "pingone-$i.yml missing.  Skipping"; \
            fi
        popd

    fi

done

