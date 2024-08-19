# AuthorizeEditorDataRulesReferenceableRuleDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**EffectSettings** | [**AuthorizeEditorDataRulesEffectSettingsDTO**](AuthorizeEditorDataRulesEffectSettingsDTO.md) |  | 
**Version** | **string** |  | 

## Methods

### NewAuthorizeEditorDataRulesReferenceableRuleDTO

`func NewAuthorizeEditorDataRulesReferenceableRuleDTO(effectSettings AuthorizeEditorDataRulesEffectSettingsDTO, version string, ) *AuthorizeEditorDataRulesReferenceableRuleDTO`

NewAuthorizeEditorDataRulesReferenceableRuleDTO instantiates a new AuthorizeEditorDataRulesReferenceableRuleDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataRulesReferenceableRuleDTOWithDefaults

`func NewAuthorizeEditorDataRulesReferenceableRuleDTOWithDefaults() *AuthorizeEditorDataRulesReferenceableRuleDTO`

NewAuthorizeEditorDataRulesReferenceableRuleDTOWithDefaults instantiates a new AuthorizeEditorDataRulesReferenceableRuleDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEffectSettings

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetEffectSettings() AuthorizeEditorDataRulesEffectSettingsDTO`

GetEffectSettings returns the EffectSettings field if non-nil, zero value otherwise.

### GetEffectSettingsOk

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetEffectSettingsOk() (*AuthorizeEditorDataRulesEffectSettingsDTO, bool)`

GetEffectSettingsOk returns a tuple with the EffectSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectSettings

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) SetEffectSettings(v AuthorizeEditorDataRulesEffectSettingsDTO)`

SetEffectSettings sets EffectSettings field to given value.


### GetVersion

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


