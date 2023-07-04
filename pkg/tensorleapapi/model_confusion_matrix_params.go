/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.342
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the ConfusionMatrixParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfusionMatrixParams{}

// ConfusionMatrixParams struct for ConfusionMatrixParams
type ConfusionMatrixParams struct {
	ProjectId string `json:"projectId"`
	XField string `json:"xField"`
	SessionRunIds []string `json:"sessionRunIds"`
	VerticalSplit *string `json:"verticalSplit,omitempty"`
	HorizontalSplit *string `json:"horizontalSplit,omitempty"`
	Threshold *float64 `json:"threshold,omitempty"`
	CustomMetricName string `json:"customMetricName"`
	Filters []ESFilter `json:"filters,omitempty"`
	XFieldAggregationType *EsBatchAggregationType `json:"xFieldAggregationType,omitempty"`
	DataDistributionType *DataDistributionType `json:"dataDistributionType,omitempty"`
	OrderByParam *string `json:"orderByParam,omitempty"`
	OrderParams *OrderType `json:"orderParams,omitempty"`
	XAxisSizeInterval float64 `json:"xAxisSizeInterval"`
	LastEpochOnly *bool `json:"lastEpochOnly,omitempty"`
}

// NewConfusionMatrixParams instantiates a new ConfusionMatrixParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfusionMatrixParams(projectId string, xField string, sessionRunIds []string, customMetricName string, xAxisSizeInterval float64) *ConfusionMatrixParams {
	this := ConfusionMatrixParams{}
	this.ProjectId = projectId
	this.XField = xField
	this.SessionRunIds = sessionRunIds
	this.CustomMetricName = customMetricName
	this.XAxisSizeInterval = xAxisSizeInterval
	return &this
}

// NewConfusionMatrixParamsWithDefaults instantiates a new ConfusionMatrixParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfusionMatrixParamsWithDefaults() *ConfusionMatrixParams {
	this := ConfusionMatrixParams{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *ConfusionMatrixParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *ConfusionMatrixParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *ConfusionMatrixParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetXField returns the XField field value
func (o *ConfusionMatrixParams) GetXField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.XField
}

// GetXFieldOk returns a tuple with the XField field value
// and a boolean to check if the value has been set.
func (o *ConfusionMatrixParams) GetXFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.XField, true
}

// SetXField sets field value
func (o *ConfusionMatrixParams) SetXField(v string) {
	o.XField = v
}

// GetSessionRunIds returns the SessionRunIds field value
func (o *ConfusionMatrixParams) GetSessionRunIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SessionRunIds
}

// GetSessionRunIdsOk returns a tuple with the SessionRunIds field value
// and a boolean to check if the value has been set.
func (o *ConfusionMatrixParams) GetSessionRunIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionRunIds, true
}

// SetSessionRunIds sets field value
func (o *ConfusionMatrixParams) SetSessionRunIds(v []string) {
	o.SessionRunIds = v
}

// GetVerticalSplit returns the VerticalSplit field value if set, zero value otherwise.
func (o *ConfusionMatrixParams) GetVerticalSplit() string {
	if o == nil || IsNil(o.VerticalSplit) {
		var ret string
		return ret
	}
	return *o.VerticalSplit
}

// GetVerticalSplitOk returns a tuple with the VerticalSplit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfusionMatrixParams) GetVerticalSplitOk() (*string, bool) {
	if o == nil || IsNil(o.VerticalSplit) {
		return nil, false
	}
	return o.VerticalSplit, true
}

// HasVerticalSplit returns a boolean if a field has been set.
func (o *ConfusionMatrixParams) HasVerticalSplit() bool {
	if o != nil && !IsNil(o.VerticalSplit) {
		return true
	}

	return false
}

// SetVerticalSplit gets a reference to the given string and assigns it to the VerticalSplit field.
func (o *ConfusionMatrixParams) SetVerticalSplit(v string) {
	o.VerticalSplit = &v
}

// GetHorizontalSplit returns the HorizontalSplit field value if set, zero value otherwise.
func (o *ConfusionMatrixParams) GetHorizontalSplit() string {
	if o == nil || IsNil(o.HorizontalSplit) {
		var ret string
		return ret
	}
	return *o.HorizontalSplit
}

// GetHorizontalSplitOk returns a tuple with the HorizontalSplit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfusionMatrixParams) GetHorizontalSplitOk() (*string, bool) {
	if o == nil || IsNil(o.HorizontalSplit) {
		return nil, false
	}
	return o.HorizontalSplit, true
}

// HasHorizontalSplit returns a boolean if a field has been set.
func (o *ConfusionMatrixParams) HasHorizontalSplit() bool {
	if o != nil && !IsNil(o.HorizontalSplit) {
		return true
	}

	return false
}

