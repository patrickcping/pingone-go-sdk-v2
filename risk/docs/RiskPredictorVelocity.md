# RiskPredictorVelocity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**Name** | **string** | A string type. A unique, friendly name for the predictor. This name is displayed in the Risk Policies UI, when the admin is asked to define the overrides and weights. | 
**CompactName** | **string** | A string type. A unique name for the predictor. This property is immutable; it cannot be modified after initial creation. The value must be alpha-numeric, with no special characters or spaces. This name is used in the API both for policy configuration, and in the Risk Evaluation response (under details). | 
**Type** | [**EnumPredictorType**](EnumPredictorType.md) |  | 
**Description** | Pointer to **string** | A string type. This specifies the desription of the risk predictor. Maximum length is 1024 characters. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was updated. | [optional] [readonly] 
**Licensed** | Pointer to **bool** | Indicates whether PingOne Risk is licensed for the environment. | [optional] 
**Deletable** | Pointer to **bool** | A boolean to indicate whether the predictor is deletable in the environment. | [optional] 
**Default** | Pointer to [**RiskPredictorDefault**](RiskPredictorDefault.md) |  | [optional] 
**Condition** | Pointer to [**RiskPredictorCondition**](RiskPredictorCondition.md) |  | [optional] 
**By** | Pointer to **[]string** |  | [optional] 
**Every** | Pointer to [**RiskPredictorVelocityAllOfEvery**](RiskPredictorVelocityAllOfEvery.md) |  | [optional] 
**Fallback** | Pointer to [**RiskPredictorVelocityAllOfFallback**](RiskPredictorVelocityAllOfFallback.md) |  | [optional] 
**MaxDelay** | Pointer to [**RiskPredictorVelocityAllOfMaxDelay**](RiskPredictorVelocityAllOfMaxDelay.md) |  | [optional] 
**Measure** | Pointer to [**EnumPredictorVelocityMeasure**](EnumPredictorVelocityMeasure.md) |  | [optional] 
**Of** | Pointer to **string** |  | [optional] 
**SlidingWindow** | Pointer to [**RiskPredictorVelocityAllOfSlidingWindow**](RiskPredictorVelocityAllOfSlidingWindow.md) |  | [optional] 
**Use** | Pointer to [**RiskPredictorVelocityAllOfUse**](RiskPredictorVelocityAllOfUse.md) |  | [optional] 

## Methods

### NewRiskPredictorVelocity

`func NewRiskPredictorVelocity(name string, compactName string, type_ EnumPredictorType, ) *RiskPredictorVelocity`

NewRiskPredictorVelocity instantiates a new RiskPredictorVelocity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPredictorVelocityWithDefaults

`func NewRiskPredictorVelocityWithDefaults() *RiskPredictorVelocity`

NewRiskPredictorVelocityWithDefaults instantiates a new RiskPredictorVelocity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RiskPredictorVelocity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskPredictorVelocity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskPredictorVelocity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskPredictorVelocity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RiskPredictorVelocity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskPredictorVelocity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskPredictorVelocity) SetName(v string)`

SetName sets Name field to given value.


### GetCompactName

`func (o *RiskPredictorVelocity) GetCompactName() string`

GetCompactName returns the CompactName field if non-nil, zero value otherwise.

### GetCompactNameOk

`func (o *RiskPredictorVelocity) GetCompactNameOk() (*string, bool)`

GetCompactNameOk returns a tuple with the CompactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactName

`func (o *RiskPredictorVelocity) SetCompactName(v string)`

SetCompactName sets CompactName field to given value.


### GetType

`func (o *RiskPredictorVelocity) GetType() EnumPredictorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskPredictorVelocity) GetTypeOk() (*EnumPredictorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskPredictorVelocity) SetType(v EnumPredictorType)`

SetType sets Type field to given value.


### GetDescription

`func (o *RiskPredictorVelocity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RiskPredictorVelocity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RiskPredictorVelocity) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RiskPredictorVelocity) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RiskPredictorVelocity) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RiskPredictorVelocity) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RiskPredictorVelocity) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RiskPredictorVelocity) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RiskPredictorVelocity) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RiskPredictorVelocity) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RiskPredictorVelocity) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RiskPredictorVelocity) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLicensed

`func (o *RiskPredictorVelocity) GetLicensed() bool`

GetLicensed returns the Licensed field if non-nil, zero value otherwise.

### GetLicensedOk

`func (o *RiskPredictorVelocity) GetLicensedOk() (*bool, bool)`

GetLicensedOk returns a tuple with the Licensed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensed

`func (o *RiskPredictorVelocity) SetLicensed(v bool)`

SetLicensed sets Licensed field to given value.

### HasLicensed

`func (o *RiskPredictorVelocity) HasLicensed() bool`

HasLicensed returns a boolean if a field has been set.

### GetDeletable

`func (o *RiskPredictorVelocity) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *RiskPredictorVelocity) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *RiskPredictorVelocity) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *RiskPredictorVelocity) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetDefault

`func (o *RiskPredictorVelocity) GetDefault() RiskPredictorDefault`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RiskPredictorVelocity) GetDefaultOk() (*RiskPredictorDefault, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RiskPredictorVelocity) SetDefault(v RiskPredictorDefault)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *RiskPredictorVelocity) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetCondition

`func (o *RiskPredictorVelocity) GetCondition() RiskPredictorCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *RiskPredictorVelocity) GetConditionOk() (*RiskPredictorCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *RiskPredictorVelocity) SetCondition(v RiskPredictorCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *RiskPredictorVelocity) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetBy

