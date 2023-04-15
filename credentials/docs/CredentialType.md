# CredentialType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardDesignTemplate** | **string** | A string that specifies an SVG formatted image containing placeholders for the credential fields that need to be displayed in the image. | 
**CardType** | Pointer to **string** | A string that specifies the descriptor of the credential type. Can be non-identity types such as proof of employment or proof of insurance. | [optional] 
**CreatedAt** | Pointer to **string** | A string that specifies the date and time the credential type was created. | [optional] 
**Description** | Pointer to **string** | A string that specifies the description of the credential type. | [optional] 
**Environment** | Pointer to [**Environment**](Environment.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the identifier (UUID) associated with the credential type. | [optional] 
**Issuer** | Pointer to [**CredentialTypeIssuer**](CredentialTypeIssuer.md) |  | [optional] 
**IssuerName** | Pointer to **string** | v issuer name associated with the card, can differ from title. | [optional] 
**Metadata** | [**CredentialTypeMetaData**](CredentialTypeMetaData.md) |  | 
**Title** | **string** | A string that specifies the title of the credential. Verification sites are expected to be able to request the issued credential from the compatible wallet app using the title. | 
**UpdatedAt** | Pointer to **string** | A string that specifies the date and time the credential type was last updated; can be null. | [optional] 

## Methods

### NewCredentialType

`func NewCredentialType(cardDesignTemplate string, metadata CredentialTypeMetaData, title string, ) *CredentialType`

NewCredentialType instantiates a new CredentialType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialTypeWithDefaults

`func NewCredentialTypeWithDefaults() *CredentialType`

NewCredentialTypeWithDefaults instantiates a new CredentialType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardDesignTemplate

`func (o *CredentialType) GetCardDesignTemplate() string`

GetCardDesignTemplate returns the CardDesignTemplate field if non-nil, zero value otherwise.

### GetCardDesignTemplateOk

`func (o *CredentialType) GetCardDesignTemplateOk() (*string, bool)`

GetCardDesignTemplateOk returns a tuple with the CardDesignTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardDesignTemplate

`func (o *CredentialType) SetCardDesignTemplate(v string)`

SetCardDesignTemplate sets CardDesignTemplate field to given value.


### GetCardType

`func (o *CredentialType) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *CredentialType) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *CredentialType) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *CredentialType) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CredentialType) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CredentialType) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CredentialType) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CredentialType) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *CredentialType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CredentialType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CredentialType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CredentialType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *CredentialType) GetEnvironment() Environment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CredentialType) GetEnvironmentOk() (*Environment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CredentialType) SetEnvironment(v Environment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CredentialType) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *CredentialType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CredentialType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CredentialType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CredentialType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssuer

`func (o *CredentialType) GetIssuer() CredentialTypeIssuer`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CredentialType) GetIssuerOk() (*CredentialTypeIssuer, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CredentialType) SetIssuer(v CredentialTypeIssuer)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *CredentialType) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetIssuerName

`func (o *CredentialType) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *CredentialType) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *CredentialType) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.

### HasIssuerName

`func (o *CredentialType) HasIssuerName() bool`

HasIssuerName returns a boolean if a field has been set.

### GetMetadata

`func (o *CredentialType) GetMetadata() CredentialTypeMetaData`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CredentialType) GetMetadataOk() (*CredentialTypeMetaData, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CredentialType) SetMetadata(v CredentialTypeMetaData)`

SetMetadata sets Metadata field to given value.


### GetTitle

`func (o *CredentialType) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CredentialType) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CredentialType) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUpdatedAt

`func (o *CredentialType) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CredentialType) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CredentialType) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CredentialType) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


