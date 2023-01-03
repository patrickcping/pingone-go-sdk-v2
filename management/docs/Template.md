# Template

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The template id | [optional] [readonly] 
**DisplayName** | **string** | The template’s display name. | 
**DeliveryMethods** | **[]string** | The delivery methods supported for this template. Valid values are &#x60;SMS&#x60;, &#x60;Voice&#x60;, &#x60;Email&#x60; and &#x60;Push&#x60;. | 
**CreatedAt** | Pointer to **string** | The time the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | The time the resource was last updated. | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the template. | [optional] 
**Variables** | **map[string]interface{}** | Lists the variables you can use in each template content. The &#x60;required&#x60; property indicates whether the variable is required in template content. If &#x60;required&#x60; is &#x60;true&#x60;, the &#x60;requiredForDeliveryMethods&#x60; property lists the &#x60;deliveryMethods&#x60; types that require the variable. Note that if &#x60;required&#x60; is &#x60;true&#x60;, but &#x60;requiredForDeliveryMethods&#x60; is not returned, all &#x60;deliveryMethods&#x60; types are required. For more information, see [Variables](https://apidocs.pingidentity.com/pingone/platform/v1/api/#notifications-templates-variables). | 
**AllowDynamicVariables** | Pointer to **bool** | Specifies whether dynamic variables can be used in the template&#39;s contents. For more information, see [Dynamic variables](https://apidocs.pingidentity.com/pingone/platform/v1/api/#notifications-templates-dynamic-variables). | [optional] 

## Methods

### NewTemplate

`func NewTemplate(displayName string, deliveryMethods []string, variables map[string]interface{}, ) *Template`

NewTemplate instantiates a new Template object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateWithDefaults

`func NewTemplateWithDefaults() *Template`

NewTemplateWithDefaults instantiates a new Template object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Template) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Template) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Template) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Template) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *Template) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Template) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Template) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDeliveryMethods

`func (o *Template) GetDeliveryMethods() []string`

GetDeliveryMethods returns the DeliveryMethods field if non-nil, zero value otherwise.

### GetDeliveryMethodsOk

`func (o *Template) GetDeliveryMethodsOk() (*[]string, bool)`

GetDeliveryMethodsOk returns a tuple with the DeliveryMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethods

`func (o *Template) SetDeliveryMethods(v []string)`

SetDeliveryMethods sets DeliveryMethods field to given value.


### GetCreatedAt

`func (o *Template) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Template) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Template) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Template) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Template) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Template) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Template) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Template) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *Template) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Template) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Template) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Template) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVariables

`func (o *Template) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *Template) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *Template) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.


### GetAllowDynamicVariables

`func (o *Template) GetAllowDynamicVariables() bool`

GetAllowDynamicVariables returns the AllowDynamicVariables field if non-nil, zero value otherwise.

### GetAllowDynamicVariablesOk

`func (o *Template) GetAllowDynamicVariablesOk() (*bool, bool)`

GetAllowDynamicVariablesOk returns a tuple with the AllowDynamicVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDynamicVariables

`func (o *Template) SetAllowDynamicVariables(v bool)`

SetAllowDynamicVariables sets AllowDynamicVariables field to given value.

### HasAllowDynamicVariables

`func (o *Template) HasAllowDynamicVariables() bool`

HasAllowDynamicVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


