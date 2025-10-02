# EntityArrayEmbeddedItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**CardDesignTemplate** | **string** | A string that specifies an SVG formatted image containing placeholders for the credential fields that need to be displayed in the image. | 
**CardType** | Pointer to **string** | A string that specifies the descriptor of the credential type. Can be non-identity types such as proof of employment or proof of insurance. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Description** | Pointer to **string** | A string that specifies the description of the credential type. | [optional] 
**DeletedAt** | Pointer to **time.Time** | A string that specifies the date and time the credential type was deleted. Note - a deletion of a credential type is a \&quot;soft delete\&quot;. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Expiration** | Pointer to [**CredentialTypeExpiration**](CredentialTypeExpiration.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Issuer** | Pointer to [**ObjectIssuer**](ObjectIssuer.md) |  | [optional] 
**IssuerName** | Pointer to **string** | A string that specifies the issuer name associated with the card, can differ from title. | [optional] 
**Management** | Pointer to [**CredentialTypeManagement**](CredentialTypeManagement.md) |  | [optional] 
**Metadata** | [**CredentialTypeMetaData**](CredentialTypeMetaData.md) |  | 
**Multiple** | Pointer to [**CredentialTypeMultiple**](CredentialTypeMultiple.md) |  | [optional] 
**OnDelete** | Pointer to [**CredentialTypeOnDelete**](CredentialTypeOnDelete.md) |  | [optional] 
**Title** | **string** | A string that specifies the title of the credential. Verification sites are expected to be able to request the issued credential from the compatible wallet app using the title. | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Version** | Pointer to [**CredentialTypeVersion**](CredentialTypeVersion.md) |  | [optional] 
**CredentialType** | Pointer to [**CredentialDigitalWalletNotificationResultsInnerNotification**](CredentialDigitalWalletNotificationResultsInnerNotification.md) |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Notification** | Pointer to [**CredentialDigitalWalletNotification**](CredentialDigitalWalletNotification.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**CredentialDigitalWalletNotificationResultsInnerNotification**](CredentialDigitalWalletNotificationResultsInnerNotification.md) |  | [optional] 
**Userdata** | Pointer to [**UserCredentialUserdata**](UserCredentialUserdata.md) |  | [optional] 

## Methods

### NewEntityArrayEmbeddedItemsInner

`func NewEntityArrayEmbeddedItemsInner(cardDesignTemplate string, metadata CredentialTypeMetaData, title string, ) *EntityArrayEmbeddedItemsInner`

NewEntityArrayEmbeddedItemsInner instantiates a new EntityArrayEmbeddedItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityArrayEmbeddedItemsInnerWithDefaults

`func NewEntityArrayEmbeddedItemsInnerWithDefaults() *EntityArrayEmbeddedItemsInner`

NewEntityArrayEmbeddedItemsInnerWithDefaults instantiates a new EntityArrayEmbeddedItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *EntityArrayEmbeddedItemsInner) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EntityArrayEmbeddedItemsInner) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EntityArrayEmbeddedItemsInner) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EntityArrayEmbeddedItemsInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCardDesignTemplate

`func (o *EntityArrayEmbeddedItemsInner) GetCardDesignTemplate() string`

GetCardDesignTemplate returns the CardDesignTemplate field if non-nil, zero value otherwise.

### GetCardDesignTemplateOk

`func (o *EntityArrayEmbeddedItemsInner) GetCardDesignTemplateOk() (*string, bool)`

GetCardDesignTemplateOk returns a tuple with the CardDesignTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardDesignTemplate

`func (o *EntityArrayEmbeddedItemsInner) SetCardDesignTemplate(v string)`

SetCardDesignTemplate sets CardDesignTemplate field to given value.


### GetCardType

