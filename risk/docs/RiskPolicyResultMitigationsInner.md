# RiskPolicyResultMitigationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**EnumMitigationAction**](EnumMitigationAction.md) |  | 
**CustomAction** | Pointer to **string** | If you set the &#x60;action&#x60; parameter to &#x60;CUSTOM&#x60;, use &#x60;customAction&#x60; to specify the custom action that you want to recommend. | [optional] 
**MfaAuthenticationPolicyId** | Pointer to **string** | If you set the &#x60;action&#x60; parameter to &#x60;MFA&#x60;, use &#x60;mfaAuthenticationPolicyId&#x60; to specify the ID of the MFA policy that should be used for authentication flows. | [optional] 
**MfaRegistrationPolicyId** | Pointer to **string** | If you set the &#x60;action&#x60; parameter to &#x60;MFA&#x60;, use &#x60;mfaRegistrationPolicyId&#x60; to specify the ID of the MFA policy that should be used for registration flows. | [optional] 
**VerifyPolicyId** | Pointer to **string** | If you set the &#x60;action&#x60; parameter to &#x60;VERIFY&#x60;, use &#x60;verifyPolicyId&#x60; to specify the ID of the Verify policy that should be used. | [optional] 

## Methods

### NewRiskPolicyResultMitigationsInner

`func NewRiskPolicyResultMitigationsInner(action EnumMitigationAction, ) *RiskPolicyResultMitigationsInner`

NewRiskPolicyResultMitigationsInner instantiates a new RiskPolicyResultMitigationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPolicyResultMitigationsInnerWithDefaults

`func NewRiskPolicyResultMitigationsInnerWithDefaults() *RiskPolicyResultMitigationsInner`

NewRiskPolicyResultMitigationsInnerWithDefaults instantiates a new RiskPolicyResultMitigationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *RiskPolicyResultMitigationsInner) GetAction() EnumMitigationAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RiskPolicyResultMitigationsInner) GetActionOk() (*EnumMitigationAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RiskPolicyResultMitigationsInner) SetAction(v EnumMitigationAction)`

SetAction sets Action field to given value.


### GetCustomAction

`func (o *RiskPolicyResultMitigationsInner) GetCustomAction() string`

GetCustomAction returns the CustomAction field if non-nil, zero value otherwise.

### GetCustomActionOk

`func (o *RiskPolicyResultMitigationsInner) GetCustomActionOk() (*string, bool)`

GetCustomActionOk returns a tuple with the CustomAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAction

`func (o *RiskPolicyResultMitigationsInner) SetCustomAction(v string)`

SetCustomAction sets CustomAction field to given value.

### HasCustomAction

`func (o *RiskPolicyResultMitigationsInner) HasCustomAction() bool`

HasCustomAction returns a boolean if a field has been set.

### GetMfaAuthenticationPolicyId

`func (o *RiskPolicyResultMitigationsInner) GetMfaAuthenticationPolicyId() string`

GetMfaAuthenticationPolicyId returns the MfaAuthenticationPolicyId field if non-nil, zero value otherwise.

### GetMfaAuthenticationPolicyIdOk

`func (o *RiskPolicyResultMitigationsInner) GetMfaAuthenticationPolicyIdOk() (*string, bool)`

GetMfaAuthenticationPolicyIdOk returns a tuple with the MfaAuthenticationPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaAuthenticationPolicyId

`func (o *RiskPolicyResultMitigationsInner) SetMfaAuthenticationPolicyId(v string)`

SetMfaAuthenticationPolicyId sets MfaAuthenticationPolicyId field to given value.

### HasMfaAuthenticationPolicyId

`func (o *RiskPolicyResultMitigationsInner) HasMfaAuthenticationPolicyId() bool`

HasMfaAuthenticationPolicyId returns a boolean if a field has been set.

### GetMfaRegistrationPolicyId

`func (o *RiskPolicyResultMitigationsInner) GetMfaRegistrationPolicyId() string`

GetMfaRegistrationPolicyId returns the MfaRegistrationPolicyId field if non-nil, zero value otherwise.

### GetMfaRegistrationPolicyIdOk

`func (o *RiskPolicyResultMitigationsInner) GetMfaRegistrationPolicyIdOk() (*string, bool)`

GetMfaRegistrationPolicyIdOk returns a tuple with the MfaRegistrationPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaRegistrationPolicyId

`func (o *RiskPolicyResultMitigationsInner) SetMfaRegistrationPolicyId(v string)`

SetMfaRegistrationPolicyId sets MfaRegistrationPolicyId field to given value.

### HasMfaRegistrationPolicyId

`func (o *RiskPolicyResultMitigationsInner) HasMfaRegistrationPolicyId() bool`

HasMfaRegistrationPolicyId returns a boolean if a field has been set.

### GetVerifyPolicyId

`func (o *RiskPolicyResultMitigationsInner) GetVerifyPolicyId() string`

GetVerifyPolicyId returns the VerifyPolicyId field if non-nil, zero value otherwise.

### GetVerifyPolicyIdOk

`func (o *RiskPolicyResultMitigationsInner) GetVerifyPolicyIdOk() (*string, bool)`

GetVerifyPolicyIdOk returns a tuple with the VerifyPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyPolicyId

`func (o *RiskPolicyResultMitigationsInner) SetVerifyPolicyId(v string)`

SetVerifyPolicyId sets VerifyPolicyId field to given value.

### HasVerifyPolicyId

`func (o *RiskPolicyResultMitigationsInner) HasVerifyPolicyId() bool`

HasVerifyPolicyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


