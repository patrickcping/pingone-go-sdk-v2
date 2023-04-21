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
		// Add retryability to typed output
		{
			fileSelectPattern: "api_*.go",
			pattern:           `func \(([a-zA-Z\* ]+)\) ([a-zA-Z])([a-zA-Z]+Execute)\(([a-zA-Z ]*)\) \(([\*a-zA-Z]*), \*http\.Response, error\) {`,
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
			pattern:           `func \(([a-zA-Z\* ]+)\) ([a-zA-Z])([a-zA-Z]+Execute)\(([a-zA-Z ]*)\) \(\*http\.Response, error\) {`,
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

		// Handle errors for linters
		// {
		// 	`	localVarHTTPResponse.Body.Close()`,
		// 	`	_ = localVarHTTPResponse.Body.Close()`,
		// },

		// Password policy model
		{
			fileSelectPattern: "model_password_policy_min_characters.go",
			pattern:           `______`,
			repl:              `SpecialChar`,
		},

		{
			fileSelectPattern: "model_password_policy_min_characters.go",
			pattern:           `json:"~!@\#\$%\^&amp;\*\(\)-_&\#x3D;\+\[]\{}\|;:,\.&lt;&gt;/\?,omitempty"`,
			repl:              `json:"specialchar,omitempty"`,
		},

		{
			fileSelectPattern: "model_password_policy_min_characters.go",
			pattern:           `toSerialize\["~!@\#\$%\^&amp;\*\(\)-_&\#x3D;\+\[]\{}\|;:,\.&lt;&gt;/\?"]\ =\ o\.SpecialChar`,
			repl:              `toSerialize["~!@#$%^&*()-_=+[]{}|;:,.<>/?"] = o.SpecialChar`,
		},

		// Certificate model
		{
			fileSelectPattern: "model_certificate.go",
			pattern:           `int64`,
			repl:              `big.Int`,
		},
		{
			fileSelectPattern: "model_certificate.go",
			pattern:           `import \(`,
			repl: `import (
	"math/big"`,
		},

		// RiskPredictor model
		{
			fileSelectPattern: "model_risk_predictor.go",
			pattern:           `(func \(dst \*RiskPredictor\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *RiskPredictor) UnmarshalJSON(data []byte) error {

	var common RiskPredictorCommon

	if err := json.Unmarshal(data, &common); err != nil { // simple model
		return err
	}

	dst.RiskPredictorAnonymousNetwork = nil
	dst.RiskPredictorComposite = nil
	dst.RiskPredictorCustom = nil
	dst.RiskPredictorGeovelocity = nil
	dst.RiskPredictorIPReputation = nil
	dst.RiskPredictorNewDevice = nil
	dst.RiskPredictorUEBA = nil
	dst.RiskPredictorUserLocationAnomaly = nil
	dst.RiskPredictorVelocity = nil

	switch common.GetType() {
	case ENUMPREDICTORTYPE_ANONYMOUS_NETWORK:
		if err := json.Unmarshal(data, &dst.RiskPredictorAnonymousNetwork); err != nil { // simple model
			return err
		}
	case ENUMPREDICTORTYPE_COMPOSITE:
		if err := json.Unmarshal(data, &dst.RiskPredictorComposite); err != nil { // simple model
			return err
		}
	case ENUMPREDICTORTYPE_MAP:
		if err := json.Unmarshal(data, &dst.RiskPredictorCustom); err != nil { // simple model
			return err
		}
	case ENUMPREDICTORTYPE_GEO_VELOCITY:
		if err := json.Unmarshal(data, &dst.RiskPredictorGeovelocity); err != nil { // simple model
			return err
		}
	case ENUMPREDICTORTYPE_IP_REPUTATION:
		if err := json.Unmarshal(data, &dst.RiskPredictorIPReputation); err != nil { // simple model
			return err
		}
	case ENUMPREDICTORTYPE_NEW_DEVICE:
		if err := json.Unmarshal(data, &dst.RiskPredictorNewDevice); err != nil { // simple model
			return err
		}
	case ENUMPREDICTORTYPE_USER_RISK_BEHAVIOR:
		if err := json.Unmarshal(data, &dst.RiskPredictorUEBA); err != nil { // simple model
			return err
		}
	case ENUMPREDICTORTYPE_USER_LOCATION_ANOMALY:
		if err := json.Unmarshal(data, &dst.RiskPredictorUserLocationAnomaly); err != nil { // simple model
			return err
		}
	case ENUMPREDICTORTYPE_VELOCITY:
		if err := json.Unmarshal(data, &dst.RiskPredictorVelocity); err != nil { // simple model
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(RiskPredictor)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},
	}
)
