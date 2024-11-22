/*
PingOne Platform API - Authorize

The PingOne Platform API covering the PingOne Authorize service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorize

import (
	"encoding/json"
)

// checks if the ReadApplicationResources200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadApplicationResources200Response{}

// ReadApplicationResources200Response struct for ReadApplicationResources200Response
type ReadApplicationResources200Response struct {
	Links map[string]LinksHATEOASValue `json:"_links,omitempty"`
	Embedded *ReadApplicationResources200ResponseEmbedded `json:"_embedded,omitempty"`
	Count *float32 `json:"count,omitempty"`
	Size *float32 `json:"size,omitempty"`
}

// NewReadApplicationResources200Response instantiates a new ReadApplicationResources200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadApplicationResources200Response() *ReadApplicationResources200Response {
	this := ReadApplicationResources200Response{}
	return &this
}

// NewReadApplicationResources200ResponseWithDefaults instantiates a new ReadApplicationResources200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadApplicationResources200ResponseWithDefaults() *ReadApplicationResources200Response {
	this := ReadApplicationResources200Response{}
	return &this
}

func (o ReadApplicationResources200Response) hasHalLink(linkIndex string) bool {
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

func (o ReadApplicationResources200Response) getHalLink(linkIndex string) LinksHATEOASValue {
	if l, ok := o.GetLinksOk(); ok && l != nil {
		links := l
		if v, ok := links[linkIndex]; ok {
			return v
		}
	}

	var ret LinksHATEOASValue
	return ret
}

func (o ReadApplicationResources200Response) getHalLinkOk(linkIndex string) (*LinksHATEOASValue, bool) {
	if l, ok := o.GetLinksOk(); ok && l != nil {
		links := l
		if v, ok := links[linkIndex]; ok {
			return &v, true
		}
	}

	return nil, false
}

func (o ReadApplicationResources200Response) IsPaginated() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_NEXT) || o.hasHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o ReadApplicationResources200Response) HasPaginationSelf() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o ReadApplicationResources200Response) GetPaginationSelfLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o ReadApplicationResources200Response) GetPaginationSelfLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o ReadApplicationResources200Response) HasPaginationNext() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o ReadApplicationResources200Response) GetPaginationNextLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o ReadApplicationResources200Response) GetPaginationNextLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o ReadApplicationResources200Response) HasPaginationPrevious() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o ReadApplicationResources200Response) GetPaginationPreviousLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o ReadApplicationResources200Response) GetPaginationPreviousLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_PREV)
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ReadApplicationResources200Response) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadApplicationResources200Response) GetLinksOk() (map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return map[string]LinksHATEOASValue{}, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ReadApplicationResources200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *ReadApplicationResources200Response) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *ReadApplicationResources200Response) GetEmbedded() ReadApplicationResources200ResponseEmbedded {
	if o == nil || IsNil(o.Embedded) {
		var ret ReadApplicationResources200ResponseEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadApplicationResources200Response) GetEmbeddedOk() (*ReadApplicationResources200ResponseEmbedded, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *ReadApplicationResources200Response) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given ReadApplicationResources200ResponseEmbedded and assigns it to the Embedded field.
func (o *ReadApplicationResources200Response) SetEmbedded(v ReadApplicationResources200ResponseEmbedded) {
	o.Embedded = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ReadApplicationResources200Response) GetCount() float32 {
	if o == nil || IsNil(o.Count) {
		var ret float32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadApplicationResources200Response) GetCountOk() (*float32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ReadApplicationResources200Response) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given float32 and assigns it to the Count field.
func (o *ReadApplicationResources200Response) SetCount(v float32) {
	o.Count = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ReadApplicationResources200Response) GetSize() float32 {
	if o == nil || IsNil(o.Size) {
		var ret float32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadApplicationResources200Response) GetSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ReadApplicationResources200Response) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given float32 and assigns it to the Size field.
func (o *ReadApplicationResources200Response) SetSize(v float32) {
	o.Size = &v
}

func (o ReadApplicationResources200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadApplicationResources200Response) ToMap() (map[string]interface{}, error) {
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

type NullableReadApplicationResources200Response struct {
	value *ReadApplicationResources200Response
	isSet bool
}

func (v NullableReadApplicationResources200Response) Get() *ReadApplicationResources200Response {
	return v.value
}

func (v *NullableReadApplicationResources200Response) Set(val *ReadApplicationResources200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableReadApplicationResources200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableReadApplicationResources200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadApplicationResources200Response(val *ReadApplicationResources200Response) *NullableReadApplicationResources200Response {
	return &NullableReadApplicationResources200Response{value: val, isSet: true}
}

func (v NullableReadApplicationResources200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadApplicationResources200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


