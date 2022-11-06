# License

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedServices** | Pointer to [**LicenseAdvancedServices**](LicenseAdvancedServices.md) |  | [optional] 
**AssignedEnvironmentsCount** | Pointer to **int32** | A read-only integer that specifies the total number of environments associated with this license. | [optional] [readonly] 
**Authorize** | Pointer to [**LicenseAuthorize**](LicenseAuthorize.md) |  | [optional] 
**BeginsAt** | Pointer to **string** | The date and time this license begins. | [optional] [readonly] 
**Credentials** | Pointer to [**LicenseCredentials**](LicenseCredentials.md) |  | [optional] 
**Environments** | Pointer to [**LicenseEnvironments**](LicenseEnvironments.md) |  | [optional] 
**ExpiresAt** | Pointer to **string** | The date and time this license expires. TRIAL licenses stop access to PingOne services at expiration. All other licenses trigger an event to send a notification when the license expires but do not block services. | [optional] [readonly] 
**Fraud** | Pointer to [**LicenseFraud**](LicenseFraud.md) |  | [optional] 
**Gateways** | Pointer to [**LicenseGateways**](LicenseGateways.md) |  | [optional] 
**Id** | Pointer to **string** | A read-only string that specifies the license resource’s unique identifier. | [optional] [readonly] 
**Intelligence** | Pointer to [**LicenseIntelligence**](LicenseIntelligence.md) |  | [optional] 
**Mfa** | Pointer to [**LicenseMfa**](LicenseMfa.md) |  | [optional] 
**Name** | **string** | A string that specifies a descriptive name for the license. This is a required property in a license name update request. Valid characters consists of any Unicode letter, mark, numeric character, forward slash, dot, apostrophe, underscore, space, or hyphen. The maximum length of a name is 255 characters. | 
**Orchestrate** | Pointer to [**LicenseOrchestrate**](LicenseOrchestrate.md) |  | [optional] 
**Organization** | Pointer to [**ObjectOrganization**](ObjectOrganization.md) |  | [optional] 
**Package** | Pointer to **string** | A string that specifies the license template on which this license is based. | [optional] 
**ReplacesLicense** | Pointer to [**LicenseReplacesLicense**](LicenseReplacesLicense.md) |  | [optional] 
**ReplacedByLicense** | Pointer to [**LicenseReplacedByLicense**](LicenseReplacedByLicense.md) |  | [optional] 
**Status** | Pointer to [**EnumLicenseStatus**](EnumLicenseStatus.md) |  | [optional] 
**TerminatesAt** | Pointer to **string** | An optional attribute that designates the exact date and time when this license terminates access to PingOne services. This attribute can be added to any licensing package. | [optional] 
**Users** | Pointer to [**LicenseUsers**](LicenseUsers.md) |  | [optional] 
**Verify** | Pointer to [**LicenseVerify**](LicenseVerify.md) |  | [optional] 

## Methods

### NewLicense

`func NewLicense(name string, ) *License`

NewLicense instantiates a new License object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseWithDefaults

`func NewLicenseWithDefaults() *License`

NewLicenseWithDefaults instantiates a new License object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedServices

`func (o *License) GetAdvancedServices() LicenseAdvancedServices`

GetAdvancedServices returns the AdvancedServices field if non-nil, zero value otherwise.

### GetAdvancedServicesOk

`func (o *License) GetAdvancedServicesOk() (*LicenseAdvancedServices, bool)`

GetAdvancedServicesOk returns a tuple with the AdvancedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedServices

`func (o *License) SetAdvancedServices(v LicenseAdvancedServices)`

SetAdvancedServices sets AdvancedServices field to given value.

### HasAdvancedServices

`func (o *License) HasAdvancedServices() bool`

HasAdvancedServices returns a boolean if a field has been set.

### GetAssignedEnvironmentsCount

`func (o *License) GetAssignedEnvironmentsCount() int32`