`func (o *EntityArrayEmbeddedItemsInner) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *EntityArrayEmbeddedItemsInner) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *EntityArrayEmbeddedItemsInner) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *EntityArrayEmbeddedItemsInner) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EntityArrayEmbeddedItemsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EntityArrayEmbeddedItemsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EntityArrayEmbeddedItemsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EntityArrayEmbeddedItemsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *EntityArrayEmbeddedItemsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntityArrayEmbeddedItemsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntityArrayEmbeddedItemsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntityArrayEmbeddedItemsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeletedAt

`func (o *EntityArrayEmbeddedItemsInner) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EntityArrayEmbeddedItemsInner) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EntityArrayEmbeddedItemsInner) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EntityArrayEmbeddedItemsInner) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetEnvironment

`func (o *EntityArrayEmbeddedItemsInner) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EntityArrayEmbeddedItemsInner) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EntityArrayEmbeddedItemsInner) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EntityArrayEmbeddedItemsInner) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetExpiration

`func (o *EntityArrayEmbeddedItemsInner) GetExpiration() CredentialTypeExpiration`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *EntityArrayEmbeddedItemsInner) GetExpirationOk() (*CredentialTypeExpiration, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *EntityArrayEmbeddedItemsInner) SetExpiration(v CredentialTypeExpiration)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *EntityArrayEmbeddedItemsInner) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetId

`func (o *EntityArrayEmbeddedItemsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityArrayEmbeddedItemsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityArrayEmbeddedItemsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntityArrayEmbeddedItemsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssuer

`func (o *EntityArrayEmbeddedItemsInner) GetIssuer() ObjectIssuer`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *EntityArrayEmbeddedItemsInner) GetIssuerOk() (*ObjectIssuer, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *EntityArrayEmbeddedItemsInner) SetIssuer(v ObjectIssuer)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *EntityArrayEmbeddedItemsInner) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetIssuerName

`func (o *EntityArrayEmbeddedItemsInner) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *EntityArrayEmbeddedItemsInner) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *EntityArrayEmbeddedItemsInner) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.

### HasIssuerName

`func (o *EntityArrayEmbeddedItemsInner) HasIssuerName() bool`

HasIssuerName returns a boolean if a field has been set.

### GetManagement

`func (o *EntityArrayEmbeddedItemsInner) GetManagement() CredentialTypeManagement`

GetManagement returns the Management field if non-nil, zero value otherwise.

### GetManagementOk

`func (o *EntityArrayEmbeddedItemsInner) GetManagementOk() (*CredentialTypeManagement, bool)`

GetManagementOk returns a tuple with the Management field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagement

`func (o *EntityArrayEmbeddedItemsInner) SetManagement(v CredentialTypeManagement)`

SetManagement sets Management field to given value.

### HasManagement

`func (o *EntityArrayEmbeddedItemsInner) HasManagement() bool`

HasManagement returns a boolean if a field has been set.

### GetMetadata

`func (o *EntityArrayEmbeddedItemsInner) GetMetadata() CredentialTypeMetaData`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EntityArrayEmbeddedItemsInner) GetMetadataOk() (*CredentialTypeMetaData, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EntityArrayEmbeddedItemsInner) SetMetadata(v CredentialTypeMetaData)`

SetMetadata sets Metadata field to given value.


### GetMultiple

`func (o *EntityArrayEmbeddedItemsInner) GetMultiple() CredentialTypeMultiple`

GetMultiple returns the Multiple field if non-nil, zero value otherwise.

### GetMultipleOk

`func (o *EntityArrayEmbeddedItemsInner) GetMultipleOk() (*CredentialTypeMultiple, bool)`

GetMultipleOk returns a tuple with the Multiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiple

`func (o *EntityArrayEmbeddedItemsInner) SetMultiple(v CredentialTypeMultiple)`

SetMultiple sets Multiple field to given value.

### HasMultiple

`func (o *EntityArrayEmbeddedItemsInner) HasMultiple() bool`

HasMultiple returns a boolean if a field has been set.

### GetOnDelete

`func (o *EntityArrayEmbeddedItemsInner) GetOnDelete() CredentialTypeOnDelete`

GetOnDelete returns the OnDelete field if non-nil, zero value otherwise.

