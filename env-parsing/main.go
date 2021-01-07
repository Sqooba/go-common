package env_parsing

import (
	"fmt"
	"strings"
)

// StringFlags extends the flags feature with a list of strings.
// shamelessly inspired from http://blog.ralch.com/tutorial/golang-custom-flags/
type StringFlags []string

func (i *StringFlags) String() string {
	return strings.Join(*i, ", ")
}

func (i *StringFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

// MapSemicolonSepValDecoder is a custom decoder for https://github.com/kelseyhightower/envconfig
type MapSemicolonSepValDecoder map[string]string

// Decode decode the environment variable values like key1=value1;key2=value2;... into a map
// { "key1": "value1", "key2": "value2", ... }
func (m *MapSemicolonSepValDecoder) Decode(value string) error {
	mapCSV := map[string]string{}
	pairs := strings.Split(value, ";")
	for _, pair := range pairs {
		kvpair := strings.Split(pair, "=")
		if len(kvpair) != 2 {
			return fmt.Errorf("invalid map item: %q", pair)
		}
		mapCSV[kvpair[0]] = kvpair[1]

	}
	*m = mapCSV
	return nil
}
