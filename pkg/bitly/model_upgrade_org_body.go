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

// UpgradeOrgBody struct for UpgradeOrgBody
type UpgradeOrgBody struct {
	BillingInfo     *BillingInfo `json:"billing_info,omitempty"`
	PaymentProvider *string      `json:"payment_provider,omitempty"`
	StripeToken     *string      `json:"stripe_token,omitempty"`
	RatePlanName    *string      `json:"rate_plan_name,omitempty"`
	OrgGuid         *string      `json:"org_guid,omitempty"`
	PaymentMethodId *string      `json:"payment_method_id,omitempty"`
}

// NewUpgradeOrgBody instantiates a new UpgradeOrgBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpgradeOrgBody() *UpgradeOrgBody {
	this := UpgradeOrgBody{}
	var paymentProvider string = "zuora"
	this.PaymentProvider = &paymentProvider
	return &this
}

// NewUpgradeOrgBodyWithDefaults instantiates a new UpgradeOrgBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpgradeOrgBodyWithDefaults() *UpgradeOrgBody {
	this := UpgradeOrgBody{}
	var paymentProvider string = "zuora"
	this.PaymentProvider = &paymentProvider
	return &this
}

// GetBillingInfo returns the BillingInfo field value if set, zero value otherwise.
func (o *UpgradeOrgBody) GetBillingInfo() BillingInfo {
	if o == nil || o.BillingInfo == nil {
		var ret BillingInfo
		return ret
	}
	return *o.BillingInfo
}

// GetBillingInfoOk returns a tuple with the BillingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeOrgBody) GetBillingInfoOk() (*BillingInfo, bool) {
	if o == nil || o.BillingInfo == nil {
		return nil, false
	}
	return o.BillingInfo, true
}

// HasBillingInfo returns a boolean if a field has been set.
func (o *UpgradeOrgBody) HasBillingInfo() bool {
	if o != nil && o.BillingInfo != nil {
		return true
	}

	return false
}

// SetBillingInfo gets a reference to the given BillingInfo and assigns it to the BillingInfo field.
func (o *UpgradeOrgBody) SetBillingInfo(v BillingInfo) {
	o.BillingInfo = &v
}

// GetPaymentProvider returns the PaymentProvider field value if set, zero value otherwise.
func (o *UpgradeOrgBody) GetPaymentProvider() string {
	if o == nil || o.PaymentProvider == nil {
		var ret string
		return ret
	}
	return *o.PaymentProvider
}

// GetPaymentProviderOk returns a tuple with the PaymentProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeOrgBody) GetPaymentProviderOk() (*string, bool) {
	if o == nil || o.PaymentProvider == nil {
		return nil, false
	}
	return o.PaymentProvider, true
}

// HasPaymentProvider returns a boolean if a field has been set.
func (o *UpgradeOrgBody) HasPaymentProvider() bool {
	if o != nil && o.PaymentProvider != nil {
		return true
	}

	return false
}

// SetPaymentProvider gets a reference to the given string and assigns it to the PaymentProvider field.
func (o *UpgradeOrgBody) SetPaymentProvider(v string) {
	o.PaymentProvider = &v
}

// GetStripeToken returns the StripeToken field value if set, zero value otherwise.
func (o *UpgradeOrgBody) GetStripeToken() string {
	if o == nil || o.StripeToken == nil {
		var ret string
		return ret
	}
	return *o.StripeToken
}

// GetStripeTokenOk returns a tuple with the StripeToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeOrgBody) GetStripeTokenOk() (*string, bool) {
	if o == nil || o.StripeToken == nil {
		return nil, false
	}
	return o.StripeToken, true
}

// HasStripeToken returns a boolean if a field has been set.
func (o *UpgradeOrgBody) HasStripeToken() bool {
	if o != nil && o.StripeToken != nil {
		return true
	}

	return false
}

// SetStripeToken gets a reference to the given string and assigns it to the StripeToken field.
func (o *UpgradeOrgBody) SetStripeToken(v string) {
	o.StripeToken = &v
}

