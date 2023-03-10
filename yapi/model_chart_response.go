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

// ChartResponse struct for ChartResponse
type ChartResponse struct {
	Chart *ChartResponseChart `json:"chart,omitempty"`
}

// NewChartResponse instantiates a new ChartResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChartResponse() *ChartResponse {
	this := ChartResponse{}
	return &this
}

// NewChartResponseWithDefaults instantiates a new ChartResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChartResponseWithDefaults() *ChartResponse {
	this := ChartResponse{}
	return &this
}

// GetChart returns the Chart field value if set, zero value otherwise.
func (o *ChartResponse) GetChart() ChartResponseChart {
	if o == nil || o.Chart == nil {
		var ret ChartResponseChart
		return ret
	}
	return *o.Chart
}

// GetChartOk returns a tuple with the Chart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponse) GetChartOk() (*ChartResponseChart, bool) {
	if o == nil || o.Chart == nil {
		return nil, false
	}
	return o.Chart, true
}

// HasChart returns a boolean if a field has been set.
func (o *ChartResponse) HasChart() bool {
	if o != nil && o.Chart != nil {
		return true
	}

	return false
}

// SetChart gets a reference to the given ChartResponseChart and assigns it to the Chart field.
func (o *ChartResponse) SetChart(v ChartResponseChart) {
	o.Chart = &v
}

func (o ChartResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Chart != nil {
		toSerialize["chart"] = o.Chart
	}
	return json.Marshal(toSerialize)
}

type NullableChartResponse struct {
	value *ChartResponse
	isSet bool
}

func (v NullableChartResponse) Get() *ChartResponse {
	return v.value
}

func (v *NullableChartResponse) Set(val *ChartResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChartResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChartResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChartResponse(val *ChartResponse) *NullableChartResponse {
	return &NullableChartResponse{value: val, isSet: true}
}

func (v NullableChartResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChartResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


