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
			repl:              `management`,
		},

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

		// Management: EnvironmentRegion model
		{
			fileSelectPattern: "model_environment_region.go",
			pattern:           `(func \(dst \*EnvironmentRegion\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *EnvironmentRegion) UnmarshalJSON(data []byte) error {

	var err error
	match := false
	// try to unmarshal data into EnumRegionCode
	err = newStrictDecoder(data).Decode(&dst.EnumRegionCode)
	if err == nil {
		jsonEnumRegionCode, _ := json.Marshal(dst.EnumRegionCode)
		if string(jsonEnumRegionCode) == "{}" { // empty struct
			dst.EnumRegionCode = nil
		} else {
			match = true
		}
	} else {
		dst.EnumRegionCode = nil
	}
	
	if !match {
		// try to unmarshal data into String
		err = newStrictDecoder(data).Decode(&dst.String)
		if err == nil {
			jsonString, _ := json.Marshal(dst.String)
			if string(jsonString) == "{}" { // empty struct
				dst.String = nil
			} else {
				match = true
			}
		} else {
			dst.String = nil
		}
	}
			
	if !match { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(EnvironmentRegion)")
	}
			
	return nil
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

	objType := common.GetType()

	if !objType.IsValid() {
		return nil
	}

	switch objType {
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

		// Management: NotificationsSettingsPhoneDeliverySettings model
		{
			fileSelectPattern: "model_notifications_settings_phone_delivery_settings.go",
			pattern:           `(func \(dst \*NotificationsSettingsPhoneDeliverySettings\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *NotificationsSettingsPhoneDeliverySettings) UnmarshalJSON(data []byte) error {

	var common NotificationsSettingsPhoneDeliverySettingsCommon

	if err := json.Unmarshal(data, &common); err != nil {
		return err
	}

	dst.NotificationsSettingsPhoneDeliverySettingsCustom = nil
	dst.NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse = nil

	objProvider := common.GetProvider()

	if !objProvider.IsValid() {
		return nil
	}

	switch objProvider {
	case ENUMNOTIFICATIONSSETTINGSPHONEDELIVERYSETTINGSPROVIDER_TWILIO:
		if err := json.Unmarshal(data, &dst.NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse); err != nil {
			return err
		}
	case ENUMNOTIFICATIONSSETTINGSPHONEDELIVERYSETTINGSPROVIDER_SYNIVERSE:
		if err := json.Unmarshal(data, &dst.NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse); err != nil {
			return err
		}
	case ENUMNOTIFICATIONSSETTINGSPHONEDELIVERYSETTINGSPROVIDER_PROVIDER:
		if err := json.Unmarshal(data, &dst.NotificationsSettingsPhoneDeliverySettingsCustom); err != nil {
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(NotificationsSettingsPhoneDeliverySettings)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},

		// EntityArrayEmbeddedGatewaysInner model
		{
			fileSelectPattern: "model_entity_array__embedded_gateways_inner.go",
			pattern:           `(func \(dst \*EntityArrayEmbeddedGatewaysInner\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *EntityArrayEmbeddedGatewaysInner) UnmarshalJSON(data []byte) error {

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

		// EntityArrayEmbeddedResourcesInner model
		{
			fileSelectPattern: "model_entity_array__embedded_resources_inner.go",
			pattern:           `(func \(dst \*EntityArrayEmbeddedResourcesInner\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *EntityArrayEmbeddedResourcesInner) UnmarshalJSON(data []byte) error {

	var err error
	// try to unmarshal JSON data into Resource
	err = json.Unmarshal(data, &dst.Resource)
	if err == nil {
		jsonResource, _ := json.Marshal(dst.Resource)
		if string(jsonResource) == "{}" { // empty struct
			dst.Resource = nil
		} else {
			if dst.Resource.Type == nil { // we expect an action for this data type
				dst.Resource = nil
			} else {
				return nil // data stored in dst.Resource, return on the first match
			}
		}
	} else {
		dst.Resource = nil
	}

	// try to unmarshal JSON data into ResourceApplicationResource
	err = json.Unmarshal(data, &dst.ResourceApplicationResource)
	if err == nil {
		jsonResourceApplicationResource, _ := json.Marshal(dst.ResourceApplicationResource)
		if string(jsonResourceApplicationResource) == "{}" { // empty struct
			dst.ResourceApplicationResource = nil
		} else {
				return nil // data stored in dst.ResourceApplicationResource, return on the first match
		}
	} else {
		dst.ResourceApplicationResource = nil
	}

	return fmt.Errorf("Data failed to match schemas in oneOf(EntityArrayEmbeddedResourcesInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},

		// EntityArrayEmbeddedRolesInner model
		{
			fileSelectPattern: "model_entity_array__embedded_roles_inner.go",
			pattern:           `(func \(dst \*EntityArrayEmbeddedRolesInner\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *EntityArrayEmbeddedRolesInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into CustomAdminRole
	err = json.Unmarshal(data, &dst.CustomAdminRole)
	if err == nil {
		jsonCustomAdminRole, _ := json.Marshal(dst.CustomAdminRole)
		if string(jsonCustomAdminRole) == "{}" { // empty struct
			dst.CustomAdminRole = nil
		} else {
				return nil // data stored in dst.CustomAdminRole, return on the first match
		}
	} else {
		dst.CustomAdminRole = nil
	}
	// try to unmarshal JSON data into Role
	err = json.Unmarshal(data, &dst.Role)
	if err == nil {
		jsonRole, _ := json.Marshal(dst.Role)
		if string(jsonRole) == "{}" { // empty struct
			dst.Role = nil
		} else {
			if dst.Role.Permissions == nil { // we expect an action for this data type
				dst.Role = nil
			} else {
				return nil // data stored in dst.Role, return on the first match
			}
		}
	} else {
		dst.Role = nil
	}
	// try to unmarshal JSON data into UserApplicationRoleAssignment
	err = json.Unmarshal(data, &dst.UserApplicationRoleAssignment)
	if err == nil {
		jsonUserApplicationRoleAssignment, _ := json.Marshal(dst.UserApplicationRoleAssignment)
		if string(jsonUserApplicationRoleAssignment) == "{}" { // empty struct
			dst.UserApplicationRoleAssignment = nil
		} else {
				return nil // data stored in dst.UserApplicationRoleAssignment, return on the first match
		}
	} else {
		dst.UserApplicationRoleAssignment = nil
	}
	return fmt.Errorf("Data failed to match schemas in oneOf(EntityArrayEmbeddedRolesInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},

		// Management: IntegrationVersion model
		{
			fileSelectPattern: "model_integration_version.go",
			pattern:           `(func \(dst \*IntegrationVersion\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *IntegrationVersion) UnmarshalJSON(data []byte) error {

	var common IntegrationVersionCommon

	if err := json.Unmarshal(data, &common); err != nil {
		return err
	}

	dst.IntegrationVersionIntegrationKit = nil
	dst.IntegrationVersionSAML = nil

	objType := common.GetType()

	if !objType.IsValid() {
		return nil
	}

	switch objType {
	case ENUMINTEGRATIONVERSIONTYPE_PRODUCT_INTEGRATION_KIT:
		if err := json.Unmarshal(data, &dst.IntegrationVersionIntegrationKit); err != nil {
			return err
		}
	case ENUMINTEGRATIONVERSIONTYPE_SAML:
		if err := json.Unmarshal(data, &dst.IntegrationVersionSAML); err != nil {
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(IntegrationVersion)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},
	}
)
