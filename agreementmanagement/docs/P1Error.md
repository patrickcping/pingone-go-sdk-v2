# P1Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier that is stored in log files and always included in an error response. This value can be used to track the error received by the client, with server-side activity included for troubleshooting purposes. | [optional] 
**Code** | Pointer to **string** | A general fault code which the client must handle to provide all exception handling routines and to localize messages for users. This code is common across all PingOne services and is human readable (such as a defined constant rather than a number). | [optional] 
**Message** | Pointer to **string** | A short description of the error. This message is intended to assist with debugging and is returned in English only. | [optional] 
**Details** | Pointer to [**[]P1ErrorDetailsInner**](P1ErrorDetailsInner.md) | Additional details about the error. Optional information to help resolve the error and to display to users. | [optional] 

## Methods

### NewP1Error

`func NewP1Error() *P1Error`

NewP1Error instantiates a new P1Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewP1ErrorWithDefaults

`func NewP1ErrorWithDefaults() *P1Error`

NewP1ErrorWithDefaults instantiates a new P1Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *P1Error) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *P1Error) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *P1Error) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *P1Error) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *P1Error) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *P1Error) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *P1Error) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *P1Error) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *P1Error) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *P1Error) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *P1Error) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *P1Error) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetails

`func (o *P1Error) GetDetails() []P1ErrorDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *P1Error) GetDetailsOk() (*[]P1ErrorDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *P1Error) SetDetails(v []P1ErrorDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *P1Error) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


