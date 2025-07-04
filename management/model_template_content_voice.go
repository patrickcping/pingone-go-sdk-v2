/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// checks if the TemplateContentVoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateContentVoice{}

// TemplateContentVoice struct for TemplateContentVoice
type TemplateContentVoice struct {
	Links *map[string]LinksHATEOASValue `json:"_links,omitempty"`
	// The template id.
	Id *string `json:"id,omitempty"`
	// The time the resource was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The time the resource was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// Specifies whether the template is a predefined default template.
	Default *bool `json:"default,omitempty"`
	// A valid case-insensitive locale, complying with the ISO-639 language code and ISO-3166 country code standards: Two-character language code, for example, \"en\". Two-character language code followed by a two-character country code, separated by an underscore or dash, for example: \"en_GB\", \"en-GB\". Cannot be changed after it is initially set in `POST /environments/{{envID}}/templates/{{templateName}}/contents`.
	Locale         string                            `json:"locale"`
	DeliveryMethod EnumTemplateContentDeliveryMethod `json:"deliveryMethod"`
	// Holds the unique user-defined name for each content variant that uses the same template + `deliveryMethod` + `locale` combination. This property is case insensitive and has a limit of 100 characters. For more information, see [Creating custom contents](https://apidocs.pingidentity.com/pingone/platform/v1/api/#notifications-templates-creating-custom-contents).
	Variant *string `json:"variant,omitempty"`
	// The voice text.  Limited to 1Kb characters.   The following substitution place holders can be embedded in the message:   * `<pause1sec>`: Pauses the message narration for 1 second.     The pause interval <pause1sec> cannot be modified. To pause the message narration for more than one second, repeat multiple `<pause1sec>` occurrences in succession, according to the desired pause interval duration. For example, `<pause1sec><pause1sec><pause1sec>` pauses the message narration for three seconds.   * `<sayCharValue> .. </sayCharValue>`: Reads the character name of each character of the enclosed string separately.   * `<repeatMessage val=x> .. </repeatMessage>`: Narrates the enclosed text `<val>` number of times.      In the following message example, `${otp}` is assigned the value `\"123456\"`, and `${email}` is assigned the value `\"joe@bxz.com\"`:   ```   Hello <pause1sec> your authentication code is    <sayCharValue>${otp}</​sayCharValue>    <repeatMessage val=2> I repeat your code is    <sayCharValue>${otp}</​sayCharValue>   </​repeatMessage> <pause1sec>    Mail <sayCharValue>${email}</​sayCharValue> for help   ```   The narrated message on the voice call sounds like:   ```   HELLO <1 second silence> YOUR AUTHENTICATION CODE IS    ONE TWO THREE FOUR FIVE SIX    I REPEAT YOUR CODE IS ONE TWO THREE FOUR FIVE SIX    I REPEAT YOUR CODE IS ONE TWO THREE FOUR FIVE SIX <1 second silence>    MAIL JAY OH EE AT BEE EX ZEE DOT SEE OH EM FOR HELP   ```
	Content string `json:"content"`
	// Voice OTP supports vendor-specific voices.   * Supported Twilio voices:     * Man, Woman       Supported locales (default: en):       en, en_GB, es, fr, de     * Alice (Twilio only)       Supported locales (default: en US):       da_DK, de_DE, en_AU, en_CA, en_GB, en_US, ca_ES, es_ES, es_MX, fi_FI, fr_CA, fr_FR, it_IT, ja_JP, ko_KR, nb_NO, nl_NL, pl_PL, pt_BR, pt_PT, ru_RU, sv_SE, zh_CN, zh_HK, zh_TW     * Amazon Polly       cy_GB, ro_RO, is_IS, hi_IN tr_TR   * Supported Syniverse voices:     * Man, Woman       Supported locales:       en_US, en_GB, es_ES, es_US, fr_FR, de_DE, it_IT, en_AU, da_DK, is_IS, nl_NL, pl_PL, pt_BR, pt_PT, ru_RU, ja_JP     * Woman only       Supported locales:       cmn_CN, cy_GB, en_IN, fr_CA, hi_IN, nb_NO, ro_RO, sv_SE, tr_TR, ko_KR, ar
	Voice *string `json:"voice,omitempty"`
}

// NewTemplateContentVoice instantiates a new TemplateContentVoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateContentVoice(locale string, deliveryMethod EnumTemplateContentDeliveryMethod, content string) *TemplateContentVoice {
	this := TemplateContentVoice{}
	this.Locale = locale
	this.DeliveryMethod = deliveryMethod
	this.Content = content
	return &this
}

