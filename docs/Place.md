# Place

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The identifier for this place | 
**Name** | Pointer to **string** | The human readable name of this place. | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**PlaceType** | Pointer to [**PlaceType**](PlaceType.md) |  | [optional] 
**FullName** | **string** |  | 
**Country** | Pointer to **string** |  | [optional] 
**ContainedWithin** | Pointer to **[]string** |  | [optional] 
**Geo** | Pointer to [**Geo**](Geo.md) |  | [optional] 

## Methods

### NewPlace

`func NewPlace(id string, fullName string, ) *Place`

NewPlace instantiates a new Place object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaceWithDefaults

`func NewPlaceWithDefaults() *Place`

NewPlaceWithDefaults instantiates a new Place object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Place) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Place) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Place) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Place) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Place) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Place) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Place) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCountryCode

`func (o *Place) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Place) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Place) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *Place) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetPlaceType

`func (o *Place) GetPlaceType() PlaceType`

GetPlaceType returns the PlaceType field if non-nil, zero value otherwise.

### GetPlaceTypeOk

`func (o *Place) GetPlaceTypeOk() (*PlaceType, bool)`

GetPlaceTypeOk returns a tuple with the PlaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceType

`func (o *Place) SetPlaceType(v PlaceType)`

SetPlaceType sets PlaceType field to given value.

### HasPlaceType

`func (o *Place) HasPlaceType() bool`

HasPlaceType returns a boolean if a field has been set.

### GetFullName

`func (o *Place) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *Place) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *Place) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetCountry

`func (o *Place) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Place) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Place) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Place) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetContainedWithin

`func (o *Place) GetContainedWithin() []string`

GetContainedWithin returns the ContainedWithin field if non-nil, zero value otherwise.

### GetContainedWithinOk

`func (o *Place) GetContainedWithinOk() (*[]string, bool)`

GetContainedWithinOk returns a tuple with the ContainedWithin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainedWithin

`func (o *Place) SetContainedWithin(v []string)`

SetContainedWithin sets ContainedWithin field to given value.

### HasContainedWithin

`func (o *Place) HasContainedWithin() bool`

HasContainedWithin returns a boolean if a field has been set.

### GetGeo

`func (o *Place) GetGeo() Geo`

GetGeo returns the Geo field if non-nil, zero value otherwise.

### GetGeoOk

`func (o *Place) GetGeoOk() (*Geo, bool)`

GetGeoOk returns a tuple with the Geo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeo

`func (o *Place) SetGeo(v Geo)`

SetGeo sets Geo field to given value.

### HasGeo

`func (o *Place) HasGeo() bool`

HasGeo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


