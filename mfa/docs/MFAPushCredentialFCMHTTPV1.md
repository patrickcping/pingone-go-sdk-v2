# MFAPushCredentialFCMHTTPV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumMFAPushCredentialAttrType**](EnumMFAPushCredentialAttrType.md) |  | 
**GoogleServiceAccountCredentials** | **string** | Used when &#x60;type&#x60; is set to &#x60;FCM_HTTP_V1&#x60;. The value should be the contents of the JSON file that represents your Service Account Credentials. | 

## Methods

### NewMFAPushCredentialFCMHTTPV1

`func NewMFAPushCredentialFCMHTTPV1(type_ EnumMFAPushCredentialAttrType, googleServiceAccountCredentials string, ) *MFAPushCredentialFCMHTTPV1`

NewMFAPushCredentialFCMHTTPV1 instantiates a new MFAPushCredentialFCMHTTPV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFAPushCredentialFCMHTTPV1WithDefaults

`func NewMFAPushCredentialFCMHTTPV1WithDefaults() *MFAPushCredentialFCMHTTPV1`

NewMFAPushCredentialFCMHTTPV1WithDefaults instantiates a new MFAPushCredentialFCMHTTPV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MFAPushCredentialFCMHTTPV1) GetType() EnumMFAPushCredentialAttrType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MFAPushCredentialFCMHTTPV1) GetTypeOk() (*EnumMFAPushCredentialAttrType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MFAPushCredentialFCMHTTPV1) SetType(v EnumMFAPushCredentialAttrType)`

SetType sets Type field to given value.


### GetGoogleServiceAccountCredentials

`func (o *MFAPushCredentialFCMHTTPV1) GetGoogleServiceAccountCredentials() string`

GetGoogleServiceAccountCredentials returns the GoogleServiceAccountCredentials field if non-nil, zero value otherwise.

### GetGoogleServiceAccountCredentialsOk

`func (o *MFAPushCredentialFCMHTTPV1) GetGoogleServiceAccountCredentialsOk() (*string, bool)`

GetGoogleServiceAccountCredentialsOk returns a tuple with the GoogleServiceAccountCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleServiceAccountCredentials

`func (o *MFAPushCredentialFCMHTTPV1) SetGoogleServiceAccountCredentials(v string)`

SetGoogleServiceAccountCredentials sets GoogleServiceAccountCredentials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


