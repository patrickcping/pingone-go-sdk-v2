# SignOnPolicyActionCommonConditionOrOrInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**[]SignOnPolicyActionCommonConditionOrOrInner**](SignOnPolicyActionCommonConditionOrOrInner.md) |  | [optional] 
**Or** | Pointer to [**[]SignOnPolicyActionCommonConditionOrOrInner**](SignOnPolicyActionCommonConditionOrOrInner.md) |  | [optional] 
**Not** | Pointer to [**SignOnPolicyActionCommonConditionAggregate**](SignOnPolicyActionCommonConditionAggregate.md) |  | [optional] 
**Greater** | **int32** | An integer that specifies the maximum number of seconds to wait since the last sign on before prompting for a new sign-on action. | 
**SecondsSince** | **string** | A string representing a condition variable.  For more information, see documenation at https://apidocs.pingidentity.com/pingone/platform/v1/api/#sign-on-policy-actions . | 
**IpRisk** | [**SignOnPolicyActionCommonConditionIPRiskIpRisk**](SignOnPolicyActionCommonConditionIPRiskIpRisk.md) |  | 
**Valid** | **string** |  | 
**Contains** | **string** |  | 
**IpRange** | **[]string** |  | 
**GeoVelocity** | **string** |  | 
**AnonymousNetwork** | **[]string** |  | 
**Value** | **string** |  | 
**Equals** | [**SignOnPolicyActionCommonConditionEqualsEquals**](SignOnPolicyActionCommonConditionEqualsEquals.md) |  | 

## Methods

### NewSignOnPolicyActionCommonConditionOrOrInner

`func NewSignOnPolicyActionCommonConditionOrOrInner(greater int32, secondsSince string, ipRisk SignOnPolicyActionCommonConditionIPRiskIpRisk, valid string, contains string, ipRange []string, geoVelocity string, anonymousNetwork []string, value string, equals SignOnPolicyActionCommonConditionEqualsEquals, ) *SignOnPolicyActionCommonConditionOrOrInner`

NewSignOnPolicyActionCommonConditionOrOrInner instantiates a new SignOnPolicyActionCommonConditionOrOrInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionCommonConditionOrOrInnerWithDefaults

`func NewSignOnPolicyActionCommonConditionOrOrInnerWithDefaults() *SignOnPolicyActionCommonConditionOrOrInner`

NewSignOnPolicyActionCommonConditionOrOrInnerWithDefaults instantiates a new SignOnPolicyActionCommonConditionOrOrInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *SignOnPolicyActionCommonConditionOrOrInner) GetAnd() []SignOnPolicyActionCommonConditionOrOrInner`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *SignOnPolicyActionCommonConditionOrOrInner) GetAndOk() (*[]SignOnPolicyActionCommonConditionOrOrInner, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *SignOnPolicyActionCommonConditionOrOrInner) SetAnd(v []SignOnPolicyActionCommonConditionOrOrInner)`

SetAnd sets And field to given value.

### HasAnd

`func (o *SignOnPolicyActionCommonConditionOrOrInner) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetOr

`func (o *SignOnPolicyActionCommonConditionOrOrInner) GetOr() []SignOnPolicyActionCommonConditionOrOrInner`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *SignOnPolicyActionCommonConditionOrOrInner) GetOrOk() (*[]SignOnPolicyActionCommonConditionOrOrInner, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *SignOnPolicyActionCommonConditionOrOrInner) SetOr(v []SignOnPolicyActionCommonConditionOrOrInner)`

SetOr sets Or field to given value.

### HasOr

`func (o *SignOnPolicyActionCommonConditionOrOrInner) HasOr() bool`

HasOr returns a boolean if a field has been set.

### GetNot

`func (o *SignOnPolicyActionCommonConditionOrOrInner) GetNot() SignOnPolicyActionCommonConditionAggregate`

GetNot returns the Not field if non-nil, zero value otherwise.

### GetNotOk

`func (o *SignOnPolicyActionCommonConditionOrOrInner) GetNotOk() (*SignOnPolicyActionCommonConditionAggregate, bool)`

GetNotOk returns a tuple with the Not field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNot

`func (o *SignOnPolicyActionCommonConditionOrOrInner) SetNot(v SignOnPolicyActionCommonConditionAggregate)`

SetNot sets Not field to given value.

### HasNot

`func (o *SignOnPolicyActionCommonConditionOrOrInner) HasNot() bool`

HasNot returns a boolean if a field has been set.

### GetGreater

`func (o *SignOnPolicyActionCommonConditionOrOrInner) GetGreater() int32`

GetGreater returns the Greater field if non-nil, zero value otherwise.

### GetGreaterOk

`func (o *SignOnPolicyActionCommonConditionOrOrInner) GetGreaterOk() (*int32, bool)`

GetGreaterOk returns a tuple with the Greater field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreater

