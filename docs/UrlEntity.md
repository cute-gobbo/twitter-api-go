# UrlEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **int32** | Index (zero-based) at which position this entity starts. | 
**End** | **int32** | Index (zero-based) at which position this entity ends. | 
**Url** | **string** | A validly formatted URL. | 
**ExpandedUrl** | Pointer to **string** | A validly formatted URL. | [optional] 
**DisplayUrl** | Pointer to **string** | The URL as displayed in the Twitter client. | [optional] 
**UnwoundUrl** | Pointer to **string** | Fully resolved url | [optional] 
**Status** | Pointer to **int32** | HTTP Status Code. | [optional] 
**Title** | Pointer to **string** | Title of the page the URL points to. | [optional] 
**Description** | Pointer to **string** | Description of the URL landing page. | [optional] 
**Images** | Pointer to [**[]URLImage**](URLImage.md) |  | [optional] 

## Methods

### NewUrlEntity

`func NewUrlEntity(start int32, end int32, url string, ) *UrlEntity`

NewUrlEntity instantiates a new UrlEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUrlEntityWithDefaults

`func NewUrlEntityWithDefaults() *UrlEntity`

NewUrlEntityWithDefaults instantiates a new UrlEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *UrlEntity) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *UrlEntity) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *UrlEntity) SetStart(v int32)`

SetStart sets Start field to given value.


### GetEnd

`func (o *UrlEntity) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *UrlEntity) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *UrlEntity) SetEnd(v int32)`

SetEnd sets End field to given value.


### GetUrl

`func (o *UrlEntity) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UrlEntity) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UrlEntity) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetExpandedUrl

`func (o *UrlEntity) GetExpandedUrl() string`

GetExpandedUrl returns the ExpandedUrl field if non-nil, zero value otherwise.

### GetExpandedUrlOk

`func (o *UrlEntity) GetExpandedUrlOk() (*string, bool)`

GetExpandedUrlOk returns a tuple with the ExpandedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandedUrl

`func (o *UrlEntity) SetExpandedUrl(v string)`

SetExpandedUrl sets ExpandedUrl field to given value.

### HasExpandedUrl

`func (o *UrlEntity) HasExpandedUrl() bool`

HasExpandedUrl returns a boolean if a field has been set.

### GetDisplayUrl

`func (o *UrlEntity) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *UrlEntity) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *UrlEntity) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.

### HasDisplayUrl

`func (o *UrlEntity) HasDisplayUrl() bool`

HasDisplayUrl returns a boolean if a field has been set.

### GetUnwoundUrl

`func (o *UrlEntity) GetUnwoundUrl() string`

GetUnwoundUrl returns the UnwoundUrl field if non-nil, zero value otherwise.

### GetUnwoundUrlOk

`func (o *UrlEntity) GetUnwoundUrlOk() (*string, bool)`

GetUnwoundUrlOk returns a tuple with the UnwoundUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnwoundUrl

`func (o *UrlEntity) SetUnwoundUrl(v string)`

SetUnwoundUrl sets UnwoundUrl field to given value.

### HasUnwoundUrl

`func (o *UrlEntity) HasUnwoundUrl() bool`

HasUnwoundUrl returns a boolean if a field has been set.

### GetStatus

`func (o *UrlEntity) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UrlEntity) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UrlEntity) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UrlEntity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *UrlEntity) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UrlEntity) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UrlEntity) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UrlEntity) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *UrlEntity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UrlEntity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UrlEntity) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UrlEntity) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImages

`func (o *UrlEntity) GetImages() []URLImage`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *UrlEntity) GetImagesOk() (*[]URLImage, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *UrlEntity) SetImages(v []URLImage)`

SetImages sets Images field to given value.

### HasImages

`func (o *UrlEntity) HasImages() bool`

HasImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


