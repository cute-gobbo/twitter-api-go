# TweetAttachments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaKeys** | Pointer to **[]string** | A list of Media Keys for each one of the media attachments (if media are attached). | [optional] 
**PollIds** | Pointer to **[]string** | A list of poll IDs (if polls are attached). | [optional] 

## Methods

### NewTweetAttachments

`func NewTweetAttachments() *TweetAttachments`

NewTweetAttachments instantiates a new TweetAttachments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTweetAttachmentsWithDefaults

`func NewTweetAttachmentsWithDefaults() *TweetAttachments`

NewTweetAttachmentsWithDefaults instantiates a new TweetAttachments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaKeys

`func (o *TweetAttachments) GetMediaKeys() []string`

GetMediaKeys returns the MediaKeys field if non-nil, zero value otherwise.

### GetMediaKeysOk

`func (o *TweetAttachments) GetMediaKeysOk() (*[]string, bool)`

GetMediaKeysOk returns a tuple with the MediaKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaKeys

`func (o *TweetAttachments) SetMediaKeys(v []string)`

SetMediaKeys sets MediaKeys field to given value.

### HasMediaKeys

`func (o *TweetAttachments) HasMediaKeys() bool`

HasMediaKeys returns a boolean if a field has been set.

### GetPollIds

`func (o *TweetAttachments) GetPollIds() []string`

GetPollIds returns the PollIds field if non-nil, zero value otherwise.

### GetPollIdsOk

`func (o *TweetAttachments) GetPollIdsOk() (*[]string, bool)`

GetPollIdsOk returns a tuple with the PollIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollIds

`func (o *TweetAttachments) SetPollIds(v []string)`

SetPollIds sets PollIds field to given value.

### HasPollIds

`func (o *TweetAttachments) HasPollIds() bool`

HasPollIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


