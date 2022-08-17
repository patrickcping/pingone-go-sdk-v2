# SignOnPolicyActionAgreementAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agreement** | [**SignOnPolicyActionAgreementAllOfAgreement**](SignOnPolicyActionAgreementAllOfAgreement.md) |  | 
**DisableDeclineOption** | Pointer to **bool** | When enabled, the &#x60;Do Not Accept&#x60; button will terminate the Flow and display an error message to the user. | [optional] 

## Methods

### NewSignOnPolicyActionAgreementAllOf

`func NewSignOnPolicyActionAgreementAllOf(agreement SignOnPolicyActionAgreementAllOfAgreement, ) *SignOnPolicyActionAgreementAllOf`

NewSignOnPolicyActionAgreementAllOf instantiates a new SignOnPolicyActionAgreementAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionAgreementAllOfWithDefaults

`func NewSignOnPolicyActionAgreementAllOfWithDefaults() *SignOnPolicyActionAgreementAllOf`

NewSignOnPolicyActionAgreementAllOfWithDefaults instantiates a new SignOnPolicyActionAgreementAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgreement

`func (o *SignOnPolicyActionAgreementAllOf) GetAgreement() SignOnPolicyActionAgreementAllOfAgreement`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *SignOnPolicyActionAgreementAllOf) GetAgreementOk() (*SignOnPolicyActionAgreementAllOfAgreement, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *SignOnPolicyActionAgreementAllOf) SetAgreement(v SignOnPolicyActionAgreementAllOfAgreement)`

SetAgreement sets Agreement field to given value.


### GetDisableDeclineOption

`func (o *SignOnPolicyActionAgreementAllOf) GetDisableDeclineOption() bool`

GetDisableDeclineOption returns the DisableDeclineOption field if non-nil, zero value otherwise.

### GetDisableDeclineOptionOk

`func (o *SignOnPolicyActionAgreementAllOf) GetDisableDeclineOptionOk() (*bool, bool)`

GetDisableDeclineOptionOk returns a tuple with the DisableDeclineOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDeclineOption

`func (o *SignOnPolicyActionAgreementAllOf) SetDisableDeclineOption(v bool)`

SetDisableDeclineOption sets DisableDeclineOption field to given value.

### HasDisableDeclineOption

`func (o *SignOnPolicyActionAgreementAllOf) HasDisableDeclineOption() bool`

HasDisableDeclineOption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


