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

// checks if the GetAllProjectSessionTestsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllProjectSessionTestsRequest{}

// GetAllProjectSessionTestsRequest struct for GetAllProjectSessionTestsRequest
type GetAllProjectSessionTestsRequest struct {
	ProjectId string `json:"projectId"`
}

// NewGetAllProjectSessionTestsRequest instantiates a new GetAllProjectSessionTestsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllProjectSessionTestsRequest(projectId string) *GetAllProjectSessionTestsRequest {
	this := GetAllProjectSessionTestsRequest{}
	this.ProjectId = projectId
	return &this
}

// NewGetAllProjectSessionTestsRequestWithDefaults instantiates a new GetAllProjectSessionTestsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllProjectSessionTestsRequestWithDefaults() *GetAllProjectSessionTestsRequest {
	this := GetAllProjectSessionTestsRequest{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *GetAllProjectSessionTestsRequest) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *GetAllProjectSessionTestsRequest) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *GetAllProjectSessionTestsRequest) SetProjectId(v string) {
	o.ProjectId = v
}

func (o GetAllProjectSessionTestsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllProjectSessionTestsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableGetAllProjectSessionTestsRequest struct {
	value *GetAllProjectSessionTestsRequest
	isSet bool
}

func (v NullableGetAllProjectSessionTestsRequest) Get() *GetAllProjectSessionTestsRequest {
	return v.value
}

func (v *NullableGetAllProjectSessionTestsRequest) Set(val *GetAllProjectSessionTestsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllProjectSessionTestsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllProjectSessionTestsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllProjectSessionTestsRequest(val *GetAllProjectSessionTestsRequest) *NullableGetAllProjectSessionTestsRequest {
	return &NullableGetAllProjectSessionTestsRequest{value: val, isSet: true}
}

func (v NullableGetAllProjectSessionTestsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllProjectSessionTestsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


