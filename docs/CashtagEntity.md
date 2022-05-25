# CashtagEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **int32** | Index (zero-based) at which position this entity starts. | 
**End** | **int32** | Index (zero-based) at which position this entity ends. | 
**Tag** | **string** |  | 

## Methods

### NewCashtagEntity

`func NewCashtagEntity(start int32, end int32, tag string, ) *CashtagEntity`

NewCashtagEntity instantiates a new CashtagEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashtagEntityWithDefaults

`func NewCashtagEntityWithDefaults() *CashtagEntity`

NewCashtagEntityWithDefaults instantiates a new CashtagEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *CashtagEntity) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CashtagEntity) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CashtagEntity) SetStart(v int32)`

SetStart sets Start field to given value.


### GetEnd

`func (o *CashtagEntity) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CashtagEntity) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CashtagEntity) SetEnd(v int32)`

SetEnd sets End field to given value.


### GetTag

`func (o *CashtagEntity) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CashtagEntity) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CashtagEntity) SetTag(v string)`

SetTag sets Tag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


