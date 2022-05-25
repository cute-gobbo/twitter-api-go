# UserLookupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]User**](User.md) |  | [optional] 
**Includes** | Pointer to [**Expansions**](Expansions.md) |  | [optional] 
**Errors** | Pointer to [**[]Problem**](Problem.md) |  | [optional] 

## Methods

### NewUserLookupResponse

`func NewUserLookupResponse() *UserLookupResponse`

NewUserLookupResponse instantiates a new UserLookupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLookupResponseWithDefaults

`func NewUserLookupResponseWithDefaults() *UserLookupResponse`

NewUserLookupResponseWithDefaults instantiates a new UserLookupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *UserLookupResponse) GetData() []User`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserLookupResponse) GetDataOk() (*[]User, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserLookupResponse) SetData(v []User)`

SetData sets Data field to given value.

### HasData

`func (o *UserLookupResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIncludes

`func (o *UserLookupResponse) GetIncludes() Expansions`

GetIncludes returns the Includes field if non-nil, zero value otherwise.

### GetIncludesOk

`func (o *UserLookupResponse) GetIncludesOk() (*Expansions, bool)`

GetIncludesOk returns a tuple with the Includes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludes

`func (o *UserLookupResponse) SetIncludes(v Expansions)`

SetIncludes sets Includes field to given value.

### HasIncludes

`func (o *UserLookupResponse) HasIncludes() bool`

HasIncludes returns a boolean if a field has been set.

### GetErrors

`func (o *UserLookupResponse) GetErrors() []Problem`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *UserLookupResponse) GetErrorsOk() (*[]Problem, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *UserLookupResponse) SetErrors(v []Problem)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *UserLookupResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


