# SignOnPolicyActionCommonCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**[]SignOnPolicyActionCommonConditionOrOrInner**](SignOnPolicyActionCommonConditionOrOrInner.md) |  | [optional] 
**Or** | Pointer to [**[]SignOnPolicyActionCommonConditionOrOrInner**](SignOnPolicyActionCommonConditionOrOrInner.md) |  | [optional] 
**Not** | Pointer to [**[]SignOnPolicyActionCommonConditionOrOrInner**](SignOnPolicyActionCommonConditionOrOrInner.md) |  | [optional] 
**Greater** | **int32** | An integer that specifies the maximum number of seconds to wait since the last sign on before prompting for a new sign-on action. | 
**SecondsSince** | **string** | A string representing a condition variable.  For more information, see documenation at https://apidocs.pingidentity.com/pingone/platform/v1/api/#sign-on-policy-actions . | 
**IpRisk** | [**SignOnPolicyActionCommonConditionIPRiskIpRisk**](SignOnPolicyActionCommonConditionIPRiskIpRisk.md) |  | 
**Valid** | **string** |  | 
**Contains** | **string** |  | 
**IpRange** | **[]string** |  | 
**GeoVelocity** | **string** |  | 
**AnonymousNetwork** | **[]string** |  | 
**Value** | **string** |  | 
**Equals** | **string** |  | 

## Methods

### NewSignOnPolicyActionCommonCondition

`func NewSignOnPolicyActionCommonCondition(greater int32, secondsSince string, ipRisk SignOnPolicyActionCommonConditionIPRiskIpRisk, valid string, contains string, ipRange []string, geoVelocity string, anonymousNetwork []string, value string, equals string, ) *SignOnPolicyActionCommonCondition`

NewSignOnPolicyActionCommonCondition instantiates a new SignOnPolicyActionCommonCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionCommonConditionWithDefaults

`func NewSignOnPolicyActionCommonConditionWithDefaults() *SignOnPolicyActionCommonCondition`

NewSignOnPolicyActionCommonConditionWithDefaults instantiates a new SignOnPolicyActionCommonCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *SignOnPolicyActionCommonCondition) GetAnd() []SignOnPolicyActionCommonConditionOrOrInner`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *SignOnPolicyActionCommonCondition) GetAndOk() (*[]SignOnPolicyActionCommonConditionOrOrInner, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *SignOnPolicyActionCommonCondition) SetAnd(v []SignOnPolicyActionCommonConditionOrOrInner)`

SetAnd sets And field to given value.

### HasAnd

`func (o *SignOnPolicyActionCommonCondition) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetOr

`func (o *SignOnPolicyActionCommonCondition) GetOr() []SignOnPolicyActionCommonConditionOrOrInner`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *SignOnPolicyActionCommonCondition) GetOrOk() (*[]SignOnPolicyActionCommonConditionOrOrInner, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *SignOnPolicyActionCommonCondition) SetOr(v []SignOnPolicyActionCommonConditionOrOrInner)`

SetOr sets Or field to given value.

### HasOr

`func (o *SignOnPolicyActionCommonCondition) HasOr() bool`

HasOr returns a boolean if a field has been set.

### GetNot

`func (o *SignOnPolicyActionCommonCondition) GetNot() []SignOnPolicyActionCommonConditionOrOrInner`

GetNot returns the Not field if non-nil, zero value otherwise.

### GetNotOk

`func (o *SignOnPolicyActionCommonCondition) GetNotOk() (*[]SignOnPolicyActionCommonConditionOrOrInner, bool)`

GetNotOk returns a tuple with the Not field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNot

`func (o *SignOnPolicyActionCommonCondition) SetNot(v []SignOnPolicyActionCommonConditionOrOrInner)`

SetNot sets Not field to given value.

### HasNot

`func (o *SignOnPolicyActionCommonCondition) HasNot() bool`

HasNot returns a boolean if a field has been set.

### GetGreater

`func (o *SignOnPolicyActionCommonCondition) GetGreater() int32`

GetGreater returns the Greater field if non-nil, zero value otherwise.

### GetGreaterOk

`func (o *SignOnPolicyActionCommonCondition) GetGreaterOk() (*int32, bool)`

GetGreaterOk returns a tuple with the Greater field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreater

