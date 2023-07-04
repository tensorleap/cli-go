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

// checks if the ScatterVizDataState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScatterVizDataState{}

// ScatterVizDataState struct for ScatterVizDataState
type ScatterVizDataState struct {
	DataState *DataStateType `json:"data_state,omitempty"`
	ScatterData [][]float64 `json:"scatter_data"`
	Labels []ScatterLabel `json:"labels"`
	Samples []SampleIdentity `json:"samples"`
	// Construct a type with a set of properties K of type T
	Metadata map[string]interface{} `json:"metadata"`
	// Construct a type with a set of properties K of type T
	MiByCluster map[string]interface{} `json:"mi_by_cluster,omitempty"`
}

// NewScatterVizDataState instantiates a new ScatterVizDataState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScatterVizDataState(scatterData [][]float64, labels []ScatterLabel, samples []SampleIdentity, metadata map[string]interface{}) *ScatterVizDataState {
	this := ScatterVizDataState{}
	this.ScatterData = scatterData
	this.Labels = labels
	this.Samples = samples
	this.Metadata = metadata
	return &this
}

// NewScatterVizDataStateWithDefaults instantiates a new ScatterVizDataState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScatterVizDataStateWithDefaults() *ScatterVizDataState {
	this := ScatterVizDataState{}
	return &this
}

// GetDataState returns the DataState field value if set, zero value otherwise.
func (o *ScatterVizDataState) GetDataState() DataStateType {
	if o == nil || IsNil(o.DataState) {
		var ret DataStateType
		return ret
	}
	return *o.DataState
}

// GetDataStateOk returns a tuple with the DataState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterVizDataState) GetDataStateOk() (*DataStateType, bool) {
	if o == nil || IsNil(o.DataState) {
		return nil, false
	}
	return o.DataState, true
}

// HasDataState returns a boolean if a field has been set.
func (o *ScatterVizDataState) HasDataState() bool {
	if o != nil && !IsNil(o.DataState) {
		return true
	}

	return false
}

// SetDataState gets a reference to the given DataStateType and assigns it to the DataState field.
func (o *ScatterVizDataState) SetDataState(v DataStateType) {
	o.DataState = &v
}

// GetScatterData returns the ScatterData field value
func (o *ScatterVizDataState) GetScatterData() [][]float64 {
	if o == nil {
		var ret [][]float64
		return ret
	}

	return o.ScatterData
}

// GetScatterDataOk returns a tuple with the ScatterData field value
// and a boolean to check if the value has been set.
func (o *ScatterVizDataState) GetScatterDataOk() ([][]float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScatterData, true
}

// SetScatterData sets field value
func (o *ScatterVizDataState) SetScatterData(v [][]float64) {
	o.ScatterData = v
}

// GetLabels returns the Labels field value
func (o *ScatterVizDataState) GetLabels() []ScatterLabel {
	if o == nil {
		var ret []ScatterLabel
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *ScatterVizDataState) GetLabelsOk() ([]ScatterLabel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *ScatterVizDataState) SetLabels(v []ScatterLabel) {
	o.Labels = v
}

// GetSamples returns the Samples field value
func (o *ScatterVizDataState) GetSamples() []SampleIdentity {
	if o == nil {
		var ret []SampleIdentity
		return ret
	}

	return o.Samples
}

// GetSamplesOk returns a tuple with the Samples field value
// and a boolean to check if the value has been set.
func (o *ScatterVizDataState) GetSamplesOk() ([]SampleIdentity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Samples, true
}

// SetSamples sets field value
func (o *ScatterVizDataState) SetSamples(v []SampleIdentity) {
	o.Samples = v
}

// GetMetadata returns the Metadata field value
func (o *ScatterVizDataState) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ScatterVizDataState) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *ScatterVizDataState) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetMiByCluster returns the MiByCluster field value if set, zero value otherwise.
func (o *ScatterVizDataState) GetMiByCluster() map[string]interface{} {
	if o == nil || IsNil(o.MiByCluster) {
		var ret map[string]interface{}
		return ret
	}
	return o.MiByCluster
}

// GetMiByClusterOk returns a tuple with the MiByCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterVizDataState) GetMiByClusterOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.MiByCluster) {
		return map[string]interface{}{}, false
	}
	return o.MiByCluster, true
}

// HasMiByCluster returns a boolean if a field has been set.
func (o *ScatterVizDataState) HasMiByCluster() bool {
	if o != nil && !IsNil(o.MiByCluster) {
		return true
	}

	return false
}

// SetMiByCluster gets a reference to the given map[string]interface{} and assigns it to the MiByCluster field.
func (o *ScatterVizDataState) SetMiByCluster(v map[string]interface{}) {
	o.MiByCluster = v
}

func (o ScatterVizDataState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScatterVizDataState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataState) {
		toSerialize["data_state"] = o.DataState
	}
	toSerialize["scatter_data"] = o.ScatterData
	toSerialize["labels"] = o.Labels
	toSerialize["samples"] = o.Samples
	toSerialize["metadata"] = o.Metadata
	if !IsNil(o.MiByCluster) {
		toSerialize["mi_by_cluster"] = o.MiByCluster
	}
	return toSerialize, nil
}

type NullableScatterVizDataState struct {
	value *ScatterVizDataState
	isSet bool
}

func (v NullableScatterVizDataState) Get() *ScatterVizDataState {
	return v.value
}

func (v *NullableScatterVizDataState) Set(val *ScatterVizDataState) {
	v.value = val
	v.isSet = true
}

func (v NullableScatterVizDataState) IsSet() bool {
	return v.isSet
}

func (v *NullableScatterVizDataState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScatterVizDataState(val *ScatterVizDataState) *NullableScatterVizDataState {
	return &NullableScatterVizDataState{value: val, isSet: true}
}

func (v NullableScatterVizDataState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScatterVizDataState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


