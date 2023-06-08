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

		/////////////////////////
		// ALL configuration.go
		/////////////////////////

		{
			fileSelectPattern: "configuration.go",
			pattern:           `"OpenAPI-Generator/([0-9]+\.[0-9]+\.[0-9]+)/go",`,
			repl:              `"PingOne-GOLANG-SDK/$1/go",`,
		},

		/////////////////////////
		// ALL API
		/////////////////////////

		// Add retryability to typed output
		{
			fileSelectPattern: "api_*.go",
			pattern:           `func \(([a-zA-Z0-9\* ]+)\) ([a-zA-Z])([a-zA-Z0-9]+Execute)\(([a-zA-Z0-9 ]*)\) \(([\*a-zA-Z0-9]*), \*http\.Response, error\) {`,
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
			pattern:           `func \(([a-zA-Z0-9\* ]+)\) ([a-zA-Z])([a-zA-Z0-9]+Execute)\(([a-zA-Z0-9 ]*)\) \(\*http\.Response, error\) {`,
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

		// Handle errors for Github code scanning
		{
			fileSelectPattern: "api_*.go",
			pattern:           `	localVarHTTPResponse\.Body\.Close\(\)`,
			repl:              `	_ = localVarHTTPResponse.Body.Close()`,
		},

		/////////////////////////
		// Management: Password policy
		/////////////////////////

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

		/////////////////////////
		// Management: Certificate
		/////////////////////////

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

		/////////////////////////
		// Risk: Risk Predictor
		/////////////////////////

		// RiskPredictor model
		{
			fileSelectPattern: "model_risk_predictor.go",
			pattern:           `(func \(dst \*RiskPredictor\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *RiskPredictor) UnmarshalJSON(data []byte) error {

	var common RiskPredictorCommon

	if err := json.Unmarshal(data, &common); err != nil {
		return err
	}

	dst.RiskPredictorAnonymousNetwork = nil
	dst.RiskPredictorComposite = nil
	dst.RiskPredictorCustom = nil
	dst.RiskPredictorGeovelocity = nil
	dst.RiskPredictorIPReputation = nil
	dst.RiskPredictorDevice = nil
	dst.RiskPredictorUserRiskBehavior = nil
	dst.RiskPredictorUserLocationAnomaly = nil
	dst.RiskPredictorVelocity = nil

	switch common.GetType() {
	case ENUMPREDICTORTYPE_ANONYMOUS_NETWORK:
		if err := json.Unmarshal(data, &dst.RiskPredictorAnonymousNetwork); err != nil {
			return err
		}
	case ENUMPREDICTORTYPE_COMPOSITE:
		if err := json.Unmarshal(data, &dst.RiskPredictorComposite); err != nil {
			return err
		}
	case ENUMPREDICTORTYPE_MAP:
		if err := json.Unmarshal(data, &dst.RiskPredictorCustom); err != nil {
			return err
		}
	case ENUMPREDICTORTYPE_GEO_VELOCITY:
		if err := json.Unmarshal(data, &dst.RiskPredictorGeovelocity); err != nil {
			return err
		}
	case ENUMPREDICTORTYPE_IP_REPUTATION:
		if err := json.Unmarshal(data, &dst.RiskPredictorIPReputation); err != nil {
			return err
		}
	case ENUMPREDICTORTYPE_DEVICE:
		if err := json.Unmarshal(data, &dst.RiskPredictorDevice); err != nil {
			return err
		}
	case ENUMPREDICTORTYPE_USER_RISK_BEHAVIOR:
		if err := json.Unmarshal(data, &dst.RiskPredictorUserRiskBehavior); err != nil {
			return err
		}
	case ENUMPREDICTORTYPE_USER_LOCATION_ANOMALY:
		if err := json.Unmarshal(data, &dst.RiskPredictorUserLocationAnomaly); err != nil {
			return err
		}
	case ENUMPREDICTORTYPE_VELOCITY:
		if err := json.Unmarshal(data, &dst.RiskPredictorVelocity); err != nil {
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(RiskPredictor)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},

		// RiskPredictorCompositeCondition model
		{
			fileSelectPattern: "model_risk_predictor_composite_condition.go",
			pattern:           `(func \(dst \*RiskPredictorCompositeCondition\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *RiskPredictorCompositeCondition) UnmarshalJSON(data []byte) error {

	match := 0
	var common map[string]interface{}

	if err := json.Unmarshal(data, &common); err != nil {
		return err
	}

	dst.RiskPredictorCompositeAnd = nil
	dst.RiskPredictorCompositeConditionOneOf = nil
	dst.RiskPredictorCompositeConditionOneOf1 = nil
	dst.RiskPredictorCompositeNot = nil
	dst.RiskPredictorCompositeOr = nil

	if common["and"] != nil {
		if err := json.Unmarshal(data, &dst.RiskPredictorCompositeAnd); err != nil {
			return err
		}
		match++
	}

	if common["or"] != nil {
		if err := json.Unmarshal(data, &dst.RiskPredictorCompositeOr); err != nil {
			return err
		}
		match++
	}

	if common["not"] != nil {
		if err := json.Unmarshal(data, &dst.RiskPredictorCompositeNot); err != nil {
			return err
		}
		match++
	}

	if v, ok := common["type"].(string); ok && v == string(ENUMPREDICTORCOMPOSITECONDITIONTYPE_STRING_LIST) {
		if err := json.Unmarshal(data, &dst.RiskPredictorCompositeConditionOneOf); err != nil {
			return err
		}
		match++
	}

	if v, ok := common["type"].(string); ok && v == string(ENUMPREDICTORCOMPOSITECONDITIONTYPE_VALUE_COMPARISON) {
		if err := json.Unmarshal(data, &dst.RiskPredictorCompositeConditionOneOf1); err != nil {
			return err
		}
		match++
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.RiskPredictorCompositeAnd = nil
		dst.RiskPredictorCompositeConditionOneOf = nil
		dst.RiskPredictorCompositeConditionOneOf1 = nil
		dst.RiskPredictorCompositeNot = nil
		dst.RiskPredictorCompositeOr = nil

		return fmt.Errorf("data matches more than one schema in oneOf(RiskPredictorCompositeCondition)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RiskPredictorCompositeCondition)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},

		// RiskPredictorCompositeConditionBase model
		{
			fileSelectPattern: "model_risk_predictor_composite_condition_base.go",
			pattern:           `(func \(dst \*RiskPredictorCompositeConditionBase\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *RiskPredictorCompositeConditionBase) UnmarshalJSON(data []byte) error {
	
	match := 0
	var common map[string]interface{}

	if err := json.Unmarshal(data, &common); err != nil {
		return err
	}

	dst.RiskPredictorCompositeAnd = nil
	dst.RiskPredictorCompositeNot = nil
	dst.RiskPredictorCompositeOr = nil

	if common["and"] != nil {
		if err := json.Unmarshal(data, &dst.RiskPredictorCompositeAnd); err != nil {
			return err
		}
		match++
	}

	if common["or"] != nil {
		if err := json.Unmarshal(data, &dst.RiskPredictorCompositeOr); err != nil {
			return err
		}
		match++
	}

	if common["not"] != nil {
		if err := json.Unmarshal(data, &dst.RiskPredictorCompositeNot); err != nil {
			return err
		}
		match++
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.RiskPredictorCompositeAnd = nil
		dst.RiskPredictorCompositeNot = nil
		dst.RiskPredictorCompositeOr = nil

		return fmt.Errorf("data matches more than one schema in oneOf(RiskPredictorCompositeConditionBase)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RiskPredictorCompositeConditionBase)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},

		/////////////////////////
		// Credentials: EntityArrayEmbeddedItemsInner
		/////////////////////////

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
			if dst.CredentialType.CardDesignTemplate != "" {
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
			if dst.UserCredential.User.Id != "" {
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

		// Management: FormField model
		{
			fileSelectPattern: "model_form_field.go",
			pattern:           `(func \(dst \*FormField\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *FormField) UnmarshalJSON(data []byte) error {

	var common FormFieldCommon

	if err := json.Unmarshal(data, &common); err != nil {
		return err
	}

	dst.FormFieldCheckbox = nil
	dst.FormFieldCombobox = nil
	dst.FormFieldDivider = nil
	dst.FormFieldDropdown = nil
	dst.FormFieldEmptyField = nil
	dst.FormFieldErrorDisplay = nil
	dst.FormFieldFlowButton = nil
	dst.FormFieldFlowLink = nil
	dst.FormFieldPassword = nil
	dst.FormFieldPasswordVerify = nil
	dst.FormFieldQrCode = nil
	dst.FormFieldRadio = nil
	dst.FormFieldRecaptchaV2 = nil
	dst.FormFieldSlateTextblob = nil
	dst.FormFieldSocialLoginButton = nil
	dst.FormFieldSubmitButton = nil
	dst.FormFieldText = nil

	switch common.GetType() {
	case ENUMFORMFIELDTYPE_TEXT:
		if err := json.Unmarshal(data, &dst.FormFieldText); err != nil {
			return err
		}
	case ENUMFORMFIELDTYPE_PASSWORD:
		if err := json.Unmarshal(data, &dst.FormFieldPassword); err != nil {
			return err
		}
	case ENUMFORMFIELDTYPE_PASSWORD_VERIFY:
		if err := json.Unmarshal(data, &dst.FormFieldPasswordVerify); err != nil {
			return err
		}
	case ENUMFORMFIELDTYPE_RADIO:
		if err := json.Unmarshal(data, &dst.FormFieldRadio); err != nil {
			return err
		}
	case ENUMFORMFIELDTYPE_CHECKBOX:
		if err := json.Unmarshal(data, &dst.FormFieldCheckbox); err != nil {
			return err
		}
	case ENUMFORMFIELDTYPE_DROPDOWN:
		if err := json.Unmarshal(data, &dst.FormFieldDropdown); err != nil {
			return err
		}
	case ENUMFORMFIELDTYPE_COMBOBOX:
		if err := json.Unmarshal(data, &dst.FormFieldCombobox); err != nil {
			return err
		}
	case ENUMFORMFIELDTYPE_DIVIDER:
		if err := json.Unmarshal(data, &dst.FormFieldDivider); err != nil {
			return err
		}
	case ENUMFORMFIELDTYPE_EMPTY_FIELD:
		if err := json.Unmarshal(data, &dst.FormFieldEmptyField); err != nil {
			return err
		}
	case ENUMFORMFIELDTYPE_TEXTBLOB:
		if err := json.Unmarshal(data, &dst.FormFieldSlateTextblob); err != nil {
			return err
		}
	case ENUMFORMFIELDTYPE_SLATE_TEXTBLOB:
		if err := json.Unmarshal(data, &dst.FormFieldSlateTextblob); err != nil {
			return err
		}
	case ENUMFORMFIELDTYPE_SUBMIT_BUTTON:
		if err := json.Unmarshal(data, &dst.FormFieldSubmitButton); err != nil {
			return err
		}
	case ENUMFORMFIELDTYPE_ERROR_DISPLAY:
		if err := json.Unmarshal(data, &dst.FormFieldErrorDisplay); err != nil {
			return err
		}
	case ENUMFORMFIELDTYPE_FLOW_LINK:
		if err := json.Unmarshal(data, &dst.FormFieldFlowLink); err != nil {
			return err
		}
	case ENUMFORMFIELDTYPE_FLOW_BUTTON:
		if err := json.Unmarshal(data, &dst.FormFieldFlowButton); err != nil {
			return err
		}
	case ENUMFORMFIELDTYPE_RECAPTCHA_V2:
		if err := json.Unmarshal(data, &dst.FormFieldRecaptchaV2); err != nil {
			return err
		}
	case ENUMFORMFIELDTYPE_QR_CODE:
		if err := json.Unmarshal(data, &dst.FormFieldQrCode); err != nil {
			return err
		}
	case ENUMFORMFIELDTYPE_SOCIAL_LOGIN_BUTTON:
		if err := json.Unmarshal(data, &dst.FormFieldSocialLoginButton); err != nil {
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(FormField)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},
	}
)
