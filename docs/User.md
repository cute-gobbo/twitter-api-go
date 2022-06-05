# User

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Creation time of this user. | [optional] [default to null]
**Name** | **string** | The friendly name of this user, as shown on their profile. | [default to null]
**Username** | **string** |  | [default to null]
**Protected** | **bool** | Indicates if this user has chosen to protect their Tweets (in other words, if this user&#x27;s Tweets are private). | [optional] [default to null]
**Verified** | **bool** | Indicate if this user is a verified Twitter User. | [optional] [default to null]
**Withheld** | [***UserWithheld**](UserWithheld.md) |  | [optional] [default to null]
**ProfileImageUrl** | **string** | The URL to the profile image for this user. | [optional] [default to null]
**Location** | **string** | The location specified in the user&#x27;s profile, if the user provided one. As this is a freeform value, it may not indicate a valid location, but it may be fuzzily evaluated when performing searches with location queries. | [optional] [default to null]
**Url** | **string** | The URL specified in the user&#x27;s profile. | [optional] [default to null]
**Description** | **string** | The text of this user&#x27;s profile description (also known as bio), if the user provided one. | [optional] [default to null]
**Entities** | [***UserEntities**](User_entities.md) |  | [optional] [default to null]
**PinnedTweetId** | **string** |  | [optional] [default to null]
**PublicMetrics** | [***UserPublicMetrics**](User_public_metrics.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

