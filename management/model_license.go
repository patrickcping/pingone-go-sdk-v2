/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"time"
)

// checks if the License type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &License{}

// License struct for License
type License struct {
	Links            *map[string]LinksHATEOASValue `json:"_links,omitempty"`
	AdvancedServices *LicenseAdvancedServices      `json:"advancedServices,omitempty"`
	// A read-only integer that specifies the total number of environments associated with this license.
	AssignedEnvironmentsCount *int32            `json:"assignedEnvironmentsCount,omitempty"`
	Authorize                 *LicenseAuthorize `json:"authorize,omitempty"`
	// The date and time this license begins.
	BeginsAt     *time.Time           `json:"beginsAt,omitempty"`
	Credentials  *LicenseCredentials  `json:"credentials,omitempty"`
	Environments *LicenseEnvironments `json:"environments,omitempty"`
	// The date and time this license expires. TRIAL licenses stop access to PingOne services at expiration. All other licenses trigger an event to send a notification when the license expires but do not block services.
	ExpiresAt *time.Time       `json:"expiresAt,omitempty"`
	Fraud     *LicenseFraud    `json:"fraud,omitempty"`
	Gateways  *LicenseGateways `json:"gateways,omitempty"`
	// A read-only string that specifies the license resource’s unique identifier.
	Id           *string              `json:"id,omitempty"`
	Intelligence *LicenseIntelligence `json:"intelligence,omitempty"`
	Mfa          *LicenseMfa          `json:"mfa,omitempty"`
	// A string that specifies a descriptive name for the license. This is a required property in a license name update request. Valid characters consists of any Unicode letter, mark, numeric character, forward slash, dot, apostrophe, underscore, space, or hyphen. The maximum length of a name is 255 characters.
	Name         string              `json:"name"`
	Orchestrate  *LicenseOrchestrate `json:"orchestrate,omitempty"`
	Organization *ObjectOrganization `json:"organization,omitempty"`
	// A string that specifies the license template on which this license is based.
	Package           *string                   `json:"package,omitempty"`
	ReplacesLicense   *LicenseReplacesLicense   `json:"replacesLicense,omitempty"`
	ReplacedByLicense *LicenseReplacedByLicense `json:"replacedByLicense,omitempty"`
	Status            *EnumLicenseStatus        `json:"status,omitempty"`
	// An optional attribute that designates the exact date and time when this license terminates access to PingOne services. This attribute can be added to any licensing package.
	TerminatesAt *time.Time     `json:"terminatesAt,omitempty"`
	Users        *LicenseUsers  `json:"users,omitempty"`
	Verify       *LicenseVerify `json:"verify,omitempty"`
}

// NewLicense instantiates a new License object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicense(name string) *License {
	this := License{}
	this.Name = name
	return &this
}

// NewLicenseWithDefaults instantiates a new License object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseWithDefaults() *License {
	this := License{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *License) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetLinksOk() (*map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *License) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *License) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = &v
}

// GetAdvancedServices returns the AdvancedServices field value if set, zero value otherwise.
func (o *License) GetAdvancedServices() LicenseAdvancedServices {
	if o == nil || IsNil(o.AdvancedServices) {
		var ret LicenseAdvancedServices
		return ret
	}
	return *o.AdvancedServices
}

// GetAdvancedServicesOk returns a tuple with the AdvancedServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetAdvancedServicesOk() (*LicenseAdvancedServices, bool) {
	if o == nil || IsNil(o.AdvancedServices) {
		return nil, false
	}
	return o.AdvancedServices, true
}

// HasAdvancedServices returns a boolean if a field has been set.
func (o *License) HasAdvancedServices() bool {
	if o != nil && !IsNil(o.AdvancedServices) {
		return true
	}

	return false
}

// SetAdvancedServices gets a reference to the given LicenseAdvancedServices and assigns it to the AdvancedServices field.
func (o *License) SetAdvancedServices(v LicenseAdvancedServices) {
	o.AdvancedServices = &v
}

