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

		{
			fileSelectPattern: "*Api.md",
			pattern:           `PACKAGENAME`,
			repl:              `mfa`,
		},

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

	objType := common.GetType()

	if !objType.IsValid() {
		return nil
	}

	switch objType {
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

		// DeviceAuthenticationPolicyPost model
		{
			fileSelectPattern: "model_device_authentication_policy_post.go",
			pattern:           `(func \(dst \*DeviceAuthenticationPolicyPost\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *DeviceAuthenticationPolicyPost) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DeviceAuthenticationPolicy
	err = json.Unmarshal(data, &dst.DeviceAuthenticationPolicy)
	if err == nil {
		if v, ok := dst.DeviceAuthenticationPolicy.GetNameOk(); ok && v != nil && *v != "" {
			match++
		} else {
			dst.DeviceAuthenticationPolicy = nil
		}
	} else {
		dst.DeviceAuthenticationPolicy = nil
	}

	// try to unmarshal data into DeviceAuthenticationPolicyMigrate
	err = json.Unmarshal(data, &dst.DeviceAuthenticationPolicyMigrate)
	if err == nil {
		if v, ok := dst.DeviceAuthenticationPolicyMigrate.GetMigrationDataOk(); ok && len(v) > 0 {
			match++
		} else {
			dst.DeviceAuthenticationPolicyMigrate = nil
		}
	} else {
		dst.DeviceAuthenticationPolicyMigrate = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DeviceAuthenticationPolicy = nil
		dst.DeviceAuthenticationPolicyMigrate = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DeviceAuthenticationPolicyPost)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DeviceAuthenticationPolicyPost)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},

		// DeviceAuthenticationPolicyPostResponse model
		{
			fileSelectPattern: "model_device_authentication_policy_post_response.go",
			pattern:           `(func \(dst \*DeviceAuthenticationPolicyPostResponse\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *DeviceAuthenticationPolicyPostResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DeviceAuthenticationPolicy
	err = json.Unmarshal(data, &dst.DeviceAuthenticationPolicy)
	if err == nil {
		if v, ok := dst.DeviceAuthenticationPolicy.GetNameOk(); ok && v != nil && *v != "" {
			match++
		} else {
			dst.DeviceAuthenticationPolicy = nil
		}
	} else {
		dst.DeviceAuthenticationPolicy = nil
	}

	// try to unmarshal data into EntityArray
	err = json.Unmarshal(data, &dst.EntityArray)
	if err == nil {
		if v, ok := dst.EntityArray.GetEmbeddedOk(); ok && v != nil {
			match++
		} else {
			dst.EntityArray = nil
		}
	} else {
		dst.EntityArray = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DeviceAuthenticationPolicy = nil
		dst.EntityArray = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DeviceAuthenticationPolicyPostResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DeviceAuthenticationPolicyPostResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},
	}
)
