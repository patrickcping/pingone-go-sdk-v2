#!/usr/bin/env bash

echo "==> Applying version..."
sed -i.backup "s/^var version = \".*\"$/var version = \"$1\"/g" pingone/client.go && rm pingone/client.go.backup