// NewTemplateContentVoiceWithDefaults instantiates a new TemplateContentVoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateContentVoiceWithDefaults() *TemplateContentVoice {
	this := TemplateContentVoice{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TemplateContentVoice) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContentVoice) GetLinksOk() (*map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TemplateContentVoice) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *TemplateContentVoice) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TemplateContentVoice) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContentVoice) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TemplateContentVoice) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TemplateContentVoice) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TemplateContentVoice) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContentVoice) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TemplateContentVoice) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *TemplateContentVoice) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TemplateContentVoice) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContentVoice) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TemplateContentVoice) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *TemplateContentVoice) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *TemplateContentVoice) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContentVoice) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *TemplateContentVoice) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *TemplateContentVoice) SetDefault(v bool) {
	o.Default = &v
}

// GetLocale returns the Locale field value
func (o *TemplateContentVoice) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *TemplateContentVoice) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *TemplateContentVoice) SetLocale(v string) {
	o.Locale = v
}

// GetDeliveryMethod returns the DeliveryMethod field value
func (o *TemplateContentVoice) GetDeliveryMethod() EnumTemplateContentDeliveryMethod {
	if o == nil {
		var ret EnumTemplateContentDeliveryMethod
		return ret
	}

	return o.DeliveryMethod
}

// GetDeliveryMethodOk returns a tuple with the DeliveryMethod field value
// and a boolean to check if the value has been set.
func (o *TemplateContentVoice) GetDeliveryMethodOk() (*EnumTemplateContentDeliveryMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliveryMethod, true
}

// SetDeliveryMethod sets field value
func (o *TemplateContentVoice) SetDeliveryMethod(v EnumTemplateContentDeliveryMethod) {
	o.DeliveryMethod = v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *TemplateContentVoice) GetVariant() string {
	if o == nil || IsNil(o.Variant) {
		var ret string
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContentVoice) GetVariantOk() (*string, bool) {
	if o == nil || IsNil(o.Variant) {
		return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *TemplateContentVoice) HasVariant() bool {
	if o != nil && !IsNil(o.Variant) {
		return true
	}

	return false
}

// SetVariant gets a reference to the given string and assigns it to the Variant field.
func (o *TemplateContentVoice) SetVariant(v string) {
	o.Variant = &v
}

// GetContent returns the Content field value
func (o *TemplateContentVoice) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *TemplateContentVoice) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *TemplateContentVoice) SetContent(v string) {
	o.Content = v
}

// GetVoice returns the Voice field value if set, zero value otherwise.
func (o *TemplateContentVoice) GetVoice() string {
	if o == nil || IsNil(o.Voice) {
		var ret string
		return ret
	}
	return *o.Voice
}

// GetVoiceOk returns a tuple with the Voice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContentVoice) GetVoiceOk() (*string, bool) {
	if o == nil || IsNil(o.Voice) {
		return nil, false
	}
	return o.Voice, true
}

// HasVoice returns a boolean if a field has been set.
func (o *TemplateContentVoice) HasVoice() bool {
	if o != nil && !IsNil(o.Voice) {
		return true
	}

	return false
}

// SetVoice gets a reference to the given string and assigns it to the Voice field.
func (o *TemplateContentVoice) SetVoice(v string) {
	o.Voice = &v
}

func (o TemplateContentVoice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateContentVoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	toSerialize["locale"] = o.Locale
	toSerialize["deliveryMethod"] = o.DeliveryMethod
	if !IsNil(o.Variant) {
		toSerialize["variant"] = o.Variant
	}
	toSerialize["content"] = o.Content
	if !IsNil(o.Voice) {
		toSerialize["voice"] = o.Voice
	}
	return toSerialize, nil
}

type NullableTemplateContentVoice struct {
	value *TemplateContentVoice
	isSet bool
}

func (v NullableTemplateContentVoice) Get() *TemplateContentVoice {
	return v.value
}

func (v *NullableTemplateContentVoice) Set(val *TemplateContentVoice) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateContentVoice) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateContentVoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateContentVoice(val *TemplateContentVoice) *NullableTemplateContentVoice {
	return &NullableTemplateContentVoice{value: val, isSet: true}
}

func (v NullableTemplateContentVoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateContentVoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
