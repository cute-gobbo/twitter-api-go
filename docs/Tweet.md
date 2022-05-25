# Tweet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of this Tweet. This is returned as a string in order to avoid complications with languages and tools that cannot handle large integers. | 
**CreatedAt** | Pointer to **time.Time** | Creation time of the Tweet. | [optional] 
**Text** | **string** | The content of the Tweet. | 
**AuthorId** | Pointer to **string** | Unique identifier of this User. This is returned as a string in order to avoid complications with languages and tools that cannot handle large integers. | [optional] 
**InReplyToUserId** | Pointer to **string** | Unique identifier of this User. This is returned as a string in order to avoid complications with languages and tools that cannot handle large integers. | [optional] 
**ReferencedTweets** | Pointer to [**[]TweetReferencedTweets**](TweetReferencedTweets.md) | A list of Tweets this Tweet refers to. For example, if the parent Tweet is a Retweet, a Quoted Tweet or a Reply, it will include the related Tweet referenced to by its parent. | [optional] 
**Attachments** | Pointer to [**TweetAttachments**](TweetAttachments.md) |  | [optional] 
**ContextAnnotations** | Pointer to [**[]ContextAnnotation**](ContextAnnotation.md) |  | [optional] 
**Withheld** | Pointer to [**TweetWithheld**](TweetWithheld.md) |  | [optional] 
**Geo** | Pointer to [**TweetGeo**](TweetGeo.md) |  | [optional] 
**Entities** | Pointer to [**FullTextEntities**](FullTextEntities.md) |  | [optional] 
**PublicMetrics** | Pointer to [**TweetPublicMetrics**](TweetPublicMetrics.md) |  | [optional] 
**PossiblySensitive** | Pointer to **bool** | Indicates if this Tweet contains URLs marked as sensitive, for example content suitable for mature audiences. | [optional] 
**Lang** | Pointer to **string** | Language of the Tweet, if detected by Twitter. Returned as a BCP47 language tag. | [optional] 
**Source** | Pointer to **string** | The name of the app the user Tweeted from. | [optional] 
**NonPublicMetrics** | Pointer to [**TweetNonPublicMetrics**](TweetNonPublicMetrics.md) |  | [optional] 
**PromotedMetrics** | Pointer to [**TweetPromotedMetrics**](TweetPromotedMetrics.md) |  | [optional] 
**OrganicMetrics** | Pointer to [**TweetOrganicMetrics**](TweetOrganicMetrics.md) |  | [optional] 

## Methods

### NewTweet

`func NewTweet(id string, text string, ) *Tweet`

NewTweet instantiates a new Tweet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTweetWithDefaults

`func NewTweetWithDefaults() *Tweet`

NewTweetWithDefaults instantiates a new Tweet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Tweet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tweet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tweet) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Tweet) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Tweet) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Tweet) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Tweet) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetText

`func (o *Tweet) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Tweet) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Tweet) SetText(v string)`

SetText sets Text field to given value.


### GetAuthorId

`func (o *Tweet) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *Tweet) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *Tweet) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *Tweet) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetInReplyToUserId

`func (o *Tweet) GetInReplyToUserId() string`

GetInReplyToUserId returns the InReplyToUserId field if non-nil, zero value otherwise.

### GetInReplyToUserIdOk

`func (o *Tweet) GetInReplyToUserIdOk() (*string, bool)`

GetInReplyToUserIdOk returns a tuple with the InReplyToUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInReplyToUserId

`func (o *Tweet) SetInReplyToUserId(v string)`

SetInReplyToUserId sets InReplyToUserId field to given value.

### HasInReplyToUserId

`func (o *Tweet) HasInReplyToUserId() bool`

HasInReplyToUserId returns a boolean if a field has been set.

### GetReferencedTweets

`func (o *Tweet) GetReferencedTweets() []TweetReferencedTweets`

GetReferencedTweets returns the ReferencedTweets field if non-nil, zero value otherwise.

### GetReferencedTweetsOk

`func (o *Tweet) GetReferencedTweetsOk() (*[]TweetReferencedTweets, bool)`

GetReferencedTweetsOk returns a tuple with the ReferencedTweets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedTweets

`func (o *Tweet) SetReferencedTweets(v []TweetReferencedTweets)`

SetReferencedTweets sets ReferencedTweets field to given value.

### HasReferencedTweets

`func (o *Tweet) HasReferencedTweets() bool`

HasReferencedTweets returns a boolean if a field has been set.

### GetAttachments

`func (o *Tweet) GetAttachments() TweetAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *Tweet) GetAttachmentsOk() (*TweetAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *Tweet) SetAttachments(v TweetAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *Tweet) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetContextAnnotations

`func (o *Tweet) GetContextAnnotations() []ContextAnnotation`

GetContextAnnotations returns the ContextAnnotations field if non-nil, zero value otherwise.

### GetContextAnnotationsOk

`func (o *Tweet) GetContextAnnotationsOk() (*[]ContextAnnotation, bool)`

GetContextAnnotationsOk returns a tuple with the ContextAnnotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextAnnotations

`func (o *Tweet) SetContextAnnotations(v []ContextAnnotation)`

SetContextAnnotations sets ContextAnnotations field to given value.

### HasContextAnnotations

