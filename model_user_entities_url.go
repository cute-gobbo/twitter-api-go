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

// UserEntitiesUrl Expanded details for the URL specified in the user's profile, with start and end indices.
type UserEntitiesUrl struct {
	Urls []UrlEntity `json:"urls,omitempty"`
}

// NewUserEntitiesUrl instantiates a new UserEntitiesUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserEntitiesUrl() *UserEntitiesUrl {
	this := UserEntitiesUrl{}
	return &this
}

// NewUserEntitiesUrlWithDefaults instantiates a new UserEntitiesUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserEntitiesUrlWithDefaults() *UserEntitiesUrl {
	this := UserEntitiesUrl{}
	return &this
}

// GetUrls returns the Urls field value if set, zero value otherwise.
func (o *UserEntitiesUrl) GetUrls() []UrlEntity {
	if o == nil || o.Urls == nil {
		var ret []UrlEntity
		return ret
	}
	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserEntitiesUrl) GetUrlsOk() ([]UrlEntity, bool) {
	if o == nil || o.Urls == nil {
		return nil, false
	}
	return o.Urls, true
}

// HasUrls returns a boolean if a field has been set.
func (o *UserEntitiesUrl) HasUrls() bool {
	if o != nil && o.Urls != nil {
		return true
	}

	return false
}

// SetUrls gets a reference to the given []UrlEntity and assigns it to the Urls field.
func (o *UserEntitiesUrl) SetUrls(v []UrlEntity) {
	o.Urls = v
}

func (o UserEntitiesUrl) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Urls != nil {
		toSerialize["urls"] = o.Urls
	}
	return json.Marshal(toSerialize)
}

type NullableUserEntitiesUrl struct {
	value *UserEntitiesUrl
	isSet bool
}

func (v NullableUserEntitiesUrl) Get() *UserEntitiesUrl {
	return v.value
}

func (v *NullableUserEntitiesUrl) Set(val *UserEntitiesUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableUserEntitiesUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableUserEntitiesUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserEntitiesUrl(val *UserEntitiesUrl) *NullableUserEntitiesUrl {
	return &NullableUserEntitiesUrl{value: val, isSet: true}
}

func (v NullableUserEntitiesUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserEntitiesUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


