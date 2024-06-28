# AgreementLanguage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Agreement** | Pointer to [**AgreementLanguageAgreement**](AgreementLanguageAgreement.md) |  | [optional] 
**CurrentRevision** | Pointer to [**AgreementLanguageCurrentRevision**](AgreementLanguageCurrentRevision.md) |  | [optional] 
**DisplayName** | **string** | A string that is used as the title of the agreement for the language presented to the user. This is a required property. | 
**Enabled** | **bool** | A boolean that maps directly with a language being enabled or displayed for the environment within the platform. This is a required property. | 
**Id** | Pointer to **string** | A string that specifies the language ID. | [optional] [readonly] 
**Locale** | **string** | A string that specifies the tag for identifying the language resource associated with this agreement consent (for example, en-US). This is a required property. For more information about language tags, see Tags for Identifying Languages. | 
**UserExperience** | Pointer to [**AgreementLanguageUserExperience**](AgreementLanguageUserExperience.md) |  | [optional] 

## Methods

### NewAgreementLanguage

`func NewAgreementLanguage(displayName string, enabled bool, locale string, ) *AgreementLanguage`

NewAgreementLanguage instantiates a new AgreementLanguage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementLanguageWithDefaults

`func NewAgreementLanguageWithDefaults() *AgreementLanguage`

NewAgreementLanguageWithDefaults instantiates a new AgreementLanguage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AgreementLanguage) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AgreementLanguage) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AgreementLanguage) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AgreementLanguage) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAgreement

`func (o *AgreementLanguage) GetAgreement() AgreementLanguageAgreement`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *AgreementLanguage) GetAgreementOk() (*AgreementLanguageAgreement, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *AgreementLanguage) SetAgreement(v AgreementLanguageAgreement)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *AgreementLanguage) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetCurrentRevision

`func (o *AgreementLanguage) GetCurrentRevision() AgreementLanguageCurrentRevision`

GetCurrentRevision returns the CurrentRevision field if non-nil, zero value otherwise.

### GetCurrentRevisionOk

`func (o *AgreementLanguage) GetCurrentRevisionOk() (*AgreementLanguageCurrentRevision, bool)`

GetCurrentRevisionOk returns a tuple with the CurrentRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRevision

`func (o *AgreementLanguage) SetCurrentRevision(v AgreementLanguageCurrentRevision)`

SetCurrentRevision sets CurrentRevision field to given value.

### HasCurrentRevision

`func (o *AgreementLanguage) HasCurrentRevision() bool`

HasCurrentRevision returns a boolean if a field has been set.

### GetDisplayName

`func (o *AgreementLanguage) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AgreementLanguage) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AgreementLanguage) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEnabled

`func (o *AgreementLanguage) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AgreementLanguage) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AgreementLanguage) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetId

`func (o *AgreementLanguage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgreementLanguage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgreementLanguage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgreementLanguage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocale

`func (o *AgreementLanguage) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *AgreementLanguage) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *AgreementLanguage) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetUserExperience

`func (o *AgreementLanguage) GetUserExperience() AgreementLanguageUserExperience`

GetUserExperience returns the UserExperience field if non-nil, zero value otherwise.

### GetUserExperienceOk

`func (o *AgreementLanguage) GetUserExperienceOk() (*AgreementLanguageUserExperience, bool)`

GetUserExperienceOk returns a tuple with the UserExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserExperience

`func (o *AgreementLanguage) SetUserExperience(v AgreementLanguageUserExperience)`

SetUserExperience sets UserExperience field to given value.

### HasUserExperience

`func (o *AgreementLanguage) HasUserExperience() bool`

HasUserExperience returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


