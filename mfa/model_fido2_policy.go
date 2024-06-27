/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
	"time"
)

// checks if the FIDO2Policy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FIDO2Policy{}

// FIDO2Policy struct for FIDO2Policy
type FIDO2Policy struct {
	Links *map[string]LinksHATEOASValue `json:"_links,omitempty"`
	// FIDO policy's UUID.
	Id *string `json:"id,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// The time the resource was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The time the resource was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	AttestationRequirements EnumFIDO2PolicyAttestationRequirements `json:"attestationRequirements"`
	AuthenticatorAttachment EnumFIDO2PolicyAuthenticatorAttachment `json:"authenticatorAttachment"`
	BackupEligibility FIDO2PolicyBackupEligibility `json:"backupEligibility"`
	// Whether this policy should serve as the default FIDO policy.
	Default *bool `json:"default,omitempty"`
	// Description of the FIDO policy.
	Description *string `json:"description,omitempty"`
	// The device authentication policies that use the relevant FIDO policy. If you include the parameter `expand=deviceAuthenticationPolicies` in the URL of the request, this array is included in the response when reading FIDO policies. Each object in the array contains the ID and the name of the device authentication policy.
	DeviceAuthenticationPolicies []string `json:"deviceAuthenticationPolicies,omitempty"`
	// The name to display for the device in registration and authentication windows. Can be up to 100 characters. If you want to use translatable text, you can use any of the keys listed on the *FIDO Policy* page of the *Self-Service* module and the *Sign On Policy* module. The value of the parameter should include only the part of the key name that comes after the module name, for example, `fidoPolicy.deviceDisplayName01` or `fidoPolicy.deviceDisplayName07`. See the pages in the UI for the full list of keys. For more information on translatable keys, see [Modifying translatable keys](https://docs.pingidentity.com/access/sources/dita/topic?category=p1&resourceid=pingone_modifying_translatable_keys) in the PingOne documentation.
	DeviceDisplayName string `json:"deviceDisplayName"`
	DiscoverableCredentials EnumFIDO2PolicyDiscoverableCredentials `json:"discoverableCredentials"`
	MdsAuthenticatorsRequirements FIDO2PolicyMdsAuthenticatorsRequirements `json:"mdsAuthenticatorsRequirements"`
	// The name to use for the FIDO policy. Can be up to 256 characters.
	Name string `json:"name"`
	// The ID of the relying party. The value should be a domain name, such as `example.com` (in lower-case characters).
	RelyingPartyId string `json:"relyingPartyId"`
	UserDisplayNameAttributes FIDO2PolicyUserDisplayNameAttributes `json:"userDisplayNameAttributes"`
	UserVerification FIDO2PolicyUserVerification `json:"userVerification"`
}

// NewFIDO2Policy instantiates a new FIDO2Policy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFIDO2Policy(attestationRequirements EnumFIDO2PolicyAttestationRequirements, authenticatorAttachment EnumFIDO2PolicyAuthenticatorAttachment, backupEligibility FIDO2PolicyBackupEligibility, deviceDisplayName string, discoverableCredentials EnumFIDO2PolicyDiscoverableCredentials, mdsAuthenticatorsRequirements FIDO2PolicyMdsAuthenticatorsRequirements, name string, relyingPartyId string, userDisplayNameAttributes FIDO2PolicyUserDisplayNameAttributes, userVerification FIDO2PolicyUserVerification) *FIDO2Policy {
	this := FIDO2Policy{}
	this.AttestationRequirements = attestationRequirements
	this.AuthenticatorAttachment = authenticatorAttachment
	this.BackupEligibility = backupEligibility
	this.DeviceDisplayName = deviceDisplayName
	this.DiscoverableCredentials = discoverableCredentials
	this.MdsAuthenticatorsRequirements = mdsAuthenticatorsRequirements
	this.Name = name
	this.RelyingPartyId = relyingPartyId
	this.UserDisplayNameAttributes = userDisplayNameAttributes
	this.UserVerification = userVerification
	return &this
}

// NewFIDO2PolicyWithDefaults instantiates a new FIDO2Policy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFIDO2PolicyWithDefaults() *FIDO2Policy {
	this := FIDO2Policy{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *FIDO2Policy) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FIDO2Policy) GetLinksOk() (*map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *FIDO2Policy) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *FIDO2Policy) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FIDO2Policy) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FIDO2Policy) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FIDO2Policy) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FIDO2Policy) SetId(v string) {
	o.Id = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *FIDO2Policy) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FIDO2Policy) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *FIDO2Policy) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *FIDO2Policy) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FIDO2Policy) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FIDO2Policy) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FIDO2Policy) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FIDO2Policy) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *FIDO2Policy) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FIDO2Policy) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FIDO2Policy) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *FIDO2Policy) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetAttestationRequirements returns the AttestationRequirements field value
func (o *FIDO2Policy) GetAttestationRequirements() EnumFIDO2PolicyAttestationRequirements {
	if o == nil {
		var ret EnumFIDO2PolicyAttestationRequirements
		return ret
	}

	return o.AttestationRequirements
}

// GetAttestationRequirementsOk returns a tuple with the AttestationRequirements field value
// and a boolean to check if the value has been set.
func (o *FIDO2Policy) GetAttestationRequirementsOk() (*EnumFIDO2PolicyAttestationRequirements, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttestationRequirements, true
}

// SetAttestationRequirements sets field value
func (o *FIDO2Policy) SetAttestationRequirements(v EnumFIDO2PolicyAttestationRequirements) {
	o.AttestationRequirements = v
}

// GetAuthenticatorAttachment returns the AuthenticatorAttachment field value
func (o *FIDO2Policy) GetAuthenticatorAttachment() EnumFIDO2PolicyAuthenticatorAttachment {
	if o == nil {
		var ret EnumFIDO2PolicyAuthenticatorAttachment
		return ret
	}

	return o.AuthenticatorAttachment
}

// GetAuthenticatorAttachmentOk returns a tuple with the AuthenticatorAttachment field value
// and a boolean to check if the value has been set.
func (o *FIDO2Policy) GetAuthenticatorAttachmentOk() (*EnumFIDO2PolicyAuthenticatorAttachment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticatorAttachment, true
}

// SetAuthenticatorAttachment sets field value
func (o *FIDO2Policy) SetAuthenticatorAttachment(v EnumFIDO2PolicyAuthenticatorAttachment) {
	o.AuthenticatorAttachment = v
}

// GetBackupEligibility returns the BackupEligibility field value
func (o *FIDO2Policy) GetBackupEligibility() FIDO2PolicyBackupEligibility {
	if o == nil {
		var ret FIDO2PolicyBackupEligibility
		return ret
	}

	return o.BackupEligibility
}

// GetBackupEligibilityOk returns a tuple with the BackupEligibility field value
// and a boolean to check if the value has been set.
func (o *FIDO2Policy) GetBackupEligibilityOk() (*FIDO2PolicyBackupEligibility, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackupEligibility, true
}

// SetBackupEligibility sets field value
func (o *FIDO2Policy) SetBackupEligibility(v FIDO2PolicyBackupEligibility) {
	o.BackupEligibility = v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *FIDO2Policy) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FIDO2Policy) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *FIDO2Policy) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *FIDO2Policy) SetDefault(v bool) {
	o.Default = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FIDO2Policy) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FIDO2Policy) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FIDO2Policy) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FIDO2Policy) SetDescription(v string) {
	o.Description = &v
}

// GetDeviceAuthenticationPolicies returns the DeviceAuthenticationPolicies field value if set, zero value otherwise.
func (o *FIDO2Policy) GetDeviceAuthenticationPolicies() []string {
	if o == nil || IsNil(o.DeviceAuthenticationPolicies) {
		var ret []string
		return ret
	}
	return o.DeviceAuthenticationPolicies
}

// GetDeviceAuthenticationPoliciesOk returns a tuple with the DeviceAuthenticationPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FIDO2Policy) GetDeviceAuthenticationPoliciesOk() ([]string, bool) {
	if o == nil || IsNil(o.DeviceAuthenticationPolicies) {
		return nil, false
	}
	return o.DeviceAuthenticationPolicies, true
}

// HasDeviceAuthenticationPolicies returns a boolean if a field has been set.
func (o *FIDO2Policy) HasDeviceAuthenticationPolicies() bool {
	if o != nil && !IsNil(o.DeviceAuthenticationPolicies) {
		return true
	}

	return false
}

// SetDeviceAuthenticationPolicies gets a reference to the given []string and assigns it to the DeviceAuthenticationPolicies field.
func (o *FIDO2Policy) SetDeviceAuthenticationPolicies(v []string) {
	o.DeviceAuthenticationPolicies = v
}

// GetDeviceDisplayName returns the DeviceDisplayName field value
func (o *FIDO2Policy) GetDeviceDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceDisplayName
}

// GetDeviceDisplayNameOk returns a tuple with the DeviceDisplayName field value
// and a boolean to check if the value has been set.
func (o *FIDO2Policy) GetDeviceDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceDisplayName, true
}

// SetDeviceDisplayName sets field value
func (o *FIDO2Policy) SetDeviceDisplayName(v string) {
	o.DeviceDisplayName = v
}

// GetDiscoverableCredentials returns the DiscoverableCredentials field value
func (o *FIDO2Policy) GetDiscoverableCredentials() EnumFIDO2PolicyDiscoverableCredentials {
	if o == nil {
		var ret EnumFIDO2PolicyDiscoverableCredentials
		return ret
	}

	return o.DiscoverableCredentials
}

// GetDiscoverableCredentialsOk returns a tuple with the DiscoverableCredentials field value
// and a boolean to check if the value has been set.
func (o *FIDO2Policy) GetDiscoverableCredentialsOk() (*EnumFIDO2PolicyDiscoverableCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscoverableCredentials, true
}

// SetDiscoverableCredentials sets field value
func (o *FIDO2Policy) SetDiscoverableCredentials(v EnumFIDO2PolicyDiscoverableCredentials) {
	o.DiscoverableCredentials = v
}

// GetMdsAuthenticatorsRequirements returns the MdsAuthenticatorsRequirements field value
func (o *FIDO2Policy) GetMdsAuthenticatorsRequirements() FIDO2PolicyMdsAuthenticatorsRequirements {
	if o == nil {
		var ret FIDO2PolicyMdsAuthenticatorsRequirements
		return ret
	}

	return o.MdsAuthenticatorsRequirements
}

// GetMdsAuthenticatorsRequirementsOk returns a tuple with the MdsAuthenticatorsRequirements field value
// and a boolean to check if the value has been set.
func (o *FIDO2Policy) GetMdsAuthenticatorsRequirementsOk() (*FIDO2PolicyMdsAuthenticatorsRequirements, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MdsAuthenticatorsRequirements, true
}

// SetMdsAuthenticatorsRequirements sets field value
func (o *FIDO2Policy) SetMdsAuthenticatorsRequirements(v FIDO2PolicyMdsAuthenticatorsRequirements) {
	o.MdsAuthenticatorsRequirements = v
}

// GetName returns the Name field value
func (o *FIDO2Policy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FIDO2Policy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FIDO2Policy) SetName(v string) {
	o.Name = v
}

// GetRelyingPartyId returns the RelyingPartyId field value
func (o *FIDO2Policy) GetRelyingPartyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelyingPartyId
}

// GetRelyingPartyIdOk returns a tuple with the RelyingPartyId field value
// and a boolean to check if the value has been set.
func (o *FIDO2Policy) GetRelyingPartyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelyingPartyId, true
}

// SetRelyingPartyId sets field value
func (o *FIDO2Policy) SetRelyingPartyId(v string) {
	o.RelyingPartyId = v
}

// GetUserDisplayNameAttributes returns the UserDisplayNameAttributes field value
func (o *FIDO2Policy) GetUserDisplayNameAttributes() FIDO2PolicyUserDisplayNameAttributes {
	if o == nil {
		var ret FIDO2PolicyUserDisplayNameAttributes
		return ret
	}

	return o.UserDisplayNameAttributes
}

// GetUserDisplayNameAttributesOk returns a tuple with the UserDisplayNameAttributes field value
// and a boolean to check if the value has been set.
func (o *FIDO2Policy) GetUserDisplayNameAttributesOk() (*FIDO2PolicyUserDisplayNameAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserDisplayNameAttributes, true
}

// SetUserDisplayNameAttributes sets field value
func (o *FIDO2Policy) SetUserDisplayNameAttributes(v FIDO2PolicyUserDisplayNameAttributes) {
	o.UserDisplayNameAttributes = v
}

// GetUserVerification returns the UserVerification field value
func (o *FIDO2Policy) GetUserVerification() FIDO2PolicyUserVerification {
	if o == nil {
		var ret FIDO2PolicyUserVerification
		return ret
	}

	return o.UserVerification
}

// GetUserVerificationOk returns a tuple with the UserVerification field value
// and a boolean to check if the value has been set.
func (o *FIDO2Policy) GetUserVerificationOk() (*FIDO2PolicyUserVerification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserVerification, true
}

// SetUserVerification sets field value
func (o *FIDO2Policy) SetUserVerification(v FIDO2PolicyUserVerification) {
	o.UserVerification = v
}

func (o FIDO2Policy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FIDO2Policy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	toSerialize["attestationRequirements"] = o.AttestationRequirements
	toSerialize["authenticatorAttachment"] = o.AuthenticatorAttachment
	toSerialize["backupEligibility"] = o.BackupEligibility
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DeviceAuthenticationPolicies) {
		toSerialize["deviceAuthenticationPolicies"] = o.DeviceAuthenticationPolicies
	}
	toSerialize["deviceDisplayName"] = o.DeviceDisplayName
	toSerialize["discoverableCredentials"] = o.DiscoverableCredentials
	toSerialize["mdsAuthenticatorsRequirements"] = o.MdsAuthenticatorsRequirements
	toSerialize["name"] = o.Name
	toSerialize["relyingPartyId"] = o.RelyingPartyId
	toSerialize["userDisplayNameAttributes"] = o.UserDisplayNameAttributes
	toSerialize["userVerification"] = o.UserVerification
	return toSerialize, nil
}

type NullableFIDO2Policy struct {
	value *FIDO2Policy
	isSet bool
}

func (v NullableFIDO2Policy) Get() *FIDO2Policy {
	return v.value
}

func (v *NullableFIDO2Policy) Set(val *FIDO2Policy) {
	v.value = val
	v.isSet = true
}

func (v NullableFIDO2Policy) IsSet() bool {
	return v.isSet
}

func (v *NullableFIDO2Policy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFIDO2Policy(val *FIDO2Policy) *NullableFIDO2Policy {
	return &NullableFIDO2Policy{value: val, isSet: true}
}

func (v NullableFIDO2Policy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFIDO2Policy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


