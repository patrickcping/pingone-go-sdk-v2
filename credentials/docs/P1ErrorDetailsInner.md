# P1ErrorDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | A general fault code which the client must handle to provide all exception handling routines and to localize messages for users. This code is common across all PingOne services and is human readable (such as a defined constant rather than a number). | [optional] 
**Target** | Pointer to **string** | The item that caused the error (such as a form field ID or an attribute inside a JSON object). | [optional] 
**Message** | Pointer to **string** | A short description of the error. This message is intended to assist with debugging and is returned in English only. | [optional] 
**InnerError** | Pointer to [**P1ErrorDetailsInnerInnerError**](P1ErrorDetailsInnerInnerError.md) |  | [optional] 

## Methods

### NewP1ErrorDetailsInner

`func NewP1ErrorDetailsInner() *P1ErrorDetailsInner`

NewP1ErrorDetailsInner instantiates a new P1ErrorDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewP1ErrorDetailsInnerWithDefaults

`func NewP1ErrorDetailsInnerWithDefaults() *P1ErrorDetailsInner`

NewP1ErrorDetailsInnerWithDefaults instantiates a new P1ErrorDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *P1ErrorDetailsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *P1ErrorDetailsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *P1ErrorDetailsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *P1ErrorDetailsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetTarget

`func (o *P1ErrorDetailsInner) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *P1ErrorDetailsInner) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *P1ErrorDetailsInner) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *P1ErrorDetailsInner) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetMessage

`func (o *P1ErrorDetailsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *P1ErrorDetailsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *P1ErrorDetailsInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *P1ErrorDetailsInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetInnerError

`func (o *P1ErrorDetailsInner) GetInnerError() P1ErrorDetailsInnerInnerError`

GetInnerError returns the InnerError field if non-nil, zero value otherwise.

### GetInnerErrorOk

`func (o *P1ErrorDetailsInner) GetInnerErrorOk() (*P1ErrorDetailsInnerInnerError, bool)`

GetInnerErrorOk returns a tuple with the InnerError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerError

`func (o *P1ErrorDetailsInner) SetInnerError(v P1ErrorDetailsInnerInnerError)`

SetInnerError sets InnerError field to given value.

### HasInnerError

`func (o *P1ErrorDetailsInner) HasInnerError() bool`

HasInnerError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


