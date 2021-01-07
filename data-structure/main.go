package data_structure

import (
	"fmt"
	"reflect"
)

// GetKeysFromMap returns the string representations of the keys of a map.
func GetKeysFromMap(m interface{}) []string {
	v := reflect.ValueOf(m)
	switch v.Kind() {
	case reflect.Map:
		keys := make([]string, 0, len(v.MapKeys()))
		for _, key := range v.MapKeys() {
			keys = append(keys, key.String())
		}
		return keys
	default:
		panic(fmt.Sprintf("unexpected type: %T", v.Interface()))
	}
}
