# URLFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | A validly formatted URL. | 
**ExpandedUrl** | Pointer to **string** | A validly formatted URL. | [optional] 
**DisplayUrl** | Pointer to **string** | The URL as displayed in the Twitter client. | [optional] 
**UnwoundUrl** | Pointer to **string** | Fully resolved url | [optional] 
**Status** | Pointer to **int32** | HTTP Status Code. | [optional] 
**Title** | Pointer to **string** | Title of the page the URL points to. | [optional] 
**Description** | Pointer to **string** | Description of the URL landing page. | [optional] 
**Images** | Pointer to [**[]URLImage**](URLImage.md) |  | [optional] 

## Methods

### NewURLFields

`func NewURLFields(url string, ) *URLFields`

NewURLFields instantiates a new URLFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewURLFieldsWithDefaults

`func NewURLFieldsWithDefaults() *URLFields`

NewURLFieldsWithDefaults instantiates a new URLFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *URLFields) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *URLFields) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *URLFields) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetExpandedUrl

`func (o *URLFields) GetExpandedUrl() string`

GetExpandedUrl returns the ExpandedUrl field if non-nil, zero value otherwise.

### GetExpandedUrlOk

`func (o *URLFields) GetExpandedUrlOk() (*string, bool)`

GetExpandedUrlOk returns a tuple with the ExpandedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandedUrl

`func (o *URLFields) SetExpandedUrl(v string)`

SetExpandedUrl sets ExpandedUrl field to given value.

### HasExpandedUrl

`func (o *URLFields) HasExpandedUrl() bool`

HasExpandedUrl returns a boolean if a field has been set.

### GetDisplayUrl

`func (o *URLFields) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *URLFields) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *URLFields) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.

### HasDisplayUrl

`func (o *URLFields) HasDisplayUrl() bool`

HasDisplayUrl returns a boolean if a field has been set.

### GetUnwoundUrl

`func (o *URLFields) GetUnwoundUrl() string`

GetUnwoundUrl returns the UnwoundUrl field if non-nil, zero value otherwise.

### GetUnwoundUrlOk

`func (o *URLFields) GetUnwoundUrlOk() (*string, bool)`

GetUnwoundUrlOk returns a tuple with the UnwoundUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnwoundUrl

`func (o *URLFields) SetUnwoundUrl(v string)`

SetUnwoundUrl sets UnwoundUrl field to given value.

### HasUnwoundUrl

`func (o *URLFields) HasUnwoundUrl() bool`

HasUnwoundUrl returns a boolean if a field has been set.

### GetStatus

`func (o *URLFields) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *URLFields) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *URLFields) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *URLFields) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *URLFields) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *URLFields) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *URLFields) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *URLFields) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *URLFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *URLFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *URLFields) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *URLFields) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImages

`func (o *URLFields) GetImages() []URLImage`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *URLFields) GetImagesOk() (*[]URLImage, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *URLFields) SetImages(v []URLImage)`

SetImages sets Images field to given value.

### HasImages

`func (o *URLFields) HasImages() bool`

HasImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


