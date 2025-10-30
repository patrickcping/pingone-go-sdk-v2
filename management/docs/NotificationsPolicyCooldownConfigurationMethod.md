# NotificationsPolicyCooldownConfigurationMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Set to true if you want to specify notification cooldown periods for the authentication method. Set to false if you don&#39;t want notification cooldown periods for this authentication method. | 
**Periods** | Pointer to [**[]NotificationsPolicyCooldownConfigurationMethodPeriodsInner**](NotificationsPolicyCooldownConfigurationMethodPeriodsInner.md) | Use the periods array to specify the amount of time the user has to wait before requesting another notification such as another OTP. The array should contain three objects: the time to wait before the first retry, the time to wait before the second retry, and the time to wait before any subsequent retries.  | [optional] 
**GroupBy** | Pointer to **string** | Since bad actors may try to target multiple users at a single email address or phone number, by default the cooldown settings (both waiting period and maximum retries) are applied to the email address or phone number. If you want the settings to be applied at the single-user level for the address/number, include the groupBy parameter in the request and set it to USER_ID. | [optional] 
**ResendLimit** | Pointer to **int32** | The maximum number of requests that a user can send to receive another notification, such as another OTP, before they are blocked for 30 minutes. | [optional] 

## Methods

### NewNotificationsPolicyCooldownConfigurationMethod

`func NewNotificationsPolicyCooldownConfigurationMethod(enabled bool, ) *NotificationsPolicyCooldownConfigurationMethod`

NewNotificationsPolicyCooldownConfigurationMethod instantiates a new NotificationsPolicyCooldownConfigurationMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsPolicyCooldownConfigurationMethodWithDefaults

`func NewNotificationsPolicyCooldownConfigurationMethodWithDefaults() *NotificationsPolicyCooldownConfigurationMethod`

NewNotificationsPolicyCooldownConfigurationMethodWithDefaults instantiates a new NotificationsPolicyCooldownConfigurationMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *NotificationsPolicyCooldownConfigurationMethod) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NotificationsPolicyCooldownConfigurationMethod) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NotificationsPolicyCooldownConfigurationMethod) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetPeriods

`func (o *NotificationsPolicyCooldownConfigurationMethod) GetPeriods() []NotificationsPolicyCooldownConfigurationMethodPeriodsInner`

GetPeriods returns the Periods field if non-nil, zero value otherwise.

### GetPeriodsOk

`func (o *NotificationsPolicyCooldownConfigurationMethod) GetPeriodsOk() (*[]NotificationsPolicyCooldownConfigurationMethodPeriodsInner, bool)`

GetPeriodsOk returns a tuple with the Periods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriods

`func (o *NotificationsPolicyCooldownConfigurationMethod) SetPeriods(v []NotificationsPolicyCooldownConfigurationMethodPeriodsInner)`

SetPeriods sets Periods field to given value.

### HasPeriods

`func (o *NotificationsPolicyCooldownConfigurationMethod) HasPeriods() bool`

HasPeriods returns a boolean if a field has been set.

### GetGroupBy

`func (o *NotificationsPolicyCooldownConfigurationMethod) GetGroupBy() string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *NotificationsPolicyCooldownConfigurationMethod) GetGroupByOk() (*string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *NotificationsPolicyCooldownConfigurationMethod) SetGroupBy(v string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *NotificationsPolicyCooldownConfigurationMethod) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetResendLimit

`func (o *NotificationsPolicyCooldownConfigurationMethod) GetResendLimit() int32`

GetResendLimit returns the ResendLimit field if non-nil, zero value otherwise.

### GetResendLimitOk

`func (o *NotificationsPolicyCooldownConfigurationMethod) GetResendLimitOk() (*int32, bool)`

GetResendLimitOk returns a tuple with the ResendLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResendLimit

`func (o *NotificationsPolicyCooldownConfigurationMethod) SetResendLimit(v int32)`

SetResendLimit sets ResendLimit field to given value.

### HasResendLimit

`func (o *NotificationsPolicyCooldownConfigurationMethod) HasResendLimit() bool`

HasResendLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


