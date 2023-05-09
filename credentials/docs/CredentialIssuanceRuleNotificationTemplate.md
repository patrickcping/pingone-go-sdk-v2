# CredentialIssuanceRuleNotificationTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locale** | Pointer to **string** | A string that specifies the ISO 2-character language code used for the notification; for example, en. | [optional] 
**Variables** | Pointer to **map[string]interface{}** | An object of key:value pairs that defines the dynamic variables used by the content variant. | [optional] 
**Variant** | Pointer to **string** | A string that specifies the unique user-defined name for the content variant that contains the message text used for the notification | [optional] 

## Methods

### NewCredentialIssuanceRuleNotificationTemplate

`func NewCredentialIssuanceRuleNotificationTemplate() *CredentialIssuanceRuleNotificationTemplate`

NewCredentialIssuanceRuleNotificationTemplate instantiates a new CredentialIssuanceRuleNotificationTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialIssuanceRuleNotificationTemplateWithDefaults

`func NewCredentialIssuanceRuleNotificationTemplateWithDefaults() *CredentialIssuanceRuleNotificationTemplate`

NewCredentialIssuanceRuleNotificationTemplateWithDefaults instantiates a new CredentialIssuanceRuleNotificationTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocale

`func (o *CredentialIssuanceRuleNotificationTemplate) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *CredentialIssuanceRuleNotificationTemplate) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *CredentialIssuanceRuleNotificationTemplate) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *CredentialIssuanceRuleNotificationTemplate) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetVariables

`func (o *CredentialIssuanceRuleNotificationTemplate) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *CredentialIssuanceRuleNotificationTemplate) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *CredentialIssuanceRuleNotificationTemplate) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *CredentialIssuanceRuleNotificationTemplate) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetVariant

`func (o *CredentialIssuanceRuleNotificationTemplate) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *CredentialIssuanceRuleNotificationTemplate) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *CredentialIssuanceRuleNotificationTemplate) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *CredentialIssuanceRuleNotificationTemplate) HasVariant() bool`

HasVariant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


