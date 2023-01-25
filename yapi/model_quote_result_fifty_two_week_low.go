/*
 * Yahoo Finance
 *
 * Yahoo Finance API specification
 *
 * API version: 1.0.8
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package yapi

import (
	"encoding/json"
)

// QuoteResultFiftyTwoWeekLow struct for QuoteResultFiftyTwoWeekLow
type QuoteResultFiftyTwoWeekLow struct {
	Raw *float32 `json:"raw,omitempty"`
	Fmt *string `json:"fmt,omitempty"`
}

// NewQuoteResultFiftyTwoWeekLow instantiates a new QuoteResultFiftyTwoWeekLow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteResultFiftyTwoWeekLow() *QuoteResultFiftyTwoWeekLow {
	this := QuoteResultFiftyTwoWeekLow{}
	return &this
}

// NewQuoteResultFiftyTwoWeekLowWithDefaults instantiates a new QuoteResultFiftyTwoWeekLow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteResultFiftyTwoWeekLowWithDefaults() *QuoteResultFiftyTwoWeekLow {
	this := QuoteResultFiftyTwoWeekLow{}
	return &this
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *QuoteResultFiftyTwoWeekLow) GetRaw() float32 {
	if o == nil || o.Raw == nil {
		var ret float32
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteResultFiftyTwoWeekLow) GetRawOk() (*float32, bool) {
	if o == nil || o.Raw == nil {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *QuoteResultFiftyTwoWeekLow) HasRaw() bool {
	if o != nil && o.Raw != nil {
		return true
	}

	return false
}

// SetRaw gets a reference to the given float32 and assigns it to the Raw field.
func (o *QuoteResultFiftyTwoWeekLow) SetRaw(v float32) {
	o.Raw = &v
}

// GetFmt returns the Fmt field value if set, zero value otherwise.
func (o *QuoteResultFiftyTwoWeekLow) GetFmt() string {
	if o == nil || o.Fmt == nil {
		var ret string
		return ret
	}
	return *o.Fmt
}

// GetFmtOk returns a tuple with the Fmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteResultFiftyTwoWeekLow) GetFmtOk() (*string, bool) {
	if o == nil || o.Fmt == nil {
		return nil, false
	}
	return o.Fmt, true
}

// HasFmt returns a boolean if a field has been set.
func (o *QuoteResultFiftyTwoWeekLow) HasFmt() bool {
	if o != nil && o.Fmt != nil {
		return true
	}

	return false
}

// SetFmt gets a reference to the given string and assigns it to the Fmt field.
func (o *QuoteResultFiftyTwoWeekLow) SetFmt(v string) {
	o.Fmt = &v
}

func (o QuoteResultFiftyTwoWeekLow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Raw != nil {
		toSerialize["raw"] = o.Raw
	}
	if o.Fmt != nil {
		toSerialize["fmt"] = o.Fmt
	}
	return json.Marshal(toSerialize)
}

type NullableQuoteResultFiftyTwoWeekLow struct {
	value *QuoteResultFiftyTwoWeekLow
	isSet bool
}

func (v NullableQuoteResultFiftyTwoWeekLow) Get() *QuoteResultFiftyTwoWeekLow {
	return v.value
}

func (v *NullableQuoteResultFiftyTwoWeekLow) Set(val *QuoteResultFiftyTwoWeekLow) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteResultFiftyTwoWeekLow) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteResultFiftyTwoWeekLow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteResultFiftyTwoWeekLow(val *QuoteResultFiftyTwoWeekLow) *NullableQuoteResultFiftyTwoWeekLow {
	return &NullableQuoteResultFiftyTwoWeekLow{value: val, isSet: true}
}

func (v NullableQuoteResultFiftyTwoWeekLow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteResultFiftyTwoWeekLow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

