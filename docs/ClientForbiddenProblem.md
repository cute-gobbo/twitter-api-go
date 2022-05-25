# ClientForbiddenProblem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**RegistrationUrl** | Pointer to **string** |  | [optional] 
**Title** | **string** |  | 
**Detail** | **string** |  | 

## Methods

### NewClientForbiddenProblem

`func NewClientForbiddenProblem(title string, detail string, ) *ClientForbiddenProblem`

NewClientForbiddenProblem instantiates a new ClientForbiddenProblem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientForbiddenProblemWithDefaults

`func NewClientForbiddenProblemWithDefaults() *ClientForbiddenProblem`

NewClientForbiddenProblemWithDefaults instantiates a new ClientForbiddenProblem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ClientForbiddenProblem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientForbiddenProblem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientForbiddenProblem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClientForbiddenProblem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetReason

`func (o *ClientForbiddenProblem) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ClientForbiddenProblem) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ClientForbiddenProblem) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ClientForbiddenProblem) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRegistrationUrl

`func (o *ClientForbiddenProblem) GetRegistrationUrl() string`

GetRegistrationUrl returns the RegistrationUrl field if non-nil, zero value otherwise.

### GetRegistrationUrlOk

`func (o *ClientForbiddenProblem) GetRegistrationUrlOk() (*string, bool)`

GetRegistrationUrlOk returns a tuple with the RegistrationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationUrl

`func (o *ClientForbiddenProblem) SetRegistrationUrl(v string)`

SetRegistrationUrl sets RegistrationUrl field to given value.

### HasRegistrationUrl

`func (o *ClientForbiddenProblem) HasRegistrationUrl() bool`

HasRegistrationUrl returns a boolean if a field has been set.

### GetTitle

`func (o *ClientForbiddenProblem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ClientForbiddenProblem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ClientForbiddenProblem) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *ClientForbiddenProblem) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ClientForbiddenProblem) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ClientForbiddenProblem) SetDetail(v string)`

SetDetail sets Detail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