GetAssignedEnvironmentsCount returns the AssignedEnvironmentsCount field if non-nil, zero value otherwise.

### GetAssignedEnvironmentsCountOk

`func (o *License) GetAssignedEnvironmentsCountOk() (*int32, bool)`

GetAssignedEnvironmentsCountOk returns a tuple with the AssignedEnvironmentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedEnvironmentsCount

`func (o *License) SetAssignedEnvironmentsCount(v int32)`

SetAssignedEnvironmentsCount sets AssignedEnvironmentsCount field to given value.

### HasAssignedEnvironmentsCount

`func (o *License) HasAssignedEnvironmentsCount() bool`

HasAssignedEnvironmentsCount returns a boolean if a field has been set.

### GetAuthorize

`func (o *License) GetAuthorize() LicenseAuthorize`

GetAuthorize returns the Authorize field if non-nil, zero value otherwise.

### GetAuthorizeOk

`func (o *License) GetAuthorizeOk() (*LicenseAuthorize, bool)`

GetAuthorizeOk returns a tuple with the Authorize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorize

`func (o *License) SetAuthorize(v LicenseAuthorize)`

SetAuthorize sets Authorize field to given value.

### HasAuthorize

`func (o *License) HasAuthorize() bool`

HasAuthorize returns a boolean if a field has been set.

### GetBeginsAt

`func (o *License) GetBeginsAt() string`

GetBeginsAt returns the BeginsAt field if non-nil, zero value otherwise.

### GetBeginsAtOk

`func (o *License) GetBeginsAtOk() (*string, bool)`

GetBeginsAtOk returns a tuple with the BeginsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginsAt

`func (o *License) SetBeginsAt(v string)`

SetBeginsAt sets BeginsAt field to given value.

### HasBeginsAt

`func (o *License) HasBeginsAt() bool`

HasBeginsAt returns a boolean if a field has been set.

### GetCredentials

`func (o *License) GetCredentials() LicenseCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *License) GetCredentialsOk() (*LicenseCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *License) SetCredentials(v LicenseCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *License) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetEnvironments

`func (o *License) GetEnvironments() LicenseEnvironments`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *License) GetEnvironmentsOk() (*LicenseEnvironments, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *License) SetEnvironments(v LicenseEnvironments)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *License) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetExpiresAt

`func (o *License) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *License) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *License) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *License) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFraud

`func (o *License) GetFraud() LicenseFraud`

GetFraud returns the Fraud field if non-nil, zero value otherwise.

### GetFraudOk

`func (o *License) GetFraudOk() (*LicenseFraud, bool)`

GetFraudOk returns a tuple with the Fraud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraud

`func (o *License) SetFraud(v LicenseFraud)`

SetFraud sets Fraud field to given value.

### HasFraud

`func (o *License) HasFraud() bool`

HasFraud returns a boolean if a field has been set.

### GetGateways

`func (o *License) GetGateways() LicenseGateways`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *License) GetGatewaysOk() (*LicenseGateways, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *License) SetGateways(v LicenseGateways)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *License) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetId

