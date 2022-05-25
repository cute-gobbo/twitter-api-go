# \UsersApi

All URIs are relative to *https://api.twitter.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindUserById**](UsersApi.md#FindUserById) | **Get** /labs/2/users/{id} | Return details for the specified users
[**FindUserByUsername**](UsersApi.md#FindUserByUsername) | **Get** /labs/2/users/by/username/{username} | Return details for the specified users
[**FindUsersById**](UsersApi.md#FindUsersById) | **Get** /labs/2/users | Return details for the specified users
[**FindUsersByUsername**](UsersApi.md#FindUsersByUsername) | **Get** /labs/2/users/by | Return details for the specified users



## FindUserById

> SingleUserLookupResponse FindUserById(ctx, id).Expansions(expansions).TweetFields(tweetFields).UserFields(userFields).MediaFields(mediaFields).PlaceFields(placeFields).PollFields(pollFields).Execute()

Return details for the specified users



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | Required. A User ID.
    expansions := []string{"Expansions_example"} // []string | A comma separated list of fields to expand. (optional)
    tweetFields := []string{"TweetFields_example"} // []string | A comma separated list of Tweet fields to display. (optional)
    userFields := []string{"UserFields_example"} // []string | A comma separated list of User fields to display. (optional)
    mediaFields := []string{"MediaFields_example"} // []string | A comma separated list of Media fields to display. (optional)
    placeFields := []string{"PlaceFields_example"} // []string | A comma separated list of Place fields to display. (optional)
    pollFields := []string{"PollFields_example"} // []string | A comma separated list of Poll fields to display. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.FindUserById(context.Background(), id).Expansions(expansions).TweetFields(tweetFields).UserFields(userFields).MediaFields(mediaFields).PlaceFields(placeFields).PollFields(pollFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.FindUserById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindUserById`: SingleUserLookupResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.FindUserById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Required. A User ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindUserByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expansions** | **[]string** | A comma separated list of fields to expand. | 
 **tweetFields** | **[]string** | A comma separated list of Tweet fields to display. | 
 **userFields** | **[]string** | A comma separated list of User fields to display. | 
 **mediaFields** | **[]string** | A comma separated list of Media fields to display. | 
 **placeFields** | **[]string** | A comma separated list of Place fields to display. | 
 **pollFields** | **[]string** | A comma separated list of Poll fields to display. | 

### Return type

[**SingleUserLookupResponse**](SingleUserLookupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindUserByUsername

> SingleUserLookupResponse FindUserByUsername(ctx, username).Expansions(expansions).TweetFields(tweetFields).UserFields(userFields).MediaFields(mediaFields).PlaceFields(placeFields).PollFields(pollFields).Execute()

Return details for the specified users



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    username := "username_example" // string | Required. A username.
    expansions := []string{"Expansions_example"} // []string | A comma separated list of fields to expand. (optional)
    tweetFields := []string{"TweetFields_example"} // []string | A comma separated list of Tweet fields to display. (optional)
    userFields := []string{"UserFields_example"} // []string | A comma separated list of User fields to display. (optional)
    mediaFields := []string{"MediaFields_example"} // []string | A comma separated list of Media fields to display. (optional)
    placeFields := []string{"PlaceFields_example"} // []string | A comma separated list of Place fields to display. (optional)
    pollFields := []string{"PollFields_example"} // []string | A comma separated list of Poll fields to display. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.FindUserByUsername(context.Background(), username).Expansions(expansions).TweetFields(tweetFields).UserFields(userFields).MediaFields(mediaFields).PlaceFields(placeFields).PollFields(pollFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.FindUserByUsername``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindUserByUsername`: SingleUserLookupResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.FindUserByUsername`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Required. A username. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindUserByUsernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expansions** | **[]string** | A comma separated list of fields to expand. | 
 **tweetFields** | **[]string** | A comma separated list of Tweet fields to display. | 
 **userFields** | **[]string** | A comma separated list of User fields to display. | 
 **mediaFields** | **[]string** | A comma separated list of Media fields to display. | 
 **placeFields** | **[]string** | A comma separated list of Place fields to display. | 
 **pollFields** | **[]string** | A comma separated list of Poll fields to display. | 

### Return type

[**SingleUserLookupResponse**](SingleUserLookupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindUsersById

> UserLookupResponse FindUsersById(ctx).Ids(ids).Expansions(expansions).TweetFields(tweetFields).UserFields(userFields).MediaFields(mediaFields).PlaceFields(placeFields).PollFields(pollFields).Execute()

Return details for the specified users



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ids := []string{"712089797812"} // []string | Required. A list of User IDs, comma-separated. You can specify up to 100 IDs.
    expansions := []string{"Expansions_example"} // []string | A comma separated list of fields to expand. (optional)
    tweetFields := []string{"TweetFields_example"} // []string | A comma separated list of Tweet fields to display. (optional)
    userFields := []string{"UserFields_example"} // []string | A comma separated list of User fields to display. (optional)
    mediaFields := []string{"MediaFields_example"} // []string | A comma separated list of Media fields to display. (optional)
    placeFields := []string{"PlaceFields_example"} // []string | A comma separated list of Place fields to display. (optional)
    pollFields := []string{"PollFields_example"} // []string | A comma separated list of Poll fields to display. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.FindUsersById(context.Background()).Ids(ids).Expansions(expansions).TweetFields(tweetFields).UserFields(userFields).MediaFields(mediaFields).PlaceFields(placeFields).PollFields(pollFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.FindUsersById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindUsersById`: UserLookupResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.FindUsersById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindUsersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | Required. A list of User IDs, comma-separated. You can specify up to 100 IDs. | 
 **expansions** | **[]string** | A comma separated list of fields to expand. | 
 **tweetFields** | **[]string** | A comma separated list of Tweet fields to display. | 
 **userFields** | **[]string** | A comma separated list of User fields to display. | 
 **mediaFields** | **[]string** | A comma separated list of Media fields to display. | 
 **placeFields** | **[]string** | A comma separated list of Place fields to display. | 
 **pollFields** | **[]string** | A comma separated list of Poll fields to display. | 

### Return type

[**UserLookupResponse**](UserLookupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindUsersByUsername

> UserLookupResponse FindUsersByUsername(ctx).Usernames(usernames).Expansions(expansions).TweetFields(tweetFields).UserFields(userFields).MediaFields(mediaFields).PlaceFields(placeFields).PollFields(pollFields).Execute()

Return details for the specified users



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    usernames := []string{"Inner_example"} // []string | Required . A list of usernames, comma-separated. You can specify up to 100 usernames.
    expansions := []string{"Expansions_example"} // []string | A comma separated list of fields to expand. (optional)
    tweetFields := []string{"TweetFields_example"} // []string | A comma separated list of Tweet fields to display. (optional)
    userFields := []string{"UserFields_example"} // []string | A comma separated list of User fields to display. (optional)
    mediaFields := []string{"MediaFields_example"} // []string | A comma separated list of Media fields to display. (optional)
    placeFields := []string{"PlaceFields_example"} // []string | A comma separated list of Place fields to display. (optional)
    pollFields := []string{"PollFields_example"} // []string | A comma separated list of Poll fields to display. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.FindUsersByUsername(context.Background()).Usernames(usernames).Expansions(expansions).TweetFields(tweetFields).UserFields(userFields).MediaFields(mediaFields).PlaceFields(placeFields).PollFields(pollFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.FindUsersByUsername``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindUsersByUsername`: UserLookupResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.FindUsersByUsername`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindUsersByUsernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usernames** | **[]string** | Required . A list of usernames, comma-separated. You can specify up to 100 usernames. | 
 **expansions** | **[]string** | A comma separated list of fields to expand. | 
 **tweetFields** | **[]string** | A comma separated list of Tweet fields to display. | 
 **userFields** | **[]string** | A comma separated list of User fields to display. | 
 **mediaFields** | **[]string** | A comma separated list of Media fields to display. | 
 **placeFields** | **[]string** | A comma separated list of Place fields to display. | 
 **pollFields** | **[]string** | A comma separated list of Poll fields to display. | 

### Return type

[**UserLookupResponse**](UserLookupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

