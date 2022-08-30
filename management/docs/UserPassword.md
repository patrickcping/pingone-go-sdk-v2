# UserPassword

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceChange** | Pointer to **bool** | A boolean that specifies whether the user is forced to change the password on the next log in. If not provided, the property is set to false. | [optional] 
**Value** | Pointer to **string** | A string that specifies the user&#39;s password value. The string is either in cleartext or pre-encoded format. | [optional] 
**External** | Pointer to [**UserPasswordExternal**](UserPasswordExternal.md) |  | [optional] 

## Methods

### NewUserPassword

`func NewUserPassword() *UserPassword`

NewUserPassword instantiates a new UserPassword object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPasswordWithDefaults

`func NewUserPasswordWithDefaults() *UserPassword`

NewUserPasswordWithDefaults instantiates a new UserPassword object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceChange

`func (o *UserPassword) GetForceChange() bool`

GetForceChange returns the ForceChange field if non-nil, zero value otherwise.

### GetForceChangeOk

`func (o *UserPassword) GetForceChangeOk() (*bool, bool)`

GetForceChangeOk returns a tuple with the ForceChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceChange

`func (o *UserPassword) SetForceChange(v bool)`

SetForceChange sets ForceChange field to given value.

### HasForceChange

`func (o *UserPassword) HasForceChange() bool`

HasForceChange returns a boolean if a field has been set.

### GetValue

`func (o *UserPassword) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UserPassword) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UserPassword) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *UserPassword) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetExternal

`func (o *UserPassword) GetExternal() UserPasswordExternal`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *UserPassword) GetExternalOk() (*UserPasswordExternal, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *UserPassword) SetExternal(v UserPasswordExternal)`

SetExternal sets External field to given value.

### HasExternal

`func (o *UserPassword) HasExternal() bool`

HasExternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


