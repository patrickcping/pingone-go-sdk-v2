# SignOnPolicyActionCommonConditionAggregate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewSignOnPolicyActionCommonConditionAggregate

`func NewSignOnPolicyActionCommonConditionAggregate(greater int32, secondsSince string, ipRisk SignOnPolicyActionCommonConditionIPRiskIpRisk, valid string, contains string, ipRange []string, geoVelocity string, anonymousNetwork []string, value string, equals SignOnPolicyActionCommonConditionEqualsEquals, ) *SignOnPolicyActionCommonConditionAggregate`

NewSignOnPolicyActionCommonConditionAggregate instantiates a new SignOnPolicyActionCommonConditionAggregate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionCommonConditionAggregateWithDefaults

`func NewSignOnPolicyActionCommonConditionAggregateWithDefaults() *SignOnPolicyActionCommonConditionAggregate`

NewSignOnPolicyActionCommonConditionAggregateWithDefaults instantiates a new SignOnPolicyActionCommonConditionAggregate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGreater

`func (o *SignOnPolicyActionCommonConditionAggregate) GetGreater() int32`

GetGreater returns the Greater field if non-nil, zero value otherwise.

### GetGreaterOk

`func (o *SignOnPolicyActionCommonConditionAggregate) GetGreaterOk() (*int32, bool)`

GetGreaterOk returns a tuple with the Greater field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreater

`func (o *SignOnPolicyActionCommonConditionAggregate) SetGreater(v int32)`

SetGreater sets Greater field to given value.


### GetSecondsSince

`func (o *SignOnPolicyActionCommonConditionAggregate) GetSecondsSince() string`

GetSecondsSince returns the SecondsSince field if non-nil, zero value otherwise.

### GetSecondsSinceOk

`func (o *SignOnPolicyActionCommonConditionAggregate) GetSecondsSinceOk() (*string, bool)`

GetSecondsSinceOk returns a tuple with the SecondsSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsSince

`func (o *SignOnPolicyActionCommonConditionAggregate) SetSecondsSince(v string)`

SetSecondsSince sets SecondsSince field to given value.


### GetIpRisk

`func (o *SignOnPolicyActionCommonConditionAggregate) GetIpRisk() SignOnPolicyActionCommonConditionIPRiskIpRisk`

GetIpRisk returns the IpRisk field if non-nil, zero value otherwise.

### GetIpRiskOk

`func (o *SignOnPolicyActionCommonConditionAggregate) GetIpRiskOk() (*SignOnPolicyActionCommonConditionIPRiskIpRisk, bool)`

GetIpRiskOk returns a tuple with the IpRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRisk

`func (o *SignOnPolicyActionCommonConditionAggregate) SetIpRisk(v SignOnPolicyActionCommonConditionIPRiskIpRisk)`

SetIpRisk sets IpRisk field to given value.


### GetValid

`func (o *SignOnPolicyActionCommonConditionAggregate) GetValid() string`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *SignOnPolicyActionCommonConditionAggregate) GetValidOk() (*string, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *SignOnPolicyActionCommonConditionAggregate) SetValid(v string)`

SetValid sets Valid field to given value.


### GetContains

`func (o *SignOnPolicyActionCommonConditionAggregate) GetContains() string`

GetContains returns the Contains field if non-nil, zero value otherwise.

### GetContainsOk

`func (o *SignOnPolicyActionCommonConditionAggregate) GetContainsOk() (*string, bool)`

GetContainsOk returns a tuple with the Contains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContains

`func (o *SignOnPolicyActionCommonConditionAggregate) SetContains(v string)`

SetContains sets Contains field to given value.


### GetIpRange

`func (o *SignOnPolicyActionCommonConditionAggregate) GetIpRange() []string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *SignOnPolicyActionCommonConditionAggregate) GetIpRangeOk() (*[]string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *SignOnPolicyActionCommonConditionAggregate) SetIpRange(v []string)`

SetIpRange sets IpRange field to given value.


### GetGeoVelocity

`func (o *SignOnPolicyActionCommonConditionAggregate) GetGeoVelocity() string`

GetGeoVelocity returns the GeoVelocity field if non-nil, zero value otherwise.

### GetGeoVelocityOk

`func (o *SignOnPolicyActionCommonConditionAggregate) GetGeoVelocityOk() (*string, bool)`

GetGeoVelocityOk returns a tuple with the GeoVelocity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoVelocity

`func (o *SignOnPolicyActionCommonConditionAggregate) SetGeoVelocity(v string)`

SetGeoVelocity sets GeoVelocity field to given value.


### GetAnonymousNetwork

`func (o *SignOnPolicyActionCommonConditionAggregate) GetAnonymousNetwork() []string`

GetAnonymousNetwork returns the AnonymousNetwork field if non-nil, zero value otherwise.

### GetAnonymousNetworkOk

`func (o *SignOnPolicyActionCommonConditionAggregate) GetAnonymousNetworkOk() (*[]string, bool)`

GetAnonymousNetworkOk returns a tuple with the AnonymousNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousNetwork

`func (o *SignOnPolicyActionCommonConditionAggregate) SetAnonymousNetwork(v []string)`

SetAnonymousNetwork sets AnonymousNetwork field to given value.


### GetValue

`func (o *SignOnPolicyActionCommonConditionAggregate) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SignOnPolicyActionCommonConditionAggregate) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SignOnPolicyActionCommonConditionAggregate) SetValue(v string)`

SetValue sets Value field to given value.


### GetEquals

`func (o *SignOnPolicyActionCommonConditionAggregate) GetEquals() SignOnPolicyActionCommonConditionEqualsEquals`

GetEquals returns the Equals field if non-nil, zero value otherwise.

### GetEqualsOk

`func (o *SignOnPolicyActionCommonConditionAggregate) GetEqualsOk() (*SignOnPolicyActionCommonConditionEqualsEquals, bool)`

GetEqualsOk returns a tuple with the Equals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquals

`func (o *SignOnPolicyActionCommonConditionAggregate) SetEquals(v SignOnPolicyActionCommonConditionEqualsEquals)`

SetEquals sets Equals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


