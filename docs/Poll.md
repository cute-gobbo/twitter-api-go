# Poll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of this poll. | 
**Options** | [**[]PollOption**](PollOption.md) |  | 
**VotingStatus** | Pointer to **string** |  | [optional] 
**EndDatetime** | Pointer to **time.Time** |  | [optional] 
**DurationMinutes** | Pointer to **int32** |  | [optional] 

## Methods

### NewPoll

`func NewPoll(id string, options []PollOption, ) *Poll`

NewPoll instantiates a new Poll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPollWithDefaults

`func NewPollWithDefaults() *Poll`

NewPollWithDefaults instantiates a new Poll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Poll) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Poll) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Poll) SetId(v string)`

SetId sets Id field to given value.


### GetOptions

`func (o *Poll) GetOptions() []PollOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Poll) GetOptionsOk() (*[]PollOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Poll) SetOptions(v []PollOption)`

SetOptions sets Options field to given value.


### GetVotingStatus

`func (o *Poll) GetVotingStatus() string`

GetVotingStatus returns the VotingStatus field if non-nil, zero value otherwise.

### GetVotingStatusOk

`func (o *Poll) GetVotingStatusOk() (*string, bool)`

GetVotingStatusOk returns a tuple with the VotingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotingStatus

`func (o *Poll) SetVotingStatus(v string)`

SetVotingStatus sets VotingStatus field to given value.

### HasVotingStatus

`func (o *Poll) HasVotingStatus() bool`

HasVotingStatus returns a boolean if a field has been set.

### GetEndDatetime

`func (o *Poll) GetEndDatetime() time.Time`

GetEndDatetime returns the EndDatetime field if non-nil, zero value otherwise.

### GetEndDatetimeOk

`func (o *Poll) GetEndDatetimeOk() (*time.Time, bool)`

GetEndDatetimeOk returns a tuple with the EndDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDatetime

`func (o *Poll) SetEndDatetime(v time.Time)`

SetEndDatetime sets EndDatetime field to given value.

### HasEndDatetime

`func (o *Poll) HasEndDatetime() bool`

HasEndDatetime returns a boolean if a field has been set.

### GetDurationMinutes

`func (o *Poll) GetDurationMinutes() int32`

GetDurationMinutes returns the DurationMinutes field if non-nil, zero value otherwise.

### GetDurationMinutesOk

`func (o *Poll) GetDurationMinutesOk() (*int32, bool)`

GetDurationMinutesOk returns a tuple with the DurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMinutes

`func (o *Poll) SetDurationMinutes(v int32)`

SetDurationMinutes sets DurationMinutes field to given value.

### HasDurationMinutes

`func (o *Poll) HasDurationMinutes() bool`

HasDurationMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


