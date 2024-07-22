# IntegrationVersionSAML

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
**AssertionConsumerService** | **string** | The URL to which PingOne sends SAML responses. Parameterize the URL using &#x60;${_paremter_}&#x60;. For example, &#x60;https://${subdomain}.slack.com&#x60;. The maximum length is 2000 characters. | 
**AssertionEncrypted** | Pointer to **bool** | The state of assertion encryption. &#x60;true&#x60; if encrypted. | [optional] 
**DefaultTarget** | Pointer to **string** | After an IdP-initiated SSO, this URL is passed in the RelayState value in the SAML response. Informs the IdP where to send its response. Parameterize the URL using &#x60;${_paremter_}&#x60;. For example, &#x60;https://${subdomain}.slack.com&#x60;. | [optional] 
**EntityId** | Pointer to **string** | Unique ID for the application. | [optional] 
**NameIdFormat** | Pointer to **string** | This can be one of the following: urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress urn:oasis:names:tc:SAML:2.0:nameid-format:persistent urn:oasis:names:tc:SAML:2.0:nameid-format:transient  | [optional] 
**ProtocolVersion** | [**EnumIntegrationVersionSAMLProtocolVersion**](EnumIntegrationVersionSAMLProtocolVersion.md) |  | 
**Slo** | Pointer to [**IntegrationVersionSAMLAllOfSlo**](IntegrationVersionSAMLAllOfSlo.md) |  | [optional] 
**ThirdParty** | Pointer to [**IntegrationVersionSAMLAllOfThirdParty**](IntegrationVersionSAMLAllOfThirdParty.md) |  | [optional] 

## Methods

### NewIntegrationVersionSAML

`func NewIntegrationVersionSAML(name string, number string, assertionConsumerService string, protocolVersion EnumIntegrationVersionSAMLProtocolVersion, ) *IntegrationVersionSAML`

NewIntegrationVersionSAML instantiates a new IntegrationVersionSAML object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationVersionSAMLWithDefaults

`func NewIntegrationVersionSAMLWithDefaults() *IntegrationVersionSAML`

NewIntegrationVersionSAMLWithDefaults instantiates a new IntegrationVersionSAML object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *IntegrationVersionSAML) GetConfiguration() IntegrationVersionCommonConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *IntegrationVersionSAML) GetConfigurationOk() (*IntegrationVersionCommonConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *IntegrationVersionSAML) SetConfiguration(v IntegrationVersionCommonConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *IntegrationVersionSAML) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDescription

`func (o *IntegrationVersionSAML) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IntegrationVersionSAML) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IntegrationVersionSAML) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IntegrationVersionSAML) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *IntegrationVersionSAML) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationVersionSAML) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationVersionSAML) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IntegrationVersionSAML) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIntegration

`func (o *IntegrationVersionSAML) GetIntegration() IntegrationVersionCommonIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *IntegrationVersionSAML) GetIntegrationOk() (*IntegrationVersionCommonIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *IntegrationVersionSAML) SetIntegration(v IntegrationVersionCommonIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *IntegrationVersionSAML) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetName

`func (o *IntegrationVersionSAML) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationVersionSAML) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationVersionSAML) SetName(v string)`

SetName sets Name field to given value.


### GetNumber

`func (o *IntegrationVersionSAML) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *IntegrationVersionSAML) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *IntegrationVersionSAML) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetType

`func (o *IntegrationVersionSAML) GetType() EnumIntegrationVersionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationVersionSAML) GetTypeOk() (*EnumIntegrationVersionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationVersionSAML) SetType(v EnumIntegrationVersionType)`

SetType sets Type field to given value.

### HasType

`func (o *IntegrationVersionSAML) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAssertionConsumerService

`func (o *IntegrationVersionSAML) GetAssertionConsumerService() string`

GetAssertionConsumerService returns the AssertionConsumerService field if non-nil, zero value otherwise.

### GetAssertionConsumerServiceOk

`func (o *IntegrationVersionSAML) GetAssertionConsumerServiceOk() (*string, bool)`

GetAssertionConsumerServiceOk returns a tuple with the AssertionConsumerService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionConsumerService

