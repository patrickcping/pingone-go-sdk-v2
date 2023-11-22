/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// checks if the TemplateContentEmail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateContentEmail{}

// TemplateContentEmail struct for TemplateContentEmail
type TemplateContentEmail struct {
	Links *LinksHATEOAS `json:"_links,omitempty"`
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
	DeliveryMethod EnumTemplateContentDeliveryMethod `json:"deliveryMethod"`
	// Holds the unique user-defined name for each content variant that uses the same template + `deliveryMethod` + `locale` combination. This property is case insensitive and has a limit of 100 characters. For more information, see [Creating custom contents](https://apidocs.pingidentity.com/pingone/platform/v1/api/#notifications-templates-creating-custom-contents).
	Variant *string `json:"variant,omitempty"`
	// The email text. Email text cannot be larger than 100 kB. Email text can contain HTML. If supported, this can include variables.
	Body string `json:"body"`
	From *TemplateContentEmailAllOfFrom `json:"from,omitempty"`
	// The email's subject line. Cannot exceed 256 characters. If supported, can include variables.
	Subject *string `json:"subject,omitempty"`
	ReplyTo *TemplateContentEmailAllOfReplyTo `json:"replyTo,omitempty"`
	// If not specified, `UTF-8` is the default value.
	Charset *string `json:"charset,omitempty"`
	// If not specified, `text/html` is the default value.
	EmailContentType *string `json:"emailContentType,omitempty"`
}

type _TemplateContentEmail TemplateContentEmail

// NewTemplateContentEmail instantiates a new TemplateContentEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateContentEmail(locale string, deliveryMethod EnumTemplateContentDeliveryMethod, body string) *TemplateContentEmail {
	this := TemplateContentEmail{}
	this.Locale = locale
	this.DeliveryMethod = deliveryMethod
	this.Body = body
	var charset string = "UTF-8"
	this.Charset = &charset
	var emailContentType string = "text/html"
	this.EmailContentType = &emailContentType
	return &this
}

// NewTemplateContentEmailWithDefaults instantiates a new TemplateContentEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateContentEmailWithDefaults() *TemplateContentEmail {
	this := TemplateContentEmail{}
	var charset string = "UTF-8"
	this.Charset = &charset
	var emailContentType string = "text/html"
	this.EmailContentType = &emailContentType
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TemplateContentEmail) GetLinks() LinksHATEOAS {
	if o == nil || IsNil(o.Links) {
		var ret LinksHATEOAS
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContentEmail) GetLinksOk() (*LinksHATEOAS, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TemplateContentEmail) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksHATEOAS and assigns it to the Links field.
func (o *TemplateContentEmail) SetLinks(v LinksHATEOAS) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TemplateContentEmail) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContentEmail) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TemplateContentEmail) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TemplateContentEmail) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TemplateContentEmail) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContentEmail) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TemplateContentEmail) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *TemplateContentEmail) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TemplateContentEmail) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContentEmail) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TemplateContentEmail) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *TemplateContentEmail) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *TemplateContentEmail) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContentEmail) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *TemplateContentEmail) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *TemplateContentEmail) SetDefault(v bool) {
	o.Default = &v
}

// GetLocale returns the Locale field value
func (o *TemplateContentEmail) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *TemplateContentEmail) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *TemplateContentEmail) SetLocale(v string) {
	o.Locale = v
}

// GetDeliveryMethod returns the DeliveryMethod field value
func (o *TemplateContentEmail) GetDeliveryMethod() EnumTemplateContentDeliveryMethod {
	if o == nil {
		var ret EnumTemplateContentDeliveryMethod
		return ret
	}

	return o.DeliveryMethod
}

// GetDeliveryMethodOk returns a tuple with the DeliveryMethod field value
// and a boolean to check if the value has been set.
func (o *TemplateContentEmail) GetDeliveryMethodOk() (*EnumTemplateContentDeliveryMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliveryMethod, true
}

// SetDeliveryMethod sets field value
func (o *TemplateContentEmail) SetDeliveryMethod(v EnumTemplateContentDeliveryMethod) {
	o.DeliveryMethod = v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *TemplateContentEmail) GetVariant() string {
	if o == nil || IsNil(o.Variant) {
		var ret string
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContentEmail) GetVariantOk() (*string, bool) {
	if o == nil || IsNil(o.Variant) {
		return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *TemplateContentEmail) HasVariant() bool {
	if o != nil && !IsNil(o.Variant) {
		return true
	}

	return false
}

// SetVariant gets a reference to the given string and assigns it to the Variant field.
func (o *TemplateContentEmail) SetVariant(v string) {
	o.Variant = &v
}

// GetBody returns the Body field value
func (o *TemplateContentEmail) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *TemplateContentEmail) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *TemplateContentEmail) SetBody(v string) {
	o.Body = v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *TemplateContentEmail) GetFrom() TemplateContentEmailAllOfFrom {
	if o == nil || IsNil(o.From) {
		var ret TemplateContentEmailAllOfFrom
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContentEmail) GetFromOk() (*TemplateContentEmailAllOfFrom, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *TemplateContentEmail) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given TemplateContentEmailAllOfFrom and assigns it to the From field.
func (o *TemplateContentEmail) SetFrom(v TemplateContentEmailAllOfFrom) {
	o.From = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *TemplateContentEmail) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContentEmail) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *TemplateContentEmail) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *TemplateContentEmail) SetSubject(v string) {
	o.Subject = &v
}

