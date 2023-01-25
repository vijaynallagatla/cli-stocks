/*
 * Yahoo Finance
 *
 * Yahoo Finance API specification
 *
 * API version: 1.0.8
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SparkResponseSparkResponse struct for SparkResponseSparkResponse
type SparkResponseSparkResponse struct {
	Meta *SparkResponseSparkMeta `json:"meta,omitempty"`
	Timestamp *[]int32 `json:"timestamp,omitempty"`
	Indicators *SparkResponseSparkIndicators `json:"indicators,omitempty"`
}

// NewSparkResponseSparkResponse instantiates a new SparkResponseSparkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSparkResponseSparkResponse() *SparkResponseSparkResponse {
	this := SparkResponseSparkResponse{}
	return &this
}

// NewSparkResponseSparkResponseWithDefaults instantiates a new SparkResponseSparkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSparkResponseSparkResponseWithDefaults() *SparkResponseSparkResponse {
	this := SparkResponseSparkResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SparkResponseSparkResponse) GetMeta() SparkResponseSparkMeta {
	if o == nil || o.Meta == nil {
		var ret SparkResponseSparkMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SparkResponseSparkResponse) GetMetaOk() (*SparkResponseSparkMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SparkResponseSparkResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given SparkResponseSparkMeta and assigns it to the Meta field.
func (o *SparkResponseSparkResponse) SetMeta(v SparkResponseSparkMeta) {
	o.Meta = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *SparkResponseSparkResponse) GetTimestamp() []int32 {
	if o == nil || o.Timestamp == nil {
		var ret []int32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SparkResponseSparkResponse) GetTimestampOk() (*[]int32, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *SparkResponseSparkResponse) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given []int32 and assigns it to the Timestamp field.
func (o *SparkResponseSparkResponse) SetTimestamp(v []int32) {
	o.Timestamp = &v
}

// GetIndicators returns the Indicators field value if set, zero value otherwise.
func (o *SparkResponseSparkResponse) GetIndicators() SparkResponseSparkIndicators {
	if o == nil || o.Indicators == nil {
		var ret SparkResponseSparkIndicators
		return ret
	}
	return *o.Indicators
}

// GetIndicatorsOk returns a tuple with the Indicators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SparkResponseSparkResponse) GetIndicatorsOk() (*SparkResponseSparkIndicators, bool) {
	if o == nil || o.Indicators == nil {
		return nil, false
	}
	return o.Indicators, true
}

// HasIndicators returns a boolean if a field has been set.
func (o *SparkResponseSparkResponse) HasIndicators() bool {
	if o != nil && o.Indicators != nil {
		return true
	}

	return false
}

// SetIndicators gets a reference to the given SparkResponseSparkIndicators and assigns it to the Indicators field.
func (o *SparkResponseSparkResponse) SetIndicators(v SparkResponseSparkIndicators) {
	o.Indicators = &v
}

func (o SparkResponseSparkResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Indicators != nil {
		toSerialize["indicators"] = o.Indicators
	}
	return json.Marshal(toSerialize)
}

type NullableSparkResponseSparkResponse struct {
	value *SparkResponseSparkResponse
	isSet bool
}

func (v NullableSparkResponseSparkResponse) Get() *SparkResponseSparkResponse {
	return v.value
}

func (v *NullableSparkResponseSparkResponse) Set(val *SparkResponseSparkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSparkResponseSparkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSparkResponseSparkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSparkResponseSparkResponse(val *SparkResponseSparkResponse) *NullableSparkResponseSparkResponse {
	return &NullableSparkResponseSparkResponse{value: val, isSet: true}
}

func (v NullableSparkResponseSparkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSparkResponseSparkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


