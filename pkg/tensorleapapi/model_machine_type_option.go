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

// checks if the MachineTypeOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MachineTypeOption{}

// MachineTypeOption struct for MachineTypeOption
type MachineTypeOption struct {
	Id string `json:"id"`
	DisplayName string `json:"displayName"`
}

// NewMachineTypeOption instantiates a new MachineTypeOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMachineTypeOption(id string, displayName string) *MachineTypeOption {
	this := MachineTypeOption{}
	this.Id = id
	this.DisplayName = displayName
	return &this
}

// NewMachineTypeOptionWithDefaults instantiates a new MachineTypeOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMachineTypeOptionWithDefaults() *MachineTypeOption {
	this := MachineTypeOption{}
	return &this
}

// GetId returns the Id field value
func (o *MachineTypeOption) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MachineTypeOption) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MachineTypeOption) SetId(v string) {
	o.Id = v
}

// GetDisplayName returns the DisplayName field value
func (o *MachineTypeOption) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *MachineTypeOption) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *MachineTypeOption) SetDisplayName(v string) {
	o.DisplayName = v
}

func (o MachineTypeOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MachineTypeOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["displayName"] = o.DisplayName
	return toSerialize, nil
}

type NullableMachineTypeOption struct {
	value *MachineTypeOption
	isSet bool
}

func (v NullableMachineTypeOption) Get() *MachineTypeOption {
	return v.value
}

func (v *NullableMachineTypeOption) Set(val *MachineTypeOption) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineTypeOption) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineTypeOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineTypeOption(val *MachineTypeOption) *NullableMachineTypeOption {
	return &NullableMachineTypeOption{value: val, isSet: true}
}

func (v NullableMachineTypeOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineTypeOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


