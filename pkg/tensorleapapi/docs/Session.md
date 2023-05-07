# Session

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ExtId** | Pointer to **string** |  | [optional] 
**ModelName** | **string** |  | 
**Type** | [**ModelType**](ModelType.md) |  | 
**CreatedAt** | **time.Time** |  | 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Organization** | **string** |  | 
**Hash** | Pointer to **NullableString** |  | [optional] 
**TrainingParams** | Pointer to [**TrainingParams**](TrainingParams.md) |  | [optional] 
**SessionRuns** | Pointer to [**[]SessionRunData**](SessionRunData.md) |  | [optional] 

## Methods

### NewSession

`func NewSession(id string, modelName string, type_ ModelType, createdAt time.Time, organization string, ) *Session`

NewSession instantiates a new Session object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionWithDefaults

`func NewSessionWithDefaults() *Session`

NewSessionWithDefaults instantiates a new Session object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Session) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Session) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Session) SetId(v string)`

SetId sets Id field to given value.


### GetExtId

`func (o *Session) GetExtId() string`

GetExtId returns the ExtId field if non-nil, zero value otherwise.

### GetExtIdOk

`func (o *Session) GetExtIdOk() (*string, bool)`

GetExtIdOk returns a tuple with the ExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtId

`func (o *Session) SetExtId(v string)`

SetExtId sets ExtId field to given value.

### HasExtId

`func (o *Session) HasExtId() bool`

HasExtId returns a boolean if a field has been set.

### GetModelName

`func (o *Session) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *Session) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *Session) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetType

`func (o *Session) GetType() ModelType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Session) GetTypeOk() (*ModelType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Session) SetType(v ModelType)`

SetType sets Type field to given value.


### GetCreatedAt

`func (o *Session) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Session) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Session) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Session) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Session) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Session) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Session) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetOrganization

`func (o *Session) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Session) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Session) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetHash

`func (o *Session) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Session) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Session) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *Session) HasHash() bool`

HasHash returns a boolean if a field has been set.

### SetHashNil

`func (o *Session) SetHashNil(b bool)`

 SetHashNil sets the value for Hash to be an explicit nil

### UnsetHash
`func (o *Session) UnsetHash()`

UnsetHash ensures that no value is present for Hash, not even an explicit nil
### GetTrainingParams

`func (o *Session) GetTrainingParams() TrainingParams`

GetTrainingParams returns the TrainingParams field if non-nil, zero value otherwise.

### GetTrainingParamsOk

`func (o *Session) GetTrainingParamsOk() (*TrainingParams, bool)`

GetTrainingParamsOk returns a tuple with the TrainingParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingParams

`func (o *Session) SetTrainingParams(v TrainingParams)`

SetTrainingParams sets TrainingParams field to given value.

### HasTrainingParams

`func (o *Session) HasTrainingParams() bool`

HasTrainingParams returns a boolean if a field has been set.

### GetSessionRuns

`func (o *Session) GetSessionRuns() []SessionRunData`

GetSessionRuns returns the SessionRuns field if non-nil, zero value otherwise.

### GetSessionRunsOk

`func (o *Session) GetSessionRunsOk() (*[]SessionRunData, bool)`

GetSessionRunsOk returns a tuple with the SessionRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRuns

`func (o *Session) SetSessionRuns(v []SessionRunData)`

SetSessionRuns sets SessionRuns field to given value.

### HasSessionRuns

`func (o *Session) HasSessionRuns() bool`

HasSessionRuns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


