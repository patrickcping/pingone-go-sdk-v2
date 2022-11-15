# TemplateContentSMS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The template id. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | The time the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | The time the resource was last updated. | [optional] [readonly] 
**Default** | Pointer to **bool** | Specifies whether the template is a predefined default template. | [optional] 
**Locale** | **string** | A valid case-insensitive locale, complying with the ISO-639 language code and ISO-3166 country code standards: Two-character language code, for example, \&quot;en\&quot;. Two-character language code followed by a two-character country code, separated by an underscore or dash, for example: \&quot;en_GB\&quot;, \&quot;en-GB\&quot;. Cannot be changed after it is initially set in &#x60;POST /environments/{{envID}}/templates/{{templateName}}/contents&#x60;.  | 
**DeliveryMethod** | **string** | The content&#39;s delivery method. Possible values are &#x60;Email&#x60;, &#x60;SMS&#x60;, &#x60;Voice&#x60; or &#x60;Push&#x60;. Cannot be changed after it is initially set in &#x60;POST /environments/{{envID}}/templates/{{templateName}}/contents&#x60;. | 
**Variant** | Pointer to **string** | Holds the unique user-defined name for each content variant that uses the same template + &#x60;deliveryMethod&#x60; + &#x60;locale&#x60; combination. This property is case insensitive and has a limit of 100 characters. For more information, see [Creating custom contents](https://apidocs.pingidentity.com/pingone/platform/v1/api/#notifications-templates-creating-custom-contents). | [optional] 
**Content** | **string** | The SMS or voice text. * **SMS**: UC-2 encoding is used for text that contains non GSM-7 characters. UC-2 encoded text cannot exceed 67 characters. GSM-7 encoded text cannot exceed 153 characters. If supported, it can include variables. * **Voice**: Limited to 1Kb characters.   The following substitution place holders can be embedded in the message:   * &#x60;&lt;pause1sec&gt;&#x60;: Pauses the message narration for 1 second.     The pause interval &lt;pause1sec&gt; cannot be modified. To pause the message narration for more than one second, repeat multiple &#x60;&lt;pause1sec&gt;&#x60; occurrences in succession, according to the desired pause interval duration. For example, &#x60;&lt;pause1sec&gt;&lt;pause1sec&gt;&lt;pause1sec&gt;&#x60; pauses the message narration for three seconds.   * &#x60;&lt;sayCharValue&gt; .. &lt;/sayCharValue&gt;&#x60;: Reads the character name of each character of the enclosed string separately.   * &#x60;&lt;repeatMessage val&#x3D;x&gt; .. &lt;/repeatMessage&gt;&#x60;: Narrates the enclosed text &#x60;&lt;val&gt;&#x60; number of times.      In the following message example, &#x60;${otp}&#x60; is assigned the value &#x60;\&quot;123456\&quot;&#x60;, and &#x60;${email}&#x60; is assigned the value &#x60;\&quot;joe@bxz.com\&quot;&#x60;:   &#x60;&#x60;&#x60;   Hello &lt;pause1sec&gt; your authentication code is    &lt;sayCharValue&gt;${otp}&lt;/​sayCharValue&gt;    &lt;repeatMessage val&#x3D;2&gt; I repeat your code is    &lt;sayCharValue&gt;${otp}&lt;/​sayCharValue&gt;   &lt;/​repeatMessage&gt; &lt;pause1sec&gt;    Mail &lt;sayCharValue&gt;${email}&lt;/​sayCharValue&gt; for help   &#x60;&#x60;&#x60;   The narrated message on the voice call sounds like:   &#x60;&#x60;&#x60;   HELLO &lt;1 second silence&gt; YOUR AUTHENTICATION CODE IS    ONE TWO THREE FOUR FIVE SIX    I REPEAT YOUR CODE IS ONE TWO THREE FOUR FIVE SIX    I REPEAT YOUR CODE IS ONE TWO THREE FOUR FIVE SIX &lt;1 second silence&gt;    MAIL JAY OH EE AT BEE EX ZEE DOT SEE OH EM FOR HELP   &#x60;&#x60;&#x60;    Voice OTP supports vendor-specific voices.   * Supported Twilio voices:     * Man, Woman       Supported locales (default: en):       en, en_GB, es, fr, de     * Alice (Twilio only)       Supported locales (default: en US):       da_DK, de_DE, en_AU, en_CA, en_GB, en_US, ca_ES, es_ES, es_MX, fi_FI, fr_CA, fr_FR, it_IT, ja_JP, ko_KR, nb_NO, nl_NL, pl_PL, pt_BR, pt_PT, ru_RU, sv_SE, zh_CN, zh_HK, zh_TW     * Amazon Polly       cy_GB, ro_RO, is_IS, hi_IN tr_TR   * Supported Syniverse voices:     * Man, Woman       Supported locales:       en_US, en_GB, es_ES, es_US, fr_FR, de_DE, it_IT, en_AU, da_DK, is_IS, nl_NL, pl_PL, pt_BR, pt_PT, ru_RU, ja_JP     * Woman only       Supported locales:       cmn_CN, cy_GB, en_IN, fr_CA, hi_IN, nb_NO, ro_RO, sv_SE, tr_TR, ko_KR, ar  | 
**Sender** | Pointer to **string** | The SMS sender ID. This property can contain only alphanumeric characters and spaces, and its length cannot exceed 11 characters. In some countries, it is impossible to send an SMS with an alphanumeric sender ID. For those countries, the sender ID must be empty. For SMS recipients in specific countries, refer to Twilio&#39;s documentation on [International support for Alphanumeric Sender ID](https://support.twilio.com/hc/en-us/articles/223133767-International-support-for-Alphanumeric-Sender-ID). | [optional] 

## Methods

### NewTemplateContentSMS

`func NewTemplateContentSMS(locale string, deliveryMethod string, content string, ) *TemplateContentSMS`

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

`func (o *TemplateContentSMS) GetDeliveryMethod() string`

GetDeliveryMethod returns the DeliveryMethod field if non-nil, zero value otherwise.

### GetDeliveryMethodOk

`func (o *TemplateContentSMS) GetDeliveryMethodOk() (*string, bool)`

GetDeliveryMethodOk returns a tuple with the DeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethod

`func (o *TemplateContentSMS) SetDeliveryMethod(v string)`

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


