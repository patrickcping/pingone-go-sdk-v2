/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-07-18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// IdentityProvider - struct for IdentityProvider
type IdentityProvider struct {
	IdentityProviderAmazon *IdentityProviderAmazon
	IdentityProviderApple *IdentityProviderApple
	IdentityProviderFacebook *IdentityProviderFacebook
	IdentityProviderGithub *IdentityProviderGithub
	IdentityProviderGoogle *IdentityProviderGoogle
	IdentityProviderLinkedIn *IdentityProviderLinkedIn
	IdentityProviderMicrosoft *IdentityProviderMicrosoft
	IdentityProviderOIDC *IdentityProviderOIDC
	IdentityProviderPaypal *IdentityProviderPaypal
	IdentityProviderSAML *IdentityProviderSAML
	IdentityProviderTwitter *IdentityProviderTwitter
	IdentityProviderYahoo *IdentityProviderYahoo
}

// IdentityProviderAmazonAsIdentityProvider is a convenience function that returns IdentityProviderAmazon wrapped in IdentityProvider
func IdentityProviderAmazonAsIdentityProvider(v *IdentityProviderAmazon) IdentityProvider {
	return IdentityProvider{
		IdentityProviderAmazon: v,
	}
}

// IdentityProviderAppleAsIdentityProvider is a convenience function that returns IdentityProviderApple wrapped in IdentityProvider
func IdentityProviderAppleAsIdentityProvider(v *IdentityProviderApple) IdentityProvider {
	return IdentityProvider{
		IdentityProviderApple: v,
	}
}

// IdentityProviderFacebookAsIdentityProvider is a convenience function that returns IdentityProviderFacebook wrapped in IdentityProvider
func IdentityProviderFacebookAsIdentityProvider(v *IdentityProviderFacebook) IdentityProvider {
	return IdentityProvider{
		IdentityProviderFacebook: v,
	}
}

// IdentityProviderGithubAsIdentityProvider is a convenience function that returns IdentityProviderGithub wrapped in IdentityProvider
func IdentityProviderGithubAsIdentityProvider(v *IdentityProviderGithub) IdentityProvider {
	return IdentityProvider{
		IdentityProviderGithub: v,
	}
}

// IdentityProviderGoogleAsIdentityProvider is a convenience function that returns IdentityProviderGoogle wrapped in IdentityProvider
func IdentityProviderGoogleAsIdentityProvider(v *IdentityProviderGoogle) IdentityProvider {
	return IdentityProvider{
		IdentityProviderGoogle: v,
	}
}

// IdentityProviderLinkedInAsIdentityProvider is a convenience function that returns IdentityProviderLinkedIn wrapped in IdentityProvider
func IdentityProviderLinkedInAsIdentityProvider(v *IdentityProviderLinkedIn) IdentityProvider {
	return IdentityProvider{
		IdentityProviderLinkedIn: v,
	}
}

// IdentityProviderMicrosoftAsIdentityProvider is a convenience function that returns IdentityProviderMicrosoft wrapped in IdentityProvider
func IdentityProviderMicrosoftAsIdentityProvider(v *IdentityProviderMicrosoft) IdentityProvider {
	return IdentityProvider{
		IdentityProviderMicrosoft: v,
	}
}

// IdentityProviderOIDCAsIdentityProvider is a convenience function that returns IdentityProviderOIDC wrapped in IdentityProvider
func IdentityProviderOIDCAsIdentityProvider(v *IdentityProviderOIDC) IdentityProvider {
	return IdentityProvider{
		IdentityProviderOIDC: v,
	}
}

// IdentityProviderPaypalAsIdentityProvider is a convenience function that returns IdentityProviderPaypal wrapped in IdentityProvider
func IdentityProviderPaypalAsIdentityProvider(v *IdentityProviderPaypal) IdentityProvider {
	return IdentityProvider{
		IdentityProviderPaypal: v,
	}
}

// IdentityProviderSAMLAsIdentityProvider is a convenience function that returns IdentityProviderSAML wrapped in IdentityProvider
func IdentityProviderSAMLAsIdentityProvider(v *IdentityProviderSAML) IdentityProvider {
	return IdentityProvider{
		IdentityProviderSAML: v,
	}
}

// IdentityProviderTwitterAsIdentityProvider is a convenience function that returns IdentityProviderTwitter wrapped in IdentityProvider
func IdentityProviderTwitterAsIdentityProvider(v *IdentityProviderTwitter) IdentityProvider {
	return IdentityProvider{
		IdentityProviderTwitter: v,
	}
}

