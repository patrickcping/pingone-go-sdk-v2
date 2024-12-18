# RiskPredictor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Name** | **string** | A string type. A unique, friendly name for the predictor. This name is displayed in the Risk Policies UI, when the admin is asked to define the overrides and weights. | 
**CompactName** | **string** | A string type. A unique name for the predictor. This property is immutable; it cannot be modified after initial creation. The value must be alpha-numeric, with no special characters or spaces. This name is used in the API both for policy configuration, and in the Risk Evaluation response (under details). | 
**Type** | [**EnumPredictorType**](EnumPredictorType.md) |  | 
**Description** | Pointer to **string** | A string type. This specifies the description of the risk predictor. Maximum length is 1024 characters. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was updated. | [optional] [readonly] 
**Licensed** | Pointer to **bool** | Indicates whether PingOne Risk is licensed for the environment. | [optional] [readonly] 
**Deletable** | Pointer to **bool** | A boolean to indicate whether the predictor is deletable in the environment. | [optional] [readonly] 
**Default** | Pointer to [**RiskPredictorCommonDefault**](RiskPredictorCommonDefault.md) |  | [optional] 
**Condition** | Pointer to [**RiskPredictorCommonCondition**](RiskPredictorCommonCondition.md) |  | [optional] 
**DomainWhiteList** | Pointer to **[]string** | A list of strings that specify legitimate domains that users will access for restricted resources. | [optional] 
**WhiteList** | Pointer to **[]string** | A list of IP addresses (CDIRs) that are ignored for the predictor results. | [optional] 
**IncludeRepeatedEventsWithoutSdk** | Pointer to **bool** | Set the value of &#x60;includeRepeatedEventsWithoutSdk&#x60; to &#x60;true&#x60; to expand the range of bot activity that PingOne Protect can detect. | [optional] 
**Compositions** | [**[]RiskPredictorCompositeAllOfCompositionsInner**](RiskPredictorCompositeAllOfCompositionsInner.md) | Contains the objects that specify the conditions to test and the risk level that should be assigned if the conditions are met. The array can contain a maximum of three elements. | 
**Map** | [**RiskPredictorCustomAllOfMap**](RiskPredictorCustomAllOfMap.md) |  | 
**Detect** | [**EnumPredictorNewDeviceDetectType**](EnumPredictorNewDeviceDetectType.md) |  | 
**ShouldValidatePayloadSignature** | Pointer to **bool** | Relevant only for Suspicious Device predictors. If &#x60;shouldValidatePayloadSignature&#x60; is set to &#x60;true&#x60;, then any risk policies that include this predictor will require that the Signals SDK payload be provided as a signed JWT whose signature will be verified before proceeding with risk evaluation. You instruct the Signals SDK to provide the payload as a signed JWT by using the &#x60;universalDeviceIdentification&#x60; flag during initialization of the SDK, or by selecting the relevant setting for the &#x60;skrisk&#x60; component in DaVinci flows. | [optional] 
**ActivationAt** | Pointer to **time.Time** | You can use the &#x60;activationAt&#x60; parameter to specify a date on which the learning process for the predictor should be restarted. This can be used in conjunction with the fallback setting (&#x60;default.result.level&#x60;) to force strong authentication when moving the predictor to production. The date should be in an RFC3339 format. Note that activation date uses UTC time. | [optional] 
**Rules** | [**[]RiskPredictorTrafficAnomalyAllOfRules**](RiskPredictorTrafficAnomalyAllOfRules.md) |  | 
**Days** | **int32** |  | 
**Radius** | [**RiskPredictorUserLocationAnomalyAllOfRadius**](RiskPredictorUserLocationAnomalyAllOfRadius.md) |  | 
**PredictionModel** | [**RiskPredictorUserRiskBehaviorAllOfPredictionModel**](RiskPredictorUserRiskBehaviorAllOfPredictionModel.md) |  | 
**By** | Pointer to **[]string** |  | [optional] 
**Every** | Pointer to [**RiskPredictorVelocityAllOfEvery**](RiskPredictorVelocityAllOfEvery.md) |  | [optional] 
**Fallback** | Pointer to [**RiskPredictorVelocityAllOfFallback**](RiskPredictorVelocityAllOfFallback.md) |  | [optional] 
**MaxDelay** | Pointer to [**RiskPredictorVelocityAllOfMaxDelay**](RiskPredictorVelocityAllOfMaxDelay.md) |  | [optional] 
**Measure** | Pointer to [**EnumPredictorVelocityMeasure**](EnumPredictorVelocityMeasure.md) |  | [optional] 
**Of** | Pointer to **string** |  | [optional] 
**SlidingWindow** | Pointer to [**RiskPredictorVelocityAllOfSlidingWindow**](RiskPredictorVelocityAllOfSlidingWindow.md) |  | [optional] 
**Use** | Pointer to [**RiskPredictorVelocityAllOfUse**](RiskPredictorVelocityAllOfUse.md) |  | [optional] 

