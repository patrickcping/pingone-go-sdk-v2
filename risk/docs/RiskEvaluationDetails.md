# RiskEvaluationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddressReputation** | Pointer to [**RiskEvaluationDetailsIpAddressReputation**](RiskEvaluationDetailsIpAddressReputation.md) |  | [optional] 
**IpVelocityByUser** | Pointer to [**RiskEvaluationDetailsIpVelocityByUser**](RiskEvaluationDetailsIpVelocityByUser.md) |  | [optional] 
**UserVelocityByIp** | Pointer to [**RiskEvaluationDetailsUserVelocityByIp**](RiskEvaluationDetailsUserVelocityByIp.md) |  | [optional] 
**ImpossibleTravel** | Pointer to **bool** | A boolean that specifies whether the distance between the location of the user in their previous successful authentication and current authentication infers that the user had to travel at a speed greater than 1000 kilometers per hour. This condition is marked as fulfilled, only if: Location data is available for the current and previous IP address of the user. This is not the first transaction that the user has performed. The user’s previous successful transaction was performed less than 24 hours ago. The user moved a distance of at least 100 kilometers. Thus, even if the user moved very fast, but moved only a distance of 90 kilometers, the condition is not fulfilled. The user moved at a speed greater than 1000 kilometers per hour. | [optional] 
**EstimatedSpeed** | Pointer to **float32** | The calculated travel speed in units of kilometers per hour. | [optional] 
**AnonymousNetworkDetected** | Pointer to **bool** | A boolean that specifies whether the current authentication originated from an anonymous network (for example, proxy or VPN). | [optional] 
**Country** | Pointer to **string** | A string that specifies the country related to current transaction from the IP address. | [optional] 
**State** | Pointer to **string** | A string that specifies the state related to current transaction from the IP address. | [optional] 
**City** | Pointer to **string** | A string that specifies the city related to current transaction from the IP address. | [optional] 
**Longitude** | Pointer to **float32** | A double-precision floating point that specifies the longitude related to current transaction from the IP address. Values range from -90 to 90. | [optional] 
**Latitude** | Pointer to **float32** | A double-precision floating point that specifies the latitude related to current transaction from the IP address. Values range from -180 to 180. | [optional] 
**PreviousSuccessfulTransaction** | Pointer to [**RiskEvaluationDetailsPreviousSuccessfulTransaction**](RiskEvaluationDetailsPreviousSuccessfulTransaction.md) |  | [optional] 
**UserBasedRiskBehavior** | Pointer to [**RiskEvaluationDetailsUserBasedRiskBehavior**](RiskEvaluationDetailsUserBasedRiskBehavior.md) |  | [optional] 
**UserRiskBehavior** | Pointer to [**RiskEvaluationDetailsUserRiskBehavior**](RiskEvaluationDetailsUserRiskBehavior.md) |  | [optional] 

## Methods

### NewRiskEvaluationDetails

`func NewRiskEvaluationDetails() *RiskEvaluationDetails`

NewRiskEvaluationDetails instantiates a new RiskEvaluationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskEvaluationDetailsWithDefaults

`func NewRiskEvaluationDetailsWithDefaults() *RiskEvaluationDetails`

NewRiskEvaluationDetailsWithDefaults instantiates a new RiskEvaluationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddressReputation

`func (o *RiskEvaluationDetails) GetIpAddressReputation() RiskEvaluationDetailsIpAddressReputation`

GetIpAddressReputation returns the IpAddressReputation field if non-nil, zero value otherwise.

### GetIpAddressReputationOk

`func (o *RiskEvaluationDetails) GetIpAddressReputationOk() (*RiskEvaluationDetailsIpAddressReputation, bool)`

GetIpAddressReputationOk returns a tuple with the IpAddressReputation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressReputation

`func (o *RiskEvaluationDetails) SetIpAddressReputation(v RiskEvaluationDetailsIpAddressReputation)`

SetIpAddressReputation sets IpAddressReputation field to given value.

### HasIpAddressReputation

`func (o *RiskEvaluationDetails) HasIpAddressReputation() bool`

HasIpAddressReputation returns a boolean if a field has been set.

