# FullTextEntities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Urls** | Pointer to [**[]UrlEntity**](UrlEntity.md) |  | [optional] 
**Hashtags** | Pointer to [**[]HashtagEntity**](HashtagEntity.md) |  | [optional] 
**Mentions** | Pointer to [**[]MentionEntity**](MentionEntity.md) |  | [optional] 
**Cashtags** | Pointer to [**[]CashtagEntity**](CashtagEntity.md) |  | [optional] 

## Methods

### NewFullTextEntities

`func NewFullTextEntities() *FullTextEntities`

NewFullTextEntities instantiates a new FullTextEntities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullTextEntitiesWithDefaults

`func NewFullTextEntitiesWithDefaults() *FullTextEntities`

NewFullTextEntitiesWithDefaults instantiates a new FullTextEntities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrls

`func (o *FullTextEntities) GetUrls() []UrlEntity`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *FullTextEntities) GetUrlsOk() (*[]UrlEntity, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *FullTextEntities) SetUrls(v []UrlEntity)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *FullTextEntities) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetHashtags

`func (o *FullTextEntities) GetHashtags() []HashtagEntity`

GetHashtags returns the Hashtags field if non-nil, zero value otherwise.

### GetHashtagsOk

`func (o *FullTextEntities) GetHashtagsOk() (*[]HashtagEntity, bool)`

GetHashtagsOk returns a tuple with the Hashtags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashtags

`func (o *FullTextEntities) SetHashtags(v []HashtagEntity)`

SetHashtags sets Hashtags field to given value.

### HasHashtags

`func (o *FullTextEntities) HasHashtags() bool`

HasHashtags returns a boolean if a field has been set.

### GetMentions

`func (o *FullTextEntities) GetMentions() []MentionEntity`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *FullTextEntities) GetMentionsOk() (*[]MentionEntity, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *FullTextEntities) SetMentions(v []MentionEntity)`

SetMentions sets Mentions field to given value.

### HasMentions

`func (o *FullTextEntities) HasMentions() bool`

HasMentions returns a boolean if a field has been set.

### GetCashtags

`func (o *FullTextEntities) GetCashtags() []CashtagEntity`

GetCashtags returns the Cashtags field if non-nil, zero value otherwise.

### GetCashtagsOk

`func (o *FullTextEntities) GetCashtagsOk() (*[]CashtagEntity, bool)`

GetCashtagsOk returns a tuple with the Cashtags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashtags

`func (o *FullTextEntities) SetCashtags(v []CashtagEntity)`

SetCashtags sets Cashtags field to given value.

### HasCashtags

`func (o *FullTextEntities) HasCashtags() bool`

HasCashtags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


