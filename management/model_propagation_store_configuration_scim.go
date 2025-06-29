/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// checks if the PropagationStoreConfigurationSCIM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropagationStoreConfigurationSCIM{}

// PropagationStoreConfigurationSCIM struct for PropagationStoreConfigurationSCIM
type PropagationStoreConfigurationSCIM struct {
	AUTHENTICATION_METHOD EnumPropagationStoreTypeSCIMAuthenticationMethod `json:"AUTHENTICATION_METHOD"`
	// The authorization header type.
	AUTHORIZATION_TYPE string `json:"AUTHORIZATION_TYPE"`
	// The password for account authentication. Required when `AUTHENTICATION_METHOD` is `Basic Authentication`, otherwise optional.
	BASIC_AUTH_PASSWORD *string `json:"BASIC_AUTH_PASSWORD,omitempty"`
	// The user name for account authentication. Required when `AUTHENTICATION_METHOD` is `Basic Authentication`, otherwise optional.
	BASIC_AUTH_USER *string `json:"BASIC_AUTH_USER,omitempty"`
	// Whether or not users are allowed to be created.
	CREATE_USERS *bool `json:"CREATE_USERS,omitempty"`
	// Whether or not users are allowed to be disabled.
	DISABLE_USERS     *bool                                        `json:"DISABLE_USERS,omitempty"`
	GROUP_NAME_SOURCE *EnumPropagationStoreTypeSCIMGroupNameSource `json:"GROUP_NAME_SOURCE,omitempty"`
	// API endpoint path to the group entity.
	GROUPS_RESOURCE *string `json:"GROUPS_RESOURCE,omitempty"`
	// OAuth access token for account authentication. Required when `AUTHENTICATION_METHOD` is `OAuth 2 Bearer Token`, otherwise optional.
	OAUTH_ACCESS_TOKEN *string `json:"OAUTH_ACCESS_TOKEN,omitempty"`
	// OAuth client ID. Required when `AUTHENTICATION_METHOD` is `OAuth 2 Client Credentials`, otherwise optional.
	OAUTH_CLIENT_ID *string `json:"OAUTH_CLIENT_ID,omitempty"`
	// OAuth client secret. Required when `AUTHENTICATION_METHOD` is `OAuth 2 Client Credentials`, otherwise optional.
	OAUTH_CLIENT_SECRET *string `json:"OAUTH_CLIENT_SECRET,omitempty"`
	// OAuth token request endpoint. Required when `AUTHENTICATION_METHOD` is `OAuth 2 Bearer Token`, otherwise optional.
	OAUTH_TOKEN_REQUEST *string                                            `json:"OAUTH_TOKEN_REQUEST,omitempty"`
	REMOVE_ACTION       *EnumPropagationStoreTypeRemoveActionDisableDelete `json:"REMOVE_ACTION,omitempty"`
	// The SCIM URL.
	SCIM_URL string `json:"SCIM_URL"`
	// The SCIM version.
	SCIM_VERSION           string                                           `json:"SCIM_VERSION"`
	UNIQUE_USER_IDENTIFIER EnumPropagationStoreTypeSCIMUniqueUserIdentifier `json:"UNIQUE_USER_IDENTIFIER"`
	// Whether or not users are allowed to be updated.
	UPDATE_USERS *bool `json:"UPDATE_USERS,omitempty"`
	// A string that specifies a SCIM filter expression.
	USER_FILTER string `json:"USER_FILTER"`
	// API endpoint path to the user entity.
	USERS_RESOURCE string `json:"USERS_RESOURCE"`
}

// NewPropagationStoreConfigurationSCIM instantiates a new PropagationStoreConfigurationSCIM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropagationStoreConfigurationSCIM(aUTHENTICATIONMETHOD EnumPropagationStoreTypeSCIMAuthenticationMethod, aUTHORIZATIONTYPE string, sCIMURL string, sCIMVERSION string, uNIQUEUSERIDENTIFIER EnumPropagationStoreTypeSCIMUniqueUserIdentifier, uSERFILTER string, uSERSRESOURCE string) *PropagationStoreConfigurationSCIM {
	this := PropagationStoreConfigurationSCIM{}
	this.AUTHENTICATION_METHOD = aUTHENTICATIONMETHOD
	this.AUTHORIZATION_TYPE = aUTHORIZATIONTYPE
	this.SCIM_URL = sCIMURL
	this.SCIM_VERSION = sCIMVERSION
	this.UNIQUE_USER_IDENTIFIER = uNIQUEUSERIDENTIFIER
	this.USER_FILTER = uSERFILTER
	this.USERS_RESOURCE = uSERSRESOURCE
	return &this
}

