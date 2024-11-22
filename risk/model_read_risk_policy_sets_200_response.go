/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"encoding/json"
)

// checks if the ReadRiskPolicySets200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadRiskPolicySets200Response{}

// ReadRiskPolicySets200Response struct for ReadRiskPolicySets200Response
type ReadRiskPolicySets200Response struct {
	Links map[string]LinksHATEOASValue `json:"_links,omitempty"`
	Embedded *ReadRiskPolicySets200ResponseEmbedded `json:"_embedded,omitempty"`
	Count *float32 `json:"count,omitempty"`
	Size *float32 `json:"size,omitempty"`
}

// NewReadRiskPolicySets200Response instantiates a new ReadRiskPolicySets200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadRiskPolicySets200Response() *ReadRiskPolicySets200Response {
	this := ReadRiskPolicySets200Response{}
	return &this
}

// NewReadRiskPolicySets200ResponseWithDefaults instantiates a new ReadRiskPolicySets200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadRiskPolicySets200ResponseWithDefaults() *ReadRiskPolicySets200Response {
	this := ReadRiskPolicySets200Response{}
	return &this
}

func (o ReadRiskPolicySets200Response) hasHalLink(linkIndex string) bool {
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

func (o ReadRiskPolicySets200Response) getHalLink(linkIndex string) LinksHATEOASValue {
	if l, ok := o.GetLinksOk(); ok && l != nil {
		links := l
		if v, ok := links[linkIndex]; ok {
			return v
		}
	}

	var ret LinksHATEOASValue
	return ret
}

func (o ReadRiskPolicySets200Response) getHalLinkOk(linkIndex string) (*LinksHATEOASValue, bool) {
	if l, ok := o.GetLinksOk(); ok && l != nil {
		links := l
		if v, ok := links[linkIndex]; ok {
			return &v, true
		}
	}

	return nil, false
}

func (o ReadRiskPolicySets200Response) IsPaginated() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_NEXT) || o.hasHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o ReadRiskPolicySets200Response) HasPaginationSelf() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o ReadRiskPolicySets200Response) GetPaginationSelfLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o ReadRiskPolicySets200Response) GetPaginationSelfLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o ReadRiskPolicySets200Response) HasPaginationNext() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o ReadRiskPolicySets200Response) GetPaginationNextLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o ReadRiskPolicySets200Response) GetPaginationNextLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o ReadRiskPolicySets200Response) HasPaginationPrevious() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o ReadRiskPolicySets200Response) GetPaginationPreviousLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o ReadRiskPolicySets200Response) GetPaginationPreviousLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_PREV)
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ReadRiskPolicySets200Response) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadRiskPolicySets200Response) GetLinksOk() (map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return map[string]LinksHATEOASValue{}, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ReadRiskPolicySets200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *ReadRiskPolicySets200Response) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *ReadRiskPolicySets200Response) GetEmbedded() ReadRiskPolicySets200ResponseEmbedded {
	if o == nil || IsNil(o.Embedded) {
		var ret ReadRiskPolicySets200ResponseEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadRiskPolicySets200Response) GetEmbeddedOk() (*ReadRiskPolicySets200ResponseEmbedded, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *ReadRiskPolicySets200Response) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given ReadRiskPolicySets200ResponseEmbedded and assigns it to the Embedded field.
func (o *ReadRiskPolicySets200Response) SetEmbedded(v ReadRiskPolicySets200ResponseEmbedded) {
	o.Embedded = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ReadRiskPolicySets200Response) GetCount() float32 {
	if o == nil || IsNil(o.Count) {
		var ret float32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadRiskPolicySets200Response) GetCountOk() (*float32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ReadRiskPolicySets200Response) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given float32 and assigns it to the Count field.
func (o *ReadRiskPolicySets200Response) SetCount(v float32) {
	o.Count = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ReadRiskPolicySets200Response) GetSize() float32 {
	if o == nil || IsNil(o.Size) {
		var ret float32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadRiskPolicySets200Response) GetSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ReadRiskPolicySets200Response) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given float32 and assigns it to the Size field.
func (o *ReadRiskPolicySets200Response) SetSize(v float32) {
	o.Size = &v
}

func (o ReadRiskPolicySets200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadRiskPolicySets200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableReadRiskPolicySets200Response struct {
	value *ReadRiskPolicySets200Response
	isSet bool
}

func (v NullableReadRiskPolicySets200Response) Get() *ReadRiskPolicySets200Response {
	return v.value
}

func (v *NullableReadRiskPolicySets200Response) Set(val *ReadRiskPolicySets200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableReadRiskPolicySets200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableReadRiskPolicySets200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadRiskPolicySets200Response(val *ReadRiskPolicySets200Response) *NullableReadRiskPolicySets200Response {
	return &NullableReadRiskPolicySets200Response{value: val, isSet: true}
}

func (v NullableReadRiskPolicySets200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadRiskPolicySets200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


