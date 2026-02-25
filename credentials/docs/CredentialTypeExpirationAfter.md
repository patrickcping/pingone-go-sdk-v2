# CredentialTypeExpirationAfter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | Length of time before transaction timeout expires; range is from a 1 hour minimum (or equivalent) to an (effectively) unbounded maximum. | 
**TimeUnit** | [**EnumTimeUnit**](EnumTimeUnit.md) |  | 

## Methods

### NewCredentialTypeExpirationAfter

`func NewCredentialTypeExpirationAfter(duration int32, timeUnit EnumTimeUnit, ) *CredentialTypeExpirationAfter`

NewCredentialTypeExpirationAfter instantiates a new CredentialTypeExpirationAfter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialTypeExpirationAfterWithDefaults

`func NewCredentialTypeExpirationAfterWithDefaults() *CredentialTypeExpirationAfter`

NewCredentialTypeExpirationAfterWithDefaults instantiates a new CredentialTypeExpirationAfter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *CredentialTypeExpirationAfter) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CredentialTypeExpirationAfter) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CredentialTypeExpirationAfter) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetTimeUnit

`func (o *CredentialTypeExpirationAfter) GetTimeUnit() EnumTimeUnit`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *CredentialTypeExpirationAfter) GetTimeUnitOk() (*EnumTimeUnit, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *CredentialTypeExpirationAfter) SetTimeUnit(v EnumTimeUnit)`

SetTimeUnit sets TimeUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


