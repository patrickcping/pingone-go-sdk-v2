# DigitalWalletApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | [**ObjectApplication**](ObjectApplication.md) |  | 
**AppOpenUrl** | **string** | A string that specifies the URL sent in notifications to the user to communicate with the service. | 
**CreatedAt** | Pointer to **time.Time** | A string that specifies the date and time the credential digital wallet app was created. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the identifier (UUID) associated with the credential digital wallet app. | [optional] [readonly] 
**Name** | **string** | A string that specifies the name associated with the digital wallet app. | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UsesPingOneWalletSDK** | Pointer to **bool** | A boolean that specifies whether the user&#39;s wallet app uses the PingOne Wallet SDK. | [optional] 

## Methods

### NewDigitalWalletApplication

`func NewDigitalWalletApplication(application ObjectApplication, appOpenUrl string, name string, ) *DigitalWalletApplication`

NewDigitalWalletApplication instantiates a new DigitalWalletApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalWalletApplicationWithDefaults

`func NewDigitalWalletApplicationWithDefaults() *DigitalWalletApplication`

NewDigitalWalletApplicationWithDefaults instantiates a new DigitalWalletApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *DigitalWalletApplication) GetApplication() ObjectApplication`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *DigitalWalletApplication) GetApplicationOk() (*ObjectApplication, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *DigitalWalletApplication) SetApplication(v ObjectApplication)`

SetApplication sets Application field to given value.


### GetAppOpenUrl

`func (o *DigitalWalletApplication) GetAppOpenUrl() string`

GetAppOpenUrl returns the AppOpenUrl field if non-nil, zero value otherwise.

### GetAppOpenUrlOk

`func (o *DigitalWalletApplication) GetAppOpenUrlOk() (*string, bool)`

GetAppOpenUrlOk returns a tuple with the AppOpenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppOpenUrl

`func (o *DigitalWalletApplication) SetAppOpenUrl(v string)`

SetAppOpenUrl sets AppOpenUrl field to given value.


### GetCreatedAt

`func (o *DigitalWalletApplication) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DigitalWalletApplication) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DigitalWalletApplication) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DigitalWalletApplication) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnvironment

`func (o *DigitalWalletApplication) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DigitalWalletApplication) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DigitalWalletApplication) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *DigitalWalletApplication) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *DigitalWalletApplication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DigitalWalletApplication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DigitalWalletApplication) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DigitalWalletApplication) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DigitalWalletApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DigitalWalletApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DigitalWalletApplication) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *DigitalWalletApplication) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DigitalWalletApplication) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DigitalWalletApplication) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DigitalWalletApplication) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsesPingOneWalletSDK

`func (o *DigitalWalletApplication) GetUsesPingOneWalletSDK() bool`

GetUsesPingOneWalletSDK returns the UsesPingOneWalletSDK field if non-nil, zero value otherwise.

### GetUsesPingOneWalletSDKOk

`func (o *DigitalWalletApplication) GetUsesPingOneWalletSDKOk() (*bool, bool)`

GetUsesPingOneWalletSDKOk returns a tuple with the UsesPingOneWalletSDK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesPingOneWalletSDK

`func (o *DigitalWalletApplication) SetUsesPingOneWalletSDK(v bool)`

SetUsesPingOneWalletSDK sets UsesPingOneWalletSDK field to given value.

### HasUsesPingOneWalletSDK

`func (o *DigitalWalletApplication) HasUsesPingOneWalletSDK() bool`

HasUsesPingOneWalletSDK returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


