# GetSecretManagerListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]SecretManager**](SecretManager.md) |  | 

## Methods

### NewGetSecretManagerListResponse

`func NewGetSecretManagerListResponse(results []SecretManager, ) *GetSecretManagerListResponse`

NewGetSecretManagerListResponse instantiates a new GetSecretManagerListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSecretManagerListResponseWithDefaults

`func NewGetSecretManagerListResponseWithDefaults() *GetSecretManagerListResponse`

NewGetSecretManagerListResponseWithDefaults instantiates a new GetSecretManagerListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *GetSecretManagerListResponse) GetResults() []SecretManager`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetSecretManagerListResponse) GetResultsOk() (*[]SecretManager, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetSecretManagerListResponse) SetResults(v []SecretManager)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


