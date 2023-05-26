# EmailPhoneConfigurationOtpDeliveries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Count of OTP deliveries. | 
**Cooldown** | [**EmailPhoneConfigurationOtpDeliveriesCooldown**](EmailPhoneConfigurationOtpDeliveriesCooldown.md) |  | 

## Methods

### NewEmailPhoneConfigurationOtpDeliveries

`func NewEmailPhoneConfigurationOtpDeliveries(count int32, cooldown EmailPhoneConfigurationOtpDeliveriesCooldown, ) *EmailPhoneConfigurationOtpDeliveries`

NewEmailPhoneConfigurationOtpDeliveries instantiates a new EmailPhoneConfigurationOtpDeliveries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailPhoneConfigurationOtpDeliveriesWithDefaults

`func NewEmailPhoneConfigurationOtpDeliveriesWithDefaults() *EmailPhoneConfigurationOtpDeliveries`

NewEmailPhoneConfigurationOtpDeliveriesWithDefaults instantiates a new EmailPhoneConfigurationOtpDeliveries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *EmailPhoneConfigurationOtpDeliveries) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *EmailPhoneConfigurationOtpDeliveries) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *EmailPhoneConfigurationOtpDeliveries) SetCount(v int32)`

SetCount sets Count field to given value.


### GetCooldown

`func (o *EmailPhoneConfigurationOtpDeliveries) GetCooldown() EmailPhoneConfigurationOtpDeliveriesCooldown`

GetCooldown returns the Cooldown field if non-nil, zero value otherwise.

### GetCooldownOk

`func (o *EmailPhoneConfigurationOtpDeliveries) GetCooldownOk() (*EmailPhoneConfigurationOtpDeliveriesCooldown, bool)`

GetCooldownOk returns a tuple with the Cooldown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldown

`func (o *EmailPhoneConfigurationOtpDeliveries) SetCooldown(v EmailPhoneConfigurationOtpDeliveriesCooldown)`

SetCooldown sets Cooldown field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


