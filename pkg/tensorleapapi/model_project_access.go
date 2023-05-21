/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.301
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// ProjectAccess the model 'ProjectAccess'
type ProjectAccess string

// List of ProjectAccess
const (
	PROJECTACCESS_PUBLIC ProjectAccess = "public"
	PROJECTACCESS_LOCAL ProjectAccess = "local"
)

// All allowed values of ProjectAccess enum
var AllowedProjectAccessEnumValues = []ProjectAccess{
	"public",
	"local",
}

func (v *ProjectAccess) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProjectAccess(value)
	for _, existing := range AllowedProjectAccessEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProjectAccess", value)
}

// NewProjectAccessFromValue returns a pointer to a valid ProjectAccess
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectAccessFromValue(v string) (*ProjectAccess, error) {
	ev := ProjectAccess(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectAccess: valid values are %v", v, AllowedProjectAccessEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectAccess) IsValid() bool {
	for _, existing := range AllowedProjectAccessEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProjectAccess value
func (v ProjectAccess) Ptr() *ProjectAccess {
	return &v
}

type NullableProjectAccess struct {
	value *ProjectAccess
	isSet bool
}

func (v NullableProjectAccess) Get() *ProjectAccess {
	return v.value
}

func (v *NullableProjectAccess) Set(val *ProjectAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectAccess(val *ProjectAccess) *NullableProjectAccess {
	return &NullableProjectAccess{value: val, isSet: true}
}

func (v NullableProjectAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

