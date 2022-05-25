# Media

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaKey** | Pointer to **string** | The Media Key identifier for this attachment. | [optional] 
**Height** | Pointer to **int32** | The height of the media in pixels | [optional] 
**Width** | Pointer to **int32** | The width of the media in pixels | [optional] 

## Methods

### NewMedia

`func NewMedia() *Media`

NewMedia instantiates a new Media object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaWithDefaults

`func NewMediaWithDefaults() *Media`

NewMediaWithDefaults instantiates a new Media object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaKey

`func (o *Media) GetMediaKey() string`

GetMediaKey returns the MediaKey field if non-nil, zero value otherwise.

### GetMediaKeyOk

`func (o *Media) GetMediaKeyOk() (*string, bool)`

GetMediaKeyOk returns a tuple with the MediaKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaKey

`func (o *Media) SetMediaKey(v string)`

SetMediaKey sets MediaKey field to given value.

### HasMediaKey

`func (o *Media) HasMediaKey() bool`

HasMediaKey returns a boolean if a field has been set.

### GetHeight

`func (o *Media) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Media) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Media) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *Media) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetWidth

`func (o *Media) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Media) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Media) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *Media) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


