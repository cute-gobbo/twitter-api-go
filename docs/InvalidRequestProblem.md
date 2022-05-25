# InvalidRequestProblem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Title** | **string** |  | 
**Detail** | **string** |  | 

## Methods

### NewInvalidRequestProblem

`func NewInvalidRequestProblem(title string, detail string, ) *InvalidRequestProblem`

NewInvalidRequestProblem instantiates a new InvalidRequestProblem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvalidRequestProblemWithDefaults

`func NewInvalidRequestProblemWithDefaults() *InvalidRequestProblem`

NewInvalidRequestProblemWithDefaults instantiates a new InvalidRequestProblem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InvalidRequestProblem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvalidRequestProblem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvalidRequestProblem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InvalidRequestProblem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetErrors

`func (o *InvalidRequestProblem) GetErrors() []map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InvalidRequestProblem) GetErrorsOk() (*[]map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InvalidRequestProblem) SetErrors(v []map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InvalidRequestProblem) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetTitle

`func (o *InvalidRequestProblem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InvalidRequestProblem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InvalidRequestProblem) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *InvalidRequestProblem) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *InvalidRequestProblem) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *InvalidRequestProblem) SetDetail(v string)`

SetDetail sets Detail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


