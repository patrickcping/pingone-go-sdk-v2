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
		// EntityArrayEmbeddedPermissionsInner model
		{
			fileSelectPattern: "model_entity_array__embedded_permissions_inner.go",
			pattern:           `(func \(dst \*EntityArrayEmbeddedPermissionsInner\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *EntityArrayEmbeddedPermissionsInner) UnmarshalJSON(data []byte) error {

	var err error
	// try to unmarshal JSON data into ApplicationRolePermission
	err = json.Unmarshal(data, &dst.ApplicationRolePermission)
	if err == nil {
		jsonApplicationRolePermission, _ := json.Marshal(dst.ApplicationRolePermission)
		if string(jsonApplicationRolePermission) == "{}" { // empty struct
			dst.ApplicationRolePermission = nil
		} else {
			if dst.ApplicationRolePermission.Key != nil {
				return nil // data stored in dst.ApplicationRolePermission, return on the first match
			} else {
				dst.ApplicationRolePermission = nil
			}
		}
	} else {
		dst.ApplicationRolePermission = nil
	}

	// try to unmarshal JSON data into ApplicationResourcePermission
	err = json.Unmarshal(data, &dst.ApplicationResourcePermission)
	if err == nil {
		jsonApplicationResourcePermission, _ := json.Marshal(dst.ApplicationResourcePermission)
		if string(jsonApplicationResourcePermission) == "{}" { // empty struct
			dst.ApplicationResourcePermission = nil
		} else {
			if dst.ApplicationResourcePermission.Action != "" { // we expect an resource for this data type
				dst.ApplicationResourcePermission = nil
			} else {
				return nil // data stored in dst.ApplicationResourcePermission, return on the first match
			}
		}
	} else {
		dst.ApplicationResourcePermission = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(EntityArrayEmbeddedPermissionsInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},
	}
)
