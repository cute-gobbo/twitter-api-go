# {{classname}}

All URIs are relative to *https://api.twitter.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserListMemberships**](ListsApi.md#GetUserListMemberships) | **Get** /2/users/{id}/list_memberships | Get a User&#x27;s List Memberships
[**ListAddMember**](ListsApi.md#ListAddMember) | **Post** /2/lists/{id}/members | Add a List member
[**ListIdCreate**](ListsApi.md#ListIdCreate) | **Post** /2/lists | Create List
[**ListIdDelete**](ListsApi.md#ListIdDelete) | **Delete** /2/lists/{id} | Delete List
[**ListIdGet**](ListsApi.md#ListIdGet) | **Get** /2/lists/{id} | List lookup by List ID
[**ListIdUpdate**](ListsApi.md#ListIdUpdate) | **Put** /2/lists/{id} | Update List
[**ListRemoveMember**](ListsApi.md#ListRemoveMember) | **Delete** /2/lists/{id}/members/{user_id} | Remove a List member
[**ListUserFollow**](ListsApi.md#ListUserFollow) | **Post** /2/users/{id}/followed_lists | Follow a List
[**ListUserOwnedLists**](ListsApi.md#ListUserOwnedLists) | **Get** /2/users/{id}/owned_lists | Get a User&#x27;s Owned Lists
[**ListUserPin**](ListsApi.md#ListUserPin) | **Post** /2/users/{id}/pinned_lists | Pin a List
[**ListUserPinnedLists**](ListsApi.md#ListUserPinnedLists) | **Get** /2/users/{id}/pinned_lists | Get a User&#x27;s Pinned Lists
[**ListUserUnfollow**](ListsApi.md#ListUserUnfollow) | **Delete** /2/users/{id}/followed_lists/{list_id} | Unfollow a List
[**ListUserUnpin**](ListsApi.md#ListUserUnpin) | **Delete** /2/users/{id}/pinned_lists/{list_id} | Unpin a List
[**UserFollowedLists**](ListsApi.md#UserFollowedLists) | **Get** /2/users/{id}/followed_lists | Get User&#x27;s Followed Lists

# **GetUserListMemberships**
> MultiListResponse GetUserListMemberships(ctx, id, optional)
Get a User's List Memberships

Get a User's List Memberships.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the user for whom to return results | 
 **optional** | ***ListsApiGetUserListMembershipsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetUserListMembershipsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Int32**| The maximum number of results | [default to 100]
 **paginationToken** | **optional.Int64**| This parameter is used to get a specified &#x27;page&#x27; of results. | 
 **listFields** | [**optional.Interface of []string**](string.md)| A comma separated list of List fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 

### Return type

[**MultiListResponse**](MultiListResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAddMember**
> ListMemberResponse ListAddMember(ctx, id, optional)
Add a List member

Causes a user to become a member of a List.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the List to add a member | 
 **optional** | ***ListsApiListAddMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiListAddMemberOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ListAddMemberRequest**](ListAddMemberRequest.md)|  | 

### Return type

[**ListMemberResponse**](ListMemberResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIdCreate**
> ListCreateResponse ListIdCreate(ctx, optional)
Create List

Creates a new List.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListsApiListIdCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiListIdCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ListCreateRequest**](ListCreateRequest.md)|  | 

### Return type

[**ListCreateResponse**](ListCreateResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIdDelete**
> ListDeleteResponse ListIdDelete(ctx, id)
Delete List

Delete a List that you own.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the List to delete | 

### Return type

[**ListDeleteResponse**](ListDeleteResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIdGet**
> SingleListLookupResponse ListIdGet(ctx, id, optional)
List lookup by List ID

Returns a List

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the List to get | 
 **optional** | ***ListsApiListIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiListIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **listFields** | [**optional.Interface of []string**](string.md)| A comma separated list of List fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 

### Return type

[**SingleListLookupResponse**](SingleListLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIdUpdate**
> ListUpdateResponse ListIdUpdate(ctx, id, optional)
Update List

Update a List that you own.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the List to modify | 
 **optional** | ***ListsApiListIdUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiListIdUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ListUpdateRequest**](ListUpdateRequest.md)|  | 

### Return type

[**ListUpdateResponse**](ListUpdateResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRemoveMember**
> ListMemberResponse ListRemoveMember(ctx, id, userId)
Remove a List member

Causes a user to be removed from the members of a List.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the List to remove a member | 
  **userId** | [**string**](.md)| The ID of user that will be removed from the List | 

### Return type

[**ListMemberResponse**](ListMemberResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUserFollow**
> ListFollowedResponse ListUserFollow(ctx, id, optional)
Follow a List

Causes a user to follow a List.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the authenticated source user that will follow the List | 
 **optional** | ***ListsApiListUserFollowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiListUserFollowOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ListFollowRequest**](ListFollowRequest.md)|  | 

### Return type

[**ListFollowedResponse**](ListFollowedResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUserOwnedLists**
> MultiListResponse ListUserOwnedLists(ctx, id, optional)
Get a User's Owned Lists

Get a User's Owned Lists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the user for whom to return results | 
 **optional** | ***ListsApiListUserOwnedListsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiListUserOwnedListsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Int32**| The maximum number of results | [default to 100]
 **paginationToken** | **optional.Int64**| This parameter is used to get a specified &#x27;page&#x27; of results. | 
 **listFields** | [**optional.Interface of []string**](string.md)| A comma separated list of List fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 

### Return type

[**MultiListResponse**](MultiListResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUserPin**
> ListPinnedResponse ListUserPin(ctx, id, optional)
Pin a List

Causes a user to pin a List.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the authenticated source user that will pin the List | 
 **optional** | ***ListsApiListUserPinOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiListUserPinOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ListPinRequest**](ListPinRequest.md)|  | 

### Return type

[**ListPinnedResponse**](ListPinnedResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUserPinnedLists**
> MultiListNoPaginationResponse ListUserPinnedLists(ctx, id, optional)
Get a User's Pinned Lists

Get a User's Pinned Lists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the user for whom to return results | 
 **optional** | ***ListsApiListUserPinnedListsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiListUserPinnedListsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **listFields** | [**optional.Interface of []string**](string.md)| A comma separated list of List fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 

### Return type

[**MultiListNoPaginationResponse**](MultiListNoPaginationResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUserUnfollow**
> ListFollowedResponse ListUserUnfollow(ctx, id, listId)
Unfollow a List

Causes a user to unfollow a List.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the authenticated source user that will unfollow the List | 
  **listId** | [**string**](.md)| The ID of the List to unfollow | 

### Return type

[**ListFollowedResponse**](ListFollowedResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUserUnpin**
> ListPinnedResponse ListUserUnpin(ctx, id, listId)
Unpin a List

Causes a user to remove a pinned List.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the authenticated source user that will remove the pinned List | 
  **listId** | [**string**](.md)| The ID of the List to unpin | 

### Return type

[**ListPinnedResponse**](ListPinnedResponse.md)

### Authorization

[OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserFollowedLists**
> MultiListResponse UserFollowedLists(ctx, id, optional)
Get User's Followed Lists

Returns a user's followed Lists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the user for whom to return results | 
 **optional** | ***ListsApiUserFollowedListsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiUserFollowedListsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Int32**| The maximum number of results | [default to 100]
 **paginationToken** | **optional.Int64**| This parameter is used to get a specified &#x27;page&#x27; of results. | 
 **listFields** | [**optional.Interface of []string**](string.md)| A comma separated list of List fields to display. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 

### Return type

[**MultiListResponse**](MultiListResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [OAuth2UserToken](../README.md#OAuth2UserToken), [UserToken](../README.md#UserToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

