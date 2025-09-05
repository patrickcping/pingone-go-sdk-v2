# RiskPredictorComposite

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
**Compositions** | [**[]RiskPredictorCompositeAllOfCompositionsInner**](RiskPredictorCompositeAllOfCompositionsInner.md) | Contains the objects that specify the conditions to test and the risk level that should be assigned if the conditions are met. The array can contain a maximum of three elements. | 

## Methods

### NewRiskPredictorComposite

`func NewRiskPredictorComposite(name string, compactName string, type_ EnumPredictorType, compositions []RiskPredictorCompositeAllOfCompositionsInner, ) *RiskPredictorComposite`

NewRiskPredictorComposite instantiates a new RiskPredictorComposite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPredictorCompositeWithDefaults

`func NewRiskPredictorCompositeWithDefaults() *RiskPredictorComposite`

NewRiskPredictorCompositeWithDefaults instantiates a new RiskPredictorComposite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *RiskPredictorComposite) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RiskPredictorComposite) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RiskPredictorComposite) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RiskPredictorComposite) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *RiskPredictorComposite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskPredictorComposite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskPredictorComposite) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskPredictorComposite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *RiskPredictorComposite) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *RiskPredictorComposite) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *RiskPredictorComposite) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *RiskPredictorComposite) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetName

`func (o *RiskPredictorComposite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskPredictorComposite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskPredictorComposite) SetName(v string)`

SetName sets Name field to given value.


### GetCompactName

`func (o *RiskPredictorComposite) GetCompactName() string`

GetCompactName returns the CompactName field if non-nil, zero value otherwise.

### GetCompactNameOk

`func (o *RiskPredictorComposite) GetCompactNameOk() (*string, bool)`

GetCompactNameOk returns a tuple with the CompactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactName

`func (o *RiskPredictorComposite) SetCompactName(v string)`

SetCompactName sets CompactName field to given value.


### GetType

`func (o *RiskPredictorComposite) GetType() EnumPredictorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskPredictorComposite) GetTypeOk() (*EnumPredictorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskPredictorComposite) SetType(v EnumPredictorType)`

SetType sets Type field to given value.


### GetDescription

`func (o *RiskPredictorComposite) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RiskPredictorComposite) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RiskPredictorComposite) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RiskPredictorComposite) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RiskPredictorComposite) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RiskPredictorComposite) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RiskPredictorComposite) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RiskPredictorComposite) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RiskPredictorComposite) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RiskPredictorComposite) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RiskPredictorComposite) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RiskPredictorComposite) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLicensed

`func (o *RiskPredictorComposite) GetLicensed() bool`

GetLicensed returns the Licensed field if non-nil, zero value otherwise.

### GetLicensedOk

`func (o *RiskPredictorComposite) GetLicensedOk() (*bool, bool)`

GetLicensedOk returns a tuple with the Licensed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensed

`func (o *RiskPredictorComposite) SetLicensed(v bool)`

SetLicensed sets Licensed field to given value.

### HasLicensed

`func (o *RiskPredictorComposite) HasLicensed() bool`

HasLicensed returns a boolean if a field has been set.

### GetDeletable

`func (o *RiskPredictorComposite) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *RiskPredictorComposite) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *RiskPredictorComposite) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *RiskPredictorComposite) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetDefault

`func (o *RiskPredictorComposite) GetDefault() RiskPredictorCommonDefault`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RiskPredictorComposite) GetDefaultOk() (*RiskPredictorCommonDefault, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RiskPredictorComposite) SetDefault(v RiskPredictorCommonDefault)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *RiskPredictorComposite) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetCondition

`func (o *RiskPredictorComposite) GetCondition() RiskPredictorCommonCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *RiskPredictorComposite) GetConditionOk() (*RiskPredictorCommonCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *RiskPredictorComposite) SetCondition(v RiskPredictorCommonCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *RiskPredictorComposite) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetCompositions

`func (o *RiskPredictorComposite) GetCompositions() []RiskPredictorCompositeAllOfCompositionsInner`

GetCompositions returns the Compositions field if non-nil, zero value otherwise.

### GetCompositionsOk

`func (o *RiskPredictorComposite) GetCompositionsOk() (*[]RiskPredictorCompositeAllOfCompositionsInner, bool)`

GetCompositionsOk returns a tuple with the Compositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositions

`func (o *RiskPredictorComposite) SetCompositions(v []RiskPredictorCompositeAllOfCompositionsInner)`

SetCompositions sets Compositions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