## Methods

### NewRiskPredictor

`func NewRiskPredictor(name string, compactName string, type_ EnumPredictorType, compositions []RiskPredictorCompositeAllOfCompositionsInner, map_ RiskPredictorCustomAllOfMap, detect EnumPredictorNewDeviceDetectType, rules []RiskPredictorTrafficAnomalyAllOfRules, days int32, radius RiskPredictorUserLocationAnomalyAllOfRadius, predictionModel RiskPredictorUserRiskBehaviorAllOfPredictionModel, ) *RiskPredictor`

NewRiskPredictor instantiates a new RiskPredictor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPredictorWithDefaults

`func NewRiskPredictorWithDefaults() *RiskPredictor`

NewRiskPredictorWithDefaults instantiates a new RiskPredictor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *RiskPredictor) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RiskPredictor) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RiskPredictor) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RiskPredictor) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *RiskPredictor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskPredictor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskPredictor) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskPredictor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *RiskPredictor) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *RiskPredictor) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *RiskPredictor) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *RiskPredictor) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetName

`func (o *RiskPredictor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskPredictor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskPredictor) SetName(v string)`

SetName sets Name field to given value.


### GetCompactName

`func (o *RiskPredictor) GetCompactName() string`

GetCompactName returns the CompactName field if non-nil, zero value otherwise.

### GetCompactNameOk

`func (o *RiskPredictor) GetCompactNameOk() (*string, bool)`

GetCompactNameOk returns a tuple with the CompactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactName

`func (o *RiskPredictor) SetCompactName(v string)`

SetCompactName sets CompactName field to given value.


### GetType

`func (o *RiskPredictor) GetType() EnumPredictorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskPredictor) GetTypeOk() (*EnumPredictorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskPredictor) SetType(v EnumPredictorType)`

SetType sets Type field to given value.


### GetDescription

`func (o *RiskPredictor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RiskPredictor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RiskPredictor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RiskPredictor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RiskPredictor) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RiskPredictor) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RiskPredictor) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RiskPredictor) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RiskPredictor) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RiskPredictor) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RiskPredictor) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RiskPredictor) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLicensed

`func (o *RiskPredictor) GetLicensed() bool`

GetLicensed returns the Licensed field if non-nil, zero value otherwise.

### GetLicensedOk

`func (o *RiskPredictor) GetLicensedOk() (*bool, bool)`

GetLicensedOk returns a tuple with the Licensed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensed

`func (o *RiskPredictor) SetLicensed(v bool)`

SetLicensed sets Licensed field to given value.

### HasLicensed

`func (o *RiskPredictor) HasLicensed() bool`

HasLicensed returns a boolean if a field has been set.

### GetDeletable

`func (o *RiskPredictor) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *RiskPredictor) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *RiskPredictor) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *RiskPredictor) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetDefault

`func (o *RiskPredictor) GetDefault() RiskPredictorCommonDefault`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RiskPredictor) GetDefaultOk() (*RiskPredictorCommonDefault, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RiskPredictor) SetDefault(v RiskPredictorCommonDefault)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *RiskPredictor) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetCondition

