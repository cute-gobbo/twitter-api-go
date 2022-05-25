# Video

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**PreviewImageUrl** | Pointer to **string** |  | [optional] 
**DurationMs** | Pointer to **int32** |  | [optional] 
**PublicMetrics** | Pointer to **map[string]interface{}** | Engagement metrics for the Media at the time of the request. | [optional] 
**NonPublicMetrics** | Pointer to **map[string]interface{}** | Nonpublic engagement metrics for the Media at the time of the request. | [optional] 
**OrganicMetrics** | Pointer to **map[string]interface{}** | Organic nonpublic engagement metrics for the Media at the time of the request. | [optional] 
**PromotedMetrics** | Pointer to **map[string]interface{}** | Promoted nonpublic engagement metrics for the Media at the time of the request. | [optional] 
**MediaKey** | Pointer to **string** | The Media Key identifier for this attachment. | [optional] 
**Height** | Pointer to **int32** | The height of the media in pixels | [optional] 
**Width** | Pointer to **int32** | The width of the media in pixels | [optional] 

## Methods

### NewVideo

`func NewVideo() *Video`

NewVideo instantiates a new Video object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoWithDefaults

`func NewVideoWithDefaults() *Video`

NewVideoWithDefaults instantiates a new Video object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Video) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Video) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Video) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Video) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPreviewImageUrl

`func (o *Video) GetPreviewImageUrl() string`

GetPreviewImageUrl returns the PreviewImageUrl field if non-nil, zero value otherwise.

### GetPreviewImageUrlOk

`func (o *Video) GetPreviewImageUrlOk() (*string, bool)`

GetPreviewImageUrlOk returns a tuple with the PreviewImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewImageUrl

`func (o *Video) SetPreviewImageUrl(v string)`

SetPreviewImageUrl sets PreviewImageUrl field to given value.

### HasPreviewImageUrl

`func (o *Video) HasPreviewImageUrl() bool`

HasPreviewImageUrl returns a boolean if a field has been set.

### GetDurationMs

`func (o *Video) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *Video) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *Video) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *Video) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetPublicMetrics

`func (o *Video) GetPublicMetrics() map[string]interface{}`

GetPublicMetrics returns the PublicMetrics field if non-nil, zero value otherwise.

### GetPublicMetricsOk

`func (o *Video) GetPublicMetricsOk() (*map[string]interface{}, bool)`

GetPublicMetricsOk returns a tuple with the PublicMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicMetrics

`func (o *Video) SetPublicMetrics(v map[string]interface{})`

SetPublicMetrics sets PublicMetrics field to given value.

### HasPublicMetrics

`func (o *Video) HasPublicMetrics() bool`

HasPublicMetrics returns a boolean if a field has been set.

### GetNonPublicMetrics

`func (o *Video) GetNonPublicMetrics() map[string]interface{}`

GetNonPublicMetrics returns the NonPublicMetrics field if non-nil, zero value otherwise.

### GetNonPublicMetricsOk

`func (o *Video) GetNonPublicMetricsOk() (*map[string]interface{}, bool)`

GetNonPublicMetricsOk returns a tuple with the NonPublicMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonPublicMetrics

`func (o *Video) SetNonPublicMetrics(v map[string]interface{})`

SetNonPublicMetrics sets NonPublicMetrics field to given value.

### HasNonPublicMetrics

`func (o *Video) HasNonPublicMetrics() bool`

HasNonPublicMetrics returns a boolean if a field has been set.

### GetOrganicMetrics

`func (o *Video) GetOrganicMetrics() map[string]interface{}`

GetOrganicMetrics returns the OrganicMetrics field if non-nil, zero value otherwise.

### GetOrganicMetricsOk

`func (o *Video) GetOrganicMetricsOk() (*map[string]interface{}, bool)`

GetOrganicMetricsOk returns a tuple with the OrganicMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganicMetrics

`func (o *Video) SetOrganicMetrics(v map[string]interface{})`

SetOrganicMetrics sets OrganicMetrics field to given value.

### HasOrganicMetrics

`func (o *Video) HasOrganicMetrics() bool`

HasOrganicMetrics returns a boolean if a field has been set.

### GetPromotedMetrics

`func (o *Video) GetPromotedMetrics() map[string]interface{}`

GetPromotedMetrics returns the PromotedMetrics field if non-nil, zero value otherwise.

### GetPromotedMetricsOk

`func (o *Video) GetPromotedMetricsOk() (*map[string]interface{}, bool)`

GetPromotedMetricsOk returns a tuple with the PromotedMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotedMetrics

`func (o *Video) SetPromotedMetrics(v map[string]interface{})`

SetPromotedMetrics sets PromotedMetrics field to given value.

### HasPromotedMetrics

`func (o *Video) HasPromotedMetrics() bool`

HasPromotedMetrics returns a boolean if a field has been set.

### GetMediaKey

`func (o *Video) GetMediaKey() string`

GetMediaKey returns the MediaKey field if non-nil, zero value otherwise.

### GetMediaKeyOk

`func (o *Video) GetMediaKeyOk() (*string, bool)`

GetMediaKeyOk returns a tuple with the MediaKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaKey

`func (o *Video) SetMediaKey(v string)`

SetMediaKey sets MediaKey field to given value.

### HasMediaKey

`func (o *Video) HasMediaKey() bool`

HasMediaKey returns a boolean if a field has been set.

### GetHeight

`func (o *Video) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Video) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Video) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *Video) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetWidth

`func (o *Video) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Video) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Video) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *Video) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


