# TweetGeo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Coordinates** | Pointer to [**Point**](Point.md) |  | [optional] 
**PlaceId** | Pointer to **string** | The identifier for this place | [optional] 

## Methods

### NewTweetGeo

`func NewTweetGeo() *TweetGeo`

NewTweetGeo instantiates a new TweetGeo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTweetGeoWithDefaults

`func NewTweetGeoWithDefaults() *TweetGeo`

NewTweetGeoWithDefaults instantiates a new TweetGeo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoordinates

`func (o *TweetGeo) GetCoordinates() Point`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *TweetGeo) GetCoordinatesOk() (*Point, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *TweetGeo) SetCoordinates(v Point)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *TweetGeo) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.

### GetPlaceId

`func (o *TweetGeo) GetPlaceId() string`

GetPlaceId returns the PlaceId field if non-nil, zero value otherwise.

### GetPlaceIdOk

`func (o *TweetGeo) GetPlaceIdOk() (*string, bool)`

GetPlaceIdOk returns a tuple with the PlaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceId

`func (o *TweetGeo) SetPlaceId(v string)`

SetPlaceId sets PlaceId field to given value.

### HasPlaceId

`func (o *TweetGeo) HasPlaceId() bool`

HasPlaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


