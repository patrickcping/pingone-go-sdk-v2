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

		// EntityArrayEmbeddedItemsInner model
		{
			fileSelectPattern: "model_entity_array__embedded_items_inner.go",
			pattern:           `(func \(dst \*EntityArrayEmbeddedItemsInner\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *EntityArrayEmbeddedItemsInner) UnmarshalJSON(data []byte) error {

	var err error
	// try to unmarshal JSON data into CredentialType
	err = json.Unmarshal(data, &dst.CredentialType);
	if err == nil {
		jsonCredentialType, _ := json.Marshal(dst.CredentialType)
		if string(jsonCredentialType) == "{}" { // empty struct
			dst.CredentialType = nil
		} else {
			if *dst.CredentialType.CardType != "" {
				return nil // data stored in dst.CredentialType, return on the first match
			} else {
				dst.CredentialType = nil
			}
		}
	} else {
		dst.CredentialType = nil
	}

	// try to unmarshal JSON data into UserCredential
	err = json.Unmarshal(data, &dst.UserCredential);
	if err == nil {
		jsonUserCredential, _ := json.Marshal(dst.UserCredential)
		if string(jsonUserCredential) == "{}" { // empty struct
			dst.UserCredential = nil
		} else {
			if dst.UserCredential.User != nil && dst.UserCredential.User.Id != "" {
				return nil // data stored in dst.UserCredential, return on the first match
			} else {
				dst.UserCredential = nil
			}
		}
	} else {
		dst.UserCredential = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(EntityArrayEmbeddedItemsInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},
	}
)
