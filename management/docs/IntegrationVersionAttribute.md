# IntegrationVersionAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Auto-generated ID of this attribute. | [optional] [readonly] 
**Integration** | Pointer to [**IntegrationVersionAttributeIntegration**](IntegrationVersionAttributeIntegration.md) |  | [optional] 
**Name** | Pointer to **string** | Attribute name the application expects. Unique within the integration version and in the form &#x60;urn:oasis:names:tc:SAML:2.0:attrname-format:uri&#x60;. | [optional] 
**Required** | Pointer to **bool** | Whether or not the attribute is required. If true, the value property must be set with a non-empty value. Default is false. | [optional] 
**Schema** | **string** | A JSON schema describing the current attribute mapping. | 
**Version** | Pointer to [**IntegrationVersionAttributeVersion**](IntegrationVersionAttributeVersion.md) |  | [optional] 

## Methods

### NewIntegrationVersionAttribute

`func NewIntegrationVersionAttribute(schema string, ) *IntegrationVersionAttribute`

NewIntegrationVersionAttribute instantiates a new IntegrationVersionAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationVersionAttributeWithDefaults

`func NewIntegrationVersionAttributeWithDefaults() *IntegrationVersionAttribute`

NewIntegrationVersionAttributeWithDefaults instantiates a new IntegrationVersionAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationVersionAttribute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationVersionAttribute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationVersionAttribute) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IntegrationVersionAttribute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIntegration

`func (o *IntegrationVersionAttribute) GetIntegration() IntegrationVersionAttributeIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *IntegrationVersionAttribute) GetIntegrationOk() (*IntegrationVersionAttributeIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *IntegrationVersionAttribute) SetIntegration(v IntegrationVersionAttributeIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *IntegrationVersionAttribute) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetName

`func (o *IntegrationVersionAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationVersionAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationVersionAttribute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IntegrationVersionAttribute) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequired

`func (o *IntegrationVersionAttribute) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *IntegrationVersionAttribute) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *IntegrationVersionAttribute) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *IntegrationVersionAttribute) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetSchema

`func (o *IntegrationVersionAttribute) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *IntegrationVersionAttribute) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *IntegrationVersionAttribute) SetSchema(v string)`

SetSchema sets Schema field to given value.


### GetVersion

`func (o *IntegrationVersionAttribute) GetVersion() IntegrationVersionAttributeVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IntegrationVersionAttribute) GetVersionOk() (*IntegrationVersionAttributeVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IntegrationVersionAttribute) SetVersion(v IntegrationVersionAttributeVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IntegrationVersionAttribute) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