`func (o *RiskPredictor) GetCondition() RiskPredictorCommonCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *RiskPredictor) GetConditionOk() (*RiskPredictorCommonCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *RiskPredictor) SetCondition(v RiskPredictorCommonCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *RiskPredictor) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetDomainWhiteList

`func (o *RiskPredictor) GetDomainWhiteList() []string`

GetDomainWhiteList returns the DomainWhiteList field if non-nil, zero value otherwise.

### GetDomainWhiteListOk

`func (o *RiskPredictor) GetDomainWhiteListOk() (*[]string, bool)`

GetDomainWhiteListOk returns a tuple with the DomainWhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainWhiteList

`func (o *RiskPredictor) SetDomainWhiteList(v []string)`

SetDomainWhiteList sets DomainWhiteList field to given value.

### HasDomainWhiteList

`func (o *RiskPredictor) HasDomainWhiteList() bool`

HasDomainWhiteList returns a boolean if a field has been set.

### GetWhiteList

`func (o *RiskPredictor) GetWhiteList() []string`

GetWhiteList returns the WhiteList field if non-nil, zero value otherwise.

### GetWhiteListOk

`func (o *RiskPredictor) GetWhiteListOk() (*[]string, bool)`

GetWhiteListOk returns a tuple with the WhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteList

`func (o *RiskPredictor) SetWhiteList(v []string)`

SetWhiteList sets WhiteList field to given value.

### HasWhiteList

`func (o *RiskPredictor) HasWhiteList() bool`

HasWhiteList returns a boolean if a field has been set.

### GetIncludeRepeatedEventsWithoutSdk

`func (o *RiskPredictor) GetIncludeRepeatedEventsWithoutSdk() bool`

GetIncludeRepeatedEventsWithoutSdk returns the IncludeRepeatedEventsWithoutSdk field if non-nil, zero value otherwise.

### GetIncludeRepeatedEventsWithoutSdkOk

`func (o *RiskPredictor) GetIncludeRepeatedEventsWithoutSdkOk() (*bool, bool)`

GetIncludeRepeatedEventsWithoutSdkOk returns a tuple with the IncludeRepeatedEventsWithoutSdk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRepeatedEventsWithoutSdk

`func (o *RiskPredictor) SetIncludeRepeatedEventsWithoutSdk(v bool)`

SetIncludeRepeatedEventsWithoutSdk sets IncludeRepeatedEventsWithoutSdk field to given value.

### HasIncludeRepeatedEventsWithoutSdk

`func (o *RiskPredictor) HasIncludeRepeatedEventsWithoutSdk() bool`

HasIncludeRepeatedEventsWithoutSdk returns a boolean if a field has been set.

### GetCompositions

`func (o *RiskPredictor) GetCompositions() []RiskPredictorCompositeAllOfCompositionsInner`

GetCompositions returns the Compositions field if non-nil, zero value otherwise.

### GetCompositionsOk

`func (o *RiskPredictor) GetCompositionsOk() (*[]RiskPredictorCompositeAllOfCompositionsInner, bool)`

GetCompositionsOk returns a tuple with the Compositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositions

`func (o *RiskPredictor) SetCompositions(v []RiskPredictorCompositeAllOfCompositionsInner)`

SetCompositions sets Compositions field to given value.


### GetMap

`func (o *RiskPredictor) GetMap() RiskPredictorCustomAllOfMap`

GetMap returns the Map field if non-nil, zero value otherwise.

### GetMapOk

`func (o *RiskPredictor) GetMapOk() (*RiskPredictorCustomAllOfMap, bool)`

GetMapOk returns a tuple with the Map field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMap

`func (o *RiskPredictor) SetMap(v RiskPredictorCustomAllOfMap)`

SetMap sets Map field to given value.


### GetDetect

`func (o *RiskPredictor) GetDetect() EnumPredictorNewDeviceDetectType`

GetDetect returns the Detect field if non-nil, zero value otherwise.

### GetDetectOk

`func (o *RiskPredictor) GetDetectOk() (*EnumPredictorNewDeviceDetectType, bool)`

GetDetectOk returns a tuple with the Detect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetect

`func (o *RiskPredictor) SetDetect(v EnumPredictorNewDeviceDetectType)`

SetDetect sets Detect field to given value.


### GetShouldValidatePayloadSignature

`func (o *RiskPredictor) GetShouldValidatePayloadSignature() bool`

GetShouldValidatePayloadSignature returns the ShouldValidatePayloadSignature field if non-nil, zero value otherwise.

### GetShouldValidatePayloadSignatureOk

`func (o *RiskPredictor) GetShouldValidatePayloadSignatureOk() (*bool, bool)`

GetShouldValidatePayloadSignatureOk returns a tuple with the ShouldValidatePayloadSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldValidatePayloadSignature

`func (o *RiskPredictor) SetShouldValidatePayloadSignature(v bool)`

SetShouldValidatePayloadSignature sets ShouldValidatePayloadSignature field to given value.

### HasShouldValidatePayloadSignature

`func (o *RiskPredictor) HasShouldValidatePayloadSignature() bool`

HasShouldValidatePayloadSignature returns a boolean if a field has been set.

### GetActivationAt

`func (o *RiskPredictor) GetActivationAt() time.Time`

GetActivationAt returns the ActivationAt field if non-nil, zero value otherwise.

### GetActivationAtOk

`func (o *RiskPredictor) GetActivationAtOk() (*time.Time, bool)`

GetActivationAtOk returns a tuple with the ActivationAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationAt

`func (o *RiskPredictor) SetActivationAt(v time.Time)`

SetActivationAt sets ActivationAt field to given value.

### HasActivationAt

`func (o *RiskPredictor) HasActivationAt() bool`

HasActivationAt returns a boolean if a field has been set.

### GetRules

`func (o *RiskPredictor) GetRules() []RiskPredictorTrafficAnomalyAllOfRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *RiskPredictor) GetRulesOk() (*[]RiskPredictorTrafficAnomalyAllOfRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *RiskPredictor) SetRules(v []RiskPredictorTrafficAnomalyAllOfRules)`

SetRules sets Rules field to given value.


### GetDays

`func (o *RiskPredictor) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *RiskPredictor) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *RiskPredictor) SetDays(v int32)`

SetDays sets Days field to given value.


### GetRadius

`func (o *RiskPredictor) GetRadius() RiskPredictorUserLocationAnomalyAllOfRadius`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *RiskPredictor) GetRadiusOk() (*RiskPredictorUserLocationAnomalyAllOfRadius, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *RiskPredictor) SetRadius(v RiskPredictorUserLocationAnomalyAllOfRadius)`

