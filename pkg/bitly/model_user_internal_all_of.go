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

// UserInternalAllOf struct for UserInternalAllOf
type UserInternalAllOf struct {
	RoleName *string `json:"role_name,omitempty"`
}

// NewUserInternalAllOf instantiates a new UserInternalAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInternalAllOf() *UserInternalAllOf {
	this := UserInternalAllOf{}
	return &this
}

// NewUserInternalAllOfWithDefaults instantiates a new UserInternalAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInternalAllOfWithDefaults() *UserInternalAllOf {
	this := UserInternalAllOf{}
	return &this
}

// GetRoleName returns the RoleName field value if set, zero value otherwise.
func (o *UserInternalAllOf) GetRoleName() string {
	if o == nil || o.RoleName == nil {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInternalAllOf) GetRoleNameOk() (*string, bool) {
	if o == nil || o.RoleName == nil {
		return nil, false
	}
	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *UserInternalAllOf) HasRoleName() bool {
	if o != nil && o.RoleName != nil {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *UserInternalAllOf) SetRoleName(v string) {
	o.RoleName = &v
}

func (o UserInternalAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RoleName != nil {
		toSerialize["role_name"] = o.RoleName
	}
	return json.Marshal(toSerialize)
}

type NullableUserInternalAllOf struct {
	value *UserInternalAllOf
	isSet bool
}

func (v NullableUserInternalAllOf) Get() *UserInternalAllOf {
	return v.value
}

func (v *NullableUserInternalAllOf) Set(val *UserInternalAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInternalAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInternalAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInternalAllOf(val *UserInternalAllOf) *NullableUserInternalAllOf {
	return &NullableUserInternalAllOf{value: val, isSet: true}
}

func (v NullableUserInternalAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInternalAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
