# TemplateContentSMS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The template id. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | The time the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | The time the resource was last updated. | [optional] [readonly] 
**Default** | Pointer to **bool** | Specifies whether the template is a predefined default template. | [optional] 
**Locale** | **string** | A valid case-insensitive locale, complying with the ISO-639 language code and ISO-3166 country code standards: Two-character language code, for example, \&quot;en\&quot;. Two-character language code followed by a two-character country code, separated by an underscore or dash, for example: \&quot;en_GB\&quot;, \&quot;en-GB\&quot;. Cannot be changed after it is initially set in &#x60;POST /environments/{{envID}}/templates/{{templateName}}/contents&#x60;.  | 
**DeliveryMethod** | [**EnumTemplateContentDeliveryMethod**](EnumTemplateContentDeliveryMethod.md) |  | 
**Variant** | Pointer to **string** | Holds the unique user-defined name for each content variant that uses the same template + &#x60;deliveryMethod&#x60; + &#x60;locale&#x60; combination. This property is case insensitive and has a limit of 100 characters. For more information, see [Creating custom contents](https://apidocs.pingidentity.com/pingone/platform/v1/api/#notifications-templates-creating-custom-contents). | [optional] 
**Content** | **string** | The SMS text. UC-2 encoding is used for text that contains non GSM-7 characters. UC-2 encoded text cannot exceed 67 characters. GSM-7 encoded text cannot exceed 153 characters. If supported, it can include variables.  | 
**Sender** | Pointer to **string** | The SMS sender ID. This property can contain only alphanumeric characters and spaces, and its length cannot exceed 11 characters. In some countries, it is impossible to send an SMS with an alphanumeric sender ID. For those countries, the sender ID must be empty. For SMS recipients in specific countries, refer to Twilio&#39;s documentation on [International support for Alphanumeric Sender ID](https://support.twilio.com/hc/en-us/articles/223133767-International-support-for-Alphanumeric-Sender-ID). | [optional] 

## Methods

### NewTemplateContentSMS

`func NewTemplateContentSMS(locale string, deliveryMethod EnumTemplateContentDeliveryMethod, content string, ) *TemplateContentSMS`

NewTemplateContentSMS instantiates a new TemplateContentSMS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateContentSMSWithDefaults

`func NewTemplateContentSMSWithDefaults() *TemplateContentSMS`

NewTemplateContentSMSWithDefaults instantiates a new TemplateContentSMS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateContentSMS) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateContentSMS) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateContentSMS) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateContentSMS) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TemplateContentSMS) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TemplateContentSMS) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TemplateContentSMS) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TemplateContentSMS) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TemplateContentSMS) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TemplateContentSMS) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TemplateContentSMS) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TemplateContentSMS) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDefault

`func (o *TemplateContentSMS) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *TemplateContentSMS) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *TemplateContentSMS) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *TemplateContentSMS) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetLocale

`func (o *TemplateContentSMS) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *TemplateContentSMS) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *TemplateContentSMS) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetDeliveryMethod

`func (o *TemplateContentSMS) GetDeliveryMethod() EnumTemplateContentDeliveryMethod`

GetDeliveryMethod returns the DeliveryMethod field if non-nil, zero value otherwise.

### GetDeliveryMethodOk

`func (o *TemplateContentSMS) GetDeliveryMethodOk() (*EnumTemplateContentDeliveryMethod, bool)`

GetDeliveryMethodOk returns a tuple with the DeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethod

`func (o *TemplateContentSMS) SetDeliveryMethod(v EnumTemplateContentDeliveryMethod)`

SetDeliveryMethod sets DeliveryMethod field to given value.


### GetVariant

`func (o *TemplateContentSMS) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *TemplateContentSMS) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *TemplateContentSMS) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *TemplateContentSMS) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetContent

`func (o *TemplateContentSMS) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *TemplateContentSMS) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *TemplateContentSMS) SetContent(v string)`

SetContent sets Content field to given value.


### GetSender

`func (o *TemplateContentSMS) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *TemplateContentSMS) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *TemplateContentSMS) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *TemplateContentSMS) HasSender() bool`

HasSender returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


