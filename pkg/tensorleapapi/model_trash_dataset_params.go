/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.301
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the TrashDatasetParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrashDatasetParams{}

// TrashDatasetParams struct for TrashDatasetParams
type TrashDatasetParams struct {
	DatasetId string `json:"datasetId"`
}

// NewTrashDatasetParams instantiates a new TrashDatasetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrashDatasetParams(datasetId string) *TrashDatasetParams {
	this := TrashDatasetParams{}
	this.DatasetId = datasetId
	return &this
}

// NewTrashDatasetParamsWithDefaults instantiates a new TrashDatasetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrashDatasetParamsWithDefaults() *TrashDatasetParams {
	this := TrashDatasetParams{}
	return &this
}

// GetDatasetId returns the DatasetId field value
func (o *TrashDatasetParams) GetDatasetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value
// and a boolean to check if the value has been set.
func (o *TrashDatasetParams) GetDatasetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetId, true
}

// SetDatasetId sets field value
func (o *TrashDatasetParams) SetDatasetId(v string) {
	o.DatasetId = v
}

func (o TrashDatasetParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrashDatasetParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datasetId"] = o.DatasetId
	return toSerialize, nil
}

type NullableTrashDatasetParams struct {
	value *TrashDatasetParams
	isSet bool
}

func (v NullableTrashDatasetParams) Get() *TrashDatasetParams {
	return v.value
}

func (v *NullableTrashDatasetParams) Set(val *TrashDatasetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableTrashDatasetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableTrashDatasetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrashDatasetParams(val *TrashDatasetParams) *NullableTrashDatasetParams {
	return &NullableTrashDatasetParams{value: val, isSet: true}
}

func (v NullableTrashDatasetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrashDatasetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


