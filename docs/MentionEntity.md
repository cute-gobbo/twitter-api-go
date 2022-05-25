# MentionEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **int32** | Index (zero-based) at which position this entity starts. | 
**End** | **int32** | Index (zero-based) at which position this entity ends. | 
**Username** | **string** | The Twitter handle (screen name) of this user. | 
**Id** | **string** | Unique identifier of this User. This is returned as a string in order to avoid complications with languages and tools that cannot handle large integers. | 

## Methods

### NewMentionEntity

`func NewMentionEntity(start int32, end int32, username string, id string, ) *MentionEntity`

NewMentionEntity instantiates a new MentionEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMentionEntityWithDefaults

`func NewMentionEntityWithDefaults() *MentionEntity`

NewMentionEntityWithDefaults instantiates a new MentionEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *MentionEntity) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *MentionEntity) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *MentionEntity) SetStart(v int32)`

SetStart sets Start field to given value.


### GetEnd

`func (o *MentionEntity) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *MentionEntity) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *MentionEntity) SetEnd(v int32)`

SetEnd sets End field to given value.


### GetUsername

`func (o *MentionEntity) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MentionEntity) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MentionEntity) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetId

`func (o *MentionEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MentionEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MentionEntity) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


