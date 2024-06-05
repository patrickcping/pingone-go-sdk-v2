# APIServerAccessControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Custom** | Pointer to [**APIServerAccessControlCustom**](APIServerAccessControlCustom.md) |  | [optional] 

## Methods

### NewAPIServerAccessControl

`func NewAPIServerAccessControl() *APIServerAccessControl`

NewAPIServerAccessControl instantiates a new APIServerAccessControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIServerAccessControlWithDefaults

`func NewAPIServerAccessControlWithDefaults() *APIServerAccessControl`

NewAPIServerAccessControlWithDefaults instantiates a new APIServerAccessControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustom

`func (o *APIServerAccessControl) GetCustom() APIServerAccessControlCustom`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *APIServerAccessControl) GetCustomOk() (*APIServerAccessControlCustom, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *APIServerAccessControl) SetCustom(v APIServerAccessControlCustom)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *APIServerAccessControl) HasCustom() bool`

HasCustom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


