# UserWithheld

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCodes** | **[]string** | Provides a list of countries where this content is not available. | 
**Scope** | Pointer to **string** | Indicates that the content being withheld is a &#x60;user&#x60;. | [optional] 

## Methods

### NewUserWithheld

`func NewUserWithheld(countryCodes []string, ) *UserWithheld`

NewUserWithheld instantiates a new UserWithheld object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithheldWithDefaults

`func NewUserWithheldWithDefaults() *UserWithheld`

NewUserWithheldWithDefaults instantiates a new UserWithheld object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCodes

`func (o *UserWithheld) GetCountryCodes() []string`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *UserWithheld) GetCountryCodesOk() (*[]string, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *UserWithheld) SetCountryCodes(v []string)`

SetCountryCodes sets CountryCodes field to given value.


### GetScope

`func (o *UserWithheld) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UserWithheld) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UserWithheld) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *UserWithheld) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


