# TweetWithheld

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Copyright** | **bool** | Indicates if the content is being withheld for on the basis of copyright infringement. | 
**CountryCodes** | **[]string** | Provides a list of countries where this content is not available. | 
**Scope** | Pointer to **string** | Indicates whether the content being withheld is the &#x60;tweet&#x60; or a &#x60;user&#x60;. | [optional] 

## Methods

### NewTweetWithheld

`func NewTweetWithheld(copyright bool, countryCodes []string, ) *TweetWithheld`

NewTweetWithheld instantiates a new TweetWithheld object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTweetWithheldWithDefaults

`func NewTweetWithheldWithDefaults() *TweetWithheld`

NewTweetWithheldWithDefaults instantiates a new TweetWithheld object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopyright

`func (o *TweetWithheld) GetCopyright() bool`

GetCopyright returns the Copyright field if non-nil, zero value otherwise.

### GetCopyrightOk

`func (o *TweetWithheld) GetCopyrightOk() (*bool, bool)`

GetCopyrightOk returns a tuple with the Copyright field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyright

`func (o *TweetWithheld) SetCopyright(v bool)`

SetCopyright sets Copyright field to given value.


### GetCountryCodes

`func (o *TweetWithheld) GetCountryCodes() []string`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *TweetWithheld) GetCountryCodesOk() (*[]string, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *TweetWithheld) SetCountryCodes(v []string)`

SetCountryCodes sets CountryCodes field to given value.


### GetScope

`func (o *TweetWithheld) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *TweetWithheld) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *TweetWithheld) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *TweetWithheld) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


