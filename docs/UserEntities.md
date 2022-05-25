# UserEntities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to [**UserEntitiesUrl**](UserEntitiesUrl.md) |  | [optional] 
**Description** | Pointer to [**FullTextEntities**](FullTextEntities.md) |  | [optional] 

## Methods

### NewUserEntities

`func NewUserEntities() *UserEntities`

NewUserEntities instantiates a new UserEntities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserEntitiesWithDefaults

`func NewUserEntitiesWithDefaults() *UserEntities`

NewUserEntitiesWithDefaults instantiates a new UserEntities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *UserEntities) GetUrl() UserEntitiesUrl`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UserEntities) GetUrlOk() (*UserEntitiesUrl, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UserEntities) SetUrl(v UserEntitiesUrl)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UserEntities) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDescription

`func (o *UserEntities) GetDescription() FullTextEntities`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserEntities) GetDescriptionOk() (*FullTextEntities, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserEntities) SetDescription(v FullTextEntities)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserEntities) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


