# CredentialTypeExpiration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**After** | Pointer to [**CredentialTypeExpirationAfter**](CredentialTypeExpirationAfter.md) |  | [optional] 
**Expression** | Pointer to **string** | PingOne Expression Language (PEL) expression that evaluates to an ISO 8601 date. For more information on PEL, see [PingOne expression language](https://docs.pingidentity.com/pingone/pingone_expression_language/p1_expression_language.html?_gl&#x3D;1*5lup6b*_gcl_au*MTU0Njc1NDk2Ny4xNzQ5MTEzMjI3*_ga*MTk5NDI3NjE3Ny4xNzQ5MTEzMjI3*_ga_V94KKVLZPT*czE3NTQ0NjU5OTQkbzEzNCRnMSR0MTc1NDQ3Mzg1OSRqNjAkbDAkaDA.). | [optional] 
**FieldName** | Pointer to **string** | On issuance, name of the field in the credential to hold the &#x60;expiration&#x60; that, when evaluated, returns an expiration date. Must be unique from all fields defined in the &#x60;metadata&#x60; object of [Create Credential Type](https://apidocs.pingidentity.com/pingone/platform/v1/api/#credentialing-metadata-object-data-model). Required when &#x60;expiration.type&#x60; is &#x60;SOFT&#x60;. Optional when &#x60;expiration.type&#x60; is &#x60;HARD&#x60; and ignored if used. | [optional] 
**Timestamp** | Pointer to **time.Time** | The date and time to expire in ISO 8601 YYYY-MM-DDTHH:MM:SS.sssZ format (milliseconds optional). | [optional] 
**Type** | [**EnumCredentialTypeExpirationType**](EnumCredentialTypeExpirationType.md) |  | 

## Methods

### NewCredentialTypeExpiration

`func NewCredentialTypeExpiration(type_ EnumCredentialTypeExpirationType, ) *CredentialTypeExpiration`

NewCredentialTypeExpiration instantiates a new CredentialTypeExpiration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialTypeExpirationWithDefaults

`func NewCredentialTypeExpirationWithDefaults() *CredentialTypeExpiration`

NewCredentialTypeExpirationWithDefaults instantiates a new CredentialTypeExpiration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfter

`func (o *CredentialTypeExpiration) GetAfter() CredentialTypeExpirationAfter`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *CredentialTypeExpiration) GetAfterOk() (*CredentialTypeExpirationAfter, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *CredentialTypeExpiration) SetAfter(v CredentialTypeExpirationAfter)`

SetAfter sets After field to given value.

### HasAfter

`func (o *CredentialTypeExpiration) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetExpression

`func (o *CredentialTypeExpiration) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *CredentialTypeExpiration) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *CredentialTypeExpiration) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *CredentialTypeExpiration) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetFieldName

`func (o *CredentialTypeExpiration) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *CredentialTypeExpiration) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *CredentialTypeExpiration) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *CredentialTypeExpiration) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetTimestamp

`func (o *CredentialTypeExpiration) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CredentialTypeExpiration) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CredentialTypeExpiration) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CredentialTypeExpiration) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *CredentialTypeExpiration) GetType() EnumCredentialTypeExpirationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CredentialTypeExpiration) GetTypeOk() (*EnumCredentialTypeExpirationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CredentialTypeExpiration) SetType(v EnumCredentialTypeExpirationType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


