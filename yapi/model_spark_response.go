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

// SparkResponse struct for SparkResponse
type SparkResponse struct {
	Spark *SparkResponseSpark `json:"spark,omitempty"`
}

// NewSparkResponse instantiates a new SparkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSparkResponse() *SparkResponse {
	this := SparkResponse{}
	return &this
}

// NewSparkResponseWithDefaults instantiates a new SparkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSparkResponseWithDefaults() *SparkResponse {
	this := SparkResponse{}
	return &this
}

// GetSpark returns the Spark field value if set, zero value otherwise.
func (o *SparkResponse) GetSpark() SparkResponseSpark {
	if o == nil || o.Spark == nil {
		var ret SparkResponseSpark
		return ret
	}
	return *o.Spark
}

// GetSparkOk returns a tuple with the Spark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SparkResponse) GetSparkOk() (*SparkResponseSpark, bool) {
	if o == nil || o.Spark == nil {
		return nil, false
	}
	return o.Spark, true
}

// HasSpark returns a boolean if a field has been set.
func (o *SparkResponse) HasSpark() bool {
	if o != nil && o.Spark != nil {
		return true
	}

	return false
}

// SetSpark gets a reference to the given SparkResponseSpark and assigns it to the Spark field.
func (o *SparkResponse) SetSpark(v SparkResponseSpark) {
	o.Spark = &v
}

func (o SparkResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Spark != nil {
		toSerialize["spark"] = o.Spark
	}
	return json.Marshal(toSerialize)
}

type NullableSparkResponse struct {
	value *SparkResponse
	isSet bool
}

func (v NullableSparkResponse) Get() *SparkResponse {
	return v.value
}

func (v *NullableSparkResponse) Set(val *SparkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSparkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSparkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSparkResponse(val *SparkResponse) *NullableSparkResponse {
	return &NullableSparkResponse{value: val, isSet: true}
}

func (v NullableSparkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSparkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


