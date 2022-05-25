# SingleUserLookupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**User**](User.md) |  | [optional] 
**Includes** | Pointer to [**Expansions**](Expansions.md) |  | [optional] 
**Errors** | Pointer to [**[]Problem**](Problem.md) |  | [optional] 

## Methods

### NewSingleUserLookupResponse

`func NewSingleUserLookupResponse() *SingleUserLookupResponse`

NewSingleUserLookupResponse instantiates a new SingleUserLookupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleUserLookupResponseWithDefaults

`func NewSingleUserLookupResponseWithDefaults() *SingleUserLookupResponse`

NewSingleUserLookupResponseWithDefaults instantiates a new SingleUserLookupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SingleUserLookupResponse) GetData() User`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SingleUserLookupResponse) GetDataOk() (*User, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SingleUserLookupResponse) SetData(v User)`

SetData sets Data field to given value.

### HasData

`func (o *SingleUserLookupResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIncludes

`func (o *SingleUserLookupResponse) GetIncludes() Expansions`

GetIncludes returns the Includes field if non-nil, zero value otherwise.

### GetIncludesOk

`func (o *SingleUserLookupResponse) GetIncludesOk() (*Expansions, bool)`

GetIncludesOk returns a tuple with the Includes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludes

`func (o *SingleUserLookupResponse) SetIncludes(v Expansions)`

SetIncludes sets Includes field to given value.

### HasIncludes

`func (o *SingleUserLookupResponse) HasIncludes() bool`

HasIncludes returns a boolean if a field has been set.

### GetErrors

`func (o *SingleUserLookupResponse) GetErrors() []Problem`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SingleUserLookupResponse) GetErrorsOk() (*[]Problem, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SingleUserLookupResponse) SetErrors(v []Problem)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SingleUserLookupResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