`func (o *License) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *License) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *License) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *License) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIntelligence

`func (o *License) GetIntelligence() LicenseIntelligence`

GetIntelligence returns the Intelligence field if non-nil, zero value otherwise.

### GetIntelligenceOk

`func (o *License) GetIntelligenceOk() (*LicenseIntelligence, bool)`

GetIntelligenceOk returns a tuple with the Intelligence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntelligence

`func (o *License) SetIntelligence(v LicenseIntelligence)`

SetIntelligence sets Intelligence field to given value.

### HasIntelligence

`func (o *License) HasIntelligence() bool`

HasIntelligence returns a boolean if a field has been set.

### GetMfa

`func (o *License) GetMfa() LicenseMfa`

GetMfa returns the Mfa field if non-nil, zero value otherwise.

### GetMfaOk

`func (o *License) GetMfaOk() (*LicenseMfa, bool)`

GetMfaOk returns a tuple with the Mfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfa

`func (o *License) SetMfa(v LicenseMfa)`

SetMfa sets Mfa field to given value.

### HasMfa

`func (o *License) HasMfa() bool`

HasMfa returns a boolean if a field has been set.

### GetName

`func (o *License) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *License) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *License) SetName(v string)`

SetName sets Name field to given value.


### GetOrchestrate

`func (o *License) GetOrchestrate() LicenseOrchestrate`

GetOrchestrate returns the Orchestrate field if non-nil, zero value otherwise.

### GetOrchestrateOk

`func (o *License) GetOrchestrateOk() (*LicenseOrchestrate, bool)`

GetOrchestrateOk returns a tuple with the Orchestrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestrate

`func (o *License) SetOrchestrate(v LicenseOrchestrate)`

SetOrchestrate sets Orchestrate field to given value.

### HasOrchestrate

`func (o *License) HasOrchestrate() bool`

HasOrchestrate returns a boolean if a field has been set.

### GetOrganization

`func (o *License) GetOrganization() ObjectOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *License) GetOrganizationOk() (*ObjectOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *License) SetOrganization(v ObjectOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *License) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPackage

`func (o *License) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *License) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *License) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *License) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetReplacesLicense

`func (o *License) GetReplacesLicense() LicenseReplacesLicense`

GetReplacesLicense returns the ReplacesLicense field if non-nil, zero value otherwise.

### GetReplacesLicenseOk

`func (o *License) GetReplacesLicenseOk() (*LicenseReplacesLicense, bool)`

GetReplacesLicenseOk returns a tuple with the ReplacesLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacesLicense

`func (o *License) SetReplacesLicense(v LicenseReplacesLicense)`

SetReplacesLicense sets ReplacesLicense field to given value.

### HasReplacesLicense

`func (o *License) HasReplacesLicense() bool`

HasReplacesLicense returns a boolean if a field has been set.

### GetReplacedByLicense

`func (o *License) GetReplacedByLicense() LicenseReplacedByLicense`

GetReplacedByLicense returns the ReplacedByLicense field if non-nil, zero value otherwise.

### GetReplacedByLicenseOk

`func (o *License) GetReplacedByLicenseOk() (*LicenseReplacedByLicense, bool)`

GetReplacedByLicenseOk returns a tuple with the ReplacedByLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedByLicense

`func (o *License) SetReplacedByLicense(v LicenseReplacedByLicense)`

SetReplacedByLicense sets ReplacedByLicense field to given value.

### HasReplacedByLicense

`func (o *License) HasReplacedByLicense() bool`

HasReplacedByLicense returns a boolean if a field has been set.

### GetStatus

`func (o *License) GetStatus() EnumLicenseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *License) GetStatusOk() (*EnumLicenseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *License) SetStatus(v EnumLicenseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *License) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTerminatesAt

`func (o *License) GetTerminatesAt() string`

GetTerminatesAt returns the TerminatesAt field if non-nil, zero value otherwise.

### GetTerminatesAtOk

`func (o *License) GetTerminatesAtOk() (*string, bool)`

GetTerminatesAtOk returns a tuple with the TerminatesAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatesAt

`func (o *License) SetTerminatesAt(v string)`

SetTerminatesAt sets TerminatesAt field to given value.

### HasTerminatesAt

`func (o *License) HasTerminatesAt() bool`

HasTerminatesAt returns a boolean if a field has been set.

### GetUsers

`func (o *License) GetUsers() LicenseUsers`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *License) GetUsersOk() (*LicenseUsers, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *License) SetUsers(v LicenseUsers)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *License) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetVerify

`func (o *License) GetVerify() LicenseVerify`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *License) GetVerifyOk() (*LicenseVerify, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *License) SetVerify(v LicenseVerify)`

SetVerify sets Verify field to given value.

### HasVerify

`func (o *License) HasVerify() bool`

HasVerify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


