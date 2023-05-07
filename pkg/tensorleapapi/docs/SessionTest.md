# SessionTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ProjectId** | **string** |  | 
**OrganizationId** | **string** |  | 
**Name** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**CreatedBy** | **string** |  | 
**TestFilter** | [**ClientFilterParams**](ClientFilterParams.md) |  | 
**DatasetFilter** | [**[]ESFilter**](ESFilter.md) |  | 

## Methods

### NewSessionTest

`func NewSessionTest(id string, projectId string, organizationId string, name string, createdAt time.Time, updatedAt time.Time, createdBy string, testFilter ClientFilterParams, datasetFilter []ESFilter, ) *SessionTest`

NewSessionTest instantiates a new SessionTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionTestWithDefaults

`func NewSessionTestWithDefaults() *SessionTest`

NewSessionTestWithDefaults instantiates a new SessionTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SessionTest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionTest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionTest) SetId(v string)`

SetId sets Id field to given value.


### GetProjectId

`func (o *SessionTest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SessionTest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SessionTest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetOrganizationId

`func (o *SessionTest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SessionTest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SessionTest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetName

`func (o *SessionTest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SessionTest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SessionTest) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *SessionTest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SessionTest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SessionTest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SessionTest) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SessionTest) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SessionTest) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedBy

`func (o *SessionTest) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SessionTest) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SessionTest) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetTestFilter

`func (o *SessionTest) GetTestFilter() ClientFilterParams`

GetTestFilter returns the TestFilter field if non-nil, zero value otherwise.

### GetTestFilterOk

`func (o *SessionTest) GetTestFilterOk() (*ClientFilterParams, bool)`

GetTestFilterOk returns a tuple with the TestFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestFilter

`func (o *SessionTest) SetTestFilter(v ClientFilterParams)`

SetTestFilter sets TestFilter field to given value.


### GetDatasetFilter

`func (o *SessionTest) GetDatasetFilter() []ESFilter`

GetDatasetFilter returns the DatasetFilter field if non-nil, zero value otherwise.

### GetDatasetFilterOk

`func (o *SessionTest) GetDatasetFilterOk() (*[]ESFilter, bool)`

GetDatasetFilterOk returns a tuple with the DatasetFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetFilter

`func (o *SessionTest) SetDatasetFilter(v []ESFilter)`

SetDatasetFilter sets DatasetFilter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


