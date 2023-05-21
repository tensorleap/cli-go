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

// ExportModelTypeEnum the model 'ExportModelTypeEnum'
type ExportModelTypeEnum string

// List of ExportModelTypeEnum
const (
	EXPORTMODELTYPEENUM_JSON_TF2 ExportModelTypeEnum = "JSON_TF2"
	EXPORTMODELTYPEENUM_H5_TF2 ExportModelTypeEnum = "H5_TF2"
	EXPORTMODELTYPEENUM_ONNX ExportModelTypeEnum = "ONNX"
	EXPORTMODELTYPEENUM_SAVED_MODEL_TF2 ExportModelTypeEnum = "SavedModel_TF2"
)

// All allowed values of ExportModelTypeEnum enum
var AllowedExportModelTypeEnumEnumValues = []ExportModelTypeEnum{
	"JSON_TF2",
	"H5_TF2",
	"ONNX",
	"SavedModel_TF2",
}

func (v *ExportModelTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExportModelTypeEnum(value)
	for _, existing := range AllowedExportModelTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExportModelTypeEnum", value)
}

// NewExportModelTypeEnumFromValue returns a pointer to a valid ExportModelTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExportModelTypeEnumFromValue(v string) (*ExportModelTypeEnum, error) {
	ev := ExportModelTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExportModelTypeEnum: valid values are %v", v, AllowedExportModelTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExportModelTypeEnum) IsValid() bool {
	for _, existing := range AllowedExportModelTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExportModelTypeEnum value
func (v ExportModelTypeEnum) Ptr() *ExportModelTypeEnum {
	return &v
}

type NullableExportModelTypeEnum struct {
	value *ExportModelTypeEnum
	isSet bool
}

func (v NullableExportModelTypeEnum) Get() *ExportModelTypeEnum {
	return v.value
}

func (v *NullableExportModelTypeEnum) Set(val *ExportModelTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableExportModelTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableExportModelTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportModelTypeEnum(val *ExportModelTypeEnum) *NullableExportModelTypeEnum {
	return &NullableExportModelTypeEnum{value: val, isSet: true}
}

func (v NullableExportModelTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportModelTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

