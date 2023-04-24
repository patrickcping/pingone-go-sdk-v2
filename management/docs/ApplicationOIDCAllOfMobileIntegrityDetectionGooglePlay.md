# ApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DecryptionKey** | Pointer to **string** | Play Integrity verdict decryption key from your Google Play Services account. This parameter must be provided if you have set &#x60;mobile.integrityDetection.googlePlay.verificationType&#x60; to &#x60;INTERNAL&#x60;. | [optional] 
**ServiceAccountCredentials** | Pointer to **string** | Contents of the JSON file that represents your Service Account Credentials. This parameter must be provided if you have set &#x60;mobile.integrityDetection.googlePlay.verificationType&#x60; to &#x60;GOOGLE&#x60;. | [optional] 
**VerificationKey** | Pointer to **string** | Play Integrity verdict signature verification key from your Google Play Services account. This parameter must be provided if you have set &#x60;mobile.integrityDetection.googlePlay.verificationType&#x60; to &#x60;INTERNAL&#x60;. | [optional] 
**VerificationType** | Pointer to [**EnumApplicationNativeGooglePlayVerificationType**](EnumApplicationNativeGooglePlayVerificationType.md) |  | [optional] 

## Methods

### NewApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay

`func NewApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay() *ApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay`

NewApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay instantiates a new ApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationOIDCAllOfMobileIntegrityDetectionGooglePlayWithDefaults

`func NewApplicationOIDCAllOfMobileIntegrityDetectionGooglePlayWithDefaults() *ApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay`

NewApplicationOIDCAllOfMobileIntegrityDetectionGooglePlayWithDefaults instantiates a new ApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecryptionKey

`func (o *ApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay) GetDecryptionKey() string`

GetDecryptionKey returns the DecryptionKey field if non-nil, zero value otherwise.

### GetDecryptionKeyOk

`func (o *ApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay) GetDecryptionKeyOk() (*string, bool)`

GetDecryptionKeyOk returns a tuple with the DecryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecryptionKey

`func (o *ApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay) SetDecryptionKey(v string)`

SetDecryptionKey sets DecryptionKey field to given value.

### HasDecryptionKey

`func (o *ApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay) HasDecryptionKey() bool`

HasDecryptionKey returns a boolean if a field has been set.

### GetServiceAccountCredentials

`func (o *ApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay) GetServiceAccountCredentials() string`

GetServiceAccountCredentials returns the ServiceAccountCredentials field if non-nil, zero value otherwise.

### GetServiceAccountCredentialsOk

`func (o *ApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay) GetServiceAccountCredentialsOk() (*string, bool)`

GetServiceAccountCredentialsOk returns a tuple with the ServiceAccountCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredentials

`func (o *ApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay) SetServiceAccountCredentials(v string)`

SetServiceAccountCredentials sets ServiceAccountCredentials field to given value.

### HasServiceAccountCredentials

`func (o *ApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay) HasServiceAccountCredentials() bool`

HasServiceAccountCredentials returns a boolean if a field has been set.

### GetVerificationKey

`func (o *ApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay) GetVerificationKey() string`

GetVerificationKey returns the VerificationKey field if non-nil, zero value otherwise.

### GetVerificationKeyOk

`func (o *ApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay) GetVerificationKeyOk() (*string, bool)`

GetVerificationKeyOk returns a tuple with the VerificationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationKey

`func (o *ApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay) SetVerificationKey(v string)`

SetVerificationKey sets VerificationKey field to given value.

### HasVerificationKey

`func (o *ApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay) HasVerificationKey() bool`

HasVerificationKey returns a boolean if a field has been set.

### GetVerificationType

`func (o *ApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay) GetVerificationType() EnumApplicationNativeGooglePlayVerificationType`

GetVerificationType returns the VerificationType field if non-nil, zero value otherwise.

### GetVerificationTypeOk

`func (o *ApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay) GetVerificationTypeOk() (*EnumApplicationNativeGooglePlayVerificationType, bool)`

GetVerificationTypeOk returns a tuple with the VerificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationType

`func (o *ApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay) SetVerificationType(v EnumApplicationNativeGooglePlayVerificationType)`

SetVerificationType sets VerificationType field to given value.

### HasVerificationType

`func (o *ApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay) HasVerificationType() bool`

HasVerificationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


