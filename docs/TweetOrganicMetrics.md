# TweetOrganicMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImpressionCount** | **int32** | Number of times this Tweet has been viewed. | 
**RetweetCount** | **int32** | Number of times this Tweet has been Retweeted. | 
**ReplyCount** | **int32** | Number of times this Tweet has been replied to. | 
**LikeCount** | **int32** | Number of times this Tweet has been liked. | 
**UserProfileClicks** | **int32** | Number of times the user&#39;s profile from this Tweet has been clicked. | 
**UrlLinkClicks** | Pointer to **int32** | Number of times links in this Tweet have been clicked. | [optional] 

## Methods

### NewTweetOrganicMetrics

`func NewTweetOrganicMetrics(impressionCount int32, retweetCount int32, replyCount int32, likeCount int32, userProfileClicks int32, ) *TweetOrganicMetrics`

NewTweetOrganicMetrics instantiates a new TweetOrganicMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTweetOrganicMetricsWithDefaults

`func NewTweetOrganicMetricsWithDefaults() *TweetOrganicMetrics`

NewTweetOrganicMetricsWithDefaults instantiates a new TweetOrganicMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImpressionCount

`func (o *TweetOrganicMetrics) GetImpressionCount() int32`

GetImpressionCount returns the ImpressionCount field if non-nil, zero value otherwise.

### GetImpressionCountOk

`func (o *TweetOrganicMetrics) GetImpressionCountOk() (*int32, bool)`

GetImpressionCountOk returns a tuple with the ImpressionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpressionCount

`func (o *TweetOrganicMetrics) SetImpressionCount(v int32)`

SetImpressionCount sets ImpressionCount field to given value.


### GetRetweetCount

`func (o *TweetOrganicMetrics) GetRetweetCount() int32`

GetRetweetCount returns the RetweetCount field if non-nil, zero value otherwise.

### GetRetweetCountOk

`func (o *TweetOrganicMetrics) GetRetweetCountOk() (*int32, bool)`

GetRetweetCountOk returns a tuple with the RetweetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetweetCount

`func (o *TweetOrganicMetrics) SetRetweetCount(v int32)`

SetRetweetCount sets RetweetCount field to given value.


### GetReplyCount

`func (o *TweetOrganicMetrics) GetReplyCount() int32`

GetReplyCount returns the ReplyCount field if non-nil, zero value otherwise.

### GetReplyCountOk

`func (o *TweetOrganicMetrics) GetReplyCountOk() (*int32, bool)`

GetReplyCountOk returns a tuple with the ReplyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyCount

`func (o *TweetOrganicMetrics) SetReplyCount(v int32)`

SetReplyCount sets ReplyCount field to given value.


### GetLikeCount

`func (o *TweetOrganicMetrics) GetLikeCount() int32`

GetLikeCount returns the LikeCount field if non-nil, zero value otherwise.

### GetLikeCountOk

`func (o *TweetOrganicMetrics) GetLikeCountOk() (*int32, bool)`

GetLikeCountOk returns a tuple with the LikeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeCount

`func (o *TweetOrganicMetrics) SetLikeCount(v int32)`

SetLikeCount sets LikeCount field to given value.


### GetUserProfileClicks

`func (o *TweetOrganicMetrics) GetUserProfileClicks() int32`

GetUserProfileClicks returns the UserProfileClicks field if non-nil, zero value otherwise.

### GetUserProfileClicksOk

`func (o *TweetOrganicMetrics) GetUserProfileClicksOk() (*int32, bool)`

GetUserProfileClicksOk returns a tuple with the UserProfileClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProfileClicks

`func (o *TweetOrganicMetrics) SetUserProfileClicks(v int32)`

SetUserProfileClicks sets UserProfileClicks field to given value.


### GetUrlLinkClicks

`func (o *TweetOrganicMetrics) GetUrlLinkClicks() int32`

GetUrlLinkClicks returns the UrlLinkClicks field if non-nil, zero value otherwise.

### GetUrlLinkClicksOk

`func (o *TweetOrganicMetrics) GetUrlLinkClicksOk() (*int32, bool)`

GetUrlLinkClicksOk returns a tuple with the UrlLinkClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlLinkClicks

`func (o *TweetOrganicMetrics) SetUrlLinkClicks(v int32)`

SetUrlLinkClicks sets UrlLinkClicks field to given value.

### HasUrlLinkClicks

`func (o *TweetOrganicMetrics) HasUrlLinkClicks() bool`

HasUrlLinkClicks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