// GetAssignedEnvironmentsCount returns the AssignedEnvironmentsCount field value if set, zero value otherwise.
func (o *License) GetAssignedEnvironmentsCount() int32 {
	if o == nil || IsNil(o.AssignedEnvironmentsCount) {
		var ret int32
		return ret
	}
	return *o.AssignedEnvironmentsCount
}

// GetAssignedEnvironmentsCountOk returns a tuple with the AssignedEnvironmentsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetAssignedEnvironmentsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AssignedEnvironmentsCount) {
		return nil, false
	}
	return o.AssignedEnvironmentsCount, true
}

// HasAssignedEnvironmentsCount returns a boolean if a field has been set.
func (o *License) HasAssignedEnvironmentsCount() bool {
	if o != nil && !IsNil(o.AssignedEnvironmentsCount) {
		return true
	}

	return false
}

// SetAssignedEnvironmentsCount gets a reference to the given int32 and assigns it to the AssignedEnvironmentsCount field.
func (o *License) SetAssignedEnvironmentsCount(v int32) {
	o.AssignedEnvironmentsCount = &v
}

// GetAuthorize returns the Authorize field value if set, zero value otherwise.
func (o *License) GetAuthorize() LicenseAuthorize {
	if o == nil || IsNil(o.Authorize) {
		var ret LicenseAuthorize
		return ret
	}
	return *o.Authorize
}

// GetAuthorizeOk returns a tuple with the Authorize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetAuthorizeOk() (*LicenseAuthorize, bool) {
	if o == nil || IsNil(o.Authorize) {
		return nil, false
	}
	return o.Authorize, true
}

// HasAuthorize returns a boolean if a field has been set.
func (o *License) HasAuthorize() bool {
	if o != nil && !IsNil(o.Authorize) {
		return true
	}

	return false
}

// SetAuthorize gets a reference to the given LicenseAuthorize and assigns it to the Authorize field.
func (o *License) SetAuthorize(v LicenseAuthorize) {
	o.Authorize = &v
}

// GetBeginsAt returns the BeginsAt field value if set, zero value otherwise.
func (o *License) GetBeginsAt() time.Time {
	if o == nil || IsNil(o.BeginsAt) {
		var ret time.Time
		return ret
	}
	return *o.BeginsAt
}

// GetBeginsAtOk returns a tuple with the BeginsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetBeginsAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.BeginsAt) {
		return nil, false
	}
	return o.BeginsAt, true
}

// HasBeginsAt returns a boolean if a field has been set.
func (o *License) HasBeginsAt() bool {
	if o != nil && !IsNil(o.BeginsAt) {
		return true
	}

	return false
}

// SetBeginsAt gets a reference to the given time.Time and assigns it to the BeginsAt field.
func (o *License) SetBeginsAt(v time.Time) {
	o.BeginsAt = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *License) GetCredentials() LicenseCredentials {
	if o == nil || IsNil(o.Credentials) {
		var ret LicenseCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetCredentialsOk() (*LicenseCredentials, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *License) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given LicenseCredentials and assigns it to the Credentials field.
func (o *License) SetCredentials(v LicenseCredentials) {
	o.Credentials = &v
}

// GetEnvironments returns the Environments field value if set, zero value otherwise.
func (o *License) GetEnvironments() LicenseEnvironments {
	if o == nil || IsNil(o.Environments) {
		var ret LicenseEnvironments
		return ret
	}
	return *o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetEnvironmentsOk() (*LicenseEnvironments, bool) {
	if o == nil || IsNil(o.Environments) {
		return nil, false
	}
	return o.Environments, true
}

// HasEnvironments returns a boolean if a field has been set.
func (o *License) HasEnvironments() bool {
	if o != nil && !IsNil(o.Environments) {
		return true
	}

	return false
}

// SetEnvironments gets a reference to the given LicenseEnvironments and assigns it to the Environments field.
func (o *License) SetEnvironments(v LicenseEnvironments) {
	o.Environments = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *License) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *License) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *License) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetFraud returns the Fraud field value if set, zero value otherwise.
func (o *License) GetFraud() LicenseFraud {
	if o == nil || IsNil(o.Fraud) {
		var ret LicenseFraud
		return ret
	}
	return *o.Fraud
}

// GetFraudOk returns a tuple with the Fraud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetFraudOk() (*LicenseFraud, bool) {
	if o == nil || IsNil(o.Fraud) {
		return nil, false
	}
	return o.Fraud, true
}

