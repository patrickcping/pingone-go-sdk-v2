# RiskPredictorDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Name** | **string** | A string type. A unique, friendly name for the predictor. This name is displayed in the Risk Policies UI, when the admin is asked to define the overrides and weights. | 
**CompactName** | **string** | A string type. A unique name for the predictor. This property is immutable; it cannot be modified after initial creation. The value must be alpha-numeric, with no special characters or spaces. This name is used in the API both for policy configuration, and in the Risk Evaluation response (under details). | 
**Type** | [**EnumPredictorType**](EnumPredictorType.md) |  | 
**Description** | Pointer to **string** | A string type. This specifies the description of the risk predictor. Maximum length is 1024 characters. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was updated. | [optional] [readonly] 
**Licensed** | Pointer to **bool** | Indicates whether PingOne Risk is licensed for the environment. | [optional] [readonly] 
**Deletable** | Pointer to **bool** | A boolean to indicate whether the predictor is deletable in the environment. | [optional] [readonly] 
**Default** | Pointer to [**RiskPredictorCommonDefault**](RiskPredictorCommonDefault.md) |  | [optional] 
**Condition** | Pointer to [**RiskPredictorCommonCondition**](RiskPredictorCommonCondition.md) |  | [optional] 
**Detect** | [**EnumPredictorNewDeviceDetectType**](EnumPredictorNewDeviceDetectType.md) |  | 
**ShouldValidatePayloadSignature** | Pointer to **bool** | Relevant only for Suspicious Device predictors. If &#x60;shouldValidatePayloadSignature&#x60; is set to &#x60;true&#x60;, then any risk policies that include this predictor will require that the Signals SDK payload be provided as a signed JWT whose signature will be verified before proceeding with risk evaluation. You instruct the Signals SDK to provide the payload as a signed JWT by using the &#x60;universalDeviceIdentification&#x60; flag during initialization of the SDK, or by selecting the relevant setting for the &#x60;skrisk&#x60; component in DaVinci flows. | [optional] 
**ActivationAt** | Pointer to **time.Time** | You can use the &#x60;activationAt&#x60; parameter to specify a date on which the learning process for the predictor should be restarted. This can be used in conjunction with the fallback setting (&#x60;default.result.level&#x60;) to force strong authentication when moving the predictor to production. The date should be in an RFC3339 format. Note that activation date uses UTC time. | [optional] 

## Methods

### NewRiskPredictorDevice

`func NewRiskPredictorDevice(name string, compactName string, type_ EnumPredictorType, detect EnumPredictorNewDeviceDetectType, ) *RiskPredictorDevice`

NewRiskPredictorDevice instantiates a new RiskPredictorDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPredictorDeviceWithDefaults

`func NewRiskPredictorDeviceWithDefaults() *RiskPredictorDevice`

NewRiskPredictorDeviceWithDefaults instantiates a new RiskPredictorDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *RiskPredictorDevice) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RiskPredictorDevice) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RiskPredictorDevice) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RiskPredictorDevice) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *RiskPredictorDevice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskPredictorDevice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskPredictorDevice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskPredictorDevice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *RiskPredictorDevice) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *RiskPredictorDevice) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *RiskPredictorDevice) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *RiskPredictorDevice) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetName

`func (o *RiskPredictorDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskPredictorDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskPredictorDevice) SetName(v string)`

SetName sets Name field to given value.


### GetCompactName

`func (o *RiskPredictorDevice) GetCompactName() string`

GetCompactName returns the CompactName field if non-nil, zero value otherwise.

### GetCompactNameOk

`func (o *RiskPredictorDevice) GetCompactNameOk() (*string, bool)`

GetCompactNameOk returns a tuple with the CompactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactName

`func (o *RiskPredictorDevice) SetCompactName(v string)`

SetCompactName sets CompactName field to given value.


### GetType

`func (o *RiskPredictorDevice) GetType() EnumPredictorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskPredictorDevice) GetTypeOk() (*EnumPredictorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskPredictorDevice) SetType(v EnumPredictorType)`