// IdentityProviderYahooAsIdentityProvider is a convenience function that returns IdentityProviderYahoo wrapped in IdentityProvider
func IdentityProviderYahooAsIdentityProvider(v *IdentityProviderYahoo) IdentityProvider {
	return IdentityProvider{
		IdentityProviderYahoo: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *IdentityProvider) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IdentityProviderAmazon
	err = newStrictDecoder(data).Decode(&dst.IdentityProviderAmazon)
	if err == nil {
		jsonIdentityProviderAmazon, _ := json.Marshal(dst.IdentityProviderAmazon)
		if string(jsonIdentityProviderAmazon) == "{}" { // empty struct
			dst.IdentityProviderAmazon = nil
		} else {
			match++
		}
	} else {
		dst.IdentityProviderAmazon = nil
	}

	// try to unmarshal data into IdentityProviderApple
	err = newStrictDecoder(data).Decode(&dst.IdentityProviderApple)
	if err == nil {
		jsonIdentityProviderApple, _ := json.Marshal(dst.IdentityProviderApple)
		if string(jsonIdentityProviderApple) == "{}" { // empty struct
			dst.IdentityProviderApple = nil
		} else {
			match++
		}
	} else {
		dst.IdentityProviderApple = nil
	}

	// try to unmarshal data into IdentityProviderFacebook
	err = newStrictDecoder(data).Decode(&dst.IdentityProviderFacebook)
	if err == nil {
		jsonIdentityProviderFacebook, _ := json.Marshal(dst.IdentityProviderFacebook)
		if string(jsonIdentityProviderFacebook) == "{}" { // empty struct
			dst.IdentityProviderFacebook = nil
		} else {
			match++
		}
	} else {
		dst.IdentityProviderFacebook = nil
	}

	// try to unmarshal data into IdentityProviderGithub
	err = newStrictDecoder(data).Decode(&dst.IdentityProviderGithub)
	if err == nil {
		jsonIdentityProviderGithub, _ := json.Marshal(dst.IdentityProviderGithub)
		if string(jsonIdentityProviderGithub) == "{}" { // empty struct
			dst.IdentityProviderGithub = nil
		} else {
			match++
		}
	} else {
		dst.IdentityProviderGithub = nil
	}

	// try to unmarshal data into IdentityProviderGoogle
	err = newStrictDecoder(data).Decode(&dst.IdentityProviderGoogle)
	if err == nil {
		jsonIdentityProviderGoogle, _ := json.Marshal(dst.IdentityProviderGoogle)
		if string(jsonIdentityProviderGoogle) == "{}" { // empty struct
			dst.IdentityProviderGoogle = nil
		} else {
			match++
		}
	} else {
		dst.IdentityProviderGoogle = nil
	}

	// try to unmarshal data into IdentityProviderLinkedIn
	err = newStrictDecoder(data).Decode(&dst.IdentityProviderLinkedIn)
	if err == nil {
		jsonIdentityProviderLinkedIn, _ := json.Marshal(dst.IdentityProviderLinkedIn)
		if string(jsonIdentityProviderLinkedIn) == "{}" { // empty struct
			dst.IdentityProviderLinkedIn = nil
		} else {
			match++
		}
	} else {
		dst.IdentityProviderLinkedIn = nil
	}

	// try to unmarshal data into IdentityProviderMicrosoft
	err = newStrictDecoder(data).Decode(&dst.IdentityProviderMicrosoft)
	if err == nil {
		jsonIdentityProviderMicrosoft, _ := json.Marshal(dst.IdentityProviderMicrosoft)
		if string(jsonIdentityProviderMicrosoft) == "{}" { // empty struct
			dst.IdentityProviderMicrosoft = nil
		} else {
			match++
		}
	} else {
		dst.IdentityProviderMicrosoft = nil
	}

	// try to unmarshal data into IdentityProviderOIDC
	err = newStrictDecoder(data).Decode(&dst.IdentityProviderOIDC)
	if err == nil {
		jsonIdentityProviderOIDC, _ := json.Marshal(dst.IdentityProviderOIDC)
		if string(jsonIdentityProviderOIDC) == "{}" { // empty struct
			dst.IdentityProviderOIDC = nil
		} else {
			match++
		}
	} else {
		dst.IdentityProviderOIDC = nil
	}

	// try to unmarshal data into IdentityProviderPaypal
	err = newStrictDecoder(data).Decode(&dst.IdentityProviderPaypal)
	if err == nil {
		jsonIdentityProviderPaypal, _ := json.Marshal(dst.IdentityProviderPaypal)
		if string(jsonIdentityProviderPaypal) == "{}" { // empty struct
			dst.IdentityProviderPaypal = nil
		} else {
			match++
		}
	} else {
		dst.IdentityProviderPaypal = nil
	}

	// try to unmarshal data into IdentityProviderSAML
	err = newStrictDecoder(data).Decode(&dst.IdentityProviderSAML)
	if err == nil {
		jsonIdentityProviderSAML, _ := json.Marshal(dst.IdentityProviderSAML)
		if string(jsonIdentityProviderSAML) == "{}" { // empty struct
			dst.IdentityProviderSAML = nil
		} else {
			match++
		}
	} else {
		dst.IdentityProviderSAML = nil
	}

	// try to unmarshal data into IdentityProviderTwitter
	err = newStrictDecoder(data).Decode(&dst.IdentityProviderTwitter)
	if err == nil {
		jsonIdentityProviderTwitter, _ := json.Marshal(dst.IdentityProviderTwitter)
		if string(jsonIdentityProviderTwitter) == "{}" { // empty struct
			dst.IdentityProviderTwitter = nil
		} else {
			match++
		}
	} else {
		dst.IdentityProviderTwitter = nil
	}

	// try to unmarshal data into IdentityProviderYahoo
	err = newStrictDecoder(data).Decode(&dst.IdentityProviderYahoo)
	if err == nil {
		jsonIdentityProviderYahoo, _ := json.Marshal(dst.IdentityProviderYahoo)
		if string(jsonIdentityProviderYahoo) == "{}" { // empty struct
			dst.IdentityProviderYahoo = nil
		} else {
			match++
		}
	} else {
		dst.IdentityProviderYahoo = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.IdentityProviderAmazon = nil
		dst.IdentityProviderApple = nil
		dst.IdentityProviderFacebook = nil
		dst.IdentityProviderGithub = nil
		dst.IdentityProviderGoogle = nil
		dst.IdentityProviderLinkedIn = nil
		dst.IdentityProviderMicrosoft = nil
		dst.IdentityProviderOIDC = nil
		dst.IdentityProviderPaypal = nil
		dst.IdentityProviderSAML = nil
		dst.IdentityProviderTwitter = nil
		dst.IdentityProviderYahoo = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(IdentityProvider)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(IdentityProvider)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IdentityProvider) MarshalJSON() ([]byte, error) {
	if src.IdentityProviderAmazon != nil {
		return json.Marshal(&src.IdentityProviderAmazon)
	}

	if src.IdentityProviderApple != nil {
		return json.Marshal(&src.IdentityProviderApple)
	}

	if src.IdentityProviderFacebook != nil {
		return json.Marshal(&src.IdentityProviderFacebook)
	}

	if src.IdentityProviderGithub != nil {
		return json.Marshal(&src.IdentityProviderGithub)
	}

	if src.IdentityProviderGoogle != nil {
		return json.Marshal(&src.IdentityProviderGoogle)
	}

	if src.IdentityProviderLinkedIn != nil {
		return json.Marshal(&src.IdentityProviderLinkedIn)
	}

	if src.IdentityProviderMicrosoft != nil {
		return json.Marshal(&src.IdentityProviderMicrosoft)
	}

	if src.IdentityProviderOIDC != nil {
		return json.Marshal(&src.IdentityProviderOIDC)
	}

	if src.IdentityProviderPaypal != nil {
		return json.Marshal(&src.IdentityProviderPaypal)
	}

	if src.IdentityProviderSAML != nil {
		return json.Marshal(&src.IdentityProviderSAML)
	}

	if src.IdentityProviderTwitter != nil {
		return json.Marshal(&src.IdentityProviderTwitter)
	}

	if src.IdentityProviderYahoo != nil {
		return json.Marshal(&src.IdentityProviderYahoo)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IdentityProvider) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.IdentityProviderAmazon != nil {
		return obj.IdentityProviderAmazon
	}

	if obj.IdentityProviderApple != nil {
		return obj.IdentityProviderApple
	}

	if obj.IdentityProviderFacebook != nil {
		return obj.IdentityProviderFacebook
	}

	if obj.IdentityProviderGithub != nil {
		return obj.IdentityProviderGithub
	}

	if obj.IdentityProviderGoogle != nil {
		return obj.IdentityProviderGoogle
	}

	if obj.IdentityProviderLinkedIn != nil {
		return obj.IdentityProviderLinkedIn
	}

	if obj.IdentityProviderMicrosoft != nil {
		return obj.IdentityProviderMicrosoft
	}

	if obj.IdentityProviderOIDC != nil {
		return obj.IdentityProviderOIDC
	}

	if obj.IdentityProviderPaypal != nil {
		return obj.IdentityProviderPaypal
	}

	if obj.IdentityProviderSAML != nil {
		return obj.IdentityProviderSAML
	}

	if obj.IdentityProviderTwitter != nil {
		return obj.IdentityProviderTwitter
	}

	if obj.IdentityProviderYahoo != nil {
		return obj.IdentityProviderYahoo
	}

	// all schemas are nil
	return nil
}

type NullableIdentityProvider struct {
	value *IdentityProvider
	isSet bool
}

func (v NullableIdentityProvider) Get() *IdentityProvider {
	return v.value
}

func (v *NullableIdentityProvider) Set(val *IdentityProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProvider(val *IdentityProvider) *NullableIdentityProvider {
	return &NullableIdentityProvider{value: val, isSet: true}
}

func (v NullableIdentityProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


