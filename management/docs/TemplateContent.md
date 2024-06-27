# TemplateContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Id** | Pointer to **string** | The template id. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | The time the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | The time the resource was last updated. | [optional] [readonly] 
**Default** | Pointer to **bool** | Specifies whether the template is a predefined default template. | [optional] 
**Locale** | **string** | A valid case-insensitive locale, complying with the ISO-639 language code and ISO-3166 country code standards: Two-character language code, for example, \&quot;en\&quot;. Two-character language code followed by a two-character country code, separated by an underscore or dash, for example: \&quot;en_GB\&quot;, \&quot;en-GB\&quot;. Cannot be changed after it is initially set in &#x60;POST /environments/{{envID}}/templates/{{templateName}}/contents&#x60;.  | 
**DeliveryMethod** | [**EnumTemplateContentDeliveryMethod**](EnumTemplateContentDeliveryMethod.md) |  | 
**Variant** | Pointer to **string** | Holds the unique user-defined name for each content variant that uses the same template + &#x60;deliveryMethod&#x60; + &#x60;locale&#x60; combination. This property is case insensitive and has a limit of 100 characters. For more information, see [Creating custom contents](https://apidocs.pingidentity.com/pingone/platform/v1/api/#notifications-templates-creating-custom-contents). | [optional] 
**Body** | **string** | The push text (maximum 400 characters for push text). If supported, this can include variables. | 
**From** | Pointer to [**TemplateContentEmailAllOfFrom**](TemplateContentEmailAllOfFrom.md) |  | [optional] 
**Subject** | Pointer to **string** | The email&#39;s subject line. Cannot exceed 256 characters. If supported, can include variables. | [optional] 
**ReplyTo** | Pointer to [**TemplateContentEmailAllOfReplyTo**](TemplateContentEmailAllOfReplyTo.md) |  | [optional] 
**Charset** | Pointer to **string** | If not specified, &#x60;UTF-8&#x60; is the default value. | [optional] [default to "UTF-8"]
**EmailContentType** | Pointer to **string** | If not specified, &#x60;text/html&#x60; is the default value. | [optional] [default to "text/html"]
**Content** | **string** | The voice text.  Limited to 1Kb characters.   The following substitution place holders can be embedded in the message:   * &#x60;&lt;pause1sec&gt;&#x60;: Pauses the message narration for 1 second.     The pause interval &lt;pause1sec&gt; cannot be modified. To pause the message narration for more than one second, repeat multiple &#x60;&lt;pause1sec&gt;&#x60; occurrences in succession, according to the desired pause interval duration. For example, &#x60;&lt;pause1sec&gt;&lt;pause1sec&gt;&lt;pause1sec&gt;&#x60; pauses the message narration for three seconds.   * &#x60;&lt;sayCharValue&gt; .. &lt;/sayCharValue&gt;&#x60;: Reads the character name of each character of the enclosed string separately.   * &#x60;&lt;repeatMessage val&#x3D;x&gt; .. &lt;/repeatMessage&gt;&#x60;: Narrates the enclosed text &#x60;&lt;val&gt;&#x60; number of times.      In the following message example, &#x60;${otp}&#x60; is assigned the value &#x60;\&quot;123456\&quot;&#x60;, and &#x60;${email}&#x60; is assigned the value &#x60;\&quot;joe@bxz.com\&quot;&#x60;:   &#x60;&#x60;&#x60;   Hello &lt;pause1sec&gt; your authentication code is    &lt;sayCharValue&gt;${otp}&lt;/​sayCharValue&gt;    &lt;repeatMessage val&#x3D;2&gt; I repeat your code is    &lt;sayCharValue&gt;${otp}&lt;/​sayCharValue&gt;   &lt;/​repeatMessage&gt; &lt;pause1sec&gt;    Mail &lt;sayCharValue&gt;${email}&lt;/​sayCharValue&gt; for help   &#x60;&#x60;&#x60;   The narrated message on the voice call sounds like:   &#x60;&#x60;&#x60;   HELLO &lt;1 second silence&gt; YOUR AUTHENTICATION CODE IS    ONE TWO THREE FOUR FIVE SIX    I REPEAT YOUR CODE IS ONE TWO THREE FOUR FIVE SIX    I REPEAT YOUR CODE IS ONE TWO THREE FOUR FIVE SIX &lt;1 second silence&gt;    MAIL JAY OH EE AT BEE EX ZEE DOT SEE OH EM FOR HELP   &#x60;&#x60;&#x60;  | 
**Sender** | Pointer to **string** | The SMS sender ID. This property can contain only alphanumeric characters and spaces, and its length cannot exceed 11 characters. In some countries, it is impossible to send an SMS with an alphanumeric sender ID. For those countries, the sender ID must be empty. For SMS recipients in specific countries, refer to Twilio&#39;s documentation on [International support for Alphanumeric Sender ID](https://support.twilio.com/hc/en-us/articles/223133767-International-support-for-Alphanumeric-Sender-ID). | [optional] 
**Title** | **string** | The push title (maximum 200 characters). If supported, this can include variables. | 
**PushCategory** | Pointer to [**EnumTemplateContentPushCategory**](EnumTemplateContentPushCategory.md) |  | [optional] [default to ENUMTEMPLATECONTENTPUSHCATEGORY_BANNER_BUTTONS]
**Voice** | Pointer to **string** | Voice OTP supports vendor-specific voices.   * Supported Twilio voices:     * Man, Woman       Supported locales (default: en):       en, en_GB, es, fr, de     * Alice (Twilio only)       Supported locales (default: en US):       da_DK, de_DE, en_AU, en_CA, en_GB, en_US, ca_ES, es_ES, es_MX, fi_FI, fr_CA, fr_FR, it_IT, ja_JP, ko_KR, nb_NO, nl_NL, pl_PL, pt_BR, pt_PT, ru_RU, sv_SE, zh_CN, zh_HK, zh_TW     * Amazon Polly       cy_GB, ro_RO, is_IS, hi_IN tr_TR   * Supported Syniverse voices:     * Man, Woman       Supported locales:       en_US, en_GB, es_ES, es_US, fr_FR, de_DE, it_IT, en_AU, da_DK, is_IS, nl_NL, pl_PL, pt_BR, pt_PT, ru_RU, ja_JP     * Woman only       Supported locales:       cmn_CN, cy_GB, en_IN, fr_CA, hi_IN, nb_NO, ro_RO, sv_SE, tr_TR, ko_KR, ar  | [optional] 