// GetRatePlanName returns the RatePlanName field value if set, zero value otherwise.
func (o *UpgradeOrgBody) GetRatePlanName() string {
	if o == nil || o.RatePlanName == nil {
		var ret string
		return ret
	}
	return *o.RatePlanName
}

// GetRatePlanNameOk returns a tuple with the RatePlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeOrgBody) GetRatePlanNameOk() (*string, bool) {
	if o == nil || o.RatePlanName == nil {
		return nil, false
	}
	return o.RatePlanName, true
}

// HasRatePlanName returns a boolean if a field has been set.
func (o *UpgradeOrgBody) HasRatePlanName() bool {
	if o != nil && o.RatePlanName != nil {
		return true
	}

	return false
}

// SetRatePlanName gets a reference to the given string and assigns it to the RatePlanName field.
func (o *UpgradeOrgBody) SetRatePlanName(v string) {
	o.RatePlanName = &v
}

// GetOrgGuid returns the OrgGuid field value if set, zero value otherwise.
func (o *UpgradeOrgBody) GetOrgGuid() string {
	if o == nil || o.OrgGuid == nil {
		var ret string
		return ret
	}
	return *o.OrgGuid
}

// GetOrgGuidOk returns a tuple with the OrgGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeOrgBody) GetOrgGuidOk() (*string, bool) {
	if o == nil || o.OrgGuid == nil {
		return nil, false
	}
	return o.OrgGuid, true
}

// HasOrgGuid returns a boolean if a field has been set.
func (o *UpgradeOrgBody) HasOrgGuid() bool {
	if o != nil && o.OrgGuid != nil {
		return true
	}

	return false
}

// SetOrgGuid gets a reference to the given string and assigns it to the OrgGuid field.
func (o *UpgradeOrgBody) SetOrgGuid(v string) {
	o.OrgGuid = &v
}

// GetPaymentMethodId returns the PaymentMethodId field value if set, zero value otherwise.
func (o *UpgradeOrgBody) GetPaymentMethodId() string {
	if o == nil || o.PaymentMethodId == nil {
		var ret string
		return ret
	}
	return *o.PaymentMethodId
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeOrgBody) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil || o.PaymentMethodId == nil {
		return nil, false
	}
	return o.PaymentMethodId, true
}

// HasPaymentMethodId returns a boolean if a field has been set.
func (o *UpgradeOrgBody) HasPaymentMethodId() bool {
	if o != nil && o.PaymentMethodId != nil {
		return true
	}

	return false
}

// SetPaymentMethodId gets a reference to the given string and assigns it to the PaymentMethodId field.
func (o *UpgradeOrgBody) SetPaymentMethodId(v string) {
	o.PaymentMethodId = &v
}

func (o UpgradeOrgBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BillingInfo != nil {
		toSerialize["billing_info"] = o.BillingInfo
	}
	if o.PaymentProvider != nil {
		toSerialize["payment_provider"] = o.PaymentProvider
	}
	if o.StripeToken != nil {
		toSerialize["stripe_token"] = o.StripeToken
	}
	if o.RatePlanName != nil {
		toSerialize["rate_plan_name"] = o.RatePlanName
	}
	if o.OrgGuid != nil {
		toSerialize["org_guid"] = o.OrgGuid
	}
	if o.PaymentMethodId != nil {
		toSerialize["payment_method_id"] = o.PaymentMethodId
	}
	return json.Marshal(toSerialize)
}

type NullableUpgradeOrgBody struct {
	value *UpgradeOrgBody
	isSet bool
}

func (v NullableUpgradeOrgBody) Get() *UpgradeOrgBody {
	return v.value
}

func (v *NullableUpgradeOrgBody) Set(val *UpgradeOrgBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpgradeOrgBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpgradeOrgBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpgradeOrgBody(val *UpgradeOrgBody) *NullableUpgradeOrgBody {
	return &NullableUpgradeOrgBody{value: val, isSet: true}
}

func (v NullableUpgradeOrgBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpgradeOrgBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
