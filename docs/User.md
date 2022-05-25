# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of this User. This is returned as a string in order to avoid complications with languages and tools that cannot handle large integers. | 
**CreatedAt** | Pointer to **time.Time** | Creation time of this user. | [optional] 
**Name** | **string** | The friendly name of this user, as shown on their profile. | 
**Username** | **string** | The Twitter handle (screen name) of this user. | 
**Protected** | Pointer to **bool** | Indicates if this user has chosen to protect their Tweets (in other words, if this user&#39;s Tweets are private). | [optional] 
**Verified** | Pointer to **bool** | Indicate if this user is a verified Twitter User. | [optional] 
**Withheld** | Pointer to [**UserWithheld**](UserWithheld.md) |  | [optional] 
**ProfileImageUrl** | Pointer to **string** | The URL to the profile image for this user. | [optional] 
**Location** | Pointer to **string** | The location specified in the user&#39;s profile, if the user provided one. As this is a freeform value, it may not indicate a valid location, but it may be fuzzily evaluated when performing searches with location queries. | [optional] 
**Url** | Pointer to **string** | The URL specified in the user&#39;s profile. | [optional] 
**Description** | Pointer to **string** | The text of this user&#39;s profile description (also known as bio), if the user provided one. | [optional] 
**Entities** | Pointer to [**UserEntities**](UserEntities.md) |  | [optional] 
**PinnedTweetId** | Pointer to **string** | Unique identifier of this Tweet. This is returned as a string in order to avoid complications with languages and tools that cannot handle large integers. | [optional] 
**PublicMetrics** | Pointer to [**UserPublicMetrics**](UserPublicMetrics.md) |  | [optional] 

## Methods

### NewUser

`func NewUser(id string, name string, username string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *User) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *User) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetName

`func (o *User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User) SetName(v string)`

SetName sets Name field to given value.


### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetProtected

`func (o *User) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *User) GetProtectedOk() (*bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *User) SetProtected(v bool)`

SetProtected sets Protected field to given value.

### HasProtected

`func (o *User) HasProtected() bool`

HasProtected returns a boolean if a field has been set.

### GetVerified

`func (o *User) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *User) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *User) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *User) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetWithheld

`func (o *User) GetWithheld() UserWithheld`

GetWithheld returns the Withheld field if non-nil, zero value otherwise.

### GetWithheldOk

`func (o *User) GetWithheldOk() (*UserWithheld, bool)`

GetWithheldOk returns a tuple with the Withheld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithheld

`func (o *User) SetWithheld(v UserWithheld)`

SetWithheld sets Withheld field to given value.

### HasWithheld

`func (o *User) HasWithheld() bool`

HasWithheld returns a boolean if a field has been set.

### GetProfileImageUrl

`func (o *User) GetProfileImageUrl() string`

GetProfileImageUrl returns the ProfileImageUrl field if non-nil, zero value otherwise.

### GetProfileImageUrlOk

`func (o *User) GetProfileImageUrlOk() (*string, bool)`

GetProfileImageUrlOk returns a tuple with the ProfileImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileImageUrl

`func (o *User) SetProfileImageUrl(v string)`

SetProfileImageUrl sets ProfileImageUrl field to given value.

### HasProfileImageUrl

`func (o *User) HasProfileImageUrl() bool`

HasProfileImageUrl returns a boolean if a field has been set.

### GetLocation

`func (o *User) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *User) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *User) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *User) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetUrl

`func (o *User) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *User) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *User) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *User) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDescription

`func (o *User) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *User) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *User) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *User) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntities

`func (o *User) GetEntities() UserEntities`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *User) GetEntitiesOk() (*UserEntities, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *User) SetEntities(v UserEntities)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *User) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetPinnedTweetId

`func (o *User) GetPinnedTweetId() string`

GetPinnedTweetId returns the PinnedTweetId field if non-nil, zero value otherwise.

### GetPinnedTweetIdOk

`func (o *User) GetPinnedTweetIdOk() (*string, bool)`

GetPinnedTweetIdOk returns a tuple with the PinnedTweetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedTweetId

`func (o *User) SetPinnedTweetId(v string)`

SetPinnedTweetId sets PinnedTweetId field to given value.

### HasPinnedTweetId

`func (o *User) HasPinnedTweetId() bool`

HasPinnedTweetId returns a boolean if a field has been set.

### GetPublicMetrics

`func (o *User) GetPublicMetrics() UserPublicMetrics`

GetPublicMetrics returns the PublicMetrics field if non-nil, zero value otherwise.

### GetPublicMetricsOk

`func (o *User) GetPublicMetricsOk() (*UserPublicMetrics, bool)`

GetPublicMetricsOk returns a tuple with the PublicMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicMetrics

`func (o *User) SetPublicMetrics(v UserPublicMetrics)`

SetPublicMetrics sets PublicMetrics field to given value.

### HasPublicMetrics

`func (o *User) HasPublicMetrics() bool`

HasPublicMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


