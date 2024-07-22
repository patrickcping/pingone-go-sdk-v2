# IntegrationVersionIntegrationKit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to [**IntegrationVersionCommonConfiguration**](IntegrationVersionCommonConfiguration.md) |  | [optional] 
**Description** | Pointer to **string** | Unicode characters. The description of this integration metadata version. | [optional] 
**Id** | Pointer to **string** | The platform-generated ID of this integration metadata version. | [optional] [readonly] 
**Integration** | Pointer to [**IntegrationVersionCommonIntegration**](IntegrationVersionCommonIntegration.md) |  | [optional] 
**Name** | **string** | A unique name for the integration metadata version. | 
**Number** | **string** | Unique number for the integration version. | 
**Type** | Pointer to [**EnumIntegrationVersionType**](EnumIntegrationVersionType.md) |  | [optional] 
**DocumentationUrl** | Pointer to **string** | Absolute URL to the documentation. | [optional] 
**EndOfLifeOn** | Pointer to **string** | EOL support date in the form yyyy-mm-dd. | [optional] 
**IntegratedWith** | Pointer to [**IntegrationVersionIntegrationKitAllOfIntegratedWith**](IntegrationVersionIntegrationKitAllOfIntegratedWith.md) |  | [optional] 
**ReleasedOn** | **string** | Release date in the form yyyy-mm-dd. | 

## Methods

### NewIntegrationVersionIntegrationKit

`func NewIntegrationVersionIntegrationKit(name string, number string, releasedOn string, ) *IntegrationVersionIntegrationKit`

NewIntegrationVersionIntegrationKit instantiates a new IntegrationVersionIntegrationKit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationVersionIntegrationKitWithDefaults

`func NewIntegrationVersionIntegrationKitWithDefaults() *IntegrationVersionIntegrationKit`

NewIntegrationVersionIntegrationKitWithDefaults instantiates a new IntegrationVersionIntegrationKit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *IntegrationVersionIntegrationKit) GetConfiguration() IntegrationVersionCommonConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *IntegrationVersionIntegrationKit) GetConfigurationOk() (*IntegrationVersionCommonConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *IntegrationVersionIntegrationKit) SetConfiguration(v IntegrationVersionCommonConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *IntegrationVersionIntegrationKit) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDescription

`func (o *IntegrationVersionIntegrationKit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IntegrationVersionIntegrationKit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IntegrationVersionIntegrationKit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IntegrationVersionIntegrationKit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *IntegrationVersionIntegrationKit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationVersionIntegrationKit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationVersionIntegrationKit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IntegrationVersionIntegrationKit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIntegration

`func (o *IntegrationVersionIntegrationKit) GetIntegration() IntegrationVersionCommonIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *IntegrationVersionIntegrationKit) GetIntegrationOk() (*IntegrationVersionCommonIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *IntegrationVersionIntegrationKit) SetIntegration(v IntegrationVersionCommonIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *IntegrationVersionIntegrationKit) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetName

`func (o *IntegrationVersionIntegrationKit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationVersionIntegrationKit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationVersionIntegrationKit) SetName(v string)`

SetName sets Name field to given value.


### GetNumber

`func (o *IntegrationVersionIntegrationKit) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *IntegrationVersionIntegrationKit) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *IntegrationVersionIntegrationKit) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetType

`func (o *IntegrationVersionIntegrationKit) GetType() EnumIntegrationVersionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationVersionIntegrationKit) GetTypeOk() (*EnumIntegrationVersionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationVersionIntegrationKit) SetType(v EnumIntegrationVersionType)`

SetType sets Type field to given value.

### HasType

`func (o *IntegrationVersionIntegrationKit) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDocumentationUrl

`func (o *IntegrationVersionIntegrationKit) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *IntegrationVersionIntegrationKit) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *IntegrationVersionIntegrationKit) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *IntegrationVersionIntegrationKit) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetEndOfLifeOn

`func (o *IntegrationVersionIntegrationKit) GetEndOfLifeOn() string`

GetEndOfLifeOn returns the EndOfLifeOn field if non-nil, zero value otherwise.

### GetEndOfLifeOnOk

`func (o *IntegrationVersionIntegrationKit) GetEndOfLifeOnOk() (*string, bool)`

GetEndOfLifeOnOk returns a tuple with the EndOfLifeOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOfLifeOn

`func (o *IntegrationVersionIntegrationKit) SetEndOfLifeOn(v string)`

SetEndOfLifeOn sets EndOfLifeOn field to given value.

### HasEndOfLifeOn

`func (o *IntegrationVersionIntegrationKit) HasEndOfLifeOn() bool`

HasEndOfLifeOn returns a boolean if a field has been set.

### GetIntegratedWith

`func (o *IntegrationVersionIntegrationKit) GetIntegratedWith() IntegrationVersionIntegrationKitAllOfIntegratedWith`

GetIntegratedWith returns the IntegratedWith field if non-nil, zero value otherwise.

### GetIntegratedWithOk

`func (o *IntegrationVersionIntegrationKit) GetIntegratedWithOk() (*IntegrationVersionIntegrationKitAllOfIntegratedWith, bool)`

GetIntegratedWithOk returns a tuple with the IntegratedWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegratedWith

`func (o *IntegrationVersionIntegrationKit) SetIntegratedWith(v IntegrationVersionIntegrationKitAllOfIntegratedWith)`

SetIntegratedWith sets IntegratedWith field to given value.

### HasIntegratedWith

`func (o *IntegrationVersionIntegrationKit) HasIntegratedWith() bool`

HasIntegratedWith returns a boolean if a field has been set.

### GetReleasedOn

`func (o *IntegrationVersionIntegrationKit) GetReleasedOn() string`

GetReleasedOn returns the ReleasedOn field if non-nil, zero value otherwise.

### GetReleasedOnOk

`func (o *IntegrationVersionIntegrationKit) GetReleasedOnOk() (*string, bool)`

GetReleasedOnOk returns a tuple with the ReleasedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedOn

`func (o *IntegrationVersionIntegrationKit) SetReleasedOn(v string)`

SetReleasedOn sets ReleasedOn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


