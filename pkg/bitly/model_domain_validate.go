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

// DomainValidate custom domain in validation queue
type DomainValidate struct {
	CustomDomain     *string `json:"custom_domain,omitempty"`
	OrganizationGuid *string `json:"organization_guid,omitempty"`
}

// NewDomainValidate instantiates a new DomainValidate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainValidate() *DomainValidate {
	this := DomainValidate{}
	return &this
}

// NewDomainValidateWithDefaults instantiates a new DomainValidate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainValidateWithDefaults() *DomainValidate {
	this := DomainValidate{}
	return &this
}

// GetCustomDomain returns the CustomDomain field value if set, zero value otherwise.
func (o *DomainValidate) GetCustomDomain() string {
	if o == nil || o.CustomDomain == nil {
		var ret string
		return ret
	}
	return *o.CustomDomain
}

// GetCustomDomainOk returns a tuple with the CustomDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainValidate) GetCustomDomainOk() (*string, bool) {
	if o == nil || o.CustomDomain == nil {
		return nil, false
	}
	return o.CustomDomain, true
}

// HasCustomDomain returns a boolean if a field has been set.
func (o *DomainValidate) HasCustomDomain() bool {
	if o != nil && o.CustomDomain != nil {
		return true
	}

	return false
}

// SetCustomDomain gets a reference to the given string and assigns it to the CustomDomain field.
func (o *DomainValidate) SetCustomDomain(v string) {
	o.CustomDomain = &v
}

// GetOrganizationGuid returns the OrganizationGuid field value if set, zero value otherwise.
func (o *DomainValidate) GetOrganizationGuid() string {
	if o == nil || o.OrganizationGuid == nil {
		var ret string
		return ret
	}
	return *o.OrganizationGuid
}

// GetOrganizationGuidOk returns a tuple with the OrganizationGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainValidate) GetOrganizationGuidOk() (*string, bool) {
	if o == nil || o.OrganizationGuid == nil {
		return nil, false
	}
	return o.OrganizationGuid, true
}

// HasOrganizationGuid returns a boolean if a field has been set.
func (o *DomainValidate) HasOrganizationGuid() bool {
	if o != nil && o.OrganizationGuid != nil {
		return true
	}

	return false
}

// SetOrganizationGuid gets a reference to the given string and assigns it to the OrganizationGuid field.
func (o *DomainValidate) SetOrganizationGuid(v string) {
	o.OrganizationGuid = &v
}

func (o DomainValidate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomDomain != nil {
		toSerialize["custom_domain"] = o.CustomDomain
	}
	if o.OrganizationGuid != nil {
		toSerialize["organization_guid"] = o.OrganizationGuid
	}
	return json.Marshal(toSerialize)
}

type NullableDomainValidate struct {
	value *DomainValidate
	isSet bool
}

func (v NullableDomainValidate) Get() *DomainValidate {
	return v.value
}

func (v *NullableDomainValidate) Set(val *DomainValidate) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainValidate) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainValidate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainValidate(val *DomainValidate) *NullableDomainValidate {
	return &NullableDomainValidate{value: val, isSet: true}
}

func (v NullableDomainValidate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainValidate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
