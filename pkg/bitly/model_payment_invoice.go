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

// PaymentInvoice struct for PaymentInvoice
type PaymentInvoice struct {
	TotalTax       *float32    `json:"total_tax,omitempty"`
	Description    *string     `json:"description,omitempty"`
	Charges        *Charges    `json:"charges,omitempty"`
	InvoiceId      *string     `json:"invoice_id,omitempty"`
	InvoiceDate    *string     `json:"invoice_date,omitempty"`
	Payments       *[]Payments `json:"payments,omitempty"`
	InvoiceNumber  *string     `json:"invoice_number,omitempty"`
	Total          *float32    `json:"total,omitempty"`
	Subtotal       *float32    `json:"subtotal,omitempty"`
	InvoiceDueDate *string     `json:"invoice_due_date,omitempty"`
}

// NewPaymentInvoice instantiates a new PaymentInvoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInvoice() *PaymentInvoice {
	this := PaymentInvoice{}
	return &this
}

// NewPaymentInvoiceWithDefaults instantiates a new PaymentInvoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInvoiceWithDefaults() *PaymentInvoice {
	this := PaymentInvoice{}
	return &this
}

// GetTotalTax returns the TotalTax field value if set, zero value otherwise.
func (o *PaymentInvoice) GetTotalTax() float32 {
	if o == nil || o.TotalTax == nil {
		var ret float32
		return ret
	}
	return *o.TotalTax
}

// GetTotalTaxOk returns a tuple with the TotalTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInvoice) GetTotalTaxOk() (*float32, bool) {
	if o == nil || o.TotalTax == nil {
		return nil, false
	}
	return o.TotalTax, true
}

// HasTotalTax returns a boolean if a field has been set.
func (o *PaymentInvoice) HasTotalTax() bool {
	if o != nil && o.TotalTax != nil {
		return true
	}

	return false
}

// SetTotalTax gets a reference to the given float32 and assigns it to the TotalTax field.
func (o *PaymentInvoice) SetTotalTax(v float32) {
	o.TotalTax = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PaymentInvoice) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInvoice) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PaymentInvoice) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PaymentInvoice) SetDescription(v string) {
	o.Description = &v
}

// GetCharges returns the Charges field value if set, zero value otherwise.
func (o *PaymentInvoice) GetCharges() Charges {
	if o == nil || o.Charges == nil {
		var ret Charges
		return ret
	}
	return *o.Charges
}

// GetChargesOk returns a tuple with the Charges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInvoice) GetChargesOk() (*Charges, bool) {
	if o == nil || o.Charges == nil {
		return nil, false
	}
	return o.Charges, true
}

// HasCharges returns a boolean if a field has been set.
func (o *PaymentInvoice) HasCharges() bool {
	if o != nil && o.Charges != nil {
		return true
	}

	return false
}

