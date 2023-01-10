# TemplateContentVoiceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | The voice text.  Limited to 1Kb characters.   The following substitution place holders can be embedded in the message:   * &#x60;&lt;pause1sec&gt;&#x60;: Pauses the message narration for 1 second.     The pause interval &lt;pause1sec&gt; cannot be modified. To pause the message narration for more than one second, repeat multiple &#x60;&lt;pause1sec&gt;&#x60; occurrences in succession, according to the desired pause interval duration. For example, &#x60;&lt;pause1sec&gt;&lt;pause1sec&gt;&lt;pause1sec&gt;&#x60; pauses the message narration for three seconds.   * &#x60;&lt;sayCharValue&gt; .. &lt;/sayCharValue&gt;&#x60;: Reads the character name of each character of the enclosed string separately.   * &#x60;&lt;repeatMessage val&#x3D;x&gt; .. &lt;/repeatMessage&gt;&#x60;: Narrates the enclosed text &#x60;&lt;val&gt;&#x60; number of times.      In the following message example, &#x60;${otp}&#x60; is assigned the value &#x60;\&quot;123456\&quot;&#x60;, and &#x60;${email}&#x60; is assigned the value &#x60;\&quot;joe@bxz.com\&quot;&#x60;:   &#x60;&#x60;&#x60;   Hello &lt;pause1sec&gt; your authentication code is    &lt;sayCharValue&gt;${otp}&lt;/​sayCharValue&gt;    &lt;repeatMessage val&#x3D;2&gt; I repeat your code is    &lt;sayCharValue&gt;${otp}&lt;/​sayCharValue&gt;   &lt;/​repeatMessage&gt; &lt;pause1sec&gt;    Mail &lt;sayCharValue&gt;${email}&lt;/​sayCharValue&gt; for help   &#x60;&#x60;&#x60;   The narrated message on the voice call sounds like:   &#x60;&#x60;&#x60;   HELLO &lt;1 second silence&gt; YOUR AUTHENTICATION CODE IS    ONE TWO THREE FOUR FIVE SIX    I REPEAT YOUR CODE IS ONE TWO THREE FOUR FIVE SIX    I REPEAT YOUR CODE IS ONE TWO THREE FOUR FIVE SIX &lt;1 second silence&gt;    MAIL JAY OH EE AT BEE EX ZEE DOT SEE OH EM FOR HELP   &#x60;&#x60;&#x60;  | 
**Voice** | Pointer to **string** | Voice OTP supports vendor-specific voices.   * Supported Twilio voices:     * Man, Woman       Supported locales (default: en):       en, en_GB, es, fr, de     * Alice (Twilio only)       Supported locales (default: en US):       da_DK, de_DE, en_AU, en_CA, en_GB, en_US, ca_ES, es_ES, es_MX, fi_FI, fr_CA, fr_FR, it_IT, ja_JP, ko_KR, nb_NO, nl_NL, pl_PL, pt_BR, pt_PT, ru_RU, sv_SE, zh_CN, zh_HK, zh_TW     * Amazon Polly       cy_GB, ro_RO, is_IS, hi_IN tr_TR   * Supported Syniverse voices:     * Man, Woman       Supported locales:       en_US, en_GB, es_ES, es_US, fr_FR, de_DE, it_IT, en_AU, da_DK, is_IS, nl_NL, pl_PL, pt_BR, pt_PT, ru_RU, ja_JP     * Woman only       Supported locales:       cmn_CN, cy_GB, en_IN, fr_CA, hi_IN, nb_NO, ro_RO, sv_SE, tr_TR, ko_KR, ar  | [optional] 

## Methods

### NewTemplateContentVoiceAllOf

`func NewTemplateContentVoiceAllOf(content string, ) *TemplateContentVoiceAllOf`

NewTemplateContentVoiceAllOf instantiates a new TemplateContentVoiceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateContentVoiceAllOfWithDefaults

`func NewTemplateContentVoiceAllOfWithDefaults() *TemplateContentVoiceAllOf`

NewTemplateContentVoiceAllOfWithDefaults instantiates a new TemplateContentVoiceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *TemplateContentVoiceAllOf) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *TemplateContentVoiceAllOf) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *TemplateContentVoiceAllOf) SetContent(v string)`

SetContent sets Content field to given value.


### GetVoice

`func (o *TemplateContentVoiceAllOf) GetVoice() string`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *TemplateContentVoiceAllOf) GetVoiceOk() (*string, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *TemplateContentVoiceAllOf) SetVoice(v string)`

SetVoice sets Voice field to given value.

### HasVoice

`func (o *TemplateContentVoiceAllOf) HasVoice() bool`

HasVoice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


