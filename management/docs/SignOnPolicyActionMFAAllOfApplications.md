# SignOnPolicyActionMFAAllOfApplications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**AutoEnrollment** | Pointer to [**SignOnPolicyActionMFAAllOfAutoEnrollment**](SignOnPolicyActionMFAAllOfAutoEnrollment.md) |  | [optional] 
**DeviceAuthorization** | Pointer to [**SignOnPolicyActionMFAAllOfDeviceAuthorization**](SignOnPolicyActionMFAAllOfDeviceAuthorization.md) |  | [optional] 

## Methods

### NewSignOnPolicyActionMFAAllOfApplications

`func NewSignOnPolicyActionMFAAllOfApplications(id string, ) *SignOnPolicyActionMFAAllOfApplications`

NewSignOnPolicyActionMFAAllOfApplications instantiates a new SignOnPolicyActionMFAAllOfApplications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionMFAAllOfApplicationsWithDefaults

`func NewSignOnPolicyActionMFAAllOfApplicationsWithDefaults() *SignOnPolicyActionMFAAllOfApplications`

NewSignOnPolicyActionMFAAllOfApplicationsWithDefaults instantiates a new SignOnPolicyActionMFAAllOfApplications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SignOnPolicyActionMFAAllOfApplications) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignOnPolicyActionMFAAllOfApplications) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignOnPolicyActionMFAAllOfApplications) SetId(v string)`

SetId sets Id field to given value.


### GetAutoEnrollment

`func (o *SignOnPolicyActionMFAAllOfApplications) GetAutoEnrollment() SignOnPolicyActionMFAAllOfAutoEnrollment`

GetAutoEnrollment returns the AutoEnrollment field if non-nil, zero value otherwise.

### GetAutoEnrollmentOk

`func (o *SignOnPolicyActionMFAAllOfApplications) GetAutoEnrollmentOk() (*SignOnPolicyActionMFAAllOfAutoEnrollment, bool)`

GetAutoEnrollmentOk returns a tuple with the AutoEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoEnrollment

`func (o *SignOnPolicyActionMFAAllOfApplications) SetAutoEnrollment(v SignOnPolicyActionMFAAllOfAutoEnrollment)`

SetAutoEnrollment sets AutoEnrollment field to given value.

### HasAutoEnrollment

`func (o *SignOnPolicyActionMFAAllOfApplications) HasAutoEnrollment() bool`

HasAutoEnrollment returns a boolean if a field has been set.

### GetDeviceAuthorization

`func (o *SignOnPolicyActionMFAAllOfApplications) GetDeviceAuthorization() SignOnPolicyActionMFAAllOfDeviceAuthorization`

GetDeviceAuthorization returns the DeviceAuthorization field if non-nil, zero value otherwise.

### GetDeviceAuthorizationOk

`func (o *SignOnPolicyActionMFAAllOfApplications) GetDeviceAuthorizationOk() (*SignOnPolicyActionMFAAllOfDeviceAuthorization, bool)`

GetDeviceAuthorizationOk returns a tuple with the DeviceAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAuthorization

`func (o *SignOnPolicyActionMFAAllOfApplications) SetDeviceAuthorization(v SignOnPolicyActionMFAAllOfDeviceAuthorization)`

SetDeviceAuthorization sets DeviceAuthorization field to given value.

### HasDeviceAuthorization

`func (o *SignOnPolicyActionMFAAllOfApplications) HasDeviceAuthorization() bool`

HasDeviceAuthorization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


