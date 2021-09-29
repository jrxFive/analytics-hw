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

// AccessFeature struct for AccessFeature
type AccessFeature struct {
	HasAccess bool   `json:"has_access"`
	Name      string `json:"name"`
	// ISO_TIMESTAMP
	Created string `json:"created"`
	// ISO_TIMESTAMP
	Modified      string `json:"modified"`
	IsTierDefault bool   `json:"is_tier_default"`
	Id            string `json:"id"`
}

// NewAccessFeature instantiates a new AccessFeature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessFeature(hasAccess bool, name string, created string, modified string, isTierDefault bool, id string) *AccessFeature {
	this := AccessFeature{}
	this.HasAccess = hasAccess
	this.Name = name
	this.Created = created
	this.Modified = modified
	this.IsTierDefault = isTierDefault
	this.Id = id
	return &this
}

// NewAccessFeatureWithDefaults instantiates a new AccessFeature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessFeatureWithDefaults() *AccessFeature {
	this := AccessFeature{}
	return &this
}

// GetHasAccess returns the HasAccess field value
func (o *AccessFeature) GetHasAccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasAccess
}

// GetHasAccessOk returns a tuple with the HasAccess field value
// and a boolean to check if the value has been set.
func (o *AccessFeature) GetHasAccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasAccess, true
}

// SetHasAccess sets field value
func (o *AccessFeature) SetHasAccess(v bool) {
	o.HasAccess = v
}

// GetName returns the Name field value
func (o *AccessFeature) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccessFeature) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccessFeature) SetName(v string) {
	o.Name = v
}

// GetCreated returns the Created field value
func (o *AccessFeature) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *AccessFeature) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *AccessFeature) SetCreated(v string) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *AccessFeature) GetModified() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *AccessFeature) GetModifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *AccessFeature) SetModified(v string) {
	o.Modified = v
}

// GetIsTierDefault returns the IsTierDefault field value
func (o *AccessFeature) GetIsTierDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsTierDefault
}

// GetIsTierDefaultOk returns a tuple with the IsTierDefault field value
// and a boolean to check if the value has been set.
func (o *AccessFeature) GetIsTierDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsTierDefault, true
}

// SetIsTierDefault sets field value
func (o *AccessFeature) SetIsTierDefault(v bool) {
	o.IsTierDefault = v
}

// GetId returns the Id field value
func (o *AccessFeature) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccessFeature) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccessFeature) SetId(v string) {
	o.Id = v
}

func (o AccessFeature) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["has_access"] = o.HasAccess
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["modified"] = o.Modified
	}
	if true {
		toSerialize["is_tier_default"] = o.IsTierDefault
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableAccessFeature struct {
	value *AccessFeature
	isSet bool
}

func (v NullableAccessFeature) Get() *AccessFeature {
	return v.value
}

func (v *NullableAccessFeature) Set(val *AccessFeature) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessFeature) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessFeature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessFeature(val *AccessFeature) *NullableAccessFeature {
	return &NullableAccessFeature{value: val, isSet: true}
}

func (v NullableAccessFeature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessFeature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}