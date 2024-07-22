# ApplicationTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | **map[string]string** | Contains a key/value map of the parameters required by the integration in Integration Catalog. | 
**Integration** | [**ApplicationTemplateIntegration**](ApplicationTemplateIntegration.md) |  | 
**Version** | [**ApplicationTemplateVersion**](ApplicationTemplateVersion.md) |  | 

## Methods

### NewApplicationTemplate

`func NewApplicationTemplate(configuration map[string]string, integration ApplicationTemplateIntegration, version ApplicationTemplateVersion, ) *ApplicationTemplate`

NewApplicationTemplate instantiates a new ApplicationTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationTemplateWithDefaults

`func NewApplicationTemplateWithDefaults() *ApplicationTemplate`

NewApplicationTemplateWithDefaults instantiates a new ApplicationTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *ApplicationTemplate) GetConfiguration() map[string]string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ApplicationTemplate) GetConfigurationOk() (*map[string]string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ApplicationTemplate) SetConfiguration(v map[string]string)`

SetConfiguration sets Configuration field to given value.


### GetIntegration

`func (o *ApplicationTemplate) GetIntegration() ApplicationTemplateIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *ApplicationTemplate) GetIntegrationOk() (*ApplicationTemplateIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *ApplicationTemplate) SetIntegration(v ApplicationTemplateIntegration)`

SetIntegration sets Integration field to given value.


### GetVersion

`func (o *ApplicationTemplate) GetVersion() ApplicationTemplateVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApplicationTemplate) GetVersionOk() (*ApplicationTemplateVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApplicationTemplate) SetVersion(v ApplicationTemplateVersion)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


