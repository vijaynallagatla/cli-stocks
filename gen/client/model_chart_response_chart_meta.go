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

// ChartResponseChartMeta struct for ChartResponseChartMeta
type ChartResponseChartMeta struct {
	Currency *string `json:"currency,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	ExchangeName *string `json:"exchangeName,omitempty"`
	InstrumentType *string `json:"instrumentType,omitempty"`
	FirstTradeDate *int32 `json:"firstTradeDate,omitempty"`
	RegularMarketTime *int32 `json:"regularMarketTime,omitempty"`
	Gmtoffset *int32 `json:"gmtoffset,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
	ExchangeTimezoneName *string `json:"exchangeTimezoneName,omitempty"`
	RegularMarketPrice *float32 `json:"regularMarketPrice,omitempty"`
	ChartPreviousClose *float32 `json:"chartPreviousClose,omitempty"`
	PreviousClose *float32 `json:"previousClose,omitempty"`
	Scale *int32 `json:"scale,omitempty"`
	PriceHint *int32 `json:"priceHint,omitempty"`
	CurrentTradingPeriod *ChartResponseChartMetaCurrentTradingPeriod `json:"currentTradingPeriod,omitempty"`
	TradingPeriods *[][]map[string]interface{} `json:"tradingPeriods,omitempty"`
	DataGranularity *string `json:"dataGranularity,omitempty"`
	Range *string `json:"range,omitempty"`
	ValidRanges *[]string `json:"validRanges,omitempty"`
}

// NewChartResponseChartMeta instantiates a new ChartResponseChartMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChartResponseChartMeta() *ChartResponseChartMeta {
	this := ChartResponseChartMeta{}
	return &this
}

// NewChartResponseChartMetaWithDefaults instantiates a new ChartResponseChartMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChartResponseChartMetaWithDefaults() *ChartResponseChartMeta {
	this := ChartResponseChartMeta{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ChartResponseChartMeta) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartMeta) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ChartResponseChartMeta) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *ChartResponseChartMeta) SetCurrency(v string) {
	o.Currency = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *ChartResponseChartMeta) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartMeta) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *ChartResponseChartMeta) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *ChartResponseChartMeta) SetSymbol(v string) {
	o.Symbol = &v
}

// GetExchangeName returns the ExchangeName field value if set, zero value otherwise.
func (o *ChartResponseChartMeta) GetExchangeName() string {
	if o == nil || o.ExchangeName == nil {
		var ret string
		return ret
	}
	return *o.ExchangeName
}

// GetExchangeNameOk returns a tuple with the ExchangeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartMeta) GetExchangeNameOk() (*string, bool) {
	if o == nil || o.ExchangeName == nil {
		return nil, false
	}
	return o.ExchangeName, true
}

// HasExchangeName returns a boolean if a field has been set.
func (o *ChartResponseChartMeta) HasExchangeName() bool {
	if o != nil && o.ExchangeName != nil {
		return true
	}

	return false
}

// SetExchangeName gets a reference to the given string and assigns it to the ExchangeName field.
func (o *ChartResponseChartMeta) SetExchangeName(v string) {
	o.ExchangeName = &v
}

// GetInstrumentType returns the InstrumentType field value if set, zero value otherwise.
func (o *ChartResponseChartMeta) GetInstrumentType() string {
	if o == nil || o.InstrumentType == nil {
		var ret string
		return ret
	}
	return *o.InstrumentType
}

// GetInstrumentTypeOk returns a tuple with the InstrumentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartMeta) GetInstrumentTypeOk() (*string, bool) {
	if o == nil || o.InstrumentType == nil {
		return nil, false
	}
	return o.InstrumentType, true
}

// HasInstrumentType returns a boolean if a field has been set.
func (o *ChartResponseChartMeta) HasInstrumentType() bool {
	if o != nil && o.InstrumentType != nil {
		return true
	}

	return false
}

// SetInstrumentType gets a reference to the given string and assigns it to the InstrumentType field.
func (o *ChartResponseChartMeta) SetInstrumentType(v string) {
	o.InstrumentType = &v
}

