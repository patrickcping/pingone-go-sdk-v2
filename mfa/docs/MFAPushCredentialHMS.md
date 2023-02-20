# MFAPushCredentialHMS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumMFAPushCredentialAttrType**](EnumMFAPushCredentialAttrType.md) |  | 
**ClientId** | **string** | Used only if type is set to HMS. OAuth 2.0 Client ID from the Huawei Developers API console. | 
**ClientSecret** | **string** | Used only if type is set to HMS. The client secret associated with the OAuth 2.0 Client ID. | 

## Methods

### NewMFAPushCredentialHMS

`func NewMFAPushCredentialHMS(type_ EnumMFAPushCredentialAttrType, clientId string, clientSecret string, ) *MFAPushCredentialHMS`

NewMFAPushCredentialHMS instantiates a new MFAPushCredentialHMS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFAPushCredentialHMSWithDefaults

`func NewMFAPushCredentialHMSWithDefaults() *MFAPushCredentialHMS`

NewMFAPushCredentialHMSWithDefaults instantiates a new MFAPushCredentialHMS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MFAPushCredentialHMS) GetType() EnumMFAPushCredentialAttrType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MFAPushCredentialHMS) GetTypeOk() (*EnumMFAPushCredentialAttrType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MFAPushCredentialHMS) SetType(v EnumMFAPushCredentialAttrType)`

SetType sets Type field to given value.


### GetClientId

`func (o *MFAPushCredentialHMS) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *MFAPushCredentialHMS) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *MFAPushCredentialHMS) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *MFAPushCredentialHMS) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *MFAPushCredentialHMS) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *MFAPushCredentialHMS) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


