# UserPublicMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FollowersCount** | **int32** | Number of users who are following this user. | 
**FollowingCount** | **int32** | Number of users this user is following. | 
**TweetCount** | **int32** | The number of Tweets (including Retweets) posted by this user. | 
**ListedCount** | **int32** | The number of lists that include this user. | 

## Methods

### NewUserPublicMetrics

`func NewUserPublicMetrics(followersCount int32, followingCount int32, tweetCount int32, listedCount int32, ) *UserPublicMetrics`

NewUserPublicMetrics instantiates a new UserPublicMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPublicMetricsWithDefaults

`func NewUserPublicMetricsWithDefaults() *UserPublicMetrics`

NewUserPublicMetricsWithDefaults instantiates a new UserPublicMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFollowersCount

`func (o *UserPublicMetrics) GetFollowersCount() int32`

GetFollowersCount returns the FollowersCount field if non-nil, zero value otherwise.

### GetFollowersCountOk

`func (o *UserPublicMetrics) GetFollowersCountOk() (*int32, bool)`

GetFollowersCountOk returns a tuple with the FollowersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowersCount

`func (o *UserPublicMetrics) SetFollowersCount(v int32)`

SetFollowersCount sets FollowersCount field to given value.


### GetFollowingCount

`func (o *UserPublicMetrics) GetFollowingCount() int32`

GetFollowingCount returns the FollowingCount field if non-nil, zero value otherwise.

### GetFollowingCountOk

`func (o *UserPublicMetrics) GetFollowingCountOk() (*int32, bool)`

GetFollowingCountOk returns a tuple with the FollowingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingCount

`func (o *UserPublicMetrics) SetFollowingCount(v int32)`

SetFollowingCount sets FollowingCount field to given value.


### GetTweetCount

`func (o *UserPublicMetrics) GetTweetCount() int32`

GetTweetCount returns the TweetCount field if non-nil, zero value otherwise.

### GetTweetCountOk

`func (o *UserPublicMetrics) GetTweetCountOk() (*int32, bool)`

GetTweetCountOk returns a tuple with the TweetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTweetCount

`func (o *UserPublicMetrics) SetTweetCount(v int32)`

SetTweetCount sets TweetCount field to given value.


### GetListedCount

`func (o *UserPublicMetrics) GetListedCount() int32`

GetListedCount returns the ListedCount field if non-nil, zero value otherwise.

### GetListedCountOk

`func (o *UserPublicMetrics) GetListedCountOk() (*int32, bool)`

GetListedCountOk returns a tuple with the ListedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListedCount

`func (o *UserPublicMetrics) SetListedCount(v int32)`

SetListedCount sets ListedCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


