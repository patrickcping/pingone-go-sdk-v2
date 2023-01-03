# ApplicationWSFEDAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudienceRestriction** | Pointer to **string** | The service provider ID. Defaults to &#x60;urn:federation:MicrosoftOnline&#x60;. | [optional] [default to "urn:federation:MicrosoftOnline"]
**DomainName** | **string** | The federated domain name (for example, the Azure custom domain). | 
**IdpSigning** | [**ApplicationWSFEDAllOfIdpSigning**](ApplicationWSFEDAllOfIdpSigning.md) |  | 
**Kerberos** | Pointer to [**ApplicationWSFEDAllOfKerberos**](ApplicationWSFEDAllOfKerberos.md) |  | [optional] 
**ReplyUrl** | **string** | The URL that the replying party (such as, Office365) uses to accept submissions of RequestSecurityTokenResponse messages that are a result of SSO requests. | 
**SloEndpoint** | Pointer to **string** | The single logout endpoint URL. | [optional] 

## Methods

### NewApplicationWSFEDAllOf

`func NewApplicationWSFEDAllOf(domainName string, idpSigning ApplicationWSFEDAllOfIdpSigning, replyUrl string, ) *ApplicationWSFEDAllOf`

NewApplicationWSFEDAllOf instantiates a new ApplicationWSFEDAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationWSFEDAllOfWithDefaults

`func NewApplicationWSFEDAllOfWithDefaults() *ApplicationWSFEDAllOf`

NewApplicationWSFEDAllOfWithDefaults instantiates a new ApplicationWSFEDAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudienceRestriction

`func (o *ApplicationWSFEDAllOf) GetAudienceRestriction() string`

GetAudienceRestriction returns the AudienceRestriction field if non-nil, zero value otherwise.

### GetAudienceRestrictionOk

`func (o *ApplicationWSFEDAllOf) GetAudienceRestrictionOk() (*string, bool)`

GetAudienceRestrictionOk returns a tuple with the AudienceRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceRestriction

`func (o *ApplicationWSFEDAllOf) SetAudienceRestriction(v string)`

SetAudienceRestriction sets AudienceRestriction field to given value.

### HasAudienceRestriction

`func (o *ApplicationWSFEDAllOf) HasAudienceRestriction() bool`

HasAudienceRestriction returns a boolean if a field has been set.

### GetDomainName

`func (o *ApplicationWSFEDAllOf) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ApplicationWSFEDAllOf) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ApplicationWSFEDAllOf) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### GetIdpSigning

`func (o *ApplicationWSFEDAllOf) GetIdpSigning() ApplicationWSFEDAllOfIdpSigning`

GetIdpSigning returns the IdpSigning field if non-nil, zero value otherwise.

### GetIdpSigningOk

`func (o *ApplicationWSFEDAllOf) GetIdpSigningOk() (*ApplicationWSFEDAllOfIdpSigning, bool)`

GetIdpSigningOk returns a tuple with the IdpSigning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpSigning

`func (o *ApplicationWSFEDAllOf) SetIdpSigning(v ApplicationWSFEDAllOfIdpSigning)`

SetIdpSigning sets IdpSigning field to given value.


### GetKerberos

`func (o *ApplicationWSFEDAllOf) GetKerberos() ApplicationWSFEDAllOfKerberos`

GetKerberos returns the Kerberos field if non-nil, zero value otherwise.

### GetKerberosOk

`func (o *ApplicationWSFEDAllOf) GetKerberosOk() (*ApplicationWSFEDAllOfKerberos, bool)`

GetKerberosOk returns a tuple with the Kerberos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberos

`func (o *ApplicationWSFEDAllOf) SetKerberos(v ApplicationWSFEDAllOfKerberos)`

SetKerberos sets Kerberos field to given value.

### HasKerberos

`func (o *ApplicationWSFEDAllOf) HasKerberos() bool`

HasKerberos returns a boolean if a field has been set.

### GetReplyUrl

`func (o *ApplicationWSFEDAllOf) GetReplyUrl() string`

GetReplyUrl returns the ReplyUrl field if non-nil, zero value otherwise.

### GetReplyUrlOk

`func (o *ApplicationWSFEDAllOf) GetReplyUrlOk() (*string, bool)`

GetReplyUrlOk returns a tuple with the ReplyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyUrl

`func (o *ApplicationWSFEDAllOf) SetReplyUrl(v string)`

SetReplyUrl sets ReplyUrl field to given value.


### GetSloEndpoint

`func (o *ApplicationWSFEDAllOf) GetSloEndpoint() string`

GetSloEndpoint returns the SloEndpoint field if non-nil, zero value otherwise.

### GetSloEndpointOk

`func (o *ApplicationWSFEDAllOf) GetSloEndpointOk() (*string, bool)`

GetSloEndpointOk returns a tuple with the SloEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloEndpoint

`func (o *ApplicationWSFEDAllOf) SetSloEndpoint(v string)`

SetSloEndpoint sets SloEndpoint field to given value.

### HasSloEndpoint

`func (o *ApplicationWSFEDAllOf) HasSloEndpoint() bool`

HasSloEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


