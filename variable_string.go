package unpuzzled

import (
	"flag"
	"os"
)

type StringVariable struct {
	Name        string
	Description string
	Default     string
	Required    bool
	Destination *string

	flagDestination *string
	envName         string
}

func (s *StringVariable) GetName() string {
	return s.Name
}

func (s *StringVariable) GetDescription() string {
	return s.Description
}

func (s *StringVariable) GetDestination() interface{} {
	return s.Destination
}

func (s *StringVariable) IsRequired() bool {
	return s.Required
}

func (s *StringVariable) apply(val interface{}) {
	if stringVal, ok := val.(string); ok {
		*s.Destination = stringVal
	}
}

func (s *StringVariable) setFlag(flagset *flag.FlagSet) {
	var destination string
	s.flagDestination = &destination
	flagset.StringVar(s.flagDestination, s.Name, s.Default, s.Description)
}

func (s *StringVariable) getFlagValue(set *flag.FlagSet) (interface{}, bool) {
	if *s.flagDestination == "" {
		return nil, false
	} else {
		return *s.flagDestination, true
	}
}

func (s *StringVariable) setEnv() (interface{}, bool) {
	s.envName = convertNameToOS(s.Name)
	if value, found := os.LookupEnv(s.envName); found {
		return value, true
	}
	return nil, false
}
