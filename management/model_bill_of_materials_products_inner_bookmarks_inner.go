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

// checks if the BillOfMaterialsProductsInnerBookmarksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillOfMaterialsProductsInnerBookmarksInner{}

// BillOfMaterialsProductsInnerBookmarksInner struct for BillOfMaterialsProductsInnerBookmarksInner
type BillOfMaterialsProductsInnerBookmarksInner struct {
	// Name of the custom bookmark. The name must be unique among the product bookmarks and be 50 characters or fewer.
	Name string `json:"name"`
	// A valid URL for the bookmark.
	Href string `json:"href"`
}

// NewBillOfMaterialsProductsInnerBookmarksInner instantiates a new BillOfMaterialsProductsInnerBookmarksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillOfMaterialsProductsInnerBookmarksInner(name string, href string) *BillOfMaterialsProductsInnerBookmarksInner {
	this := BillOfMaterialsProductsInnerBookmarksInner{}
	this.Name = name
	this.Href = href
	return &this
}

// NewBillOfMaterialsProductsInnerBookmarksInnerWithDefaults instantiates a new BillOfMaterialsProductsInnerBookmarksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillOfMaterialsProductsInnerBookmarksInnerWithDefaults() *BillOfMaterialsProductsInnerBookmarksInner {
	this := BillOfMaterialsProductsInnerBookmarksInner{}
	return &this
}

// GetName returns the Name field value
func (o *BillOfMaterialsProductsInnerBookmarksInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BillOfMaterialsProductsInnerBookmarksInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BillOfMaterialsProductsInnerBookmarksInner) SetName(v string) {
	o.Name = v
}

// GetHref returns the Href field value
func (o *BillOfMaterialsProductsInnerBookmarksInner) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *BillOfMaterialsProductsInnerBookmarksInner) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *BillOfMaterialsProductsInnerBookmarksInner) SetHref(v string) {
	o.Href = v
}

func (o BillOfMaterialsProductsInnerBookmarksInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillOfMaterialsProductsInnerBookmarksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["href"] = o.Href
	return toSerialize, nil
}

type NullableBillOfMaterialsProductsInnerBookmarksInner struct {
	value *BillOfMaterialsProductsInnerBookmarksInner
	isSet bool
}

func (v NullableBillOfMaterialsProductsInnerBookmarksInner) Get() *BillOfMaterialsProductsInnerBookmarksInner {
	return v.value
}

func (v *NullableBillOfMaterialsProductsInnerBookmarksInner) Set(val *BillOfMaterialsProductsInnerBookmarksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBillOfMaterialsProductsInnerBookmarksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBillOfMaterialsProductsInnerBookmarksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillOfMaterialsProductsInnerBookmarksInner(val *BillOfMaterialsProductsInnerBookmarksInner) *NullableBillOfMaterialsProductsInnerBookmarksInner {
	return &NullableBillOfMaterialsProductsInnerBookmarksInner{value: val, isSet: true}
}

func (v NullableBillOfMaterialsProductsInnerBookmarksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillOfMaterialsProductsInnerBookmarksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
