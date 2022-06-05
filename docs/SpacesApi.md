# {{classname}}

All URIs are relative to *https://api.twitter.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindSpaceById**](SpacesApi.md#FindSpaceById) | **Get** /2/spaces/{id} | Space lookup by Space ID
[**FindSpacesByCreatorIds**](SpacesApi.md#FindSpacesByCreatorIds) | **Get** /2/spaces/by/creator_ids | Space lookup by their creators
[**FindSpacesByIds**](SpacesApi.md#FindSpacesByIds) | **Get** /2/spaces | Space lookup up Space IDs
[**SearchSpaces**](SpacesApi.md#SearchSpaces) | **Get** /2/spaces/search | Search for Spaces
[**SpaceBuyers**](SpacesApi.md#SpaceBuyers) | **Get** /2/spaces/{id}/buyers | Retrieve the list of users who purchased a ticket to the given space
[**SpaceTweets**](SpacesApi.md#SpaceTweets) | **Get** /2/spaces/{id}/tweets | Retrieve tweets from a Space

# **FindSpaceById**
> SingleSpaceLookupResponse FindSpaceById(ctx, id, optional)
Space lookup by Space ID

Returns a variety of information about the Space specified by the requested ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The space id to be retrieved | 
 **optional** | ***SpacesApiFindSpaceByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpacesApiFindSpaceByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spaceFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Space fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 

### Return type

[**SingleSpaceLookupResponse**](SingleSpaceLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindSpacesByCreatorIds**
> MultiSpaceLookupResponse FindSpacesByCreatorIds(ctx, userIds, optional)
Space lookup by their creators

Returns a variety of information about the Spaces created by the provided User IDs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userIds** | [**[]string**](string.md)| The users to search through | 
 **optional** | ***SpacesApiFindSpacesByCreatorIdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpacesApiFindSpacesByCreatorIdsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spaceFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Space fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 

### Return type

[**MultiSpaceLookupResponse**](MultiSpaceLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindSpacesByIds**
> MultiSpaceLookupResponse FindSpacesByIds(ctx, ids, optional)
Space lookup up Space IDs

Returns a variety of information about the Spaces specified by the requested IDs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ids** | [**[]string**](string.md)| A list of space ids | 
 **optional** | ***SpacesApiFindSpacesByIdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpacesApiFindSpacesByIdsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spaceFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Space fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 

### Return type

[**MultiSpaceLookupResponse**](MultiSpaceLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSpaces**
> MultiSpaceLookupResponse SearchSpaces(ctx, query, optional)
Search for Spaces

Returns Spaces that match the provided query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| The search query | 
 **optional** | ***SpacesApiSearchSpacesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpacesApiSearchSpacesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **optional.String**| The state of spaces to search for | [default to all]
 **maxResults** | [**optional.Interface of int32**](.md)| The number of results to return. The maximum for this value is 100. | 
 **spaceFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Space fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 

### Return type

[**MultiSpaceLookupResponse**](MultiSpaceLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken)

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
 **optional** | ***SpacesApiSpaceBuyersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpacesApiSpaceBuyersOpts struct
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
 **optional** | ***SpacesApiSpaceTweetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpacesApiSpaceTweetsOpts struct
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

