# IntegrationVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to [**IntegrationVersionCommonConfiguration**](IntegrationVersionCommonConfiguration.md) |  | [optional] 
**Description** | Pointer to **string** | Unicode characters. The description of this integration metadata version. | [optional] 
**Id** | Pointer to **string** | The platform-generated ID of this integration metadata version. | [optional] [readonly] 
**Integration** | Pointer to [**IntegrationVersionCommonIntegration**](IntegrationVersionCommonIntegration.md) |  | [optional] 
**Name** | **string** | A unique name for the integration metadata version. | 
**Number** | **string** | A unique number for the integration version. | 
**Type** | Pointer to [**EnumIntegrationVersionType**](EnumIntegrationVersionType.md) |  | [optional] 
**DocumentationUrl** | Pointer to **string** | Absolute URL to the documentation. | [optional] 
**EndOfLifeOn** | Pointer to **string** | EOL support date in the form yyyy-mm-dd. | [optional] 
**IntegratedWith** | Pointer to [**IntegrationVersionIntegrationKitAllOfIntegratedWith**](IntegrationVersionIntegrationKitAllOfIntegratedWith.md) |  | [optional] 
**ReleasedOn** | **string** | Release date in the form yyyy-mm-dd. | 
**AssertionConsumerService** | **string** | The URL to which PingOne sends SAML responses. Parameterize the URL using &#x60;${_paremter_}&#x60;. For example, &#x60;https://${subdomain}.slack.com&#x60;. The maximum length is 2000 characters. | 
**AssertionEncrypted** | Pointer to **bool** | The state of assertion encryption. &#x60;true&#x60; if encrypted. | [optional] 
**DefaultTarget** | Pointer to **string** | After an IdP-initiated SSO, this URL is passed in the RelayState value in the SAML response. Informs the IdP where to send its response. Parameterize the URL using &#x60;${_paremter_}&#x60;. For example, &#x60;https://${subdomain}.slack.com&#x60;. | [optional] 
**EntityId** | Pointer to **string** | Unique ID for the application. | [optional] 
**NameIdFormat** | Pointer to **string** | This can be one of the following: urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress urn:oasis:names:tc:SAML:2.0:nameid-format:persistent urn:oasis:names:tc:SAML:2.0:nameid-format:transient  | [optional] 
**ProtocolVersion** | [**EnumIntegrationVersionSAMLProtocolVersion**](EnumIntegrationVersionSAMLProtocolVersion.md) |  | 
**Slo** | Pointer to [**IntegrationVersionSAMLAllOfSlo**](IntegrationVersionSAMLAllOfSlo.md) |  | [optional] 
**ThirdParty** | Pointer to [**IntegrationVersionSAMLAllOfThirdParty**](IntegrationVersionSAMLAllOfThirdParty.md) |  | [optional] 

## Methods

### NewIntegrationVersion

`func NewIntegrationVersion(name string, number string, releasedOn string, assertionConsumerService string, protocolVersion EnumIntegrationVersionSAMLProtocolVersion, ) *IntegrationVersion`

NewIntegrationVersion instantiates a new IntegrationVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationVersionWithDefaults

`func NewIntegrationVersionWithDefaults() *IntegrationVersion`

NewIntegrationVersionWithDefaults instantiates a new IntegrationVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *IntegrationVersion) GetConfiguration() IntegrationVersionCommonConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *IntegrationVersion) GetConfigurationOk() (*IntegrationVersionCommonConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *IntegrationVersion) SetConfiguration(v IntegrationVersionCommonConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *IntegrationVersion) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDescription

`func (o *IntegrationVersion) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IntegrationVersion) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IntegrationVersion) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IntegrationVersion) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *IntegrationVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationVersion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IntegrationVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIntegration

`func (o *IntegrationVersion) GetIntegration() IntegrationVersionCommonIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *IntegrationVersion) GetIntegrationOk() (*IntegrationVersionCommonIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *IntegrationVersion) SetIntegration(v IntegrationVersionCommonIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *IntegrationVersion) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetName

`func (o *IntegrationVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationVersion) SetName(v string)`

SetName sets Name field to given value.


### GetNumber

`func (o *IntegrationVersion) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *IntegrationVersion) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *IntegrationVersion) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetType

`func (o *IntegrationVersion) GetType() EnumIntegrationVersionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationVersion) GetTypeOk() (*EnumIntegrationVersionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationVersion) SetType(v EnumIntegrationVersionType)`

SetType sets Type field to given value.

### HasType

`func (o *IntegrationVersion) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDocumentationUrl

