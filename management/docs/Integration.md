# Integration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to **[]string** | The list of categories used to classify the integration. Valid strings include alphanumeric characters, underscore, hyphen, and space. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation date of the integration. | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the integration in HTML to be used for the integration listing. You can use class and style attributes for inline styling. There&#39;s a 4000 character limit for the description. | [optional] 
**Id** | Pointer to **string** | The platform-generated ID of this integration. | [optional] [readonly] 
**Logo** | Pointer to [**IntegrationLogo**](IntegrationLogo.md) |  | [optional] 
**MarketingLandingPageUrl** | Pointer to **string** | Absolute URL link to the marketing landing page. | [optional] 
**Name** | **string** | Name of the integration. | 
**PingProductNames** | [**[]EnumIntegrationPingProductName**](EnumIntegrationPingProductName.md) | The Ping product associated with the integration. Can include &#x60;PINGID&#x60;, &#x60;PINGONE_ENTERPRISE&#x60;, &#x60;PINGONE&#x60;, &#x60;PINGACCESS&#x60;, &#x60;PINGFEDERATE&#x60;, &#x60;PINGDIRECTORY&#x60;, &#x60;PINGDATAGOVERNANCE&#x60;, &#x60;PINGINTELLIGENCE_FOR_APIS&#x60; or &#x60;PINGONE_ADVANCED_SERVICES&#x60; | 
**Publisher** | **string** | Name of the publisher. | 
**Tags** | Pointer to [**[]EnumIntegrationTag**](EnumIntegrationTag.md) | Tags to apply to the integration metadata. Can include &#x60;SSO&#x60;, &#x60;AUTHENTICATION&#x60;, &#x60;MFA&#x60;, &#x60;INTELLIGENCE&#x60;, &#x60;GOVERNANCE&#x60;, &#x60;IDAAS&#x60;, &#x60;ACCESS&#x60;, &#x60;DIRECTORY&#x60;, or &#x60;PROVISIONING&#x60;. | [optional] 
**ThirdParty** | Pointer to [**IntegrationThirdParty**](IntegrationThirdParty.md) |  | [optional] 

## Methods

### NewIntegration

`func NewIntegration(name string, pingProductNames []EnumIntegrationPingProductName, publisher string, ) *Integration`

NewIntegration instantiates a new Integration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationWithDefaults

`func NewIntegrationWithDefaults() *Integration`

NewIntegrationWithDefaults instantiates a new Integration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *Integration) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Integration) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Integration) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *Integration) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Integration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Integration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Integration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Integration) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *Integration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Integration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Integration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Integration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *Integration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Integration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Integration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Integration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogo

`func (o *Integration) GetLogo() IntegrationLogo`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *Integration) GetLogoOk() (*IntegrationLogo, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *Integration) SetLogo(v IntegrationLogo)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *Integration) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetMarketingLandingPageUrl

`func (o *Integration) GetMarketingLandingPageUrl() string`

GetMarketingLandingPageUrl returns the MarketingLandingPageUrl field if non-nil, zero value otherwise.

### GetMarketingLandingPageUrlOk

`func (o *Integration) GetMarketingLandingPageUrlOk() (*string, bool)`

GetMarketingLandingPageUrlOk returns a tuple with the MarketingLandingPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingLandingPageUrl

`func (o *Integration) SetMarketingLandingPageUrl(v string)`

SetMarketingLandingPageUrl sets MarketingLandingPageUrl field to given value.

### HasMarketingLandingPageUrl

`func (o *Integration) HasMarketingLandingPageUrl() bool`

HasMarketingLandingPageUrl returns a boolean if a field has been set.

### GetName

`func (o *Integration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Integration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Integration) SetName(v string)`

SetName sets Name field to given value.


### GetPingProductNames

`func (o *Integration) GetPingProductNames() []EnumIntegrationPingProductName`

GetPingProductNames returns the PingProductNames field if non-nil, zero value otherwise.

### GetPingProductNamesOk

`func (o *Integration) GetPingProductNamesOk() (*[]EnumIntegrationPingProductName, bool)`

GetPingProductNamesOk returns a tuple with the PingProductNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingProductNames

`func (o *Integration) SetPingProductNames(v []EnumIntegrationPingProductName)`

SetPingProductNames sets PingProductNames field to given value.


### GetPublisher

`func (o *Integration) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *Integration) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *Integration) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.


### GetTags

`func (o *Integration) GetTags() []EnumIntegrationTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Integration) GetTagsOk() (*[]EnumIntegrationTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Integration) SetTags(v []EnumIntegrationTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Integration) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThirdParty

`func (o *Integration) GetThirdParty() IntegrationThirdParty`

GetThirdParty returns the ThirdParty field if non-nil, zero value otherwise.

### GetThirdPartyOk

`func (o *Integration) GetThirdPartyOk() (*IntegrationThirdParty, bool)`

GetThirdPartyOk returns a tuple with the ThirdParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdParty

`func (o *Integration) SetThirdParty(v IntegrationThirdParty)`

SetThirdParty sets ThirdParty field to given value.

### HasThirdParty

`func (o *Integration) HasThirdParty() bool`

HasThirdParty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