// NewPropagationStoreConfigurationSCIMWithDefaults instantiates a new PropagationStoreConfigurationSCIM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropagationStoreConfigurationSCIMWithDefaults() *PropagationStoreConfigurationSCIM {
	this := PropagationStoreConfigurationSCIM{}
	return &this
}

// GetAUTHENTICATION_METHOD returns the AUTHENTICATION_METHOD field value
func (o *PropagationStoreConfigurationSCIM) GetAUTHENTICATION_METHOD() EnumPropagationStoreTypeSCIMAuthenticationMethod {
	if o == nil {
		var ret EnumPropagationStoreTypeSCIMAuthenticationMethod
		return ret
	}

	return o.AUTHENTICATION_METHOD
}

// GetAUTHENTICATION_METHODOk returns a tuple with the AUTHENTICATION_METHOD field value
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationSCIM) GetAUTHENTICATION_METHODOk() (*EnumPropagationStoreTypeSCIMAuthenticationMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AUTHENTICATION_METHOD, true
}

// SetAUTHENTICATION_METHOD sets field value
func (o *PropagationStoreConfigurationSCIM) SetAUTHENTICATION_METHOD(v EnumPropagationStoreTypeSCIMAuthenticationMethod) {
	o.AUTHENTICATION_METHOD = v
}

// GetAUTHORIZATION_TYPE returns the AUTHORIZATION_TYPE field value
func (o *PropagationStoreConfigurationSCIM) GetAUTHORIZATION_TYPE() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AUTHORIZATION_TYPE
}

// GetAUTHORIZATION_TYPEOk returns a tuple with the AUTHORIZATION_TYPE field value
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationSCIM) GetAUTHORIZATION_TYPEOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AUTHORIZATION_TYPE, true
}

// SetAUTHORIZATION_TYPE sets field value
func (o *PropagationStoreConfigurationSCIM) SetAUTHORIZATION_TYPE(v string) {
	o.AUTHORIZATION_TYPE = v
}

// GetBASIC_AUTH_PASSWORD returns the BASIC_AUTH_PASSWORD field value if set, zero value otherwise.
func (o *PropagationStoreConfigurationSCIM) GetBASIC_AUTH_PASSWORD() string {
	if o == nil || IsNil(o.BASIC_AUTH_PASSWORD) {
		var ret string
		return ret
	}
	return *o.BASIC_AUTH_PASSWORD
}

// GetBASIC_AUTH_PASSWORDOk returns a tuple with the BASIC_AUTH_PASSWORD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationSCIM) GetBASIC_AUTH_PASSWORDOk() (*string, bool) {
	if o == nil || IsNil(o.BASIC_AUTH_PASSWORD) {
		return nil, false
	}
	return o.BASIC_AUTH_PASSWORD, true
}

// HasBASIC_AUTH_PASSWORD returns a boolean if a field has been set.
func (o *PropagationStoreConfigurationSCIM) HasBASIC_AUTH_PASSWORD() bool {
	if o != nil && !IsNil(o.BASIC_AUTH_PASSWORD) {
		return true
	}

	return false
}

// SetBASIC_AUTH_PASSWORD gets a reference to the given string and assigns it to the BASIC_AUTH_PASSWORD field.
func (o *PropagationStoreConfigurationSCIM) SetBASIC_AUTH_PASSWORD(v string) {
	o.BASIC_AUTH_PASSWORD = &v
}

// GetBASIC_AUTH_USER returns the BASIC_AUTH_USER field value if set, zero value otherwise.
func (o *PropagationStoreConfigurationSCIM) GetBASIC_AUTH_USER() string {
	if o == nil || IsNil(o.BASIC_AUTH_USER) {
		var ret string
		return ret
	}
	return *o.BASIC_AUTH_USER
}

