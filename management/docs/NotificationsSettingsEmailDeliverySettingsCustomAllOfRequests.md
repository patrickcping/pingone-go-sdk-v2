# NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Required if method is set to &#x60;POST&#x60;. Use body to provide the content of the body for the request sent to the email provider. The body that you define must include the mandatory PingOne variables for notifications: &#x60;${from}&#x60;, &#x60;${to}&#x60;, and &#x60;${message}&#x60;. You can also include any optional notification variables. If you are including the header &#x60;content-type: application/json&#x60; in &#x60;headers&#x60;: - The content of &#x60;body&#x60; must be valid JSON - You must enclose the JSON object in quotation marks and escape all quotation marks within the string, for example, &#x60;\&quot;{\\\&quot;From\\\&quot;: \\\&quot;${from}\\\&quot;,\\\&quot;To\\\&quot;: \\\&quot;${to}\\\&quot;,\\\&quot;message2\\\&quot;: \\\&quot;${message}\\\&quot;}\&quot;&#x60;.  | [optional] 
**DeliveryMethod** | [**EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod**](EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod.md) |  | 
**Headers** | Pointer to **map[string]string** | Use this object to specify the headers that your email provider&#39;s API expects. | [optional] 
**Method** | [**EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsMethod**](EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsMethod.md) |  | 
**Url** | **string** | Use url to specify the endpoint for your email provider, for example, &#x60;https://api.example.com/email&#x60;. If &#x60;method&#x60; is set to &#x60;GET&#x60;, append to the URL the various query parameters that the email provider&#39;s API requires. The URL must also include the required PingOne variables, as described for the body parameter. | 

## Methods

### NewNotificationsSettingsEmailDeliverySettingsCustomAllOfRequests

`func NewNotificationsSettingsEmailDeliverySettingsCustomAllOfRequests(deliveryMethod EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod, method EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsMethod, url string, ) *NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests`

NewNotificationsSettingsEmailDeliverySettingsCustomAllOfRequests instantiates a new NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsSettingsEmailDeliverySettingsCustomAllOfRequestsWithDefaults

`func NewNotificationsSettingsEmailDeliverySettingsCustomAllOfRequestsWithDefaults() *NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests`

NewNotificationsSettingsEmailDeliverySettingsCustomAllOfRequestsWithDefaults instantiates a new NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetDeliveryMethod

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests) GetDeliveryMethod() EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod`

GetDeliveryMethod returns the DeliveryMethod field if non-nil, zero value otherwise.

### GetDeliveryMethodOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests) GetDeliveryMethodOk() (*EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod, bool)`

GetDeliveryMethodOk returns a tuple with the DeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethod

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests) SetDeliveryMethod(v EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod)`

SetDeliveryMethod sets DeliveryMethod field to given value.


### GetHeaders

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetMethod

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests) GetMethod() EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests) GetMethodOk() (*EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests) SetMethod(v EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsMethod)`

SetMethod sets Method field to given value.


### GetUrl

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


