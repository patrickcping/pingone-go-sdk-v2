package main

import (
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	// Get the target directory from the command line argument
	if len(os.Args) < 2 {
		println("Usage: go run script.go <directory>")
		return
	}
	dir := os.Args[1]

	for _, rule := range replaceRules {
		// Get a list of all files with the given extension in the directory
		files, err := filepath.Glob(filepath.Join(dir, rule.fileSelectPattern))
		if err != nil {
			panic(err)
		}

		// Iterate over the files and apply the regex replacement rules
		for _, file := range files {
			// Read the file contents
			content, err := os.ReadFile(filepath.Clean(file))
			if err != nil {
				panic(err)
			}

			// Apply the regex replacement rule
			re := regexp.MustCompile(rule.pattern)
			content = re.ReplaceAll(content, []byte(rule.repl))

			// Write the updated file contents
			err = os.WriteFile(file, content, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
	}
}

var (
	// Define the full list of regex replacement rules
	replaceRules = []struct {
		fileSelectPattern string
		pattern           string
		repl              string
	}{

		/////////////////////////
		// ALL configuration.go
		/////////////////////////

		{
			fileSelectPattern: "configuration.go",
			pattern:           `"OpenAPI-Generator/([0-9]+\.[0-9]+\.[0-9]+)/go",`,
			repl:              `"PingOne-GOLANG-SDK/$1/go",`,
		},

		/////////////////////////
		// ALL API
		/////////////////////////

		// Add retryability to typed output
		{
			fileSelectPattern: "api_*.go",
			pattern:           `func \(([a-zA-Z0-9\* ]+)\) ([a-zA-Z])([a-zA-Z0-9]+Execute)\(([a-zA-Z0-9 ]*)\) \(([\*a-zA-Z0-9]*), \*http\.Response, error\) {`,
			repl: `func ($1) $2$3($4) ($5, *http.Response, error) {
	obj, response, error := processResponse(
		func() (interface{}, *http.Response, error) {
			return r.ApiService.internal$2$3(r)
		},
	)
	return obj.($5), response, error
}
			
func ($1) internal$2$3($4) ($5, *http.Response, error) {`,
		},

		// Add retryability to non-typed outputs
		{
			fileSelectPattern: "api_*.go",
			pattern:           `func \(([a-zA-Z0-9\* ]+)\) ([a-zA-Z])([a-zA-Z0-9]+Execute)\(([a-zA-Z0-9 ]*)\) \(\*http\.Response, error\) {`,
			repl: `func ($1) $2$3($4) (*http.Response, error) {
	_, response, error := processResponse(
		func() (interface{}, *http.Response, error) {
			resp, err := r.ApiService.internal$2$3(r)
			return nil, resp, err
		},
	)
	return response, error
}
			
func ($1) internal$2$3($4) (*http.Response, error) {`,
		},

		// Handle errors for Github code scanning
		{
			fileSelectPattern: "api_*.go",
			pattern:           `	localVarHTTPResponse\.Body\.Close\(\)`,
			repl:              `	_ = localVarHTTPResponse.Body.Close()`,
		},
	}
)
