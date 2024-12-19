/*
Svix API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// StatisticsPeriod Period length for a statistics data point.
type StatisticsPeriod string

// List of StatisticsPeriod
const (
	STATISTICSPERIOD_ONE_DAY StatisticsPeriod = "OneDay"
	STATISTICSPERIOD_FIVE_MINUTES StatisticsPeriod = "FiveMinutes"
)

// All allowed values of StatisticsPeriod enum
var AllowedStatisticsPeriodEnumValues = []StatisticsPeriod{
	"OneDay",
	"FiveMinutes",
}

func (v *StatisticsPeriod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StatisticsPeriod(value)
	for _, existing := range AllowedStatisticsPeriodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StatisticsPeriod", value)
}

// NewStatisticsPeriodFromValue returns a pointer to a valid StatisticsPeriod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStatisticsPeriodFromValue(v string) (*StatisticsPeriod, error) {
	ev := StatisticsPeriod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StatisticsPeriod: valid values are %v", v, AllowedStatisticsPeriodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StatisticsPeriod) IsValid() bool {
	for _, existing := range AllowedStatisticsPeriodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StatisticsPeriod value
func (v StatisticsPeriod) Ptr() *StatisticsPeriod {
	return &v
}

type NullableStatisticsPeriod struct {
	value *StatisticsPeriod
	isSet bool
}

func (v NullableStatisticsPeriod) Get() *StatisticsPeriod {
	return v.value
}

func (v *NullableStatisticsPeriod) Set(val *StatisticsPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableStatisticsPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableStatisticsPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatisticsPeriod(val *StatisticsPeriod) *NullableStatisticsPeriod {
	return &NullableStatisticsPeriod{value: val, isSet: true}
}

func (v NullableStatisticsPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatisticsPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

