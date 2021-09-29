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

// CampaignClicks struct for CampaignClicks
type CampaignClicks struct {
	UnitReference *string             `json:"unit_reference,omitempty"`
	Rollup        *bool               `json:"rollup,omitempty"`
	Limit         *int32              `json:"limit,omitempty"`
	Units         *int32              `json:"units,omitempty"`
	Data          *CampaignClicksData `json:"data,omitempty"`
	Unit          *string             `json:"unit,omitempty"`
}

// NewCampaignClicks instantiates a new CampaignClicks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignClicks() *CampaignClicks {
	this := CampaignClicks{}
	return &this
}

// NewCampaignClicksWithDefaults instantiates a new CampaignClicks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignClicksWithDefaults() *CampaignClicks {
	this := CampaignClicks{}
	return &this
}

// GetUnitReference returns the UnitReference field value if set, zero value otherwise.
func (o *CampaignClicks) GetUnitReference() string {
	if o == nil || o.UnitReference == nil {
		var ret string
		return ret
	}
	return *o.UnitReference
}

// GetUnitReferenceOk returns a tuple with the UnitReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignClicks) GetUnitReferenceOk() (*string, bool) {
	if o == nil || o.UnitReference == nil {
		return nil, false
	}
	return o.UnitReference, true
}

// HasUnitReference returns a boolean if a field has been set.
func (o *CampaignClicks) HasUnitReference() bool {
	if o != nil && o.UnitReference != nil {
		return true
	}

	return false
}

// SetUnitReference gets a reference to the given string and assigns it to the UnitReference field.
func (o *CampaignClicks) SetUnitReference(v string) {
	o.UnitReference = &v
}

// GetRollup returns the Rollup field value if set, zero value otherwise.
func (o *CampaignClicks) GetRollup() bool {
	if o == nil || o.Rollup == nil {
		var ret bool
		return ret
	}
	return *o.Rollup
}

// GetRollupOk returns a tuple with the Rollup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignClicks) GetRollupOk() (*bool, bool) {
	if o == nil || o.Rollup == nil {
		return nil, false
	}
	return o.Rollup, true
}

// HasRollup returns a boolean if a field has been set.
func (o *CampaignClicks) HasRollup() bool {
	if o != nil && o.Rollup != nil {
		return true
	}

	return false
}

// SetRollup gets a reference to the given bool and assigns it to the Rollup field.
func (o *CampaignClicks) SetRollup(v bool) {
	o.Rollup = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *CampaignClicks) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignClicks) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *CampaignClicks) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *CampaignClicks) SetLimit(v int32) {
	o.Limit = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *CampaignClicks) GetUnits() int32 {
	if o == nil || o.Units == nil {
		var ret int32
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignClicks) GetUnitsOk() (*int32, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *CampaignClicks) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given int32 and assigns it to the Units field.
func (o *CampaignClicks) SetUnits(v int32) {
	o.Units = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CampaignClicks) GetData() CampaignClicksData {
	if o == nil || o.Data == nil {
		var ret CampaignClicksData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignClicks) GetDataOk() (*CampaignClicksData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CampaignClicks) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given CampaignClicksData and assigns it to the Data field.
func (o *CampaignClicks) SetData(v CampaignClicksData) {
	o.Data = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *CampaignClicks) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignClicks) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *CampaignClicks) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *CampaignClicks) SetUnit(v string) {
	o.Unit = &v
}

func (o CampaignClicks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnitReference != nil {
		toSerialize["unit_reference"] = o.UnitReference
	}
	if o.Rollup != nil {
		toSerialize["rollup"] = o.Rollup
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableCampaignClicks struct {
	value *CampaignClicks
	isSet bool
}

func (v NullableCampaignClicks) Get() *CampaignClicks {
	return v.value
}

func (v *NullableCampaignClicks) Set(val *CampaignClicks) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignClicks) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignClicks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignClicks(val *CampaignClicks) *NullableCampaignClicks {
	return &NullableCampaignClicks{value: val, isSet: true}
}

func (v NullableCampaignClicks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignClicks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
