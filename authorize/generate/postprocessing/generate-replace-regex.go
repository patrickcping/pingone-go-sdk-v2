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
		// EntityArrayEmbeddedPermissionsInner model
		{
			fileSelectPattern: "model_entity_array__embedded_permissions_inner.go",
			pattern:           `(func \(dst \*EntityArrayEmbeddedPermissionsInner\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *EntityArrayEmbeddedPermissionsInner) UnmarshalJSON(data []byte) error {

	var err error
	// try to unmarshal JSON data into Gateway
	err = json.Unmarshal(data, &dst.Gateway)
	if err == nil {
		jsonGateway, _ := json.Marshal(dst.Gateway)
		if string(jsonGateway) == "{}" { // empty struct
			dst.Gateway = nil
		} else {
			switch dst.Gateway.Type {
			case ENUMGATEWAYTYPE_LDAP, ENUMGATEWAYTYPE_RADIUS:
				dst.Gateway = nil
			default:
				return nil // data stored in dst.Gateway, return on the first match
			}
		}
	} else {
		dst.Gateway = nil
	}

	// try to unmarshal JSON data into GatewayTypeLDAP
	err = json.Unmarshal(data, &dst.GatewayTypeLDAP)
	if err == nil {
		jsonGatewayLDAP, _ := json.Marshal(dst.GatewayTypeLDAP)
		if string(jsonGatewayLDAP) == "{}" { // empty struct
			dst.GatewayTypeLDAP = nil
		} else {
			if dst.GatewayTypeLDAP.Type == ENUMGATEWAYTYPE_LDAP {
				return nil // data stored in dst.GatewayLDAP, return on the first match
			} else {
				dst.GatewayTypeLDAP = nil
			}
		}
	} else {
		dst.GatewayTypeLDAP = nil
	}

	// try to unmarshal JSON data into GatewayTypeRADIUS
	err = json.Unmarshal(data, &dst.GatewayTypeRADIUS)
	if err == nil {
		jsonGatewayRADIUS, _ := json.Marshal(dst.GatewayTypeRADIUS)
		if string(jsonGatewayRADIUS) == "{}" { // empty struct
			dst.GatewayTypeRADIUS = nil
		} else {
			if dst.GatewayTypeRADIUS.Type == ENUMGATEWAYTYPE_RADIUS {
				return nil // data stored in dst.GatewayTypeRADIUS, return on the first match
			} else {
				dst.GatewayTypeRADIUS = nil
			}
		}
	} else {
		dst.GatewayTypeRADIUS = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(EntityArrayEmbeddedGatewaysInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},
	}
)
