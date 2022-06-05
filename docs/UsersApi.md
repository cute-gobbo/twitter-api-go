# {{classname}}

All URIs are relative to *https://api.twitter.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindMyUser**](UsersApi.md#FindMyUser) | **Get** /2/users/me | User lookup me
[**FindUserById**](UsersApi.md#FindUserById) | **Get** /2/users/{id} | User lookup by ID
[**FindUserByUsername**](UsersApi.md#FindUserByUsername) | **Get** /2/users/by/username/{username} | User lookup by username
[**FindUsersById**](UsersApi.md#FindUsersById) | **Get** /2/users | User lookup by IDs
[**FindUsersByUsername**](UsersApi.md#FindUsersByUsername) | **Get** /2/users/by | User lookup by usernames
[**ListGetFollowers**](UsersApi.md#ListGetFollowers) | **Get** /2/lists/{id}/followers | Returns user objects that follow a List by the provided List ID
[**ListGetMembers**](UsersApi.md#ListGetMembers) | **Get** /2/lists/{id}/members | Returns user objects that are members of a List by the provided List ID
[**TweetsIdLikingUsers**](UsersApi.md#TweetsIdLikingUsers) | **Get** /2/tweets/{id}/liking_users | Returns user objects that have liked the provided Tweet ID
[**TweetsIdRetweetingUsers**](UsersApi.md#TweetsIdRetweetingUsers) | **Get** /2/tweets/{id}/retweeted_by | Returns user objects that have retweeted the provided Tweet ID
[**UsersIdBlock**](UsersApi.md#UsersIdBlock) | **Post** /2/users/{id}/blocking | Block User by User ID
[**UsersIdBlocking**](UsersApi.md#UsersIdBlocking) | **Get** /2/users/{id}/blocking | Returns user objects that are blocked by provided user ID
[**UsersIdFollow**](UsersApi.md#UsersIdFollow) | **Post** /2/users/{id}/following | Follow User
[**UsersIdFollowers**](UsersApi.md#UsersIdFollowers) | **Get** /2/users/{id}/followers | Returns user objects that follow the provided user ID
[**UsersIdFollowing**](UsersApi.md#UsersIdFollowing) | **Get** /2/users/{id}/following | Following by User ID
[**UsersIdMute**](UsersApi.md#UsersIdMute) | **Post** /2/users/{id}/muting | Mute User by User ID
[**UsersIdMuting**](UsersApi.md#UsersIdMuting) | **Get** /2/users/{id}/muting | Returns user objects that are muted by the provided user ID
[**UsersIdUnblock**](UsersApi.md#UsersIdUnblock) | **Delete** /2/users/{source_user_id}/blocking/{target_user_id} | Unblock User by User ID
[**UsersIdUnfollow**](UsersApi.md#UsersIdUnfollow) | **Delete** /2/users/{source_user_id}/following/{target_user_id} | Unfollow User
[**UsersIdUnmute**](UsersApi.md#UsersIdUnmute) | **Delete** /2/users/{source_user_id}/muting/{target_user_id} | Unmute User by User ID

# **FindMyUser**
> SingleUserLookupResponse FindMyUser(ctx, optional)
User lookup me

This endpoint returns information about the requesting user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiFindMyUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiFindMyUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 

### Return type

[**SingleUserLookupResponse**](SingleUserLookupResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindUserById**
> SingleUserLookupResponse FindUserById(ctx, id, optional)
User lookup by ID

This endpoint returns information about a user. Specify user by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Required. A User ID. | 
 **optional** | ***UsersApiFindUserByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiFindUserByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 

### Return type

[**SingleUserLookupResponse**](SingleUserLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindUserByUsername**
> SingleUserLookupResponse FindUserByUsername(ctx, username, optional)
User lookup by username

This endpoint returns information about a user. Specify user by username.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | [**string**](.md)| Required. A username. | 
 **optional** | ***UsersApiFindUserByUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiFindUserByUsernameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 

### Return type

[**SingleUserLookupResponse**](SingleUserLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindUsersById**
> MultiUserLookupResponse FindUsersById(ctx, ids, optional)
User lookup by IDs

This endpoint returns information about users. Specify users by their ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ids** | [**[]string**](string.md)| Required. A list of User IDs, comma-separated. You can specify up to 100 IDs. | 
 **optional** | ***UsersApiFindUsersByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiFindUsersByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 

### Return type

[**MultiUserLookupResponse**](MultiUserLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindUsersByUsername**
> MultiUserLookupResponse FindUsersByUsername(ctx, usernames, optional)
User lookup by usernames

This endpoint returns information about users. Specify users by their username.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **usernames** | [**[]string**](string.md)| Required . A list of usernames, comma-separated. You can specify up to 100 usernames. | 
 **optional** | ***UsersApiFindUsersByUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiFindUsersByUsernameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 

### Return type

[**MultiUserLookupResponse**](MultiUserLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGetFollowers**
> ListLookupMultipleUsersLookupResponse ListGetFollowers(ctx, id, optional)
Returns user objects that follow a List by the provided List ID

Returns a list of users that follow a List by the provided List ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the List for which to return followers | 
 **optional** | ***UsersApiListGetFollowersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiListGetFollowersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Int32**| The maximum number of results | [default to 100]
 **paginationToken** | **optional.Int64**| This parameter is used to get a specified &#x27;page&#x27; of results. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 

### Return type

[**ListLookupMultipleUsersLookupResponse**](ListLookupMultipleUsersLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGetMembers**
> ListLookupMultipleUsersLookupResponse ListGetMembers(ctx, id, optional)
Returns user objects that are members of a List by the provided List ID

Returns a list of users that are members of a List by the provided List ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the List for which to return members | 
 **optional** | ***UsersApiListGetMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiListGetMembersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Int32**| The maximum number of results | [default to 100]
 **paginationToken** | **optional.Int64**| This parameter is used to get a specified &#x27;page&#x27; of results. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 

### Return type

[**ListLookupMultipleUsersLookupResponse**](ListLookupMultipleUsersLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TweetsIdLikingUsers**
> GenericMultipleUsersLookupResponse TweetsIdLikingUsers(ctx, id, optional)
Returns user objects that have liked the provided Tweet ID

Returns a list of users that have liked the provided Tweet ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the Tweet for which to return results | 
 **optional** | ***UsersApiTweetsIdLikingUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiTweetsIdLikingUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Int32**| The maximum number of results | [default to 100]
 **paginationToken** | **optional.String**| This parameter is used to get the next &#x27;page&#x27; of results. | 

### Return type

[**GenericMultipleUsersLookupResponse**](GenericMultipleUsersLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TweetsIdRetweetingUsers**
> GenericMultipleUsersLookupResponse TweetsIdRetweetingUsers(ctx, id, optional)
Returns user objects that have retweeted the provided Tweet ID

Returns a list of users that have retweeted the provided Tweet ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the Tweet for which to return results | 
 **optional** | ***UsersApiTweetsIdRetweetingUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiTweetsIdRetweetingUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Int32**| The maximum number of results | [default to 100]
 **paginationToken** | **optional.String**| This parameter is used to get the next &#x27;page&#x27; of results. | 

### Return type

[**GenericMultipleUsersLookupResponse**](GenericMultipleUsersLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdBlock**
> UsersBlockingMutationResponse UsersIdBlock(ctx, id, optional)
Block User by User ID

Causes the user (in the path) to block the target user. The user (in the path) must match the user context authorizing the request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the user that is requesting to block the target user | 
 **optional** | ***UsersApiUsersIdBlockOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersIdBlockOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdBlockingBody**](IdBlockingBody.md)|  | 

### Return type

[**UsersBlockingMutationResponse**](UsersBlockingMutationResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdBlocking**
> GenericMultipleUsersLookupResponse UsersIdBlocking(ctx, id, optional)
Returns user objects that are blocked by provided user ID

Returns a list of users that are blocked by the provided user ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the user for whom to return results | 
 **optional** | ***UsersApiUsersIdBlockingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersIdBlockingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Int32**| The maximum number of results | 
 **paginationToken** | **optional.String**| This value is populated by passing the &#x27;next_token&#x27; or &#x27;previous_token&#x27; returned in a request to paginate through results. | 

### Return type

[**GenericMultipleUsersLookupResponse**](GenericMultipleUsersLookupResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdFollow**
> UsersFollowingCreateResponse UsersIdFollow(ctx, id, optional)
Follow User

Causes the user(in the path) to follow, or “request to follow” for protected users, the target user. The user(in the path) must match the user context authorizing the request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the user that is requesting to follow the target user | 
 **optional** | ***UsersApiUsersIdFollowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersIdFollowOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdFollowingBody**](IdFollowingBody.md)|  | 

### Return type

[**UsersFollowingCreateResponse**](UsersFollowingCreateResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdFollowers**
> GenericMultipleUsersLookupResponse UsersIdFollowers(ctx, id, optional)
Returns user objects that follow the provided user ID

Returns a list of users that follow the provided user ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the user for whom to return results | 
 **optional** | ***UsersApiUsersIdFollowersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersIdFollowersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Int32**| The maximum number of results | 
 **paginationToken** | **optional.String**| This value is populated by passing the &#x27;next_token&#x27; or &#x27;previous_token&#x27; returned in a request to paginate through results. | 

### Return type

[**GenericMultipleUsersLookupResponse**](GenericMultipleUsersLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdFollowing**
> UsersFollowingLookupResponse UsersIdFollowing(ctx, id, optional)
Following by User ID

Returns a list of users that are being followed by the provided user ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the user for whom to return results | 
 **optional** | ***UsersApiUsersIdFollowingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersIdFollowingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Int32**| The maximum number of results | 
 **paginationToken** | **optional.String**| This value is populated by passing the &#x27;next_token&#x27; or &#x27;previous_token&#x27; returned in a request to paginate through results. | 

### Return type

[**UsersFollowingLookupResponse**](UsersFollowingLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdMute**
> UsersMutingMutationResponse UsersIdMute(ctx, id, optional)
Mute User by User ID

Causes the user (in the path) to mute the target user. The user (in the path) must match the user context authorizing the request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the user that is requesting to mute the target user | 
 **optional** | ***UsersApiUsersIdMuteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersIdMuteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdMutingBody**](IdMutingBody.md)|  | 

### Return type

[**UsersMutingMutationResponse**](UsersMutingMutationResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdMuting**
> GenericMultipleUsersLookupResponse UsersIdMuting(ctx, id, optional)
Returns user objects that are muted by the provided user ID

Returns a list of users that are muted by the provided user ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the user for whom to return results | 
 **optional** | ***UsersApiUsersIdMutingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersIdMutingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Int32**| The maximum number of results | [default to 100]
 **paginationToken** | **optional.String**| This parameter is used to get the next &#x27;page&#x27; of results. | 

### Return type

[**GenericMultipleUsersLookupResponse**](GenericMultipleUsersLookupResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdUnblock**
> UsersBlockingMutationResponse UsersIdUnblock(ctx, sourceUserId, targetUserId)
Unblock User by User ID

Causes the source user to unblock the target user. The source user must match the user context authorizing the request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sourceUserId** | [**string**](.md)| The ID of the user that is requesting to unblock the target user | 
  **targetUserId** | [**string**](.md)| The ID of the user that the source user is requesting to unblock | 

### Return type

[**UsersBlockingMutationResponse**](UsersBlockingMutationResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdUnfollow**
> UsersFollowingDeleteResponse UsersIdUnfollow(ctx, sourceUserId, targetUserId)
Unfollow User

Causes the source user to unfollow the target user. The source user must match the user context authorizing the request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sourceUserId** | [**string**](.md)| The ID of the user that is requesting to unfollow the target user | 
  **targetUserId** | [**string**](.md)| The ID of the user that the source user is requesting to unfollow | 

### Return type

[**UsersFollowingDeleteResponse**](UsersFollowingDeleteResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdUnmute**
> UsersMutingMutationResponse UsersIdUnmute(ctx, sourceUserId, targetUserId)
Unmute User by User ID

Causes the source user to unmute the target user. The source user must match the user context authorizing the request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sourceUserId** | [**string**](.md)| The ID of the user that is requesting to unmute the target user | 
  **targetUserId** | [**string**](.md)| The ID of the user that the source user is requesting to unmute | 

### Return type

[**UsersMutingMutationResponse**](UsersMutingMutationResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

