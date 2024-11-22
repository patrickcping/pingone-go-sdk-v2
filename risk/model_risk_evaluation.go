/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RiskEvaluation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskEvaluation{}

// RiskEvaluation struct for RiskEvaluation
type RiskEvaluation struct {
	Links map[string]LinksHATEOASValue `json:"_links,omitempty"`
	// The time the resource was created (format ISO-8061).
	CreatedAt *string `json:"createdAt,omitempty"`
	// A details object that provides additional information about the risk evaluation.
	Details *RiskEvaluationDetails `json:"details,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// An object that specifies the attributes to identify the event. This is a required property. For more information about event attributes, see the Event Data Model table.
	Event RiskEvaluationEvent `json:"event"`
	// A string that specifies the resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	RiskPolicySet *RiskEvaluationRiskPolicySet `json:"riskPolicySet,omitempty"`
	// A result object that specifies the result that corresponds to the risk policy that evaluates as true. If there are several risk policies that evaluate as true, the result that corresponds to the highest priority risk policy is returned. If no risk policy evaluates as true, the result is the defaultResult of the policy set.
	Result *RiskEvaluationResult `json:"result,omitempty"`
	// The time the resource was last updated (format ISO-8061).
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type _RiskEvaluation RiskEvaluation

// NewRiskEvaluation instantiates a new RiskEvaluation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskEvaluation(event RiskEvaluationEvent) *RiskEvaluation {
	this := RiskEvaluation{}
	this.Event = event
	return &this
}

// NewRiskEvaluationWithDefaults instantiates a new RiskEvaluation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskEvaluationWithDefaults() *RiskEvaluation {
	this := RiskEvaluation{}
	return &this
}

func (o RiskEvaluation) hasHalLink(linkIndex string) bool {
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

func (o RiskEvaluation) getHalLink(linkIndex string) LinksHATEOASValue {
	if l, ok := o.GetLinksOk(); ok && l != nil {
		links := l
		if v, ok := links[linkIndex]; ok {
			return v
		}
	}

	var ret LinksHATEOASValue
	return ret
}

func (o RiskEvaluation) getHalLinkOk(linkIndex string) (*LinksHATEOASValue, bool) {
	if l, ok := o.GetLinksOk(); ok && l != nil {
		links := l
		if v, ok := links[linkIndex]; ok {
			return &v, true
		}
	}

	return nil, false
}

func (o RiskEvaluation) IsPaginated() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_NEXT) || o.hasHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o RiskEvaluation) HasPaginationSelf() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o RiskEvaluation) GetPaginationSelfLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o RiskEvaluation) GetPaginationSelfLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o RiskEvaluation) HasPaginationNext() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o RiskEvaluation) GetPaginationNextLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o RiskEvaluation) GetPaginationNextLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o RiskEvaluation) HasPaginationPrevious() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o RiskEvaluation) GetPaginationPreviousLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o RiskEvaluation) GetPaginationPreviousLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_PREV)
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RiskEvaluation) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluation) GetLinksOk() (map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return map[string]LinksHATEOASValue{}, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RiskEvaluation) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *RiskEvaluation) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RiskEvaluation) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluation) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RiskEvaluation) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *RiskEvaluation) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *RiskEvaluation) GetDetails() RiskEvaluationDetails {
	if o == nil || IsNil(o.Details) {
		var ret RiskEvaluationDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluation) GetDetailsOk() (*RiskEvaluationDetails, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *RiskEvaluation) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given RiskEvaluationDetails and assigns it to the Details field.
func (o *RiskEvaluation) SetDetails(v RiskEvaluationDetails) {
	o.Details = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *RiskEvaluation) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluation) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *RiskEvaluation) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *RiskEvaluation) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetEvent returns the Event field value
func (o *RiskEvaluation) GetEvent() RiskEvaluationEvent {
	if o == nil {
		var ret RiskEvaluationEvent
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *RiskEvaluation) GetEventOk() (*RiskEvaluationEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *RiskEvaluation) SetEvent(v RiskEvaluationEvent) {
	o.Event = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RiskEvaluation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RiskEvaluation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RiskEvaluation) SetId(v string) {
	o.Id = &v
}

// GetRiskPolicySet returns the RiskPolicySet field value if set, zero value otherwise.
func (o *RiskEvaluation) GetRiskPolicySet() RiskEvaluationRiskPolicySet {
	if o == nil || IsNil(o.RiskPolicySet) {
		var ret RiskEvaluationRiskPolicySet
		return ret
	}
	return *o.RiskPolicySet
}

// GetRiskPolicySetOk returns a tuple with the RiskPolicySet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluation) GetRiskPolicySetOk() (*RiskEvaluationRiskPolicySet, bool) {
	if o == nil || IsNil(o.RiskPolicySet) {
		return nil, false
	}
	return o.RiskPolicySet, true
}

// HasRiskPolicySet returns a boolean if a field has been set.
func (o *RiskEvaluation) HasRiskPolicySet() bool {
	if o != nil && !IsNil(o.RiskPolicySet) {
		return true
	}

	return false
}

// SetRiskPolicySet gets a reference to the given RiskEvaluationRiskPolicySet and assigns it to the RiskPolicySet field.
func (o *RiskEvaluation) SetRiskPolicySet(v RiskEvaluationRiskPolicySet) {
	o.RiskPolicySet = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *RiskEvaluation) GetResult() RiskEvaluationResult {
	if o == nil || IsNil(o.Result) {
		var ret RiskEvaluationResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluation) GetResultOk() (*RiskEvaluationResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *RiskEvaluation) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given RiskEvaluationResult and assigns it to the Result field.
func (o *RiskEvaluation) SetResult(v RiskEvaluationResult) {
	o.Result = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RiskEvaluation) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluation) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RiskEvaluation) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *RiskEvaluation) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o RiskEvaluation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskEvaluation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	toSerialize["event"] = o.Event
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.RiskPolicySet) {
		toSerialize["riskPolicySet"] = o.RiskPolicySet
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

func (o *RiskEvaluation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRiskEvaluation := _RiskEvaluation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRiskEvaluation)

	if err != nil {
		return err
	}

	*o = RiskEvaluation(varRiskEvaluation)

	return err
}

type NullableRiskEvaluation struct {
	value *RiskEvaluation
	isSet bool
}

func (v NullableRiskEvaluation) Get() *RiskEvaluation {
	return v.value
}

func (v *NullableRiskEvaluation) Set(val *RiskEvaluation) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskEvaluation) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskEvaluation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskEvaluation(val *RiskEvaluation) *NullableRiskEvaluation {
	return &NullableRiskEvaluation{value: val, isSet: true}
}

func (v NullableRiskEvaluation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskEvaluation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


