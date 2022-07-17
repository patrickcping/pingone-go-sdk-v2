# MFAPushCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumMFAPushCredentialAttrType**](EnumMFAPushCredentialAttrType.md) |  | 
**Key** | **string** | A string that Apple uses as an identifier to identify an authentication key.  Mandatory. | 

## Methods

### NewMFAPushCredential

`func NewMFAPushCredential(type_ EnumMFAPushCredentialAttrType, key string, ) *MFAPushCredential`

NewMFAPushCredential instantiates a new MFAPushCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFAPushCredentialWithDefaults

`func NewMFAPushCredentialWithDefaults() *MFAPushCredential`

NewMFAPushCredentialWithDefaults instantiates a new MFAPushCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MFAPushCredential) GetType() EnumMFAPushCredentialAttrType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MFAPushCredential) GetTypeOk() (*EnumMFAPushCredentialAttrType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MFAPushCredential) SetType(v EnumMFAPushCredentialAttrType)`

SetType sets Type field to given value.


### GetKey

`func (o *MFAPushCredential) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MFAPushCredential) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MFAPushCredential) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