### GetIpVelocityByUser

`func (o *RiskEvaluationDetails) GetIpVelocityByUser() RiskEvaluationDetailsIpVelocityByUser`

GetIpVelocityByUser returns the IpVelocityByUser field if non-nil, zero value otherwise.

### GetIpVelocityByUserOk

`func (o *RiskEvaluationDetails) GetIpVelocityByUserOk() (*RiskEvaluationDetailsIpVelocityByUser, bool)`

GetIpVelocityByUserOk returns a tuple with the IpVelocityByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVelocityByUser

`func (o *RiskEvaluationDetails) SetIpVelocityByUser(v RiskEvaluationDetailsIpVelocityByUser)`

SetIpVelocityByUser sets IpVelocityByUser field to given value.

### HasIpVelocityByUser

`func (o *RiskEvaluationDetails) HasIpVelocityByUser() bool`

HasIpVelocityByUser returns a boolean if a field has been set.

### GetUserVelocityByIp

`func (o *RiskEvaluationDetails) GetUserVelocityByIp() RiskEvaluationDetailsUserVelocityByIp`

GetUserVelocityByIp returns the UserVelocityByIp field if non-nil, zero value otherwise.

### GetUserVelocityByIpOk

`func (o *RiskEvaluationDetails) GetUserVelocityByIpOk() (*RiskEvaluationDetailsUserVelocityByIp, bool)`

GetUserVelocityByIpOk returns a tuple with the UserVelocityByIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVelocityByIp

`func (o *RiskEvaluationDetails) SetUserVelocityByIp(v RiskEvaluationDetailsUserVelocityByIp)`

SetUserVelocityByIp sets UserVelocityByIp field to given value.

### HasUserVelocityByIp

`func (o *RiskEvaluationDetails) HasUserVelocityByIp() bool`

HasUserVelocityByIp returns a boolean if a field has been set.

### GetImpossibleTravel

`func (o *RiskEvaluationDetails) GetImpossibleTravel() bool`

GetImpossibleTravel returns the ImpossibleTravel field if non-nil, zero value otherwise.

### GetImpossibleTravelOk

`func (o *RiskEvaluationDetails) GetImpossibleTravelOk() (*bool, bool)`

GetImpossibleTravelOk returns a tuple with the ImpossibleTravel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpossibleTravel

`func (o *RiskEvaluationDetails) SetImpossibleTravel(v bool)`

SetImpossibleTravel sets ImpossibleTravel field to given value.

### HasImpossibleTravel

`func (o *RiskEvaluationDetails) HasImpossibleTravel() bool`

HasImpossibleTravel returns a boolean if a field has been set.

### GetEstimatedSpeed

`func (o *RiskEvaluationDetails) GetEstimatedSpeed() float32`

GetEstimatedSpeed returns the EstimatedSpeed field if non-nil, zero value otherwise.

### GetEstimatedSpeedOk

`func (o *RiskEvaluationDetails) GetEstimatedSpeedOk() (*float32, bool)`

GetEstimatedSpeedOk returns a tuple with the EstimatedSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedSpeed

`func (o *RiskEvaluationDetails) SetEstimatedSpeed(v float32)`

SetEstimatedSpeed sets EstimatedSpeed field to given value.

### HasEstimatedSpeed

`func (o *RiskEvaluationDetails) HasEstimatedSpeed() bool`

HasEstimatedSpeed returns a boolean if a field has been set.

### GetAnonymousNetworkDetected

`func (o *RiskEvaluationDetails) GetAnonymousNetworkDetected() bool`

GetAnonymousNetworkDetected returns the AnonymousNetworkDetected field if non-nil, zero value otherwise.

### GetAnonymousNetworkDetectedOk

`func (o *RiskEvaluationDetails) GetAnonymousNetworkDetectedOk() (*bool, bool)`

GetAnonymousNetworkDetectedOk returns a tuple with the AnonymousNetworkDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousNetworkDetected

`func (o *RiskEvaluationDetails) SetAnonymousNetworkDetected(v bool)`

SetAnonymousNetworkDetected sets AnonymousNetworkDetected field to given value.