// GetFirstTradeDate returns the FirstTradeDate field value if set, zero value otherwise.
func (o *ChartResponseChartMeta) GetFirstTradeDate() int32 {
	if o == nil || o.FirstTradeDate == nil {
		var ret int32
		return ret
	}
	return *o.FirstTradeDate
}

// GetFirstTradeDateOk returns a tuple with the FirstTradeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartMeta) GetFirstTradeDateOk() (*int32, bool) {
	if o == nil || o.FirstTradeDate == nil {
		return nil, false
	}
	return o.FirstTradeDate, true
}

// HasFirstTradeDate returns a boolean if a field has been set.
func (o *ChartResponseChartMeta) HasFirstTradeDate() bool {
	if o != nil && o.FirstTradeDate != nil {
		return true
	}

	return false
}

// SetFirstTradeDate gets a reference to the given int32 and assigns it to the FirstTradeDate field.
func (o *ChartResponseChartMeta) SetFirstTradeDate(v int32) {
	o.FirstTradeDate = &v
}

// GetRegularMarketTime returns the RegularMarketTime field value if set, zero value otherwise.
func (o *ChartResponseChartMeta) GetRegularMarketTime() int32 {
	if o == nil || o.RegularMarketTime == nil {
		var ret int32
		return ret
	}
	return *o.RegularMarketTime
}

// GetRegularMarketTimeOk returns a tuple with the RegularMarketTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartMeta) GetRegularMarketTimeOk() (*int32, bool) {
	if o == nil || o.RegularMarketTime == nil {
		return nil, false
	}
	return o.RegularMarketTime, true
}

// HasRegularMarketTime returns a boolean if a field has been set.
func (o *ChartResponseChartMeta) HasRegularMarketTime() bool {
	if o != nil && o.RegularMarketTime != nil {
		return true
	}

	return false
}

// SetRegularMarketTime gets a reference to the given int32 and assigns it to the RegularMarketTime field.
func (o *ChartResponseChartMeta) SetRegularMarketTime(v int32) {
	o.RegularMarketTime = &v
}

// GetGmtoffset returns the Gmtoffset field value if set, zero value otherwise.
func (o *ChartResponseChartMeta) GetGmtoffset() int32 {
	if o == nil || o.Gmtoffset == nil {
		var ret int32
		return ret
	}
	return *o.Gmtoffset
}

// GetGmtoffsetOk returns a tuple with the Gmtoffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartMeta) GetGmtoffsetOk() (*int32, bool) {
	if o == nil || o.Gmtoffset == nil {
		return nil, false
	}
	return o.Gmtoffset, true
}

// HasGmtoffset returns a boolean if a field has been set.
func (o *ChartResponseChartMeta) HasGmtoffset() bool {
	if o != nil && o.Gmtoffset != nil {
		return true
	}

	return false
}

// SetGmtoffset gets a reference to the given int32 and assigns it to the Gmtoffset field.
func (o *ChartResponseChartMeta) SetGmtoffset(v int32) {
	o.Gmtoffset = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *ChartResponseChartMeta) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartMeta) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *ChartResponseChartMeta) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *ChartResponseChartMeta) SetTimezone(v string) {
	o.Timezone = &v
}

// GetExchangeTimezoneName returns the ExchangeTimezoneName field value if set, zero value otherwise.
func (o *ChartResponseChartMeta) GetExchangeTimezoneName() string {
	if o == nil || o.ExchangeTimezoneName == nil {
		var ret string
		return ret
	}
	return *o.ExchangeTimezoneName
}

// GetExchangeTimezoneNameOk returns a tuple with the ExchangeTimezoneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartMeta) GetExchangeTimezoneNameOk() (*string, bool) {
	if o == nil || o.ExchangeTimezoneName == nil {
		return nil, false
	}
	return o.ExchangeTimezoneName, true
}

// HasExchangeTimezoneName returns a boolean if a field has been set.
func (o *ChartResponseChartMeta) HasExchangeTimezoneName() bool {
	if o != nil && o.ExchangeTimezoneName != nil {
		return true
	}

	return false
}

