# SignOnPolicyActionProgressiveProfilingAttributesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A string that specifies the name and path of the user profile attribute as defined in the user schema (for example, email or address.postalCode). This property is required. | 
**Required** | **bool** | A boolean that specifies whether the user is required to provide a value for the attribute. This property is required. | 

## Methods

### NewSignOnPolicyActionProgressiveProfilingAttributesInner

`func NewSignOnPolicyActionProgressiveProfilingAttributesInner(name string, required bool, ) *SignOnPolicyActionProgressiveProfilingAttributesInner`

NewSignOnPolicyActionProgressiveProfilingAttributesInner instantiates a new SignOnPolicyActionProgressiveProfilingAttributesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionProgressiveProfilingAttributesInnerWithDefaults

`func NewSignOnPolicyActionProgressiveProfilingAttributesInnerWithDefaults() *SignOnPolicyActionProgressiveProfilingAttributesInner`

NewSignOnPolicyActionProgressiveProfilingAttributesInnerWithDefaults instantiates a new SignOnPolicyActionProgressiveProfilingAttributesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SignOnPolicyActionProgressiveProfilingAttributesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SignOnPolicyActionProgressiveProfilingAttributesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SignOnPolicyActionProgressiveProfilingAttributesInner) SetName(v string)`

SetName sets Name field to given value.


### GetRequired

`func (o *SignOnPolicyActionProgressiveProfilingAttributesInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *SignOnPolicyActionProgressiveProfilingAttributesInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *SignOnPolicyActionProgressiveProfilingAttributesInner) SetRequired(v bool)`

SetRequired sets Required field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


