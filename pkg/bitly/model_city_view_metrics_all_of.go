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

// CityViewMetricsAllOf struct for CityViewMetricsAllOf
type CityViewMetricsAllOf struct {
	LaunchpadViews *[]ViewMetric     `json:"launchpad_views,omitempty"`
	OtherMetrics   *OtherViewMetrics `json:"other_metrics,omitempty"`
}

// NewCityViewMetricsAllOf instantiates a new CityViewMetricsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCityViewMetricsAllOf() *CityViewMetricsAllOf {
	this := CityViewMetricsAllOf{}
	return &this
}

// NewCityViewMetricsAllOfWithDefaults instantiates a new CityViewMetricsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCityViewMetricsAllOfWithDefaults() *CityViewMetricsAllOf {
	this := CityViewMetricsAllOf{}
	return &this
}

// GetLaunchpadViews returns the LaunchpadViews field value if set, zero value otherwise.
func (o *CityViewMetricsAllOf) GetLaunchpadViews() []ViewMetric {
	if o == nil || o.LaunchpadViews == nil {
		var ret []ViewMetric
		return ret
	}
	return *o.LaunchpadViews
}

// GetLaunchpadViewsOk returns a tuple with the LaunchpadViews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CityViewMetricsAllOf) GetLaunchpadViewsOk() (*[]ViewMetric, bool) {
	if o == nil || o.LaunchpadViews == nil {
		return nil, false
	}
	return o.LaunchpadViews, true
}

// HasLaunchpadViews returns a boolean if a field has been set.
func (o *CityViewMetricsAllOf) HasLaunchpadViews() bool {
	if o != nil && o.LaunchpadViews != nil {
		return true
	}

	return false
}

// SetLaunchpadViews gets a reference to the given []ViewMetric and assigns it to the LaunchpadViews field.
func (o *CityViewMetricsAllOf) SetLaunchpadViews(v []ViewMetric) {
	o.LaunchpadViews = &v
}

// GetOtherMetrics returns the OtherMetrics field value if set, zero value otherwise.
func (o *CityViewMetricsAllOf) GetOtherMetrics() OtherViewMetrics {
	if o == nil || o.OtherMetrics == nil {
		var ret OtherViewMetrics
		return ret
	}
	return *o.OtherMetrics
}

// GetOtherMetricsOk returns a tuple with the OtherMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CityViewMetricsAllOf) GetOtherMetricsOk() (*OtherViewMetrics, bool) {
	if o == nil || o.OtherMetrics == nil {
		return nil, false
	}
	return o.OtherMetrics, true
}

// HasOtherMetrics returns a boolean if a field has been set.
func (o *CityViewMetricsAllOf) HasOtherMetrics() bool {
	if o != nil && o.OtherMetrics != nil {
		return true
	}

	return false
}

// SetOtherMetrics gets a reference to the given OtherViewMetrics and assigns it to the OtherMetrics field.
func (o *CityViewMetricsAllOf) SetOtherMetrics(v OtherViewMetrics) {
	o.OtherMetrics = &v
}

func (o CityViewMetricsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LaunchpadViews != nil {
		toSerialize["launchpad_views"] = o.LaunchpadViews
	}
	if o.OtherMetrics != nil {
		toSerialize["other_metrics"] = o.OtherMetrics
	}
	return json.Marshal(toSerialize)
}

type NullableCityViewMetricsAllOf struct {
	value *CityViewMetricsAllOf
	isSet bool
}

func (v NullableCityViewMetricsAllOf) Get() *CityViewMetricsAllOf {
	return v.value
}

func (v *NullableCityViewMetricsAllOf) Set(val *CityViewMetricsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCityViewMetricsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCityViewMetricsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCityViewMetricsAllOf(val *CityViewMetricsAllOf) *NullableCityViewMetricsAllOf {
	return &NullableCityViewMetricsAllOf{value: val, isSet: true}
}

func (v NullableCityViewMetricsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCityViewMetricsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
