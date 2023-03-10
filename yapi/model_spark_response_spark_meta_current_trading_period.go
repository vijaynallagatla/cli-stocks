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

// SparkResponseSparkMetaCurrentTradingPeriod struct for SparkResponseSparkMetaCurrentTradingPeriod
type SparkResponseSparkMetaCurrentTradingPeriod struct {
	Pre *SparkResponseSparkMetaCurrentTradingPeriodPre `json:"pre,omitempty"`
	Regular *SparkResponseSparkMetaCurrentTradingPeriodPre `json:"regular,omitempty"`
	Post *SparkResponseSparkMetaCurrentTradingPeriodPre `json:"post,omitempty"`
}

// NewSparkResponseSparkMetaCurrentTradingPeriod instantiates a new SparkResponseSparkMetaCurrentTradingPeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSparkResponseSparkMetaCurrentTradingPeriod() *SparkResponseSparkMetaCurrentTradingPeriod {
	this := SparkResponseSparkMetaCurrentTradingPeriod{}
	return &this
}

// NewSparkResponseSparkMetaCurrentTradingPeriodWithDefaults instantiates a new SparkResponseSparkMetaCurrentTradingPeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSparkResponseSparkMetaCurrentTradingPeriodWithDefaults() *SparkResponseSparkMetaCurrentTradingPeriod {
	this := SparkResponseSparkMetaCurrentTradingPeriod{}
	return &this
}

// GetPre returns the Pre field value if set, zero value otherwise.
func (o *SparkResponseSparkMetaCurrentTradingPeriod) GetPre() SparkResponseSparkMetaCurrentTradingPeriodPre {
	if o == nil || o.Pre == nil {
		var ret SparkResponseSparkMetaCurrentTradingPeriodPre
		return ret
	}
	return *o.Pre
}

// GetPreOk returns a tuple with the Pre field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SparkResponseSparkMetaCurrentTradingPeriod) GetPreOk() (*SparkResponseSparkMetaCurrentTradingPeriodPre, bool) {
	if o == nil || o.Pre == nil {
		return nil, false
	}
	return o.Pre, true
}

// HasPre returns a boolean if a field has been set.
func (o *SparkResponseSparkMetaCurrentTradingPeriod) HasPre() bool {
	if o != nil && o.Pre != nil {
		return true
	}

	return false
}

// SetPre gets a reference to the given SparkResponseSparkMetaCurrentTradingPeriodPre and assigns it to the Pre field.
func (o *SparkResponseSparkMetaCurrentTradingPeriod) SetPre(v SparkResponseSparkMetaCurrentTradingPeriodPre) {
	o.Pre = &v
}

// GetRegular returns the Regular field value if set, zero value otherwise.
func (o *SparkResponseSparkMetaCurrentTradingPeriod) GetRegular() SparkResponseSparkMetaCurrentTradingPeriodPre {
	if o == nil || o.Regular == nil {
		var ret SparkResponseSparkMetaCurrentTradingPeriodPre
		return ret
	}
	return *o.Regular
}

// GetRegularOk returns a tuple with the Regular field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SparkResponseSparkMetaCurrentTradingPeriod) GetRegularOk() (*SparkResponseSparkMetaCurrentTradingPeriodPre, bool) {
	if o == nil || o.Regular == nil {
		return nil, false
	}
	return o.Regular, true
}

// HasRegular returns a boolean if a field has been set.
func (o *SparkResponseSparkMetaCurrentTradingPeriod) HasRegular() bool {
	if o != nil && o.Regular != nil {
		return true
	}

	return false
}

// SetRegular gets a reference to the given SparkResponseSparkMetaCurrentTradingPeriodPre and assigns it to the Regular field.
func (o *SparkResponseSparkMetaCurrentTradingPeriod) SetRegular(v SparkResponseSparkMetaCurrentTradingPeriodPre) {
	o.Regular = &v
}

// GetPost returns the Post field value if set, zero value otherwise.
func (o *SparkResponseSparkMetaCurrentTradingPeriod) GetPost() SparkResponseSparkMetaCurrentTradingPeriodPre {
	if o == nil || o.Post == nil {
		var ret SparkResponseSparkMetaCurrentTradingPeriodPre
		return ret
	}
	return *o.Post
}

// GetPostOk returns a tuple with the Post field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SparkResponseSparkMetaCurrentTradingPeriod) GetPostOk() (*SparkResponseSparkMetaCurrentTradingPeriodPre, bool) {
	if o == nil || o.Post == nil {
		return nil, false
	}
	return o.Post, true
}

// HasPost returns a boolean if a field has been set.
func (o *SparkResponseSparkMetaCurrentTradingPeriod) HasPost() bool {
	if o != nil && o.Post != nil {
		return true
	}

	return false
}

// SetPost gets a reference to the given SparkResponseSparkMetaCurrentTradingPeriodPre and assigns it to the Post field.
func (o *SparkResponseSparkMetaCurrentTradingPeriod) SetPost(v SparkResponseSparkMetaCurrentTradingPeriodPre) {
	o.Post = &v
}

func (o SparkResponseSparkMetaCurrentTradingPeriod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pre != nil {
		toSerialize["pre"] = o.Pre
	}
	if o.Regular != nil {
		toSerialize["regular"] = o.Regular
	}
	if o.Post != nil {
		toSerialize["post"] = o.Post
	}
	return json.Marshal(toSerialize)
}

type NullableSparkResponseSparkMetaCurrentTradingPeriod struct {
	value *SparkResponseSparkMetaCurrentTradingPeriod
	isSet bool
}

func (v NullableSparkResponseSparkMetaCurrentTradingPeriod) Get() *SparkResponseSparkMetaCurrentTradingPeriod {
	return v.value
}

func (v *NullableSparkResponseSparkMetaCurrentTradingPeriod) Set(val *SparkResponseSparkMetaCurrentTradingPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableSparkResponseSparkMetaCurrentTradingPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableSparkResponseSparkMetaCurrentTradingPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSparkResponseSparkMetaCurrentTradingPeriod(val *SparkResponseSparkMetaCurrentTradingPeriod) *NullableSparkResponseSparkMetaCurrentTradingPeriod {
	return &NullableSparkResponseSparkMetaCurrentTradingPeriod{value: val, isSet: true}
}

func (v NullableSparkResponseSparkMetaCurrentTradingPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSparkResponseSparkMetaCurrentTradingPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


