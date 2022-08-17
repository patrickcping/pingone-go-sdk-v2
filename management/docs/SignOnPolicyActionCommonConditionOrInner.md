# SignOnPolicyActionCommonConditionOrInner

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
**Equals** | **string** |  | 

## Methods

### NewSignOnPolicyActionCommonConditionOrInner

`func NewSignOnPolicyActionCommonConditionOrInner(greater int32, secondsSince string, ipRisk SignOnPolicyActionCommonConditionIPRiskIpRisk, valid string, contains string, ipRange []string, geoVelocity string, anonymousNetwork []string, value string, equals string, ) *SignOnPolicyActionCommonConditionOrInner`

NewSignOnPolicyActionCommonConditionOrInner instantiates a new SignOnPolicyActionCommonConditionOrInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionCommonConditionOrInnerWithDefaults

`func NewSignOnPolicyActionCommonConditionOrInnerWithDefaults() *SignOnPolicyActionCommonConditionOrInner`

NewSignOnPolicyActionCommonConditionOrInnerWithDefaults instantiates a new SignOnPolicyActionCommonConditionOrInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGreater

`func (o *SignOnPolicyActionCommonConditionOrInner) GetGreater() int32`

GetGreater returns the Greater field if non-nil, zero value otherwise.

### GetGreaterOk

`func (o *SignOnPolicyActionCommonConditionOrInner) GetGreaterOk() (*int32, bool)`

GetGreaterOk returns a tuple with the Greater field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreater

`func (o *SignOnPolicyActionCommonConditionOrInner) SetGreater(v int32)`

SetGreater sets Greater field to given value.


### GetSecondsSince

`func (o *SignOnPolicyActionCommonConditionOrInner) GetSecondsSince() string`

GetSecondsSince returns the SecondsSince field if non-nil, zero value otherwise.

### GetSecondsSinceOk

`func (o *SignOnPolicyActionCommonConditionOrInner) GetSecondsSinceOk() (*string, bool)`

GetSecondsSinceOk returns a tuple with the SecondsSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsSince

`func (o *SignOnPolicyActionCommonConditionOrInner) SetSecondsSince(v string)`

SetSecondsSince sets SecondsSince field to given value.


### GetIpRisk

`func (o *SignOnPolicyActionCommonConditionOrInner) GetIpRisk() SignOnPolicyActionCommonConditionIPRiskIpRisk`

GetIpRisk returns the IpRisk field if non-nil, zero value otherwise.

### GetIpRiskOk

`func (o *SignOnPolicyActionCommonConditionOrInner) GetIpRiskOk() (*SignOnPolicyActionCommonConditionIPRiskIpRisk, bool)`

GetIpRiskOk returns a tuple with the IpRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRisk

`func (o *SignOnPolicyActionCommonConditionOrInner) SetIpRisk(v SignOnPolicyActionCommonConditionIPRiskIpRisk)`

SetIpRisk sets IpRisk field to given value.


### GetValid

`func (o *SignOnPolicyActionCommonConditionOrInner) GetValid() string`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *SignOnPolicyActionCommonConditionOrInner) GetValidOk() (*string, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *SignOnPolicyActionCommonConditionOrInner) SetValid(v string)`

SetValid sets Valid field to given value.


### GetContains

`func (o *SignOnPolicyActionCommonConditionOrInner) GetContains() string`

GetContains returns the Contains field if non-nil, zero value otherwise.

### GetContainsOk

`func (o *SignOnPolicyActionCommonConditionOrInner) GetContainsOk() (*string, bool)`

GetContainsOk returns a tuple with the Contains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContains

`func (o *SignOnPolicyActionCommonConditionOrInner) SetContains(v string)`

SetContains sets Contains field to given value.


### GetIpRange

`func (o *SignOnPolicyActionCommonConditionOrInner) GetIpRange() []string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *SignOnPolicyActionCommonConditionOrInner) GetIpRangeOk() (*[]string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *SignOnPolicyActionCommonConditionOrInner) SetIpRange(v []string)`

SetIpRange sets IpRange field to given value.


### GetGeoVelocity

`func (o *SignOnPolicyActionCommonConditionOrInner) GetGeoVelocity() string`

GetGeoVelocity returns the GeoVelocity field if non-nil, zero value otherwise.

### GetGeoVelocityOk

`func (o *SignOnPolicyActionCommonConditionOrInner) GetGeoVelocityOk() (*string, bool)`

GetGeoVelocityOk returns a tuple with the GeoVelocity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoVelocity

`func (o *SignOnPolicyActionCommonConditionOrInner) SetGeoVelocity(v string)`

SetGeoVelocity sets GeoVelocity field to given value.


### GetAnonymousNetwork

`func (o *SignOnPolicyActionCommonConditionOrInner) GetAnonymousNetwork() []string`

GetAnonymousNetwork returns the AnonymousNetwork field if non-nil, zero value otherwise.

### GetAnonymousNetworkOk

`func (o *SignOnPolicyActionCommonConditionOrInner) GetAnonymousNetworkOk() (*[]string, bool)`

GetAnonymousNetworkOk returns a tuple with the AnonymousNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousNetwork

`func (o *SignOnPolicyActionCommonConditionOrInner) SetAnonymousNetwork(v []string)`

SetAnonymousNetwork sets AnonymousNetwork field to given value.


### GetValue

`func (o *SignOnPolicyActionCommonConditionOrInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SignOnPolicyActionCommonConditionOrInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SignOnPolicyActionCommonConditionOrInner) SetValue(v string)`

SetValue sets Value field to given value.


### GetEquals

`func (o *SignOnPolicyActionCommonConditionOrInner) GetEquals() string`

GetEquals returns the Equals field if non-nil, zero value otherwise.

### GetEqualsOk

`func (o *SignOnPolicyActionCommonConditionOrInner) GetEqualsOk() (*string, bool)`

GetEqualsOk returns a tuple with the Equals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquals

`func (o *SignOnPolicyActionCommonConditionOrInner) SetEquals(v string)`

SetEquals sets Equals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