// GetBASIC_AUTH_USEROk returns a tuple with the BASIC_AUTH_USER field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationSCIM) GetBASIC_AUTH_USEROk() (*string, bool) {
	if o == nil || IsNil(o.BASIC_AUTH_USER) {
		return nil, false
	}
	return o.BASIC_AUTH_USER, true
}

// HasBASIC_AUTH_USER returns a boolean if a field has been set.
func (o *PropagationStoreConfigurationSCIM) HasBASIC_AUTH_USER() bool {
	if o != nil && !IsNil(o.BASIC_AUTH_USER) {
		return true
	}

	return false
}

// SetBASIC_AUTH_USER gets a reference to the given string and assigns it to the BASIC_AUTH_USER field.
func (o *PropagationStoreConfigurationSCIM) SetBASIC_AUTH_USER(v string) {
	o.BASIC_AUTH_USER = &v
}

// GetCREATE_USERS returns the CREATE_USERS field value if set, zero value otherwise.
func (o *PropagationStoreConfigurationSCIM) GetCREATE_USERS() bool {
	if o == nil || IsNil(o.CREATE_USERS) {
		var ret bool
		return ret
	}
	return *o.CREATE_USERS
}

// GetCREATE_USERSOk returns a tuple with the CREATE_USERS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationSCIM) GetCREATE_USERSOk() (*bool, bool) {
	if o == nil || IsNil(o.CREATE_USERS) {
		return nil, false
	}
	return o.CREATE_USERS, true
}

// HasCREATE_USERS returns a boolean if a field has been set.
func (o *PropagationStoreConfigurationSCIM) HasCREATE_USERS() bool {
	if o != nil && !IsNil(o.CREATE_USERS) {
		return true
	}

	return false
}

// SetCREATE_USERS gets a reference to the given bool and assigns it to the CREATE_USERS field.
func (o *PropagationStoreConfigurationSCIM) SetCREATE_USERS(v bool) {
	o.CREATE_USERS = &v
}

// GetDISABLE_USERS returns the DISABLE_USERS field value if set, zero value otherwise.
func (o *PropagationStoreConfigurationSCIM) GetDISABLE_USERS() bool {
	if o == nil || IsNil(o.DISABLE_USERS) {
		var ret bool
		return ret
	}
	return *o.DISABLE_USERS
}

// GetDISABLE_USERSOk returns a tuple with the DISABLE_USERS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationSCIM) GetDISABLE_USERSOk() (*bool, bool) {
	if o == nil || IsNil(o.DISABLE_USERS) {
		return nil, false
	}
	return o.DISABLE_USERS, true
}

// HasDISABLE_USERS returns a boolean if a field has been set.
func (o *PropagationStoreConfigurationSCIM) HasDISABLE_USERS() bool {
	if o != nil && !IsNil(o.DISABLE_USERS) {
		return true
	}

	return false
}

// SetDISABLE_USERS gets a reference to the given bool and assigns it to the DISABLE_USERS field.
func (o *PropagationStoreConfigurationSCIM) SetDISABLE_USERS(v bool) {
	o.DISABLE_USERS = &v
}

// GetGROUP_NAME_SOURCE returns the GROUP_NAME_SOURCE field value if set, zero value otherwise.
func (o *PropagationStoreConfigurationSCIM) GetGROUP_NAME_SOURCE() EnumPropagationStoreTypeSCIMGroupNameSource {
	if o == nil || IsNil(o.GROUP_NAME_SOURCE) {
		var ret EnumPropagationStoreTypeSCIMGroupNameSource
		return ret
	}
	return *o.GROUP_NAME_SOURCE
}

// GetGROUP_NAME_SOURCEOk returns a tuple with the GROUP_NAME_SOURCE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationSCIM) GetGROUP_NAME_SOURCEOk() (*EnumPropagationStoreTypeSCIMGroupNameSource, bool) {
	if o == nil || IsNil(o.GROUP_NAME_SOURCE) {
		return nil, false
	}
	return o.GROUP_NAME_SOURCE, true
}

// HasGROUP_NAME_SOURCE returns a boolean if a field has been set.
func (o *PropagationStoreConfigurationSCIM) HasGROUP_NAME_SOURCE() bool {
	if o != nil && !IsNil(o.GROUP_NAME_SOURCE) {
		return true
	}

	return false
}

