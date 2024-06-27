# Language

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Default** | **bool** | Specifies whether this language is the default for the environment. This property value must be set to false when creating a language resource. It can be set to true only after the language is enabled and after the localization of an agreement resource is complete when agreements are used for the environment. | 
**Enabled** | **bool** | Specifies whether this language is enabled for the environment. This property value must be set to false when creating a language. | 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | The resource’s unique identifier. | [optional] [readonly] 
**Locale** | **string** | An ISO standard language code. For more information about standard language codes, see ISO Language Code Table. | 
**Name** | Pointer to **string** | The language name. If omitted, the ISO standard name is used. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the language resource was created. | [optional] [readonly] 
**CustomerAdded** | Pointer to **bool** | Specifies whether this language was added by a customer administrator. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The time the language resource was last updated. | [optional] [readonly] 

## Methods

### NewLanguage

`func NewLanguage(default_ bool, enabled bool, locale string, ) *Language`

NewLanguage instantiates a new Language object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanguageWithDefaults

`func NewLanguageWithDefaults() *Language`

NewLanguageWithDefaults instantiates a new Language object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Language) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Language) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Language) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Language) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetDefault

`func (o *Language) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Language) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Language) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetEnabled

`func (o *Language) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Language) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Language) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnvironment

`func (o *Language) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Language) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Language) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *Language) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *Language) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Language) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Language) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Language) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocale

`func (o *Language) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *Language) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *Language) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetName

`func (o *Language) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Language) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Language) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Language) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Language) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Language) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Language) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Language) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomerAdded

`func (o *Language) GetCustomerAdded() bool`

GetCustomerAdded returns the CustomerAdded field if non-nil, zero value otherwise.

### GetCustomerAddedOk

`func (o *Language) GetCustomerAddedOk() (*bool, bool)`

GetCustomerAddedOk returns a tuple with the CustomerAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAdded

`func (o *Language) SetCustomerAdded(v bool)`

SetCustomerAdded sets CustomerAdded field to given value.

### HasCustomerAdded

`func (o *Language) HasCustomerAdded() bool`

HasCustomerAdded returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Language) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Language) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Language) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Language) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