// HasFraud returns a boolean if a field has been set.
func (o *License) HasFraud() bool {
	if o != nil && !IsNil(o.Fraud) {
		return true
	}

	return false
}

// SetFraud gets a reference to the given LicenseFraud and assigns it to the Fraud field.
func (o *License) SetFraud(v LicenseFraud) {
	o.Fraud = &v
}

// GetGateways returns the Gateways field value if set, zero value otherwise.
func (o *License) GetGateways() LicenseGateways {
	if o == nil || IsNil(o.Gateways) {
		var ret LicenseGateways
		return ret
	}
	return *o.Gateways
}

// GetGatewaysOk returns a tuple with the Gateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetGatewaysOk() (*LicenseGateways, bool) {
	if o == nil || IsNil(o.Gateways) {
		return nil, false
	}
	return o.Gateways, true
}

// HasGateways returns a boolean if a field has been set.
func (o *License) HasGateways() bool {
	if o != nil && !IsNil(o.Gateways) {
		return true
	}

	return false
}

// SetGateways gets a reference to the given LicenseGateways and assigns it to the Gateways field.
func (o *License) SetGateways(v LicenseGateways) {
	o.Gateways = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *License) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *License) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *License) SetId(v string) {
	o.Id = &v
}

// GetIntelligence returns the Intelligence field value if set, zero value otherwise.
func (o *License) GetIntelligence() LicenseIntelligence {
	if o == nil || IsNil(o.Intelligence) {
		var ret LicenseIntelligence
		return ret
	}
	return *o.Intelligence
}

// GetIntelligenceOk returns a tuple with the Intelligence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetIntelligenceOk() (*LicenseIntelligence, bool) {
	if o == nil || IsNil(o.Intelligence) {
		return nil, false
	}
	return o.Intelligence, true
}

// HasIntelligence returns a boolean if a field has been set.
func (o *License) HasIntelligence() bool {
	if o != nil && !IsNil(o.Intelligence) {
		return true
	}

	return false
}

// SetIntelligence gets a reference to the given LicenseIntelligence and assigns it to the Intelligence field.
func (o *License) SetIntelligence(v LicenseIntelligence) {
	o.Intelligence = &v
}

// GetMfa returns the Mfa field value if set, zero value otherwise.
func (o *License) GetMfa() LicenseMfa {
	if o == nil || IsNil(o.Mfa) {
		var ret LicenseMfa
		return ret
	}
	return *o.Mfa
}

// GetMfaOk returns a tuple with the Mfa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetMfaOk() (*LicenseMfa, bool) {
	if o == nil || IsNil(o.Mfa) {
		return nil, false
	}
	return o.Mfa, true
}

// HasMfa returns a boolean if a field has been set.
func (o *License) HasMfa() bool {
	if o != nil && !IsNil(o.Mfa) {
		return true
	}

	return false
}

// SetMfa gets a reference to the given LicenseMfa and assigns it to the Mfa field.
func (o *License) SetMfa(v LicenseMfa) {
	o.Mfa = &v
}

// GetName returns the Name field value
func (o *License) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *License) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *License) SetName(v string) {
	o.Name = v
}

// GetOrchestrate returns the Orchestrate field value if set, zero value otherwise.
func (o *License) GetOrchestrate() LicenseOrchestrate {
	if o == nil || IsNil(o.Orchestrate) {
		var ret LicenseOrchestrate
		return ret
	}
	return *o.Orchestrate
}