// SetGROUP_NAME_SOURCE gets a reference to the given EnumPropagationStoreTypeSCIMGroupNameSource and assigns it to the GROUP_NAME_SOURCE field.
func (o *PropagationStoreConfigurationSCIM) SetGROUP_NAME_SOURCE(v EnumPropagationStoreTypeSCIMGroupNameSource) {
	o.GROUP_NAME_SOURCE = &v
}

// GetGROUPS_RESOURCE returns the GROUPS_RESOURCE field value if set, zero value otherwise.
func (o *PropagationStoreConfigurationSCIM) GetGROUPS_RESOURCE() string {
	if o == nil || IsNil(o.GROUPS_RESOURCE) {
		var ret string
		return ret
	}
	return *o.GROUPS_RESOURCE
}

// GetGROUPS_RESOURCEOk returns a tuple with the GROUPS_RESOURCE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationSCIM) GetGROUPS_RESOURCEOk() (*string, bool) {
	if o == nil || IsNil(o.GROUPS_RESOURCE) {
		return nil, false
	}
	return o.GROUPS_RESOURCE, true
}

// HasGROUPS_RESOURCE returns a boolean if a field has been set.
func (o *PropagationStoreConfigurationSCIM) HasGROUPS_RESOURCE() bool {
	if o != nil && !IsNil(o.GROUPS_RESOURCE) {
		return true
	}

	return false
}

// SetGROUPS_RESOURCE gets a reference to the given string and assigns it to the GROUPS_RESOURCE field.
func (o *PropagationStoreConfigurationSCIM) SetGROUPS_RESOURCE(v string) {
	o.GROUPS_RESOURCE = &v
}

// GetOAUTH_ACCESS_TOKEN returns the OAUTH_ACCESS_TOKEN field value if set, zero value otherwise.
func (o *PropagationStoreConfigurationSCIM) GetOAUTH_ACCESS_TOKEN() string {
	if o == nil || IsNil(o.OAUTH_ACCESS_TOKEN) {
		var ret string
		return ret
	}
	return *o.OAUTH_ACCESS_TOKEN
}

// GetOAUTH_ACCESS_TOKENOk returns a tuple with the OAUTH_ACCESS_TOKEN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationSCIM) GetOAUTH_ACCESS_TOKENOk() (*string, bool) {
	if o == nil || IsNil(o.OAUTH_ACCESS_TOKEN) {
		return nil, false
	}
	return o.OAUTH_ACCESS_TOKEN, true
}

// HasOAUTH_ACCESS_TOKEN returns a boolean if a field has been set.
func (o *PropagationStoreConfigurationSCIM) HasOAUTH_ACCESS_TOKEN() bool {
	if o != nil && !IsNil(o.OAUTH_ACCESS_TOKEN) {
		return true
	}

	return false
}

// SetOAUTH_ACCESS_TOKEN gets a reference to the given string and assigns it to the OAUTH_ACCESS_TOKEN field.
func (o *PropagationStoreConfigurationSCIM) SetOAUTH_ACCESS_TOKEN(v string) {
	o.OAUTH_ACCESS_TOKEN = &v
}

// GetOAUTH_CLIENT_ID returns the OAUTH_CLIENT_ID field value if set, zero value otherwise.
func (o *PropagationStoreConfigurationSCIM) GetOAUTH_CLIENT_ID() string {
	if o == nil || IsNil(o.OAUTH_CLIENT_ID) {
		var ret string
		return ret
	}
	return *o.OAUTH_CLIENT_ID
}

// GetOAUTH_CLIENT_IDOk returns a tuple with the OAUTH_CLIENT_ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationSCIM) GetOAUTH_CLIENT_IDOk() (*string, bool) {
	if o == nil || IsNil(o.OAUTH_CLIENT_ID) {
		return nil, false
	}
	return o.OAUTH_CLIENT_ID, true
}

// HasOAUTH_CLIENT_ID returns a boolean if a field has been set.
func (o *PropagationStoreConfigurationSCIM) HasOAUTH_CLIENT_ID() bool {
	if o != nil && !IsNil(o.OAUTH_CLIENT_ID) {
		return true
	}

	return false
}

