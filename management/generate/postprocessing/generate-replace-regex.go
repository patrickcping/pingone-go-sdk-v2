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

	switch common.GetProvider() {
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
	}
)
