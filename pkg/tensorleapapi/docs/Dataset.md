# Dataset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** |  | 
**Name** | **NullableString** |  | 
**TeamId** | **string** |  | 
**SourceDatasetId** | Pointer to **string** |  | [optional] 
**Access** | [**DatasetAccess**](DatasetAccess.md) |  | 
**CreatedAt** | **string** |  | 
**LatestValidVersion** | Pointer to [**DatasetVersion**](DatasetVersion.md) |  | [optional] 
**LatestVersion** | Pointer to [**DatasetVersion**](DatasetVersion.md) |  | [optional] 

## Methods

### NewDataset

`func NewDataset(cid string, name NullableString, teamId string, access DatasetAccess, createdAt string, ) *Dataset`

NewDataset instantiates a new Dataset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetWithDefaults

`func NewDatasetWithDefaults() *Dataset`

NewDatasetWithDefaults instantiates a new Dataset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *Dataset) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *Dataset) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *Dataset) SetCid(v string)`

SetCid sets Cid field to given value.


### GetName

`func (o *Dataset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dataset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dataset) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Dataset) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Dataset) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTeamId

`func (o *Dataset) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *Dataset) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *Dataset) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetSourceDatasetId

`func (o *Dataset) GetSourceDatasetId() string`

GetSourceDatasetId returns the SourceDatasetId field if non-nil, zero value otherwise.

### GetSourceDatasetIdOk

`func (o *Dataset) GetSourceDatasetIdOk() (*string, bool)`

GetSourceDatasetIdOk returns a tuple with the SourceDatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDatasetId

`func (o *Dataset) SetSourceDatasetId(v string)`

SetSourceDatasetId sets SourceDatasetId field to given value.

### HasSourceDatasetId

`func (o *Dataset) HasSourceDatasetId() bool`

HasSourceDatasetId returns a boolean if a field has been set.

### GetAccess

`func (o *Dataset) GetAccess() DatasetAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *Dataset) GetAccessOk() (*DatasetAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *Dataset) SetAccess(v DatasetAccess)`

SetAccess sets Access field to given value.


### GetCreatedAt

`func (o *Dataset) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Dataset) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Dataset) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetLatestValidVersion

`func (o *Dataset) GetLatestValidVersion() DatasetVersion`

GetLatestValidVersion returns the LatestValidVersion field if non-nil, zero value otherwise.

### GetLatestValidVersionOk

`func (o *Dataset) GetLatestValidVersionOk() (*DatasetVersion, bool)`

GetLatestValidVersionOk returns a tuple with the LatestValidVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestValidVersion

`func (o *Dataset) SetLatestValidVersion(v DatasetVersion)`

SetLatestValidVersion sets LatestValidVersion field to given value.

### HasLatestValidVersion

`func (o *Dataset) HasLatestValidVersion() bool`

HasLatestValidVersion returns a boolean if a field has been set.

### GetLatestVersion

`func (o *Dataset) GetLatestVersion() DatasetVersion`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *Dataset) GetLatestVersionOk() (*DatasetVersion, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *Dataset) SetLatestVersion(v DatasetVersion)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *Dataset) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