// SetHorizontalSplit gets a reference to the given string and assigns it to the HorizontalSplit field.
func (o *ConfusionMatrixParams) SetHorizontalSplit(v string) {
	o.HorizontalSplit = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *ConfusionMatrixParams) GetThreshold() float64 {
	if o == nil || IsNil(o.Threshold) {
		var ret float64
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfusionMatrixParams) GetThresholdOk() (*float64, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *ConfusionMatrixParams) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given float64 and assigns it to the Threshold field.
func (o *ConfusionMatrixParams) SetThreshold(v float64) {
	o.Threshold = &v
}

// GetCustomMetricName returns the CustomMetricName field value
func (o *ConfusionMatrixParams) GetCustomMetricName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomMetricName
}

// GetCustomMetricNameOk returns a tuple with the CustomMetricName field value
// and a boolean to check if the value has been set.
func (o *ConfusionMatrixParams) GetCustomMetricNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomMetricName, true
}

// SetCustomMetricName sets field value
func (o *ConfusionMatrixParams) SetCustomMetricName(v string) {
	o.CustomMetricName = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ConfusionMatrixParams) GetFilters() []ESFilter {
	if o == nil || IsNil(o.Filters) {
		var ret []ESFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfusionMatrixParams) GetFiltersOk() ([]ESFilter, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ConfusionMatrixParams) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []ESFilter and assigns it to the Filters field.
func (o *ConfusionMatrixParams) SetFilters(v []ESFilter) {
	o.Filters = v
}

// GetXFieldAggregationType returns the XFieldAggregationType field value if set, zero value otherwise.
func (o *ConfusionMatrixParams) GetXFieldAggregationType() EsBatchAggregationType {
	if o == nil || IsNil(o.XFieldAggregationType) {
		var ret EsBatchAggregationType
		return ret
	}
	return *o.XFieldAggregationType
}

// GetXFieldAggregationTypeOk returns a tuple with the XFieldAggregationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfusionMatrixParams) GetXFieldAggregationTypeOk() (*EsBatchAggregationType, bool) {
	if o == nil || IsNil(o.XFieldAggregationType) {
		return nil, false
	}
	return o.XFieldAggregationType, true
}

// HasXFieldAggregationType returns a boolean if a field has been set.
func (o *ConfusionMatrixParams) HasXFieldAggregationType() bool {
	if o != nil && !IsNil(o.XFieldAggregationType) {
		return true
	}

	return false
}

// SetXFieldAggregationType gets a reference to the given EsBatchAggregationType and assigns it to the XFieldAggregationType field.
func (o *ConfusionMatrixParams) SetXFieldAggregationType(v EsBatchAggregationType) {
	o.XFieldAggregationType = &v
}

// GetDataDistributionType returns the DataDistributionType field value if set, zero value otherwise.
func (o *ConfusionMatrixParams) GetDataDistributionType() DataDistributionType {
	if o == nil || IsNil(o.DataDistributionType) {
		var ret DataDistributionType
		return ret
	}
	return *o.DataDistributionType
}

// GetDataDistributionTypeOk returns a tuple with the DataDistributionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfusionMatrixParams) GetDataDistributionTypeOk() (*DataDistributionType, bool) {
	if o == nil || IsNil(o.DataDistributionType) {
		return nil, false
	}
	return o.DataDistributionType, true
}

// HasDataDistributionType returns a boolean if a field has been set.
func (o *ConfusionMatrixParams) HasDataDistributionType() bool {
	if o != nil && !IsNil(o.DataDistributionType) {
		return true
	}

	return false
}

// SetDataDistributionType gets a reference to the given DataDistributionType and assigns it to the DataDistributionType field.
func (o *ConfusionMatrixParams) SetDataDistributionType(v DataDistributionType) {
	o.DataDistributionType = &v
}

// GetOrderByParam returns the OrderByParam field value if set, zero value otherwise.
func (o *ConfusionMatrixParams) GetOrderByParam() string {
	if o == nil || IsNil(o.OrderByParam) {
		var ret string
		return ret
	}
	return *o.OrderByParam
}

// GetOrderByParamOk returns a tuple with the OrderByParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfusionMatrixParams) GetOrderByParamOk() (*string, bool) {
	if o == nil || IsNil(o.OrderByParam) {
		return nil, false
	}
	return o.OrderByParam, true
}

// HasOrderByParam returns a boolean if a field has been set.
func (o *ConfusionMatrixParams) HasOrderByParam() bool {
	if o != nil && !IsNil(o.OrderByParam) {
		return true
	}

	return false
}