SetRadius sets Radius field to given value.


### GetPredictionModel

`func (o *RiskPredictor) GetPredictionModel() RiskPredictorUserRiskBehaviorAllOfPredictionModel`

GetPredictionModel returns the PredictionModel field if non-nil, zero value otherwise.

### GetPredictionModelOk

`func (o *RiskPredictor) GetPredictionModelOk() (*RiskPredictorUserRiskBehaviorAllOfPredictionModel, bool)`

GetPredictionModelOk returns a tuple with the PredictionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictionModel

`func (o *RiskPredictor) SetPredictionModel(v RiskPredictorUserRiskBehaviorAllOfPredictionModel)`

SetPredictionModel sets PredictionModel field to given value.


### GetBy

`func (o *RiskPredictor) GetBy() []string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *RiskPredictor) GetByOk() (*[]string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *RiskPredictor) SetBy(v []string)`

SetBy sets By field to given value.

### HasBy

`func (o *RiskPredictor) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetEvery

`func (o *RiskPredictor) GetEvery() RiskPredictorVelocityAllOfEvery`

GetEvery returns the Every field if non-nil, zero value otherwise.

### GetEveryOk

`func (o *RiskPredictor) GetEveryOk() (*RiskPredictorVelocityAllOfEvery, bool)`

GetEveryOk returns a tuple with the Every field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvery

`func (o *RiskPredictor) SetEvery(v RiskPredictorVelocityAllOfEvery)`

SetEvery sets Every field to given value.

### HasEvery

`func (o *RiskPredictor) HasEvery() bool`

HasEvery returns a boolean if a field has been set.

### GetFallback

`func (o *RiskPredictor) GetFallback() RiskPredictorVelocityAllOfFallback`

GetFallback returns the Fallback field if non-nil, zero value otherwise.

### GetFallbackOk

`func (o *RiskPredictor) GetFallbackOk() (*RiskPredictorVelocityAllOfFallback, bool)`

GetFallbackOk returns a tuple with the Fallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallback

`func (o *RiskPredictor) SetFallback(v RiskPredictorVelocityAllOfFallback)`

SetFallback sets Fallback field to given value.

### HasFallback

`func (o *RiskPredictor) HasFallback() bool`

HasFallback returns a boolean if a field has been set.

### GetMaxDelay

`func (o *RiskPredictor) GetMaxDelay() RiskPredictorVelocityAllOfMaxDelay`

GetMaxDelay returns the MaxDelay field if non-nil, zero value otherwise.

### GetMaxDelayOk

`func (o *RiskPredictor) GetMaxDelayOk() (*RiskPredictorVelocityAllOfMaxDelay, bool)`

GetMaxDelayOk returns a tuple with the MaxDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDelay

`func (o *RiskPredictor) SetMaxDelay(v RiskPredictorVelocityAllOfMaxDelay)`

SetMaxDelay sets MaxDelay field to given value.

### HasMaxDelay

`func (o *RiskPredictor) HasMaxDelay() bool`

HasMaxDelay returns a boolean if a field has been set.

### GetMeasure

`func (o *RiskPredictor) GetMeasure() EnumPredictorVelocityMeasure`

GetMeasure returns the Measure field if non-nil, zero value otherwise.

### GetMeasureOk

`func (o *RiskPredictor) GetMeasureOk() (*EnumPredictorVelocityMeasure, bool)`

GetMeasureOk returns a tuple with the Measure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasure

`func (o *RiskPredictor) SetMeasure(v EnumPredictorVelocityMeasure)`

SetMeasure sets Measure field to given value.

### HasMeasure

`func (o *RiskPredictor) HasMeasure() bool`

HasMeasure returns a boolean if a field has been set.

### GetOf

`func (o *RiskPredictor) GetOf() string`

GetOf returns the Of field if non-nil, zero value otherwise.

### GetOfOk

`func (o *RiskPredictor) GetOfOk() (*string, bool)`

GetOfOk returns a tuple with the Of field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOf

`func (o *RiskPredictor) SetOf(v string)`

SetOf sets Of field to given value.

### HasOf

`func (o *RiskPredictor) HasOf() bool`

HasOf returns a boolean if a field has been set.

### GetSlidingWindow

`func (o *RiskPredictor) GetSlidingWindow() RiskPredictorVelocityAllOfSlidingWindow`

GetSlidingWindow returns the SlidingWindow field if non-nil, zero value otherwise.

### GetSlidingWindowOk

`func (o *RiskPredictor) GetSlidingWindowOk() (*RiskPredictorVelocityAllOfSlidingWindow, bool)`

GetSlidingWindowOk returns a tuple with the SlidingWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlidingWindow

`func (o *RiskPredictor) SetSlidingWindow(v RiskPredictorVelocityAllOfSlidingWindow)`

SetSlidingWindow sets SlidingWindow field to given value.

### HasSlidingWindow

`func (o *RiskPredictor) HasSlidingWindow() bool`

HasSlidingWindow returns a boolean if a field has been set.

### GetUse

`func (o *RiskPredictor) GetUse() RiskPredictorVelocityAllOfUse`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *RiskPredictor) GetUseOk() (*RiskPredictorVelocityAllOfUse, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *RiskPredictor) SetUse(v RiskPredictorVelocityAllOfUse)`

SetUse sets Use field to given value.

### HasUse

`func (o *RiskPredictor) HasUse() bool`

HasUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


