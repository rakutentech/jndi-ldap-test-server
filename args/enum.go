package args

import (
	"fmt"
	"strings"
)

// EnumValues creates an Enum Value with the first argument being the default value
func EnumValues(defaultValue string, values ...string) *EnumValue {
	return &EnumValue{
		Enum: append([]string{defaultValue}, values...),
		Default: defaultValue,
	}
}

type EnumValue struct {
	Enum     []string
	Default  string
	selected string
}

func (e *EnumValue) Set(value string) error {
	for _, enum := range e.Enum {
		if enum == value {
			e.selected = value
			return nil
		}
	}

	return fmt.Errorf("allowed values are %s", strings.Join(e.Enum, ", "))
}

func (e EnumValue) String() string {
	if e.selected == "" {
		return e.Default
	}
	return e.selected
}
