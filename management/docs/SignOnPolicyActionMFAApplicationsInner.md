# SignOnPolicyActionMFAApplicationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**AutoEnrollment** | Pointer to [**SignOnPolicyActionMFAApplicationsInnerAutoEnrollment**](SignOnPolicyActionMFAApplicationsInnerAutoEnrollment.md) |  | [optional] 
**DeviceAuthorization** | Pointer to [**SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization**](SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization.md) |  | [optional] 

## Methods

### NewSignOnPolicyActionMFAApplicationsInner

`func NewSignOnPolicyActionMFAApplicationsInner(id string, ) *SignOnPolicyActionMFAApplicationsInner`

NewSignOnPolicyActionMFAApplicationsInner instantiates a new SignOnPolicyActionMFAApplicationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionMFAApplicationsInnerWithDefaults

`func NewSignOnPolicyActionMFAApplicationsInnerWithDefaults() *SignOnPolicyActionMFAApplicationsInner`

NewSignOnPolicyActionMFAApplicationsInnerWithDefaults instantiates a new SignOnPolicyActionMFAApplicationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SignOnPolicyActionMFAApplicationsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignOnPolicyActionMFAApplicationsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignOnPolicyActionMFAApplicationsInner) SetId(v string)`

SetId sets Id field to given value.


### GetAutoEnrollment

`func (o *SignOnPolicyActionMFAApplicationsInner) GetAutoEnrollment() SignOnPolicyActionMFAApplicationsInnerAutoEnrollment`

GetAutoEnrollment returns the AutoEnrollment field if non-nil, zero value otherwise.

### GetAutoEnrollmentOk

`func (o *SignOnPolicyActionMFAApplicationsInner) GetAutoEnrollmentOk() (*SignOnPolicyActionMFAApplicationsInnerAutoEnrollment, bool)`

GetAutoEnrollmentOk returns a tuple with the AutoEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoEnrollment

`func (o *SignOnPolicyActionMFAApplicationsInner) SetAutoEnrollment(v SignOnPolicyActionMFAApplicationsInnerAutoEnrollment)`

SetAutoEnrollment sets AutoEnrollment field to given value.

### HasAutoEnrollment

`func (o *SignOnPolicyActionMFAApplicationsInner) HasAutoEnrollment() bool`

HasAutoEnrollment returns a boolean if a field has been set.

### GetDeviceAuthorization

`func (o *SignOnPolicyActionMFAApplicationsInner) GetDeviceAuthorization() SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization`

GetDeviceAuthorization returns the DeviceAuthorization field if non-nil, zero value otherwise.

### GetDeviceAuthorizationOk

`func (o *SignOnPolicyActionMFAApplicationsInner) GetDeviceAuthorizationOk() (*SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization, bool)`

GetDeviceAuthorizationOk returns a tuple with the DeviceAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAuthorization

`func (o *SignOnPolicyActionMFAApplicationsInner) SetDeviceAuthorization(v SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization)`

SetDeviceAuthorization sets DeviceAuthorization field to given value.

### HasDeviceAuthorization

`func (o *SignOnPolicyActionMFAApplicationsInner) HasDeviceAuthorization() bool`

HasDeviceAuthorization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


