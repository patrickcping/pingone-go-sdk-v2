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

// checks if the EmailDomainTrustedEmail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailDomainTrustedEmail{}

// EmailDomainTrustedEmail struct for EmailDomainTrustedEmail
type EmailDomainTrustedEmail struct {
	// A string that specifies the auto generated ID of the trusted email address.
	Id *string `json:"id,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	Domain *EmailDomainTrustedEmailDomain `json:"domain,omitempty"`
	// A string that specifies the trusted email address, for example john.smith@shopco.com.
	EmailAddress string `json:"emailAddress"`
	Status *EnumTrustedEmailStatus `json:"status,omitempty"`
	// A string that specifies the trusted email domain resource’s unique identifier associated with the resource.
	DomainId *string `json:"domainId,omitempty"`
}

// NewEmailDomainTrustedEmail instantiates a new EmailDomainTrustedEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailDomainTrustedEmail(emailAddress string) *EmailDomainTrustedEmail {
	this := EmailDomainTrustedEmail{}
	this.EmailAddress = emailAddress
	return &this
}

// NewEmailDomainTrustedEmailWithDefaults instantiates a new EmailDomainTrustedEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailDomainTrustedEmailWithDefaults() *EmailDomainTrustedEmail {
	this := EmailDomainTrustedEmail{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EmailDomainTrustedEmail) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainTrustedEmail) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EmailDomainTrustedEmail) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EmailDomainTrustedEmail) SetId(v string) {
	o.Id = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *EmailDomainTrustedEmail) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainTrustedEmail) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *EmailDomainTrustedEmail) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *EmailDomainTrustedEmail) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *EmailDomainTrustedEmail) GetDomain() EmailDomainTrustedEmailDomain {
	if o == nil || IsNil(o.Domain) {
		var ret EmailDomainTrustedEmailDomain
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainTrustedEmail) GetDomainOk() (*EmailDomainTrustedEmailDomain, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *EmailDomainTrustedEmail) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given EmailDomainTrustedEmailDomain and assigns it to the Domain field.
func (o *EmailDomainTrustedEmail) SetDomain(v EmailDomainTrustedEmailDomain) {
	o.Domain = &v
}

// GetEmailAddress returns the EmailAddress field value
func (o *EmailDomainTrustedEmail) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *EmailDomainTrustedEmail) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *EmailDomainTrustedEmail) SetEmailAddress(v string) {
	o.EmailAddress = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EmailDomainTrustedEmail) GetStatus() EnumTrustedEmailStatus {
	if o == nil || IsNil(o.Status) {
		var ret EnumTrustedEmailStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainTrustedEmail) GetStatusOk() (*EnumTrustedEmailStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EmailDomainTrustedEmail) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given EnumTrustedEmailStatus and assigns it to the Status field.
func (o *EmailDomainTrustedEmail) SetStatus(v EnumTrustedEmailStatus) {
	o.Status = &v
}

// GetDomainId returns the DomainId field value if set, zero value otherwise.
func (o *EmailDomainTrustedEmail) GetDomainId() string {
	if o == nil || IsNil(o.DomainId) {
		var ret string
		return ret
	}
	return *o.DomainId
}

// GetDomainIdOk returns a tuple with the DomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainTrustedEmail) GetDomainIdOk() (*string, bool) {
	if o == nil || IsNil(o.DomainId) {
		return nil, false
	}
	return o.DomainId, true
}

// HasDomainId returns a boolean if a field has been set.
func (o *EmailDomainTrustedEmail) HasDomainId() bool {
	if o != nil && !IsNil(o.DomainId) {
		return true
	}

	return false
}

// SetDomainId gets a reference to the given string and assigns it to the DomainId field.
func (o *EmailDomainTrustedEmail) SetDomainId(v string) {
	o.DomainId = &v
}

func (o EmailDomainTrustedEmail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailDomainTrustedEmail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	toSerialize["emailAddress"] = o.EmailAddress
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	// skip: domainId is readOnly
	return toSerialize, nil
}

type NullableEmailDomainTrustedEmail struct {
	value *EmailDomainTrustedEmail
	isSet bool
}

func (v NullableEmailDomainTrustedEmail) Get() *EmailDomainTrustedEmail {
	return v.value
}

func (v *NullableEmailDomainTrustedEmail) Set(val *EmailDomainTrustedEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailDomainTrustedEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailDomainTrustedEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailDomainTrustedEmail(val *EmailDomainTrustedEmail) *NullableEmailDomainTrustedEmail {
	return &NullableEmailDomainTrustedEmail{value: val, isSet: true}
}

func (v NullableEmailDomainTrustedEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailDomainTrustedEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


