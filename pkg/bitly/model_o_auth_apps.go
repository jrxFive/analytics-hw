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

// OAuthApps struct for OAuthApps
type OAuthApps struct {
	Applications *[]OAuthApp `json:"applications,omitempty"`
}

// NewOAuthApps instantiates a new OAuthApps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthApps() *OAuthApps {
	this := OAuthApps{}
	return &this
}

// NewOAuthAppsWithDefaults instantiates a new OAuthApps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthAppsWithDefaults() *OAuthApps {
	this := OAuthApps{}
	return &this
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *OAuthApps) GetApplications() []OAuthApp {
	if o == nil || o.Applications == nil {
		var ret []OAuthApp
		return ret
	}
	return *o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthApps) GetApplicationsOk() (*[]OAuthApp, bool) {
	if o == nil || o.Applications == nil {
		return nil, false
	}
	return o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *OAuthApps) HasApplications() bool {
	if o != nil && o.Applications != nil {
		return true
	}

	return false
}

// SetApplications gets a reference to the given []OAuthApp and assigns it to the Applications field.
func (o *OAuthApps) SetApplications(v []OAuthApp) {
	o.Applications = &v
}

func (o OAuthApps) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Applications != nil {
		toSerialize["applications"] = o.Applications
	}
	return json.Marshal(toSerialize)
}

type NullableOAuthApps struct {
	value *OAuthApps
	isSet bool
}

func (v NullableOAuthApps) Get() *OAuthApps {
	return v.value
}

func (v *NullableOAuthApps) Set(val *OAuthApps) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthApps) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthApps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthApps(val *OAuthApps) *NullableOAuthApps {
	return &NullableOAuthApps{value: val, isSet: true}
}

func (v NullableOAuthApps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthApps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