### HasAnonymousNetworkDetected

`func (o *RiskEvaluationDetails) HasAnonymousNetworkDetected() bool`

HasAnonymousNetworkDetected returns a boolean if a field has been set.

### GetCountry

`func (o *RiskEvaluationDetails) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *RiskEvaluationDetails) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *RiskEvaluationDetails) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *RiskEvaluationDetails) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetState

`func (o *RiskEvaluationDetails) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RiskEvaluationDetails) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RiskEvaluationDetails) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *RiskEvaluationDetails) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCity

`func (o *RiskEvaluationDetails) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *RiskEvaluationDetails) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *RiskEvaluationDetails) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *RiskEvaluationDetails) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetLongitude

`func (o *RiskEvaluationDetails) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *RiskEvaluationDetails) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *RiskEvaluationDetails) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *RiskEvaluationDetails) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetLatitude

`func (o *RiskEvaluationDetails) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *RiskEvaluationDetails) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *RiskEvaluationDetails) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *RiskEvaluationDetails) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetPreviousSuccessfulTransaction

`func (o *RiskEvaluationDetails) GetPreviousSuccessfulTransaction() RiskEvaluationDetailsPreviousSuccessfulTransaction`

GetPreviousSuccessfulTransaction returns the PreviousSuccessfulTransaction field if non-nil, zero value otherwise.

### GetPreviousSuccessfulTransactionOk

`func (o *RiskEvaluationDetails) GetPreviousSuccessfulTransactionOk() (*RiskEvaluationDetailsPreviousSuccessfulTransaction, bool)`

GetPreviousSuccessfulTransactionOk returns a tuple with the PreviousSuccessfulTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousSuccessfulTransaction

`func (o *RiskEvaluationDetails) SetPreviousSuccessfulTransaction(v RiskEvaluationDetailsPreviousSuccessfulTransaction)`

SetPreviousSuccessfulTransaction sets PreviousSuccessfulTransaction field to given value.

### HasPreviousSuccessfulTransaction

`func (o *RiskEvaluationDetails) HasPreviousSuccessfulTransaction() bool`

HasPreviousSuccessfulTransaction returns a boolean if a field has been set.

### GetUserBasedRiskBehavior

`func (o *RiskEvaluationDetails) GetUserBasedRiskBehavior() RiskEvaluationDetailsUserBasedRiskBehavior`

GetUserBasedRiskBehavior returns the UserBasedRiskBehavior field if non-nil, zero value otherwise.

### GetUserBasedRiskBehaviorOk

`func (o *RiskEvaluationDetails) GetUserBasedRiskBehaviorOk() (*RiskEvaluationDetailsUserBasedRiskBehavior, bool)`

GetUserBasedRiskBehaviorOk returns a tuple with the UserBasedRiskBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBasedRiskBehavior

`func (o *RiskEvaluationDetails) SetUserBasedRiskBehavior(v RiskEvaluationDetailsUserBasedRiskBehavior)`

SetUserBasedRiskBehavior sets UserBasedRiskBehavior field to given value.

### HasUserBasedRiskBehavior

`func (o *RiskEvaluationDetails) HasUserBasedRiskBehavior() bool`

HasUserBasedRiskBehavior returns a boolean if a field has been set.

### GetUserRiskBehavior

`func (o *RiskEvaluationDetails) GetUserRiskBehavior() RiskEvaluationDetailsUserRiskBehavior`

GetUserRiskBehavior returns the UserRiskBehavior field if non-nil, zero value otherwise.

### GetUserRiskBehaviorOk

`func (o *RiskEvaluationDetails) GetUserRiskBehaviorOk() (*RiskEvaluationDetailsUserRiskBehavior, bool)`

GetUserRiskBehaviorOk returns a tuple with the UserRiskBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRiskBehavior

`func (o *RiskEvaluationDetails) SetUserRiskBehavior(v RiskEvaluationDetailsUserRiskBehavior)`

SetUserRiskBehavior sets UserRiskBehavior field to given value.

### HasUserRiskBehavior

`func (o *RiskEvaluationDetails) HasUserRiskBehavior() bool`

HasUserRiskBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