`func (o *SignOnPolicyActionCommonConditionOrOrInner) SetGreater(v int32)`

SetGreater sets Greater field to given value.


### GetSecondsSince

`func (o *SignOnPolicyActionCommonConditionOrOrInner) GetSecondsSince() string`

GetSecondsSince returns the SecondsSince field if non-nil, zero value otherwise.

### GetSecondsSinceOk

`func (o *SignOnPolicyActionCommonConditionOrOrInner) GetSecondsSinceOk() (*string, bool)`

GetSecondsSinceOk returns a tuple with the SecondsSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsSince

`func (o *SignOnPolicyActionCommonConditionOrOrInner) SetSecondsSince(v string)`

SetSecondsSince sets SecondsSince field to given value.


### GetIpRisk

`func (o *SignOnPolicyActionCommonConditionOrOrInner) GetIpRisk() SignOnPolicyActionCommonConditionIPRiskIpRisk`

GetIpRisk returns the IpRisk field if non-nil, zero value otherwise.

### GetIpRiskOk

`func (o *SignOnPolicyActionCommonConditionOrOrInner) GetIpRiskOk() (*SignOnPolicyActionCommonConditionIPRiskIpRisk, bool)`

GetIpRiskOk returns a tuple with the IpRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRisk

`func (o *SignOnPolicyActionCommonConditionOrOrInner) SetIpRisk(v SignOnPolicyActionCommonConditionIPRiskIpRisk)`

SetIpRisk sets IpRisk field to given value.


### GetValid

`func (o *SignOnPolicyActionCommonConditionOrOrInner) GetValid() string`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *SignOnPolicyActionCommonConditionOrOrInner) GetValidOk() (*string, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *SignOnPolicyActionCommonConditionOrOrInner) SetValid(v string)`

SetValid sets Valid field to given value.


### GetContains

`func (o *SignOnPolicyActionCommonConditionOrOrInner) GetContains() string`

GetContains returns the Contains field if non-nil, zero value otherwise.

### GetContainsOk

`func (o *SignOnPolicyActionCommonConditionOrOrInner) GetContainsOk() (*string, bool)`

GetContainsOk returns a tuple with the Contains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContains

`func (o *SignOnPolicyActionCommonConditionOrOrInner) SetContains(v string)`

SetContains sets Contains field to given value.


### GetIpRange

`func (o *SignOnPolicyActionCommonConditionOrOrInner) GetIpRange() []string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *SignOnPolicyActionCommonConditionOrOrInner) GetIpRangeOk() (*[]string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *SignOnPolicyActionCommonConditionOrOrInner) SetIpRange(v []string)`

SetIpRange sets IpRange field to given value.


### GetGeoVelocity

`func (o *SignOnPolicyActionCommonConditionOrOrInner) GetGeoVelocity() string`

GetGeoVelocity returns the GeoVelocity field if non-nil, zero value otherwise.

### GetGeoVelocityOk

`func (o *SignOnPolicyActionCommonConditionOrOrInner) GetGeoVelocityOk() (*string, bool)`

GetGeoVelocityOk returns a tuple with the GeoVelocity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoVelocity

`func (o *SignOnPolicyActionCommonConditionOrOrInner) SetGeoVelocity(v string)`

SetGeoVelocity sets GeoVelocity field to given value.


### GetAnonymousNetwork

`func (o *SignOnPolicyActionCommonConditionOrOrInner) GetAnonymousNetwork() []string`

GetAnonymousNetwork returns the AnonymousNetwork field if non-nil, zero value otherwise.

### GetAnonymousNetworkOk

`func (o *SignOnPolicyActionCommonConditionOrOrInner) GetAnonymousNetworkOk() (*[]string, bool)`

GetAnonymousNetworkOk returns a tuple with the AnonymousNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousNetwork

`func (o *SignOnPolicyActionCommonConditionOrOrInner) SetAnonymousNetwork(v []string)`

SetAnonymousNetwork sets AnonymousNetwork field to given value.


### GetValue

`func (o *SignOnPolicyActionCommonConditionOrOrInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SignOnPolicyActionCommonConditionOrOrInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SignOnPolicyActionCommonConditionOrOrInner) SetValue(v string)`

SetValue sets Value field to given value.


### GetEquals

`func (o *SignOnPolicyActionCommonConditionOrOrInner) GetEquals() SignOnPolicyActionCommonConditionEqualsEquals`

GetEquals returns the Equals field if non-nil, zero value otherwise.

### GetEqualsOk

`func (o *SignOnPolicyActionCommonConditionOrOrInner) GetEqualsOk() (*SignOnPolicyActionCommonConditionEqualsEquals, bool)`

GetEqualsOk returns a tuple with the Equals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquals

`func (o *SignOnPolicyActionCommonConditionOrOrInner) SetEquals(v SignOnPolicyActionCommonConditionEqualsEquals)`

SetEquals sets Equals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


