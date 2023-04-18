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

	// Define the target extension
	ext := ".go"

	// Get a list of all files with the given extension in the directory
	files, err := filepath.Glob(filepath.Join(dir, "*"+ext))
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

		// Apply the regex replacement rules
		for _, rule := range replaceRules {
			re := regexp.MustCompile(rule.pattern)
			content = re.ReplaceAll(content, []byte(rule.repl))
		}

		// Write the updated file contents
		err = os.WriteFile(file, content, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

var (
	// Define the full list of regex replacement rules
	replaceRules = []struct {
		pattern string
		repl    string
	}{
		// Add retryability to typed output
		{
			`func \(([a-zA-Z\* ]+)\) ([a-zA-Z])([a-zA-Z]+Execute)\(([a-zA-Z ]*)\) \(([\*a-zA-Z]*), \*http\.Response, error\) {`,
			`func ($1) $2$3($4) ($5, *http.Response, error) {
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
			`func \(([a-zA-Z\* ]+)\) ([a-zA-Z])([a-zA-Z]+Execute)\(([a-zA-Z ]*)\) \(\*http\.Response, error\) {`,
			`func ($1) $2$3($4) (*http.Response, error) {
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

		// Handle errors for linters
		// {
		// 	`	localVarHTTPResponse.Body.Close()`,
		// 	`	_ = localVarHTTPResponse.Body.Close()`,
		// },

		{
			`______`,
			`SpecialChar`,
		},

		{
			`json:"~!@\#\$%\^&amp;\*\(\)-_&\#x3D;\+\[]\{}\|;:,\.&lt;&gt;/\?,omitempty"`,
			`json:"specialchar,omitempty"`,
		},

		{
			`toSerialize\["~!@\#\$%\^&amp;\*\(\)-_&\#x3D;\+\[]\{}\|;:,\.&lt;&gt;/\?"]\ =\ o\.SpecialChar`,
			`toSerialize["~!@#$%^&*()-_=+[]{}|;:,.<>/?"] = o.SpecialChar`,
		},
	}
)
