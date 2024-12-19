/*
Svix API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RotatePollerTokenIn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RotatePollerTokenIn{}

// RotatePollerTokenIn struct for RotatePollerTokenIn
type RotatePollerTokenIn struct {
	// How long the token will be valid for, in seconds. Can be up to 31,536,000 seconds (1 year).
	Expiry *int64 `json:"expiry,omitempty"`
	// Updates the previous token's expiration, in seconds. If set to 0, the old token will immediately be revoked. Must be between 0 and 86,400 seconds (1 day).  Defaults to 300 seconds (5 minutes).
	OldTokenExpiry *int64 `json:"oldTokenExpiry,omitempty"`
}

// NewRotatePollerTokenIn instantiates a new RotatePollerTokenIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRotatePollerTokenIn() *RotatePollerTokenIn {
	this := RotatePollerTokenIn{}
	var oldTokenExpiry int64 = 300
	this.OldTokenExpiry = &oldTokenExpiry
	return &this
}

// NewRotatePollerTokenInWithDefaults instantiates a new RotatePollerTokenIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRotatePollerTokenInWithDefaults() *RotatePollerTokenIn {
	this := RotatePollerTokenIn{}
	var oldTokenExpiry int64 = 300
	this.OldTokenExpiry = &oldTokenExpiry
	return &this
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *RotatePollerTokenIn) GetExpiry() int64 {
	if o == nil || IsNil(o.Expiry) {
		var ret int64
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatePollerTokenIn) GetExpiryOk() (*int64, bool) {
	if o == nil || IsNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *RotatePollerTokenIn) HasExpiry() bool {
	if o != nil && !IsNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given int64 and assigns it to the Expiry field.
func (o *RotatePollerTokenIn) SetExpiry(v int64) {
	o.Expiry = &v
}

// GetOldTokenExpiry returns the OldTokenExpiry field value if set, zero value otherwise.
func (o *RotatePollerTokenIn) GetOldTokenExpiry() int64 {
	if o == nil || IsNil(o.OldTokenExpiry) {
		var ret int64
		return ret
	}
	return *o.OldTokenExpiry
}

// GetOldTokenExpiryOk returns a tuple with the OldTokenExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatePollerTokenIn) GetOldTokenExpiryOk() (*int64, bool) {
	if o == nil || IsNil(o.OldTokenExpiry) {
		return nil, false
	}
	return o.OldTokenExpiry, true
}

// HasOldTokenExpiry returns a boolean if a field has been set.
func (o *RotatePollerTokenIn) HasOldTokenExpiry() bool {
	if o != nil && !IsNil(o.OldTokenExpiry) {
		return true
	}

	return false
}

// SetOldTokenExpiry gets a reference to the given int64 and assigns it to the OldTokenExpiry field.
func (o *RotatePollerTokenIn) SetOldTokenExpiry(v int64) {
	o.OldTokenExpiry = &v
}

func (o RotatePollerTokenIn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RotatePollerTokenIn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !IsNil(o.OldTokenExpiry) {
		toSerialize["oldTokenExpiry"] = o.OldTokenExpiry
	}
	return toSerialize, nil
}

type NullableRotatePollerTokenIn struct {
	value *RotatePollerTokenIn
	isSet bool
}

func (v NullableRotatePollerTokenIn) Get() *RotatePollerTokenIn {
	return v.value
}

func (v *NullableRotatePollerTokenIn) Set(val *RotatePollerTokenIn) {
	v.value = val
	v.isSet = true
}

func (v NullableRotatePollerTokenIn) IsSet() bool {
	return v.isSet
}

func (v *NullableRotatePollerTokenIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRotatePollerTokenIn(val *RotatePollerTokenIn) *NullableRotatePollerTokenIn {
	return &NullableRotatePollerTokenIn{value: val, isSet: true}
}

func (v NullableRotatePollerTokenIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRotatePollerTokenIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


