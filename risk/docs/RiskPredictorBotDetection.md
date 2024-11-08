# RiskPredictorBotDetection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
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
**IncludeRepeatedEventsWithoutSdk** | Pointer to **bool** | Set the value of &#x60;includeRepeatedEventsWithoutSdk&#x60; to &#x60;true&#x60; to expand the range of bot activity that PingOne Protect can detect. | [optional] 

## Methods

### NewRiskPredictorBotDetection

`func NewRiskPredictorBotDetection(name string, compactName string, type_ EnumPredictorType, ) *RiskPredictorBotDetection`

NewRiskPredictorBotDetection instantiates a new RiskPredictorBotDetection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPredictorBotDetectionWithDefaults

`func NewRiskPredictorBotDetectionWithDefaults() *RiskPredictorBotDetection`

NewRiskPredictorBotDetectionWithDefaults instantiates a new RiskPredictorBotDetection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *RiskPredictorBotDetection) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RiskPredictorBotDetection) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RiskPredictorBotDetection) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RiskPredictorBotDetection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *RiskPredictorBotDetection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskPredictorBotDetection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskPredictorBotDetection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskPredictorBotDetection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RiskPredictorBotDetection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskPredictorBotDetection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskPredictorBotDetection) SetName(v string)`

SetName sets Name field to given value.


### GetCompactName

`func (o *RiskPredictorBotDetection) GetCompactName() string`

GetCompactName returns the CompactName field if non-nil, zero value otherwise.

### GetCompactNameOk

`func (o *RiskPredictorBotDetection) GetCompactNameOk() (*string, bool)`

GetCompactNameOk returns a tuple with the CompactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactName

`func (o *RiskPredictorBotDetection) SetCompactName(v string)`

SetCompactName sets CompactName field to given value.


### GetType

`func (o *RiskPredictorBotDetection) GetType() EnumPredictorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskPredictorBotDetection) GetTypeOk() (*EnumPredictorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskPredictorBotDetection) SetType(v EnumPredictorType)`

SetType sets Type field to given value.


### GetDescription

`func (o *RiskPredictorBotDetection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RiskPredictorBotDetection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RiskPredictorBotDetection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RiskPredictorBotDetection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RiskPredictorBotDetection) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RiskPredictorBotDetection) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RiskPredictorBotDetection) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RiskPredictorBotDetection) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RiskPredictorBotDetection) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RiskPredictorBotDetection) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RiskPredictorBotDetection) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RiskPredictorBotDetection) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLicensed

`func (o *RiskPredictorBotDetection) GetLicensed() bool`

GetLicensed returns the Licensed field if non-nil, zero value otherwise.

### GetLicensedOk

`func (o *RiskPredictorBotDetection) GetLicensedOk() (*bool, bool)`

GetLicensedOk returns a tuple with the Licensed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensed

`func (o *RiskPredictorBotDetection) SetLicensed(v bool)`

SetLicensed sets Licensed field to given value.

### HasLicensed

`func (o *RiskPredictorBotDetection) HasLicensed() bool`

HasLicensed returns a boolean if a field has been set.

### GetDeletable

`func (o *RiskPredictorBotDetection) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *RiskPredictorBotDetection) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *RiskPredictorBotDetection) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *RiskPredictorBotDetection) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetDefault

`func (o *RiskPredictorBotDetection) GetDefault() RiskPredictorCommonDefault`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RiskPredictorBotDetection) GetDefaultOk() (*RiskPredictorCommonDefault, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RiskPredictorBotDetection) SetDefault(v RiskPredictorCommonDefault)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *RiskPredictorBotDetection) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetCondition

`func (o *RiskPredictorBotDetection) GetCondition() RiskPredictorCommonCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *RiskPredictorBotDetection) GetConditionOk() (*RiskPredictorCommonCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *RiskPredictorBotDetection) SetCondition(v RiskPredictorCommonCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *RiskPredictorBotDetection) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetIncludeRepeatedEventsWithoutSdk

`func (o *RiskPredictorBotDetection) GetIncludeRepeatedEventsWithoutSdk() bool`

GetIncludeRepeatedEventsWithoutSdk returns the IncludeRepeatedEventsWithoutSdk field if non-nil, zero value otherwise.

### GetIncludeRepeatedEventsWithoutSdkOk

`func (o *RiskPredictorBotDetection) GetIncludeRepeatedEventsWithoutSdkOk() (*bool, bool)`

GetIncludeRepeatedEventsWithoutSdkOk returns a tuple with the IncludeRepeatedEventsWithoutSdk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRepeatedEventsWithoutSdk

`func (o *RiskPredictorBotDetection) SetIncludeRepeatedEventsWithoutSdk(v bool)`

SetIncludeRepeatedEventsWithoutSdk sets IncludeRepeatedEventsWithoutSdk field to given value.

### HasIncludeRepeatedEventsWithoutSdk

`func (o *RiskPredictorBotDetection) HasIncludeRepeatedEventsWithoutSdk() bool`

HasIncludeRepeatedEventsWithoutSdk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


