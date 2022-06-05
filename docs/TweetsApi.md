# {{classname}}

All URIs are relative to *https://api.twitter.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrDeleteRules**](TweetsApi.md#AddOrDeleteRules) | **Post** /2/tweets/search/stream/rules | Add/Delete rules
[**CreateTweet**](TweetsApi.md#CreateTweet) | **Post** /2/tweets | Creation of a Tweet
[**DeleteTweetById**](TweetsApi.md#DeleteTweetById) | **Delete** /2/tweets/{id} | Tweet delete by Tweet ID
[**FindTweetById**](TweetsApi.md#FindTweetById) | **Get** /2/tweets/{id} | Tweet lookup by Tweet ID
[**FindTweetsById**](TweetsApi.md#FindTweetsById) | **Get** /2/tweets | Tweet lookup by Tweet IDs
[**FindTweetsThatQuoteATweet**](TweetsApi.md#FindTweetsThatQuoteATweet) | **Get** /2/tweets/{id}/quote_tweets | Retrieve tweets that quote a tweet.
[**GetRules**](TweetsApi.md#GetRules) | **Get** /2/tweets/search/stream/rules | Rules lookup
[**HideReplyById**](TweetsApi.md#HideReplyById) | **Put** /2/tweets/{id}/hidden | Hide replies
[**ListsIdTweets**](TweetsApi.md#ListsIdTweets) | **Get** /2/lists/{id}/tweets | List Tweets timeline by List ID
[**SampleStream**](TweetsApi.md#SampleStream) | **Get** /2/tweets/sample/stream | Sample stream
[**SearchStream**](TweetsApi.md#SearchStream) | **Get** /2/tweets/search/stream | Filtered stream
[**SpaceBuyers**](TweetsApi.md#SpaceBuyers) | **Get** /2/spaces/{id}/buyers | Retrieve the list of users who purchased a ticket to the given space
[**SpaceTweets**](TweetsApi.md#SpaceTweets) | **Get** /2/spaces/{id}/tweets | Retrieve tweets from a Space
[**TweetCountsFullArchiveSearch**](TweetsApi.md#TweetCountsFullArchiveSearch) | **Get** /2/tweets/counts/all | Full archive search counts
[**TweetCountsRecentSearch**](TweetsApi.md#TweetCountsRecentSearch) | **Get** /2/tweets/counts/recent | Recent search counts
[**TweetsFullarchiveSearch**](TweetsApi.md#TweetsFullarchiveSearch) | **Get** /2/tweets/search/all | Full-archive search
[**TweetsRecentSearch**](TweetsApi.md#TweetsRecentSearch) | **Get** /2/tweets/search/recent | Recent search
[**UsersIdLike**](TweetsApi.md#UsersIdLike) | **Post** /2/users/{id}/likes | Causes the user (in the path) to like the specified tweet
[**UsersIdLikedTweets**](TweetsApi.md#UsersIdLikedTweets) | **Get** /2/users/{id}/liked_tweets | Returns Tweet objects liked by the provided User ID
[**UsersIdMentions**](TweetsApi.md#UsersIdMentions) | **Get** /2/users/{id}/mentions | User mention timeline by User ID
[**UsersIdRetweets**](TweetsApi.md#UsersIdRetweets) | **Post** /2/users/{id}/retweets | Causes the user (in the path) to retweet the specified tweet
[**UsersIdTimeline**](TweetsApi.md#UsersIdTimeline) | **Get** /2/users/{id}/timelines/reverse_chronological | User home timeline by User ID
[**UsersIdTweets**](TweetsApi.md#UsersIdTweets) | **Get** /2/users/{id}/tweets | User Tweets timeline by User ID
[**UsersIdUnlike**](TweetsApi.md#UsersIdUnlike) | **Delete** /2/users/{id}/likes/{tweet_id} | Causes the user (in the path) to unlike the specified tweet
[**UsersIdUnretweets**](TweetsApi.md#UsersIdUnretweets) | **Delete** /2/users/{id}/retweets/{source_tweet_id} | Causes the user (in the path) to unretweet the specified tweet

# **AddOrDeleteRules**
> AddOrDeleteRulesResponse AddOrDeleteRules(ctx, body, optional)
Add/Delete rules

Add or delete rules from a user's active rule set. Users can provide unique, optionally tagged rules to add. Users can delete their entire rule set or a subset specified by rule ids or values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddOrDeleteRulesRequest**](AddOrDeleteRulesRequest.md)|  | 
 **optional** | ***TweetsApiAddOrDeleteRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiAddOrDeleteRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dryRun** | **optional.**| Dry Run can be used with both the add and delete action, with the expected result given, but without actually taking any action in the system (meaning the end state will always be as it was when the request was submitted). This is particularly useful to validate rule changes. | 

### Return type

[**AddOrDeleteRulesResponse**](AddOrDeleteRulesResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTweet**
> TweetCreateResponse CreateTweet(ctx, optional)
Creation of a Tweet

Causes the user to create a tweet under the authorized account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TweetsApiCreateTweetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiCreateTweetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Model2TweetsBody**](Model2TweetsBody.md)|  | 

### Return type

[**TweetCreateResponse**](TweetCreateResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTweetById**
> TweetDeleteResponse DeleteTweetById(ctx, id)
Tweet delete by Tweet ID

Delete specified Tweet (in the path) by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the Tweet to be deleted. | 

### Return type

[**TweetDeleteResponse**](TweetDeleteResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindTweetById**
> SingleTweetLookupResponse FindTweetById(ctx, id, optional)
Tweet lookup by Tweet ID

Returns a variety of information about the Tweet specified by the requested ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| A single Tweet ID. | 
 **optional** | ***TweetsApiFindTweetByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiFindTweetByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 

### Return type

[**SingleTweetLookupResponse**](SingleTweetLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindTweetsById**
> MultiTweetLookupResponse FindTweetsById(ctx, ids, optional)
Tweet lookup by Tweet IDs

Returns a variety of information about the Tweet specified by the requested ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ids** | [**[]string**](string.md)| A comma separated list of Tweet IDs. Up to 100 are allowed in a single request. | 
 **optional** | ***TweetsApiFindTweetsByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiFindTweetsByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 

### Return type

[**MultiTweetLookupResponse**](MultiTweetLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindTweetsThatQuoteATweet**
> QuoteTweetLookupResponse FindTweetsThatQuoteATweet(ctx, id, optional)
Retrieve tweets that quote a tweet.

Returns a variety of information about each tweet that quotes the Tweet specified by the requested ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the Quoted Tweet. | 
 **optional** | ***TweetsApiFindTweetsThatQuoteATweetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiFindTweetsThatQuoteATweetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Int32**| The maximum number of results to be returned. | [default to 10]
 **paginationToken** | **optional.String**| This parameter is used to get the next &#x27;page&#x27; of results. | 
 **exclude** | [**optional.Interface of []string**](string.md)| The set of entities to exclude (e.g. &#x27;replies&#x27; or &#x27;retweets&#x27;) | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 

### Return type

[**QuoteTweetLookupResponse**](QuoteTweetLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRules**
> InlineResponse2001 GetRules(ctx, optional)
Rules lookup

Returns rules from a user's active rule set. Users can fetch all of their rules or a subset, specified by the provided rule ids.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TweetsApiGetRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiGetRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**optional.Interface of []string**](string.md)| A comma-separated list of Rule IDs. | 
 **maxResults** | **optional.Int32**| The maximum number of results | [default to 1000]
 **paginationToken** | **optional.String**| This value is populated by passing the &#x27;next_token&#x27; returned in a request to paginate through results. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HideReplyById**
> InlineResponse200 HideReplyById(ctx, id, optional)
Hide replies

Hides or unhides a reply to an owned conversation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the reply that you want to hide or unhide. | 
 **optional** | ***TweetsApiHideReplyByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiHideReplyByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdHiddenBody**](IdHiddenBody.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListsIdTweets**
> InlineResponse2002 ListsIdTweets(ctx, id, optional)
List Tweets timeline by List ID

Returns a list of Tweets associated with the provided List ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the List to list Tweets of | 
 **optional** | ***TweetsApiListsIdTweetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiListsIdTweetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Int32**| The maximum number of results | [default to 100]
 **paginationToken** | **optional.String**| This parameter is used to get the next &#x27;page&#x27; of results. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SampleStream**
> StreamingTweet SampleStream(ctx, optional)
Sample stream

Streams a deterministic 1% of public Tweets.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TweetsApiSampleStreamOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiSampleStreamOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 
 **backfillMinutes** | **optional.Int32**| The number of minutes of backfill requested | 

### Return type

[**StreamingTweet**](StreamingTweet.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchStream**
> FilteredStreamingTweet SearchStream(ctx, optional)
Filtered stream

Streams Tweets matching the stream's active rule set.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TweetsApiSearchStreamOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiSearchStreamOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 
 **backfillMinutes** | **optional.Int32**| The number of minutes of backfill requested | 

### Return type

[**FilteredStreamingTweet**](FilteredStreamingTweet.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SpaceBuyers**
> MultiUserLookupResponse SpaceBuyers(ctx, id, optional)
Retrieve the list of users who purchased a ticket to the given space

Retrieves the list of users who purchased a ticket to the given space

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The space id from which tweets will be retrieved | 
 **optional** | ***TweetsApiSpaceBuyersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiSpaceBuyersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 

### Return type

[**MultiUserLookupResponse**](MultiUserLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SpaceTweets**
> MultiTweetLookupResponse SpaceTweets(ctx, id, optional)
Retrieve tweets from a Space

Retrieves tweets shared in the specified space

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The space id from which tweets will be retrieved | 
 **optional** | ***TweetsApiSpaceTweetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiSpaceTweetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | [**optional.Interface of int32**](.md)| The number of tweets to fetch from the provided space. If not provided, the value will default to the maximum of 100 | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 

### Return type

[**MultiTweetLookupResponse**](MultiTweetLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TweetCountsFullArchiveSearch**
> TweetCountsResponse TweetCountsFullArchiveSearch(ctx, query, optional)
Full archive search counts

Returns Tweet Counts that match a search query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| One query/rule/filter for matching Tweets. Refer to https://t.co/rulelength to identify the max query length | 
 **optional** | ***TweetsApiTweetCountsFullArchiveSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiTweetCountsFullArchiveSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The oldest UTC timestamp (from most recent 7 days) from which the Tweets will be provided. Timestamp is in second granularity and is inclusive (i.e. 12:00:01 includes the first second of the minute). | 
 **endTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The newest, most recent UTC timestamp to which the Tweets will be provided. Timestamp is in second granularity and is exclusive (i.e. 12:00:01 excludes the first second of the minute). | 
 **sinceId** | [**optional.Interface of string**](.md)| Returns results with a Tweet ID greater than (that is, more recent than) the specified ID. | 
 **untilId** | [**optional.Interface of string**](.md)| Returns results with a Tweet ID less than (that is, older than) the specified ID. | 
 **nextToken** | **optional.String**| This parameter is used to get the next &#x27;page&#x27; of results. The value used with the parameter is pulled directly from the response provided by the API, and should not be modified. | 
 **paginationToken** | **optional.String**| This parameter is used to get the next &#x27;page&#x27; of results. The value used with the parameter is pulled directly from the response provided by the API, and should not be modified. | 
 **granularity** | [**optional.Interface of Granularity**](.md)| The granularity for the search counts results. | 

### Return type

[**TweetCountsResponse**](TweetCountsResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TweetCountsRecentSearch**
> TweetCountsResponse TweetCountsRecentSearch(ctx, query, optional)
Recent search counts

Returns Tweet Counts from the last 7 days that match a search query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| One query/rule/filter for matching Tweets. Refer to https://t.co/rulelength to identify the max query length | 
 **optional** | ***TweetsApiTweetCountsRecentSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiTweetCountsRecentSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The oldest UTC timestamp (from most recent 7 days) from which the Tweets will be provided. Timestamp is in second granularity and is inclusive (i.e. 12:00:01 includes the first second of the minute). | 
 **endTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The newest, most recent UTC timestamp to which the Tweets will be provided. Timestamp is in second granularity and is exclusive (i.e. 12:00:01 excludes the first second of the minute). | 
 **sinceId** | [**optional.Interface of string**](.md)| Returns results with a Tweet ID greater than (that is, more recent than) the specified ID. | 
 **untilId** | [**optional.Interface of string**](.md)| Returns results with a Tweet ID less than (that is, older than) the specified ID. | 
 **nextToken** | **optional.String**| This parameter is used to get the next &#x27;page&#x27; of results. The value used with the parameter is pulled directly from the response provided by the API, and should not be modified. | 
 **paginationToken** | **optional.String**| This parameter is used to get the next &#x27;page&#x27; of results. The value used with the parameter is pulled directly from the response provided by the API, and should not be modified. | 
 **granularity** | [**optional.Interface of Granularity**](.md)| The granularity for the search counts results. | 

### Return type

[**TweetCountsResponse**](TweetCountsResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TweetsFullarchiveSearch**
> TweetSearchResponse TweetsFullarchiveSearch(ctx, query, optional)
Full-archive search

Returns Tweets that match a search query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| One query/rule/filter for matching Tweets. Refer to https://t.co/rulelength to identify the max query length | 
 **optional** | ***TweetsApiTweetsFullarchiveSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiTweetsFullarchiveSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The oldest UTC timestamp from which the Tweets will be provided. Timestamp is in second granularity and is inclusive (i.e. 12:00:01 includes the first second of the minute). | 
 **endTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The newest, most recent UTC timestamp to which the Tweets will be provided. Timestamp is in second granularity and is exclusive (i.e. 12:00:01 excludes the first second of the minute). | 
 **sinceId** | [**optional.Interface of string**](.md)| Returns results with a Tweet ID greater than (that is, more recent than) the specified ID. | 
 **untilId** | [**optional.Interface of string**](.md)| Returns results with a Tweet ID less than (that is, older than) the specified ID. | 
 **maxResults** | **optional.Int32**| The maximum number of search results to be returned by a request. | [default to 10]
 **sortOrder** | **optional.String**| This order in which to return results. | 
 **nextToken** | **optional.String**| This parameter is used to get the next &#x27;page&#x27; of results. The value used with the parameter is pulled directly from the response provided by the API, and should not be modified. | 
 **paginationToken** | **optional.String**| This parameter is used to get the next &#x27;page&#x27; of results. The value used with the parameter is pulled directly from the response provided by the API, and should not be modified. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 

### Return type

[**TweetSearchResponse**](TweetSearchResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TweetsRecentSearch**
> TweetSearchResponse TweetsRecentSearch(ctx, query, optional)
Recent search

Returns Tweets from the last 7 days that match a search query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| One query/rule/filter for matching Tweets. Refer to https://t.co/rulelength to identify the max query length | 
 **optional** | ***TweetsApiTweetsRecentSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiTweetsRecentSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The oldest UTC timestamp (from most recent 7 days) from which the Tweets will be provided. Timestamp is in second granularity and is inclusive (i.e. 12:00:01 includes the first second of the minute). | 
 **endTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The newest, most recent UTC timestamp to which the Tweets will be provided. Timestamp is in second granularity and is exclusive (i.e. 12:00:01 excludes the first second of the minute). | 
 **sinceId** | [**optional.Interface of string**](.md)| Returns results with a Tweet ID greater than (that is, more recent than) the specified ID. | 
 **untilId** | [**optional.Interface of string**](.md)| Returns results with a Tweet ID less than (that is, older than) the specified ID. | 
 **maxResults** | **optional.Int32**| The maximum number of search results to be returned by a request. | [default to 10]
 **sortOrder** | **optional.String**| This order in which to return results. | 
 **nextToken** | **optional.String**| This parameter is used to get the next &#x27;page&#x27; of results. The value used with the parameter is pulled directly from the response provided by the API, and should not be modified. | 
 **paginationToken** | **optional.String**| This parameter is used to get the next &#x27;page&#x27; of results. The value used with the parameter is pulled directly from the response provided by the API, and should not be modified. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 

### Return type

[**TweetSearchResponse**](TweetSearchResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdLike**
> UsersLikesCreateResponse UsersIdLike(ctx, id, optional)
Causes the user (in the path) to like the specified tweet

Causes the user (in the path) to like the specified tweet. The user in the path must match the user context authorizing the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the user that is requesting to like the tweet | 
 **optional** | ***TweetsApiUsersIdLikeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiUsersIdLikeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UsersLikesCreateRequest**](UsersLikesCreateRequest.md)|  | 

### Return type

[**UsersLikesCreateResponse**](UsersLikesCreateResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdLikedTweets**
> InlineResponse2002 UsersIdLikedTweets(ctx, id, optional)
Returns Tweet objects liked by the provided User ID

Returns a list of Tweets liked by the provided User ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the User to list the liked Tweets of | 
 **optional** | ***TweetsApiUsersIdLikedTweetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiUsersIdLikedTweetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Int32**| The maximum number of results | 
 **paginationToken** | **optional.String**| This parameter is used to get the next &#x27;page&#x27; of results. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdMentions**
> GenericTweetsTimelineResponse UsersIdMentions(ctx, id, optional)
User mention timeline by User ID

Returns Tweet objects that mention username associated to the provided User ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the User to list mentions of | 
 **optional** | ***TweetsApiUsersIdMentionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiUsersIdMentionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sinceId** | [**optional.Interface of string**](.md)| The minimum Tweet ID to be included in the result set. This parameter takes precedence over start_time if both are specified. | 
 **untilId** | [**optional.Interface of string**](.md)| The maximum Tweet ID to be included in the result set. This parameter takes precedence over end_time if both are specified. | 
 **maxResults** | **optional.Int32**| The maximum number of results | 
 **paginationToken** | **optional.String**| This parameter is used to get the next &#x27;page&#x27; of results. | 
 **startTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The earliest UTC timestamp from which the Tweets will be provided. The since_id parameter takes precedence if it is also specified. | 
 **endTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The latest UTC timestamp to which the Tweets will be provided. The until_id parameter takes precedence if it is also specified. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 

### Return type

[**GenericTweetsTimelineResponse**](GenericTweetsTimelineResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdRetweets**
> UsersRetweetsCreateResponse UsersIdRetweets(ctx, id, optional)
Causes the user (in the path) to retweet the specified tweet

Causes the user (in the path) to retweet the specified tweet. The user in the path must match the user context authorizing the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the user that is requesting to retweet the tweet | 
 **optional** | ***TweetsApiUsersIdRetweetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiUsersIdRetweetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UsersRetweetsCreateRequest**](UsersRetweetsCreateRequest.md)|  | 

### Return type

[**UsersRetweetsCreateResponse**](UsersRetweetsCreateResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdTimeline**
> GenericTweetsTimelineResponse UsersIdTimeline(ctx, id, optional)
User home timeline by User ID

Returns Tweet objects that appears in the provided User ID's home timeline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the User to list Reverse Chronological Timeline Tweets of | 
 **optional** | ***TweetsApiUsersIdTimelineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiUsersIdTimelineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sinceId** | [**optional.Interface of string**](.md)| The minimum Tweet ID to be included in the result set. This parameter takes precedence over start_time if both are specified. | 
 **untilId** | [**optional.Interface of string**](.md)| The maximum Tweet ID to be included in the result set. This parameter takes precedence over end_time if both are specified. | 
 **maxResults** | **optional.Int32**| The maximum number of results | 
 **exclude** | [**optional.Interface of []string**](string.md)| The set of entities to exclude (e.g. &#x27;replies&#x27; or &#x27;retweets&#x27;) | 
 **paginationToken** | **optional.String**| This parameter is used to get the next &#x27;page&#x27; of results. | 
 **startTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The earliest UTC timestamp from which the Tweets will be provided. The since_id parameter takes precedence if it is also specified. | 
 **endTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The latest UTC timestamp to which the Tweets will be provided. The until_id parameter takes precedence if it is also specified. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 

### Return type

[**GenericTweetsTimelineResponse**](GenericTweetsTimelineResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdTweets**
> GenericTweetsTimelineResponse UsersIdTweets(ctx, id, optional)
User Tweets timeline by User ID

Returns a list of Tweets authored by the provided User ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the User to list Tweets of | 
 **optional** | ***TweetsApiUsersIdTweetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TweetsApiUsersIdTweetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sinceId** | [**optional.Interface of string**](.md)| The minimum Tweet ID to be included in the result set. This parameter takes precedence over start_time if both are specified. | 
 **untilId** | [**optional.Interface of string**](.md)| The maximum Tweet ID to be included in the result set. This parameter takes precedence over end_time if both are specified. | 
 **maxResults** | **optional.Int32**| The maximum number of results | 
 **exclude** | [**optional.Interface of []string**](string.md)| The set of entities to exclude (e.g. &#x27;replies&#x27; or &#x27;retweets&#x27;) | 
 **paginationToken** | **optional.String**| This parameter is used to get the next &#x27;page&#x27; of results. | 
 **startTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The earliest UTC timestamp from which the Tweets will be provided. The since_id parameter takes precedence if it is also specified. | 
 **endTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The latest UTC timestamp to which the Tweets will be provided. The until_id parameter takes precedence if it is also specified. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 

### Return type

[**GenericTweetsTimelineResponse**](GenericTweetsTimelineResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdUnlike**
> UsersLikesDeleteResponse UsersIdUnlike(ctx, id, tweetId)
Causes the user (in the path) to unlike the specified tweet

Causes the user (in the path) to unlike the specified tweet. The user must match the user context authorizing the request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the user that is requesting to unlike the tweet | 
  **tweetId** | [**string**](.md)| The ID of the tweet that the user is requesting to unlike | 

### Return type

[**UsersLikesDeleteResponse**](UsersLikesDeleteResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdUnretweets**
> UsersRetweetsDeleteResponse UsersIdUnretweets(ctx, id, sourceTweetId)
Causes the user (in the path) to unretweet the specified tweet

Causes the user (in the path) to unretweet the specified tweet. The user must match the user context authorizing the request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the user that is requesting to unretweet the tweet | 
  **sourceTweetId** | [**string**](.md)| The ID of the tweet that the user is requesting to unretweet | 

### Return type

[**UsersRetweetsDeleteResponse**](UsersRetweetsDeleteResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

