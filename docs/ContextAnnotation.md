# ContextAnnotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | [**ContextAnnotationDomainFields**](ContextAnnotationDomainFields.md) |  | 
**Entity** | [**ContextAnnotationEntityFields**](ContextAnnotationEntityFields.md) |  | 

## Methods

### NewContextAnnotation

`func NewContextAnnotation(domain ContextAnnotationDomainFields, entity ContextAnnotationEntityFields, ) *ContextAnnotation`

NewContextAnnotation instantiates a new ContextAnnotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextAnnotationWithDefaults

`func NewContextAnnotationWithDefaults() *ContextAnnotation`

NewContextAnnotationWithDefaults instantiates a new ContextAnnotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *ContextAnnotation) GetDomain() ContextAnnotationDomainFields`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ContextAnnotation) GetDomainOk() (*ContextAnnotationDomainFields, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ContextAnnotation) SetDomain(v ContextAnnotationDomainFields)`

SetDomain sets Domain field to given value.


### GetEntity

`func (o *ContextAnnotation) GetEntity() ContextAnnotationEntityFields`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *ContextAnnotation) GetEntityOk() (*ContextAnnotationEntityFields, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *ContextAnnotation) SetEntity(v ContextAnnotationEntityFields)`

SetEntity sets Entity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


