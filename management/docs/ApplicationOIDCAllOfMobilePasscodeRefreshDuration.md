# ApplicationOIDCAllOfMobilePasscodeRefreshDuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | The amount of time a passcode should be displayed before being replaced with a new passcode - must be between 30 and 60. | [default to 30]
**TimeUnit** | [**EnumPasscodeRefreshTimeUnit**](EnumPasscodeRefreshTimeUnit.md) |  | [default to ENUMPASSCODEREFRESHTIMEUNIT_SECONDS]

## Methods

### NewApplicationOIDCAllOfMobilePasscodeRefreshDuration

`func NewApplicationOIDCAllOfMobilePasscodeRefreshDuration(duration int32, timeUnit EnumPasscodeRefreshTimeUnit, ) *ApplicationOIDCAllOfMobilePasscodeRefreshDuration`

NewApplicationOIDCAllOfMobilePasscodeRefreshDuration instantiates a new ApplicationOIDCAllOfMobilePasscodeRefreshDuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationOIDCAllOfMobilePasscodeRefreshDurationWithDefaults

`func NewApplicationOIDCAllOfMobilePasscodeRefreshDurationWithDefaults() *ApplicationOIDCAllOfMobilePasscodeRefreshDuration`

NewApplicationOIDCAllOfMobilePasscodeRefreshDurationWithDefaults instantiates a new ApplicationOIDCAllOfMobilePasscodeRefreshDuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *ApplicationOIDCAllOfMobilePasscodeRefreshDuration) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ApplicationOIDCAllOfMobilePasscodeRefreshDuration) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ApplicationOIDCAllOfMobilePasscodeRefreshDuration) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetTimeUnit

`func (o *ApplicationOIDCAllOfMobilePasscodeRefreshDuration) GetTimeUnit() EnumPasscodeRefreshTimeUnit`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *ApplicationOIDCAllOfMobilePasscodeRefreshDuration) GetTimeUnitOk() (*EnumPasscodeRefreshTimeUnit, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *ApplicationOIDCAllOfMobilePasscodeRefreshDuration) SetTimeUnit(v EnumPasscodeRefreshTimeUnit)`

SetTimeUnit sets TimeUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