`func (o *SignOnPolicyActionCommonCondition) SetGreater(v int32)`

SetGreater sets Greater field to given value.


### GetSecondsSince

`func (o *SignOnPolicyActionCommonCondition) GetSecondsSince() string`

GetSecondsSince returns the SecondsSince field if non-nil, zero value otherwise.

### GetSecondsSinceOk

`func (o *SignOnPolicyActionCommonCondition) GetSecondsSinceOk() (*string, bool)`

GetSecondsSinceOk returns a tuple with the SecondsSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsSince

`func (o *SignOnPolicyActionCommonCondition) SetSecondsSince(v string)`

SetSecondsSince sets SecondsSince field to given value.


### GetIpRisk

`func (o *SignOnPolicyActionCommonCondition) GetIpRisk() SignOnPolicyActionCommonConditionIPRiskIpRisk`

GetIpRisk returns the IpRisk field if non-nil, zero value otherwise.

### GetIpRiskOk

`func (o *SignOnPolicyActionCommonCondition) GetIpRiskOk() (*SignOnPolicyActionCommonConditionIPRiskIpRisk, bool)`

GetIpRiskOk returns a tuple with the IpRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRisk

`func (o *SignOnPolicyActionCommonCondition) SetIpRisk(v SignOnPolicyActionCommonConditionIPRiskIpRisk)`

SetIpRisk sets IpRisk field to given value.


### GetValid

`func (o *SignOnPolicyActionCommonCondition) GetValid() string`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *SignOnPolicyActionCommonCondition) GetValidOk() (*string, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *SignOnPolicyActionCommonCondition) SetValid(v string)`

SetValid sets Valid field to given value.


### GetContains

`func (o *SignOnPolicyActionCommonCondition) GetContains() string`

GetContains returns the Contains field if non-nil, zero value otherwise.

### GetContainsOk

`func (o *SignOnPolicyActionCommonCondition) GetContainsOk() (*string, bool)`

GetContainsOk returns a tuple with the Contains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContains

`func (o *SignOnPolicyActionCommonCondition) SetContains(v string)`

SetContains sets Contains field to given value.


### GetIpRange

`func (o *SignOnPolicyActionCommonCondition) GetIpRange() []string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *SignOnPolicyActionCommonCondition) GetIpRangeOk() (*[]string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *SignOnPolicyActionCommonCondition) SetIpRange(v []string)`

SetIpRange sets IpRange field to given value.


### GetGeoVelocity

`func (o *SignOnPolicyActionCommonCondition) GetGeoVelocity() string`

GetGeoVelocity returns the GeoVelocity field if non-nil, zero value otherwise.

### GetGeoVelocityOk

`func (o *SignOnPolicyActionCommonCondition) GetGeoVelocityOk() (*string, bool)`

GetGeoVelocityOk returns a tuple with the GeoVelocity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoVelocity

`func (o *SignOnPolicyActionCommonCondition) SetGeoVelocity(v string)`

SetGeoVelocity sets GeoVelocity field to given value.


### GetAnonymousNetwork

`func (o *SignOnPolicyActionCommonCondition) GetAnonymousNetwork() []string`

GetAnonymousNetwork returns the AnonymousNetwork field if non-nil, zero value otherwise.

### GetAnonymousNetworkOk

`func (o *SignOnPolicyActionCommonCondition) GetAnonymousNetworkOk() (*[]string, bool)`

GetAnonymousNetworkOk returns a tuple with the AnonymousNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousNetwork

`func (o *SignOnPolicyActionCommonCondition) SetAnonymousNetwork(v []string)`

SetAnonymousNetwork sets AnonymousNetwork field to given value.


### GetValue

`func (o *SignOnPolicyActionCommonCondition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SignOnPolicyActionCommonCondition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SignOnPolicyActionCommonCondition) SetValue(v string)`

SetValue sets Value field to given value.


### GetEquals

`func (o *SignOnPolicyActionCommonCondition) GetEquals() string`

GetEquals returns the Equals field if non-nil, zero value otherwise.

### GetEqualsOk

`func (o *SignOnPolicyActionCommonCondition) GetEqualsOk() (*string, bool)`

GetEqualsOk returns a tuple with the Equals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquals

`func (o *SignOnPolicyActionCommonCondition) SetEquals(v string)`

SetEquals sets Equals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


