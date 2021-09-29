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

// DefaultAccessFeature the tier default values for a single access feature
type DefaultAccessFeature struct {
	HasAccess *bool   `json:"has_access,omitempty"`
	Name      *string `json:"name,omitempty"`
}

// NewDefaultAccessFeature instantiates a new DefaultAccessFeature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultAccessFeature() *DefaultAccessFeature {
	this := DefaultAccessFeature{}
	return &this
}

// NewDefaultAccessFeatureWithDefaults instantiates a new DefaultAccessFeature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultAccessFeatureWithDefaults() *DefaultAccessFeature {
	this := DefaultAccessFeature{}
	return &this
}

// GetHasAccess returns the HasAccess field value if set, zero value otherwise.
func (o *DefaultAccessFeature) GetHasAccess() bool {
	if o == nil || o.HasAccess == nil {
		var ret bool
		return ret
	}
	return *o.HasAccess
}

// GetHasAccessOk returns a tuple with the HasAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultAccessFeature) GetHasAccessOk() (*bool, bool) {
	if o == nil || o.HasAccess == nil {
		return nil, false
	}
	return o.HasAccess, true
}

// HasHasAccess returns a boolean if a field has been set.
func (o *DefaultAccessFeature) HasHasAccess() bool {
	if o != nil && o.HasAccess != nil {
		return true
	}

	return false
}

// SetHasAccess gets a reference to the given bool and assigns it to the HasAccess field.
func (o *DefaultAccessFeature) SetHasAccess(v bool) {
	o.HasAccess = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DefaultAccessFeature) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultAccessFeature) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DefaultAccessFeature) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DefaultAccessFeature) SetName(v string) {
	o.Name = &v
}

func (o DefaultAccessFeature) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HasAccess != nil {
		toSerialize["has_access"] = o.HasAccess
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDefaultAccessFeature struct {
	value *DefaultAccessFeature
	isSet bool
}

func (v NullableDefaultAccessFeature) Get() *DefaultAccessFeature {
	return v.value
}

func (v *NullableDefaultAccessFeature) Set(val *DefaultAccessFeature) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultAccessFeature) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultAccessFeature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultAccessFeature(val *DefaultAccessFeature) *NullableDefaultAccessFeature {
	return &NullableDefaultAccessFeature{value: val, isSet: true}
}

func (v NullableDefaultAccessFeature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultAccessFeature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
