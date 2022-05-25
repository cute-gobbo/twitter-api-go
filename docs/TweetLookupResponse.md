# TweetLookupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Tweet**](Tweet.md) |  | [optional] 
**Includes** | Pointer to [**Expansions**](Expansions.md) |  | [optional] 
**Errors** | Pointer to [**[]Problem**](Problem.md) |  | [optional] 

## Methods

### NewTweetLookupResponse

`func NewTweetLookupResponse() *TweetLookupResponse`

NewTweetLookupResponse instantiates a new TweetLookupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTweetLookupResponseWithDefaults

`func NewTweetLookupResponseWithDefaults() *TweetLookupResponse`

NewTweetLookupResponseWithDefaults instantiates a new TweetLookupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TweetLookupResponse) GetData() []Tweet`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TweetLookupResponse) GetDataOk() (*[]Tweet, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TweetLookupResponse) SetData(v []Tweet)`

SetData sets Data field to given value.

### HasData

`func (o *TweetLookupResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIncludes

`func (o *TweetLookupResponse) GetIncludes() Expansions`

GetIncludes returns the Includes field if non-nil, zero value otherwise.

### GetIncludesOk

`func (o *TweetLookupResponse) GetIncludesOk() (*Expansions, bool)`

GetIncludesOk returns a tuple with the Includes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludes

`func (o *TweetLookupResponse) SetIncludes(v Expansions)`

SetIncludes sets Includes field to given value.

### HasIncludes

`func (o *TweetLookupResponse) HasIncludes() bool`

HasIncludes returns a boolean if a field has been set.

### GetErrors

`func (o *TweetLookupResponse) GetErrors() []Problem`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *TweetLookupResponse) GetErrorsOk() (*[]Problem, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *TweetLookupResponse) SetErrors(v []Problem)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *TweetLookupResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


