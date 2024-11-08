# RiskPredictorTrafficAnomaly

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
**Rules** | [**[]RiskPredictorTrafficAnomalyAllOfRules**](RiskPredictorTrafficAnomalyAllOfRules.md) |  | 

## Methods

### NewRiskPredictorTrafficAnomaly

`func NewRiskPredictorTrafficAnomaly(name string, compactName string, type_ EnumPredictorType, rules []RiskPredictorTrafficAnomalyAllOfRules, ) *RiskPredictorTrafficAnomaly`

NewRiskPredictorTrafficAnomaly instantiates a new RiskPredictorTrafficAnomaly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPredictorTrafficAnomalyWithDefaults

`func NewRiskPredictorTrafficAnomalyWithDefaults() *RiskPredictorTrafficAnomaly`

NewRiskPredictorTrafficAnomalyWithDefaults instantiates a new RiskPredictorTrafficAnomaly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *RiskPredictorTrafficAnomaly) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RiskPredictorTrafficAnomaly) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RiskPredictorTrafficAnomaly) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RiskPredictorTrafficAnomaly) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *RiskPredictorTrafficAnomaly) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskPredictorTrafficAnomaly) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskPredictorTrafficAnomaly) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskPredictorTrafficAnomaly) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RiskPredictorTrafficAnomaly) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskPredictorTrafficAnomaly) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskPredictorTrafficAnomaly) SetName(v string)`

SetName sets Name field to given value.


### GetCompactName

`func (o *RiskPredictorTrafficAnomaly) GetCompactName() string`

GetCompactName returns the CompactName field if non-nil, zero value otherwise.

### GetCompactNameOk

`func (o *RiskPredictorTrafficAnomaly) GetCompactNameOk() (*string, bool)`

GetCompactNameOk returns a tuple with the CompactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactName

`func (o *RiskPredictorTrafficAnomaly) SetCompactName(v string)`

SetCompactName sets CompactName field to given value.


### GetType

`func (o *RiskPredictorTrafficAnomaly) GetType() EnumPredictorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskPredictorTrafficAnomaly) GetTypeOk() (*EnumPredictorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskPredictorTrafficAnomaly) SetType(v EnumPredictorType)`

SetType sets Type field to given value.


### GetDescription

`func (o *RiskPredictorTrafficAnomaly) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RiskPredictorTrafficAnomaly) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RiskPredictorTrafficAnomaly) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RiskPredictorTrafficAnomaly) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RiskPredictorTrafficAnomaly) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RiskPredictorTrafficAnomaly) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RiskPredictorTrafficAnomaly) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RiskPredictorTrafficAnomaly) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RiskPredictorTrafficAnomaly) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RiskPredictorTrafficAnomaly) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RiskPredictorTrafficAnomaly) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RiskPredictorTrafficAnomaly) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLicensed

`func (o *RiskPredictorTrafficAnomaly) GetLicensed() bool`

GetLicensed returns the Licensed field if non-nil, zero value otherwise.

### GetLicensedOk

`func (o *RiskPredictorTrafficAnomaly) GetLicensedOk() (*bool, bool)`

GetLicensedOk returns a tuple with the Licensed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensed

`func (o *RiskPredictorTrafficAnomaly) SetLicensed(v bool)`

SetLicensed sets Licensed field to given value.

### HasLicensed

`func (o *RiskPredictorTrafficAnomaly) HasLicensed() bool`

HasLicensed returns a boolean if a field has been set.

### GetDeletable

`func (o *RiskPredictorTrafficAnomaly) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *RiskPredictorTrafficAnomaly) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *RiskPredictorTrafficAnomaly) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *RiskPredictorTrafficAnomaly) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetDefault

`func (o *RiskPredictorTrafficAnomaly) GetDefault() RiskPredictorCommonDefault`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RiskPredictorTrafficAnomaly) GetDefaultOk() (*RiskPredictorCommonDefault, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RiskPredictorTrafficAnomaly) SetDefault(v RiskPredictorCommonDefault)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *RiskPredictorTrafficAnomaly) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetCondition

`func (o *RiskPredictorTrafficAnomaly) GetCondition() RiskPredictorCommonCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *RiskPredictorTrafficAnomaly) GetConditionOk() (*RiskPredictorCommonCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *RiskPredictorTrafficAnomaly) SetCondition(v RiskPredictorCommonCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *RiskPredictorTrafficAnomaly) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetRules

`func (o *RiskPredictorTrafficAnomaly) GetRules() []RiskPredictorTrafficAnomalyAllOfRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *RiskPredictorTrafficAnomaly) GetRulesOk() (*[]RiskPredictorTrafficAnomalyAllOfRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *RiskPredictorTrafficAnomaly) SetRules(v []RiskPredictorTrafficAnomalyAllOfRules)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


