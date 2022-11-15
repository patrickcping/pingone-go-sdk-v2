/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// TemplateContent struct for TemplateContent
type TemplateContent struct {
	// The template id.
	Id *string `json:"id,omitempty"`
	// The time the resource was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The time the resource was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// Specifies whether the template is a predefined default template.
	Default *bool `json:"default,omitempty"`
	// A valid case-insensitive locale, complying with the ISO-639 language code and ISO-3166 country code standards: Two-character language code, for example, \"en\". Two-character language code followed by a two-character country code, separated by an underscore or dash, for example: \"en_GB\", \"en-GB\". Cannot be changed after it is initially set in `POST /environments/{{envID}}/templates/{{templateName}}/contents`. 
	Locale string `json:"locale"`
	// The content's delivery method. Possible values are `Email`, `SMS`, `Voice` or `Push`. Cannot be changed after it is initially set in `POST /environments/{{envID}}/templates/{{templateName}}/contents`.
	DeliveryMethod string `json:"deliveryMethod"`
	// Holds the unique user-defined name for each content variant that uses the same template + `deliveryMethod` + `locale` combination. This property is case insensitive and has a limit of 100 characters. For more information, see [Creating custom contents](https://apidocs.pingidentity.com/pingone/platform/v1/api/#notifications-templates-creating-custom-contents).
	Variant *string `json:"variant,omitempty"`
}

// NewTemplateContent instantiates a new TemplateContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateContent(locale string, deliveryMethod string) *TemplateContent {
	this := TemplateContent{}
	this.Locale = locale
	this.DeliveryMethod = deliveryMethod
	return &this
}

// NewTemplateContentWithDefaults instantiates a new TemplateContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateContentWithDefaults() *TemplateContent {
	this := TemplateContent{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TemplateContent) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContent) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TemplateContent) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TemplateContent) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TemplateContent) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContent) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TemplateContent) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *TemplateContent) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TemplateContent) GetUpdatedAt() string {
	if o == nil || isNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContent) GetUpdatedAtOk() (*string, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TemplateContent) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *TemplateContent) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *TemplateContent) GetDefault() bool {
	if o == nil || isNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContent) GetDefaultOk() (*bool, bool) {
	if o == nil || isNil(o.Default) {
    return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *TemplateContent) HasDefault() bool {
	if o != nil && !isNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *TemplateContent) SetDefault(v bool) {
	o.Default = &v
}

// GetLocale returns the Locale field value
func (o *TemplateContent) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *TemplateContent) GetLocaleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *TemplateContent) SetLocale(v string) {
	o.Locale = v
}

// GetDeliveryMethod returns the DeliveryMethod field value
func (o *TemplateContent) GetDeliveryMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeliveryMethod
}

// GetDeliveryMethodOk returns a tuple with the DeliveryMethod field value
// and a boolean to check if the value has been set.
func (o *TemplateContent) GetDeliveryMethodOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DeliveryMethod, true
}

// SetDeliveryMethod sets field value
func (o *TemplateContent) SetDeliveryMethod(v string) {
	o.DeliveryMethod = v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *TemplateContent) GetVariant() string {
	if o == nil || isNil(o.Variant) {
		var ret string
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContent) GetVariantOk() (*string, bool) {
	if o == nil || isNil(o.Variant) {
    return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *TemplateContent) HasVariant() bool {
	if o != nil && !isNil(o.Variant) {
		return true
	}

	return false
}

// SetVariant gets a reference to the given string and assigns it to the Variant field.
func (o *TemplateContent) SetVariant(v string) {
	o.Variant = &v
}

func (o TemplateContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !isNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if true {
		toSerialize["locale"] = o.Locale
	}
	if true {
		toSerialize["deliveryMethod"] = o.DeliveryMethod
	}
	if !isNil(o.Variant) {
		toSerialize["variant"] = o.Variant
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateContent struct {
	value *TemplateContent
	isSet bool
}

func (v NullableTemplateContent) Get() *TemplateContent {
	return v.value
}

func (v *NullableTemplateContent) Set(val *TemplateContent) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateContent) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateContent(val *TemplateContent) *NullableTemplateContent {
	return &NullableTemplateContent{value: val, isSet: true}
}

func (v NullableTemplateContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

