# UserData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** |  | 
**RecordSession** | **bool** |  | 
**Role** | [**UserRole**](UserRole.md) |  | 
**Local** | [**UserDataLocal**](UserDataLocal.md) |  | 
**Metadata** | Pointer to [**UserDataMetadata**](UserDataMetadata.md) |  | [optional] 
**TeamName** | **string** |  | 
**SingleUserMode** | Pointer to [**SingleUserModeSettings**](SingleUserModeSettings.md) |  | [optional] 
**Activated** | **bool** |  | 

## Methods

### NewUserData

`func NewUserData(cid string, recordSession bool, role UserRole, local UserDataLocal, teamName string, activated bool, ) *UserData`

NewUserData instantiates a new UserData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDataWithDefaults

`func NewUserDataWithDefaults() *UserData`

NewUserDataWithDefaults instantiates a new UserData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *UserData) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *UserData) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *UserData) SetCid(v string)`

SetCid sets Cid field to given value.


### GetRecordSession

`func (o *UserData) GetRecordSession() bool`

GetRecordSession returns the RecordSession field if non-nil, zero value otherwise.

### GetRecordSessionOk

`func (o *UserData) GetRecordSessionOk() (*bool, bool)`

GetRecordSessionOk returns a tuple with the RecordSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordSession

`func (o *UserData) SetRecordSession(v bool)`

SetRecordSession sets RecordSession field to given value.


### GetRole

`func (o *UserData) GetRole() UserRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserData) GetRoleOk() (*UserRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserData) SetRole(v UserRole)`

SetRole sets Role field to given value.


### GetLocal

`func (o *UserData) GetLocal() UserDataLocal`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *UserData) GetLocalOk() (*UserDataLocal, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *UserData) SetLocal(v UserDataLocal)`

SetLocal sets Local field to given value.


### GetMetadata

`func (o *UserData) GetMetadata() UserDataMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UserData) GetMetadataOk() (*UserDataMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UserData) SetMetadata(v UserDataMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UserData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTeamName

`func (o *UserData) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *UserData) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *UserData) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.


### GetSingleUserMode

`func (o *UserData) GetSingleUserMode() SingleUserModeSettings`

GetSingleUserMode returns the SingleUserMode field if non-nil, zero value otherwise.

### GetSingleUserModeOk

`func (o *UserData) GetSingleUserModeOk() (*SingleUserModeSettings, bool)`

GetSingleUserModeOk returns a tuple with the SingleUserMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleUserMode

`func (o *UserData) SetSingleUserMode(v SingleUserModeSettings)`

SetSingleUserMode sets SingleUserMode field to given value.

### HasSingleUserMode

`func (o *UserData) HasSingleUserMode() bool`

HasSingleUserMode returns a boolean if a field has been set.

### GetActivated

`func (o *UserData) GetActivated() bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *UserData) GetActivatedOk() (*bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *UserData) SetActivated(v bool)`

SetActivated sets Activated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


