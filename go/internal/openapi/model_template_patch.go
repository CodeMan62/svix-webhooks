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

// checks if the TemplatePatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplatePatch{}

// TemplatePatch struct for TemplatePatch
type TemplatePatch struct {
	Description *string `json:"description,omitempty"`
	FeatureFlag NullableString `json:"featureFlag,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-_.]+$"`
	FilterTypes []string `json:"filterTypes,omitempty"`
	Instructions *string `json:"instructions,omitempty"`
	InstructionsLink NullableString `json:"instructionsLink,omitempty"`
	Kind *ConnectorKind `json:"kind,omitempty"`
	Logo *string `json:"logo,omitempty"`
	Name *string `json:"name,omitempty"`
	Transformation *string `json:"transformation,omitempty"`
}

// NewTemplatePatch instantiates a new TemplatePatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplatePatch() *TemplatePatch {
	this := TemplatePatch{}
	return &this
}

// NewTemplatePatchWithDefaults instantiates a new TemplatePatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplatePatchWithDefaults() *TemplatePatch {
	this := TemplatePatch{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TemplatePatch) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatePatch) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TemplatePatch) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TemplatePatch) SetDescription(v string) {
	o.Description = &v
}

// GetFeatureFlag returns the FeatureFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplatePatch) GetFeatureFlag() string {
	if o == nil || IsNil(o.FeatureFlag.Get()) {
		var ret string
		return ret
	}
	return *o.FeatureFlag.Get()
}

// GetFeatureFlagOk returns a tuple with the FeatureFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatePatch) GetFeatureFlagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FeatureFlag.Get(), o.FeatureFlag.IsSet()
}

// HasFeatureFlag returns a boolean if a field has been set.
func (o *TemplatePatch) HasFeatureFlag() bool {
	if o != nil && o.FeatureFlag.IsSet() {
		return true
	}

	return false
}

// SetFeatureFlag gets a reference to the given NullableString and assigns it to the FeatureFlag field.
func (o *TemplatePatch) SetFeatureFlag(v string) {
	o.FeatureFlag.Set(&v)
}
// SetFeatureFlagNil sets the value for FeatureFlag to be an explicit nil
func (o *TemplatePatch) SetFeatureFlagNil() {
	o.FeatureFlag.Set(nil)
}

// UnsetFeatureFlag ensures that no value is present for FeatureFlag, not even an explicit nil
func (o *TemplatePatch) UnsetFeatureFlag() {
	o.FeatureFlag.Unset()
}

// GetFilterTypes returns the FilterTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplatePatch) GetFilterTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.FilterTypes
}

// GetFilterTypesOk returns a tuple with the FilterTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatePatch) GetFilterTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.FilterTypes) {
		return nil, false
	}
	return o.FilterTypes, true
}

// HasFilterTypes returns a boolean if a field has been set.
func (o *TemplatePatch) HasFilterTypes() bool {
	if o != nil && !IsNil(o.FilterTypes) {
		return true
	}

	return false
}

// SetFilterTypes gets a reference to the given []string and assigns it to the FilterTypes field.
func (o *TemplatePatch) SetFilterTypes(v []string) {
	o.FilterTypes = v
}

// GetInstructions returns the Instructions field value if set, zero value otherwise.
func (o *TemplatePatch) GetInstructions() string {
	if o == nil || IsNil(o.Instructions) {
		var ret string
		return ret
	}
	return *o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatePatch) GetInstructionsOk() (*string, bool) {
	if o == nil || IsNil(o.Instructions) {
		return nil, false
	}
	return o.Instructions, true
}

// HasInstructions returns a boolean if a field has been set.
func (o *TemplatePatch) HasInstructions() bool {
	if o != nil && !IsNil(o.Instructions) {
		return true
	}

	return false
}

// SetInstructions gets a reference to the given string and assigns it to the Instructions field.
func (o *TemplatePatch) SetInstructions(v string) {
	o.Instructions = &v
}

