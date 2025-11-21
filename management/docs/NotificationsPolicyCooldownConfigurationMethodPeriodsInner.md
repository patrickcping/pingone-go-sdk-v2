# NotificationsPolicyCooldownConfigurationMethodPeriodsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | Used in conjunction with timeUnit to specify the waiting period. Minimum is ten seconds, maximum is ten minutes. | 
**TimeUnit** | [**EnumNotificationsPolicyCooldownConfigurationMethodPeriodTimeUnit**](EnumNotificationsPolicyCooldownConfigurationMethodPeriodTimeUnit.md) |  | 

## Methods

### NewNotificationsPolicyCooldownConfigurationMethodPeriodsInner

`func NewNotificationsPolicyCooldownConfigurationMethodPeriodsInner(duration int32, timeUnit EnumNotificationsPolicyCooldownConfigurationMethodPeriodTimeUnit, ) *NotificationsPolicyCooldownConfigurationMethodPeriodsInner`

NewNotificationsPolicyCooldownConfigurationMethodPeriodsInner instantiates a new NotificationsPolicyCooldownConfigurationMethodPeriodsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsPolicyCooldownConfigurationMethodPeriodsInnerWithDefaults

`func NewNotificationsPolicyCooldownConfigurationMethodPeriodsInnerWithDefaults() *NotificationsPolicyCooldownConfigurationMethodPeriodsInner`

NewNotificationsPolicyCooldownConfigurationMethodPeriodsInnerWithDefaults instantiates a new NotificationsPolicyCooldownConfigurationMethodPeriodsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *NotificationsPolicyCooldownConfigurationMethodPeriodsInner) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *NotificationsPolicyCooldownConfigurationMethodPeriodsInner) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *NotificationsPolicyCooldownConfigurationMethodPeriodsInner) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetTimeUnit

`func (o *NotificationsPolicyCooldownConfigurationMethodPeriodsInner) GetTimeUnit() EnumNotificationsPolicyCooldownConfigurationMethodPeriodTimeUnit`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *NotificationsPolicyCooldownConfigurationMethodPeriodsInner) GetTimeUnitOk() (*EnumNotificationsPolicyCooldownConfigurationMethodPeriodTimeUnit, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *NotificationsPolicyCooldownConfigurationMethodPeriodsInner) SetTimeUnit(v EnumNotificationsPolicyCooldownConfigurationMethodPeriodTimeUnit)`

SetTimeUnit sets TimeUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


