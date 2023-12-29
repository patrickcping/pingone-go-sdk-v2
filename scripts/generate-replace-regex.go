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

		{
			fileSelectPattern: "go.mod",
			pattern:           `go 1.18`,
			repl:              `go 1.21`,
		},

		/////////////////////////
		// ALL configuration.go
		/////////////////////////

		{
			fileSelectPattern: "configuration.go",
			pattern:           `\/\/ AddDefaultHeader adds a new HTTP header to the default header in the request`,
			repl: `func (c *Configuration) SetDebug(debug bool) {
	c.Debug = debug
}

func (c *Configuration) SetUserAgent(userAgent string) {
	c.UserAgent = userAgent
}

func (c *Configuration) AppendUserAgent(userAgent string) {
	c.UserAgent += fmt.Sprintf(" %s", userAgent)
}

func (c *Configuration) SetDefaultServerIndex(defaultServerIndex int) {
	c.DefaultServerIndex = defaultServerIndex
}

func (c *Configuration) SetDefaultServerVariableDefaultValue(variable string, value string) error {
	return c.SetServerVariableDefaultValue(c.DefaultServerIndex, variable, value)
}

func (c *Configuration) SetServerVariableDefaultValue(serverIndex int, variable string, value string) error {
	if serverIndex >= 0 && serverIndex < len(c.Servers) {
		if entry, ok := c.Servers[serverIndex].Variables[variable]; ok {
			entry.DefaultValue = value
			c.Servers[serverIndex].Variables[variable] = entry
			return nil
		} else {
			return fmt.Errorf("variable %v not defined in server %v", variable, serverIndex)
		}
	} else {
		return fmt.Errorf("server index %v out of range %v", serverIndex, len(c.Servers)-1)
	}
}

// AddDefaultHeader adds a new HTTP header to the default header in the request`,
		},

		{
			fileSelectPattern: "configuration.go",
			pattern:           `Debug            bool              ` + "`" + `json:"debug,omitempty"` + "`" + `\n	Servers          ServerConfigurations\n`,
			repl: `Debug            bool              ` + "`" + `json:"debug,omitempty"` + "`" + `
	DefaultServerIndex int             ` + "`" + `json:"defaultServerIndex,omitempty"` + "`" + `
	ProxyURL         *string           ` + "`" + `json:"proxyURL,omitempty"` + "`" + `
	Servers          ServerConfigurations
`,
		},

		{
			fileSelectPattern: "configuration.go",
			pattern:           `Debug:            false,\n		Servers:          ServerConfigurations{\n`,
			repl: `Debug:            false,
		DefaultServerIndex: 0,
		Servers:          ServerConfigurations{
`,
		},

		{
			fileSelectPattern: "configuration.go",
			pattern:           `return 0`,
			repl:              `return defaultServerIndex`,
		},

		{
			fileSelectPattern: "configuration.go",
			pattern:           `func getServerIndex\(ctx context.Context\) \(int, error\)`,
			repl:              `func getServerIndex(ctx context.Context, defaultServerIndex int) (int, error)`,
		},

		{
			fileSelectPattern: "configuration.go",
			pattern:           `func getServerOperationIndex\(ctx context.Context, endpoint string\) \(int, error\)`,
			repl:              `func getServerOperationIndex(ctx context.Context, endpoint string, defaultServerIndex int) (int, error)`,
		},

		{
			fileSelectPattern: "configuration.go",
			pattern:           `return getServerIndex\(ctx\)`,
			repl:              `return getServerIndex(ctx, defaultServerIndex)`,
		},

		{
			fileSelectPattern: "configuration.go",
			pattern:           `return sc.URL\(0,`,
			repl:              `return sc.URL(c.DefaultServerIndex,`,
		},

		{
			fileSelectPattern: "configuration.go",
			pattern:           `getServerOperationIndex\(ctx, endpoint\)`,
			repl:              `getServerOperationIndex(ctx, endpoint, c.DefaultServerIndex)`,
		},

		/////////////////////////
		// ALL API
		/////////////////////////

		// Add retryability to typed output
		{
			fileSelectPattern: "api_*.go",
			pattern:           `func \(([a-zA-Z0-9\* ]+)\) ([a-zA-Z])([a-zA-Z0-9]+Execute)\(([a-zA-Z0-9 ]*)\) \(([\*a-zA-Z0-9]*), \*http\.Response, error\) {`,
			repl: `func ($1) $2$3($4) ($5, *http.Response, error) {
	var (
		err                  error
		response             *http.Response
		localVarReturnValue  $5
	)
	
	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internal$2$3(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}
			
func ($1) internal$2$3($4) ($5, *http.Response, error) {`,
		},

		// Add retryability to non-typed outputs
		{
			fileSelectPattern: "api_*.go",
			pattern:           `func \(([a-zA-Z0-9\* ]+)\) ([a-zA-Z])([a-zA-Z0-9]+Execute)\(([a-zA-Z0-9 ]*)\) \(\*http\.Response, error\) {`,
			repl: `func ($1) $2$3($4) (*http.Response, error) {
	var (
		err      error
		response *http.Response
	)
	
	response, err = processResponse(
		func() (any, *http.Response, error) {
			resp, err := r.ApiService.internal$2$3(r)
			return nil, resp, err
		},
		nil,
	)
	return response, err
}
			
func ($1) internal$2$3($4) (*http.Response, error) {`,
		},

		// Handle errors for Github code scanning
		{
			fileSelectPattern: "api_*.go",
			pattern:           `	localVarHTTPResponse\.Body\.Close\(\)`,
			repl:              `	_ = localVarHTTPResponse.Body.Close()`,
		},

		/////////////////////////
		// ALL ENUM Models
		/////////////////////////
		// Suppress enum unmarshalling errors
		{
			fileSelectPattern: "model_enum_*.go",
			pattern:           `return fmt\.Errorf\("%\+v is not a valid (Enum[a-zA-Z]+)", value\)`,
			repl: `*v = $1(fmt.Sprintf("%s", "UNKNOWN"))
	return nil`,
		},
	}
)
