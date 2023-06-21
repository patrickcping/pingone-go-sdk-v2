# FIDO2PolicyUserDisplayNameAttributesAttributesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the attribute to use for the user display name. | 
**SubAttributes** | Pointer to [**[]FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner**](FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner.md) | The sub-attributes to use for the user display name. | [optional] 

## Methods

### NewFIDO2PolicyUserDisplayNameAttributesAttributesInner

`func NewFIDO2PolicyUserDisplayNameAttributesAttributesInner(name string, ) *FIDO2PolicyUserDisplayNameAttributesAttributesInner`

NewFIDO2PolicyUserDisplayNameAttributesAttributesInner instantiates a new FIDO2PolicyUserDisplayNameAttributesAttributesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFIDO2PolicyUserDisplayNameAttributesAttributesInnerWithDefaults

`func NewFIDO2PolicyUserDisplayNameAttributesAttributesInnerWithDefaults() *FIDO2PolicyUserDisplayNameAttributesAttributesInner`

NewFIDO2PolicyUserDisplayNameAttributesAttributesInnerWithDefaults instantiates a new FIDO2PolicyUserDisplayNameAttributesAttributesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FIDO2PolicyUserDisplayNameAttributesAttributesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FIDO2PolicyUserDisplayNameAttributesAttributesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FIDO2PolicyUserDisplayNameAttributesAttributesInner) SetName(v string)`

SetName sets Name field to given value.


### GetSubAttributes

`func (o *FIDO2PolicyUserDisplayNameAttributesAttributesInner) GetSubAttributes() []FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner`

GetSubAttributes returns the SubAttributes field if non-nil, zero value otherwise.

### GetSubAttributesOk

`func (o *FIDO2PolicyUserDisplayNameAttributesAttributesInner) GetSubAttributesOk() (*[]FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner, bool)`

GetSubAttributesOk returns a tuple with the SubAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAttributes

`func (o *FIDO2PolicyUserDisplayNameAttributesAttributesInner) SetSubAttributes(v []FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner)`

SetSubAttributes sets SubAttributes field to given value.

### HasSubAttributes

`func (o *FIDO2PolicyUserDisplayNameAttributesAttributesInner) HasSubAttributes() bool`

HasSubAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


