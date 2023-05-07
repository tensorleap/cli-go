# DatasetVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Organization** | **string** |  | 
**DatasetId** | **string** |  | 
**Note** | **string** |  | 
**CreatedAt** | **string** |  | 
**CreatedBy** | **string** |  | 
**TestStatus** | [**TestStatus**](TestStatus.md) |  | 
**Metadata** | [**DatasetMetadata**](DatasetMetadata.md) |  | 
**CodeUrl** | **string** |  | 
**CodeEntryFile** | **string** |  | 

## Methods

### NewDatasetVersion

`func NewDatasetVersion(id string, organization string, datasetId string, note string, createdAt string, createdBy string, testStatus TestStatus, metadata DatasetMetadata, codeUrl string, codeEntryFile string, ) *DatasetVersion`

NewDatasetVersion instantiates a new DatasetVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetVersionWithDefaults

`func NewDatasetVersionWithDefaults() *DatasetVersion`

NewDatasetVersionWithDefaults instantiates a new DatasetVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DatasetVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatasetVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatasetVersion) SetId(v string)`

SetId sets Id field to given value.


### GetOrganization

`func (o *DatasetVersion) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *DatasetVersion) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *DatasetVersion) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetDatasetId

`func (o *DatasetVersion) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *DatasetVersion) GetDatasetIdOk() (*string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetId

`func (o *DatasetVersion) SetDatasetId(v string)`

SetDatasetId sets DatasetId field to given value.


### GetNote

`func (o *DatasetVersion) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *DatasetVersion) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *DatasetVersion) SetNote(v string)`

SetNote sets Note field to given value.


### GetCreatedAt

`func (o *DatasetVersion) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DatasetVersion) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DatasetVersion) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *DatasetVersion) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DatasetVersion) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DatasetVersion) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetTestStatus

`func (o *DatasetVersion) GetTestStatus() TestStatus`

GetTestStatus returns the TestStatus field if non-nil, zero value otherwise.

### GetTestStatusOk

`func (o *DatasetVersion) GetTestStatusOk() (*TestStatus, bool)`

GetTestStatusOk returns a tuple with the TestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestStatus

`func (o *DatasetVersion) SetTestStatus(v TestStatus)`

SetTestStatus sets TestStatus field to given value.


### GetMetadata

`func (o *DatasetVersion) GetMetadata() DatasetMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DatasetVersion) GetMetadataOk() (*DatasetMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DatasetVersion) SetMetadata(v DatasetMetadata)`

SetMetadata sets Metadata field to given value.


### GetCodeUrl

`func (o *DatasetVersion) GetCodeUrl() string`

GetCodeUrl returns the CodeUrl field if non-nil, zero value otherwise.

### GetCodeUrlOk

`func (o *DatasetVersion) GetCodeUrlOk() (*string, bool)`

GetCodeUrlOk returns a tuple with the CodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeUrl

`func (o *DatasetVersion) SetCodeUrl(v string)`

SetCodeUrl sets CodeUrl field to given value.


### GetCodeEntryFile

`func (o *DatasetVersion) GetCodeEntryFile() string`

GetCodeEntryFile returns the CodeEntryFile field if non-nil, zero value otherwise.

### GetCodeEntryFileOk

`func (o *DatasetVersion) GetCodeEntryFileOk() (*string, bool)`

GetCodeEntryFileOk returns a tuple with the CodeEntryFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeEntryFile

`func (o *DatasetVersion) SetCodeEntryFile(v string)`

SetCodeEntryFile sets CodeEntryFile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