`func (o *RiskPredictorVelocity) GetBy() []string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *RiskPredictorVelocity) GetByOk() (*[]string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *RiskPredictorVelocity) SetBy(v []string)`

SetBy sets By field to given value.

### HasBy

`func (o *RiskPredictorVelocity) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetEvery

`func (o *RiskPredictorVelocity) GetEvery() RiskPredictorVelocityAllOfEvery`

GetEvery returns the Every field if non-nil, zero value otherwise.

### GetEveryOk

`func (o *RiskPredictorVelocity) GetEveryOk() (*RiskPredictorVelocityAllOfEvery, bool)`

GetEveryOk returns a tuple with the Every field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvery

`func (o *RiskPredictorVelocity) SetEvery(v RiskPredictorVelocityAllOfEvery)`

SetEvery sets Every field to given value.

### HasEvery

`func (o *RiskPredictorVelocity) HasEvery() bool`

HasEvery returns a boolean if a field has been set.

### GetFallback

`func (o *RiskPredictorVelocity) GetFallback() RiskPredictorVelocityAllOfFallback`

GetFallback returns the Fallback field if non-nil, zero value otherwise.

### GetFallbackOk

`func (o *RiskPredictorVelocity) GetFallbackOk() (*RiskPredictorVelocityAllOfFallback, bool)`

GetFallbackOk returns a tuple with the Fallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallback

`func (o *RiskPredictorVelocity) SetFallback(v RiskPredictorVelocityAllOfFallback)`

SetFallback sets Fallback field to given value.

### HasFallback

`func (o *RiskPredictorVelocity) HasFallback() bool`

HasFallback returns a boolean if a field has been set.

### GetMaxDelay

`func (o *RiskPredictorVelocity) GetMaxDelay() RiskPredictorVelocityAllOfMaxDelay`

GetMaxDelay returns the MaxDelay field if non-nil, zero value otherwise.

### GetMaxDelayOk

`func (o *RiskPredictorVelocity) GetMaxDelayOk() (*RiskPredictorVelocityAllOfMaxDelay, bool)`

GetMaxDelayOk returns a tuple with the MaxDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDelay

`func (o *RiskPredictorVelocity) SetMaxDelay(v RiskPredictorVelocityAllOfMaxDelay)`

SetMaxDelay sets MaxDelay field to given value.

### HasMaxDelay

`func (o *RiskPredictorVelocity) HasMaxDelay() bool`

HasMaxDelay returns a boolean if a field has been set.

### GetMeasure

`func (o *RiskPredictorVelocity) GetMeasure() EnumPredictorVelocityMeasure`

GetMeasure returns the Measure field if non-nil, zero value otherwise.

### GetMeasureOk

`func (o *RiskPredictorVelocity) GetMeasureOk() (*EnumPredictorVelocityMeasure, bool)`

GetMeasureOk returns a tuple with the Measure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasure

`func (o *RiskPredictorVelocity) SetMeasure(v EnumPredictorVelocityMeasure)`

SetMeasure sets Measure field to given value.

### HasMeasure

`func (o *RiskPredictorVelocity) HasMeasure() bool`

HasMeasure returns a boolean if a field has been set.

### GetOf

`func (o *RiskPredictorVelocity) GetOf() string`

GetOf returns the Of field if non-nil, zero value otherwise.

### GetOfOk

`func (o *RiskPredictorVelocity) GetOfOk() (*string, bool)`

GetOfOk returns a tuple with the Of field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOf

`func (o *RiskPredictorVelocity) SetOf(v string)`

SetOf sets Of field to given value.

### HasOf

`func (o *RiskPredictorVelocity) HasOf() bool`

HasOf returns a boolean if a field has been set.

### GetSlidingWindow

`func (o *RiskPredictorVelocity) GetSlidingWindow() RiskPredictorVelocityAllOfSlidingWindow`

GetSlidingWindow returns the SlidingWindow field if non-nil, zero value otherwise.

### GetSlidingWindowOk

`func (o *RiskPredictorVelocity) GetSlidingWindowOk() (*RiskPredictorVelocityAllOfSlidingWindow, bool)`

GetSlidingWindowOk returns a tuple with the SlidingWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlidingWindow

`func (o *RiskPredictorVelocity) SetSlidingWindow(v RiskPredictorVelocityAllOfSlidingWindow)`

SetSlidingWindow sets SlidingWindow field to given value.

### HasSlidingWindow

`func (o *RiskPredictorVelocity) HasSlidingWindow() bool`

HasSlidingWindow returns a boolean if a field has been set.

### GetUse

`func (o *RiskPredictorVelocity) GetUse() RiskPredictorVelocityAllOfUse`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *RiskPredictorVelocity) GetUseOk() (*RiskPredictorVelocityAllOfUse, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *RiskPredictorVelocity) SetUse(v RiskPredictorVelocityAllOfUse)`

SetUse sets Use field to given value.

### HasUse

`func (o *RiskPredictorVelocity) HasUse() bool`

HasUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


