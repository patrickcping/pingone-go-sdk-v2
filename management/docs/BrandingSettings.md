# BrandingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Specifies the resource’s unique identifier. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**CompanyName** | Pointer to **string** | The company name associated with the specified environment. | [optional] 
**Logo** | Pointer to [**BrandingSettingsLogo**](BrandingSettingsLogo.md) |  | [optional] 

## Methods

### NewBrandingSettings

`func NewBrandingSettings() *BrandingSettings`

NewBrandingSettings instantiates a new BrandingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandingSettingsWithDefaults

`func NewBrandingSettingsWithDefaults() *BrandingSettings`

NewBrandingSettingsWithDefaults instantiates a new BrandingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BrandingSettings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BrandingSettings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BrandingSettings) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BrandingSettings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *BrandingSettings) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *BrandingSettings) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *BrandingSettings) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *BrandingSettings) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetCompanyName

`func (o *BrandingSettings) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *BrandingSettings) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *BrandingSettings) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *BrandingSettings) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetLogo

`func (o *BrandingSettings) GetLogo() BrandingSettingsLogo`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *BrandingSettings) GetLogoOk() (*BrandingSettingsLogo, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *BrandingSettings) SetLogo(v BrandingSettingsLogo)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *BrandingSettings) HasLogo() bool`

HasLogo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


