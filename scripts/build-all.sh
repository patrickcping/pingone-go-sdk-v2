#!/usr/bin/env bash

## declare an array variable
declare -a arr=("agreementmanagement" "authorize" "credentials" "davinci" "management" "mfa" "risk" "verify")

## now loop through the above array
for i in "${arr[@]}"
do
   echo "$i ...."
   
    if [ ! -d "$i" ]; then
        echo "Module $i not found. Skipping..."

    else
        echo "Module $i found.  Running checks.."

        pushd $i
            go mod download
            go mod tidy
            go mod vendor
            make build
            make depscheck
        popd

    fi

done

