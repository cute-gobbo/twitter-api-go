# \TweetsApi

All URIs are relative to *https://api.twitter.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindTweetById**](TweetsApi.md#FindTweetById) | **Get** /labs/2/tweets/{id} | Returns hydrated Tweet objects
[**FindTweetsById**](TweetsApi.md#FindTweetsById) | **Get** /labs/2/tweets | Returns hydrated Tweet objects



## FindTweetById

> SingleTweetLookupResponse FindTweetById(ctx, id).Expansions(expansions).TweetFields(tweetFields).UserFields(userFields).MediaFields(mediaFields).PlaceFields(placeFields).PollFields(pollFields).Execute()

Returns hydrated Tweet objects



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
    id := "id_example" // string | A single Tweet ID.
    expansions := []string{"Expansions_example"} // []string | A comma separated list of fields to expand. (optional)
    tweetFields := []string{"TweetFields_example"} // []string | A comma separated list of Tweet fields to display. (optional)
    userFields := []string{"UserFields_example"} // []string | A comma separated list of User fields to display. (optional)
    mediaFields := []string{"MediaFields_example"} // []string | A comma separated list of Media fields to display. (optional)
    placeFields := []string{"PlaceFields_example"} // []string | A comma separated list of Place fields to display. (optional)
    pollFields := []string{"PollFields_example"} // []string | A comma separated list of Poll fields to display. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TweetsApi.FindTweetById(context.Background(), id).Expansions(expansions).TweetFields(tweetFields).UserFields(userFields).MediaFields(mediaFields).PlaceFields(placeFields).PollFields(pollFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TweetsApi.FindTweetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindTweetById`: SingleTweetLookupResponse
    fmt.Fprintf(os.Stdout, "Response from `TweetsApi.FindTweetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A single Tweet ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindTweetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expansions** | **[]string** | A comma separated list of fields to expand. | 
 **tweetFields** | **[]string** | A comma separated list of Tweet fields to display. | 
 **userFields** | **[]string** | A comma separated list of User fields to display. | 
 **mediaFields** | **[]string** | A comma separated list of Media fields to display. | 
 **placeFields** | **[]string** | A comma separated list of Place fields to display. | 
 **pollFields** | **[]string** | A comma separated list of Poll fields to display. | 

### Return type

[**SingleTweetLookupResponse**](SingleTweetLookupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindTweetsById

> TweetLookupResponse FindTweetsById(ctx).Ids(ids).Expansions(expansions).TweetFields(tweetFields).UserFields(userFields).MediaFields(mediaFields).PlaceFields(placeFields).PollFields(pollFields).Execute()

Returns hydrated Tweet objects



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
    ids := []string{"120897978112909812"} // []string | A comma separated list of Tweet IDs. Up to 100 are allowed in a single request.
    expansions := []string{"Expansions_example"} // []string | A comma separated list of fields to expand. (optional)
    tweetFields := []string{"TweetFields_example"} // []string | A comma separated list of Tweet fields to display. (optional)
    userFields := []string{"UserFields_example"} // []string | A comma separated list of User fields to display. (optional)
    mediaFields := []string{"MediaFields_example"} // []string | A comma separated list of Media fields to display. (optional)
    placeFields := []string{"PlaceFields_example"} // []string | A comma separated list of Place fields to display. (optional)
    pollFields := []string{"PollFields_example"} // []string | A comma separated list of Poll fields to display. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TweetsApi.FindTweetsById(context.Background()).Ids(ids).Expansions(expansions).TweetFields(tweetFields).UserFields(userFields).MediaFields(mediaFields).PlaceFields(placeFields).PollFields(pollFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TweetsApi.FindTweetsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindTweetsById`: TweetLookupResponse
    fmt.Fprintf(os.Stdout, "Response from `TweetsApi.FindTweetsById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindTweetsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | A comma separated list of Tweet IDs. Up to 100 are allowed in a single request. | 
 **expansions** | **[]string** | A comma separated list of fields to expand. | 
 **tweetFields** | **[]string** | A comma separated list of Tweet fields to display. | 
 **userFields** | **[]string** | A comma separated list of User fields to display. | 
 **mediaFields** | **[]string** | A comma separated list of Media fields to display. | 
 **placeFields** | **[]string** | A comma separated list of Place fields to display. | 
 **pollFields** | **[]string** | A comma separated list of Poll fields to display. | 

### Return type

[**TweetLookupResponse**](TweetLookupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

