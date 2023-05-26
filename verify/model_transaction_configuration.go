/*
PingOne Platform API - PingOne Verify

The PingOne Platform API covering the PingOne Verify service

API version: 2023-05-26
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package verify

import (
	"encoding/json"
)

// checks if the TransactionConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionConfiguration{}

// TransactionConfiguration struct for TransactionConfiguration
type TransactionConfiguration struct {
	Transaction *TransactionConfigurationTransaction `json:"transaction,omitempty"`
}

// NewTransactionConfiguration instantiates a new TransactionConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionConfiguration() *TransactionConfiguration {
	this := TransactionConfiguration{}
	return &this
}

// NewTransactionConfigurationWithDefaults instantiates a new TransactionConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionConfigurationWithDefaults() *TransactionConfiguration {
	this := TransactionConfiguration{}
	return &this
}

// GetTransaction returns the Transaction field value if set, zero value otherwise.
func (o *TransactionConfiguration) GetTransaction() TransactionConfigurationTransaction {
	if o == nil || IsNil(o.Transaction) {
		var ret TransactionConfigurationTransaction
		return ret
	}
	return *o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionConfiguration) GetTransactionOk() (*TransactionConfigurationTransaction, bool) {
	if o == nil || IsNil(o.Transaction) {
		return nil, false
	}
	return o.Transaction, true
}

// HasTransaction returns a boolean if a field has been set.
func (o *TransactionConfiguration) HasTransaction() bool {
	if o != nil && !IsNil(o.Transaction) {
		return true
	}

	return false
}

// SetTransaction gets a reference to the given TransactionConfigurationTransaction and assigns it to the Transaction field.
func (o *TransactionConfiguration) SetTransaction(v TransactionConfigurationTransaction) {
	o.Transaction = &v
}

func (o TransactionConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Transaction) {
		toSerialize["transaction"] = o.Transaction
	}
	return toSerialize, nil
}

type NullableTransactionConfiguration struct {
	value *TransactionConfiguration
	isSet bool
}

func (v NullableTransactionConfiguration) Get() *TransactionConfiguration {
	return v.value
}

func (v *NullableTransactionConfiguration) Set(val *TransactionConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionConfiguration(val *TransactionConfiguration) *NullableTransactionConfiguration {
	return &NullableTransactionConfiguration{value: val, isSet: true}
}

func (v NullableTransactionConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