## Methods

### NewTemplateContent

`func NewTemplateContent(locale string, deliveryMethod EnumTemplateContentDeliveryMethod, body string, content string, title string, ) *TemplateContent`

NewTemplateContent instantiates a new TemplateContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateContentWithDefaults

`func NewTemplateContentWithDefaults() *TemplateContent`

NewTemplateContentWithDefaults instantiates a new TemplateContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *TemplateContent) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TemplateContent) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TemplateContent) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TemplateContent) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *TemplateContent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateContent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateContent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateContent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TemplateContent) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TemplateContent) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TemplateContent) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TemplateContent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TemplateContent) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TemplateContent) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TemplateContent) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TemplateContent) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDefault

`func (o *TemplateContent) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *TemplateContent) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *TemplateContent) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *TemplateContent) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetLocale

`func (o *TemplateContent) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *TemplateContent) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *TemplateContent) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetDeliveryMethod

`func (o *TemplateContent) GetDeliveryMethod() EnumTemplateContentDeliveryMethod`

GetDeliveryMethod returns the DeliveryMethod field if non-nil, zero value otherwise.

### GetDeliveryMethodOk

`func (o *TemplateContent) GetDeliveryMethodOk() (*EnumTemplateContentDeliveryMethod, bool)`

GetDeliveryMethodOk returns a tuple with the DeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethod

`func (o *TemplateContent) SetDeliveryMethod(v EnumTemplateContentDeliveryMethod)`

SetDeliveryMethod sets DeliveryMethod field to given value.


### GetVariant

`func (o *TemplateContent) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *TemplateContent) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *TemplateContent) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *TemplateContent) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetBody

`func (o *TemplateContent) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TemplateContent) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TemplateContent) SetBody(v string)`

SetBody sets Body field to given value.


### GetFrom

`func (o *TemplateContent) GetFrom() TemplateContentEmailAllOfFrom`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TemplateContent) GetFromOk() (*TemplateContentEmailAllOfFrom, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TemplateContent) SetFrom(v TemplateContentEmailAllOfFrom)`

SetFrom sets From field to given value.

### HasFrom

`func (o *TemplateContent) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetSubject

`func (o *TemplateContent) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TemplateContent) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TemplateContent) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *TemplateContent) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetReplyTo

`func (o *TemplateContent) GetReplyTo() TemplateContentEmailAllOfReplyTo`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *TemplateContent) GetReplyToOk() (*TemplateContentEmailAllOfReplyTo, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *TemplateContent) SetReplyTo(v TemplateContentEmailAllOfReplyTo)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *TemplateContent) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetCharset

`func (o *TemplateContent) GetCharset() string`

GetCharset returns the Charset field if non-nil, zero value otherwise.

### GetCharsetOk

`func (o *TemplateContent) GetCharsetOk() (*string, bool)`

GetCharsetOk returns a tuple with the Charset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharset

`func (o *TemplateContent) SetCharset(v string)`

SetCharset sets Charset field to given value.

### HasCharset

`func (o *TemplateContent) HasCharset() bool`

HasCharset returns a boolean if a field has been set.

### GetEmailContentType

`func (o *TemplateContent) GetEmailContentType() string`

GetEmailContentType returns the EmailContentType field if non-nil, zero value otherwise.

### GetEmailContentTypeOk

`func (o *TemplateContent) GetEmailContentTypeOk() (*string, bool)`

GetEmailContentTypeOk returns a tuple with the EmailContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailContentType

`func (o *TemplateContent) SetEmailContentType(v string)`

SetEmailContentType sets EmailContentType field to given value.

### HasEmailContentType

`func (o *TemplateContent) HasEmailContentType() bool`

HasEmailContentType returns a boolean if a field has been set.

### GetContent

`func (o *TemplateContent) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *TemplateContent) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *TemplateContent) SetContent(v string)`

SetContent sets Content field to given value.


### GetSender

`func (o *TemplateContent) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *TemplateContent) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *TemplateContent) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *TemplateContent) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetTitle

`func (o *TemplateContent) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TemplateContent) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TemplateContent) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPushCategory

`func (o *TemplateContent) GetPushCategory() EnumTemplateContentPushCategory`

GetPushCategory returns the PushCategory field if non-nil, zero value otherwise.

### GetPushCategoryOk

`func (o *TemplateContent) GetPushCategoryOk() (*EnumTemplateContentPushCategory, bool)`

GetPushCategoryOk returns a tuple with the PushCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushCategory

`func (o *TemplateContent) SetPushCategory(v EnumTemplateContentPushCategory)`

SetPushCategory sets PushCategory field to given value.

### HasPushCategory

`func (o *TemplateContent) HasPushCategory() bool`

HasPushCategory returns a boolean if a field has been set.

### GetVoice

`func (o *TemplateContent) GetVoice() string`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *TemplateContent) GetVoiceOk() (*string, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *TemplateContent) SetVoice(v string)`

SetVoice sets Voice field to given value.

### HasVoice

`func (o *TemplateContent) HasVoice() bool`

HasVoice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


