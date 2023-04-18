# RiskPredictorAnonymous

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
**Licensed** | Pointer to **bool** | Indicates whether PingOne Risk is licensed for the environment. | [optional] [readonly] 
**Deletable** | Pointer to **bool** | A boolean to indicate whether the predictor is deletable in the environment. | [optional] [readonly] 
**Default** | Pointer to [**RiskPredictorDefault**](RiskPredictorDefault.md) |  | [optional] 
**Condition** | Pointer to [**RiskPredictorCondition**](RiskPredictorCondition.md) |  | [optional] 
**WhiteList** | Pointer to **[]string** | A list of IP addresses (CDIRs) that are ignored for the predictor results. | [optional] 

## Methods

### NewRiskPredictorAnonymous

`func NewRiskPredictorAnonymous(name string, compactName string, type_ EnumPredictorType, ) *RiskPredictorAnonymous`

NewRiskPredictorAnonymous instantiates a new RiskPredictorAnonymous object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPredictorAnonymousWithDefaults

`func NewRiskPredictorAnonymousWithDefaults() *RiskPredictorAnonymous`

NewRiskPredictorAnonymousWithDefaults instantiates a new RiskPredictorAnonymous object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RiskPredictorAnonymous) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskPredictorAnonymous) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskPredictorAnonymous) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskPredictorAnonymous) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RiskPredictorAnonymous) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskPredictorAnonymous) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskPredictorAnonymous) SetName(v string)`

SetName sets Name field to given value.


### GetCompactName

`func (o *RiskPredictorAnonymous) GetCompactName() string`

GetCompactName returns the CompactName field if non-nil, zero value otherwise.

### GetCompactNameOk

`func (o *RiskPredictorAnonymous) GetCompactNameOk() (*string, bool)`

GetCompactNameOk returns a tuple with the CompactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactName

`func (o *RiskPredictorAnonymous) SetCompactName(v string)`

SetCompactName sets CompactName field to given value.


### GetType

`func (o *RiskPredictorAnonymous) GetType() EnumPredictorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskPredictorAnonymous) GetTypeOk() (*EnumPredictorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskPredictorAnonymous) SetType(v EnumPredictorType)`

SetType sets Type field to given value.


### GetDescription

`func (o *RiskPredictorAnonymous) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RiskPredictorAnonymous) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RiskPredictorAnonymous) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RiskPredictorAnonymous) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RiskPredictorAnonymous) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RiskPredictorAnonymous) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RiskPredictorAnonymous) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RiskPredictorAnonymous) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RiskPredictorAnonymous) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RiskPredictorAnonymous) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RiskPredictorAnonymous) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RiskPredictorAnonymous) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLicensed

`func (o *RiskPredictorAnonymous) GetLicensed() bool`

GetLicensed returns the Licensed field if non-nil, zero value otherwise.

### GetLicensedOk

`func (o *RiskPredictorAnonymous) GetLicensedOk() (*bool, bool)`

GetLicensedOk returns a tuple with the Licensed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensed

`func (o *RiskPredictorAnonymous) SetLicensed(v bool)`

SetLicensed sets Licensed field to given value.

### HasLicensed

`func (o *RiskPredictorAnonymous) HasLicensed() bool`

HasLicensed returns a boolean if a field has been set.

### GetDeletable

`func (o *RiskPredictorAnonymous) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *RiskPredictorAnonymous) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *RiskPredictorAnonymous) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *RiskPredictorAnonymous) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetDefault

`func (o *RiskPredictorAnonymous) GetDefault() RiskPredictorDefault`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RiskPredictorAnonymous) GetDefaultOk() (*RiskPredictorDefault, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RiskPredictorAnonymous) SetDefault(v RiskPredictorDefault)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *RiskPredictorAnonymous) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetCondition

`func (o *RiskPredictorAnonymous) GetCondition() RiskPredictorCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *RiskPredictorAnonymous) GetConditionOk() (*RiskPredictorCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *RiskPredictorAnonymous) SetCondition(v RiskPredictorCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *RiskPredictorAnonymous) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetWhiteList

`func (o *RiskPredictorAnonymous) GetWhiteList() []string`

GetWhiteList returns the WhiteList field if non-nil, zero value otherwise.

### GetWhiteListOk

`func (o *RiskPredictorAnonymous) GetWhiteListOk() (*[]string, bool)`

GetWhiteListOk returns a tuple with the WhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteList

`func (o *RiskPredictorAnonymous) SetWhiteList(v []string)`

SetWhiteList sets WhiteList field to given value.

### HasWhiteList

`func (o *RiskPredictorAnonymous) HasWhiteList() bool`

HasWhiteList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


