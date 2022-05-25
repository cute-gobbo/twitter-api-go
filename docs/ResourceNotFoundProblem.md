# ResourceNotFoundProblem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Parameter** | **string** |  | 
**Value** | **interface{}** | Value will match the schema of the field. | 
**ResourceId** | **string** |  | 
**ResourceType** | **string** |  | 
**Title** | **string** |  | 
**Detail** | **string** |  | 

## Methods

### NewResourceNotFoundProblem

`func NewResourceNotFoundProblem(parameter string, value interface{}, resourceId string, resourceType string, title string, detail string, ) *ResourceNotFoundProblem`

NewResourceNotFoundProblem instantiates a new ResourceNotFoundProblem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceNotFoundProblemWithDefaults

`func NewResourceNotFoundProblemWithDefaults() *ResourceNotFoundProblem`

NewResourceNotFoundProblemWithDefaults instantiates a new ResourceNotFoundProblem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResourceNotFoundProblem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceNotFoundProblem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceNotFoundProblem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResourceNotFoundProblem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetParameter

`func (o *ResourceNotFoundProblem) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *ResourceNotFoundProblem) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *ResourceNotFoundProblem) SetParameter(v string)`

SetParameter sets Parameter field to given value.


### GetValue

`func (o *ResourceNotFoundProblem) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ResourceNotFoundProblem) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ResourceNotFoundProblem) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *ResourceNotFoundProblem) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ResourceNotFoundProblem) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetResourceId

`func (o *ResourceNotFoundProblem) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ResourceNotFoundProblem) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ResourceNotFoundProblem) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceType

`func (o *ResourceNotFoundProblem) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceNotFoundProblem) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceNotFoundProblem) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetTitle

`func (o *ResourceNotFoundProblem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResourceNotFoundProblem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResourceNotFoundProblem) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *ResourceNotFoundProblem) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ResourceNotFoundProblem) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ResourceNotFoundProblem) SetDetail(v string)`

SetDetail sets Detail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


