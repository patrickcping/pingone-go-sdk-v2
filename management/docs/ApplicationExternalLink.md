# ApplicationExternalLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**AccessControl** | Pointer to [**ApplicationAccessControl**](ApplicationAccessControl.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the resource was created. | [optional] [readonly] 
**Description** | Pointer to **string** | A string that specifies the description of the application. | [optional] 
**Enabled** | **bool** | A string that specifies the current enabled state of the application. Options are ENABLED or DISABLED. | 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**HiddenFromAppPortal** | Pointer to **bool** | A boolean to specify whether the application is hidden in the application portal despite the configured group access policy. | [optional] 
**Icon** | Pointer to [**ApplicationIcon**](ApplicationIcon.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the application ID. | [optional] [readonly] 
**LoginPageUrl** | Pointer to **string** | A string that specifies the custom login page URL for the application. If you set the loginPageUrl property for applications in an environment that sets a custom domain, the URL should include the top-level domain and at least one additional domain level. Warning To avoid issues with third-party cookies in some browsers, a custom domain must be used, giving your PingOne environment the same parent domain as your authentication application. For more information about custom domains, see Custom domains. | [optional] 
**Name** | **string** | A string that specifies the name of the application. This is a required property. | 
**Protocol** | [**EnumApplicationProtocol**](EnumApplicationProtocol.md) |  | 
**Type** | [**EnumApplicationType**](EnumApplicationType.md) |  | 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated. | [optional] [readonly] 
**HomePageUrl** | **string** | A string that specifies the custom home page URL for the application. | 

## Methods

### NewApplicationExternalLink

`func NewApplicationExternalLink(enabled bool, name string, protocol EnumApplicationProtocol, type_ EnumApplicationType, homePageUrl string, ) *ApplicationExternalLink`

NewApplicationExternalLink instantiates a new ApplicationExternalLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationExternalLinkWithDefaults

`func NewApplicationExternalLinkWithDefaults() *ApplicationExternalLink`

NewApplicationExternalLinkWithDefaults instantiates a new ApplicationExternalLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ApplicationExternalLink) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApplicationExternalLink) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApplicationExternalLink) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApplicationExternalLink) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAccessControl

`func (o *ApplicationExternalLink) GetAccessControl() ApplicationAccessControl`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *ApplicationExternalLink) GetAccessControlOk() (*ApplicationAccessControl, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *ApplicationExternalLink) SetAccessControl(v ApplicationAccessControl)`

SetAccessControl sets AccessControl field to given value.

### HasAccessControl

`func (o *ApplicationExternalLink) HasAccessControl() bool`

HasAccessControl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApplicationExternalLink) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApplicationExternalLink) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApplicationExternalLink) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApplicationExternalLink) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationExternalLink) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationExternalLink) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationExternalLink) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationExternalLink) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ApplicationExternalLink) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApplicationExternalLink) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApplicationExternalLink) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnvironment

`func (o *ApplicationExternalLink) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ApplicationExternalLink) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ApplicationExternalLink) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ApplicationExternalLink) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetHiddenFromAppPortal

`func (o *ApplicationExternalLink) GetHiddenFromAppPortal() bool`

GetHiddenFromAppPortal returns the HiddenFromAppPortal field if non-nil, zero value otherwise.

### GetHiddenFromAppPortalOk

`func (o *ApplicationExternalLink) GetHiddenFromAppPortalOk() (*bool, bool)`

GetHiddenFromAppPortalOk returns a tuple with the HiddenFromAppPortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFromAppPortal

`func (o *ApplicationExternalLink) SetHiddenFromAppPortal(v bool)`

SetHiddenFromAppPortal sets HiddenFromAppPortal field to given value.

### HasHiddenFromAppPortal

`func (o *ApplicationExternalLink) HasHiddenFromAppPortal() bool`

HasHiddenFromAppPortal returns a boolean if a field has been set.

### GetIcon

`func (o *ApplicationExternalLink) GetIcon() ApplicationIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ApplicationExternalLink) GetIconOk() (*ApplicationIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ApplicationExternalLink) SetIcon(v ApplicationIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ApplicationExternalLink) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetId

`func (o *ApplicationExternalLink) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationExternalLink) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationExternalLink) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationExternalLink) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoginPageUrl

`func (o *ApplicationExternalLink) GetLoginPageUrl() string`

GetLoginPageUrl returns the LoginPageUrl field if non-nil, zero value otherwise.

### GetLoginPageUrlOk

`func (o *ApplicationExternalLink) GetLoginPageUrlOk() (*string, bool)`

GetLoginPageUrlOk returns a tuple with the LoginPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginPageUrl

`func (o *ApplicationExternalLink) SetLoginPageUrl(v string)`

SetLoginPageUrl sets LoginPageUrl field to given value.

### HasLoginPageUrl

`func (o *ApplicationExternalLink) HasLoginPageUrl() bool`

HasLoginPageUrl returns a boolean if a field has been set.

### GetName

`func (o *ApplicationExternalLink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationExternalLink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationExternalLink) SetName(v string)`

SetName sets Name field to given value.


### GetProtocol

`func (o *ApplicationExternalLink) GetProtocol() EnumApplicationProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ApplicationExternalLink) GetProtocolOk() (*EnumApplicationProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ApplicationExternalLink) SetProtocol(v EnumApplicationProtocol)`

SetProtocol sets Protocol field to given value.


### GetType

`func (o *ApplicationExternalLink) GetType() EnumApplicationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationExternalLink) GetTypeOk() (*EnumApplicationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationExternalLink) SetType(v EnumApplicationType)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *ApplicationExternalLink) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ApplicationExternalLink) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ApplicationExternalLink) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ApplicationExternalLink) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetHomePageUrl

`func (o *ApplicationExternalLink) GetHomePageUrl() string`

GetHomePageUrl returns the HomePageUrl field if non-nil, zero value otherwise.

### GetHomePageUrlOk

`func (o *ApplicationExternalLink) GetHomePageUrlOk() (*string, bool)`

GetHomePageUrlOk returns a tuple with the HomePageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePageUrl

`func (o *ApplicationExternalLink) SetHomePageUrl(v string)`

SetHomePageUrl sets HomePageUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


