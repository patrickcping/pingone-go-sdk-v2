# NotificationsPolicyQuotasInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumNotificationsPolicyQuotaItemType**](EnumNotificationsPolicyQuotaItemType.md) |  | 
**DeliveryMethods** | [**[]EnumNotificationsPolicyQuotaDeliveryMethods**](EnumNotificationsPolicyQuotaDeliveryMethods.md) | The delivery methods for which the limit is being defined. The value can be &#x60;Email&#x60; or &#x60;SMS,Voice&#x60;. When you use the &#x60;SMS&#x60;, &#x60;Voice&#x60; option, it means that the combined total of SMS and voice notifications must be below the limit defined. If you are limiting both email and SMS/voice, each limit should be represented by a different object in the &#x60;quotas&#x60; array, for example: &#x60;\&quot;quotas\&quot;: [{\&quot;type\&quot;: \&quot;USER\&quot;,\&quot;deliveryMethods\&quot;: [\&quot;SMS\&quot;,\&quot;Voice\&quot;],\&quot;total\&quot;: 30},{\&quot;type\&quot;: \&quot;USER\&quot;,\&quot;deliveryMethods\&quot;: [\&quot;Email\&quot;],\&quot;total\&quot;: 30}]&#x60;  | 
**Total** | Pointer to **int32** | The maximum number of notifications allowed per day. | [optional] 
**Claimed** | Pointer to **int32** | The maximum number of notifications that can be received and responded to each day. Used in conjunction with unclaimed in place of the single field total. | [optional] 
**Unclaimed** | Pointer to **int32** | The maximum number of notifications that can be received and not responded to each day. Used in conjunction with claimed in place of the single field total. | [optional] 

## Methods

### NewNotificationsPolicyQuotasInner

`func NewNotificationsPolicyQuotasInner(type_ EnumNotificationsPolicyQuotaItemType, deliveryMethods []EnumNotificationsPolicyQuotaDeliveryMethods, ) *NotificationsPolicyQuotasInner`

NewNotificationsPolicyQuotasInner instantiates a new NotificationsPolicyQuotasInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsPolicyQuotasInnerWithDefaults

`func NewNotificationsPolicyQuotasInnerWithDefaults() *NotificationsPolicyQuotasInner`

NewNotificationsPolicyQuotasInnerWithDefaults instantiates a new NotificationsPolicyQuotasInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotificationsPolicyQuotasInner) GetType() EnumNotificationsPolicyQuotaItemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationsPolicyQuotasInner) GetTypeOk() (*EnumNotificationsPolicyQuotaItemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationsPolicyQuotasInner) SetType(v EnumNotificationsPolicyQuotaItemType)`

SetType sets Type field to given value.


### GetDeliveryMethods

`func (o *NotificationsPolicyQuotasInner) GetDeliveryMethods() []EnumNotificationsPolicyQuotaDeliveryMethods`

GetDeliveryMethods returns the DeliveryMethods field if non-nil, zero value otherwise.

### GetDeliveryMethodsOk

`func (o *NotificationsPolicyQuotasInner) GetDeliveryMethodsOk() (*[]EnumNotificationsPolicyQuotaDeliveryMethods, bool)`

GetDeliveryMethodsOk returns a tuple with the DeliveryMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethods

`func (o *NotificationsPolicyQuotasInner) SetDeliveryMethods(v []EnumNotificationsPolicyQuotaDeliveryMethods)`

SetDeliveryMethods sets DeliveryMethods field to given value.


### GetTotal

`func (o *NotificationsPolicyQuotasInner) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *NotificationsPolicyQuotasInner) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *NotificationsPolicyQuotasInner) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *NotificationsPolicyQuotasInner) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetClaimed

`func (o *NotificationsPolicyQuotasInner) GetClaimed() int32`

GetClaimed returns the Claimed field if non-nil, zero value otherwise.

### GetClaimedOk

`func (o *NotificationsPolicyQuotasInner) GetClaimedOk() (*int32, bool)`

GetClaimedOk returns a tuple with the Claimed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimed

`func (o *NotificationsPolicyQuotasInner) SetClaimed(v int32)`

SetClaimed sets Claimed field to given value.

### HasClaimed

`func (o *NotificationsPolicyQuotasInner) HasClaimed() bool`

HasClaimed returns a boolean if a field has been set.

### GetUnclaimed

`func (o *NotificationsPolicyQuotasInner) GetUnclaimed() int32`

GetUnclaimed returns the Unclaimed field if non-nil, zero value otherwise.

### GetUnclaimedOk

`func (o *NotificationsPolicyQuotasInner) GetUnclaimedOk() (*int32, bool)`

GetUnclaimedOk returns a tuple with the Unclaimed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnclaimed

`func (o *NotificationsPolicyQuotasInner) SetUnclaimed(v int32)`

SetUnclaimed sets Unclaimed field to given value.

### HasUnclaimed

`func (o *NotificationsPolicyQuotasInner) HasUnclaimed() bool`

HasUnclaimed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


