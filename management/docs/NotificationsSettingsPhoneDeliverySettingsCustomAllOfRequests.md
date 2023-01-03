# NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryMethod** | [**EnumNotificationsSettingsPhoneDeliverySettingsCustomDeliveryMethod**](EnumNotificationsSettingsPhoneDeliverySettingsCustomDeliveryMethod.md) |  | 
**Url** | **string** | The provider&#39;s remote gateway or customer gateway URL. For requests using the POST method, use the provider&#39;s remote gateway URL. For requests using the GET method, use the provider&#39;s remote gateway URL, including the &#x60;${to}&#x60; and &#x60;${message}&#x60; mandatory variables, and the optional &#x60;${from}&#x60; variable, for example: &#x60;https://api.transmitsms.com/send-sms.json?to&#x3D;${to}&amp;from&#x3D;${from}&amp;message&#x3D;${message}&#x60;  | 
**Body** | **string** | The notification&#39;s request body. The body should include the ${to} and ${message} mandatory variables. For some vendors, the optional ${from} variable may also be required. For example: &#x60;messageType&#x3D;ARN&amp;message&#x3D;${message}&amp;phoneNumber&#x3D;${to}&amp;sender&#x3D;${from}&#x60; In addition, you can use the following optional variables: &#x60;${voice}&#x60; - the type of voice configured for notifications &#x60;${locale}&#x60; - locale &#x60;${otp}&#x60; - OTP &#x60;${user.user.name}&#x60; - user&#39;s username &#x60;${user.name.given}&#x60; - user&#39;s given name &#x60;${user.name.family}&#x60; - user&#39;s family name You can also use dynamic variables in the body. For more information, see [Dynamic variables](https://apidocs.pingidentity.com/pingone/platform/v1/api/#notifications-templates-dynamic-variables).  | 
**Headers** | **string** | The notification&#39;s request header, matching the format of the request body. The header can be one of: Form &#x60;(content-type&#x3D;application/x-www-form-urlencoded)&#x60; JSON &#x60;(content-type&#x3D;application/json)&#x60;  | 
**Method** | [**EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod**](EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod.md) |  | 
**PhoneNumberFormat** | [**EnumNotificationsSettingsPhoneDeliverySettingsCustomNumberFormat**](EnumNotificationsSettingsPhoneDeliverySettingsCustomNumberFormat.md) |  | [default to ENUMNOTIFICATIONSSETTINGSPHONEDELIVERYSETTINGSCUSTOMNUMBERFORMAT_FULL]
**BeforeTag** | Pointer to **string** | For voice OTP notifications only. An opening tag which is commonly used by custom providers for defining a pause between each number in the OTP number string. Possible value: &#x60;&lt;Say&gt;&#x60;  | [optional] 
**AfterTag** | Pointer to **string** | For voice OTP notifications only. A closing tag which is commonly used by custom providers for defining a pause between each number in the OTP number string. Possible value: &#x60;&lt;/Say&gt; &lt;Pause length&#x3D;\&quot;1\&quot;/&gt;&#x60;  | [optional] 

## Methods

### NewNotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests

`func NewNotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests(deliveryMethod EnumNotificationsSettingsPhoneDeliverySettingsCustomDeliveryMethod, url string, body string, headers string, method EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod, phoneNumberFormat EnumNotificationsSettingsPhoneDeliverySettingsCustomNumberFormat, ) *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests`

NewNotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests instantiates a new NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsSettingsPhoneDeliverySettingsCustomAllOfRequestsWithDefaults

`func NewNotificationsSettingsPhoneDeliverySettingsCustomAllOfRequestsWithDefaults() *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests`

NewNotificationsSettingsPhoneDeliverySettingsCustomAllOfRequestsWithDefaults instantiates a new NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryMethod

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) GetDeliveryMethod() EnumNotificationsSettingsPhoneDeliverySettingsCustomDeliveryMethod`

GetDeliveryMethod returns the DeliveryMethod field if non-nil, zero value otherwise.

### GetDeliveryMethodOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) GetDeliveryMethodOk() (*EnumNotificationsSettingsPhoneDeliverySettingsCustomDeliveryMethod, bool)`

GetDeliveryMethodOk returns a tuple with the DeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethod

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) SetDeliveryMethod(v EnumNotificationsSettingsPhoneDeliverySettingsCustomDeliveryMethod)`

SetDeliveryMethod sets DeliveryMethod field to given value.


### GetUrl

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetBody

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) SetBody(v string)`

SetBody sets Body field to given value.


### GetHeaders

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) GetHeaders() string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) GetHeadersOk() (*string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) SetHeaders(v string)`

SetHeaders sets Headers field to given value.


### GetMethod

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) GetMethod() EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) GetMethodOk() (*EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) SetMethod(v EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod)`

SetMethod sets Method field to given value.


### GetPhoneNumberFormat

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) GetPhoneNumberFormat() EnumNotificationsSettingsPhoneDeliverySettingsCustomNumberFormat`

GetPhoneNumberFormat returns the PhoneNumberFormat field if non-nil, zero value otherwise.

### GetPhoneNumberFormatOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) GetPhoneNumberFormatOk() (*EnumNotificationsSettingsPhoneDeliverySettingsCustomNumberFormat, bool)`

GetPhoneNumberFormatOk returns a tuple with the PhoneNumberFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberFormat

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) SetPhoneNumberFormat(v EnumNotificationsSettingsPhoneDeliverySettingsCustomNumberFormat)`

SetPhoneNumberFormat sets PhoneNumberFormat field to given value.


### GetBeforeTag

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) GetBeforeTag() string`

GetBeforeTag returns the BeforeTag field if non-nil, zero value otherwise.

### GetBeforeTagOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) GetBeforeTagOk() (*string, bool)`

GetBeforeTagOk returns a tuple with the BeforeTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeTag

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) SetBeforeTag(v string)`

SetBeforeTag sets BeforeTag field to given value.

### HasBeforeTag

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) HasBeforeTag() bool`

HasBeforeTag returns a boolean if a field has been set.

### GetAfterTag

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) GetAfterTag() string`

GetAfterTag returns the AfterTag field if non-nil, zero value otherwise.

### GetAfterTagOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) GetAfterTagOk() (*string, bool)`

GetAfterTagOk returns a tuple with the AfterTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterTag

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) SetAfterTag(v string)`

SetAfterTag sets AfterTag field to given value.

### HasAfterTag

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) HasAfterTag() bool`

HasAfterTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


