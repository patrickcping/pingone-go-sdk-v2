# RecaptchaConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | The time the configuration was created. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Modified** | Pointer to **time.Time** | The time the configuration was last updated. | [optional] [readonly] 
**SiteKey** | **string** | A string that specifies the public site key for the Recaptcha configuration provided by Google. | 
**SecretKey** | **string** | A string that specifies the confidential secret key for the Recaptcha configuration provided by Google. | 

## Methods

### NewRecaptchaConfiguration

`func NewRecaptchaConfiguration(siteKey string, secretKey string, ) *RecaptchaConfiguration`

NewRecaptchaConfiguration instantiates a new RecaptchaConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecaptchaConfigurationWithDefaults

`func NewRecaptchaConfigurationWithDefaults() *RecaptchaConfiguration`

NewRecaptchaConfigurationWithDefaults instantiates a new RecaptchaConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *RecaptchaConfiguration) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RecaptchaConfiguration) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RecaptchaConfiguration) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RecaptchaConfiguration) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEnvironment

`func (o *RecaptchaConfiguration) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *RecaptchaConfiguration) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *RecaptchaConfiguration) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *RecaptchaConfiguration) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetModified

`func (o *RecaptchaConfiguration) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *RecaptchaConfiguration) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *RecaptchaConfiguration) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *RecaptchaConfiguration) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetSiteKey

`func (o *RecaptchaConfiguration) GetSiteKey() string`

GetSiteKey returns the SiteKey field if non-nil, zero value otherwise.

### GetSiteKeyOk

`func (o *RecaptchaConfiguration) GetSiteKeyOk() (*string, bool)`

GetSiteKeyOk returns a tuple with the SiteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteKey

`func (o *RecaptchaConfiguration) SetSiteKey(v string)`

SetSiteKey sets SiteKey field to given value.


### GetSecretKey

`func (o *RecaptchaConfiguration) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *RecaptchaConfiguration) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *RecaptchaConfiguration) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


