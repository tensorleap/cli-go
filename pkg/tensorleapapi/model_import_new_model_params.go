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

// checks if the ImportNewModelParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportNewModelParams{}

// ImportNewModelParams struct for ImportNewModelParams
type ImportNewModelParams struct {
	ProjectId string `json:"projectId"`
	FileName string `json:"fileName"`
	DatasetId *string `json:"datasetId,omitempty"`
	ModelName string `json:"modelName"`
	VersionName string `json:"versionName"`
	BranchName *string `json:"branchName,omitempty"`
	ModelType ImportModelType `json:"model_type"`
	TransformInputs *bool `json:"transform_inputs,omitempty"`
}

// NewImportNewModelParams instantiates a new ImportNewModelParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportNewModelParams(projectId string, fileName string, modelName string, versionName string, modelType ImportModelType) *ImportNewModelParams {
	this := ImportNewModelParams{}
	this.ProjectId = projectId
	this.FileName = fileName
	this.ModelName = modelName
	this.VersionName = versionName
	this.ModelType = modelType
	return &this
}

// NewImportNewModelParamsWithDefaults instantiates a new ImportNewModelParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportNewModelParamsWithDefaults() *ImportNewModelParams {
	this := ImportNewModelParams{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *ImportNewModelParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *ImportNewModelParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *ImportNewModelParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetFileName returns the FileName field value
func (o *ImportNewModelParams) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *ImportNewModelParams) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *ImportNewModelParams) SetFileName(v string) {
	o.FileName = v
}

// GetDatasetId returns the DatasetId field value if set, zero value otherwise.
func (o *ImportNewModelParams) GetDatasetId() string {
	if o == nil || IsNil(o.DatasetId) {
		var ret string
		return ret
	}
	return *o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportNewModelParams) GetDatasetIdOk() (*string, bool) {
	if o == nil || IsNil(o.DatasetId) {
		return nil, false
	}
	return o.DatasetId, true
}

// HasDatasetId returns a boolean if a field has been set.
func (o *ImportNewModelParams) HasDatasetId() bool {
	if o != nil && !IsNil(o.DatasetId) {
		return true
	}

	return false
}

// SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.
func (o *ImportNewModelParams) SetDatasetId(v string) {
	o.DatasetId = &v
}

// GetModelName returns the ModelName field value
func (o *ImportNewModelParams) GetModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value
// and a boolean to check if the value has been set.
func (o *ImportNewModelParams) GetModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelName, true
}

// SetModelName sets field value
func (o *ImportNewModelParams) SetModelName(v string) {
	o.ModelName = v
}

// GetVersionName returns the VersionName field value
func (o *ImportNewModelParams) GetVersionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionName
}

// GetVersionNameOk returns a tuple with the VersionName field value
// and a boolean to check if the value has been set.
func (o *ImportNewModelParams) GetVersionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionName, true
}

// SetVersionName sets field value
func (o *ImportNewModelParams) SetVersionName(v string) {
	o.VersionName = v
}

// GetBranchName returns the BranchName field value if set, zero value otherwise.
func (o *ImportNewModelParams) GetBranchName() string {
	if o == nil || IsNil(o.BranchName) {
		var ret string
		return ret
	}
	return *o.BranchName
}

// GetBranchNameOk returns a tuple with the BranchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportNewModelParams) GetBranchNameOk() (*string, bool) {
	if o == nil || IsNil(o.BranchName) {
		return nil, false
	}
	return o.BranchName, true
}

// HasBranchName returns a boolean if a field has been set.
func (o *ImportNewModelParams) HasBranchName() bool {
	if o != nil && !IsNil(o.BranchName) {
		return true
	}

	return false
}

// SetBranchName gets a reference to the given string and assigns it to the BranchName field.
func (o *ImportNewModelParams) SetBranchName(v string) {
	o.BranchName = &v
}

// GetModelType returns the ModelType field value
func (o *ImportNewModelParams) GetModelType() ImportModelType {
	if o == nil {
		var ret ImportModelType
		return ret
	}

	return o.ModelType
}

// GetModelTypeOk returns a tuple with the ModelType field value
// and a boolean to check if the value has been set.
func (o *ImportNewModelParams) GetModelTypeOk() (*ImportModelType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelType, true
}

// SetModelType sets field value
func (o *ImportNewModelParams) SetModelType(v ImportModelType) {
	o.ModelType = v
}

// GetTransformInputs returns the TransformInputs field value if set, zero value otherwise.
func (o *ImportNewModelParams) GetTransformInputs() bool {
	if o == nil || IsNil(o.TransformInputs) {
		var ret bool
		return ret
	}
	return *o.TransformInputs
}

// GetTransformInputsOk returns a tuple with the TransformInputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportNewModelParams) GetTransformInputsOk() (*bool, bool) {
	if o == nil || IsNil(o.TransformInputs) {
		return nil, false
	}
	return o.TransformInputs, true
}

// HasTransformInputs returns a boolean if a field has been set.
func (o *ImportNewModelParams) HasTransformInputs() bool {
	if o != nil && !IsNil(o.TransformInputs) {
		return true
	}

	return false
}

// SetTransformInputs gets a reference to the given bool and assigns it to the TransformInputs field.
func (o *ImportNewModelParams) SetTransformInputs(v bool) {
	o.TransformInputs = &v
}

func (o ImportNewModelParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportNewModelParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["fileName"] = o.FileName
	if !IsNil(o.DatasetId) {
		toSerialize["datasetId"] = o.DatasetId
	}
	toSerialize["modelName"] = o.ModelName
	toSerialize["versionName"] = o.VersionName
	if !IsNil(o.BranchName) {
		toSerialize["branchName"] = o.BranchName
	}
	toSerialize["model_type"] = o.ModelType
	if !IsNil(o.TransformInputs) {
		toSerialize["transform_inputs"] = o.TransformInputs
	}
	return toSerialize, nil
}

type NullableImportNewModelParams struct {
	value *ImportNewModelParams
	isSet bool
}

func (v NullableImportNewModelParams) Get() *ImportNewModelParams {
	return v.value
}

func (v *NullableImportNewModelParams) Set(val *ImportNewModelParams) {
	v.value = val
	v.isSet = true
}

func (v NullableImportNewModelParams) IsSet() bool {
	return v.isSet
}

func (v *NullableImportNewModelParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportNewModelParams(val *ImportNewModelParams) *NullableImportNewModelParams {
	return &NullableImportNewModelParams{value: val, isSet: true}
}

func (v NullableImportNewModelParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportNewModelParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


