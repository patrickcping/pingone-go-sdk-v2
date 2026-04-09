# EntityArrayEmbedded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationAttributes** | Pointer to [**[]AuthorizeEditorDataDefinitionsAttributeDefinitionDTO**](AuthorizeEditorDataDefinitionsAttributeDefinitionDTO.md) |  | [optional] 
**AuthorizationConditions** | Pointer to [**[]AuthorizeEditorDataDefinitionsConditionDefinitionDTO**](AuthorizeEditorDataDefinitionsConditionDefinitionDTO.md) |  | [optional] 
**AuthorizationConnectorTemplates** | Pointer to [**[]AuthorizeEditorDataConnectorsConnectorTemplateDTO**](AuthorizeEditorDataConnectorsConnectorTemplateDTO.md) |  | [optional] 
**AuthorizationPolicies** | Pointer to [**[]AuthorizeEditorDataPoliciesReferenceablePolicyDTO**](AuthorizeEditorDataPoliciesReferenceablePolicyDTO.md) |  | [optional] 
**AuthorizationProcessors** | Pointer to [**[]AuthorizeEditorDataDefinitionsProcessorDefinitionDTO**](AuthorizeEditorDataDefinitionsProcessorDefinitionDTO.md) |  | [optional] 
**AuthorizationRules** | Pointer to [**[]AuthorizeEditorDataRulesReferenceableRuleDTO**](AuthorizeEditorDataRulesReferenceableRuleDTO.md) |  | [optional] 
**AuthorizationServices** | Pointer to [**[]AuthorizeEditorDataDefinitionsServiceDefinitionDTO**](AuthorizeEditorDataDefinitionsServiceDefinitionDTO.md) |  | [optional] 
**AuthorizationStatements** | Pointer to [**[]AuthorizeEditorDataStatementsReferenceableStatementDTO**](AuthorizeEditorDataStatementsReferenceableStatementDTO.md) |  | [optional] 
**AuthorizationVersions** | Pointer to [**[]AuthorizeEditorDataAuthorizationVersionDTO**](AuthorizeEditorDataAuthorizationVersionDTO.md) |  | [optional] 

## Methods

### NewEntityArrayEmbedded

`func NewEntityArrayEmbedded() *EntityArrayEmbedded`

NewEntityArrayEmbedded instantiates a new EntityArrayEmbedded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityArrayEmbeddedWithDefaults

`func NewEntityArrayEmbeddedWithDefaults() *EntityArrayEmbedded`

NewEntityArrayEmbeddedWithDefaults instantiates a new EntityArrayEmbedded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationAttributes

`func (o *EntityArrayEmbedded) GetAuthorizationAttributes() []AuthorizeEditorDataDefinitionsAttributeDefinitionDTO`

GetAuthorizationAttributes returns the AuthorizationAttributes field if non-nil, zero value otherwise.

### GetAuthorizationAttributesOk

`func (o *EntityArrayEmbedded) GetAuthorizationAttributesOk() (*[]AuthorizeEditorDataDefinitionsAttributeDefinitionDTO, bool)`

GetAuthorizationAttributesOk returns a tuple with the AuthorizationAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationAttributes

`func (o *EntityArrayEmbedded) SetAuthorizationAttributes(v []AuthorizeEditorDataDefinitionsAttributeDefinitionDTO)`

SetAuthorizationAttributes sets AuthorizationAttributes field to given value.

### HasAuthorizationAttributes

`func (o *EntityArrayEmbedded) HasAuthorizationAttributes() bool`

HasAuthorizationAttributes returns a boolean if a field has been set.

### GetAuthorizationConditions

`func (o *EntityArrayEmbedded) GetAuthorizationConditions() []AuthorizeEditorDataDefinitionsConditionDefinitionDTO`

GetAuthorizationConditions returns the AuthorizationConditions field if non-nil, zero value otherwise.

### GetAuthorizationConditionsOk

`func (o *EntityArrayEmbedded) GetAuthorizationConditionsOk() (*[]AuthorizeEditorDataDefinitionsConditionDefinitionDTO, bool)`

GetAuthorizationConditionsOk returns a tuple with the AuthorizationConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationConditions

`func (o *EntityArrayEmbedded) SetAuthorizationConditions(v []AuthorizeEditorDataDefinitionsConditionDefinitionDTO)`

SetAuthorizationConditions sets AuthorizationConditions field to given value.

### HasAuthorizationConditions

`func (o *EntityArrayEmbedded) HasAuthorizationConditions() bool`

HasAuthorizationConditions returns a boolean if a field has been set.

### GetAuthorizationConnectorTemplates

`func (o *EntityArrayEmbedded) GetAuthorizationConnectorTemplates() []AuthorizeEditorDataConnectorsConnectorTemplateDTO`

GetAuthorizationConnectorTemplates returns the AuthorizationConnectorTemplates field if non-nil, zero value otherwise.

### GetAuthorizationConnectorTemplatesOk

`func (o *EntityArrayEmbedded) GetAuthorizationConnectorTemplatesOk() (*[]AuthorizeEditorDataConnectorsConnectorTemplateDTO, bool)`

GetAuthorizationConnectorTemplatesOk returns a tuple with the AuthorizationConnectorTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationConnectorTemplates

`func (o *EntityArrayEmbedded) SetAuthorizationConnectorTemplates(v []AuthorizeEditorDataConnectorsConnectorTemplateDTO)`

SetAuthorizationConnectorTemplates sets AuthorizationConnectorTemplates field to given value.

### HasAuthorizationConnectorTemplates

`func (o *EntityArrayEmbedded) HasAuthorizationConnectorTemplates() bool`

HasAuthorizationConnectorTemplates returns a boolean if a field has been set.

### GetAuthorizationPolicies

