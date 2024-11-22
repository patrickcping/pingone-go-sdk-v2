/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
)

// checks if the ReadAllMFAPushCredentials200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadAllMFAPushCredentials200Response{}

// ReadAllMFAPushCredentials200Response struct for ReadAllMFAPushCredentials200Response
type ReadAllMFAPushCredentials200Response struct {
	Links map[string]LinksHATEOASValue `json:"_links,omitempty"`
	Embedded *ReadAllMFAPushCredentials200ResponseEmbedded `json:"_embedded,omitempty"`
	Count *float32 `json:"count,omitempty"`
	Size *float32 `json:"size,omitempty"`
}

// NewReadAllMFAPushCredentials200Response instantiates a new ReadAllMFAPushCredentials200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadAllMFAPushCredentials200Response() *ReadAllMFAPushCredentials200Response {
	this := ReadAllMFAPushCredentials200Response{}
	return &this
}

// NewReadAllMFAPushCredentials200ResponseWithDefaults instantiates a new ReadAllMFAPushCredentials200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadAllMFAPushCredentials200ResponseWithDefaults() *ReadAllMFAPushCredentials200Response {
	this := ReadAllMFAPushCredentials200Response{}
	return &this
}

func (o ReadAllMFAPushCredentials200Response) hasHalLink(linkIndex string) bool {
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

func (o ReadAllMFAPushCredentials200Response) getHalLink(linkIndex string) LinksHATEOASValue {
	if l, ok := o.GetLinksOk(); ok && l != nil {
		links := l
		if v, ok := links[linkIndex]; ok {
			return v
		}
	}

	var ret LinksHATEOASValue
	return ret
}

func (o ReadAllMFAPushCredentials200Response) getHalLinkOk(linkIndex string) (*LinksHATEOASValue, bool) {
	if l, ok := o.GetLinksOk(); ok && l != nil {
		links := l
		if v, ok := links[linkIndex]; ok {
			return &v, true
		}
	}

	return nil, false
}

func (o ReadAllMFAPushCredentials200Response) IsPaginated() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_NEXT) || o.hasHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o ReadAllMFAPushCredentials200Response) HasPaginationSelf() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o ReadAllMFAPushCredentials200Response) GetPaginationSelfLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o ReadAllMFAPushCredentials200Response) GetPaginationSelfLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o ReadAllMFAPushCredentials200Response) HasPaginationNext() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o ReadAllMFAPushCredentials200Response) GetPaginationNextLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o ReadAllMFAPushCredentials200Response) GetPaginationNextLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o ReadAllMFAPushCredentials200Response) HasPaginationPrevious() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o ReadAllMFAPushCredentials200Response) GetPaginationPreviousLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o ReadAllMFAPushCredentials200Response) GetPaginationPreviousLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_PREV)
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ReadAllMFAPushCredentials200Response) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadAllMFAPushCredentials200Response) GetLinksOk() (map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return map[string]LinksHATEOASValue{}, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ReadAllMFAPushCredentials200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *ReadAllMFAPushCredentials200Response) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *ReadAllMFAPushCredentials200Response) GetEmbedded() ReadAllMFAPushCredentials200ResponseEmbedded {
	if o == nil || IsNil(o.Embedded) {
		var ret ReadAllMFAPushCredentials200ResponseEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadAllMFAPushCredentials200Response) GetEmbeddedOk() (*ReadAllMFAPushCredentials200ResponseEmbedded, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *ReadAllMFAPushCredentials200Response) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given ReadAllMFAPushCredentials200ResponseEmbedded and assigns it to the Embedded field.
func (o *ReadAllMFAPushCredentials200Response) SetEmbedded(v ReadAllMFAPushCredentials200ResponseEmbedded) {
	o.Embedded = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ReadAllMFAPushCredentials200Response) GetCount() float32 {
	if o == nil || IsNil(o.Count) {
		var ret float32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadAllMFAPushCredentials200Response) GetCountOk() (*float32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ReadAllMFAPushCredentials200Response) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given float32 and assigns it to the Count field.
func (o *ReadAllMFAPushCredentials200Response) SetCount(v float32) {
	o.Count = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ReadAllMFAPushCredentials200Response) GetSize() float32 {
	if o == nil || IsNil(o.Size) {
		var ret float32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadAllMFAPushCredentials200Response) GetSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ReadAllMFAPushCredentials200Response) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given float32 and assigns it to the Size field.
func (o *ReadAllMFAPushCredentials200Response) SetSize(v float32) {
	o.Size = &v
}

func (o ReadAllMFAPushCredentials200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadAllMFAPushCredentials200Response) ToMap() (map[string]interface{}, error) {
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

type NullableReadAllMFAPushCredentials200Response struct {
	value *ReadAllMFAPushCredentials200Response
	isSet bool
}

func (v NullableReadAllMFAPushCredentials200Response) Get() *ReadAllMFAPushCredentials200Response {
	return v.value
}

func (v *NullableReadAllMFAPushCredentials200Response) Set(val *ReadAllMFAPushCredentials200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableReadAllMFAPushCredentials200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableReadAllMFAPushCredentials200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadAllMFAPushCredentials200Response(val *ReadAllMFAPushCredentials200Response) *NullableReadAllMFAPushCredentials200Response {
	return &NullableReadAllMFAPushCredentials200Response{value: val, isSet: true}
}

func (v NullableReadAllMFAPushCredentials200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadAllMFAPushCredentials200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


