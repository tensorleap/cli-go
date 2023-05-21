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

// IssueActionType the model 'IssueActionType'
type IssueActionType string

// List of IssueActionType
const (
	ISSUEACTIONTYPE_TAGS_ADDED IssueActionType = "TagsAdded"
	ISSUEACTIONTYPE_ASSIGNMENT IssueActionType = "Assignment"
	ISSUEACTIONTYPE_COMMENT IssueActionType = "Comment"
)

// All allowed values of IssueActionType enum
var AllowedIssueActionTypeEnumValues = []IssueActionType{
	"TagsAdded",
	"Assignment",
	"Comment",
}

func (v *IssueActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IssueActionType(value)
	for _, existing := range AllowedIssueActionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IssueActionType", value)
}

// NewIssueActionTypeFromValue returns a pointer to a valid IssueActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIssueActionTypeFromValue(v string) (*IssueActionType, error) {
	ev := IssueActionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IssueActionType: valid values are %v", v, AllowedIssueActionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IssueActionType) IsValid() bool {
	for _, existing := range AllowedIssueActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssueActionType value
func (v IssueActionType) Ptr() *IssueActionType {
	return &v
}

type NullableIssueActionType struct {
	value *IssueActionType
	isSet bool
}

func (v NullableIssueActionType) Get() *IssueActionType {
	return v.value
}

func (v *NullableIssueActionType) Set(val *IssueActionType) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueActionType) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueActionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueActionType(val *IssueActionType) *NullableIssueActionType {
	return &NullableIssueActionType{value: val, isSet: true}
}

func (v NullableIssueActionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueActionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