`func (o *EntityArrayEmbedded) GetAuthorizationPolicies() []AuthorizeEditorDataPoliciesReferenceablePolicyDTO`

GetAuthorizationPolicies returns the AuthorizationPolicies field if non-nil, zero value otherwise.

### GetAuthorizationPoliciesOk

`func (o *EntityArrayEmbedded) GetAuthorizationPoliciesOk() (*[]AuthorizeEditorDataPoliciesReferenceablePolicyDTO, bool)`

GetAuthorizationPoliciesOk returns a tuple with the AuthorizationPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationPolicies

`func (o *EntityArrayEmbedded) SetAuthorizationPolicies(v []AuthorizeEditorDataPoliciesReferenceablePolicyDTO)`

SetAuthorizationPolicies sets AuthorizationPolicies field to given value.

### HasAuthorizationPolicies

`func (o *EntityArrayEmbedded) HasAuthorizationPolicies() bool`

HasAuthorizationPolicies returns a boolean if a field has been set.

### GetAuthorizationProcessors

`func (o *EntityArrayEmbedded) GetAuthorizationProcessors() []AuthorizeEditorDataDefinitionsProcessorDefinitionDTO`

GetAuthorizationProcessors returns the AuthorizationProcessors field if non-nil, zero value otherwise.

### GetAuthorizationProcessorsOk

`func (o *EntityArrayEmbedded) GetAuthorizationProcessorsOk() (*[]AuthorizeEditorDataDefinitionsProcessorDefinitionDTO, bool)`

GetAuthorizationProcessorsOk returns a tuple with the AuthorizationProcessors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationProcessors

`func (o *EntityArrayEmbedded) SetAuthorizationProcessors(v []AuthorizeEditorDataDefinitionsProcessorDefinitionDTO)`

SetAuthorizationProcessors sets AuthorizationProcessors field to given value.

### HasAuthorizationProcessors

`func (o *EntityArrayEmbedded) HasAuthorizationProcessors() bool`

HasAuthorizationProcessors returns a boolean if a field has been set.

### GetAuthorizationRules

`func (o *EntityArrayEmbedded) GetAuthorizationRules() []AuthorizeEditorDataRulesReferenceableRuleDTO`

GetAuthorizationRules returns the AuthorizationRules field if non-nil, zero value otherwise.

### GetAuthorizationRulesOk

`func (o *EntityArrayEmbedded) GetAuthorizationRulesOk() (*[]AuthorizeEditorDataRulesReferenceableRuleDTO, bool)`

GetAuthorizationRulesOk returns a tuple with the AuthorizationRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationRules

`func (o *EntityArrayEmbedded) SetAuthorizationRules(v []AuthorizeEditorDataRulesReferenceableRuleDTO)`

SetAuthorizationRules sets AuthorizationRules field to given value.

### HasAuthorizationRules

`func (o *EntityArrayEmbedded) HasAuthorizationRules() bool`

HasAuthorizationRules returns a boolean if a field has been set.

### GetAuthorizationServices

`func (o *EntityArrayEmbedded) GetAuthorizationServices() []AuthorizeEditorDataDefinitionsServiceDefinitionDTO`

GetAuthorizationServices returns the AuthorizationServices field if non-nil, zero value otherwise.

### GetAuthorizationServicesOk

`func (o *EntityArrayEmbedded) GetAuthorizationServicesOk() (*[]AuthorizeEditorDataDefinitionsServiceDefinitionDTO, bool)`

GetAuthorizationServicesOk returns a tuple with the AuthorizationServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationServices

`func (o *EntityArrayEmbedded) SetAuthorizationServices(v []AuthorizeEditorDataDefinitionsServiceDefinitionDTO)`

SetAuthorizationServices sets AuthorizationServices field to given value.

### HasAuthorizationServices

`func (o *EntityArrayEmbedded) HasAuthorizationServices() bool`

HasAuthorizationServices returns a boolean if a field has been set.

### GetAuthorizationStatements

`func (o *EntityArrayEmbedded) GetAuthorizationStatements() []AuthorizeEditorDataStatementsReferenceableStatementDTO`

GetAuthorizationStatements returns the AuthorizationStatements field if non-nil, zero value otherwise.

### GetAuthorizationStatementsOk

`func (o *EntityArrayEmbedded) GetAuthorizationStatementsOk() (*[]AuthorizeEditorDataStatementsReferenceableStatementDTO, bool)`

GetAuthorizationStatementsOk returns a tuple with the AuthorizationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationStatements

`func (o *EntityArrayEmbedded) SetAuthorizationStatements(v []AuthorizeEditorDataStatementsReferenceableStatementDTO)`

SetAuthorizationStatements sets AuthorizationStatements field to given value.

### HasAuthorizationStatements

`func (o *EntityArrayEmbedded) HasAuthorizationStatements() bool`

HasAuthorizationStatements returns a boolean if a field has been set.

### GetAuthorizationVersions

`func (o *EntityArrayEmbedded) GetAuthorizationVersions() []AuthorizeEditorDataAuthorizationVersionDTO`

GetAuthorizationVersions returns the AuthorizationVersions field if non-nil, zero value otherwise.

### GetAuthorizationVersionsOk

`func (o *EntityArrayEmbedded) GetAuthorizationVersionsOk() (*[]AuthorizeEditorDataAuthorizationVersionDTO, bool)`

GetAuthorizationVersionsOk returns a tuple with the AuthorizationVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationVersions

`func (o *EntityArrayEmbedded) SetAuthorizationVersions(v []AuthorizeEditorDataAuthorizationVersionDTO)`

SetAuthorizationVersions sets AuthorizationVersions field to given value.

### HasAuthorizationVersions

`func (o *EntityArrayEmbedded) HasAuthorizationVersions() bool`

HasAuthorizationVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