// SetExchangeTimezoneName gets a reference to the given string and assigns it to the ExchangeTimezoneName field.
func (o *ChartResponseChartMeta) SetExchangeTimezoneName(v string) {
	o.ExchangeTimezoneName = &v
}

// GetRegularMarketPrice returns the RegularMarketPrice field value if set, zero value otherwise.
func (o *ChartResponseChartMeta) GetRegularMarketPrice() float32 {
	if o == nil || o.RegularMarketPrice == nil {
		var ret float32
		return ret
	}
	return *o.RegularMarketPrice
}

// GetRegularMarketPriceOk returns a tuple with the RegularMarketPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartMeta) GetRegularMarketPriceOk() (*float32, bool) {
	if o == nil || o.RegularMarketPrice == nil {
		return nil, false
	}
	return o.RegularMarketPrice, true
}

// HasRegularMarketPrice returns a boolean if a field has been set.
func (o *ChartResponseChartMeta) HasRegularMarketPrice() bool {
	if o != nil && o.RegularMarketPrice != nil {
		return true
	}

	return false
}

// SetRegularMarketPrice gets a reference to the given float32 and assigns it to the RegularMarketPrice field.
func (o *ChartResponseChartMeta) SetRegularMarketPrice(v float32) {
	o.RegularMarketPrice = &v
}

// GetChartPreviousClose returns the ChartPreviousClose field value if set, zero value otherwise.
func (o *ChartResponseChartMeta) GetChartPreviousClose() float32 {
	if o == nil || o.ChartPreviousClose == nil {
		var ret float32
		return ret
	}
	return *o.ChartPreviousClose
}

// GetChartPreviousCloseOk returns a tuple with the ChartPreviousClose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartMeta) GetChartPreviousCloseOk() (*float32, bool) {
	if o == nil || o.ChartPreviousClose == nil {
		return nil, false
	}
	return o.ChartPreviousClose, true
}

// HasChartPreviousClose returns a boolean if a field has been set.
func (o *ChartResponseChartMeta) HasChartPreviousClose() bool {
	if o != nil && o.ChartPreviousClose != nil {
		return true
	}

	return false
}

// SetChartPreviousClose gets a reference to the given float32 and assigns it to the ChartPreviousClose field.
func (o *ChartResponseChartMeta) SetChartPreviousClose(v float32) {
	o.ChartPreviousClose = &v
}

// GetPreviousClose returns the PreviousClose field value if set, zero value otherwise.
func (o *ChartResponseChartMeta) GetPreviousClose() float32 {
	if o == nil || o.PreviousClose == nil {
		var ret float32
		return ret
	}
	return *o.PreviousClose
}

// GetPreviousCloseOk returns a tuple with the PreviousClose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartMeta) GetPreviousCloseOk() (*float32, bool) {
	if o == nil || o.PreviousClose == nil {
		return nil, false
	}
	return o.PreviousClose, true
}

// HasPreviousClose returns a boolean if a field has been set.
func (o *ChartResponseChartMeta) HasPreviousClose() bool {
	if o != nil && o.PreviousClose != nil {
		return true
	}

	return false
}

// SetPreviousClose gets a reference to the given float32 and assigns it to the PreviousClose field.
func (o *ChartResponseChartMeta) SetPreviousClose(v float32) {
	o.PreviousClose = &v
}

// GetScale returns the Scale field value if set, zero value otherwise.
func (o *ChartResponseChartMeta) GetScale() int32 {
	if o == nil || o.Scale == nil {
		var ret int32
		return ret
	}
	return *o.Scale
}

// GetScaleOk returns a tuple with the Scale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartMeta) GetScaleOk() (*int32, bool) {
	if o == nil || o.Scale == nil {
		return nil, false
	}
	return o.Scale, true
}

// HasScale returns a boolean if a field has been set.
func (o *ChartResponseChartMeta) HasScale() bool {
	if o != nil && o.Scale != nil {
		return true
	}

	return false
}

// SetScale gets a reference to the given int32 and assigns it to the Scale field.
func (o *ChartResponseChartMeta) SetScale(v int32) {
	o.Scale = &v
}

