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

// checks if the CustomLossInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomLossInstance{}

// CustomLossInstance struct for CustomLossInstance
type CustomLossInstance struct {
	Name string `json:"name"`
	ArgNames []string `json:"arg_names"`
}

// NewCustomLossInstance instantiates a new CustomLossInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomLossInstance(name string, argNames []string) *CustomLossInstance {
	this := CustomLossInstance{}
	this.Name = name
	this.ArgNames = argNames
	return &this
}

// NewCustomLossInstanceWithDefaults instantiates a new CustomLossInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomLossInstanceWithDefaults() *CustomLossInstance {
	this := CustomLossInstance{}
	return &this
}

// GetName returns the Name field value
func (o *CustomLossInstance) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomLossInstance) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomLossInstance) SetName(v string) {
	o.Name = v
}

// GetArgNames returns the ArgNames field value
func (o *CustomLossInstance) GetArgNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ArgNames
}

// GetArgNamesOk returns a tuple with the ArgNames field value
// and a boolean to check if the value has been set.
func (o *CustomLossInstance) GetArgNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArgNames, true
}

// SetArgNames sets field value
func (o *CustomLossInstance) SetArgNames(v []string) {
	o.ArgNames = v
}

func (o CustomLossInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomLossInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["arg_names"] = o.ArgNames
	return toSerialize, nil
}

type NullableCustomLossInstance struct {
	value *CustomLossInstance
	isSet bool
}

func (v NullableCustomLossInstance) Get() *CustomLossInstance {
	return v.value
}

func (v *NullableCustomLossInstance) Set(val *CustomLossInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomLossInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomLossInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomLossInstance(val *CustomLossInstance) *NullableCustomLossInstance {
	return &NullableCustomLossInstance{value: val, isSet: true}
}

func (v NullableCustomLossInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomLossInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


