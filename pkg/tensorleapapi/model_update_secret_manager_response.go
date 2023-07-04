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

// checks if the UpdateSecretManagerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSecretManagerResponse{}

// UpdateSecretManagerResponse struct for UpdateSecretManagerResponse
type UpdateSecretManagerResponse struct {
	Success bool `json:"success"`
}

// NewUpdateSecretManagerResponse instantiates a new UpdateSecretManagerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSecretManagerResponse(success bool) *UpdateSecretManagerResponse {
	this := UpdateSecretManagerResponse{}
	this.Success = success
	return &this
}

// NewUpdateSecretManagerResponseWithDefaults instantiates a new UpdateSecretManagerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSecretManagerResponseWithDefaults() *UpdateSecretManagerResponse {
	this := UpdateSecretManagerResponse{}
	return &this
}

// GetSuccess returns the Success field value
func (o *UpdateSecretManagerResponse) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *UpdateSecretManagerResponse) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *UpdateSecretManagerResponse) SetSuccess(v bool) {
	o.Success = v
}

func (o UpdateSecretManagerResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSecretManagerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	return toSerialize, nil
}

type NullableUpdateSecretManagerResponse struct {
	value *UpdateSecretManagerResponse
	isSet bool
}

func (v NullableUpdateSecretManagerResponse) Get() *UpdateSecretManagerResponse {
	return v.value
}

func (v *NullableUpdateSecretManagerResponse) Set(val *UpdateSecretManagerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSecretManagerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSecretManagerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSecretManagerResponse(val *UpdateSecretManagerResponse) *NullableUpdateSecretManagerResponse {
	return &NullableUpdateSecretManagerResponse{value: val, isSet: true}
}

func (v NullableUpdateSecretManagerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSecretManagerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


