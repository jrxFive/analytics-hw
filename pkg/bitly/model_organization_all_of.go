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

// OrganizationAllOf struct for OrganizationAllOf
type OrganizationAllOf struct {
	Name            *string   `json:"name,omitempty"`
	Bsds            *[]string `json:"bsds,omitempty"`
	Created         *string   `json:"created,omitempty"`
	IsActive        *bool     `json:"is_active,omitempty"`
	Modified        *string   `json:"modified,omitempty"`
	TierDisplayName *string   `json:"tier_display_name,omitempty"`
	TierFamily      *string   `json:"tier_family,omitempty"`
	Tier            *string   `json:"tier,omitempty"`
	Role            *string   `json:"role,omitempty"`
	Guid            *string   `json:"guid,omitempty"`
}

// NewOrganizationAllOf instantiates a new OrganizationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationAllOf() *OrganizationAllOf {
	this := OrganizationAllOf{}
	return &this
}

// NewOrganizationAllOfWithDefaults instantiates a new OrganizationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationAllOfWithDefaults() *OrganizationAllOf {
	this := OrganizationAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationAllOf) SetName(v string) {
	o.Name = &v
}

// GetBsds returns the Bsds field value if set, zero value otherwise.
func (o *OrganizationAllOf) GetBsds() []string {
	if o == nil || o.Bsds == nil {
		var ret []string
		return ret
	}
	return *o.Bsds
}

// GetBsdsOk returns a tuple with the Bsds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAllOf) GetBsdsOk() (*[]string, bool) {
	if o == nil || o.Bsds == nil {
		return nil, false
	}
	return o.Bsds, true
}

// HasBsds returns a boolean if a field has been set.
func (o *OrganizationAllOf) HasBsds() bool {
	if o != nil && o.Bsds != nil {
		return true
	}

	return false
}

// SetBsds gets a reference to the given []string and assigns it to the Bsds field.
func (o *OrganizationAllOf) SetBsds(v []string) {
	o.Bsds = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *OrganizationAllOf) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAllOf) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *OrganizationAllOf) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *OrganizationAllOf) SetCreated(v string) {
	o.Created = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *OrganizationAllOf) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAllOf) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *OrganizationAllOf) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *OrganizationAllOf) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *OrganizationAllOf) GetModified() string {
	if o == nil || o.Modified == nil {
		var ret string
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAllOf) GetModifiedOk() (*string, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *OrganizationAllOf) HasModified() bool {
	if o != nil && o.Modified != nil {
		return true
	}

	return false
}

// SetModified gets a reference to the given string and assigns it to the Modified field.
func (o *OrganizationAllOf) SetModified(v string) {
	o.Modified = &v
}

// GetTierDisplayName returns the TierDisplayName field value if set, zero value otherwise.
func (o *OrganizationAllOf) GetTierDisplayName() string {
	if o == nil || o.TierDisplayName == nil {
		var ret string
		return ret
	}
	return *o.TierDisplayName
}

// GetTierDisplayNameOk returns a tuple with the TierDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAllOf) GetTierDisplayNameOk() (*string, bool) {
	if o == nil || o.TierDisplayName == nil {
		return nil, false
	}
	return o.TierDisplayName, true
}

// HasTierDisplayName returns a boolean if a field has been set.
func (o *OrganizationAllOf) HasTierDisplayName() bool {
	if o != nil && o.TierDisplayName != nil {
		return true
	}

	return false
}

// SetTierDisplayName gets a reference to the given string and assigns it to the TierDisplayName field.
func (o *OrganizationAllOf) SetTierDisplayName(v string) {
	o.TierDisplayName = &v
}

// GetTierFamily returns the TierFamily field value if set, zero value otherwise.
func (o *OrganizationAllOf) GetTierFamily() string {
	if o == nil || o.TierFamily == nil {
		var ret string
		return ret
	}
	return *o.TierFamily
}

// GetTierFamilyOk returns a tuple with the TierFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAllOf) GetTierFamilyOk() (*string, bool) {
	if o == nil || o.TierFamily == nil {
		return nil, false
	}
	return o.TierFamily, true
}

// HasTierFamily returns a boolean if a field has been set.
func (o *OrganizationAllOf) HasTierFamily() bool {
	if o != nil && o.TierFamily != nil {
		return true
	}

	return false
}

// SetTierFamily gets a reference to the given string and assigns it to the TierFamily field.
func (o *OrganizationAllOf) SetTierFamily(v string) {
	o.TierFamily = &v
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *OrganizationAllOf) GetTier() string {
	if o == nil || o.Tier == nil {
		var ret string
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAllOf) GetTierOk() (*string, bool) {
	if o == nil || o.Tier == nil {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *OrganizationAllOf) HasTier() bool {
	if o != nil && o.Tier != nil {
		return true
	}

	return false
}

// SetTier gets a reference to the given string and assigns it to the Tier field.
func (o *OrganizationAllOf) SetTier(v string) {
	o.Tier = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *OrganizationAllOf) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAllOf) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *OrganizationAllOf) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *OrganizationAllOf) SetRole(v string) {
	o.Role = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *OrganizationAllOf) GetGuid() string {
	if o == nil || o.Guid == nil {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAllOf) GetGuidOk() (*string, bool) {
	if o == nil || o.Guid == nil {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *OrganizationAllOf) HasGuid() bool {
	if o != nil && o.Guid != nil {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *OrganizationAllOf) SetGuid(v string) {
	o.Guid = &v
}

func (o OrganizationAllOf) MarshalJSON() ([]byte, error) {
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
	if o.TierDisplayName != nil {
		toSerialize["tier_display_name"] = o.TierDisplayName
	}
	if o.TierFamily != nil {
		toSerialize["tier_family"] = o.TierFamily
	}
	if o.Tier != nil {
		toSerialize["tier"] = o.Tier
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Guid != nil {
		toSerialize["guid"] = o.Guid
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationAllOf struct {
	value *OrganizationAllOf
	isSet bool
}

func (v NullableOrganizationAllOf) Get() *OrganizationAllOf {
	return v.value
}

func (v *NullableOrganizationAllOf) Set(val *OrganizationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationAllOf(val *OrganizationAllOf) *NullableOrganizationAllOf {
	return &NullableOrganizationAllOf{value: val, isSet: true}
}

func (v NullableOrganizationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
