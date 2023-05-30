# SignOnPolicyActionLoginAllOfNewUserProvisioning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateways** | [**[]SignOnPolicyActionLoginAllOfNewUserProvisioningGateways**](SignOnPolicyActionLoginAllOfNewUserProvisioningGateways.md) | Allows a set of preconfigured gateways or &#x60;userType&#x60; pairs that are specified in the [Gateway Management](https://apidocs.pingidentity.com/pingone/platform/v1/api/#gateway-management) schema to determine how to find and migrate user entries existing in an external directory. | 

## Methods

### NewSignOnPolicyActionLoginAllOfNewUserProvisioning

`func NewSignOnPolicyActionLoginAllOfNewUserProvisioning(gateways []SignOnPolicyActionLoginAllOfNewUserProvisioningGateways, ) *SignOnPolicyActionLoginAllOfNewUserProvisioning`

NewSignOnPolicyActionLoginAllOfNewUserProvisioning instantiates a new SignOnPolicyActionLoginAllOfNewUserProvisioning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionLoginAllOfNewUserProvisioningWithDefaults

`func NewSignOnPolicyActionLoginAllOfNewUserProvisioningWithDefaults() *SignOnPolicyActionLoginAllOfNewUserProvisioning`

NewSignOnPolicyActionLoginAllOfNewUserProvisioningWithDefaults instantiates a new SignOnPolicyActionLoginAllOfNewUserProvisioning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateways

`func (o *SignOnPolicyActionLoginAllOfNewUserProvisioning) GetGateways() []SignOnPolicyActionLoginAllOfNewUserProvisioningGateways`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *SignOnPolicyActionLoginAllOfNewUserProvisioning) GetGatewaysOk() (*[]SignOnPolicyActionLoginAllOfNewUserProvisioningGateways, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *SignOnPolicyActionLoginAllOfNewUserProvisioning) SetGateways(v []SignOnPolicyActionLoginAllOfNewUserProvisioningGateways)`

SetGateways sets Gateways field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


