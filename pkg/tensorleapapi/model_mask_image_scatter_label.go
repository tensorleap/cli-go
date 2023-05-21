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

// checks if the MaskImageScatterLabel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaskImageScatterLabel{}

// MaskImageScatterLabel struct for MaskImageScatterLabel
type MaskImageScatterLabel struct {
	Data [][]float64 `json:"data"`
	Src string `json:"src"`
	Blob string `json:"blob"`
	Type string `json:"type"`
	MaskSrc string `json:"mask_src"`
	MaskBlob string `json:"mask_blob"`
	Labels []string `json:"labels"`
	InputName *string `json:"input_name,omitempty"`
}

// NewMaskImageScatterLabel instantiates a new MaskImageScatterLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaskImageScatterLabel(data [][]float64, src string, blob string, type_ string, maskSrc string, maskBlob string, labels []string) *MaskImageScatterLabel {
	this := MaskImageScatterLabel{}
	this.Data = data
	this.Src = src
	this.Blob = blob
	this.Type = type_
	this.MaskSrc = maskSrc
	this.MaskBlob = maskBlob
	this.Labels = labels
	return &this
}

// NewMaskImageScatterLabelWithDefaults instantiates a new MaskImageScatterLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaskImageScatterLabelWithDefaults() *MaskImageScatterLabel {
	this := MaskImageScatterLabel{}
	return &this
}

// GetData returns the Data field value
func (o *MaskImageScatterLabel) GetData() [][]float64 {
	if o == nil {
		var ret [][]float64
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MaskImageScatterLabel) GetDataOk() ([][]float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *MaskImageScatterLabel) SetData(v [][]float64) {
	o.Data = v
}

// GetSrc returns the Src field value
func (o *MaskImageScatterLabel) GetSrc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Src
}

// GetSrcOk returns a tuple with the Src field value
// and a boolean to check if the value has been set.
func (o *MaskImageScatterLabel) GetSrcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Src, true
}

// SetSrc sets field value
func (o *MaskImageScatterLabel) SetSrc(v string) {
	o.Src = v
}

// GetBlob returns the Blob field value
func (o *MaskImageScatterLabel) GetBlob() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Blob
}

// GetBlobOk returns a tuple with the Blob field value
// and a boolean to check if the value has been set.
func (o *MaskImageScatterLabel) GetBlobOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blob, true
}

// SetBlob sets field value
func (o *MaskImageScatterLabel) SetBlob(v string) {
	o.Blob = v
}

// GetType returns the Type field value
func (o *MaskImageScatterLabel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MaskImageScatterLabel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MaskImageScatterLabel) SetType(v string) {
	o.Type = v
}

// GetMaskSrc returns the MaskSrc field value
func (o *MaskImageScatterLabel) GetMaskSrc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaskSrc
}

// GetMaskSrcOk returns a tuple with the MaskSrc field value
// and a boolean to check if the value has been set.
func (o *MaskImageScatterLabel) GetMaskSrcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaskSrc, true
}

// SetMaskSrc sets field value
func (o *MaskImageScatterLabel) SetMaskSrc(v string) {
	o.MaskSrc = v
}

// GetMaskBlob returns the MaskBlob field value
func (o *MaskImageScatterLabel) GetMaskBlob() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaskBlob
}

// GetMaskBlobOk returns a tuple with the MaskBlob field value
// and a boolean to check if the value has been set.
func (o *MaskImageScatterLabel) GetMaskBlobOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaskBlob, true
}

// SetMaskBlob sets field value
func (o *MaskImageScatterLabel) SetMaskBlob(v string) {
	o.MaskBlob = v
}

// GetLabels returns the Labels field value
func (o *MaskImageScatterLabel) GetLabels() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *MaskImageScatterLabel) GetLabelsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *MaskImageScatterLabel) SetLabels(v []string) {
	o.Labels = v
}

// GetInputName returns the InputName field value if set, zero value otherwise.
func (o *MaskImageScatterLabel) GetInputName() string {
	if o == nil || IsNil(o.InputName) {
		var ret string
		return ret
	}
	return *o.InputName
}

// GetInputNameOk returns a tuple with the InputName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskImageScatterLabel) GetInputNameOk() (*string, bool) {
	if o == nil || IsNil(o.InputName) {
		return nil, false
	}
	return o.InputName, true
}

// HasInputName returns a boolean if a field has been set.
func (o *MaskImageScatterLabel) HasInputName() bool {
	if o != nil && !IsNil(o.InputName) {
		return true
	}

	return false
}

// SetInputName gets a reference to the given string and assigns it to the InputName field.
func (o *MaskImageScatterLabel) SetInputName(v string) {
	o.InputName = &v
}

func (o MaskImageScatterLabel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaskImageScatterLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["src"] = o.Src
	toSerialize["blob"] = o.Blob
	toSerialize["type"] = o.Type
	toSerialize["mask_src"] = o.MaskSrc
	toSerialize["mask_blob"] = o.MaskBlob
	toSerialize["labels"] = o.Labels
	if !IsNil(o.InputName) {
		toSerialize["input_name"] = o.InputName
	}
	return toSerialize, nil
}

type NullableMaskImageScatterLabel struct {
	value *MaskImageScatterLabel
	isSet bool
}

func (v NullableMaskImageScatterLabel) Get() *MaskImageScatterLabel {
	return v.value
}

func (v *NullableMaskImageScatterLabel) Set(val *MaskImageScatterLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableMaskImageScatterLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableMaskImageScatterLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaskImageScatterLabel(val *MaskImageScatterLabel) *NullableMaskImageScatterLabel {
	return &NullableMaskImageScatterLabel{value: val, isSet: true}
}

func (v NullableMaskImageScatterLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaskImageScatterLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


