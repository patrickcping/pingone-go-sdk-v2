# DeviceAuthenticationPolicyPingIDAllOfMobileIpPairingConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnyIPAdress** | Pointer to **bool** | If you want to limit users to specific IP addresses when pairing their device, set &#x60;anyIPAdress&#x60; to false, and use &#x60;ipPairingConfiguration.onlyTheseIpAddresses&#x60; to specify the valid IP addresses. | [optional] 
**OnlyTheseIpAddresses** | Pointer to **[]string** | If you set &#x60;ipPairingConfiguration.anyIPAdress&#x60; to &#x60;false&#x60;, use &#x60;onlyTheseIpAddresses&#x60; to specify the IP addresses from which users can pair their devices. Each item in the array should be an IP address or an address range using CIDR notation, for example, &#x60;192.168.0.1/24&#x60; | [optional] 

## Methods

### NewDeviceAuthenticationPolicyPingIDAllOfMobileIpPairingConfiguration

`func NewDeviceAuthenticationPolicyPingIDAllOfMobileIpPairingConfiguration() *DeviceAuthenticationPolicyPingIDAllOfMobileIpPairingConfiguration`

NewDeviceAuthenticationPolicyPingIDAllOfMobileIpPairingConfiguration instantiates a new DeviceAuthenticationPolicyPingIDAllOfMobileIpPairingConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyPingIDAllOfMobileIpPairingConfigurationWithDefaults

`func NewDeviceAuthenticationPolicyPingIDAllOfMobileIpPairingConfigurationWithDefaults() *DeviceAuthenticationPolicyPingIDAllOfMobileIpPairingConfiguration`

NewDeviceAuthenticationPolicyPingIDAllOfMobileIpPairingConfigurationWithDefaults instantiates a new DeviceAuthenticationPolicyPingIDAllOfMobileIpPairingConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnyIPAdress

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileIpPairingConfiguration) GetAnyIPAdress() bool`

GetAnyIPAdress returns the AnyIPAdress field if non-nil, zero value otherwise.

### GetAnyIPAdressOk

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileIpPairingConfiguration) GetAnyIPAdressOk() (*bool, bool)`

GetAnyIPAdressOk returns a tuple with the AnyIPAdress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIPAdress

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileIpPairingConfiguration) SetAnyIPAdress(v bool)`

SetAnyIPAdress sets AnyIPAdress field to given value.

### HasAnyIPAdress

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileIpPairingConfiguration) HasAnyIPAdress() bool`

HasAnyIPAdress returns a boolean if a field has been set.

### GetOnlyTheseIpAddresses

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileIpPairingConfiguration) GetOnlyTheseIpAddresses() []string`

GetOnlyTheseIpAddresses returns the OnlyTheseIpAddresses field if non-nil, zero value otherwise.

### GetOnlyTheseIpAddressesOk

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileIpPairingConfiguration) GetOnlyTheseIpAddressesOk() (*[]string, bool)`

GetOnlyTheseIpAddressesOk returns a tuple with the OnlyTheseIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyTheseIpAddresses

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileIpPairingConfiguration) SetOnlyTheseIpAddresses(v []string)`

SetOnlyTheseIpAddresses sets OnlyTheseIpAddresses field to given value.

### HasOnlyTheseIpAddresses

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileIpPairingConfiguration) HasOnlyTheseIpAddresses() bool`

HasOnlyTheseIpAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


