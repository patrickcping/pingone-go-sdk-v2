# RegistrationDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Device type. Must be one of [EMAIL, FIDO2, MOBILE, SMS, TOTP, VOICE, WHATSAPP]. | 
**Title** | **string** | Title for the device. | 
**Description** | Pointer to **string** | Description for the device (Max 1000 characters). | [optional] 
**IconSrc** | Pointer to **string** | Icon image source to display for the device (Max 500 characters). | [optional] 

## Methods

### NewRegistrationDevice

`func NewRegistrationDevice(type_ string, title string, ) *RegistrationDevice`

NewRegistrationDevice instantiates a new RegistrationDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationDeviceWithDefaults

`func NewRegistrationDeviceWithDefaults() *RegistrationDevice`

NewRegistrationDeviceWithDefaults instantiates a new RegistrationDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RegistrationDevice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegistrationDevice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegistrationDevice) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *RegistrationDevice) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RegistrationDevice) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RegistrationDevice) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *RegistrationDevice) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RegistrationDevice) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RegistrationDevice) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RegistrationDevice) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIconSrc

`func (o *RegistrationDevice) GetIconSrc() string`

GetIconSrc returns the IconSrc field if non-nil, zero value otherwise.

### GetIconSrcOk

`func (o *RegistrationDevice) GetIconSrcOk() (*string, bool)`

GetIconSrcOk returns a tuple with the IconSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconSrc

`func (o *RegistrationDevice) SetIconSrc(v string)`

SetIconSrc sets IconSrc field to given value.

### HasIconSrc

`func (o *RegistrationDevice) HasIconSrc() bool`

HasIconSrc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