`func (o *IntegrationVersion) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *IntegrationVersion) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *IntegrationVersion) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *IntegrationVersion) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetEndOfLifeOn

`func (o *IntegrationVersion) GetEndOfLifeOn() string`

GetEndOfLifeOn returns the EndOfLifeOn field if non-nil, zero value otherwise.

### GetEndOfLifeOnOk

`func (o *IntegrationVersion) GetEndOfLifeOnOk() (*string, bool)`

GetEndOfLifeOnOk returns a tuple with the EndOfLifeOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOfLifeOn

`func (o *IntegrationVersion) SetEndOfLifeOn(v string)`

SetEndOfLifeOn sets EndOfLifeOn field to given value.

### HasEndOfLifeOn

`func (o *IntegrationVersion) HasEndOfLifeOn() bool`

HasEndOfLifeOn returns a boolean if a field has been set.

### GetIntegratedWith

`func (o *IntegrationVersion) GetIntegratedWith() IntegrationVersionIntegrationKitAllOfIntegratedWith`

GetIntegratedWith returns the IntegratedWith field if non-nil, zero value otherwise.

### GetIntegratedWithOk

`func (o *IntegrationVersion) GetIntegratedWithOk() (*IntegrationVersionIntegrationKitAllOfIntegratedWith, bool)`

GetIntegratedWithOk returns a tuple with the IntegratedWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegratedWith

`func (o *IntegrationVersion) SetIntegratedWith(v IntegrationVersionIntegrationKitAllOfIntegratedWith)`

SetIntegratedWith sets IntegratedWith field to given value.

### HasIntegratedWith

`func (o *IntegrationVersion) HasIntegratedWith() bool`

HasIntegratedWith returns a boolean if a field has been set.

### GetReleasedOn

`func (o *IntegrationVersion) GetReleasedOn() string`

GetReleasedOn returns the ReleasedOn field if non-nil, zero value otherwise.

### GetReleasedOnOk

`func (o *IntegrationVersion) GetReleasedOnOk() (*string, bool)`

GetReleasedOnOk returns a tuple with the ReleasedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedOn

`func (o *IntegrationVersion) SetReleasedOn(v string)`

SetReleasedOn sets ReleasedOn field to given value.


### GetAssertionConsumerService

`func (o *IntegrationVersion) GetAssertionConsumerService() string`

GetAssertionConsumerService returns the AssertionConsumerService field if non-nil, zero value otherwise.

### GetAssertionConsumerServiceOk

`func (o *IntegrationVersion) GetAssertionConsumerServiceOk() (*string, bool)`

GetAssertionConsumerServiceOk returns a tuple with the AssertionConsumerService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionConsumerService

`func (o *IntegrationVersion) SetAssertionConsumerService(v string)`

SetAssertionConsumerService sets AssertionConsumerService field to given value.


### GetAssertionEncrypted

`func (o *IntegrationVersion) GetAssertionEncrypted() bool`

GetAssertionEncrypted returns the AssertionEncrypted field if non-nil, zero value otherwise.

### GetAssertionEncryptedOk

`func (o *IntegrationVersion) GetAssertionEncryptedOk() (*bool, bool)`

GetAssertionEncryptedOk returns a tuple with the AssertionEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionEncrypted

`func (o *IntegrationVersion) SetAssertionEncrypted(v bool)`

SetAssertionEncrypted sets AssertionEncrypted field to given value.

### HasAssertionEncrypted

`func (o *IntegrationVersion) HasAssertionEncrypted() bool`

HasAssertionEncrypted returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *IntegrationVersion) GetDefaultTarget() string`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *IntegrationVersion) GetDefaultTargetOk() (*string, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *IntegrationVersion) SetDefaultTarget(v string)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *IntegrationVersion) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetEntityId

`func (o *IntegrationVersion) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *IntegrationVersion) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *IntegrationVersion) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *IntegrationVersion) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetNameIdFormat

`func (o *IntegrationVersion) GetNameIdFormat() string`

GetNameIdFormat returns the NameIdFormat field if non-nil, zero value otherwise.

### GetNameIdFormatOk

`func (o *IntegrationVersion) GetNameIdFormatOk() (*string, bool)`

GetNameIdFormatOk returns a tuple with the NameIdFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameIdFormat

`func (o *IntegrationVersion) SetNameIdFormat(v string)`

SetNameIdFormat sets NameIdFormat field to given value.

### HasNameIdFormat

`func (o *IntegrationVersion) HasNameIdFormat() bool`

HasNameIdFormat returns a boolean if a field has been set.

### GetProtocolVersion

`func (o *IntegrationVersion) GetProtocolVersion() EnumIntegrationVersionSAMLProtocolVersion`

GetProtocolVersion returns the ProtocolVersion field if non-nil, zero value otherwise.

### GetProtocolVersionOk

`func (o *IntegrationVersion) GetProtocolVersionOk() (*EnumIntegrationVersionSAMLProtocolVersion, bool)`

GetProtocolVersionOk returns a tuple with the ProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolVersion

`func (o *IntegrationVersion) SetProtocolVersion(v EnumIntegrationVersionSAMLProtocolVersion)`

SetProtocolVersion sets ProtocolVersion field to given value.


### GetSlo

`func (o *IntegrationVersion) GetSlo() IntegrationVersionSAMLAllOfSlo`

GetSlo returns the Slo field if non-nil, zero value otherwise.

### GetSloOk

`func (o *IntegrationVersion) GetSloOk() (*IntegrationVersionSAMLAllOfSlo, bool)`

GetSloOk returns a tuple with the Slo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlo

`func (o *IntegrationVersion) SetSlo(v IntegrationVersionSAMLAllOfSlo)`

SetSlo sets Slo field to given value.

### HasSlo

`func (o *IntegrationVersion) HasSlo() bool`

HasSlo returns a boolean if a field has been set.

### GetThirdParty

`func (o *IntegrationVersion) GetThirdParty() IntegrationVersionSAMLAllOfThirdParty`

GetThirdParty returns the ThirdParty field if non-nil, zero value otherwise.

### GetThirdPartyOk

`func (o *IntegrationVersion) GetThirdPartyOk() (*IntegrationVersionSAMLAllOfThirdParty, bool)`

GetThirdPartyOk returns a tuple with the ThirdParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdParty

`func (o *IntegrationVersion) SetThirdParty(v IntegrationVersionSAMLAllOfThirdParty)`

SetThirdParty sets ThirdParty field to given value.

### HasThirdParty

`func (o *IntegrationVersion) HasThirdParty() bool`

HasThirdParty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


