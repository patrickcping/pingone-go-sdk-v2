# RiskPredictorTrafficAnomalyAllOfRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Set to true to use the defined rule in the predictor. | 
**Interval** | [**RiskPredictorTrafficAnomalyAllOfInterval**](RiskPredictorTrafficAnomalyAllOfInterval.md) |  | 
**Threshold** | [**RiskPredictorTrafficAnomalyAllOfThreshold**](RiskPredictorTrafficAnomalyAllOfThreshold.md) |  | 
**Type** | [**EnumRiskPredictorTrafficAnomalyRuleType**](EnumRiskPredictorTrafficAnomalyRuleType.md) |  | 

## Methods

### NewRiskPredictorTrafficAnomalyAllOfRules

`func NewRiskPredictorTrafficAnomalyAllOfRules(enabled bool, interval RiskPredictorTrafficAnomalyAllOfInterval, threshold RiskPredictorTrafficAnomalyAllOfThreshold, type_ EnumRiskPredictorTrafficAnomalyRuleType, ) *RiskPredictorTrafficAnomalyAllOfRules`

NewRiskPredictorTrafficAnomalyAllOfRules instantiates a new RiskPredictorTrafficAnomalyAllOfRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPredictorTrafficAnomalyAllOfRulesWithDefaults

`func NewRiskPredictorTrafficAnomalyAllOfRulesWithDefaults() *RiskPredictorTrafficAnomalyAllOfRules`

NewRiskPredictorTrafficAnomalyAllOfRulesWithDefaults instantiates a new RiskPredictorTrafficAnomalyAllOfRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *RiskPredictorTrafficAnomalyAllOfRules) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RiskPredictorTrafficAnomalyAllOfRules) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RiskPredictorTrafficAnomalyAllOfRules) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInterval

`func (o *RiskPredictorTrafficAnomalyAllOfRules) GetInterval() RiskPredictorTrafficAnomalyAllOfInterval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *RiskPredictorTrafficAnomalyAllOfRules) GetIntervalOk() (*RiskPredictorTrafficAnomalyAllOfInterval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *RiskPredictorTrafficAnomalyAllOfRules) SetInterval(v RiskPredictorTrafficAnomalyAllOfInterval)`

SetInterval sets Interval field to given value.


### GetThreshold

`func (o *RiskPredictorTrafficAnomalyAllOfRules) GetThreshold() RiskPredictorTrafficAnomalyAllOfThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *RiskPredictorTrafficAnomalyAllOfRules) GetThresholdOk() (*RiskPredictorTrafficAnomalyAllOfThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *RiskPredictorTrafficAnomalyAllOfRules) SetThreshold(v RiskPredictorTrafficAnomalyAllOfThreshold)`

SetThreshold sets Threshold field to given value.


### GetType

`func (o *RiskPredictorTrafficAnomalyAllOfRules) GetType() EnumRiskPredictorTrafficAnomalyRuleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskPredictorTrafficAnomalyAllOfRules) GetTypeOk() (*EnumRiskPredictorTrafficAnomalyRuleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskPredictorTrafficAnomalyAllOfRules) SetType(v EnumRiskPredictorTrafficAnomalyRuleType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


