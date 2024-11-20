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
			err = os.WriteFile(file, content, os.ModeAppend)
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

		/////////////////////////
		// ALL API
		/////////////////////////

		// Add paging to EntityArray APIs (Execute function)
		{
			fileSelectPattern: "*Api.md",
			pattern:           `EntityArray`,
			repl:              `EntityArrayPagedIterator`,
		},
		{
			fileSelectPattern: "*Api.md",
			pattern:           `## ([a-zA-Z0-9]+)\n\n> EntityArrayPagedIterator ([a-zA-Z0-9]+)\(([a-zA-Z0-9, ]*)\)\.Execute\(\)\n\n([a-zA-Z0-9, ]*)\n\n`,
			repl: `## $1

$4

### Paged Response (Recommended)

> EntityArrayPagedIterator $2($3).Execute()

#### Example

^^^^^^^^^go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/PACKAGENAME"
)

func main() {
    environmentID := "environmentID_example" // string | 
	// ... other parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
	api := apiClient. // .... API function
    pagedIterator := api.$2(context.Background(), environmentID, /* ... other parameters */).Execute()
	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling ^^^api.$2^^^^^^: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
			break
		}

		// response from ^^^$2^^^: EntityArrayPagedIterator
		fmt.Fprintf(os.Stdout, "Response from ^^^api.$2^^^: %v\n", pageCursor.EntityArray)
	}
}
^^^^^^^^^

### Initial Page Response

> EntityArray $2($3).ExecuteInitialPage()

#`,
		},
		{
			fileSelectPattern: "README.md",
			pattern:           ` - \[P1ErrorDetailsInnerInnerError\]\(docs\/P1ErrorDetailsInnerInnerError\.md\)`,
			repl: ` - [P1ErrorDetailsInnerInnerError](docs/P1ErrorDetailsInnerInnerError.md)
 - [PagedCursor](docs/PagedCursor.md)`,
		},
	}
)
