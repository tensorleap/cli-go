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

// checks if the GetJobsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetJobsResponse{}

// GetJobsResponse struct for GetJobsResponse
type GetJobsResponse struct {
	Jobs []Job `json:"jobs"`
	// Construct a type with a set of properties K of type T
	MapVersionToProjectName map[string]interface{} `json:"mapVersionToProjectName"`
	IncludeProject bool `json:"includeProject"`
	IsModelJobs bool `json:"isModelJobs"`
}

// NewGetJobsResponse instantiates a new GetJobsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetJobsResponse(jobs []Job, mapVersionToProjectName map[string]interface{}, includeProject bool, isModelJobs bool) *GetJobsResponse {
	this := GetJobsResponse{}
	this.Jobs = jobs
	this.MapVersionToProjectName = mapVersionToProjectName
	this.IncludeProject = includeProject
	this.IsModelJobs = isModelJobs
	return &this
}

// NewGetJobsResponseWithDefaults instantiates a new GetJobsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetJobsResponseWithDefaults() *GetJobsResponse {
	this := GetJobsResponse{}
	return &this
}

// GetJobs returns the Jobs field value
func (o *GetJobsResponse) GetJobs() []Job {
	if o == nil {
		var ret []Job
		return ret
	}

	return o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value
// and a boolean to check if the value has been set.
func (o *GetJobsResponse) GetJobsOk() ([]Job, bool) {
	if o == nil {
		return nil, false
	}
	return o.Jobs, true
}

// SetJobs sets field value
func (o *GetJobsResponse) SetJobs(v []Job) {
	o.Jobs = v
}

// GetMapVersionToProjectName returns the MapVersionToProjectName field value
func (o *GetJobsResponse) GetMapVersionToProjectName() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.MapVersionToProjectName
}

// GetMapVersionToProjectNameOk returns a tuple with the MapVersionToProjectName field value
// and a boolean to check if the value has been set.
func (o *GetJobsResponse) GetMapVersionToProjectNameOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.MapVersionToProjectName, true
}

// SetMapVersionToProjectName sets field value
func (o *GetJobsResponse) SetMapVersionToProjectName(v map[string]interface{}) {
	o.MapVersionToProjectName = v
}

// GetIncludeProject returns the IncludeProject field value
func (o *GetJobsResponse) GetIncludeProject() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IncludeProject
}

// GetIncludeProjectOk returns a tuple with the IncludeProject field value
// and a boolean to check if the value has been set.
func (o *GetJobsResponse) GetIncludeProjectOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludeProject, true
}

// SetIncludeProject sets field value
func (o *GetJobsResponse) SetIncludeProject(v bool) {
	o.IncludeProject = v
}

// GetIsModelJobs returns the IsModelJobs field value
func (o *GetJobsResponse) GetIsModelJobs() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsModelJobs
}

// GetIsModelJobsOk returns a tuple with the IsModelJobs field value
// and a boolean to check if the value has been set.
func (o *GetJobsResponse) GetIsModelJobsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsModelJobs, true
}

// SetIsModelJobs sets field value
func (o *GetJobsResponse) SetIsModelJobs(v bool) {
	o.IsModelJobs = v
}

func (o GetJobsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetJobsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jobs"] = o.Jobs
	toSerialize["mapVersionToProjectName"] = o.MapVersionToProjectName
	toSerialize["includeProject"] = o.IncludeProject
	toSerialize["isModelJobs"] = o.IsModelJobs
	return toSerialize, nil
}

type NullableGetJobsResponse struct {
	value *GetJobsResponse
	isSet bool
}

func (v NullableGetJobsResponse) Get() *GetJobsResponse {
	return v.value
}

func (v *NullableGetJobsResponse) Set(val *GetJobsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetJobsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetJobsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetJobsResponse(val *GetJobsResponse) *NullableGetJobsResponse {
	return &NullableGetJobsResponse{value: val, isSet: true}
}

func (v NullableGetJobsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetJobsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