// GetInstructionsLink returns the InstructionsLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplatePatch) GetInstructionsLink() string {
	if o == nil || IsNil(o.InstructionsLink.Get()) {
		var ret string
		return ret
	}
	return *o.InstructionsLink.Get()
}

// GetInstructionsLinkOk returns a tuple with the InstructionsLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatePatch) GetInstructionsLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstructionsLink.Get(), o.InstructionsLink.IsSet()
}

// HasInstructionsLink returns a boolean if a field has been set.
func (o *TemplatePatch) HasInstructionsLink() bool {
	if o != nil && o.InstructionsLink.IsSet() {
		return true
	}

	return false
}

// SetInstructionsLink gets a reference to the given NullableString and assigns it to the InstructionsLink field.
func (o *TemplatePatch) SetInstructionsLink(v string) {
	o.InstructionsLink.Set(&v)
}
// SetInstructionsLinkNil sets the value for InstructionsLink to be an explicit nil
func (o *TemplatePatch) SetInstructionsLinkNil() {
	o.InstructionsLink.Set(nil)
}

// UnsetInstructionsLink ensures that no value is present for InstructionsLink, not even an explicit nil
func (o *TemplatePatch) UnsetInstructionsLink() {
	o.InstructionsLink.Unset()
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *TemplatePatch) GetKind() ConnectorKind {
	if o == nil || IsNil(o.Kind) {
		var ret ConnectorKind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatePatch) GetKindOk() (*ConnectorKind, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *TemplatePatch) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given ConnectorKind and assigns it to the Kind field.
func (o *TemplatePatch) SetKind(v ConnectorKind) {
	o.Kind = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *TemplatePatch) GetLogo() string {
	if o == nil || IsNil(o.Logo) {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatePatch) GetLogoOk() (*string, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *TemplatePatch) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *TemplatePatch) SetLogo(v string) {
	o.Logo = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TemplatePatch) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatePatch) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TemplatePatch) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TemplatePatch) SetName(v string) {
	o.Name = &v
}

// GetTransformation returns the Transformation field value if set, zero value otherwise.
func (o *TemplatePatch) GetTransformation() string {
	if o == nil || IsNil(o.Transformation) {
		var ret string
		return ret
	}
	return *o.Transformation
}

// GetTransformationOk returns a tuple with the Transformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatePatch) GetTransformationOk() (*string, bool) {
	if o == nil || IsNil(o.Transformation) {
		return nil, false
	}
	return o.Transformation, true
}

// HasTransformation returns a boolean if a field has been set.
func (o *TemplatePatch) HasTransformation() bool {
	if o != nil && !IsNil(o.Transformation) {
		return true
	}

	return false
}

// SetTransformation gets a reference to the given string and assigns it to the Transformation field.
func (o *TemplatePatch) SetTransformation(v string) {
	o.Transformation = &v
}

func (o TemplatePatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplatePatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.FeatureFlag.IsSet() {
		toSerialize["featureFlag"] = o.FeatureFlag.Get()
	}
	if o.FilterTypes != nil {
		toSerialize["filterTypes"] = o.FilterTypes
	}
	if !IsNil(o.Instructions) {
		toSerialize["instructions"] = o.Instructions
	}
	if o.InstructionsLink.IsSet() {
		toSerialize["instructionsLink"] = o.InstructionsLink.Get()
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Transformation) {
		toSerialize["transformation"] = o.Transformation
	}
	return toSerialize, nil
}

type NullableTemplatePatch struct {
	value *TemplatePatch
	isSet bool
}

func (v NullableTemplatePatch) Get() *TemplatePatch {
	return v.value
}

func (v *NullableTemplatePatch) Set(val *TemplatePatch) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplatePatch) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplatePatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplatePatch(val *TemplatePatch) *NullableTemplatePatch {
	return &NullableTemplatePatch{value: val, isSet: true}
}

func (v NullableTemplatePatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplatePatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


