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

// CityViewMetrics struct for CityViewMetrics
type CityViewMetrics struct {
	Units          *int32            `json:"units,omitempty"`
	Facet          *string           `json:"facet,omitempty"`
	UnitReference  *string           `json:"unit_reference,omitempty"`
	Unit           *string           `json:"unit,omitempty"`
	LaunchpadViews *[]ViewMetric     `json:"launchpad_views,omitempty"`
	OtherMetrics   *OtherViewMetrics `json:"other_metrics,omitempty"`
}

// NewCityViewMetrics instantiates a new CityViewMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCityViewMetrics() *CityViewMetrics {
	this := CityViewMetrics{}
	return &this
}

// NewCityViewMetricsWithDefaults instantiates a new CityViewMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCityViewMetricsWithDefaults() *CityViewMetrics {
	this := CityViewMetrics{}
	return &this
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *CityViewMetrics) GetUnits() int32 {
	if o == nil || o.Units == nil {
		var ret int32
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CityViewMetrics) GetUnitsOk() (*int32, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *CityViewMetrics) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given int32 and assigns it to the Units field.
func (o *CityViewMetrics) SetUnits(v int32) {
	o.Units = &v
}

// GetFacet returns the Facet field value if set, zero value otherwise.
func (o *CityViewMetrics) GetFacet() string {
	if o == nil || o.Facet == nil {
		var ret string
		return ret
	}
	return *o.Facet
}

// GetFacetOk returns a tuple with the Facet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CityViewMetrics) GetFacetOk() (*string, bool) {
	if o == nil || o.Facet == nil {
		return nil, false
	}
	return o.Facet, true
}

// HasFacet returns a boolean if a field has been set.
func (o *CityViewMetrics) HasFacet() bool {
	if o != nil && o.Facet != nil {
		return true
	}

	return false
}

// SetFacet gets a reference to the given string and assigns it to the Facet field.
func (o *CityViewMetrics) SetFacet(v string) {
	o.Facet = &v
}

// GetUnitReference returns the UnitReference field value if set, zero value otherwise.
func (o *CityViewMetrics) GetUnitReference() string {
	if o == nil || o.UnitReference == nil {
		var ret string
		return ret
	}
	return *o.UnitReference
}

// GetUnitReferenceOk returns a tuple with the UnitReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CityViewMetrics) GetUnitReferenceOk() (*string, bool) {
	if o == nil || o.UnitReference == nil {
		return nil, false
	}
	return o.UnitReference, true
}

// HasUnitReference returns a boolean if a field has been set.
func (o *CityViewMetrics) HasUnitReference() bool {
	if o != nil && o.UnitReference != nil {
		return true
	}

	return false
}

// SetUnitReference gets a reference to the given string and assigns it to the UnitReference field.
func (o *CityViewMetrics) SetUnitReference(v string) {
	o.UnitReference = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *CityViewMetrics) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CityViewMetrics) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *CityViewMetrics) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *CityViewMetrics) SetUnit(v string) {
	o.Unit = &v
}

// GetLaunchpadViews returns the LaunchpadViews field value if set, zero value otherwise.
func (o *CityViewMetrics) GetLaunchpadViews() []ViewMetric {
	if o == nil || o.LaunchpadViews == nil {
		var ret []ViewMetric
		return ret
	}
	return *o.LaunchpadViews
}

// GetLaunchpadViewsOk returns a tuple with the LaunchpadViews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CityViewMetrics) GetLaunchpadViewsOk() (*[]ViewMetric, bool) {
	if o == nil || o.LaunchpadViews == nil {
		return nil, false
	}
	return o.LaunchpadViews, true
}

// HasLaunchpadViews returns a boolean if a field has been set.
func (o *CityViewMetrics) HasLaunchpadViews() bool {
	if o != nil && o.LaunchpadViews != nil {
		return true
	}

	return false
}

// SetLaunchpadViews gets a reference to the given []ViewMetric and assigns it to the LaunchpadViews field.
func (o *CityViewMetrics) SetLaunchpadViews(v []ViewMetric) {
	o.LaunchpadViews = &v
}

// GetOtherMetrics returns the OtherMetrics field value if set, zero value otherwise.
func (o *CityViewMetrics) GetOtherMetrics() OtherViewMetrics {
	if o == nil || o.OtherMetrics == nil {
		var ret OtherViewMetrics
		return ret
	}
	return *o.OtherMetrics
}

// GetOtherMetricsOk returns a tuple with the OtherMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CityViewMetrics) GetOtherMetricsOk() (*OtherViewMetrics, bool) {
	if o == nil || o.OtherMetrics == nil {
		return nil, false
	}
	return o.OtherMetrics, true
}

// HasOtherMetrics returns a boolean if a field has been set.
func (o *CityViewMetrics) HasOtherMetrics() bool {
	if o != nil && o.OtherMetrics != nil {
		return true
	}

	return false
}

// SetOtherMetrics gets a reference to the given OtherViewMetrics and assigns it to the OtherMetrics field.
func (o *CityViewMetrics) SetOtherMetrics(v OtherViewMetrics) {
	o.OtherMetrics = &v
}

func (o CityViewMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	if o.Facet != nil {
		toSerialize["facet"] = o.Facet
	}
	if o.UnitReference != nil {
		toSerialize["unit_reference"] = o.UnitReference
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.LaunchpadViews != nil {
		toSerialize["launchpad_views"] = o.LaunchpadViews
	}
	if o.OtherMetrics != nil {
		toSerialize["other_metrics"] = o.OtherMetrics
	}
	return json.Marshal(toSerialize)
}

type NullableCityViewMetrics struct {
	value *CityViewMetrics
	isSet bool
}

func (v NullableCityViewMetrics) Get() *CityViewMetrics {
	return v.value
}

func (v *NullableCityViewMetrics) Set(val *CityViewMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableCityViewMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableCityViewMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCityViewMetrics(val *CityViewMetrics) *NullableCityViewMetrics {
	return &NullableCityViewMetrics{value: val, isSet: true}
}

func (v NullableCityViewMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCityViewMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