// GetPriceHint returns the PriceHint field value if set, zero value otherwise.
func (o *ChartResponseChartMeta) GetPriceHint() int32 {
	if o == nil || o.PriceHint == nil {
		var ret int32
		return ret
	}
	return *o.PriceHint
}

// GetPriceHintOk returns a tuple with the PriceHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartMeta) GetPriceHintOk() (*int32, bool) {
	if o == nil || o.PriceHint == nil {
		return nil, false
	}
	return o.PriceHint, true
}

// HasPriceHint returns a boolean if a field has been set.
func (o *ChartResponseChartMeta) HasPriceHint() bool {
	if o != nil && o.PriceHint != nil {
		return true
	}

	return false
}

// SetPriceHint gets a reference to the given int32 and assigns it to the PriceHint field.
func (o *ChartResponseChartMeta) SetPriceHint(v int32) {
	o.PriceHint = &v
}

// GetCurrentTradingPeriod returns the CurrentTradingPeriod field value if set, zero value otherwise.
func (o *ChartResponseChartMeta) GetCurrentTradingPeriod() ChartResponseChartMetaCurrentTradingPeriod {
	if o == nil || o.CurrentTradingPeriod == nil {
		var ret ChartResponseChartMetaCurrentTradingPeriod
		return ret
	}
	return *o.CurrentTradingPeriod
}

// GetCurrentTradingPeriodOk returns a tuple with the CurrentTradingPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartMeta) GetCurrentTradingPeriodOk() (*ChartResponseChartMetaCurrentTradingPeriod, bool) {
	if o == nil || o.CurrentTradingPeriod == nil {
		return nil, false
	}
	return o.CurrentTradingPeriod, true
}

// HasCurrentTradingPeriod returns a boolean if a field has been set.
func (o *ChartResponseChartMeta) HasCurrentTradingPeriod() bool {
	if o != nil && o.CurrentTradingPeriod != nil {
		return true
	}

	return false
}

// SetCurrentTradingPeriod gets a reference to the given ChartResponseChartMetaCurrentTradingPeriod and assigns it to the CurrentTradingPeriod field.
func (o *ChartResponseChartMeta) SetCurrentTradingPeriod(v ChartResponseChartMetaCurrentTradingPeriod) {
	o.CurrentTradingPeriod = &v
}

// GetTradingPeriods returns the TradingPeriods field value if set, zero value otherwise.
func (o *ChartResponseChartMeta) GetTradingPeriods() [][]map[string]interface{} {
	if o == nil || o.TradingPeriods == nil {
		var ret [][]map[string]interface{}
		return ret
	}
	return *o.TradingPeriods
}

// GetTradingPeriodsOk returns a tuple with the TradingPeriods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartMeta) GetTradingPeriodsOk() (*[][]map[string]interface{}, bool) {
	if o == nil || o.TradingPeriods == nil {
		return nil, false
	}
	return o.TradingPeriods, true
}

// HasTradingPeriods returns a boolean if a field has been set.
func (o *ChartResponseChartMeta) HasTradingPeriods() bool {
	if o != nil && o.TradingPeriods != nil {
		return true
	}

	return false
}

// SetTradingPeriods gets a reference to the given [][]map[string]interface{} and assigns it to the TradingPeriods field.
func (o *ChartResponseChartMeta) SetTradingPeriods(v [][]map[string]interface{}) {
	o.TradingPeriods = &v
}

// GetDataGranularity returns the DataGranularity field value if set, zero value otherwise.
func (o *ChartResponseChartMeta) GetDataGranularity() string {
	if o == nil || o.DataGranularity == nil {
		var ret string
		return ret
	}
	return *o.DataGranularity
}

// GetDataGranularityOk returns a tuple with the DataGranularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartMeta) GetDataGranularityOk() (*string, bool) {
	if o == nil || o.DataGranularity == nil {
		return nil, false
	}
	return o.DataGranularity, true
}

// HasDataGranularity returns a boolean if a field has been set.
func (o *ChartResponseChartMeta) HasDataGranularity() bool {
	if o != nil && o.DataGranularity != nil {
		return true
	}

	return false
}

