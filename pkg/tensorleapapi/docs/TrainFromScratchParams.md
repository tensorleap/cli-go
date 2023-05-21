# TrainFromScratchParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionId** | **string** |  | 
**SessionName** | **string** |  | 
**TrainingParams** | [**TrainingParams**](TrainingParams.md) |  | 
**ShouldRunPopulationExploration** | **bool** |  | 

## Methods

### NewTrainFromScratchParams

`func NewTrainFromScratchParams(versionId string, sessionName string, trainingParams TrainingParams, shouldRunPopulationExploration bool, ) *TrainFromScratchParams`

NewTrainFromScratchParams instantiates a new TrainFromScratchParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainFromScratchParamsWithDefaults

`func NewTrainFromScratchParamsWithDefaults() *TrainFromScratchParams`

NewTrainFromScratchParamsWithDefaults instantiates a new TrainFromScratchParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionId

`func (o *TrainFromScratchParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *TrainFromScratchParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *TrainFromScratchParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetSessionName

`func (o *TrainFromScratchParams) GetSessionName() string`

GetSessionName returns the SessionName field if non-nil, zero value otherwise.

### GetSessionNameOk

`func (o *TrainFromScratchParams) GetSessionNameOk() (*string, bool)`

GetSessionNameOk returns a tuple with the SessionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionName

`func (o *TrainFromScratchParams) SetSessionName(v string)`

SetSessionName sets SessionName field to given value.


### GetTrainingParams

`func (o *TrainFromScratchParams) GetTrainingParams() TrainingParams`

GetTrainingParams returns the TrainingParams field if non-nil, zero value otherwise.

### GetTrainingParamsOk

`func (o *TrainFromScratchParams) GetTrainingParamsOk() (*TrainingParams, bool)`

GetTrainingParamsOk returns a tuple with the TrainingParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingParams

`func (o *TrainFromScratchParams) SetTrainingParams(v TrainingParams)`

SetTrainingParams sets TrainingParams field to given value.


### GetShouldRunPopulationExploration

`func (o *TrainFromScratchParams) GetShouldRunPopulationExploration() bool`

GetShouldRunPopulationExploration returns the ShouldRunPopulationExploration field if non-nil, zero value otherwise.

### GetShouldRunPopulationExplorationOk

`func (o *TrainFromScratchParams) GetShouldRunPopulationExplorationOk() (*bool, bool)`

GetShouldRunPopulationExplorationOk returns a tuple with the ShouldRunPopulationExploration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldRunPopulationExploration

`func (o *TrainFromScratchParams) SetShouldRunPopulationExploration(v bool)`

SetShouldRunPopulationExploration sets ShouldRunPopulationExploration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