// SetOrderByParam gets a reference to the given string and assigns it to the OrderByParam field.
func (o *ConfusionMatrixParams) SetOrderByParam(v string) {
	o.OrderByParam = &v
}

// GetOrderParams returns the OrderParams field value if set, zero value otherwise.
func (o *ConfusionMatrixParams) GetOrderParams() OrderType {
	if o == nil || IsNil(o.OrderParams) {
		var ret OrderType
		return ret
	}
	return *o.OrderParams
}

// GetOrderParamsOk returns a tuple with the OrderParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfusionMatrixParams) GetOrderParamsOk() (*OrderType, bool) {
	if o == nil || IsNil(o.OrderParams) {
		return nil, false
	}
	return o.OrderParams, true
}

// HasOrderParams returns a boolean if a field has been set.
func (o *ConfusionMatrixParams) HasOrderParams() bool {
	if o != nil && !IsNil(o.OrderParams) {
		return true
	}

	return false
}

// SetOrderParams gets a reference to the given OrderType and assigns it to the OrderParams field.
func (o *ConfusionMatrixParams) SetOrderParams(v OrderType) {
	o.OrderParams = &v
}

// GetXAxisSizeInterval returns the XAxisSizeInterval field value
func (o *ConfusionMatrixParams) GetXAxisSizeInterval() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.XAxisSizeInterval
}

// GetXAxisSizeIntervalOk returns a tuple with the XAxisSizeInterval field value
// and a boolean to check if the value has been set.
func (o *ConfusionMatrixParams) GetXAxisSizeIntervalOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.XAxisSizeInterval, true
}

// SetXAxisSizeInterval sets field value
func (o *ConfusionMatrixParams) SetXAxisSizeInterval(v float64) {
	o.XAxisSizeInterval = v
}

// GetLastEpochOnly returns the LastEpochOnly field value if set, zero value otherwise.
func (o *ConfusionMatrixParams) GetLastEpochOnly() bool {
	if o == nil || IsNil(o.LastEpochOnly) {
		var ret bool
		return ret
	}
	return *o.LastEpochOnly
}

// GetLastEpochOnlyOk returns a tuple with the LastEpochOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfusionMatrixParams) GetLastEpochOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.LastEpochOnly) {
		return nil, false
	}
	return o.LastEpochOnly, true
}

// HasLastEpochOnly returns a boolean if a field has been set.
func (o *ConfusionMatrixParams) HasLastEpochOnly() bool {
	if o != nil && !IsNil(o.LastEpochOnly) {
		return true
	}

	return false
}

// SetLastEpochOnly gets a reference to the given bool and assigns it to the LastEpochOnly field.
func (o *ConfusionMatrixParams) SetLastEpochOnly(v bool) {
	o.LastEpochOnly = &v
}

func (o ConfusionMatrixParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfusionMatrixParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["xField"] = o.XField
	toSerialize["sessionRunIds"] = o.SessionRunIds
	if !IsNil(o.VerticalSplit) {
		toSerialize["verticalSplit"] = o.VerticalSplit
	}
	if !IsNil(o.HorizontalSplit) {
		toSerialize["horizontalSplit"] = o.HorizontalSplit
	}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	toSerialize["customMetricName"] = o.CustomMetricName
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.XFieldAggregationType) {
		toSerialize["xFieldAggregationType"] = o.XFieldAggregationType
	}
	if !IsNil(o.DataDistributionType) {
		toSerialize["dataDistributionType"] = o.DataDistributionType
	}
	if !IsNil(o.OrderByParam) {
		toSerialize["orderByParam"] = o.OrderByParam
	}
	if !IsNil(o.OrderParams) {
		toSerialize["orderParams"] = o.OrderParams
	}
	toSerialize["xAxisSizeInterval"] = o.XAxisSizeInterval
	if !IsNil(o.LastEpochOnly) {
		toSerialize["lastEpochOnly"] = o.LastEpochOnly
	}
	return toSerialize, nil
}

type NullableConfusionMatrixParams struct {
	value *ConfusionMatrixParams
	isSet bool
}

func (v NullableConfusionMatrixParams) Get() *ConfusionMatrixParams {
	return v.value
}

func (v *NullableConfusionMatrixParams) Set(val *ConfusionMatrixParams) {
	v.value = val
	v.isSet = true
}

func (v NullableConfusionMatrixParams) IsSet() bool {
	return v.isSet
}

func (v *NullableConfusionMatrixParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfusionMatrixParams(val *ConfusionMatrixParams) *NullableConfusionMatrixParams {
	return &NullableConfusionMatrixParams{value: val, isSet: true}
}

func (v NullableConfusionMatrixParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfusionMatrixParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


