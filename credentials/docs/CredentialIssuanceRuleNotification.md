# CredentialIssuanceRuleNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Methods** | Pointer to [**[]EnumCredentialIssuanceRuleNotificationMethod**](EnumCredentialIssuanceRuleNotificationMethod.md) |  | [optional] 
**Locale** | Pointer to **string** | A string that specifies the ISO 2-character language code used for the notification; for example, en. | [optional] 
**Variables** | Pointer to **map[string]interface{}** | An object of key:value pairs that defines the dynamic variables used by the content variant. | [optional] 
**Variant** | Pointer to **string** | A string that specifies the unique user-defined name for the content variant that contains the message text used for the notification | [optional] 

## Methods

### NewCredentialIssuanceRuleNotification

`func NewCredentialIssuanceRuleNotification() *CredentialIssuanceRuleNotification`

NewCredentialIssuanceRuleNotification instantiates a new CredentialIssuanceRuleNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialIssuanceRuleNotificationWithDefaults

`func NewCredentialIssuanceRuleNotificationWithDefaults() *CredentialIssuanceRuleNotification`

NewCredentialIssuanceRuleNotificationWithDefaults instantiates a new CredentialIssuanceRuleNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethods

`func (o *CredentialIssuanceRuleNotification) GetMethods() []EnumCredentialIssuanceRuleNotificationMethod`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *CredentialIssuanceRuleNotification) GetMethodsOk() (*[]EnumCredentialIssuanceRuleNotificationMethod, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *CredentialIssuanceRuleNotification) SetMethods(v []EnumCredentialIssuanceRuleNotificationMethod)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *CredentialIssuanceRuleNotification) HasMethods() bool`

HasMethods returns a boolean if a field has been set.

### GetLocale

`func (o *CredentialIssuanceRuleNotification) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *CredentialIssuanceRuleNotification) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *CredentialIssuanceRuleNotification) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *CredentialIssuanceRuleNotification) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetVariables

`func (o *CredentialIssuanceRuleNotification) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *CredentialIssuanceRuleNotification) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *CredentialIssuanceRuleNotification) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *CredentialIssuanceRuleNotification) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetVariant

`func (o *CredentialIssuanceRuleNotification) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *CredentialIssuanceRuleNotification) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *CredentialIssuanceRuleNotification) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *CredentialIssuanceRuleNotification) HasVariant() bool`

HasVariant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


