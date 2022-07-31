# Agreement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsentCountsUpdatedAt** | Pointer to **string** | The time the consent count metric was last updated. This value is typically updated once every 24 hours. | [optional] [readonly] 
**Description** | Pointer to **string** | A string that specifies the description of the agreement. | [optional] 
**Enabled** | **bool** | A boolean that specifies the current enabled state of the agreement. This is a required property. The agreement must support the default language to be enabled. It cannot be disabled if it is referenced by a sign-on action. When an agreement is disabled, it is not used anywhere it is configured across PingOne. | 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**ExpiredUserConsents** | Pointer to **int32** | An integer that identifies the number of users who have consented to the agreement, but their consent has expired. This value is last calculated at the consentCountsUpdatedAt time. | [optional] [readonly] 
**Id** | Pointer to **string** | A string that specifies the agreement ID. | [optional] [readonly] 
**Name** | **string** | A string that specifies the name of the agreement resource. This is a required property. | 
**ReconsentPeriodDays** | Pointer to **int32** | A number that represents the number of days until a consent to this agreement expires. | [optional] 
**TotalUserConsents** | Pointer to **int32** | An integer that identifies the total number of users who have consented to the agreement. This value is last calculated at the consentCountsUpdatedAt time. | [optional] [readonly] 

## Methods

### NewAgreement

`func NewAgreement(enabled bool, name string, ) *Agreement`

NewAgreement instantiates a new Agreement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementWithDefaults

`func NewAgreementWithDefaults() *Agreement`

NewAgreementWithDefaults instantiates a new Agreement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsentCountsUpdatedAt

`func (o *Agreement) GetConsentCountsUpdatedAt() string`

GetConsentCountsUpdatedAt returns the ConsentCountsUpdatedAt field if non-nil, zero value otherwise.

### GetConsentCountsUpdatedAtOk

`func (o *Agreement) GetConsentCountsUpdatedAtOk() (*string, bool)`

GetConsentCountsUpdatedAtOk returns a tuple with the ConsentCountsUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentCountsUpdatedAt

`func (o *Agreement) SetConsentCountsUpdatedAt(v string)`

SetConsentCountsUpdatedAt sets ConsentCountsUpdatedAt field to given value.

### HasConsentCountsUpdatedAt

`func (o *Agreement) HasConsentCountsUpdatedAt() bool`

HasConsentCountsUpdatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *Agreement) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Agreement) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Agreement) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Agreement) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *Agreement) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Agreement) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Agreement) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnvironment

`func (o *Agreement) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Agreement) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Agreement) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *Agreement) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetExpiredUserConsents

`func (o *Agreement) GetExpiredUserConsents() int32`

GetExpiredUserConsents returns the ExpiredUserConsents field if non-nil, zero value otherwise.

### GetExpiredUserConsentsOk

`func (o *Agreement) GetExpiredUserConsentsOk() (*int32, bool)`

GetExpiredUserConsentsOk returns a tuple with the ExpiredUserConsents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredUserConsents

`func (o *Agreement) SetExpiredUserConsents(v int32)`

SetExpiredUserConsents sets ExpiredUserConsents field to given value.

### HasExpiredUserConsents

`func (o *Agreement) HasExpiredUserConsents() bool`

HasExpiredUserConsents returns a boolean if a field has been set.

### GetId

`func (o *Agreement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Agreement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Agreement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Agreement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Agreement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Agreement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Agreement) SetName(v string)`

SetName sets Name field to given value.


### GetReconsentPeriodDays

`func (o *Agreement) GetReconsentPeriodDays() int32`

GetReconsentPeriodDays returns the ReconsentPeriodDays field if non-nil, zero value otherwise.

### GetReconsentPeriodDaysOk

`func (o *Agreement) GetReconsentPeriodDaysOk() (*int32, bool)`

GetReconsentPeriodDaysOk returns a tuple with the ReconsentPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconsentPeriodDays

`func (o *Agreement) SetReconsentPeriodDays(v int32)`

SetReconsentPeriodDays sets ReconsentPeriodDays field to given value.

### HasReconsentPeriodDays

`func (o *Agreement) HasReconsentPeriodDays() bool`

HasReconsentPeriodDays returns a boolean if a field has been set.

### GetTotalUserConsents

`func (o *Agreement) GetTotalUserConsents() int32`

GetTotalUserConsents returns the TotalUserConsents field if non-nil, zero value otherwise.

### GetTotalUserConsentsOk

`func (o *Agreement) GetTotalUserConsentsOk() (*int32, bool)`

GetTotalUserConsentsOk returns a tuple with the TotalUserConsents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUserConsents

`func (o *Agreement) SetTotalUserConsents(v int32)`

SetTotalUserConsents sets TotalUserConsents field to given value.

### HasTotalUserConsents

`func (o *Agreement) HasTotalUserConsents() bool`

HasTotalUserConsents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