// GetReplyTo returns the ReplyTo field value if set, zero value otherwise.
func (o *TemplateContentEmail) GetReplyTo() TemplateContentEmailAllOfReplyTo {
	if o == nil || IsNil(o.ReplyTo) {
		var ret TemplateContentEmailAllOfReplyTo
		return ret
	}
	return *o.ReplyTo
}

// GetReplyToOk returns a tuple with the ReplyTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContentEmail) GetReplyToOk() (*TemplateContentEmailAllOfReplyTo, bool) {
	if o == nil || IsNil(o.ReplyTo) {
		return nil, false
	}
	return o.ReplyTo, true
}

// HasReplyTo returns a boolean if a field has been set.
func (o *TemplateContentEmail) HasReplyTo() bool {
	if o != nil && !IsNil(o.ReplyTo) {
		return true
	}

	return false
}

// SetReplyTo gets a reference to the given TemplateContentEmailAllOfReplyTo and assigns it to the ReplyTo field.
func (o *TemplateContentEmail) SetReplyTo(v TemplateContentEmailAllOfReplyTo) {
	o.ReplyTo = &v
}

// GetCharset returns the Charset field value if set, zero value otherwise.
func (o *TemplateContentEmail) GetCharset() string {
	if o == nil || IsNil(o.Charset) {
		var ret string
		return ret
	}
	return *o.Charset
}

// GetCharsetOk returns a tuple with the Charset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContentEmail) GetCharsetOk() (*string, bool) {
	if o == nil || IsNil(o.Charset) {
		return nil, false
	}
	return o.Charset, true
}

// HasCharset returns a boolean if a field has been set.
func (o *TemplateContentEmail) HasCharset() bool {
	if o != nil && !IsNil(o.Charset) {
		return true
	}

	return false
}

// SetCharset gets a reference to the given string and assigns it to the Charset field.
func (o *TemplateContentEmail) SetCharset(v string) {
	o.Charset = &v
}

// GetEmailContentType returns the EmailContentType field value if set, zero value otherwise.
func (o *TemplateContentEmail) GetEmailContentType() string {
	if o == nil || IsNil(o.EmailContentType) {
		var ret string
		return ret
	}
	return *o.EmailContentType
}

// GetEmailContentTypeOk returns a tuple with the EmailContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContentEmail) GetEmailContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EmailContentType) {
		return nil, false
	}
	return o.EmailContentType, true
}

// HasEmailContentType returns a boolean if a field has been set.
func (o *TemplateContentEmail) HasEmailContentType() bool {
	if o != nil && !IsNil(o.EmailContentType) {
		return true
	}

	return false
}

// SetEmailContentType gets a reference to the given string and assigns it to the EmailContentType field.
func (o *TemplateContentEmail) SetEmailContentType(v string) {
	o.EmailContentType = &v
}

func (o TemplateContentEmail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateContentEmail) ToMap() (map[string]interface{}, error) {
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
	toSerialize["body"] = o.Body
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.ReplyTo) {
		toSerialize["replyTo"] = o.ReplyTo
	}
	if !IsNil(o.Charset) {
		toSerialize["charset"] = o.Charset
	}
	if !IsNil(o.EmailContentType) {
		toSerialize["emailContentType"] = o.EmailContentType
	}
	return toSerialize, nil
}

func (o *TemplateContentEmail) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"locale",
		"deliveryMethod",
		"body",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTemplateContentEmail := _TemplateContentEmail{}

	err = json.Unmarshal(bytes, &varTemplateContentEmail)

	if err != nil {
		return err
	}

	*o = TemplateContentEmail(varTemplateContentEmail)

	return err
}

type NullableTemplateContentEmail struct {
	value *TemplateContentEmail
	isSet bool
}

func (v NullableTemplateContentEmail) Get() *TemplateContentEmail {
	return v.value
}

func (v *NullableTemplateContentEmail) Set(val *TemplateContentEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateContentEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateContentEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateContentEmail(val *TemplateContentEmail) *NullableTemplateContentEmail {
	return &NullableTemplateContentEmail{value: val, isSet: true}
}

func (v NullableTemplateContentEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateContentEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


