/*
Tweets and Users

API Reference — Labs v2

API version: 2.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package twitterapi

import (
	"encoding/json"
)

// TweetGeo The location tagged on the Tweet, if the user provided one.
type TweetGeo struct {
	Coordinates *Point `json:"coordinates,omitempty"`
	// The identifier for this place
	PlaceId *string `json:"place_id,omitempty"`
}

// NewTweetGeo instantiates a new TweetGeo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTweetGeo() *TweetGeo {
	this := TweetGeo{}
	return &this
}

// NewTweetGeoWithDefaults instantiates a new TweetGeo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTweetGeoWithDefaults() *TweetGeo {
	this := TweetGeo{}
	return &this
}

// GetCoordinates returns the Coordinates field value if set, zero value otherwise.
func (o *TweetGeo) GetCoordinates() Point {
	if o == nil || o.Coordinates == nil {
		var ret Point
		return ret
	}
	return *o.Coordinates
}

// GetCoordinatesOk returns a tuple with the Coordinates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetGeo) GetCoordinatesOk() (*Point, bool) {
	if o == nil || o.Coordinates == nil {
		return nil, false
	}
	return o.Coordinates, true
}

// HasCoordinates returns a boolean if a field has been set.
func (o *TweetGeo) HasCoordinates() bool {
	if o != nil && o.Coordinates != nil {
		return true
	}

	return false
}

// SetCoordinates gets a reference to the given Point and assigns it to the Coordinates field.
func (o *TweetGeo) SetCoordinates(v Point) {
	o.Coordinates = &v
}

// GetPlaceId returns the PlaceId field value if set, zero value otherwise.
func (o *TweetGeo) GetPlaceId() string {
	if o == nil || o.PlaceId == nil {
		var ret string
		return ret
	}
	return *o.PlaceId
}

// GetPlaceIdOk returns a tuple with the PlaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetGeo) GetPlaceIdOk() (*string, bool) {
	if o == nil || o.PlaceId == nil {
		return nil, false
	}
	return o.PlaceId, true
}

// HasPlaceId returns a boolean if a field has been set.
func (o *TweetGeo) HasPlaceId() bool {
	if o != nil && o.PlaceId != nil {
		return true
	}

	return false
}

// SetPlaceId gets a reference to the given string and assigns it to the PlaceId field.
func (o *TweetGeo) SetPlaceId(v string) {
	o.PlaceId = &v
}

func (o TweetGeo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Coordinates != nil {
		toSerialize["coordinates"] = o.Coordinates
	}
	if o.PlaceId != nil {
		toSerialize["place_id"] = o.PlaceId
	}
	return json.Marshal(toSerialize)
}

type NullableTweetGeo struct {
	value *TweetGeo
	isSet bool
}

func (v NullableTweetGeo) Get() *TweetGeo {
	return v.value
}

func (v *NullableTweetGeo) Set(val *TweetGeo) {
	v.value = val
	v.isSet = true
}

func (v NullableTweetGeo) IsSet() bool {
	return v.isSet
}

func (v *NullableTweetGeo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTweetGeo(val *TweetGeo) *NullableTweetGeo {
	return &NullableTweetGeo{value: val, isSet: true}
}

func (v NullableTweetGeo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTweetGeo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


