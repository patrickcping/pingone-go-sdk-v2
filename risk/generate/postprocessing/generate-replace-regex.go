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
	}
)
