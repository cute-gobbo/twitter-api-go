# TweetPublicMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetweetCount** | **int32** | Number of times this Tweet has been Retweeted. | 
**ReplyCount** | **int32** | Number of times this Tweet has been replied to. | 
**LikeCount** | **int32** | Number of times this Tweet has been liked. | 
**QuoteCount** | Pointer to **int32** | Number of times this Tweet has been quoted. | [optional] 

## Methods

### NewTweetPublicMetrics

`func NewTweetPublicMetrics(retweetCount int32, replyCount int32, likeCount int32, ) *TweetPublicMetrics`

NewTweetPublicMetrics instantiates a new TweetPublicMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTweetPublicMetricsWithDefaults

`func NewTweetPublicMetricsWithDefaults() *TweetPublicMetrics`

NewTweetPublicMetricsWithDefaults instantiates a new TweetPublicMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetweetCount

`func (o *TweetPublicMetrics) GetRetweetCount() int32`

GetRetweetCount returns the RetweetCount field if non-nil, zero value otherwise.

### GetRetweetCountOk

`func (o *TweetPublicMetrics) GetRetweetCountOk() (*int32, bool)`

GetRetweetCountOk returns a tuple with the RetweetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetweetCount

`func (o *TweetPublicMetrics) SetRetweetCount(v int32)`

SetRetweetCount sets RetweetCount field to given value.


### GetReplyCount

`func (o *TweetPublicMetrics) GetReplyCount() int32`

GetReplyCount returns the ReplyCount field if non-nil, zero value otherwise.

### GetReplyCountOk

`func (o *TweetPublicMetrics) GetReplyCountOk() (*int32, bool)`

GetReplyCountOk returns a tuple with the ReplyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyCount

`func (o *TweetPublicMetrics) SetReplyCount(v int32)`

SetReplyCount sets ReplyCount field to given value.


### GetLikeCount

`func (o *TweetPublicMetrics) GetLikeCount() int32`

GetLikeCount returns the LikeCount field if non-nil, zero value otherwise.

### GetLikeCountOk

`func (o *TweetPublicMetrics) GetLikeCountOk() (*int32, bool)`

GetLikeCountOk returns a tuple with the LikeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeCount

`func (o *TweetPublicMetrics) SetLikeCount(v int32)`

SetLikeCount sets LikeCount field to given value.


### GetQuoteCount

`func (o *TweetPublicMetrics) GetQuoteCount() int32`

GetQuoteCount returns the QuoteCount field if non-nil, zero value otherwise.

### GetQuoteCountOk

`func (o *TweetPublicMetrics) GetQuoteCountOk() (*int32, bool)`

GetQuoteCountOk returns a tuple with the QuoteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCount

`func (o *TweetPublicMetrics) SetQuoteCount(v int32)`

SetQuoteCount sets QuoteCount field to given value.

### HasQuoteCount

`func (o *TweetPublicMetrics) HasQuoteCount() bool`

HasQuoteCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


