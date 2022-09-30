# LanguageLocalizationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | The time the language localization status resource was created. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | The resource’s unique identifier. | [optional] [readonly] 
**Locale** | Pointer to [**LanguageLocalizationStatusLocale**](LanguageLocalizationStatusLocale.md) |  | [optional] 
**Service** | **string** | The name of the service associated with this resource. | 
**LocalizationComplete** | Pointer to **bool** |  | [optional] 
**StatusDetails** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** | The time the language localization status resource was last updated. | [optional] [readonly] 

## Methods

### NewLanguageLocalizationStatus

`func NewLanguageLocalizationStatus(service string, ) *LanguageLocalizationStatus`

NewLanguageLocalizationStatus instantiates a new LanguageLocalizationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanguageLocalizationStatusWithDefaults

`func NewLanguageLocalizationStatusWithDefaults() *LanguageLocalizationStatus`

NewLanguageLocalizationStatusWithDefaults instantiates a new LanguageLocalizationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *LanguageLocalizationStatus) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LanguageLocalizationStatus) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LanguageLocalizationStatus) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LanguageLocalizationStatus) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnvironment

`func (o *LanguageLocalizationStatus) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *LanguageLocalizationStatus) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *LanguageLocalizationStatus) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *LanguageLocalizationStatus) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *LanguageLocalizationStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LanguageLocalizationStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LanguageLocalizationStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LanguageLocalizationStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocale

`func (o *LanguageLocalizationStatus) GetLocale() LanguageLocalizationStatusLocale`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *LanguageLocalizationStatus) GetLocaleOk() (*LanguageLocalizationStatusLocale, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *LanguageLocalizationStatus) SetLocale(v LanguageLocalizationStatusLocale)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *LanguageLocalizationStatus) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetService

`func (o *LanguageLocalizationStatus) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *LanguageLocalizationStatus) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *LanguageLocalizationStatus) SetService(v string)`

SetService sets Service field to given value.


### GetLocalizationComplete

`func (o *LanguageLocalizationStatus) GetLocalizationComplete() bool`

GetLocalizationComplete returns the LocalizationComplete field if non-nil, zero value otherwise.

### GetLocalizationCompleteOk

`func (o *LanguageLocalizationStatus) GetLocalizationCompleteOk() (*bool, bool)`

GetLocalizationCompleteOk returns a tuple with the LocalizationComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizationComplete

`func (o *LanguageLocalizationStatus) SetLocalizationComplete(v bool)`

SetLocalizationComplete sets LocalizationComplete field to given value.

### HasLocalizationComplete

`func (o *LanguageLocalizationStatus) HasLocalizationComplete() bool`

HasLocalizationComplete returns a boolean if a field has been set.

### GetStatusDetails

`func (o *LanguageLocalizationStatus) GetStatusDetails() string`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *LanguageLocalizationStatus) GetStatusDetailsOk() (*string, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *LanguageLocalizationStatus) SetStatusDetails(v string)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *LanguageLocalizationStatus) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LanguageLocalizationStatus) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LanguageLocalizationStatus) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LanguageLocalizationStatus) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LanguageLocalizationStatus) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


