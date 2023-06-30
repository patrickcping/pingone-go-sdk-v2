#!/usr/bin/env bash

## declare an array variable
declare -a arr=("agreementmanagement" "authorize" "credentials" "davinci" "management" "mfa" "risk" "verify")

## now loop through the above array
for i in "${arr[@]}"
do
    if [ ! -d "$i" ]; then
        echo "Module $i not found. Skipping..."

    else
        pushd $i
            make generate
        popd
    fi
done