// SetOAUTH_CLIENT_ID gets a reference to the given string and assigns it to the OAUTH_CLIENT_ID field.
func (o *PropagationStoreConfigurationSCIM) SetOAUTH_CLIENT_ID(v string) {
	o.OAUTH_CLIENT_ID = &v
}

// GetOAUTH_CLIENT_SECRET returns the OAUTH_CLIENT_SECRET field value if set, zero value otherwise.
func (o *PropagationStoreConfigurationSCIM) GetOAUTH_CLIENT_SECRET() string {
	if o == nil || IsNil(o.OAUTH_CLIENT_SECRET) {
		var ret string
		return ret
	}
	return *o.OAUTH_CLIENT_SECRET
}

// GetOAUTH_CLIENT_SECRETOk returns a tuple with the OAUTH_CLIENT_SECRET field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationSCIM) GetOAUTH_CLIENT_SECRETOk() (*string, bool) {
	if o == nil || IsNil(o.OAUTH_CLIENT_SECRET) {
		return nil, false
	}
	return o.OAUTH_CLIENT_SECRET, true
}

// HasOAUTH_CLIENT_SECRET returns a boolean if a field has been set.
func (o *PropagationStoreConfigurationSCIM) HasOAUTH_CLIENT_SECRET() bool {
	if o != nil && !IsNil(o.OAUTH_CLIENT_SECRET) {
		return true
	}

	return false
}

// SetOAUTH_CLIENT_SECRET gets a reference to the given string and assigns it to the OAUTH_CLIENT_SECRET field.
func (o *PropagationStoreConfigurationSCIM) SetOAUTH_CLIENT_SECRET(v string) {
	o.OAUTH_CLIENT_SECRET = &v
}

// GetOAUTH_TOKEN_REQUEST returns the OAUTH_TOKEN_REQUEST field value if set, zero value otherwise.
func (o *PropagationStoreConfigurationSCIM) GetOAUTH_TOKEN_REQUEST() string {
	if o == nil || IsNil(o.OAUTH_TOKEN_REQUEST) {
		var ret string
		return ret
	}
	return *o.OAUTH_TOKEN_REQUEST
}

// GetOAUTH_TOKEN_REQUESTOk returns a tuple with the OAUTH_TOKEN_REQUEST field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationSCIM) GetOAUTH_TOKEN_REQUESTOk() (*string, bool) {
	if o == nil || IsNil(o.OAUTH_TOKEN_REQUEST) {
		return nil, false
	}
	return o.OAUTH_TOKEN_REQUEST, true
}

// HasOAUTH_TOKEN_REQUEST returns a boolean if a field has been set.
func (o *PropagationStoreConfigurationSCIM) HasOAUTH_TOKEN_REQUEST() bool {
	if o != nil && !IsNil(o.OAUTH_TOKEN_REQUEST) {
		return true
	}

	return false
}

// SetOAUTH_TOKEN_REQUEST gets a reference to the given string and assigns it to the OAUTH_TOKEN_REQUEST field.
func (o *PropagationStoreConfigurationSCIM) SetOAUTH_TOKEN_REQUEST(v string) {
	o.OAUTH_TOKEN_REQUEST = &v
}

// GetREMOVE_ACTION returns the REMOVE_ACTION field value if set, zero value otherwise.
func (o *PropagationStoreConfigurationSCIM) GetREMOVE_ACTION() EnumPropagationStoreTypeRemoveActionDisableDelete {
	if o == nil || IsNil(o.REMOVE_ACTION) {
		var ret EnumPropagationStoreTypeRemoveActionDisableDelete
		return ret
	}
	return *o.REMOVE_ACTION
}

// GetREMOVE_ACTIONOk returns a tuple with the REMOVE_ACTION field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationSCIM) GetREMOVE_ACTIONOk() (*EnumPropagationStoreTypeRemoveActionDisableDelete, bool) {
	if o == nil || IsNil(o.REMOVE_ACTION) {
		return nil, false
	}
	return o.REMOVE_ACTION, true
}

// HasREMOVE_ACTION returns a boolean if a field has been set.
func (o *PropagationStoreConfigurationSCIM) HasREMOVE_ACTION() bool {
	if o != nil && !IsNil(o.REMOVE_ACTION) {
		return true
	}

	return false
}

