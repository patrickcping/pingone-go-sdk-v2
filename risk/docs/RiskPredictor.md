# RiskPredictor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**Name** | **string** | A string type. A unique, friendly name for the predictor. This name is displayed in the Risk Policies UI, when the admin is asked to define the overrides and weights. | 
**CompactName** | **string** | A string type. A unique name for the predictor. This property is immutable; it cannot be modified after initial creation. The value must be alpha-numeric, with no special characters or spaces. This name is used in the API both for policy configuration, and in the Risk Evaluation response (under details). | 
**Type** | [**EnumPredictorType**](EnumPredictorType.md) |  | 
**Description** | Pointer to **string** | A string type. This specifies the desription of the risk predictor. Maximum length is 1024 characters. | [optional] 
**CreatedAt** | Pointer to **string** | The time the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | The time the resource was updated. | [optional] [readonly] 
**Default** | Pointer to [**RiskPredictorDefault**](RiskPredictorDefault.md) |  | [optional] 
**Map** | Pointer to [**RiskPredictorMap**](RiskPredictorMap.md) |  | [optional] 

## Methods

### NewRiskPredictor

`func NewRiskPredictor(name string, compactName string, type_ EnumPredictorType, ) *RiskPredictor`

NewRiskPredictor instantiates a new RiskPredictor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPredictorWithDefaults

`func NewRiskPredictorWithDefaults() *RiskPredictor`

NewRiskPredictorWithDefaults instantiates a new RiskPredictor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RiskPredictor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskPredictor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskPredictor) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskPredictor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RiskPredictor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskPredictor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskPredictor) SetName(v string)`

SetName sets Name field to given value.


### GetCompactName

`func (o *RiskPredictor) GetCompactName() string`

GetCompactName returns the CompactName field if non-nil, zero value otherwise.

### GetCompactNameOk

`func (o *RiskPredictor) GetCompactNameOk() (*string, bool)`

GetCompactNameOk returns a tuple with the CompactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactName

`func (o *RiskPredictor) SetCompactName(v string)`

SetCompactName sets CompactName field to given value.


### GetType

`func (o *RiskPredictor) GetType() EnumPredictorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskPredictor) GetTypeOk() (*EnumPredictorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskPredictor) SetType(v EnumPredictorType)`

SetType sets Type field to given value.


### GetDescription

`func (o *RiskPredictor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RiskPredictor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RiskPredictor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RiskPredictor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RiskPredictor) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RiskPredictor) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RiskPredictor) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RiskPredictor) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RiskPredictor) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RiskPredictor) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RiskPredictor) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RiskPredictor) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDefault

`func (o *RiskPredictor) GetDefault() RiskPredictorDefault`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RiskPredictor) GetDefaultOk() (*RiskPredictorDefault, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RiskPredictor) SetDefault(v RiskPredictorDefault)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *RiskPredictor) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetMap

`func (o *RiskPredictor) GetMap() RiskPredictorMap`

GetMap returns the Map field if non-nil, zero value otherwise.

### GetMapOk

`func (o *RiskPredictor) GetMapOk() (*RiskPredictorMap, bool)`

GetMapOk returns a tuple with the Map field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMap

`func (o *RiskPredictor) SetMap(v RiskPredictorMap)`

SetMap sets Map field to given value.

### HasMap

`func (o *RiskPredictor) HasMap() bool`

HasMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


