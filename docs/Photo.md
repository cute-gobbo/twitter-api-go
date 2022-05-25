# Photo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**MediaKey** | Pointer to **string** | The Media Key identifier for this attachment. | [optional] 
**Height** | Pointer to **int32** | The height of the media in pixels | [optional] 
**Width** | Pointer to **int32** | The width of the media in pixels | [optional] 

## Methods

### NewPhoto

`func NewPhoto() *Photo`

NewPhoto instantiates a new Photo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhotoWithDefaults

`func NewPhotoWithDefaults() *Photo`

NewPhotoWithDefaults instantiates a new Photo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Photo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Photo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Photo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Photo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *Photo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Photo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Photo) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Photo) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMediaKey

`func (o *Photo) GetMediaKey() string`

GetMediaKey returns the MediaKey field if non-nil, zero value otherwise.

### GetMediaKeyOk

`func (o *Photo) GetMediaKeyOk() (*string, bool)`

GetMediaKeyOk returns a tuple with the MediaKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaKey

`func (o *Photo) SetMediaKey(v string)`

SetMediaKey sets MediaKey field to given value.

### HasMediaKey

`func (o *Photo) HasMediaKey() bool`

HasMediaKey returns a boolean if a field has been set.

### GetHeight

`func (o *Photo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Photo) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Photo) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *Photo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetWidth

`func (o *Photo) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Photo) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Photo) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *Photo) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


