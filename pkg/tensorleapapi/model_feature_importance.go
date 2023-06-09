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

// checks if the FeatureImportance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureImportance{}

// FeatureImportance struct for FeatureImportance
type FeatureImportance struct {
	GradCam []FeatureImportanceItem `json:"grad_cam"`
	GradByLoss []FeatureImportanceItem `json:"grad_by_loss"`
}

// NewFeatureImportance instantiates a new FeatureImportance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureImportance(gradCam []FeatureImportanceItem, gradByLoss []FeatureImportanceItem) *FeatureImportance {
	this := FeatureImportance{}
	this.GradCam = gradCam
	this.GradByLoss = gradByLoss
	return &this
}

// NewFeatureImportanceWithDefaults instantiates a new FeatureImportance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureImportanceWithDefaults() *FeatureImportance {
	this := FeatureImportance{}
	return &this
}

// GetGradCam returns the GradCam field value
func (o *FeatureImportance) GetGradCam() []FeatureImportanceItem {
	if o == nil {
		var ret []FeatureImportanceItem
		return ret
	}

	return o.GradCam
}

// GetGradCamOk returns a tuple with the GradCam field value
// and a boolean to check if the value has been set.
func (o *FeatureImportance) GetGradCamOk() ([]FeatureImportanceItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.GradCam, true
}

// SetGradCam sets field value
func (o *FeatureImportance) SetGradCam(v []FeatureImportanceItem) {
	o.GradCam = v
}

// GetGradByLoss returns the GradByLoss field value
func (o *FeatureImportance) GetGradByLoss() []FeatureImportanceItem {
	if o == nil {
		var ret []FeatureImportanceItem
		return ret
	}

	return o.GradByLoss
}

// GetGradByLossOk returns a tuple with the GradByLoss field value
// and a boolean to check if the value has been set.
func (o *FeatureImportance) GetGradByLossOk() ([]FeatureImportanceItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.GradByLoss, true
}

// SetGradByLoss sets field value
func (o *FeatureImportance) SetGradByLoss(v []FeatureImportanceItem) {
	o.GradByLoss = v
}

func (o FeatureImportance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureImportance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["grad_cam"] = o.GradCam
	toSerialize["grad_by_loss"] = o.GradByLoss
	return toSerialize, nil
}

type NullableFeatureImportance struct {
	value *FeatureImportance
	isSet bool
}

func (v NullableFeatureImportance) Get() *FeatureImportance {
	return v.value
}

func (v *NullableFeatureImportance) Set(val *FeatureImportance) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureImportance) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureImportance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureImportance(val *FeatureImportance) *NullableFeatureImportance {
	return &NullableFeatureImportance{value: val, isSet: true}
}

func (v NullableFeatureImportance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureImportance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