### GetOnDeleteOk

`func (o *EntityArrayEmbeddedItemsInner) GetOnDeleteOk() (*CredentialTypeOnDelete, bool)`

GetOnDeleteOk returns a tuple with the OnDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDelete

`func (o *EntityArrayEmbeddedItemsInner) SetOnDelete(v CredentialTypeOnDelete)`

SetOnDelete sets OnDelete field to given value.

### HasOnDelete

`func (o *EntityArrayEmbeddedItemsInner) HasOnDelete() bool`

HasOnDelete returns a boolean if a field has been set.

### GetTitle

`func (o *EntityArrayEmbeddedItemsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EntityArrayEmbeddedItemsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EntityArrayEmbeddedItemsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUpdatedAt

`func (o *EntityArrayEmbeddedItemsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EntityArrayEmbeddedItemsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EntityArrayEmbeddedItemsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EntityArrayEmbeddedItemsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *EntityArrayEmbeddedItemsInner) GetVersion() CredentialTypeVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EntityArrayEmbeddedItemsInner) GetVersionOk() (*CredentialTypeVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EntityArrayEmbeddedItemsInner) SetVersion(v CredentialTypeVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EntityArrayEmbeddedItemsInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCredentialType

`func (o *EntityArrayEmbeddedItemsInner) GetCredentialType() CredentialDigitalWalletNotificationResultsInnerNotification`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *EntityArrayEmbeddedItemsInner) GetCredentialTypeOk() (*CredentialDigitalWalletNotificationResultsInnerNotification, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *EntityArrayEmbeddedItemsInner) SetCredentialType(v CredentialDigitalWalletNotificationResultsInnerNotification)`

SetCredentialType sets CredentialType field to given value.

### HasCredentialType

`func (o *EntityArrayEmbeddedItemsInner) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.

### GetExpiresAt

`func (o *EntityArrayEmbeddedItemsInner) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *EntityArrayEmbeddedItemsInner) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *EntityArrayEmbeddedItemsInner) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *EntityArrayEmbeddedItemsInner) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetNotification

`func (o *EntityArrayEmbeddedItemsInner) GetNotification() CredentialDigitalWalletNotification`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *EntityArrayEmbeddedItemsInner) GetNotificationOk() (*CredentialDigitalWalletNotification, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *EntityArrayEmbeddedItemsInner) SetNotification(v CredentialDigitalWalletNotification)`

SetNotification sets Notification field to given value.

### HasNotification

`func (o *EntityArrayEmbeddedItemsInner) HasNotification() bool`

HasNotification returns a boolean if a field has been set.

### GetStatus

`func (o *EntityArrayEmbeddedItemsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EntityArrayEmbeddedItemsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EntityArrayEmbeddedItemsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EntityArrayEmbeddedItemsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUser

`func (o *EntityArrayEmbeddedItemsInner) GetUser() CredentialDigitalWalletNotificationResultsInnerNotification`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EntityArrayEmbeddedItemsInner) GetUserOk() (*CredentialDigitalWalletNotificationResultsInnerNotification, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EntityArrayEmbeddedItemsInner) SetUser(v CredentialDigitalWalletNotificationResultsInnerNotification)`

SetUser sets User field to given value.

### HasUser

`func (o *EntityArrayEmbeddedItemsInner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserdata

`func (o *EntityArrayEmbeddedItemsInner) GetUserdata() UserCredentialUserdata`

GetUserdata returns the Userdata field if non-nil, zero value otherwise.

### GetUserdataOk

`func (o *EntityArrayEmbeddedItemsInner) GetUserdataOk() (*UserCredentialUserdata, bool)`

GetUserdataOk returns a tuple with the Userdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserdata

`func (o *EntityArrayEmbeddedItemsInner) SetUserdata(v UserCredentialUserdata)`

SetUserdata sets Userdata field to given value.

### HasUserdata

`func (o *EntityArrayEmbeddedItemsInner) HasUserdata() bool`

HasUserdata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


