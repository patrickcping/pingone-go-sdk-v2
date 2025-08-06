# GovernmentIdConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailExpiredId** | Pointer to **bool** | Indicates whether verification should fail if the ID is expired. | [optional] 
**InspectionType** | Pointer to [**EnumInspectionType**](EnumInspectionType.md) |  | [optional] 
**Provider** | Pointer to [**GovernmentIdConfigurationProvider**](GovernmentIdConfigurationProvider.md) |  | [optional] 
**Retry** | Pointer to [**ObjectRetry**](ObjectRetry.md) |  | [optional] 
**Verify** | [**EnumVerify**](EnumVerify.md) |  | 
**VerifyAamva** | Pointer to **bool** | Whether [AAMVA DLDV](https://apidocs.pingidentity.com/pingone/platform/v1/api/#us-based-driver-licenses) verification is enabled | [optional] 

## Methods

### NewGovernmentIdConfiguration

`func NewGovernmentIdConfiguration(verify EnumVerify, ) *GovernmentIdConfiguration`

NewGovernmentIdConfiguration instantiates a new GovernmentIdConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGovernmentIdConfigurationWithDefaults

`func NewGovernmentIdConfigurationWithDefaults() *GovernmentIdConfiguration`

NewGovernmentIdConfigurationWithDefaults instantiates a new GovernmentIdConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailExpiredId

`func (o *GovernmentIdConfiguration) GetFailExpiredId() bool`

GetFailExpiredId returns the FailExpiredId field if non-nil, zero value otherwise.

### GetFailExpiredIdOk

`func (o *GovernmentIdConfiguration) GetFailExpiredIdOk() (*bool, bool)`

GetFailExpiredIdOk returns a tuple with the FailExpiredId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailExpiredId

`func (o *GovernmentIdConfiguration) SetFailExpiredId(v bool)`

SetFailExpiredId sets FailExpiredId field to given value.

### HasFailExpiredId

`func (o *GovernmentIdConfiguration) HasFailExpiredId() bool`

HasFailExpiredId returns a boolean if a field has been set.

### GetInspectionType

`func (o *GovernmentIdConfiguration) GetInspectionType() EnumInspectionType`

GetInspectionType returns the InspectionType field if non-nil, zero value otherwise.

### GetInspectionTypeOk

`func (o *GovernmentIdConfiguration) GetInspectionTypeOk() (*EnumInspectionType, bool)`

GetInspectionTypeOk returns a tuple with the InspectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectionType

`func (o *GovernmentIdConfiguration) SetInspectionType(v EnumInspectionType)`

SetInspectionType sets InspectionType field to given value.

### HasInspectionType

`func (o *GovernmentIdConfiguration) HasInspectionType() bool`

HasInspectionType returns a boolean if a field has been set.

### GetProvider

`func (o *GovernmentIdConfiguration) GetProvider() GovernmentIdConfigurationProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GovernmentIdConfiguration) GetProviderOk() (*GovernmentIdConfigurationProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GovernmentIdConfiguration) SetProvider(v GovernmentIdConfigurationProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *GovernmentIdConfiguration) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRetry

`func (o *GovernmentIdConfiguration) GetRetry() ObjectRetry`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *GovernmentIdConfiguration) GetRetryOk() (*ObjectRetry, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *GovernmentIdConfiguration) SetRetry(v ObjectRetry)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *GovernmentIdConfiguration) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### GetVerify

`func (o *GovernmentIdConfiguration) GetVerify() EnumVerify`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *GovernmentIdConfiguration) GetVerifyOk() (*EnumVerify, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *GovernmentIdConfiguration) SetVerify(v EnumVerify)`

SetVerify sets Verify field to given value.


### GetVerifyAamva

`func (o *GovernmentIdConfiguration) GetVerifyAamva() bool`

GetVerifyAamva returns the VerifyAamva field if non-nil, zero value otherwise.

### GetVerifyAamvaOk

`func (o *GovernmentIdConfiguration) GetVerifyAamvaOk() (*bool, bool)`

GetVerifyAamvaOk returns a tuple with the VerifyAamva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyAamva

`func (o *GovernmentIdConfiguration) SetVerifyAamva(v bool)`

SetVerifyAamva sets VerifyAamva field to given value.

### HasVerifyAamva

`func (o *GovernmentIdConfiguration) HasVerifyAamva() bool`

HasVerifyAamva returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


