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

// checks if the SignOnPolicyActionIDP type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignOnPolicyActionIDP{}

// SignOnPolicyActionIDP struct for SignOnPolicyActionIDP
type SignOnPolicyActionIDP struct {
	Links *LinksHATEOAS `json:"_links,omitempty"`
	Condition *SignOnPolicyActionCommonConditionOrOrInner `json:"condition,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// A string that specifies the sign-on policy assignment resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	// An integer that specifies the order in which the policy referenced by this assignment is evaluated during an authentication flow relative to other policies. An assignment with a lower priority will be evaluated first. This is a required property.
	Priority int32 `json:"priority"`
	SignOnPolicy *SignOnPolicyActionCommonSignOnPolicy `json:"signOnPolicy,omitempty"`
	Type EnumSignOnPolicyType `json:"type"`
	// A string that designates the sign-on policies included in the authorization flow request. Options can include the PingOne predefined sign-on policies, Single_Factor and Multi_Factor, or any custom defined sign-on policy names. Sign-on policy names should be listed in order of preference, and they must be assigned to the application. This property can be configured on the identity provider action and is passed to the identity provider if the identity provider is of type `SAML` or `OPENID_CONNECT`.
	AcrValues *string `json:"acrValues,omitempty"`
	IdentityProvider SignOnPolicyActionIDPAllOfIdentityProvider `json:"identityProvider"`
	// A boolean that specifies whether to pass in a login hint to the identity provider on the authentication request. Based on user context, the login hint is set if (1) the user is set on the flow, and (2) the user already has an account link for the identity provider. If both of these conditions are true, then the user is sent to the identity provider with a login hint equal to their externalId for the identity provider (saved on the account link). If these conditions are not true, then the API checks see if there is an OIDC login hint on the flow. If so, that login hint is used. If none of these conditions are true, the login hint parameter is not included on the authorization request to the identity provider.
	PassUserContext *bool `json:"passUserContext,omitempty"`
	Registration *SignOnPolicyActionIDPAllOfRegistration `json:"registration,omitempty"`
}

// NewSignOnPolicyActionIDP instantiates a new SignOnPolicyActionIDP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionIDP(priority int32, type_ EnumSignOnPolicyType, identityProvider SignOnPolicyActionIDPAllOfIdentityProvider) *SignOnPolicyActionIDP {
	this := SignOnPolicyActionIDP{}
	this.Priority = priority
	this.Type = type_
	this.IdentityProvider = identityProvider
	return &this
}

// NewSignOnPolicyActionIDPWithDefaults instantiates a new SignOnPolicyActionIDP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionIDPWithDefaults() *SignOnPolicyActionIDP {
	this := SignOnPolicyActionIDP{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDP) GetLinks() LinksHATEOAS {
	if o == nil || IsNil(o.Links) {
		var ret LinksHATEOAS
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDP) GetLinksOk() (*LinksHATEOAS, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDP) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksHATEOAS and assigns it to the Links field.
func (o *SignOnPolicyActionIDP) SetLinks(v LinksHATEOAS) {
	o.Links = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDP) GetCondition() SignOnPolicyActionCommonConditionOrOrInner {
	if o == nil || IsNil(o.Condition) {
		var ret SignOnPolicyActionCommonConditionOrOrInner
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDP) GetConditionOk() (*SignOnPolicyActionCommonConditionOrOrInner, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDP) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given SignOnPolicyActionCommonConditionOrOrInner and assigns it to the Condition field.
func (o *SignOnPolicyActionIDP) SetCondition(v SignOnPolicyActionCommonConditionOrOrInner) {
	o.Condition = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDP) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDP) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDP) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *SignOnPolicyActionIDP) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDP) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDP) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDP) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SignOnPolicyActionIDP) SetId(v string) {
	o.Id = &v
}

// GetPriority returns the Priority field value
func (o *SignOnPolicyActionIDP) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDP) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *SignOnPolicyActionIDP) SetPriority(v int32) {
	o.Priority = v
}

// GetSignOnPolicy returns the SignOnPolicy field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDP) GetSignOnPolicy() SignOnPolicyActionCommonSignOnPolicy {
	if o == nil || IsNil(o.SignOnPolicy) {
		var ret SignOnPolicyActionCommonSignOnPolicy
		return ret
	}
	return *o.SignOnPolicy
}

// GetSignOnPolicyOk returns a tuple with the SignOnPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDP) GetSignOnPolicyOk() (*SignOnPolicyActionCommonSignOnPolicy, bool) {
	if o == nil || IsNil(o.SignOnPolicy) {
		return nil, false
	}
	return o.SignOnPolicy, true
}

// HasSignOnPolicy returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDP) HasSignOnPolicy() bool {
	if o != nil && !IsNil(o.SignOnPolicy) {
		return true
	}

	return false
}

// SetSignOnPolicy gets a reference to the given SignOnPolicyActionCommonSignOnPolicy and assigns it to the SignOnPolicy field.
func (o *SignOnPolicyActionIDP) SetSignOnPolicy(v SignOnPolicyActionCommonSignOnPolicy) {
	o.SignOnPolicy = &v
}

// GetType returns the Type field value
func (o *SignOnPolicyActionIDP) GetType() EnumSignOnPolicyType {
	if o == nil {
		var ret EnumSignOnPolicyType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDP) GetTypeOk() (*EnumSignOnPolicyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SignOnPolicyActionIDP) SetType(v EnumSignOnPolicyType) {
	o.Type = v
}

// GetAcrValues returns the AcrValues field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDP) GetAcrValues() string {
	if o == nil || IsNil(o.AcrValues) {
		var ret string
		return ret
	}
	return *o.AcrValues
}

// GetAcrValuesOk returns a tuple with the AcrValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDP) GetAcrValuesOk() (*string, bool) {
	if o == nil || IsNil(o.AcrValues) {
		return nil, false
	}
	return o.AcrValues, true
}

// HasAcrValues returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDP) HasAcrValues() bool {
	if o != nil && !IsNil(o.AcrValues) {
		return true
	}

	return false
}

// SetAcrValues gets a reference to the given string and assigns it to the AcrValues field.
func (o *SignOnPolicyActionIDP) SetAcrValues(v string) {
	o.AcrValues = &v
}

// GetIdentityProvider returns the IdentityProvider field value
func (o *SignOnPolicyActionIDP) GetIdentityProvider() SignOnPolicyActionIDPAllOfIdentityProvider {
	if o == nil {
		var ret SignOnPolicyActionIDPAllOfIdentityProvider
		return ret
	}

	return o.IdentityProvider
}

// GetIdentityProviderOk returns a tuple with the IdentityProvider field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDP) GetIdentityProviderOk() (*SignOnPolicyActionIDPAllOfIdentityProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityProvider, true
}

// SetIdentityProvider sets field value
func (o *SignOnPolicyActionIDP) SetIdentityProvider(v SignOnPolicyActionIDPAllOfIdentityProvider) {
	o.IdentityProvider = v
}

// GetPassUserContext returns the PassUserContext field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDP) GetPassUserContext() bool {
	if o == nil || IsNil(o.PassUserContext) {
		var ret bool
		return ret
	}
	return *o.PassUserContext
}

// GetPassUserContextOk returns a tuple with the PassUserContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDP) GetPassUserContextOk() (*bool, bool) {
	if o == nil || IsNil(o.PassUserContext) {
		return nil, false
	}
	return o.PassUserContext, true
}

// HasPassUserContext returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDP) HasPassUserContext() bool {
	if o != nil && !IsNil(o.PassUserContext) {
		return true
	}

	return false
}

// SetPassUserContext gets a reference to the given bool and assigns it to the PassUserContext field.
func (o *SignOnPolicyActionIDP) SetPassUserContext(v bool) {
	o.PassUserContext = &v
}

// GetRegistration returns the Registration field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDP) GetRegistration() SignOnPolicyActionIDPAllOfRegistration {
	if o == nil || IsNil(o.Registration) {
		var ret SignOnPolicyActionIDPAllOfRegistration
		return ret
	}
	return *o.Registration
}

// GetRegistrationOk returns a tuple with the Registration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDP) GetRegistrationOk() (*SignOnPolicyActionIDPAllOfRegistration, bool) {
	if o == nil || IsNil(o.Registration) {
		return nil, false
	}
	return o.Registration, true
}

// HasRegistration returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDP) HasRegistration() bool {
	if o != nil && !IsNil(o.Registration) {
		return true
	}

	return false
}

// SetRegistration gets a reference to the given SignOnPolicyActionIDPAllOfRegistration and assigns it to the Registration field.
func (o *SignOnPolicyActionIDP) SetRegistration(v SignOnPolicyActionIDPAllOfRegistration) {
	o.Registration = &v
}

func (o SignOnPolicyActionIDP) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignOnPolicyActionIDP) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	// skip: id is readOnly
	toSerialize["priority"] = o.Priority
	if !IsNil(o.SignOnPolicy) {
		toSerialize["signOnPolicy"] = o.SignOnPolicy
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.AcrValues) {
		toSerialize["acrValues"] = o.AcrValues
	}
	toSerialize["identityProvider"] = o.IdentityProvider
	if !IsNil(o.PassUserContext) {
		toSerialize["passUserContext"] = o.PassUserContext
	}
	if !IsNil(o.Registration) {
		toSerialize["registration"] = o.Registration
	}
	return toSerialize, nil
}

type NullableSignOnPolicyActionIDP struct {
	value *SignOnPolicyActionIDP
	isSet bool
}

func (v NullableSignOnPolicyActionIDP) Get() *SignOnPolicyActionIDP {
	return v.value
}

func (v *NullableSignOnPolicyActionIDP) Set(val *SignOnPolicyActionIDP) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionIDP) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionIDP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionIDP(val *SignOnPolicyActionIDP) *NullableSignOnPolicyActionIDP {
	return &NullableSignOnPolicyActionIDP{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionIDP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionIDP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


