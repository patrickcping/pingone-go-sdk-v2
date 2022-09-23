# APIServerOperationsValuePathsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pattern** | **string** | A string that specifies the pattern used to identify the path or paths for the operation. The semantics of the pattern are determined by the type. For any type, the pattern can contain characters that are otherwise invalid in a URL path. Invalid characters are handled by performing matching against a percent-decoded HTTP request target path. This allows an administrator to configure patterns without worrying about percent encoding special characters. When the type is PARAMETER, the syntax outlined in the table below is enforced. Additionally, the pattern must contain a wildcard, double wildcard or parameter capture. When the type is EXACT, the pattern can be any byte sequence except for ASCII control characters such as line feeds or carriage returns. The length of the pattern cannot exceed 2048 characters. The path pattern must not contain empty path segments such as /../, //, and /./. | 
**Type** | [**EnumAPIServerPatternType**](EnumAPIServerPatternType.md) |  | 

## Methods

### NewAPIServerOperationsValuePathsInner

`func NewAPIServerOperationsValuePathsInner(pattern string, type_ EnumAPIServerPatternType, ) *APIServerOperationsValuePathsInner`

NewAPIServerOperationsValuePathsInner instantiates a new APIServerOperationsValuePathsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIServerOperationsValuePathsInnerWithDefaults

`func NewAPIServerOperationsValuePathsInnerWithDefaults() *APIServerOperationsValuePathsInner`

NewAPIServerOperationsValuePathsInnerWithDefaults instantiates a new APIServerOperationsValuePathsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPattern

`func (o *APIServerOperationsValuePathsInner) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *APIServerOperationsValuePathsInner) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *APIServerOperationsValuePathsInner) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetType

`func (o *APIServerOperationsValuePathsInner) GetType() EnumAPIServerPatternType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *APIServerOperationsValuePathsInner) GetTypeOk() (*EnumAPIServerPatternType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *APIServerOperationsValuePathsInner) SetType(v EnumAPIServerPatternType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


