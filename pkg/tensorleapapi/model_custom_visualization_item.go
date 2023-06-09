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

// checks if the CustomVisualizationItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomVisualizationItem{}

// CustomVisualizationItem struct for CustomVisualizationItem
type CustomVisualizationItem struct {
	Cid string `json:"cid"`
	Layout SizedLayout `json:"layout"`
	Type string `json:"type"`
	Data CustomVisualizationData `json:"data"`
}

// NewCustomVisualizationItem instantiates a new CustomVisualizationItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomVisualizationItem(cid string, layout SizedLayout, type_ string, data CustomVisualizationData) *CustomVisualizationItem {
	this := CustomVisualizationItem{}
	this.Cid = cid
	this.Layout = layout
	this.Type = type_
	this.Data = data
	return &this
}

// NewCustomVisualizationItemWithDefaults instantiates a new CustomVisualizationItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomVisualizationItemWithDefaults() *CustomVisualizationItem {
	this := CustomVisualizationItem{}
	return &this
}

// GetCid returns the Cid field value
func (o *CustomVisualizationItem) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *CustomVisualizationItem) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *CustomVisualizationItem) SetCid(v string) {
	o.Cid = v
}

// GetLayout returns the Layout field value
func (o *CustomVisualizationItem) GetLayout() SizedLayout {
	if o == nil {
		var ret SizedLayout
		return ret
	}

	return o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value
// and a boolean to check if the value has been set.
func (o *CustomVisualizationItem) GetLayoutOk() (*SizedLayout, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Layout, true
}

// SetLayout sets field value
func (o *CustomVisualizationItem) SetLayout(v SizedLayout) {
	o.Layout = v
}

// GetType returns the Type field value
func (o *CustomVisualizationItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomVisualizationItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomVisualizationItem) SetType(v string) {
	o.Type = v
}

// GetData returns the Data field value
func (o *CustomVisualizationItem) GetData() CustomVisualizationData {
	if o == nil {
		var ret CustomVisualizationData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CustomVisualizationItem) GetDataOk() (*CustomVisualizationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CustomVisualizationItem) SetData(v CustomVisualizationData) {
	o.Data = v
}

func (o CustomVisualizationItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomVisualizationItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["layout"] = o.Layout
	toSerialize["type"] = o.Type
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableCustomVisualizationItem struct {
	value *CustomVisualizationItem
	isSet bool
}

func (v NullableCustomVisualizationItem) Get() *CustomVisualizationItem {
	return v.value
}

func (v *NullableCustomVisualizationItem) Set(val *CustomVisualizationItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomVisualizationItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomVisualizationItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomVisualizationItem(val *CustomVisualizationItem) *NullableCustomVisualizationItem {
	return &NullableCustomVisualizationItem{value: val, isSet: true}
}

func (v NullableCustomVisualizationItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomVisualizationItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


