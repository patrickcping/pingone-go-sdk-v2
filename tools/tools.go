//go:build tools
// +build tools

package tools

//go:generate go install github.com/golangci/golangci-lint/cmd/golangci-lint
//go:generate go install github.com/pavius/impi/cmd/impi
//go:generate go install github.com/securego/gosec/v2/cmd/gosec@latest

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/pavius/impi/cmd/impi"
	_ "github.com/securego/gosec/v2/cmd/gosec"
)
