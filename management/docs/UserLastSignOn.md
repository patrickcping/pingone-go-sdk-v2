# UserLastSignOn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | The time of the last successful login of the user through the PingOne flow API. | [optional] 
**RemoteIp** | Pointer to **string** | The IP address of the last successful login of the user through the PingOne flow API. | [optional] 

## Methods

### NewUserLastSignOn

`func NewUserLastSignOn() *UserLastSignOn`

NewUserLastSignOn instantiates a new UserLastSignOn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLastSignOnWithDefaults

`func NewUserLastSignOnWithDefaults() *UserLastSignOn`

NewUserLastSignOnWithDefaults instantiates a new UserLastSignOn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *UserLastSignOn) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *UserLastSignOn) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *UserLastSignOn) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *UserLastSignOn) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetRemoteIp

`func (o *UserLastSignOn) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *UserLastSignOn) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *UserLastSignOn) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.

### HasRemoteIp

`func (o *UserLastSignOn) HasRemoteIp() bool`

HasRemoteIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


