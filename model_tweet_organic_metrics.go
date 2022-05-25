/*
Tweets and Users

API Reference — Labs v2

API version: 2.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package twitterapi

import (
	"encoding/json"
)

// TweetOrganicMetrics Organic nonpublic engagement metrics for the Tweet at the time of the request.
type TweetOrganicMetrics struct {
	// Number of times this Tweet has been viewed.
	ImpressionCount int32 `json:"impression_count"`
	// Number of times this Tweet has been Retweeted.
	RetweetCount int32 `json:"retweet_count"`
	// Number of times this Tweet has been replied to.
	ReplyCount int32 `json:"reply_count"`
	// Number of times this Tweet has been liked.
	LikeCount int32 `json:"like_count"`
	// Number of times the user's profile from this Tweet has been clicked.
	UserProfileClicks int32 `json:"user_profile_clicks"`
	// Number of times links in this Tweet have been clicked.
	UrlLinkClicks *int32 `json:"url_link_clicks,omitempty"`
}

// NewTweetOrganicMetrics instantiates a new TweetOrganicMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTweetOrganicMetrics(impressionCount int32, retweetCount int32, replyCount int32, likeCount int32, userProfileClicks int32) *TweetOrganicMetrics {
	this := TweetOrganicMetrics{}
	this.ImpressionCount = impressionCount
	this.RetweetCount = retweetCount
	this.ReplyCount = replyCount
	this.LikeCount = likeCount
	this.UserProfileClicks = userProfileClicks
	return &this
}

// NewTweetOrganicMetricsWithDefaults instantiates a new TweetOrganicMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTweetOrganicMetricsWithDefaults() *TweetOrganicMetrics {
	this := TweetOrganicMetrics{}
	return &this
}

// GetImpressionCount returns the ImpressionCount field value
func (o *TweetOrganicMetrics) GetImpressionCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ImpressionCount
}

// GetImpressionCountOk returns a tuple with the ImpressionCount field value
// and a boolean to check if the value has been set.
func (o *TweetOrganicMetrics) GetImpressionCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ImpressionCount, true
}

// SetImpressionCount sets field value
func (o *TweetOrganicMetrics) SetImpressionCount(v int32) {
	o.ImpressionCount = v
}

// GetRetweetCount returns the RetweetCount field value
func (o *TweetOrganicMetrics) GetRetweetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RetweetCount
}

// GetRetweetCountOk returns a tuple with the RetweetCount field value
// and a boolean to check if the value has been set.
func (o *TweetOrganicMetrics) GetRetweetCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RetweetCount, true
}

// SetRetweetCount sets field value
func (o *TweetOrganicMetrics) SetRetweetCount(v int32) {
	o.RetweetCount = v
}

// GetReplyCount returns the ReplyCount field value
func (o *TweetOrganicMetrics) GetReplyCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReplyCount
}

// GetReplyCountOk returns a tuple with the ReplyCount field value
// and a boolean to check if the value has been set.
func (o *TweetOrganicMetrics) GetReplyCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReplyCount, true
}

// SetReplyCount sets field value
func (o *TweetOrganicMetrics) SetReplyCount(v int32) {
	o.ReplyCount = v
}

// GetLikeCount returns the LikeCount field value
func (o *TweetOrganicMetrics) GetLikeCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LikeCount
}

// GetLikeCountOk returns a tuple with the LikeCount field value
// and a boolean to check if the value has been set.
func (o *TweetOrganicMetrics) GetLikeCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LikeCount, true
}

// SetLikeCount sets field value
func (o *TweetOrganicMetrics) SetLikeCount(v int32) {
	o.LikeCount = v
}

// GetUserProfileClicks returns the UserProfileClicks field value
func (o *TweetOrganicMetrics) GetUserProfileClicks() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserProfileClicks
}

// GetUserProfileClicksOk returns a tuple with the UserProfileClicks field value
// and a boolean to check if the value has been set.
func (o *TweetOrganicMetrics) GetUserProfileClicksOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserProfileClicks, true
}

// SetUserProfileClicks sets field value
func (o *TweetOrganicMetrics) SetUserProfileClicks(v int32) {
	o.UserProfileClicks = v
}

// GetUrlLinkClicks returns the UrlLinkClicks field value if set, zero value otherwise.
func (o *TweetOrganicMetrics) GetUrlLinkClicks() int32 {
	if o == nil || o.UrlLinkClicks == nil {
		var ret int32
		return ret
	}
	return *o.UrlLinkClicks
}

// GetUrlLinkClicksOk returns a tuple with the UrlLinkClicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetOrganicMetrics) GetUrlLinkClicksOk() (*int32, bool) {
	if o == nil || o.UrlLinkClicks == nil {
		return nil, false
	}
	return o.UrlLinkClicks, true
}

// HasUrlLinkClicks returns a boolean if a field has been set.
func (o *TweetOrganicMetrics) HasUrlLinkClicks() bool {
	if o != nil && o.UrlLinkClicks != nil {
		return true
	}

	return false
}

// SetUrlLinkClicks gets a reference to the given int32 and assigns it to the UrlLinkClicks field.
func (o *TweetOrganicMetrics) SetUrlLinkClicks(v int32) {
	o.UrlLinkClicks = &v
}

func (o TweetOrganicMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["impression_count"] = o.ImpressionCount
	}
	if true {
		toSerialize["retweet_count"] = o.RetweetCount
	}
	if true {
		toSerialize["reply_count"] = o.ReplyCount
	}
	if true {
		toSerialize["like_count"] = o.LikeCount
	}
	if true {
		toSerialize["user_profile_clicks"] = o.UserProfileClicks
	}
	if o.UrlLinkClicks != nil {
		toSerialize["url_link_clicks"] = o.UrlLinkClicks
	}
	return json.Marshal(toSerialize)
}

type NullableTweetOrganicMetrics struct {
	value *TweetOrganicMetrics
	isSet bool
}

func (v NullableTweetOrganicMetrics) Get() *TweetOrganicMetrics {
	return v.value
}

func (v *NullableTweetOrganicMetrics) Set(val *TweetOrganicMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableTweetOrganicMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableTweetOrganicMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTweetOrganicMetrics(val *TweetOrganicMetrics) *NullableTweetOrganicMetrics {
	return &NullableTweetOrganicMetrics{value: val, isSet: true}
}

func (v NullableTweetOrganicMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTweetOrganicMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


