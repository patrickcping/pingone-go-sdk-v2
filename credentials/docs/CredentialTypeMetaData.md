# CredentialTypeMetaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackgroundImage** | Pointer to **string** | A string that specifies the URL to an image of the background to show in the credential. | [optional] 
**BgOpacityPercent** | Pointer to **int32** | A string that specifies the percent opacity of the background image in the credential. High percentage opacity may make reading text difficult. | [optional] 
**CardColor** | Pointer to **string** | A string that specifies the color to show on the credential. | [optional] 
**Columns** | Pointer to **int32** | An integer value between 1-3 that specifies the vertical layout of displayed fields in the credential type. | [optional] 
**Description** | Pointer to **string** | A string that specifies the description of the credential. | [optional] 
**Fields** | Pointer to [**[]CredentialTypeMetaDataFieldsInner**](CredentialTypeMetaDataFieldsInner.md) | An array of objects that specifies the fields on the credential. | [optional] 
**LogoImage** | Pointer to **string** | A string that specifies the URL to an image of the logo to show in the credential. | [optional] 
**Name** | Pointer to **string** | A string that specifies the name of the credential. | [optional] 
**TextColor** | Pointer to **string** | A string that specifies the color of the text to show on the credential. | [optional] 
**Version** | Pointer to **int32** | An integer that specifies the version of this credential. If not provided, the service assigns a version. | [optional] 

## Methods

### NewCredentialTypeMetaData

`func NewCredentialTypeMetaData() *CredentialTypeMetaData`

NewCredentialTypeMetaData instantiates a new CredentialTypeMetaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialTypeMetaDataWithDefaults

`func NewCredentialTypeMetaDataWithDefaults() *CredentialTypeMetaData`

NewCredentialTypeMetaDataWithDefaults instantiates a new CredentialTypeMetaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackgroundImage

`func (o *CredentialTypeMetaData) GetBackgroundImage() string`

GetBackgroundImage returns the BackgroundImage field if non-nil, zero value otherwise.

### GetBackgroundImageOk

`func (o *CredentialTypeMetaData) GetBackgroundImageOk() (*string, bool)`

GetBackgroundImageOk returns a tuple with the BackgroundImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundImage

`func (o *CredentialTypeMetaData) SetBackgroundImage(v string)`

SetBackgroundImage sets BackgroundImage field to given value.

### HasBackgroundImage

`func (o *CredentialTypeMetaData) HasBackgroundImage() bool`

HasBackgroundImage returns a boolean if a field has been set.

### GetBgOpacityPercent

`func (o *CredentialTypeMetaData) GetBgOpacityPercent() int32`

GetBgOpacityPercent returns the BgOpacityPercent field if non-nil, zero value otherwise.

### GetBgOpacityPercentOk

`func (o *CredentialTypeMetaData) GetBgOpacityPercentOk() (*int32, bool)`

GetBgOpacityPercentOk returns a tuple with the BgOpacityPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgOpacityPercent

`func (o *CredentialTypeMetaData) SetBgOpacityPercent(v int32)`

SetBgOpacityPercent sets BgOpacityPercent field to given value.

### HasBgOpacityPercent

`func (o *CredentialTypeMetaData) HasBgOpacityPercent() bool`

HasBgOpacityPercent returns a boolean if a field has been set.

### GetCardColor

`func (o *CredentialTypeMetaData) GetCardColor() string`

GetCardColor returns the CardColor field if non-nil, zero value otherwise.

### GetCardColorOk

`func (o *CredentialTypeMetaData) GetCardColorOk() (*string, bool)`

GetCardColorOk returns a tuple with the CardColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardColor

`func (o *CredentialTypeMetaData) SetCardColor(v string)`

SetCardColor sets CardColor field to given value.

### HasCardColor

`func (o *CredentialTypeMetaData) HasCardColor() bool`

HasCardColor returns a boolean if a field has been set.

### GetColumns

`func (o *CredentialTypeMetaData) GetColumns() int32`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *CredentialTypeMetaData) GetColumnsOk() (*int32, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *CredentialTypeMetaData) SetColumns(v int32)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *CredentialTypeMetaData) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetDescription

`func (o *CredentialTypeMetaData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CredentialTypeMetaData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CredentialTypeMetaData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CredentialTypeMetaData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFields

`func (o *CredentialTypeMetaData) GetFields() []CredentialTypeMetaDataFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CredentialTypeMetaData) GetFieldsOk() (*[]CredentialTypeMetaDataFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CredentialTypeMetaData) SetFields(v []CredentialTypeMetaDataFieldsInner)`

SetFields sets Fields field to given value.

### HasFields

`func (o *CredentialTypeMetaData) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetLogoImage

`func (o *CredentialTypeMetaData) GetLogoImage() string`

GetLogoImage returns the LogoImage field if non-nil, zero value otherwise.

### GetLogoImageOk

`func (o *CredentialTypeMetaData) GetLogoImageOk() (*string, bool)`

GetLogoImageOk returns a tuple with the LogoImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoImage

`func (o *CredentialTypeMetaData) SetLogoImage(v string)`

SetLogoImage sets LogoImage field to given value.

### HasLogoImage

`func (o *CredentialTypeMetaData) HasLogoImage() bool`

HasLogoImage returns a boolean if a field has been set.

### GetName

`func (o *CredentialTypeMetaData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CredentialTypeMetaData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CredentialTypeMetaData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CredentialTypeMetaData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTextColor

`func (o *CredentialTypeMetaData) GetTextColor() string`

GetTextColor returns the TextColor field if non-nil, zero value otherwise.

### GetTextColorOk

`func (o *CredentialTypeMetaData) GetTextColorOk() (*string, bool)`

GetTextColorOk returns a tuple with the TextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextColor

`func (o *CredentialTypeMetaData) SetTextColor(v string)`

SetTextColor sets TextColor field to given value.

### HasTextColor

`func (o *CredentialTypeMetaData) HasTextColor() bool`

HasTextColor returns a boolean if a field has been set.

### GetVersion

`func (o *CredentialTypeMetaData) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CredentialTypeMetaData) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CredentialTypeMetaData) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CredentialTypeMetaData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


