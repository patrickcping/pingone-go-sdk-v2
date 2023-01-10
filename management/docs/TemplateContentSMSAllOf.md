# TemplateContentSMSAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | The SMS text. UC-2 encoding is used for text that contains non GSM-7 characters. UC-2 encoded text cannot exceed 67 characters. GSM-7 encoded text cannot exceed 153 characters. If supported, it can include variables.  | 
**Sender** | Pointer to **string** | The SMS sender ID. This property can contain only alphanumeric characters and spaces, and its length cannot exceed 11 characters. In some countries, it is impossible to send an SMS with an alphanumeric sender ID. For those countries, the sender ID must be empty. For SMS recipients in specific countries, refer to Twilio&#39;s documentation on [International support for Alphanumeric Sender ID](https://support.twilio.com/hc/en-us/articles/223133767-International-support-for-Alphanumeric-Sender-ID). | [optional] 

## Methods

### NewTemplateContentSMSAllOf

`func NewTemplateContentSMSAllOf(content string, ) *TemplateContentSMSAllOf`

NewTemplateContentSMSAllOf instantiates a new TemplateContentSMSAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateContentSMSAllOfWithDefaults

`func NewTemplateContentSMSAllOfWithDefaults() *TemplateContentSMSAllOf`

NewTemplateContentSMSAllOfWithDefaults instantiates a new TemplateContentSMSAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *TemplateContentSMSAllOf) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *TemplateContentSMSAllOf) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *TemplateContentSMSAllOf) SetContent(v string)`

SetContent sets Content field to given value.


### GetSender

`func (o *TemplateContentSMSAllOf) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *TemplateContentSMSAllOf) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *TemplateContentSMSAllOf) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *TemplateContentSMSAllOf) HasSender() bool`

HasSender returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


