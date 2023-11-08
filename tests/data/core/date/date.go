// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"
import "fmt"
import "github.com/atombender/go-jsonschema/pkg/types"
import yaml "gopkg.in/yaml.v3"

type Date struct {
	// MyObject corresponds to the JSON schema field "myObject".
	MyObject *DateMyObject `json:"myObject,omitempty" yaml:"myObject,omitempty" mapstructure:"myObject,omitempty"`
}

type DateMyObject struct {
	// MyDate corresponds to the JSON schema field "myDate".
	MyDate types.SerializableDate `json:"myDate" yaml:"myDate" mapstructure:"myDate"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DateMyObject) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["myDate"]; raw != nil && !ok {
		return fmt.Errorf("field myDate in DateMyObject: required")
	}
	type Plain DateMyObject
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	*j = DateMyObject(plain)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *DateMyObject) UnmarshalYAML(value *yaml.Node) error {
	var raw map[string]interface{}
	if err := value.Decode(&raw); err != nil {
		return err
	}
	if v, ok := raw["myDate"]; !ok || v == nil {
		return fmt.Errorf("field myDate in DateMyObject: required")
	}
	type Plain DateMyObject
	var plain Plain
	if err := value.Decode(&plain); err != nil {
		return err
	}
	*j = DateMyObject(plain)
	return nil
}