// SetREMOVE_ACTION gets a reference to the given EnumPropagationStoreTypeRemoveActionDisableDelete and assigns it to the REMOVE_ACTION field.
func (o *PropagationStoreConfigurationSCIM) SetREMOVE_ACTION(v EnumPropagationStoreTypeRemoveActionDisableDelete) {
	o.REMOVE_ACTION = &v
}

// GetSCIM_URL returns the SCIM_URL field value
func (o *PropagationStoreConfigurationSCIM) GetSCIM_URL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SCIM_URL
}

// GetSCIM_URLOk returns a tuple with the SCIM_URL field value
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationSCIM) GetSCIM_URLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SCIM_URL, true
}

// SetSCIM_URL sets field value
func (o *PropagationStoreConfigurationSCIM) SetSCIM_URL(v string) {
	o.SCIM_URL = v
}

// GetSCIM_VERSION returns the SCIM_VERSION field value
func (o *PropagationStoreConfigurationSCIM) GetSCIM_VERSION() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SCIM_VERSION
}

// GetSCIM_VERSIONOk returns a tuple with the SCIM_VERSION field value
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationSCIM) GetSCIM_VERSIONOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SCIM_VERSION, true
}

// SetSCIM_VERSION sets field value
func (o *PropagationStoreConfigurationSCIM) SetSCIM_VERSION(v string) {
	o.SCIM_VERSION = v
}

// GetUNIQUE_USER_IDENTIFIER returns the UNIQUE_USER_IDENTIFIER field value
func (o *PropagationStoreConfigurationSCIM) GetUNIQUE_USER_IDENTIFIER() EnumPropagationStoreTypeSCIMUniqueUserIdentifier {
	if o == nil {
		var ret EnumPropagationStoreTypeSCIMUniqueUserIdentifier
		return ret
	}

	return o.UNIQUE_USER_IDENTIFIER
}

// GetUNIQUE_USER_IDENTIFIEROk returns a tuple with the UNIQUE_USER_IDENTIFIER field value
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationSCIM) GetUNIQUE_USER_IDENTIFIEROk() (*EnumPropagationStoreTypeSCIMUniqueUserIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UNIQUE_USER_IDENTIFIER, true
}

// SetUNIQUE_USER_IDENTIFIER sets field value
func (o *PropagationStoreConfigurationSCIM) SetUNIQUE_USER_IDENTIFIER(v EnumPropagationStoreTypeSCIMUniqueUserIdentifier) {
	o.UNIQUE_USER_IDENTIFIER = v
}

// GetUPDATE_USERS returns the UPDATE_USERS field value if set, zero value otherwise.
func (o *PropagationStoreConfigurationSCIM) GetUPDATE_USERS() bool {
	if o == nil || IsNil(o.UPDATE_USERS) {
		var ret bool
		return ret
	}
	return *o.UPDATE_USERS
}

// GetUPDATE_USERSOk returns a tuple with the UPDATE_USERS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationSCIM) GetUPDATE_USERSOk() (*bool, bool) {
	if o == nil || IsNil(o.UPDATE_USERS) {
		return nil, false
	}
	return o.UPDATE_USERS, true
}

// HasUPDATE_USERS returns a boolean if a field has been set.
func (o *PropagationStoreConfigurationSCIM) HasUPDATE_USERS() bool {
	if o != nil && !IsNil(o.UPDATE_USERS) {
		return true
	}

	return false
}

// SetUPDATE_USERS gets a reference to the given bool and assigns it to the UPDATE_USERS field.
func (o *PropagationStoreConfigurationSCIM) SetUPDATE_USERS(v bool) {
	o.UPDATE_USERS = &v
}

// GetUSER_FILTER returns the USER_FILTER field value
func (o *PropagationStoreConfigurationSCIM) GetUSER_FILTER() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.USER_FILTER
}

// GetUSER_FILTEROk returns a tuple with the USER_FILTER field value
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationSCIM) GetUSER_FILTEROk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.USER_FILTER, true
}

// SetUSER_FILTER sets field value
func (o *PropagationStoreConfigurationSCIM) SetUSER_FILTER(v string) {
	o.USER_FILTER = v
}