// SetCharges gets a reference to the given Charges and assigns it to the Charges field.
func (o *PaymentInvoice) SetCharges(v Charges) {
	o.Charges = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *PaymentInvoice) GetInvoiceId() string {
	if o == nil || o.InvoiceId == nil {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInvoice) GetInvoiceIdOk() (*string, bool) {
	if o == nil || o.InvoiceId == nil {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *PaymentInvoice) HasInvoiceId() bool {
	if o != nil && o.InvoiceId != nil {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *PaymentInvoice) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetInvoiceDate returns the InvoiceDate field value if set, zero value otherwise.
func (o *PaymentInvoice) GetInvoiceDate() string {
	if o == nil || o.InvoiceDate == nil {
		var ret string
		return ret
	}
	return *o.InvoiceDate
}

// GetInvoiceDateOk returns a tuple with the InvoiceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInvoice) GetInvoiceDateOk() (*string, bool) {
	if o == nil || o.InvoiceDate == nil {
		return nil, false
	}
	return o.InvoiceDate, true
}

// HasInvoiceDate returns a boolean if a field has been set.
func (o *PaymentInvoice) HasInvoiceDate() bool {
	if o != nil && o.InvoiceDate != nil {
		return true
	}

	return false
}

// SetInvoiceDate gets a reference to the given string and assigns it to the InvoiceDate field.
func (o *PaymentInvoice) SetInvoiceDate(v string) {
	o.InvoiceDate = &v
}

// GetPayments returns the Payments field value if set, zero value otherwise.
func (o *PaymentInvoice) GetPayments() []Payments {
	if o == nil || o.Payments == nil {
		var ret []Payments
		return ret
	}
	return *o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInvoice) GetPaymentsOk() (*[]Payments, bool) {
	if o == nil || o.Payments == nil {
		return nil, false
	}
	return o.Payments, true
}

// HasPayments returns a boolean if a field has been set.
func (o *PaymentInvoice) HasPayments() bool {
	if o != nil && o.Payments != nil {
		return true
	}

	return false
}

// SetPayments gets a reference to the given []Payments and assigns it to the Payments field.
func (o *PaymentInvoice) SetPayments(v []Payments) {
	o.Payments = &v
}

// GetInvoiceNumber returns the InvoiceNumber field value if set, zero value otherwise.
func (o *PaymentInvoice) GetInvoiceNumber() string {
	if o == nil || o.InvoiceNumber == nil {
		var ret string
		return ret
	}
	return *o.InvoiceNumber
}

// GetInvoiceNumberOk returns a tuple with the InvoiceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInvoice) GetInvoiceNumberOk() (*string, bool) {
	if o == nil || o.InvoiceNumber == nil {
		return nil, false
	}
	return o.InvoiceNumber, true
}

// HasInvoiceNumber returns a boolean if a field has been set.
func (o *PaymentInvoice) HasInvoiceNumber() bool {
	if o != nil && o.InvoiceNumber != nil {
		return true
	}

	return false
}

// SetInvoiceNumber gets a reference to the given string and assigns it to the InvoiceNumber field.
func (o *PaymentInvoice) SetInvoiceNumber(v string) {
	o.InvoiceNumber = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *PaymentInvoice) GetTotal() float32 {
	if o == nil || o.Total == nil {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInvoice) GetTotalOk() (*float32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *PaymentInvoice) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *PaymentInvoice) SetTotal(v float32) {
	o.Total = &v
}

// GetSubtotal returns the Subtotal field value if set, zero value otherwise.
func (o *PaymentInvoice) GetSubtotal() float32 {
	if o == nil || o.Subtotal == nil {
		var ret float32
		return ret
	}
	return *o.Subtotal
}

// GetSubtotalOk returns a tuple with the Subtotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInvoice) GetSubtotalOk() (*float32, bool) {
	if o == nil || o.Subtotal == nil {
		return nil, false
	}
	return o.Subtotal, true
}

// HasSubtotal returns a boolean if a field has been set.
func (o *PaymentInvoice) HasSubtotal() bool {
	if o != nil && o.Subtotal != nil {
		return true
	}

	return false
}

// SetSubtotal gets a reference to the given float32 and assigns it to the Subtotal field.
func (o *PaymentInvoice) SetSubtotal(v float32) {
	o.Subtotal = &v
}

// GetInvoiceDueDate returns the InvoiceDueDate field value if set, zero value otherwise.
func (o *PaymentInvoice) GetInvoiceDueDate() string {
	if o == nil || o.InvoiceDueDate == nil {
		var ret string
		return ret
	}
	return *o.InvoiceDueDate
}

// GetInvoiceDueDateOk returns a tuple with the InvoiceDueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInvoice) GetInvoiceDueDateOk() (*string, bool) {
	if o == nil || o.InvoiceDueDate == nil {
		return nil, false
	}
	return o.InvoiceDueDate, true
}

// HasInvoiceDueDate returns a boolean if a field has been set.
func (o *PaymentInvoice) HasInvoiceDueDate() bool {
	if o != nil && o.InvoiceDueDate != nil {
		return true
	}

	return false
}

// SetInvoiceDueDate gets a reference to the given string and assigns it to the InvoiceDueDate field.
func (o *PaymentInvoice) SetInvoiceDueDate(v string) {
	o.InvoiceDueDate = &v
}

func (o PaymentInvoice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalTax != nil {
		toSerialize["total_tax"] = o.TotalTax
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Charges != nil {
		toSerialize["charges"] = o.Charges
	}
	if o.InvoiceId != nil {
		toSerialize["invoice_id"] = o.InvoiceId
	}
	if o.InvoiceDate != nil {
		toSerialize["invoice_date"] = o.InvoiceDate
	}
	if o.Payments != nil {
		toSerialize["payments"] = o.Payments
	}
	if o.InvoiceNumber != nil {
		toSerialize["invoice_number"] = o.InvoiceNumber
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Subtotal != nil {
		toSerialize["subtotal"] = o.Subtotal
	}
	if o.InvoiceDueDate != nil {
		toSerialize["invoice_due_date"] = o.InvoiceDueDate
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentInvoice struct {
	value *PaymentInvoice
	isSet bool
}

func (v NullablePaymentInvoice) Get() *PaymentInvoice {
	return v.value
}

func (v *NullablePaymentInvoice) Set(val *PaymentInvoice) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInvoice(val *PaymentInvoice) *NullablePaymentInvoice {
	return &NullablePaymentInvoice{value: val, isSet: true}
}

func (v NullablePaymentInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}