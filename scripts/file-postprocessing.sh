#!/usr/bin/env bash

gofmt -w -s $1
goimports -w $1