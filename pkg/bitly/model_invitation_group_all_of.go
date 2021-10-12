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

// InvitationGroupAllOf struct for InvitationGroupAllOf
type InvitationGroupAllOf struct {
	GroupGuid *string `json:"group_guid,omitempty"`
	RoleName  *string `json:"role_name,omitempty"`
	Created   *string `json:"created,omitempty"`
}

// NewInvitationGroupAllOf instantiates a new InvitationGroupAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitationGroupAllOf() *InvitationGroupAllOf {
	this := InvitationGroupAllOf{}
	return &this
}

// NewInvitationGroupAllOfWithDefaults instantiates a new InvitationGroupAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationGroupAllOfWithDefaults() *InvitationGroupAllOf {
	this := InvitationGroupAllOf{}
	return &this
}

// GetGroupGuid returns the GroupGuid field value if set, zero value otherwise.
func (o *InvitationGroupAllOf) GetGroupGuid() string {
	if o == nil || o.GroupGuid == nil {
		var ret string
		return ret
	}
	return *o.GroupGuid
}

// GetGroupGuidOk returns a tuple with the GroupGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationGroupAllOf) GetGroupGuidOk() (*string, bool) {
	if o == nil || o.GroupGuid == nil {
		return nil, false
	}
	return o.GroupGuid, true
}

// HasGroupGuid returns a boolean if a field has been set.
func (o *InvitationGroupAllOf) HasGroupGuid() bool {
	if o != nil && o.GroupGuid != nil {
		return true
	}

	return false
}

// SetGroupGuid gets a reference to the given string and assigns it to the GroupGuid field.
func (o *InvitationGroupAllOf) SetGroupGuid(v string) {
	o.GroupGuid = &v
}

// GetRoleName returns the RoleName field value if set, zero value otherwise.
func (o *InvitationGroupAllOf) GetRoleName() string {
	if o == nil || o.RoleName == nil {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationGroupAllOf) GetRoleNameOk() (*string, bool) {
	if o == nil || o.RoleName == nil {
		return nil, false
	}
	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *InvitationGroupAllOf) HasRoleName() bool {
	if o != nil && o.RoleName != nil {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *InvitationGroupAllOf) SetRoleName(v string) {
	o.RoleName = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *InvitationGroupAllOf) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationGroupAllOf) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *InvitationGroupAllOf) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *InvitationGroupAllOf) SetCreated(v string) {
	o.Created = &v
}

func (o InvitationGroupAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupGuid != nil {
		toSerialize["group_guid"] = o.GroupGuid
	}
	if o.RoleName != nil {
		toSerialize["role_name"] = o.RoleName
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	return json.Marshal(toSerialize)
}

type NullableInvitationGroupAllOf struct {
	value *InvitationGroupAllOf
	isSet bool
}

func (v NullableInvitationGroupAllOf) Get() *InvitationGroupAllOf {
	return v.value
}

func (v *NullableInvitationGroupAllOf) Set(val *InvitationGroupAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitationGroupAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitationGroupAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitationGroupAllOf(val *InvitationGroupAllOf) *NullableInvitationGroupAllOf {
	return &NullableInvitationGroupAllOf{value: val, isSet: true}
}

func (v NullableInvitationGroupAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitationGroupAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
