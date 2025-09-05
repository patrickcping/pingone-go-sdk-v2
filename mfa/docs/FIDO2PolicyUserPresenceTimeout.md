# FIDO2PolicyUserPresenceTimeout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int32** | The amount of time a user presence gesture will be accepted for the authentication request. Minimum is one minute, maximum is ten minutes. Can be set in minutes or seconds. | [optional] 
**TimeUnit** | Pointer to [**EnumTimeUnit**](EnumTimeUnit.md) |  | [optional] 

## Methods

### NewFIDO2PolicyUserPresenceTimeout

`func NewFIDO2PolicyUserPresenceTimeout() *FIDO2PolicyUserPresenceTimeout`

NewFIDO2PolicyUserPresenceTimeout instantiates a new FIDO2PolicyUserPresenceTimeout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFIDO2PolicyUserPresenceTimeoutWithDefaults

`func NewFIDO2PolicyUserPresenceTimeoutWithDefaults() *FIDO2PolicyUserPresenceTimeout`

NewFIDO2PolicyUserPresenceTimeoutWithDefaults instantiates a new FIDO2PolicyUserPresenceTimeout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *FIDO2PolicyUserPresenceTimeout) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *FIDO2PolicyUserPresenceTimeout) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *FIDO2PolicyUserPresenceTimeout) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *FIDO2PolicyUserPresenceTimeout) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetTimeUnit

`func (o *FIDO2PolicyUserPresenceTimeout) GetTimeUnit() EnumTimeUnit`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *FIDO2PolicyUserPresenceTimeout) GetTimeUnitOk() (*EnumTimeUnit, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *FIDO2PolicyUserPresenceTimeout) SetTimeUnit(v EnumTimeUnit)`

SetTimeUnit sets TimeUnit field to given value.

### HasTimeUnit

`func (o *FIDO2PolicyUserPresenceTimeout) HasTimeUnit() bool`

HasTimeUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


