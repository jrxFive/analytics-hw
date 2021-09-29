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

// Invitation struct for Invitation
type Invitation struct {
	RoleName string             `json:"role_name"`
	Email    string             `json:"email"`
	Groups   *[]InvitationGroup `json:"groups,omitempty"`
	Created  *string            `json:"created,omitempty"`
}

// NewInvitation instantiates a new Invitation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitation(roleName string, email string) *Invitation {
	this := Invitation{}
	this.RoleName = roleName
	this.Email = email
	return &this
}

// NewInvitationWithDefaults instantiates a new Invitation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationWithDefaults() *Invitation {
	this := Invitation{}
	return &this
}

// GetRoleName returns the RoleName field value
func (o *Invitation) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *Invitation) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value
func (o *Invitation) SetRoleName(v string) {
	o.RoleName = v
}

// GetEmail returns the Email field value
func (o *Invitation) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *Invitation) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *Invitation) SetEmail(v string) {
	o.Email = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *Invitation) GetGroups() []InvitationGroup {
	if o == nil || o.Groups == nil {
		var ret []InvitationGroup
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetGroupsOk() (*[]InvitationGroup, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *Invitation) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []InvitationGroup and assigns it to the Groups field.
func (o *Invitation) SetGroups(v []InvitationGroup) {
	o.Groups = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Invitation) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Invitation) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *Invitation) SetCreated(v string) {
	o.Created = &v
}

func (o Invitation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["role_name"] = o.RoleName
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	return json.Marshal(toSerialize)
}

type NullableInvitation struct {
	value *Invitation
	isSet bool
}

func (v NullableInvitation) Get() *Invitation {
	return v.value
}

func (v *NullableInvitation) Set(val *Invitation) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitation) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitation(val *Invitation) *NullableInvitation {
	return &NullableInvitation{value: val, isSet: true}
}

func (v NullableInvitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
