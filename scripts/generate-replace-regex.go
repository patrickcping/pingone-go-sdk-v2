package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	// Define the list of regex replacement rules
	replaceRules := []struct {
		pattern string
		repl    string
	}{
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
	}

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
		content, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}

		// Apply the regex replacement rules
		for _, rule := range replaceRules {
			re := regexp.MustCompile(rule.pattern)
			content = re.ReplaceAll(content, []byte(rule.repl))
		}

		// Write the updated file contents
		err = ioutil.WriteFile(file, content, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}