// GetOrchestrateOk returns a tuple with the Orchestrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetOrchestrateOk() (*LicenseOrchestrate, bool) {
	if o == nil || IsNil(o.Orchestrate) {
		return nil, false
	}
	return o.Orchestrate, true
}

// HasOrchestrate returns a boolean if a field has been set.
func (o *License) HasOrchestrate() bool {
	if o != nil && !IsNil(o.Orchestrate) {
		return true
	}

	return false
}

// SetOrchestrate gets a reference to the given LicenseOrchestrate and assigns it to the Orchestrate field.
func (o *License) SetOrchestrate(v LicenseOrchestrate) {
	o.Orchestrate = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *License) GetOrganization() ObjectOrganization {
	if o == nil || IsNil(o.Organization) {
		var ret ObjectOrganization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetOrganizationOk() (*ObjectOrganization, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *License) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given ObjectOrganization and assigns it to the Organization field.
func (o *License) SetOrganization(v ObjectOrganization) {
	o.Organization = &v
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *License) GetPackage() string {
	if o == nil || IsNil(o.Package) {
		var ret string
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetPackageOk() (*string, bool) {
	if o == nil || IsNil(o.Package) {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *License) HasPackage() bool {
	if o != nil && !IsNil(o.Package) {
		return true
	}

	return false
}

// SetPackage gets a reference to the given string and assigns it to the Package field.
func (o *License) SetPackage(v string) {
	o.Package = &v
}

// GetReplacesLicense returns the ReplacesLicense field value if set, zero value otherwise.
func (o *License) GetReplacesLicense() LicenseReplacesLicense {
	if o == nil || IsNil(o.ReplacesLicense) {
		var ret LicenseReplacesLicense
		return ret
	}
	return *o.ReplacesLicense
}

// GetReplacesLicenseOk returns a tuple with the ReplacesLicense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetReplacesLicenseOk() (*LicenseReplacesLicense, bool) {
	if o == nil || IsNil(o.ReplacesLicense) {
		return nil, false
	}
	return o.ReplacesLicense, true
}

// HasReplacesLicense returns a boolean if a field has been set.
func (o *License) HasReplacesLicense() bool {
	if o != nil && !IsNil(o.ReplacesLicense) {
		return true
	}

	return false
}

// SetReplacesLicense gets a reference to the given LicenseReplacesLicense and assigns it to the ReplacesLicense field.
func (o *License) SetReplacesLicense(v LicenseReplacesLicense) {
	o.ReplacesLicense = &v
}

// GetReplacedByLicense returns the ReplacedByLicense field value if set, zero value otherwise.
func (o *License) GetReplacedByLicense() LicenseReplacedByLicense {
	if o == nil || IsNil(o.ReplacedByLicense) {
		var ret LicenseReplacedByLicense
		return ret
	}
	return *o.ReplacedByLicense
}

// GetReplacedByLicenseOk returns a tuple with the ReplacedByLicense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetReplacedByLicenseOk() (*LicenseReplacedByLicense, bool) {
	if o == nil || IsNil(o.ReplacedByLicense) {
		return nil, false
	}
	return o.ReplacedByLicense, true
}

// HasReplacedByLicense returns a boolean if a field has been set.
func (o *License) HasReplacedByLicense() bool {
	if o != nil && !IsNil(o.ReplacedByLicense) {
		return true
	}

	return false
}

// SetReplacedByLicense gets a reference to the given LicenseReplacedByLicense and assigns it to the ReplacedByLicense field.
func (o *License) SetReplacedByLicense(v LicenseReplacedByLicense) {
	o.ReplacedByLicense = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *License) GetStatus() EnumLicenseStatus {
	if o == nil || IsNil(o.Status) {
		var ret EnumLicenseStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetStatusOk() (*EnumLicenseStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *License) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given EnumLicenseStatus and assigns it to the Status field.
func (o *License) SetStatus(v EnumLicenseStatus) {
	o.Status = &v
}

// GetTerminatesAt returns the TerminatesAt field value if set, zero value otherwise.
func (o *License) GetTerminatesAt() time.Time {
	if o == nil || IsNil(o.TerminatesAt) {
		var ret time.Time
		return ret
	}
	return *o.TerminatesAt
}

// GetTerminatesAtOk returns a tuple with the TerminatesAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetTerminatesAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TerminatesAt) {
		return nil, false
	}
	return o.TerminatesAt, true
}

// HasTerminatesAt returns a boolean if a field has been set.
func (o *License) HasTerminatesAt() bool {
	if o != nil && !IsNil(o.TerminatesAt) {
		return true
	}

	return false
}

// SetTerminatesAt gets a reference to the given time.Time and assigns it to the TerminatesAt field.
func (o *License) SetTerminatesAt(v time.Time) {
	o.TerminatesAt = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *License) GetUsers() LicenseUsers {
	if o == nil || IsNil(o.Users) {
		var ret LicenseUsers
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetUsersOk() (*LicenseUsers, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *License) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given LicenseUsers and assigns it to the Users field.
func (o *License) SetUsers(v LicenseUsers) {
	o.Users = &v
}

// GetVerify returns the Verify field value if set, zero value otherwise.
func (o *License) GetVerify() LicenseVerify {
	if o == nil || IsNil(o.Verify) {
		var ret LicenseVerify
		return ret
	}
	return *o.Verify
}

// GetVerifyOk returns a tuple with the Verify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetVerifyOk() (*LicenseVerify, bool) {
	if o == nil || IsNil(o.Verify) {
		return nil, false
	}
	return o.Verify, true
}

// HasVerify returns a boolean if a field has been set.
func (o *License) HasVerify() bool {
	if o != nil && !IsNil(o.Verify) {
		return true
	}

	return false
}

// SetVerify gets a reference to the given LicenseVerify and assigns it to the Verify field.
func (o *License) SetVerify(v LicenseVerify) {
	o.Verify = &v
}

func (o License) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o License) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.AdvancedServices) {
		toSerialize["advancedServices"] = o.AdvancedServices
	}
	if !IsNil(o.AssignedEnvironmentsCount) {
		toSerialize["assignedEnvironmentsCount"] = o.AssignedEnvironmentsCount
	}
	if !IsNil(o.Authorize) {
		toSerialize["authorize"] = o.Authorize
	}
	if !IsNil(o.BeginsAt) {
		toSerialize["beginsAt"] = o.BeginsAt
	}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.Environments) {
		toSerialize["environments"] = o.Environments
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if !IsNil(o.Fraud) {
		toSerialize["fraud"] = o.Fraud
	}
	if !IsNil(o.Gateways) {
		toSerialize["gateways"] = o.Gateways
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Intelligence) {
		toSerialize["intelligence"] = o.Intelligence
	}
	if !IsNil(o.Mfa) {
		toSerialize["mfa"] = o.Mfa
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Orchestrate) {
		toSerialize["orchestrate"] = o.Orchestrate
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.Package) {
		toSerialize["package"] = o.Package
	}
	if !IsNil(o.ReplacesLicense) {
		toSerialize["replacesLicense"] = o.ReplacesLicense
	}
	if !IsNil(o.ReplacedByLicense) {
		toSerialize["replacedByLicense"] = o.ReplacedByLicense
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TerminatesAt) {
		toSerialize["terminatesAt"] = o.TerminatesAt
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !IsNil(o.Verify) {
		toSerialize["verify"] = o.Verify
	}
	return toSerialize, nil
}

type NullableLicense struct {
	value *License
	isSet bool
}

func (v NullableLicense) Get() *License {
	return v.value
}

func (v *NullableLicense) Set(val *License) {
	v.value = val
	v.isSet = true
}

func (v NullableLicense) IsSet() bool {
	return v.isSet
}

func (v *NullableLicense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicense(val *License) *NullableLicense {
	return &NullableLicense{value: val, isSet: true}
}

func (v NullableLicense) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
