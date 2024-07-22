# IntegrationThirdParty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyName** | **string** | Name of the third party company for the listing. | 
**Products** | Pointer to **[]string** | Names of the third party products for the integration. | [optional] 

## Methods

### NewIntegrationThirdParty

`func NewIntegrationThirdParty(companyName string, ) *IntegrationThirdParty`

NewIntegrationThirdParty instantiates a new IntegrationThirdParty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationThirdPartyWithDefaults

`func NewIntegrationThirdPartyWithDefaults() *IntegrationThirdParty`

NewIntegrationThirdPartyWithDefaults instantiates a new IntegrationThirdParty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyName

`func (o *IntegrationThirdParty) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *IntegrationThirdParty) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *IntegrationThirdParty) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetProducts

`func (o *IntegrationThirdParty) GetProducts() []string`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *IntegrationThirdParty) GetProductsOk() (*[]string, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *IntegrationThirdParty) SetProducts(v []string)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *IntegrationThirdParty) HasProducts() bool`

HasProducts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


