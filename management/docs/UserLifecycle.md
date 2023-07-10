# UserLifecycle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**EnumUserLifecycleStatus**](EnumUserLifecycleStatus.md) |  | [optional] 
**SuppressVerificationCode** | Pointer to **bool** | Used when importing a user and the lifecycle.status is set to VERIFICATION_REQUIRED. If this property is set to true, no verification email is sent to the user. If this property is omitted or set to false, a verification email is sent automatically to the user. | [optional] 

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

`func (o *UserLifecycle) GetStatus() EnumUserLifecycleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserLifecycle) GetStatusOk() (*EnumUserLifecycleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserLifecycle) SetStatus(v EnumUserLifecycleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserLifecycle) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuppressVerificationCode

`func (o *UserLifecycle) GetSuppressVerificationCode() bool`

GetSuppressVerificationCode returns the SuppressVerificationCode field if non-nil, zero value otherwise.

### GetSuppressVerificationCodeOk

`func (o *UserLifecycle) GetSuppressVerificationCodeOk() (*bool, bool)`

GetSuppressVerificationCodeOk returns a tuple with the SuppressVerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressVerificationCode

`func (o *UserLifecycle) SetSuppressVerificationCode(v bool)`

SetSuppressVerificationCode sets SuppressVerificationCode field to given value.

### HasSuppressVerificationCode

`func (o *UserLifecycle) HasSuppressVerificationCode() bool`

HasSuppressVerificationCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


