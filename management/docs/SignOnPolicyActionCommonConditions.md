# SignOnPolicyActionCommonConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpRange** | Pointer to **string** | A string that specifies the supported network IP addresses expressed as classless inter-domain routing (CIDR) strings. | [optional] 
**SecondsSince** | Pointer to **string** | An integer that specifies the maximum number of minutes to wait since the last sign on before prompting for a new sign-on action. | [optional] 

## Methods

### NewSignOnPolicyActionCommonConditions

`func NewSignOnPolicyActionCommonConditions() *SignOnPolicyActionCommonConditions`

NewSignOnPolicyActionCommonConditions instantiates a new SignOnPolicyActionCommonConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionCommonConditionsWithDefaults

`func NewSignOnPolicyActionCommonConditionsWithDefaults() *SignOnPolicyActionCommonConditions`

NewSignOnPolicyActionCommonConditionsWithDefaults instantiates a new SignOnPolicyActionCommonConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpRange

`func (o *SignOnPolicyActionCommonConditions) GetIpRange() string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *SignOnPolicyActionCommonConditions) GetIpRangeOk() (*string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *SignOnPolicyActionCommonConditions) SetIpRange(v string)`

SetIpRange sets IpRange field to given value.

### HasIpRange

`func (o *SignOnPolicyActionCommonConditions) HasIpRange() bool`

HasIpRange returns a boolean if a field has been set.

### GetSecondsSince

`func (o *SignOnPolicyActionCommonConditions) GetSecondsSince() string`

GetSecondsSince returns the SecondsSince field if non-nil, zero value otherwise.

### GetSecondsSinceOk

`func (o *SignOnPolicyActionCommonConditions) GetSecondsSinceOk() (*string, bool)`

GetSecondsSinceOk returns a tuple with the SecondsSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsSince

`func (o *SignOnPolicyActionCommonConditions) SetSecondsSince(v string)`

SetSecondsSince sets SecondsSince field to given value.

### HasSecondsSince

`func (o *SignOnPolicyActionCommonConditions) HasSecondsSince() bool`

HasSecondsSince returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


