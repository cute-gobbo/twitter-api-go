# URLImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | A validly formatted URL. | [optional] 
**Height** | Pointer to **int32** | The height of the media in pixels | [optional] 
**Width** | Pointer to **int32** | The width of the media in pixels | [optional] 

## Methods

### NewURLImage

`func NewURLImage() *URLImage`

NewURLImage instantiates a new URLImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewURLImageWithDefaults

`func NewURLImageWithDefaults() *URLImage`

NewURLImageWithDefaults instantiates a new URLImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *URLImage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *URLImage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *URLImage) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *URLImage) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHeight

`func (o *URLImage) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *URLImage) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *URLImage) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *URLImage) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetWidth

`func (o *URLImage) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *URLImage) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *URLImage) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *URLImage) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


