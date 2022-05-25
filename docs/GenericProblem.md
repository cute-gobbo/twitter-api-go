# GenericProblem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Status** | **int32** |  | 
**Title** | **string** |  | 
**Detail** | **string** |  | 

## Methods

### NewGenericProblem

`func NewGenericProblem(status int32, title string, detail string, ) *GenericProblem`

NewGenericProblem instantiates a new GenericProblem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericProblemWithDefaults

`func NewGenericProblemWithDefaults() *GenericProblem`

NewGenericProblemWithDefaults instantiates a new GenericProblem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GenericProblem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GenericProblem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GenericProblem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GenericProblem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *GenericProblem) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GenericProblem) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GenericProblem) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *GenericProblem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GenericProblem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GenericProblem) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *GenericProblem) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *GenericProblem) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *GenericProblem) SetDetail(v string)`

SetDetail sets Detail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


