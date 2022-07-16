# UserLifecycle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | A string that specifies information about the account lifecycle. Options for status are &#x60;ACCOUNT_OK&#x60; and &#x60;VERIFICATION_REQUIRED&#x60;. This property value is only allowed to be set when importing a user to set the initial account status. If the initial status is set to &#x60;VERIFICATION_REQUIRED&#x60; and an email address is provided, a verification email is sent. | [optional] 

## Methods

### NewUserLifecycle

`func NewUserLifecycle() *UserLifecycle`

NewUserLifecycle instantiates a new UserLifecycle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLifecycleWithDefaults

`func NewUserLifecycleWithDefaults() *UserLifecycle`

NewUserLifecycleWithDefaults instantiates a new UserLifecycle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UserLifecycle) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserLifecycle) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserLifecycle) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserLifecycle) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


