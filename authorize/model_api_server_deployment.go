/*
PingOne Platform API - Authorize

The PingOne Platform API covering the PingOne Authorize service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorize

import (
	"encoding/json"
	"time"
)

// checks if the APIServerDeployment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIServerDeployment{}

// APIServerDeployment struct for APIServerDeployment
type APIServerDeployment struct {
	Links map[string]LinksHATEOASValue `json:"_links,omitempty"`
	AccessControl *APIServerDeploymentAccessControl `json:"accessControl,omitempty"`
	AuthorizationVersion *APIServerDeploymentAuthorizationVersion `json:"authorizationVersion,omitempty"`
	DecisionEndpoint *APIServerDeploymentDecisionEndpoint `json:"decisionEndpoint,omitempty"`
	// The time of most recent successful deployment. Null if the API service has never been successfully deployed.
	DeployedAt *time.Time `json:"deployedAt,omitempty"`
	Policy *APIServerDeploymentPolicy `json:"policy,omitempty"`
	Status *APIServerDeploymentStatus `json:"status,omitempty"`
}

// NewAPIServerDeployment instantiates a new APIServerDeployment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIServerDeployment() *APIServerDeployment {
	this := APIServerDeployment{}
	return &this
}

// NewAPIServerDeploymentWithDefaults instantiates a new APIServerDeployment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIServerDeploymentWithDefaults() *APIServerDeployment {
	this := APIServerDeployment{}
	return &this
}

func (o APIServerDeployment) hasHalLink(linkIndex string) bool {
	if l, ok := o.GetLinksOk(); ok && l != nil {
		links := l
		if v, ok := links[linkIndex]; ok {
			if h, ok := v.GetHrefOk(); ok && h != nil && *h != "" {
				return true
			}
		}
	}
	return false
}

func (o APIServerDeployment) getHalLink(linkIndex string) LinksHATEOASValue {
	if l, ok := o.GetLinksOk(); ok && l != nil {
		links := l
		if v, ok := links[linkIndex]; ok {
			return v
		}
	}

	var ret LinksHATEOASValue
	return ret
}

func (o APIServerDeployment) getHalLinkOk(linkIndex string) (*LinksHATEOASValue, bool) {
	if l, ok := o.GetLinksOk(); ok && l != nil {
		links := l
		if v, ok := links[linkIndex]; ok {
			return &v, true
		}
	}

	return nil, false
}

func (o APIServerDeployment) IsPaginated() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_NEXT) || o.hasHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o APIServerDeployment) HasPaginationSelf() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o APIServerDeployment) GetPaginationSelfLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o APIServerDeployment) GetPaginationSelfLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o APIServerDeployment) HasPaginationNext() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o APIServerDeployment) GetPaginationNextLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o APIServerDeployment) GetPaginationNextLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o APIServerDeployment) HasPaginationPrevious() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o APIServerDeployment) GetPaginationPreviousLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o APIServerDeployment) GetPaginationPreviousLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_PREV)
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *APIServerDeployment) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServerDeployment) GetLinksOk() (map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return map[string]LinksHATEOASValue{}, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *APIServerDeployment) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *APIServerDeployment) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = v
}

// GetAccessControl returns the AccessControl field value if set, zero value otherwise.
func (o *APIServerDeployment) GetAccessControl() APIServerDeploymentAccessControl {
	if o == nil || IsNil(o.AccessControl) {
		var ret APIServerDeploymentAccessControl
		return ret
	}
	return *o.AccessControl
}

// GetAccessControlOk returns a tuple with the AccessControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServerDeployment) GetAccessControlOk() (*APIServerDeploymentAccessControl, bool) {
	if o == nil || IsNil(o.AccessControl) {
		return nil, false
	}
	return o.AccessControl, true
}

// HasAccessControl returns a boolean if a field has been set.
func (o *APIServerDeployment) HasAccessControl() bool {
	if o != nil && !IsNil(o.AccessControl) {
		return true
	}

	return false
}

// SetAccessControl gets a reference to the given APIServerDeploymentAccessControl and assigns it to the AccessControl field.
func (o *APIServerDeployment) SetAccessControl(v APIServerDeploymentAccessControl) {
	o.AccessControl = &v
}

// GetAuthorizationVersion returns the AuthorizationVersion field value if set, zero value otherwise.
func (o *APIServerDeployment) GetAuthorizationVersion() APIServerDeploymentAuthorizationVersion {
	if o == nil || IsNil(o.AuthorizationVersion) {
		var ret APIServerDeploymentAuthorizationVersion
		return ret
	}
	return *o.AuthorizationVersion
}

// GetAuthorizationVersionOk returns a tuple with the AuthorizationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServerDeployment) GetAuthorizationVersionOk() (*APIServerDeploymentAuthorizationVersion, bool) {
	if o == nil || IsNil(o.AuthorizationVersion) {
		return nil, false
	}
	return o.AuthorizationVersion, true
}

// HasAuthorizationVersion returns a boolean if a field has been set.
func (o *APIServerDeployment) HasAuthorizationVersion() bool {
	if o != nil && !IsNil(o.AuthorizationVersion) {
		return true
	}

	return false
}

// SetAuthorizationVersion gets a reference to the given APIServerDeploymentAuthorizationVersion and assigns it to the AuthorizationVersion field.
func (o *APIServerDeployment) SetAuthorizationVersion(v APIServerDeploymentAuthorizationVersion) {
	o.AuthorizationVersion = &v
}

// GetDecisionEndpoint returns the DecisionEndpoint field value if set, zero value otherwise.
func (o *APIServerDeployment) GetDecisionEndpoint() APIServerDeploymentDecisionEndpoint {
	if o == nil || IsNil(o.DecisionEndpoint) {
		var ret APIServerDeploymentDecisionEndpoint
		return ret
	}
	return *o.DecisionEndpoint
}

// GetDecisionEndpointOk returns a tuple with the DecisionEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServerDeployment) GetDecisionEndpointOk() (*APIServerDeploymentDecisionEndpoint, bool) {
	if o == nil || IsNil(o.DecisionEndpoint) {
		return nil, false
	}
	return o.DecisionEndpoint, true
}

// HasDecisionEndpoint returns a boolean if a field has been set.
func (o *APIServerDeployment) HasDecisionEndpoint() bool {
	if o != nil && !IsNil(o.DecisionEndpoint) {
		return true
	}

	return false
}

// SetDecisionEndpoint gets a reference to the given APIServerDeploymentDecisionEndpoint and assigns it to the DecisionEndpoint field.
func (o *APIServerDeployment) SetDecisionEndpoint(v APIServerDeploymentDecisionEndpoint) {
	o.DecisionEndpoint = &v
}

// GetDeployedAt returns the DeployedAt field value if set, zero value otherwise.
func (o *APIServerDeployment) GetDeployedAt() time.Time {
	if o == nil || IsNil(o.DeployedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeployedAt
}

// GetDeployedAtOk returns a tuple with the DeployedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServerDeployment) GetDeployedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeployedAt) {
		return nil, false
	}
	return o.DeployedAt, true
}

// HasDeployedAt returns a boolean if a field has been set.
func (o *APIServerDeployment) HasDeployedAt() bool {
	if o != nil && !IsNil(o.DeployedAt) {
		return true
	}

	return false
}

// SetDeployedAt gets a reference to the given time.Time and assigns it to the DeployedAt field.
func (o *APIServerDeployment) SetDeployedAt(v time.Time) {
	o.DeployedAt = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *APIServerDeployment) GetPolicy() APIServerDeploymentPolicy {
	if o == nil || IsNil(o.Policy) {
		var ret APIServerDeploymentPolicy
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServerDeployment) GetPolicyOk() (*APIServerDeploymentPolicy, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *APIServerDeployment) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given APIServerDeploymentPolicy and assigns it to the Policy field.
func (o *APIServerDeployment) SetPolicy(v APIServerDeploymentPolicy) {
	o.Policy = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *APIServerDeployment) GetStatus() APIServerDeploymentStatus {
	if o == nil || IsNil(o.Status) {
		var ret APIServerDeploymentStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServerDeployment) GetStatusOk() (*APIServerDeploymentStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *APIServerDeployment) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given APIServerDeploymentStatus and assigns it to the Status field.
func (o *APIServerDeployment) SetStatus(v APIServerDeploymentStatus) {
	o.Status = &v
}

func (o APIServerDeployment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIServerDeployment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.AccessControl) {
		toSerialize["accessControl"] = o.AccessControl
	}
	if !IsNil(o.AuthorizationVersion) {
		toSerialize["authorizationVersion"] = o.AuthorizationVersion
	}
	if !IsNil(o.DecisionEndpoint) {
		toSerialize["decisionEndpoint"] = o.DecisionEndpoint
	}
	if !IsNil(o.DeployedAt) {
		toSerialize["deployedAt"] = o.DeployedAt
	}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableAPIServerDeployment struct {
	value *APIServerDeployment
	isSet bool
}

func (v NullableAPIServerDeployment) Get() *APIServerDeployment {
	return v.value
}

func (v *NullableAPIServerDeployment) Set(val *APIServerDeployment) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIServerDeployment) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIServerDeployment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIServerDeployment(val *APIServerDeployment) *NullableAPIServerDeployment {
	return &NullableAPIServerDeployment{value: val, isSet: true}
}

func (v NullableAPIServerDeployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIServerDeployment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


