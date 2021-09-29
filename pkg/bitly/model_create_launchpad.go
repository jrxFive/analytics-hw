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

// CreateLaunchpad struct for CreateLaunchpad
type CreateLaunchpad struct {
	GroupGuid string `json:"group_guid"`
	Domain    string `json:"domain"`
	Keyword   string `json:"keyword"`
}

// NewCreateLaunchpad instantiates a new CreateLaunchpad object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLaunchpad(groupGuid string, domain string, keyword string) *CreateLaunchpad {
	this := CreateLaunchpad{}
	this.GroupGuid = groupGuid
	this.Domain = domain
	this.Keyword = keyword
	return &this
}

// NewCreateLaunchpadWithDefaults instantiates a new CreateLaunchpad object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLaunchpadWithDefaults() *CreateLaunchpad {
	this := CreateLaunchpad{}
	return &this
}

// GetGroupGuid returns the GroupGuid field value
func (o *CreateLaunchpad) GetGroupGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupGuid
}

// GetGroupGuidOk returns a tuple with the GroupGuid field value
// and a boolean to check if the value has been set.
func (o *CreateLaunchpad) GetGroupGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupGuid, true
}

// SetGroupGuid sets field value
func (o *CreateLaunchpad) SetGroupGuid(v string) {
	o.GroupGuid = v
}

// GetDomain returns the Domain field value
func (o *CreateLaunchpad) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *CreateLaunchpad) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *CreateLaunchpad) SetDomain(v string) {
	o.Domain = v
}

// GetKeyword returns the Keyword field value
func (o *CreateLaunchpad) GetKeyword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Keyword
}

// GetKeywordOk returns a tuple with the Keyword field value
// and a boolean to check if the value has been set.
func (o *CreateLaunchpad) GetKeywordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Keyword, true
}

// SetKeyword sets field value
func (o *CreateLaunchpad) SetKeyword(v string) {
	o.Keyword = v
}

func (o CreateLaunchpad) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["group_guid"] = o.GroupGuid
	}
	if true {
		toSerialize["domain"] = o.Domain
	}
	if true {
		toSerialize["keyword"] = o.Keyword
	}
	return json.Marshal(toSerialize)
}

type NullableCreateLaunchpad struct {
	value *CreateLaunchpad
	isSet bool
}

func (v NullableCreateLaunchpad) Get() *CreateLaunchpad {
	return v.value
}

func (v *NullableCreateLaunchpad) Set(val *CreateLaunchpad) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLaunchpad) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLaunchpad) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLaunchpad(val *CreateLaunchpad) *NullableCreateLaunchpad {
	return &NullableCreateLaunchpad{value: val, isSet: true}
}

func (v NullableCreateLaunchpad) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLaunchpad) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
