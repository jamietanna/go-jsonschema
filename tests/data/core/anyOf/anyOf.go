// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"
import "errors"
import "fmt"
import yaml "gopkg.in/yaml.v3"

type AnyOf struct {
	// Configurations corresponds to the JSON schema field "configurations".
	Configurations []AnyOfConfigurationsElem `json:"configurations,omitempty" yaml:"configurations,omitempty" mapstructure:"configurations,omitempty"`

	// Flags corresponds to the JSON schema field "flags".
	Flags interface{} `json:"flags,omitempty" yaml:"flags,omitempty" mapstructure:"flags,omitempty"`
}

type AnyOfConfigurationsElem struct {
	// Bar corresponds to the JSON schema field "bar".
	Bar float64 `json:"bar" yaml:"bar" mapstructure:"bar"`

	// Baz corresponds to the JSON schema field "baz".
	Baz *bool `json:"baz,omitempty" yaml:"baz,omitempty" mapstructure:"baz,omitempty"`

	// Foo corresponds to the JSON schema field "foo".
	Foo string `json:"foo" yaml:"foo" mapstructure:"foo"`
}

type AnyOfConfigurationsElem_0 struct {
	// Foo corresponds to the JSON schema field "foo".
	Foo string `json:"foo" yaml:"foo" mapstructure:"foo"`
}

type AnyOfConfigurationsElem_1 struct {
	// Bar corresponds to the JSON schema field "bar".
	Bar float64 `json:"bar" yaml:"bar" mapstructure:"bar"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AnyOfConfigurationsElem_1) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if v, ok := raw["bar"]; !ok || v == nil {
		return fmt.Errorf("field bar in AnyOfConfigurationsElem_1: required")
	}
	type Plain AnyOfConfigurationsElem_1
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	*j = AnyOfConfigurationsElem_1(plain)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *AnyOfConfigurationsElem_1) UnmarshalYAML(value *yaml.Node) error {
	var raw map[string]interface{}
	if err := value.Decode(&raw); err != nil {
		return err
	}
	if v, ok := raw["bar"]; !ok || v == nil {
		return fmt.Errorf("field bar in AnyOfConfigurationsElem_1: required")
	}
	type Plain AnyOfConfigurationsElem_1
	var plain Plain
	if err := value.Decode(&plain); err != nil {
		return err
	}
	*j = AnyOfConfigurationsElem_1(plain)
	return nil
}

type AnyOfConfigurationsElem_2 struct {
	// Baz corresponds to the JSON schema field "baz".
	Baz *bool `json:"baz,omitempty" yaml:"baz,omitempty" mapstructure:"baz,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AnyOfConfigurationsElem_2) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	type Plain AnyOfConfigurationsElem_2
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	*j = AnyOfConfigurationsElem_2(plain)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *AnyOfConfigurationsElem_2) UnmarshalYAML(value *yaml.Node) error {
	var raw map[string]interface{}
	if err := value.Decode(&raw); err != nil {
		return err
	}
	type Plain AnyOfConfigurationsElem_2
	var plain Plain
	if err := value.Decode(&plain); err != nil {
		return err
	}
	*j = AnyOfConfigurationsElem_2(plain)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *AnyOfConfigurationsElem_0) UnmarshalYAML(value *yaml.Node) error {
	var raw map[string]interface{}
	if err := value.Decode(&raw); err != nil {
		return err
	}
	if v, ok := raw["foo"]; !ok || v == nil {
		return fmt.Errorf("field foo in AnyOfConfigurationsElem_0: required")
	}
	type Plain AnyOfConfigurationsElem_0
	var plain Plain
	if err := value.Decode(&plain); err != nil {
		return err
	}
	*j = AnyOfConfigurationsElem_0(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AnyOfConfigurationsElem) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	var anyOfConfigurationsElem_0 AnyOfConfigurationsElem_0
	var anyOfConfigurationsElem_1 AnyOfConfigurationsElem_1
	var anyOfConfigurationsElem_2 AnyOfConfigurationsElem_2
	var errs []error
	if err := anyOfConfigurationsElem_0.UnmarshalJSON(value); err != nil {
		errs = append(errs, err)
	}
	if err := anyOfConfigurationsElem_1.UnmarshalJSON(value); err != nil {
		errs = append(errs, err)
	}
	if err := anyOfConfigurationsElem_2.UnmarshalJSON(value); err != nil {
		errs = append(errs, err)
	}
	if len(errs) == 3 {
		return fmt.Errorf("all validators failed: %s", errors.Join(errs...))
	}
	type Plain AnyOfConfigurationsElem
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	*j = AnyOfConfigurationsElem(plain)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *AnyOfConfigurationsElem) UnmarshalYAML(value *yaml.Node) error {
	var raw map[string]interface{}
	if err := value.Decode(&raw); err != nil {
		return err
	}
	var anyOfConfigurationsElem_0 AnyOfConfigurationsElem_0
	var anyOfConfigurationsElem_1 AnyOfConfigurationsElem_1
	var anyOfConfigurationsElem_2 AnyOfConfigurationsElem_2
	var errs []error
	if err := anyOfConfigurationsElem_0.UnmarshalYAML(value); err != nil {
		errs = append(errs, err)
	}
	if err := anyOfConfigurationsElem_1.UnmarshalYAML(value); err != nil {
		errs = append(errs, err)
	}
	if err := anyOfConfigurationsElem_2.UnmarshalYAML(value); err != nil {
		errs = append(errs, err)
	}
	if len(errs) == 3 {
		return fmt.Errorf("all validators failed: %s", errors.Join(errs...))
	}
	type Plain AnyOfConfigurationsElem
	var plain Plain
	if err := value.Decode(&plain); err != nil {
		return err
	}
	*j = AnyOfConfigurationsElem(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AnyOfConfigurationsElem_0) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if v, ok := raw["foo"]; !ok || v == nil {
		return fmt.Errorf("field foo in AnyOfConfigurationsElem_0: required")
	}
	type Plain AnyOfConfigurationsElem_0
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	*j = AnyOfConfigurationsElem_0(plain)
	return nil
}
