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
        echo "Module $i found.  Running checks.."

        pushd $i
            make lint
        popd

    fi

done

