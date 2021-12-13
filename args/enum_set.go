package args

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"strings"
)

type EnumValueSet struct {
	Enum       []string
	Default    []string
	slice      []string
	hasBeenSet bool
}

// Set appends the enum value to the list of values
func (s *EnumValueSet) Set(value string) error {
	if !isEnumValue(s.Enum, value) {
		return fmt.Errorf("allowed values are %s", strings.Join(s.Enum, ", "))
	}

	if !s.hasBeenSet {
		s.slice = []string{}
		s.hasBeenSet = true
	}

	s.slice = append(s.slice, value)

	return nil
}

func (s *EnumValueSet) Value() []string {
	if s.hasBeenSet {
		return s.slice
	}
	return s.Default
}

// String returns a readable representation of this value (for usage defaults)
func (s *EnumValueSet) String() string {
	return fmt.Sprintf("%s", s.slice)
}

func GetEnumValueSet(c *cli.Context, name string) []string {
	generic, ok := c.Generic(name).(*EnumValueSet)
	if !ok {
		return nil
	}
	return generic.Value()
}
