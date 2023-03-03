# EntityArrayEmbeddedLanguagesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agreement** | Pointer to [**AgreementLanguageAgreement**](AgreementLanguageAgreement.md) |  | [optional] 
**CurrentRevision** | Pointer to [**AgreementLanguageCurrentRevision**](AgreementLanguageCurrentRevision.md) |  | [optional] 
**DisplayName** | **string** | A string that is used as the title of the agreement for the language presented to the user. This is a required property. | 
**Enabled** | **bool** | Specifies whether this language is enabled for the environment. This property value must be set to false when creating a language. | 
**Id** | Pointer to **string** | The resource’s unique identifier. | [optional] [readonly] 
**Locale** | **string** | An ISO standard language code. For more information about standard language codes, see ISO Language Code Table. | 
**UserExperience** | Pointer to [**AgreementLanguageUserExperience**](AgreementLanguageUserExperience.md) |  | [optional] 
**Default** | **bool** | Specifies whether this language is the default for the environment. This property value must be set to false when creating a language resource. It can be set to true only after the language is enabled and after the localization of an agreement resource is complete when agreements are used for the environment. | 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Name** | Pointer to **string** | The language name. If omitted, the ISO standard name is used. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the language resource was created. | [optional] [readonly] 
**CustomerAdded** | Pointer to **bool** | Specifies whether this language was added by a customer administrator. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The time the language resource was last updated. | [optional] [readonly] 

## Methods

### NewEntityArrayEmbeddedLanguagesInner

`func NewEntityArrayEmbeddedLanguagesInner(displayName string, enabled bool, locale string, default_ bool, ) *EntityArrayEmbeddedLanguagesInner`

NewEntityArrayEmbeddedLanguagesInner instantiates a new EntityArrayEmbeddedLanguagesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityArrayEmbeddedLanguagesInnerWithDefaults

`func NewEntityArrayEmbeddedLanguagesInnerWithDefaults() *EntityArrayEmbeddedLanguagesInner`

NewEntityArrayEmbeddedLanguagesInnerWithDefaults instantiates a new EntityArrayEmbeddedLanguagesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgreement

`func (o *EntityArrayEmbeddedLanguagesInner) GetAgreement() AgreementLanguageAgreement`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *EntityArrayEmbeddedLanguagesInner) GetAgreementOk() (*AgreementLanguageAgreement, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *EntityArrayEmbeddedLanguagesInner) SetAgreement(v AgreementLanguageAgreement)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *EntityArrayEmbeddedLanguagesInner) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetCurrentRevision

`func (o *EntityArrayEmbeddedLanguagesInner) GetCurrentRevision() AgreementLanguageCurrentRevision`

GetCurrentRevision returns the CurrentRevision field if non-nil, zero value otherwise.

### GetCurrentRevisionOk

`func (o *EntityArrayEmbeddedLanguagesInner) GetCurrentRevisionOk() (*AgreementLanguageCurrentRevision, bool)`

GetCurrentRevisionOk returns a tuple with the CurrentRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRevision

`func (o *EntityArrayEmbeddedLanguagesInner) SetCurrentRevision(v AgreementLanguageCurrentRevision)`

SetCurrentRevision sets CurrentRevision field to given value.

### HasCurrentRevision

`func (o *EntityArrayEmbeddedLanguagesInner) HasCurrentRevision() bool`

HasCurrentRevision returns a boolean if a field has been set.

### GetDisplayName

`func (o *EntityArrayEmbeddedLanguagesInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EntityArrayEmbeddedLanguagesInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EntityArrayEmbeddedLanguagesInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEnabled

`func (o *EntityArrayEmbeddedLanguagesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EntityArrayEmbeddedLanguagesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EntityArrayEmbeddedLanguagesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetId

`func (o *EntityArrayEmbeddedLanguagesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityArrayEmbeddedLanguagesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityArrayEmbeddedLanguagesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntityArrayEmbeddedLanguagesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocale

`func (o *EntityArrayEmbeddedLanguagesInner) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *EntityArrayEmbeddedLanguagesInner) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *EntityArrayEmbeddedLanguagesInner) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetUserExperience

`func (o *EntityArrayEmbeddedLanguagesInner) GetUserExperience() AgreementLanguageUserExperience`

GetUserExperience returns the UserExperience field if non-nil, zero value otherwise.

### GetUserExperienceOk

`func (o *EntityArrayEmbeddedLanguagesInner) GetUserExperienceOk() (*AgreementLanguageUserExperience, bool)`

GetUserExperienceOk returns a tuple with the UserExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserExperience

`func (o *EntityArrayEmbeddedLanguagesInner) SetUserExperience(v AgreementLanguageUserExperience)`

SetUserExperience sets UserExperience field to given value.

### HasUserExperience

`func (o *EntityArrayEmbeddedLanguagesInner) HasUserExperience() bool`

HasUserExperience returns a boolean if a field has been set.

### GetDefault

`func (o *EntityArrayEmbeddedLanguagesInner) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *EntityArrayEmbeddedLanguagesInner) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *EntityArrayEmbeddedLanguagesInner) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetEnvironment

`func (o *EntityArrayEmbeddedLanguagesInner) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EntityArrayEmbeddedLanguagesInner) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EntityArrayEmbeddedLanguagesInner) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EntityArrayEmbeddedLanguagesInner) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetName

`func (o *EntityArrayEmbeddedLanguagesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityArrayEmbeddedLanguagesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityArrayEmbeddedLanguagesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntityArrayEmbeddedLanguagesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EntityArrayEmbeddedLanguagesInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EntityArrayEmbeddedLanguagesInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EntityArrayEmbeddedLanguagesInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EntityArrayEmbeddedLanguagesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomerAdded

`func (o *EntityArrayEmbeddedLanguagesInner) GetCustomerAdded() bool`

GetCustomerAdded returns the CustomerAdded field if non-nil, zero value otherwise.

### GetCustomerAddedOk

`func (o *EntityArrayEmbeddedLanguagesInner) GetCustomerAddedOk() (*bool, bool)`

GetCustomerAddedOk returns a tuple with the CustomerAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAdded

`func (o *EntityArrayEmbeddedLanguagesInner) SetCustomerAdded(v bool)`

SetCustomerAdded sets CustomerAdded field to given value.

### HasCustomerAdded

`func (o *EntityArrayEmbeddedLanguagesInner) HasCustomerAdded() bool`

HasCustomerAdded returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EntityArrayEmbeddedLanguagesInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EntityArrayEmbeddedLanguagesInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EntityArrayEmbeddedLanguagesInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EntityArrayEmbeddedLanguagesInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


