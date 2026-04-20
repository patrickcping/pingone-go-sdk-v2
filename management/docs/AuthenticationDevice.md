# AuthenticationDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Device type. Must be one of [EMAIL, FIDO2, MAGIC_LINK, MOBILE, SMS, TOTP, VOICE, WHATSAPP]. | 
**Title** | **string** | Title for the device. | 
**Description** | Pointer to **string** | Description for the device (Max 1000 characters). | [optional] 
**IconSrc** | Pointer to **string** | Icon image source to display for the device (Max 500 characters). | [optional] 

## Methods

### NewAuthenticationDevice

`func NewAuthenticationDevice(type_ string, title string, ) *AuthenticationDevice`

NewAuthenticationDevice instantiates a new AuthenticationDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationDeviceWithDefaults

`func NewAuthenticationDeviceWithDefaults() *AuthenticationDevice`

NewAuthenticationDeviceWithDefaults instantiates a new AuthenticationDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthenticationDevice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthenticationDevice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthenticationDevice) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *AuthenticationDevice) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AuthenticationDevice) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AuthenticationDevice) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *AuthenticationDevice) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthenticationDevice) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthenticationDevice) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthenticationDevice) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIconSrc

`func (o *AuthenticationDevice) GetIconSrc() string`

GetIconSrc returns the IconSrc field if non-nil, zero value otherwise.

### GetIconSrcOk

`func (o *AuthenticationDevice) GetIconSrcOk() (*string, bool)`

GetIconSrcOk returns a tuple with the IconSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconSrc

`func (o *AuthenticationDevice) SetIconSrc(v string)`

SetIconSrc sets IconSrc field to given value.

### HasIconSrc

`func (o *AuthenticationDevice) HasIconSrc() bool`

HasIconSrc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


