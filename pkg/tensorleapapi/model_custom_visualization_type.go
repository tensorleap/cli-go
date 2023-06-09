/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.342
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// CustomVisualizationType the model 'CustomVisualizationType'
type CustomVisualizationType string

// List of CustomVisualizationType
const (
	CUSTOMVISUALIZATIONTYPE_BAR CustomVisualizationType = "Bar"
	CUSTOMVISUALIZATIONTYPE_LINE CustomVisualizationType = "Line"
	CUSTOMVISUALIZATIONTYPE_AREA CustomVisualizationType = "Area"
	CUSTOMVISUALIZATIONTYPE_TABLE CustomVisualizationType = "Table"
	CUSTOMVISUALIZATIONTYPE_HEATMAP CustomVisualizationType = "Heatmap"
	CUSTOMVISUALIZATIONTYPE_DONUT CustomVisualizationType = "Donut"
	CUSTOMVISUALIZATIONTYPE_CONFUSION_MATRIX CustomVisualizationType = "Confusion Matrix"
)

// All allowed values of CustomVisualizationType enum
var AllowedCustomVisualizationTypeEnumValues = []CustomVisualizationType{
	"Bar",
	"Line",
	"Area",
	"Table",
	"Heatmap",
	"Donut",
	"Confusion Matrix",
}

func (v *CustomVisualizationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomVisualizationType(value)
	for _, existing := range AllowedCustomVisualizationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomVisualizationType", value)
}

// NewCustomVisualizationTypeFromValue returns a pointer to a valid CustomVisualizationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomVisualizationTypeFromValue(v string) (*CustomVisualizationType, error) {
	ev := CustomVisualizationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomVisualizationType: valid values are %v", v, AllowedCustomVisualizationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomVisualizationType) IsValid() bool {
	for _, existing := range AllowedCustomVisualizationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomVisualizationType value
func (v CustomVisualizationType) Ptr() *CustomVisualizationType {
	return &v
}

type NullableCustomVisualizationType struct {
	value *CustomVisualizationType
	isSet bool
}

func (v NullableCustomVisualizationType) Get() *CustomVisualizationType {
	return v.value
}

func (v *NullableCustomVisualizationType) Set(val *CustomVisualizationType) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomVisualizationType) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomVisualizationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomVisualizationType(val *CustomVisualizationType) *NullableCustomVisualizationType {
	return &NullableCustomVisualizationType{value: val, isSet: true}
}

func (v NullableCustomVisualizationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomVisualizationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

