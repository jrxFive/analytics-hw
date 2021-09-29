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

// MinimalDeeplinkApp Deep Linking object
type MinimalDeeplinkApp struct {
	Guid *string `json:"guid,omitempty"`
	Os   *string `json:"os,omitempty"`
}

// NewMinimalDeeplinkApp instantiates a new MinimalDeeplinkApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinimalDeeplinkApp() *MinimalDeeplinkApp {
	this := MinimalDeeplinkApp{}
	return &this
}

// NewMinimalDeeplinkAppWithDefaults instantiates a new MinimalDeeplinkApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinimalDeeplinkAppWithDefaults() *MinimalDeeplinkApp {
	this := MinimalDeeplinkApp{}
	return &this
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *MinimalDeeplinkApp) GetGuid() string {
	if o == nil || o.Guid == nil {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeeplinkApp) GetGuidOk() (*string, bool) {
	if o == nil || o.Guid == nil {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *MinimalDeeplinkApp) HasGuid() bool {
	if o != nil && o.Guid != nil {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *MinimalDeeplinkApp) SetGuid(v string) {
	o.Guid = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *MinimalDeeplinkApp) GetOs() string {
	if o == nil || o.Os == nil {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeeplinkApp) GetOsOk() (*string, bool) {
	if o == nil || o.Os == nil {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *MinimalDeeplinkApp) HasOs() bool {
	if o != nil && o.Os != nil {
		return true
	}

	return false
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *MinimalDeeplinkApp) SetOs(v string) {
	o.Os = &v
}

func (o MinimalDeeplinkApp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Guid != nil {
		toSerialize["guid"] = o.Guid
	}
	if o.Os != nil {
		toSerialize["os"] = o.Os
	}
	return json.Marshal(toSerialize)
}

type NullableMinimalDeeplinkApp struct {
	value *MinimalDeeplinkApp
	isSet bool
}

func (v NullableMinimalDeeplinkApp) Get() *MinimalDeeplinkApp {
	return v.value
}

func (v *NullableMinimalDeeplinkApp) Set(val *MinimalDeeplinkApp) {
	v.value = val
	v.isSet = true
}

func (v NullableMinimalDeeplinkApp) IsSet() bool {
	return v.isSet
}

func (v *NullableMinimalDeeplinkApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinimalDeeplinkApp(val *MinimalDeeplinkApp) *NullableMinimalDeeplinkApp {
	return &NullableMinimalDeeplinkApp{value: val, isSet: true}
}

func (v NullableMinimalDeeplinkApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinimalDeeplinkApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
