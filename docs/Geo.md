# Geo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Bbox** | **[]float64** |  | 
**Geometry** | Pointer to [**Point**](Point.md) |  | [optional] 
**Properties** | **map[string]interface{}** |  | 

## Methods

### NewGeo

`func NewGeo(type_ string, bbox []float64, properties map[string]interface{}, ) *Geo`

NewGeo instantiates a new Geo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoWithDefaults

`func NewGeoWithDefaults() *Geo`

NewGeoWithDefaults instantiates a new Geo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Geo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Geo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Geo) SetType(v string)`

SetType sets Type field to given value.


### GetBbox

`func (o *Geo) GetBbox() []float64`

GetBbox returns the Bbox field if non-nil, zero value otherwise.

### GetBboxOk

`func (o *Geo) GetBboxOk() (*[]float64, bool)`

GetBboxOk returns a tuple with the Bbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBbox

`func (o *Geo) SetBbox(v []float64)`

SetBbox sets Bbox field to given value.


### GetGeometry

`func (o *Geo) GetGeometry() Point`

GetGeometry returns the Geometry field if non-nil, zero value otherwise.

### GetGeometryOk

`func (o *Geo) GetGeometryOk() (*Point, bool)`

GetGeometryOk returns a tuple with the Geometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometry

`func (o *Geo) SetGeometry(v Point)`

SetGeometry sets Geometry field to given value.

### HasGeometry

`func (o *Geo) HasGeometry() bool`

HasGeometry returns a boolean if a field has been set.

### GetProperties

`func (o *Geo) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Geo) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Geo) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


