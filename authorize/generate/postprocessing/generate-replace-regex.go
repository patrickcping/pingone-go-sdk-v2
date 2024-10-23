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

		// AuthorizeEditorDataProcessorDTO model
		{
			fileSelectPattern: "model_authorize_editor_data_processor_dto.go",
			pattern:           `(func \(dst \*AuthorizeEditorDataProcessorDTO\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *AuthorizeEditorDataProcessorDTO) UnmarshalJSON(data []byte) error {

	var common AuthorizeEditorDataProcessorDTOCommon

	if err := json.Unmarshal(data, &common); err != nil { // simple model
		return err
	}

	dst.AuthorizeEditorDataProcessorsChainProcessorDTO = nil
	dst.AuthorizeEditorDataProcessorsCollectionFilterProcessorDTO = nil
	dst.AuthorizeEditorDataProcessorsCollectionTransformProcessorDTO = nil
	dst.AuthorizeEditorDataProcessorsJsonPathProcessorDTO = nil
	dst.AuthorizeEditorDataProcessorsReferenceProcessorDTO = nil
	dst.AuthorizeEditorDataProcessorsSpelProcessorDTO = nil
	dst.AuthorizeEditorDataProcessorsXPathProcessorDTO = nil

	switch common.GetType() {
	case ENUMAUTHORIZEEDITORDATAPROCESSORDTOTYPE_CHAIN:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataProcessorsChainProcessorDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATAPROCESSORDTOTYPE_COLLECTION_FILTER:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataProcessorsCollectionFilterProcessorDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATAPROCESSORDTOTYPE_COLLECTION_TRANSFORM:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataProcessorsCollectionTransformProcessorDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATAPROCESSORDTOTYPE_JSON_PATH:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataProcessorsJsonPathProcessorDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATAPROCESSORDTOTYPE_REFERENCE:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataProcessorsReferenceProcessorDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATAPROCESSORDTOTYPE_SPEL:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataProcessorsSpelProcessorDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATAPROCESSORDTOTYPE_XPATH:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataProcessorsXPathProcessorDTO); err != nil { // simple model
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(AuthorizeEditorDataProcessorDTO)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},

		// AuthorizeEditorDataConditionDTO model
		{
			fileSelectPattern: "model_authorize_editor_data_condition_dto.go",
			pattern:           `(func \(dst \*AuthorizeEditorDataConditionDTO\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *AuthorizeEditorDataConditionDTO) UnmarshalJSON(data []byte) error {

	var common AuthorizeEditorDataConditionDTOCommon

	if err := json.Unmarshal(data, &common); err != nil { // simple model
		return err
	}

	dst.AuthorizeEditorDataConditionsAndConditionDTO = nil
	dst.AuthorizeEditorDataConditionsComparisonConditionDTO = nil
	dst.AuthorizeEditorDataConditionsEmptyConditionDTO = nil
	dst.AuthorizeEditorDataConditionsNotConditionDTO = nil
	dst.AuthorizeEditorDataConditionsOrConditionDTO = nil
	dst.AuthorizeEditorDataConditionsReferenceConditionDTO = nil

	switch common.GetType() {
	case ENUMAUTHORIZEEDITORDATACONDITIONDTOTYPE_AND:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataConditionsAndConditionDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATACONDITIONDTOTYPE_COMPARISON:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataConditionsComparisonConditionDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATACONDITIONDTOTYPE_EMPTY:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataConditionsEmptyConditionDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATACONDITIONDTOTYPE_NOT:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataConditionsNotConditionDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATACONDITIONDTOTYPE_OR:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataConditionsOrConditionDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATACONDITIONDTOTYPE_REFERENCE:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataConditionsReferenceConditionDTO); err != nil { // simple model
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(AuthorizeEditorDataConditionDTO)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},

		// AuthorizeEditorDataDefinitionsServiceDefinitionDTO model
		{
			fileSelectPattern: "model_authorize_editor_data_definitions_service_definition_dto.go",
			pattern:           `(func \(dst \*AuthorizeEditorDataDefinitionsServiceDefinitionDTO\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) UnmarshalJSON(data []byte) error {

	var common AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon

	if err := json.Unmarshal(data, &common); err != nil { // simple model
		return err
	}

	dst.AuthorizeEditorDataServicesConnectorServiceDefinitionDTO = nil
	dst.AuthorizeEditorDataServicesHttpServiceDefinitionDTO = nil
	dst.AuthorizeEditorDataServicesNoneServiceDefinitionDTO = nil

	switch common.GetServiceType() {
	case ENUMAUTHORIZEEDITORDATADEFINITIONSSERVICEDEFINITIONDTOSERVICETYPE_CONNECTOR:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataServicesConnectorServiceDefinitionDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATADEFINITIONSSERVICEDEFINITIONDTOSERVICETYPE_HTTP:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataServicesHttpServiceDefinitionDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATADEFINITIONSSERVICEDEFINITIONDTOSERVICETYPE_NONE:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataServicesNoneServiceDefinitionDTO); err != nil { // simple model
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(AuthorizeEditorDataDefinitionsServiceDefinitionDTO)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},

		// CreateService201Response model
		{
			fileSelectPattern: "model_create_service_201_response.go",
			pattern:           `(func \(dst \*CreateService201Response\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *CreateService201Response) UnmarshalJSON(data []byte) error {

	var common CreateService201ResponseCommon

	if err := json.Unmarshal(data, &common); err != nil { // simple model
		return err
	}

	dst.AuthorizeEditorDataServicesConnectorServiceDefinitionDTO = nil
	dst.AuthorizeEditorDataServicesHttpServiceDefinitionDTO = nil
	dst.AuthorizeEditorDataServicesNoneServiceDefinitionDTO = nil

	switch common.GetServiceType() {
	case ENUMAUTHORIZEEDITORDATADEFINITIONSSERVICEDEFINITIONDTOSERVICETYPE_CONNECTOR:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataServicesConnectorServiceDefinitionDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATADEFINITIONSSERVICEDEFINITIONDTOSERVICETYPE_HTTP:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataServicesHttpServiceDefinitionDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATADEFINITIONSSERVICEDEFINITIONDTOSERVICETYPE_NONE:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataServicesNoneServiceDefinitionDTO); err != nil { // simple model
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(CreateService201Response)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},
	}
)
