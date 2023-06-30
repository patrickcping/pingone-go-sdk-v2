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

		// MFAPushCredentialRequest model
		{
			fileSelectPattern: "model_mfa_push_credential_request.go",
			pattern:           `(func \(dst \*MFAPushCredentialRequest\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *MFAPushCredentialRequest) UnmarshalJSON(data []byte) error {

	var common MFAPushCredential

	if err := json.Unmarshal(data, &common); err != nil {
		return err
	}

	dst.MFAPushCredentialAPNS = nil
	dst.MFAPushCredentialFCM = nil
	dst.MFAPushCredentialFCMHTTPV1 = nil
	dst.MFAPushCredentialHMS = nil

	switch common.GetType() {
	case ENUMMFAPUSHCREDENTIALATTRTYPE_APNS:
		if err := json.Unmarshal(data, &dst.MFAPushCredentialAPNS); err != nil {
			return err
		}
	case ENUMMFAPUSHCREDENTIALATTRTYPE_FCM:
		if err := json.Unmarshal(data, &dst.MFAPushCredentialFCM); err != nil {
			return err
		}
	case ENUMMFAPUSHCREDENTIALATTRTYPE_HMS:
		if err := json.Unmarshal(data, &dst.MFAPushCredentialHMS); err != nil {
			return err
		}
	case ENUMMFAPUSHCREDENTIALATTRTYPE_FCM_HTTP_V1:
		if err := json.Unmarshal(data, &dst.MFAPushCredentialFCMHTTPV1); err != nil {
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(MFAPushCredentialRequest)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},
	}
)
