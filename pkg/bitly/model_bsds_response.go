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

// BSDsResponse struct for BSDsResponse
type BSDsResponse struct {
	Bsds *[]string `json:"bsds,omitempty"`
}

// NewBSDsResponse instantiates a new BSDsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBSDsResponse() *BSDsResponse {
	this := BSDsResponse{}
	return &this
}

// NewBSDsResponseWithDefaults instantiates a new BSDsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBSDsResponseWithDefaults() *BSDsResponse {
	this := BSDsResponse{}
	return &this
}

// GetBsds returns the Bsds field value if set, zero value otherwise.
func (o *BSDsResponse) GetBsds() []string {
	if o == nil || o.Bsds == nil {
		var ret []string
		return ret
	}
	return *o.Bsds
}

// GetBsdsOk returns a tuple with the Bsds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BSDsResponse) GetBsdsOk() (*[]string, bool) {
	if o == nil || o.Bsds == nil {
		return nil, false
	}
	return o.Bsds, true
}

// HasBsds returns a boolean if a field has been set.
func (o *BSDsResponse) HasBsds() bool {
	if o != nil && o.Bsds != nil {
		return true
	}

	return false
}

// SetBsds gets a reference to the given []string and assigns it to the Bsds field.
func (o *BSDsResponse) SetBsds(v []string) {
	o.Bsds = &v
}

func (o BSDsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bsds != nil {
		toSerialize["bsds"] = o.Bsds
	}
	return json.Marshal(toSerialize)
}

type NullableBSDsResponse struct {
	value *BSDsResponse
	isSet bool
}

func (v NullableBSDsResponse) Get() *BSDsResponse {
	return v.value
}

func (v *NullableBSDsResponse) Set(val *BSDsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBSDsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBSDsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBSDsResponse(val *BSDsResponse) *NullableBSDsResponse {
	return &NullableBSDsResponse{value: val, isSet: true}
}

func (v NullableBSDsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBSDsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
