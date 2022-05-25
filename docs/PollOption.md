# PollOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Position** | **int32** | Position of this choice in the poll. | 
**Label** | **string** | The text of a poll choice. | 
**Votes** | **int32** | Number of users who voted for this choice. | 

## Methods

### NewPollOption

`func NewPollOption(position int32, label string, votes int32, ) *PollOption`

NewPollOption instantiates a new PollOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPollOptionWithDefaults

`func NewPollOptionWithDefaults() *PollOption`

NewPollOptionWithDefaults instantiates a new PollOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosition

`func (o *PollOption) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PollOption) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PollOption) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetLabel

`func (o *PollOption) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PollOption) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PollOption) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetVotes

`func (o *PollOption) GetVotes() int32`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *PollOption) GetVotesOk() (*int32, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *PollOption) SetVotes(v int32)`

SetVotes sets Votes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


