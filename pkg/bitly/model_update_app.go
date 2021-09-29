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

// UpdateApp struct for UpdateApp
type UpdateApp struct {
	IpAllowlist  *[]string `json:"ip_allowlist,omitempty"`
	RedirectUris *[]string `json:"redirect_uris,omitempty"`
	Link         *string   `json:"link,omitempty"`
	Name         *string   `json:"name,omitempty"`
	Description  *string   `json:"description,omitempty"`
}

// NewUpdateApp instantiates a new UpdateApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateApp() *UpdateApp {
	this := UpdateApp{}
	return &this
}

// NewUpdateAppWithDefaults instantiates a new UpdateApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAppWithDefaults() *UpdateApp {
	this := UpdateApp{}
	return &this
}

// GetIpAllowlist returns the IpAllowlist field value if set, zero value otherwise.
func (o *UpdateApp) GetIpAllowlist() []string {
	if o == nil || o.IpAllowlist == nil {
		var ret []string
		return ret
	}
	return *o.IpAllowlist
}

// GetIpAllowlistOk returns a tuple with the IpAllowlist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApp) GetIpAllowlistOk() (*[]string, bool) {
	if o == nil || o.IpAllowlist == nil {
		return nil, false
	}
	return o.IpAllowlist, true
}

// HasIpAllowlist returns a boolean if a field has been set.
func (o *UpdateApp) HasIpAllowlist() bool {
	if o != nil && o.IpAllowlist != nil {
		return true
	}

	return false
}

// SetIpAllowlist gets a reference to the given []string and assigns it to the IpAllowlist field.
func (o *UpdateApp) SetIpAllowlist(v []string) {
	o.IpAllowlist = &v
}

// GetRedirectUris returns the RedirectUris field value if set, zero value otherwise.
func (o *UpdateApp) GetRedirectUris() []string {
	if o == nil || o.RedirectUris == nil {
		var ret []string
		return ret
	}
	return *o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApp) GetRedirectUrisOk() (*[]string, bool) {
	if o == nil || o.RedirectUris == nil {
		return nil, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *UpdateApp) HasRedirectUris() bool {
	if o != nil && o.RedirectUris != nil {
		return true
	}

	return false
}

// SetRedirectUris gets a reference to the given []string and assigns it to the RedirectUris field.
func (o *UpdateApp) SetRedirectUris(v []string) {
	o.RedirectUris = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *UpdateApp) GetLink() string {
	if o == nil || o.Link == nil {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApp) GetLinkOk() (*string, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *UpdateApp) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *UpdateApp) SetLink(v string) {
	o.Link = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateApp) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApp) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateApp) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateApp) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateApp) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApp) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateApp) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateApp) SetDescription(v string) {
	o.Description = &v
}

func (o UpdateApp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IpAllowlist != nil {
		toSerialize["ip_allowlist"] = o.IpAllowlist
	}
	if o.RedirectUris != nil {
		toSerialize["redirect_uris"] = o.RedirectUris
	}
	if o.Link != nil {
		toSerialize["link"] = o.Link
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateApp struct {
	value *UpdateApp
	isSet bool
}

func (v NullableUpdateApp) Get() *UpdateApp {
	return v.value
}

func (v *NullableUpdateApp) Set(val *UpdateApp) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateApp) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateApp(val *UpdateApp) *NullableUpdateApp {
	return &NullableUpdateApp{value: val, isSet: true}
}

func (v NullableUpdateApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
