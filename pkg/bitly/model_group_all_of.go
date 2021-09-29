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

// GroupAllOf struct for GroupAllOf
type GroupAllOf struct {
	Name             *string   `json:"name,omitempty"`
	Bsds             *[]string `json:"bsds,omitempty"`
	Created          *string   `json:"created,omitempty"`
	IsActive         *bool     `json:"is_active,omitempty"`
	Modified         *string   `json:"modified,omitempty"`
	OrganizationGuid *string   `json:"organization_guid,omitempty"`
	Role             *string   `json:"role,omitempty"`
	Guid             *string   `json:"guid,omitempty"`
}

// NewGroupAllOf instantiates a new GroupAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupAllOf() *GroupAllOf {
	this := GroupAllOf{}
	return &this
}

// NewGroupAllOfWithDefaults instantiates a new GroupAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupAllOfWithDefaults() *GroupAllOf {
	this := GroupAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GroupAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupAllOf) SetName(v string) {
	o.Name = &v
}

// GetBsds returns the Bsds field value if set, zero value otherwise.
func (o *GroupAllOf) GetBsds() []string {
	if o == nil || o.Bsds == nil {
		var ret []string
		return ret
	}
	return *o.Bsds
}

// GetBsdsOk returns a tuple with the Bsds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupAllOf) GetBsdsOk() (*[]string, bool) {
	if o == nil || o.Bsds == nil {
		return nil, false
	}
	return o.Bsds, true
}

// HasBsds returns a boolean if a field has been set.
func (o *GroupAllOf) HasBsds() bool {
	if o != nil && o.Bsds != nil {
		return true
	}

	return false
}

// SetBsds gets a reference to the given []string and assigns it to the Bsds field.
func (o *GroupAllOf) SetBsds(v []string) {
	o.Bsds = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *GroupAllOf) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupAllOf) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *GroupAllOf) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *GroupAllOf) SetCreated(v string) {
	o.Created = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *GroupAllOf) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupAllOf) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *GroupAllOf) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *GroupAllOf) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *GroupAllOf) GetModified() string {
	if o == nil || o.Modified == nil {
		var ret string
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupAllOf) GetModifiedOk() (*string, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *GroupAllOf) HasModified() bool {
	if o != nil && o.Modified != nil {
		return true
	}

	return false
}

// SetModified gets a reference to the given string and assigns it to the Modified field.
func (o *GroupAllOf) SetModified(v string) {
	o.Modified = &v
}

// GetOrganizationGuid returns the OrganizationGuid field value if set, zero value otherwise.
func (o *GroupAllOf) GetOrganizationGuid() string {
	if o == nil || o.OrganizationGuid == nil {
		var ret string
		return ret
	}
	return *o.OrganizationGuid
}

// GetOrganizationGuidOk returns a tuple with the OrganizationGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupAllOf) GetOrganizationGuidOk() (*string, bool) {
	if o == nil || o.OrganizationGuid == nil {
		return nil, false
	}
	return o.OrganizationGuid, true
}

// HasOrganizationGuid returns a boolean if a field has been set.
func (o *GroupAllOf) HasOrganizationGuid() bool {
	if o != nil && o.OrganizationGuid != nil {
		return true
	}

	return false
}

// SetOrganizationGuid gets a reference to the given string and assigns it to the OrganizationGuid field.
func (o *GroupAllOf) SetOrganizationGuid(v string) {
	o.OrganizationGuid = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *GroupAllOf) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupAllOf) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *GroupAllOf) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *GroupAllOf) SetRole(v string) {
	o.Role = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *GroupAllOf) GetGuid() string {
	if o == nil || o.Guid == nil {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupAllOf) GetGuidOk() (*string, bool) {
	if o == nil || o.Guid == nil {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *GroupAllOf) HasGuid() bool {
	if o != nil && o.Guid != nil {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *GroupAllOf) SetGuid(v string) {
	o.Guid = &v
}

func (o GroupAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Bsds != nil {
		toSerialize["bsds"] = o.Bsds
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.IsActive != nil {
		toSerialize["is_active"] = o.IsActive
	}
	if o.Modified != nil {
		toSerialize["modified"] = o.Modified
	}
	if o.OrganizationGuid != nil {
		toSerialize["organization_guid"] = o.OrganizationGuid
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Guid != nil {
		toSerialize["guid"] = o.Guid
	}
	return json.Marshal(toSerialize)
}

type NullableGroupAllOf struct {
	value *GroupAllOf
	isSet bool
}

func (v NullableGroupAllOf) Get() *GroupAllOf {
	return v.value
}

func (v *NullableGroupAllOf) Set(val *GroupAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupAllOf(val *GroupAllOf) *NullableGroupAllOf {
	return &NullableGroupAllOf{value: val, isSet: true}
}

func (v NullableGroupAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