// GetUSERS_RESOURCE returns the USERS_RESOURCE field value
func (o *PropagationStoreConfigurationSCIM) GetUSERS_RESOURCE() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.USERS_RESOURCE
}

// GetUSERS_RESOURCEOk returns a tuple with the USERS_RESOURCE field value
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationSCIM) GetUSERS_RESOURCEOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.USERS_RESOURCE, true
}

// SetUSERS_RESOURCE sets field value
func (o *PropagationStoreConfigurationSCIM) SetUSERS_RESOURCE(v string) {
	o.USERS_RESOURCE = v
}

func (o PropagationStoreConfigurationSCIM) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropagationStoreConfigurationSCIM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["AUTHENTICATION_METHOD"] = o.AUTHENTICATION_METHOD
	toSerialize["AUTHORIZATION_TYPE"] = o.AUTHORIZATION_TYPE
	if !IsNil(o.BASIC_AUTH_PASSWORD) {
		toSerialize["BASIC_AUTH_PASSWORD"] = o.BASIC_AUTH_PASSWORD
	}
	if !IsNil(o.BASIC_AUTH_USER) {
		toSerialize["BASIC_AUTH_USER"] = o.BASIC_AUTH_USER
	}
	if !IsNil(o.CREATE_USERS) {
		toSerialize["CREATE_USERS"] = o.CREATE_USERS
	}
	if !IsNil(o.DISABLE_USERS) {
		toSerialize["DISABLE_USERS"] = o.DISABLE_USERS
	}
	if !IsNil(o.GROUP_NAME_SOURCE) {
		toSerialize["GROUP_NAME_SOURCE"] = o.GROUP_NAME_SOURCE
	}
	if !IsNil(o.GROUPS_RESOURCE) {
		toSerialize["GROUPS_RESOURCE"] = o.GROUPS_RESOURCE
	}
	if !IsNil(o.OAUTH_ACCESS_TOKEN) {
		toSerialize["OAUTH_ACCESS_TOKEN"] = o.OAUTH_ACCESS_TOKEN
	}
	if !IsNil(o.OAUTH_CLIENT_ID) {
		toSerialize["OAUTH_CLIENT_ID"] = o.OAUTH_CLIENT_ID
	}
	if !IsNil(o.OAUTH_CLIENT_SECRET) {
		toSerialize["OAUTH_CLIENT_SECRET"] = o.OAUTH_CLIENT_SECRET
	}
	if !IsNil(o.OAUTH_TOKEN_REQUEST) {
		toSerialize["OAUTH_TOKEN_REQUEST"] = o.OAUTH_TOKEN_REQUEST
	}
	if !IsNil(o.REMOVE_ACTION) {
		toSerialize["REMOVE_ACTION"] = o.REMOVE_ACTION
	}
	toSerialize["SCIM_URL"] = o.SCIM_URL
	toSerialize["SCIM_VERSION"] = o.SCIM_VERSION
	toSerialize["UNIQUE_USER_IDENTIFIER"] = o.UNIQUE_USER_IDENTIFIER
	if !IsNil(o.UPDATE_USERS) {
		toSerialize["UPDATE_USERS"] = o.UPDATE_USERS
	}
	toSerialize["USER_FILTER"] = o.USER_FILTER
	toSerialize["USERS_RESOURCE"] = o.USERS_RESOURCE
	return toSerialize, nil
}

type NullablePropagationStoreConfigurationSCIM struct {
	value *PropagationStoreConfigurationSCIM
	isSet bool
}

func (v NullablePropagationStoreConfigurationSCIM) Get() *PropagationStoreConfigurationSCIM {
	return v.value
}

func (v *NullablePropagationStoreConfigurationSCIM) Set(val *PropagationStoreConfigurationSCIM) {
	v.value = val
	v.isSet = true
}

func (v NullablePropagationStoreConfigurationSCIM) IsSet() bool {
	return v.isSet
}

func (v *NullablePropagationStoreConfigurationSCIM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropagationStoreConfigurationSCIM(val *PropagationStoreConfigurationSCIM) *NullablePropagationStoreConfigurationSCIM {
	return &NullablePropagationStoreConfigurationSCIM{value: val, isSet: true}
}

func (v NullablePropagationStoreConfigurationSCIM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropagationStoreConfigurationSCIM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