// SetDataGranularity gets a reference to the given string and assigns it to the DataGranularity field.
func (o *ChartResponseChartMeta) SetDataGranularity(v string) {
	o.DataGranularity = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *ChartResponseChartMeta) GetRange() string {
	if o == nil || o.Range == nil {
		var ret string
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartMeta) GetRangeOk() (*string, bool) {
	if o == nil || o.Range == nil {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *ChartResponseChartMeta) HasRange() bool {
	if o != nil && o.Range != nil {
		return true
	}

	return false
}

// SetRange gets a reference to the given string and assigns it to the Range field.
func (o *ChartResponseChartMeta) SetRange(v string) {
	o.Range = &v
}

// GetValidRanges returns the ValidRanges field value if set, zero value otherwise.
func (o *ChartResponseChartMeta) GetValidRanges() []string {
	if o == nil || o.ValidRanges == nil {
		var ret []string
		return ret
	}
	return *o.ValidRanges
}

// GetValidRangesOk returns a tuple with the ValidRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartMeta) GetValidRangesOk() (*[]string, bool) {
	if o == nil || o.ValidRanges == nil {
		return nil, false
	}
	return o.ValidRanges, true
}

// HasValidRanges returns a boolean if a field has been set.
func (o *ChartResponseChartMeta) HasValidRanges() bool {
	if o != nil && o.ValidRanges != nil {
		return true
	}

	return false
}

// SetValidRanges gets a reference to the given []string and assigns it to the ValidRanges field.
func (o *ChartResponseChartMeta) SetValidRanges(v []string) {
	o.ValidRanges = &v
}

func (o ChartResponseChartMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.ExchangeName != nil {
		toSerialize["exchangeName"] = o.ExchangeName
	}
	if o.InstrumentType != nil {
		toSerialize["instrumentType"] = o.InstrumentType
	}
	if o.FirstTradeDate != nil {
		toSerialize["firstTradeDate"] = o.FirstTradeDate
	}
	if o.RegularMarketTime != nil {
		toSerialize["regularMarketTime"] = o.RegularMarketTime
	}
	if o.Gmtoffset != nil {
		toSerialize["gmtoffset"] = o.Gmtoffset
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	if o.ExchangeTimezoneName != nil {
		toSerialize["exchangeTimezoneName"] = o.ExchangeTimezoneName
	}
	if o.RegularMarketPrice != nil {
		toSerialize["regularMarketPrice"] = o.RegularMarketPrice
	}
	if o.ChartPreviousClose != nil {
		toSerialize["chartPreviousClose"] = o.ChartPreviousClose
	}
	if o.PreviousClose != nil {
		toSerialize["previousClose"] = o.PreviousClose
	}
	if o.Scale != nil {
		toSerialize["scale"] = o.Scale
	}
	if o.PriceHint != nil {
		toSerialize["priceHint"] = o.PriceHint
	}
	if o.CurrentTradingPeriod != nil {
		toSerialize["currentTradingPeriod"] = o.CurrentTradingPeriod
	}
	if o.TradingPeriods != nil {
		toSerialize["tradingPeriods"] = o.TradingPeriods
	}
	if o.DataGranularity != nil {
		toSerialize["dataGranularity"] = o.DataGranularity
	}
	if o.Range != nil {
		toSerialize["range"] = o.Range
	}
	if o.ValidRanges != nil {
		toSerialize["validRanges"] = o.ValidRanges
	}
	return json.Marshal(toSerialize)
}

type NullableChartResponseChartMeta struct {
	value *ChartResponseChartMeta
	isSet bool
}

func (v NullableChartResponseChartMeta) Get() *ChartResponseChartMeta {
	return v.value
}

func (v *NullableChartResponseChartMeta) Set(val *ChartResponseChartMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableChartResponseChartMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableChartResponseChartMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChartResponseChartMeta(val *ChartResponseChartMeta) *NullableChartResponseChartMeta {
	return &NullableChartResponseChartMeta{value: val, isSet: true}
}

func (v NullableChartResponseChartMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChartResponseChartMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


