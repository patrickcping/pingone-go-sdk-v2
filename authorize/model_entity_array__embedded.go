/*
PingOne Platform API - Authorize

The PingOne Platform API covering the PingOne Authorize service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorize

import (
	"encoding/json"
)

// checks if the EntityArrayEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityArrayEmbedded{}

// EntityArrayEmbedded struct for EntityArrayEmbedded
type EntityArrayEmbedded struct {
	ApiServers []APIServer `json:"apiServers,omitempty"`
	Assignments []ApplicationRoleAssignment `json:"assignments,omitempty"`
	AuthorizationAttributes []AuthorizeEditorDataDefinitionsAttributeDefinitionDTO `json:"authorizationAttributes,omitempty"`
	AuthorizationConditions []AuthorizeEditorDataDefinitionsConditionDefinitionDTO `json:"authorizationConditions,omitempty"`
	AuthorizationConnectorTemplates []AuthorizeEditorDataConnectorsConnectorTemplateDTO `json:"authorizationConnectorTemplates,omitempty"`
	AuthorizationPolicies []AuthorizeEditorDataPoliciesReferenceablePolicyDTO `json:"authorizationPolicies,omitempty"`
	AuthorizationProcessors []AuthorizeEditorDataDefinitionsProcessorDefinitionDTO `json:"authorizationProcessors,omitempty"`
	AuthorizationRules []AuthorizeEditorDataRulesReferenceableRuleDTO `json:"authorizationRules,omitempty"`
	AuthorizationServices []AuthorizeEditorDataDefinitionsServiceDefinitionDTO `json:"authorizationServices,omitempty"`
	AuthorizationStatements []AuthorizeEditorDataStatementsReferenceableStatementDTO `json:"authorizationStatements,omitempty"`
	AuthorizationVersions []AuthorizeEditorDataAuthorizationVersionDTO `json:"authorizationVersions,omitempty"`
	DecisionEndpoints []DecisionEndpoint `json:"decisionEndpoints,omitempty"`
	Operations []APIServerOperation `json:"operations,omitempty"`
	Permissions []EntityArrayEmbeddedPermissionsInner `json:"permissions,omitempty"`
	Resources []ApplicationResource `json:"resources,omitempty"`
	Roles []ApplicationRole `json:"roles,omitempty"`
}

// NewEntityArrayEmbedded instantiates a new EntityArrayEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityArrayEmbedded() *EntityArrayEmbedded {
	this := EntityArrayEmbedded{}
	return &this
}

// NewEntityArrayEmbeddedWithDefaults instantiates a new EntityArrayEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityArrayEmbeddedWithDefaults() *EntityArrayEmbedded {
	this := EntityArrayEmbedded{}
	return &this
}

// GetApiServers returns the ApiServers field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetApiServers() []APIServer {
	if o == nil || IsNil(o.ApiServers) {
		var ret []APIServer
		return ret
	}
	return o.ApiServers
}

// GetApiServersOk returns a tuple with the ApiServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetApiServersOk() ([]APIServer, bool) {
	if o == nil || IsNil(o.ApiServers) {
		return nil, false
	}
	return o.ApiServers, true
}

// HasApiServers returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasApiServers() bool {
	if o != nil && !IsNil(o.ApiServers) {
		return true
	}

	return false
}

// SetApiServers gets a reference to the given []APIServer and assigns it to the ApiServers field.
func (o *EntityArrayEmbedded) SetApiServers(v []APIServer) {
	o.ApiServers = v
}

// GetAssignments returns the Assignments field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetAssignments() []ApplicationRoleAssignment {
	if o == nil || IsNil(o.Assignments) {
		var ret []ApplicationRoleAssignment
		return ret
	}
	return o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetAssignmentsOk() ([]ApplicationRoleAssignment, bool) {
	if o == nil || IsNil(o.Assignments) {
		return nil, false
	}
	return o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasAssignments() bool {
	if o != nil && !IsNil(o.Assignments) {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []ApplicationRoleAssignment and assigns it to the Assignments field.
func (o *EntityArrayEmbedded) SetAssignments(v []ApplicationRoleAssignment) {
	o.Assignments = v
}

// GetAuthorizationAttributes returns the AuthorizationAttributes field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetAuthorizationAttributes() []AuthorizeEditorDataDefinitionsAttributeDefinitionDTO {
	if o == nil || IsNil(o.AuthorizationAttributes) {
		var ret []AuthorizeEditorDataDefinitionsAttributeDefinitionDTO
		return ret
	}
	return o.AuthorizationAttributes
}

// GetAuthorizationAttributesOk returns a tuple with the AuthorizationAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetAuthorizationAttributesOk() ([]AuthorizeEditorDataDefinitionsAttributeDefinitionDTO, bool) {
	if o == nil || IsNil(o.AuthorizationAttributes) {
		return nil, false
	}
	return o.AuthorizationAttributes, true
}

// HasAuthorizationAttributes returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasAuthorizationAttributes() bool {
	if o != nil && !IsNil(o.AuthorizationAttributes) {
		return true
	}

	return false
}

// SetAuthorizationAttributes gets a reference to the given []AuthorizeEditorDataDefinitionsAttributeDefinitionDTO and assigns it to the AuthorizationAttributes field.
func (o *EntityArrayEmbedded) SetAuthorizationAttributes(v []AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) {
	o.AuthorizationAttributes = v
}

// GetAuthorizationConditions returns the AuthorizationConditions field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetAuthorizationConditions() []AuthorizeEditorDataDefinitionsConditionDefinitionDTO {
	if o == nil || IsNil(o.AuthorizationConditions) {
		var ret []AuthorizeEditorDataDefinitionsConditionDefinitionDTO
		return ret
	}
	return o.AuthorizationConditions
}

// GetAuthorizationConditionsOk returns a tuple with the AuthorizationConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetAuthorizationConditionsOk() ([]AuthorizeEditorDataDefinitionsConditionDefinitionDTO, bool) {
	if o == nil || IsNil(o.AuthorizationConditions) {
		return nil, false
	}
	return o.AuthorizationConditions, true
}

// HasAuthorizationConditions returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasAuthorizationConditions() bool {
	if o != nil && !IsNil(o.AuthorizationConditions) {
		return true
	}

	return false
}

// SetAuthorizationConditions gets a reference to the given []AuthorizeEditorDataDefinitionsConditionDefinitionDTO and assigns it to the AuthorizationConditions field.
func (o *EntityArrayEmbedded) SetAuthorizationConditions(v []AuthorizeEditorDataDefinitionsConditionDefinitionDTO) {
	o.AuthorizationConditions = v
}

// GetAuthorizationConnectorTemplates returns the AuthorizationConnectorTemplates field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetAuthorizationConnectorTemplates() []AuthorizeEditorDataConnectorsConnectorTemplateDTO {
	if o == nil || IsNil(o.AuthorizationConnectorTemplates) {
		var ret []AuthorizeEditorDataConnectorsConnectorTemplateDTO
		return ret
	}
	return o.AuthorizationConnectorTemplates
}

// GetAuthorizationConnectorTemplatesOk returns a tuple with the AuthorizationConnectorTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetAuthorizationConnectorTemplatesOk() ([]AuthorizeEditorDataConnectorsConnectorTemplateDTO, bool) {
	if o == nil || IsNil(o.AuthorizationConnectorTemplates) {
		return nil, false
	}
	return o.AuthorizationConnectorTemplates, true
}

// HasAuthorizationConnectorTemplates returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasAuthorizationConnectorTemplates() bool {
	if o != nil && !IsNil(o.AuthorizationConnectorTemplates) {
		return true
	}

	return false
}

// SetAuthorizationConnectorTemplates gets a reference to the given []AuthorizeEditorDataConnectorsConnectorTemplateDTO and assigns it to the AuthorizationConnectorTemplates field.
func (o *EntityArrayEmbedded) SetAuthorizationConnectorTemplates(v []AuthorizeEditorDataConnectorsConnectorTemplateDTO) {
	o.AuthorizationConnectorTemplates = v
}

// GetAuthorizationPolicies returns the AuthorizationPolicies field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetAuthorizationPolicies() []AuthorizeEditorDataPoliciesReferenceablePolicyDTO {
	if o == nil || IsNil(o.AuthorizationPolicies) {
		var ret []AuthorizeEditorDataPoliciesReferenceablePolicyDTO
		return ret
	}
	return o.AuthorizationPolicies
}

// GetAuthorizationPoliciesOk returns a tuple with the AuthorizationPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetAuthorizationPoliciesOk() ([]AuthorizeEditorDataPoliciesReferenceablePolicyDTO, bool) {
	if o == nil || IsNil(o.AuthorizationPolicies) {
		return nil, false
	}
	return o.AuthorizationPolicies, true
}

// HasAuthorizationPolicies returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasAuthorizationPolicies() bool {
	if o != nil && !IsNil(o.AuthorizationPolicies) {
		return true
	}

	return false
}

// SetAuthorizationPolicies gets a reference to the given []AuthorizeEditorDataPoliciesReferenceablePolicyDTO and assigns it to the AuthorizationPolicies field.
func (o *EntityArrayEmbedded) SetAuthorizationPolicies(v []AuthorizeEditorDataPoliciesReferenceablePolicyDTO) {
	o.AuthorizationPolicies = v
}

// GetAuthorizationProcessors returns the AuthorizationProcessors field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetAuthorizationProcessors() []AuthorizeEditorDataDefinitionsProcessorDefinitionDTO {
	if o == nil || IsNil(o.AuthorizationProcessors) {
		var ret []AuthorizeEditorDataDefinitionsProcessorDefinitionDTO
		return ret
	}
	return o.AuthorizationProcessors
}

// GetAuthorizationProcessorsOk returns a tuple with the AuthorizationProcessors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetAuthorizationProcessorsOk() ([]AuthorizeEditorDataDefinitionsProcessorDefinitionDTO, bool) {
	if o == nil || IsNil(o.AuthorizationProcessors) {
		return nil, false
	}
	return o.AuthorizationProcessors, true
}

// HasAuthorizationProcessors returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasAuthorizationProcessors() bool {
	if o != nil && !IsNil(o.AuthorizationProcessors) {
		return true
	}

	return false
}

// SetAuthorizationProcessors gets a reference to the given []AuthorizeEditorDataDefinitionsProcessorDefinitionDTO and assigns it to the AuthorizationProcessors field.
func (o *EntityArrayEmbedded) SetAuthorizationProcessors(v []AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) {
	o.AuthorizationProcessors = v
}

// GetAuthorizationRules returns the AuthorizationRules field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetAuthorizationRules() []AuthorizeEditorDataRulesReferenceableRuleDTO {
	if o == nil || IsNil(o.AuthorizationRules) {
		var ret []AuthorizeEditorDataRulesReferenceableRuleDTO
		return ret
	}
	return o.AuthorizationRules
}

// GetAuthorizationRulesOk returns a tuple with the AuthorizationRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetAuthorizationRulesOk() ([]AuthorizeEditorDataRulesReferenceableRuleDTO, bool) {
	if o == nil || IsNil(o.AuthorizationRules) {
		return nil, false
	}
	return o.AuthorizationRules, true
}

// HasAuthorizationRules returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasAuthorizationRules() bool {
	if o != nil && !IsNil(o.AuthorizationRules) {
		return true
	}

	return false
}

// SetAuthorizationRules gets a reference to the given []AuthorizeEditorDataRulesReferenceableRuleDTO and assigns it to the AuthorizationRules field.
func (o *EntityArrayEmbedded) SetAuthorizationRules(v []AuthorizeEditorDataRulesReferenceableRuleDTO) {
	o.AuthorizationRules = v
}

// GetAuthorizationServices returns the AuthorizationServices field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetAuthorizationServices() []AuthorizeEditorDataDefinitionsServiceDefinitionDTO {
	if o == nil || IsNil(o.AuthorizationServices) {
		var ret []AuthorizeEditorDataDefinitionsServiceDefinitionDTO
		return ret
	}
	return o.AuthorizationServices
}

// GetAuthorizationServicesOk returns a tuple with the AuthorizationServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetAuthorizationServicesOk() ([]AuthorizeEditorDataDefinitionsServiceDefinitionDTO, bool) {
	if o == nil || IsNil(o.AuthorizationServices) {
		return nil, false
	}
	return o.AuthorizationServices, true
}

// HasAuthorizationServices returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasAuthorizationServices() bool {
	if o != nil && !IsNil(o.AuthorizationServices) {
		return true
	}

	return false
}

// SetAuthorizationServices gets a reference to the given []AuthorizeEditorDataDefinitionsServiceDefinitionDTO and assigns it to the AuthorizationServices field.
func (o *EntityArrayEmbedded) SetAuthorizationServices(v []AuthorizeEditorDataDefinitionsServiceDefinitionDTO) {
	o.AuthorizationServices = v
}

// GetAuthorizationStatements returns the AuthorizationStatements field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetAuthorizationStatements() []AuthorizeEditorDataStatementsReferenceableStatementDTO {
	if o == nil || IsNil(o.AuthorizationStatements) {
		var ret []AuthorizeEditorDataStatementsReferenceableStatementDTO
		return ret
	}
	return o.AuthorizationStatements
}

// GetAuthorizationStatementsOk returns a tuple with the AuthorizationStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetAuthorizationStatementsOk() ([]AuthorizeEditorDataStatementsReferenceableStatementDTO, bool) {
	if o == nil || IsNil(o.AuthorizationStatements) {
		return nil, false
	}
	return o.AuthorizationStatements, true
}

// HasAuthorizationStatements returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasAuthorizationStatements() bool {
	if o != nil && !IsNil(o.AuthorizationStatements) {
		return true
	}

	return false
}

// SetAuthorizationStatements gets a reference to the given []AuthorizeEditorDataStatementsReferenceableStatementDTO and assigns it to the AuthorizationStatements field.
func (o *EntityArrayEmbedded) SetAuthorizationStatements(v []AuthorizeEditorDataStatementsReferenceableStatementDTO) {
	o.AuthorizationStatements = v
}

// GetAuthorizationVersions returns the AuthorizationVersions field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetAuthorizationVersions() []AuthorizeEditorDataAuthorizationVersionDTO {
	if o == nil || IsNil(o.AuthorizationVersions) {
		var ret []AuthorizeEditorDataAuthorizationVersionDTO
		return ret
	}
	return o.AuthorizationVersions
}

// GetAuthorizationVersionsOk returns a tuple with the AuthorizationVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetAuthorizationVersionsOk() ([]AuthorizeEditorDataAuthorizationVersionDTO, bool) {
	if o == nil || IsNil(o.AuthorizationVersions) {
		return nil, false
	}
	return o.AuthorizationVersions, true
}

// HasAuthorizationVersions returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasAuthorizationVersions() bool {
	if o != nil && !IsNil(o.AuthorizationVersions) {
		return true
	}

	return false
}

// SetAuthorizationVersions gets a reference to the given []AuthorizeEditorDataAuthorizationVersionDTO and assigns it to the AuthorizationVersions field.
func (o *EntityArrayEmbedded) SetAuthorizationVersions(v []AuthorizeEditorDataAuthorizationVersionDTO) {
	o.AuthorizationVersions = v
}

// GetDecisionEndpoints returns the DecisionEndpoints field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetDecisionEndpoints() []DecisionEndpoint {
	if o == nil || IsNil(o.DecisionEndpoints) {
		var ret []DecisionEndpoint
		return ret
	}
	return o.DecisionEndpoints
}

// GetDecisionEndpointsOk returns a tuple with the DecisionEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetDecisionEndpointsOk() ([]DecisionEndpoint, bool) {
	if o == nil || IsNil(o.DecisionEndpoints) {
		return nil, false
	}
	return o.DecisionEndpoints, true
}

// HasDecisionEndpoints returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasDecisionEndpoints() bool {
	if o != nil && !IsNil(o.DecisionEndpoints) {
		return true
	}

	return false
}

// SetDecisionEndpoints gets a reference to the given []DecisionEndpoint and assigns it to the DecisionEndpoints field.
func (o *EntityArrayEmbedded) SetDecisionEndpoints(v []DecisionEndpoint) {
	o.DecisionEndpoints = v
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetOperations() []APIServerOperation {
	if o == nil || IsNil(o.Operations) {
		var ret []APIServerOperation
		return ret
	}
	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetOperationsOk() ([]APIServerOperation, bool) {
	if o == nil || IsNil(o.Operations) {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasOperations() bool {
	if o != nil && !IsNil(o.Operations) {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []APIServerOperation and assigns it to the Operations field.
func (o *EntityArrayEmbedded) SetOperations(v []APIServerOperation) {
	o.Operations = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetPermissions() []EntityArrayEmbeddedPermissionsInner {
	if o == nil || IsNil(o.Permissions) {
		var ret []EntityArrayEmbeddedPermissionsInner
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetPermissionsOk() ([]EntityArrayEmbeddedPermissionsInner, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []EntityArrayEmbeddedPermissionsInner and assigns it to the Permissions field.
func (o *EntityArrayEmbedded) SetPermissions(v []EntityArrayEmbeddedPermissionsInner) {
	o.Permissions = v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetResources() []ApplicationResource {
	if o == nil || IsNil(o.Resources) {
		var ret []ApplicationResource
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetResourcesOk() ([]ApplicationResource, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []ApplicationResource and assigns it to the Resources field.
func (o *EntityArrayEmbedded) SetResources(v []ApplicationResource) {
	o.Resources = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetRoles() []ApplicationRole {
	if o == nil || IsNil(o.Roles) {
		var ret []ApplicationRole
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetRolesOk() ([]ApplicationRole, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []ApplicationRole and assigns it to the Roles field.
func (o *EntityArrayEmbedded) SetRoles(v []ApplicationRole) {
	o.Roles = v
}

func (o EntityArrayEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntityArrayEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiServers) {
		toSerialize["apiServers"] = o.ApiServers
	}
	if !IsNil(o.Assignments) {
		toSerialize["assignments"] = o.Assignments
	}
	if !IsNil(o.AuthorizationAttributes) {
		toSerialize["authorizationAttributes"] = o.AuthorizationAttributes
	}
	if !IsNil(o.AuthorizationConditions) {
		toSerialize["authorizationConditions"] = o.AuthorizationConditions
	}
	if !IsNil(o.AuthorizationConnectorTemplates) {
		toSerialize["authorizationConnectorTemplates"] = o.AuthorizationConnectorTemplates
	}
	if !IsNil(o.AuthorizationPolicies) {
		toSerialize["authorizationPolicies"] = o.AuthorizationPolicies
	}
	if !IsNil(o.AuthorizationProcessors) {
		toSerialize["authorizationProcessors"] = o.AuthorizationProcessors
	}
	if !IsNil(o.AuthorizationRules) {
		toSerialize["authorizationRules"] = o.AuthorizationRules
	}
	if !IsNil(o.AuthorizationServices) {
		toSerialize["authorizationServices"] = o.AuthorizationServices
	}
	if !IsNil(o.AuthorizationStatements) {
		toSerialize["authorizationStatements"] = o.AuthorizationStatements
	}
	if !IsNil(o.AuthorizationVersions) {
		toSerialize["authorizationVersions"] = o.AuthorizationVersions
	}
	if !IsNil(o.DecisionEndpoints) {
		toSerialize["decisionEndpoints"] = o.DecisionEndpoints
	}
	if !IsNil(o.Operations) {
		toSerialize["operations"] = o.Operations
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableEntityArrayEmbedded struct {
	value *EntityArrayEmbedded
	isSet bool
}

func (v NullableEntityArrayEmbedded) Get() *EntityArrayEmbedded {
	return v.value
}

func (v *NullableEntityArrayEmbedded) Set(val *EntityArrayEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityArrayEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityArrayEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityArrayEmbedded(val *EntityArrayEmbedded) *NullableEntityArrayEmbedded {
	return &NullableEntityArrayEmbedded{value: val, isSet: true}
}

func (v NullableEntityArrayEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityArrayEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