`func (o *Tweet) HasContextAnnotations() bool`

HasContextAnnotations returns a boolean if a field has been set.

### GetWithheld

`func (o *Tweet) GetWithheld() TweetWithheld`

GetWithheld returns the Withheld field if non-nil, zero value otherwise.

### GetWithheldOk

`func (o *Tweet) GetWithheldOk() (*TweetWithheld, bool)`

GetWithheldOk returns a tuple with the Withheld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithheld

`func (o *Tweet) SetWithheld(v TweetWithheld)`

SetWithheld sets Withheld field to given value.

### HasWithheld

`func (o *Tweet) HasWithheld() bool`

HasWithheld returns a boolean if a field has been set.

### GetGeo

`func (o *Tweet) GetGeo() TweetGeo`

GetGeo returns the Geo field if non-nil, zero value otherwise.

### GetGeoOk

`func (o *Tweet) GetGeoOk() (*TweetGeo, bool)`

GetGeoOk returns a tuple with the Geo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeo

`func (o *Tweet) SetGeo(v TweetGeo)`

SetGeo sets Geo field to given value.

### HasGeo

`func (o *Tweet) HasGeo() bool`

HasGeo returns a boolean if a field has been set.

### GetEntities

`func (o *Tweet) GetEntities() FullTextEntities`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *Tweet) GetEntitiesOk() (*FullTextEntities, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *Tweet) SetEntities(v FullTextEntities)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *Tweet) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetPublicMetrics

`func (o *Tweet) GetPublicMetrics() TweetPublicMetrics`

GetPublicMetrics returns the PublicMetrics field if non-nil, zero value otherwise.

### GetPublicMetricsOk

`func (o *Tweet) GetPublicMetricsOk() (*TweetPublicMetrics, bool)`

GetPublicMetricsOk returns a tuple with the PublicMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicMetrics

`func (o *Tweet) SetPublicMetrics(v TweetPublicMetrics)`

SetPublicMetrics sets PublicMetrics field to given value.

### HasPublicMetrics

`func (o *Tweet) HasPublicMetrics() bool`

HasPublicMetrics returns a boolean if a field has been set.

### GetPossiblySensitive

`func (o *Tweet) GetPossiblySensitive() bool`

GetPossiblySensitive returns the PossiblySensitive field if non-nil, zero value otherwise.

### GetPossiblySensitiveOk

`func (o *Tweet) GetPossiblySensitiveOk() (*bool, bool)`

GetPossiblySensitiveOk returns a tuple with the PossiblySensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossiblySensitive

`func (o *Tweet) SetPossiblySensitive(v bool)`

SetPossiblySensitive sets PossiblySensitive field to given value.

### HasPossiblySensitive

`func (o *Tweet) HasPossiblySensitive() bool`

HasPossiblySensitive returns a boolean if a field has been set.

### GetLang

`func (o *Tweet) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *Tweet) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *Tweet) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *Tweet) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetSource

`func (o *Tweet) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Tweet) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Tweet) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Tweet) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetNonPublicMetrics

`func (o *Tweet) GetNonPublicMetrics() TweetNonPublicMetrics`

GetNonPublicMetrics returns the NonPublicMetrics field if non-nil, zero value otherwise.

### GetNonPublicMetricsOk

`func (o *Tweet) GetNonPublicMetricsOk() (*TweetNonPublicMetrics, bool)`

GetNonPublicMetricsOk returns a tuple with the NonPublicMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonPublicMetrics

`func (o *Tweet) SetNonPublicMetrics(v TweetNonPublicMetrics)`

SetNonPublicMetrics sets NonPublicMetrics field to given value.

### HasNonPublicMetrics

`func (o *Tweet) HasNonPublicMetrics() bool`

HasNonPublicMetrics returns a boolean if a field has been set.

### GetPromotedMetrics

`func (o *Tweet) GetPromotedMetrics() TweetPromotedMetrics`

GetPromotedMetrics returns the PromotedMetrics field if non-nil, zero value otherwise.

### GetPromotedMetricsOk

`func (o *Tweet) GetPromotedMetricsOk() (*TweetPromotedMetrics, bool)`

GetPromotedMetricsOk returns a tuple with the PromotedMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotedMetrics

`func (o *Tweet) SetPromotedMetrics(v TweetPromotedMetrics)`

SetPromotedMetrics sets PromotedMetrics field to given value.

### HasPromotedMetrics

`func (o *Tweet) HasPromotedMetrics() bool`

HasPromotedMetrics returns a boolean if a field has been set.

### GetOrganicMetrics

`func (o *Tweet) GetOrganicMetrics() TweetOrganicMetrics`

GetOrganicMetrics returns the OrganicMetrics field if non-nil, zero value otherwise.

### GetOrganicMetricsOk

`func (o *Tweet) GetOrganicMetricsOk() (*TweetOrganicMetrics, bool)`

GetOrganicMetricsOk returns a tuple with the OrganicMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganicMetrics

`func (o *Tweet) SetOrganicMetrics(v TweetOrganicMetrics)`

SetOrganicMetrics sets OrganicMetrics field to given value.

### HasOrganicMetrics

`func (o *Tweet) HasOrganicMetrics() bool`

HasOrganicMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


