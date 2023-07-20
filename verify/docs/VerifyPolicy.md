# VerifyPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Name** | **string** | Name displayed in PingOne Admin UI. | 
**Description** | Pointer to **string** | Description displayed in PingOne Admin UI, 1-1024 characters. | [optional] 
**Default** | Pointer to **bool** | Required as true to set this verify policy as the default policy for the environment; otherwise optional and defaults to false. | [optional] 
**GovernmentId** | Pointer to [**GovernmentIdConfiguration**](GovernmentIdConfiguration.md) |  | [optional] 
**FacialComparison** | Pointer to [**FacialComparisonConfiguration**](FacialComparisonConfiguration.md) |  | [optional] 
**Liveness** | Pointer to [**LivenessConfiguration**](LivenessConfiguration.md) |  | [optional] 
**Email** | Pointer to [**OTPDeviceConfiguration**](OTPDeviceConfiguration.md) |  | [optional] 
**Phone** | Pointer to [**OTPDeviceConfiguration**](OTPDeviceConfiguration.md) |  | [optional] 
**Transaction** | Pointer to [**TransactionConfiguration**](TransactionConfiguration.md) |  | [optional] 
**Voice** | Pointer to [**VoiceConfiguration**](VoiceConfiguration.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewVerifyPolicy

`func NewVerifyPolicy(name string, ) *VerifyPolicy`

NewVerifyPolicy instantiates a new VerifyPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyPolicyWithDefaults

`func NewVerifyPolicyWithDefaults() *VerifyPolicy`

NewVerifyPolicyWithDefaults instantiates a new VerifyPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VerifyPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VerifyPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VerifyPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VerifyPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *VerifyPolicy) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *VerifyPolicy) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *VerifyPolicy) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *VerifyPolicy) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetName

`func (o *VerifyPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VerifyPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VerifyPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VerifyPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VerifyPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VerifyPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VerifyPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefault

`func (o *VerifyPolicy) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *VerifyPolicy) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *VerifyPolicy) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *VerifyPolicy) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetGovernmentId

`func (o *VerifyPolicy) GetGovernmentId() GovernmentIdConfiguration`

GetGovernmentId returns the GovernmentId field if non-nil, zero value otherwise.

### GetGovernmentIdOk

`func (o *VerifyPolicy) GetGovernmentIdOk() (*GovernmentIdConfiguration, bool)`

GetGovernmentIdOk returns a tuple with the GovernmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovernmentId

`func (o *VerifyPolicy) SetGovernmentId(v GovernmentIdConfiguration)`

SetGovernmentId sets GovernmentId field to given value.

### HasGovernmentId

`func (o *VerifyPolicy) HasGovernmentId() bool`

HasGovernmentId returns a boolean if a field has been set.

### GetFacialComparison

`func (o *VerifyPolicy) GetFacialComparison() FacialComparisonConfiguration`

GetFacialComparison returns the FacialComparison field if non-nil, zero value otherwise.

### GetFacialComparisonOk

`func (o *VerifyPolicy) GetFacialComparisonOk() (*FacialComparisonConfiguration, bool)`

GetFacialComparisonOk returns a tuple with the FacialComparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacialComparison

`func (o *VerifyPolicy) SetFacialComparison(v FacialComparisonConfiguration)`

SetFacialComparison sets FacialComparison field to given value.

### HasFacialComparison

`func (o *VerifyPolicy) HasFacialComparison() bool`

HasFacialComparison returns a boolean if a field has been set.

### GetLiveness

`func (o *VerifyPolicy) GetLiveness() LivenessConfiguration`

GetLiveness returns the Liveness field if non-nil, zero value otherwise.

### GetLivenessOk

`func (o *VerifyPolicy) GetLivenessOk() (*LivenessConfiguration, bool)`

GetLivenessOk returns a tuple with the Liveness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveness

`func (o *VerifyPolicy) SetLiveness(v LivenessConfiguration)`

SetLiveness sets Liveness field to given value.

### HasLiveness

`func (o *VerifyPolicy) HasLiveness() bool`

HasLiveness returns a boolean if a field has been set.

### GetEmail

`func (o *VerifyPolicy) GetEmail() OTPDeviceConfiguration`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *VerifyPolicy) GetEmailOk() (*OTPDeviceConfiguration, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *VerifyPolicy) SetEmail(v OTPDeviceConfiguration)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *VerifyPolicy) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *VerifyPolicy) GetPhone() OTPDeviceConfiguration`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *VerifyPolicy) GetPhoneOk() (*OTPDeviceConfiguration, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *VerifyPolicy) SetPhone(v OTPDeviceConfiguration)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *VerifyPolicy) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetTransaction

`func (o *VerifyPolicy) GetTransaction() TransactionConfiguration`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *VerifyPolicy) GetTransactionOk() (*TransactionConfiguration, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *VerifyPolicy) SetTransaction(v TransactionConfiguration)`

SetTransaction sets Transaction field to given value.

### HasTransaction

`func (o *VerifyPolicy) HasTransaction() bool`

HasTransaction returns a boolean if a field has been set.

### GetVoice

`func (o *VerifyPolicy) GetVoice() VoiceConfiguration`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *VerifyPolicy) GetVoiceOk() (*VoiceConfiguration, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *VerifyPolicy) SetVoice(v VoiceConfiguration)`

SetVoice sets Voice field to given value.

### HasVoice

`func (o *VerifyPolicy) HasVoice() bool`

HasVoice returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VerifyPolicy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VerifyPolicy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VerifyPolicy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VerifyPolicy) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VerifyPolicy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VerifyPolicy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VerifyPolicy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VerifyPolicy) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


