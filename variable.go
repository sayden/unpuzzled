package cli

import (
	"flag"
	"regexp"
	"strings"
)

type Variable interface {
	GetName() string
	setFlag(*flag.FlagSet)
	setEnv() (interface{}, bool)
	getFlagValue(*flag.FlagSet) (interface{}, bool)
	IsRequired() bool
}

var osToEnvReplaceRegexp = regexp.MustCompile(`[\.\-]`)

func convertNameToOS(name string) string {
	a := osToEnvReplaceRegexp.ReplaceAllString(strings.ToUpper(name), "_")
	return a
}