`func (o *IntegrationVersionSAML) SetAssertionConsumerService(v string)`

SetAssertionConsumerService sets AssertionConsumerService field to given value.


### GetAssertionEncrypted

`func (o *IntegrationVersionSAML) GetAssertionEncrypted() bool`

GetAssertionEncrypted returns the AssertionEncrypted field if non-nil, zero value otherwise.

### GetAssertionEncryptedOk

`func (o *IntegrationVersionSAML) GetAssertionEncryptedOk() (*bool, bool)`

GetAssertionEncryptedOk returns a tuple with the AssertionEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionEncrypted

`func (o *IntegrationVersionSAML) SetAssertionEncrypted(v bool)`

SetAssertionEncrypted sets AssertionEncrypted field to given value.

### HasAssertionEncrypted

`func (o *IntegrationVersionSAML) HasAssertionEncrypted() bool`

HasAssertionEncrypted returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *IntegrationVersionSAML) GetDefaultTarget() string`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *IntegrationVersionSAML) GetDefaultTargetOk() (*string, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *IntegrationVersionSAML) SetDefaultTarget(v string)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *IntegrationVersionSAML) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetEntityId

`func (o *IntegrationVersionSAML) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *IntegrationVersionSAML) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *IntegrationVersionSAML) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *IntegrationVersionSAML) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetNameIdFormat

`func (o *IntegrationVersionSAML) GetNameIdFormat() string`

GetNameIdFormat returns the NameIdFormat field if non-nil, zero value otherwise.

### GetNameIdFormatOk

`func (o *IntegrationVersionSAML) GetNameIdFormatOk() (*string, bool)`

GetNameIdFormatOk returns a tuple with the NameIdFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameIdFormat

`func (o *IntegrationVersionSAML) SetNameIdFormat(v string)`

SetNameIdFormat sets NameIdFormat field to given value.

### HasNameIdFormat

`func (o *IntegrationVersionSAML) HasNameIdFormat() bool`

HasNameIdFormat returns a boolean if a field has been set.

### GetProtocolVersion

`func (o *IntegrationVersionSAML) GetProtocolVersion() EnumIntegrationVersionSAMLProtocolVersion`

GetProtocolVersion returns the ProtocolVersion field if non-nil, zero value otherwise.

### GetProtocolVersionOk

`func (o *IntegrationVersionSAML) GetProtocolVersionOk() (*EnumIntegrationVersionSAMLProtocolVersion, bool)`

GetProtocolVersionOk returns a tuple with the ProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolVersion

`func (o *IntegrationVersionSAML) SetProtocolVersion(v EnumIntegrationVersionSAMLProtocolVersion)`

SetProtocolVersion sets ProtocolVersion field to given value.


### GetSlo

`func (o *IntegrationVersionSAML) GetSlo() IntegrationVersionSAMLAllOfSlo`

GetSlo returns the Slo field if non-nil, zero value otherwise.

### GetSloOk

`func (o *IntegrationVersionSAML) GetSloOk() (*IntegrationVersionSAMLAllOfSlo, bool)`

GetSloOk returns a tuple with the Slo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlo

`func (o *IntegrationVersionSAML) SetSlo(v IntegrationVersionSAMLAllOfSlo)`

SetSlo sets Slo field to given value.

### HasSlo

`func (o *IntegrationVersionSAML) HasSlo() bool`

HasSlo returns a boolean if a field has been set.

### GetThirdParty

`func (o *IntegrationVersionSAML) GetThirdParty() IntegrationVersionSAMLAllOfThirdParty`

GetThirdParty returns the ThirdParty field if non-nil, zero value otherwise.

### GetThirdPartyOk

`func (o *IntegrationVersionSAML) GetThirdPartyOk() (*IntegrationVersionSAMLAllOfThirdParty, bool)`

GetThirdPartyOk returns a tuple with the ThirdParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdParty

`func (o *IntegrationVersionSAML) SetThirdParty(v IntegrationVersionSAMLAllOfThirdParty)`

SetThirdParty sets ThirdParty field to given value.

### HasThirdParty

`func (o *IntegrationVersionSAML) HasThirdParty() bool`

HasThirdParty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