SetType sets Type field to given value.


### GetDescription

`func (o *RiskPredictorDevice) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RiskPredictorDevice) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RiskPredictorDevice) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RiskPredictorDevice) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RiskPredictorDevice) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RiskPredictorDevice) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RiskPredictorDevice) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RiskPredictorDevice) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RiskPredictorDevice) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RiskPredictorDevice) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RiskPredictorDevice) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RiskPredictorDevice) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLicensed

`func (o *RiskPredictorDevice) GetLicensed() bool`

GetLicensed returns the Licensed field if non-nil, zero value otherwise.

### GetLicensedOk

`func (o *RiskPredictorDevice) GetLicensedOk() (*bool, bool)`

GetLicensedOk returns a tuple with the Licensed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensed

`func (o *RiskPredictorDevice) SetLicensed(v bool)`

SetLicensed sets Licensed field to given value.

### HasLicensed

`func (o *RiskPredictorDevice) HasLicensed() bool`

HasLicensed returns a boolean if a field has been set.

### GetDeletable

`func (o *RiskPredictorDevice) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *RiskPredictorDevice) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *RiskPredictorDevice) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *RiskPredictorDevice) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetDefault

`func (o *RiskPredictorDevice) GetDefault() RiskPredictorCommonDefault`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RiskPredictorDevice) GetDefaultOk() (*RiskPredictorCommonDefault, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RiskPredictorDevice) SetDefault(v RiskPredictorCommonDefault)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *RiskPredictorDevice) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetCondition

`func (o *RiskPredictorDevice) GetCondition() RiskPredictorCommonCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *RiskPredictorDevice) GetConditionOk() (*RiskPredictorCommonCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *RiskPredictorDevice) SetCondition(v RiskPredictorCommonCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *RiskPredictorDevice) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetDetect

`func (o *RiskPredictorDevice) GetDetect() EnumPredictorNewDeviceDetectType`

GetDetect returns the Detect field if non-nil, zero value otherwise.

### GetDetectOk

`func (o *RiskPredictorDevice) GetDetectOk() (*EnumPredictorNewDeviceDetectType, bool)`

GetDetectOk returns a tuple with the Detect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetect

`func (o *RiskPredictorDevice) SetDetect(v EnumPredictorNewDeviceDetectType)`

SetDetect sets Detect field to given value.


### GetShouldValidatePayloadSignature

`func (o *RiskPredictorDevice) GetShouldValidatePayloadSignature() bool`

GetShouldValidatePayloadSignature returns the ShouldValidatePayloadSignature field if non-nil, zero value otherwise.

### GetShouldValidatePayloadSignatureOk

`func (o *RiskPredictorDevice) GetShouldValidatePayloadSignatureOk() (*bool, bool)`

GetShouldValidatePayloadSignatureOk returns a tuple with the ShouldValidatePayloadSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldValidatePayloadSignature

`func (o *RiskPredictorDevice) SetShouldValidatePayloadSignature(v bool)`

SetShouldValidatePayloadSignature sets ShouldValidatePayloadSignature field to given value.

### HasShouldValidatePayloadSignature

`func (o *RiskPredictorDevice) HasShouldValidatePayloadSignature() bool`

HasShouldValidatePayloadSignature returns a boolean if a field has been set.

### GetActivationAt

`func (o *RiskPredictorDevice) GetActivationAt() time.Time`

GetActivationAt returns the ActivationAt field if non-nil, zero value otherwise.

### GetActivationAtOk

`func (o *RiskPredictorDevice) GetActivationAtOk() (*time.Time, bool)`

GetActivationAtOk returns a tuple with the ActivationAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationAt

`func (o *RiskPredictorDevice) SetActivationAt(v time.Time)`

SetActivationAt sets ActivationAt field to given value.

### HasActivationAt

`func (o *RiskPredictorDevice) HasActivationAt() bool`

HasActivationAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


