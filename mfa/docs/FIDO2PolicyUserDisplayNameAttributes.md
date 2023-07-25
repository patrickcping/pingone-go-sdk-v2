# FIDO2PolicyUserDisplayNameAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**[]FIDO2PolicyUserDisplayNameAttributesAttributesInner**](FIDO2PolicyUserDisplayNameAttributesAttributesInner.md) | List of strings associated with the users&#39;s account that can be displayed during registration and authentication. Each object in the array is a name:value pair where the first part is \&quot;name\&quot; and the second is the name of the user attribute, for example, &#x60;{\&quot;name\&quot;: \&quot;username\&quot;}&#x60;, &#x60;{\&quot;name\&quot;: \&quot;email\&quot;}&#x60;. If you want to use the \&quot;name\&quot; attribute for the user, you must also specify the \&quot;subAttributes\&quot;, which can be either the \&quot;given\&quot; and \&quot;family\&quot; user attributes or the \&quot;formatted\&quot; user attribute. For example, &#x60;{\&quot;name\&quot;: “name”, “subAttributes”:[{“name”: “given”}, {“name”: “family”}]}, {\&quot;name\&quot;: \&quot;email\&quot;}&#x60; or &#x60;{\&quot;name\&quot;: “name”, “subAttributes”:[{“name”: “formatted”}]}, {\&quot;name\&quot;: \&quot;email\&quot;}&#x60;. - The content of the list should reflect the preferred order. - If the first attribute is empty for the user, PingOne will continue through the list until a non-empty attribute is found. - You can specify any user attribute (including custom attributes) that meet the following criteria: attribute type must be String, validation cannot be set to enumerated values. - The array must contain the user attribute \&quot;username\&quot; - to ensure that there is at least one non-empty attribute. - You can have a maximum of six user attributes in the list.  | 

## Methods

### NewFIDO2PolicyUserDisplayNameAttributes

`func NewFIDO2PolicyUserDisplayNameAttributes(attributes []FIDO2PolicyUserDisplayNameAttributesAttributesInner, ) *FIDO2PolicyUserDisplayNameAttributes`

NewFIDO2PolicyUserDisplayNameAttributes instantiates a new FIDO2PolicyUserDisplayNameAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFIDO2PolicyUserDisplayNameAttributesWithDefaults

`func NewFIDO2PolicyUserDisplayNameAttributesWithDefaults() *FIDO2PolicyUserDisplayNameAttributes`

NewFIDO2PolicyUserDisplayNameAttributesWithDefaults instantiates a new FIDO2PolicyUserDisplayNameAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *FIDO2PolicyUserDisplayNameAttributes) GetAttributes() []FIDO2PolicyUserDisplayNameAttributesAttributesInner`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FIDO2PolicyUserDisplayNameAttributes) GetAttributesOk() (*[]FIDO2PolicyUserDisplayNameAttributesAttributesInner, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FIDO2PolicyUserDisplayNameAttributes) SetAttributes(v []FIDO2PolicyUserDisplayNameAttributesAttributesInner)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


