/*
Bitly API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.0
Contact: api@bitly.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PublicSSLCert SSL Certification
type PublicSSLCert struct {
	CustomCert *bool   `json:"custom_cert,omitempty"`
	ValidEnd   *int32  `json:"valid_end,omitempty"`
	Issuer     *string `json:"issuer,omitempty"`
}

// NewPublicSSLCert instantiates a new PublicSSLCert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicSSLCert() *PublicSSLCert {
	this := PublicSSLCert{}
	return &this
}

// NewPublicSSLCertWithDefaults instantiates a new PublicSSLCert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicSSLCertWithDefaults() *PublicSSLCert {
	this := PublicSSLCert{}
	return &this
}

// GetCustomCert returns the CustomCert field value if set, zero value otherwise.
func (o *PublicSSLCert) GetCustomCert() bool {
	if o == nil || o.CustomCert == nil {
		var ret bool
		return ret
	}
	return *o.CustomCert
}

// GetCustomCertOk returns a tuple with the CustomCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicSSLCert) GetCustomCertOk() (*bool, bool) {
	if o == nil || o.CustomCert == nil {
		return nil, false
	}
	return o.CustomCert, true
}

// HasCustomCert returns a boolean if a field has been set.
func (o *PublicSSLCert) HasCustomCert() bool {
	if o != nil && o.CustomCert != nil {
		return true
	}

	return false
}

// SetCustomCert gets a reference to the given bool and assigns it to the CustomCert field.
func (o *PublicSSLCert) SetCustomCert(v bool) {
	o.CustomCert = &v
}

// GetValidEnd returns the ValidEnd field value if set, zero value otherwise.
func (o *PublicSSLCert) GetValidEnd() int32 {
	if o == nil || o.ValidEnd == nil {
		var ret int32
		return ret
	}
	return *o.ValidEnd
}

// GetValidEndOk returns a tuple with the ValidEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicSSLCert) GetValidEndOk() (*int32, bool) {
	if o == nil || o.ValidEnd == nil {
		return nil, false
	}
	return o.ValidEnd, true
}

// HasValidEnd returns a boolean if a field has been set.
func (o *PublicSSLCert) HasValidEnd() bool {
	if o != nil && o.ValidEnd != nil {
		return true
	}

	return false
}

// SetValidEnd gets a reference to the given int32 and assigns it to the ValidEnd field.
func (o *PublicSSLCert) SetValidEnd(v int32) {
	o.ValidEnd = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *PublicSSLCert) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicSSLCert) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *PublicSSLCert) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *PublicSSLCert) SetIssuer(v string) {
	o.Issuer = &v
}

func (o PublicSSLCert) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomCert != nil {
		toSerialize["custom_cert"] = o.CustomCert
	}
	if o.ValidEnd != nil {
		toSerialize["valid_end"] = o.ValidEnd
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	return json.Marshal(toSerialize)
}

type NullablePublicSSLCert struct {
	value *PublicSSLCert
	isSet bool
}

func (v NullablePublicSSLCert) Get() *PublicSSLCert {
	return v.value
}

func (v *NullablePublicSSLCert) Set(val *PublicSSLCert) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicSSLCert) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicSSLCert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicSSLCert(val *PublicSSLCert) *NullablePublicSSLCert {
	return &NullablePublicSSLCert{value: val, isSet: true}
}

func (v NullablePublicSSLCert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicSSLCert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}