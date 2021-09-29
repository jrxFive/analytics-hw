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

// LaunchpadLinkPerformance struct for LaunchpadLinkPerformance
type LaunchpadLinkPerformance struct {
	UnitReference    *string            `json:"unit_reference,omitempty"`
	LinkClicks       *[]LaunchpadClicks `json:"link_clicks,omitempty"`
	PerformanceStart *string            `json:"performance_start,omitempty"`
	PerformanceEnd   *string            `json:"performance_end,omitempty"`
	Units            *int32             `json:"units,omitempty"`
	Unit             *string            `json:"unit,omitempty"`
}

// NewLaunchpadLinkPerformance instantiates a new LaunchpadLinkPerformance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLaunchpadLinkPerformance() *LaunchpadLinkPerformance {
	this := LaunchpadLinkPerformance{}
	return &this
}

// NewLaunchpadLinkPerformanceWithDefaults instantiates a new LaunchpadLinkPerformance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLaunchpadLinkPerformanceWithDefaults() *LaunchpadLinkPerformance {
	this := LaunchpadLinkPerformance{}
	return &this
}

// GetUnitReference returns the UnitReference field value if set, zero value otherwise.
func (o *LaunchpadLinkPerformance) GetUnitReference() string {
	if o == nil || o.UnitReference == nil {
		var ret string
		return ret
	}
	return *o.UnitReference
}

// GetUnitReferenceOk returns a tuple with the UnitReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadLinkPerformance) GetUnitReferenceOk() (*string, bool) {
	if o == nil || o.UnitReference == nil {
		return nil, false
	}
	return o.UnitReference, true
}

// HasUnitReference returns a boolean if a field has been set.
func (o *LaunchpadLinkPerformance) HasUnitReference() bool {
	if o != nil && o.UnitReference != nil {
		return true
	}

	return false
}

// SetUnitReference gets a reference to the given string and assigns it to the UnitReference field.
func (o *LaunchpadLinkPerformance) SetUnitReference(v string) {
	o.UnitReference = &v
}

// GetLinkClicks returns the LinkClicks field value if set, zero value otherwise.
func (o *LaunchpadLinkPerformance) GetLinkClicks() []LaunchpadClicks {
	if o == nil || o.LinkClicks == nil {
		var ret []LaunchpadClicks
		return ret
	}
	return *o.LinkClicks
}

// GetLinkClicksOk returns a tuple with the LinkClicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadLinkPerformance) GetLinkClicksOk() (*[]LaunchpadClicks, bool) {
	if o == nil || o.LinkClicks == nil {
		return nil, false
	}
	return o.LinkClicks, true
}

// HasLinkClicks returns a boolean if a field has been set.
func (o *LaunchpadLinkPerformance) HasLinkClicks() bool {
	if o != nil && o.LinkClicks != nil {
		return true
	}

	return false
}

// SetLinkClicks gets a reference to the given []LaunchpadClicks and assigns it to the LinkClicks field.
func (o *LaunchpadLinkPerformance) SetLinkClicks(v []LaunchpadClicks) {
	o.LinkClicks = &v
}

// GetPerformanceStart returns the PerformanceStart field value if set, zero value otherwise.
func (o *LaunchpadLinkPerformance) GetPerformanceStart() string {
	if o == nil || o.PerformanceStart == nil {
		var ret string
		return ret
	}
	return *o.PerformanceStart
}

// GetPerformanceStartOk returns a tuple with the PerformanceStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadLinkPerformance) GetPerformanceStartOk() (*string, bool) {
	if o == nil || o.PerformanceStart == nil {
		return nil, false
	}
	return o.PerformanceStart, true
}

// HasPerformanceStart returns a boolean if a field has been set.
func (o *LaunchpadLinkPerformance) HasPerformanceStart() bool {
	if o != nil && o.PerformanceStart != nil {
		return true
	}

	return false
}

// SetPerformanceStart gets a reference to the given string and assigns it to the PerformanceStart field.
func (o *LaunchpadLinkPerformance) SetPerformanceStart(v string) {
	o.PerformanceStart = &v
}

// GetPerformanceEnd returns the PerformanceEnd field value if set, zero value otherwise.
func (o *LaunchpadLinkPerformance) GetPerformanceEnd() string {
	if o == nil || o.PerformanceEnd == nil {
		var ret string
		return ret
	}
	return *o.PerformanceEnd
}

// GetPerformanceEndOk returns a tuple with the PerformanceEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadLinkPerformance) GetPerformanceEndOk() (*string, bool) {
	if o == nil || o.PerformanceEnd == nil {
		return nil, false
	}
	return o.PerformanceEnd, true
}

// HasPerformanceEnd returns a boolean if a field has been set.
func (o *LaunchpadLinkPerformance) HasPerformanceEnd() bool {
	if o != nil && o.PerformanceEnd != nil {
		return true
	}

	return false
}

// SetPerformanceEnd gets a reference to the given string and assigns it to the PerformanceEnd field.
func (o *LaunchpadLinkPerformance) SetPerformanceEnd(v string) {
	o.PerformanceEnd = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *LaunchpadLinkPerformance) GetUnits() int32 {
	if o == nil || o.Units == nil {
		var ret int32
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadLinkPerformance) GetUnitsOk() (*int32, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *LaunchpadLinkPerformance) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given int32 and assigns it to the Units field.
func (o *LaunchpadLinkPerformance) SetUnits(v int32) {
	o.Units = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *LaunchpadLinkPerformance) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadLinkPerformance) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *LaunchpadLinkPerformance) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *LaunchpadLinkPerformance) SetUnit(v string) {
	o.Unit = &v
}

func (o LaunchpadLinkPerformance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnitReference != nil {
		toSerialize["unit_reference"] = o.UnitReference
	}
	if o.LinkClicks != nil {
		toSerialize["link_clicks"] = o.LinkClicks
	}
	if o.PerformanceStart != nil {
		toSerialize["performance_start"] = o.PerformanceStart
	}
	if o.PerformanceEnd != nil {
		toSerialize["performance_end"] = o.PerformanceEnd
	}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableLaunchpadLinkPerformance struct {
	value *LaunchpadLinkPerformance
	isSet bool
}

func (v NullableLaunchpadLinkPerformance) Get() *LaunchpadLinkPerformance {
	return v.value
}

func (v *NullableLaunchpadLinkPerformance) Set(val *LaunchpadLinkPerformance) {
	v.value = val
	v.isSet = true
}

func (v NullableLaunchpadLinkPerformance) IsSet() bool {
	return v.isSet
}

func (v *NullableLaunchpadLinkPerformance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLaunchpadLinkPerformance(val *LaunchpadLinkPerformance) *NullableLaunchpadLinkPerformance {
	return &NullableLaunchpadLinkPerformance{value: val, isSet: true}
}

func (v NullableLaunchpadLinkPerformance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLaunchpadLinkPerformance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
