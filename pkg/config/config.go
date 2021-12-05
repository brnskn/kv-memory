// This package loads config keys as a map and handles default values
package config

import (
	"os"
	"strconv"
	"strings"
)

var config map[string]string

// Load configs from os environment
func Load() {
	config = make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		config[pair[0]] = pair[1]
	}
}

// Gets the value of the given key as string. If the key's value is
// not set returns the given default value.
func Get(key string, default_value string) string {
	value, found := config[key]
	if !found {
		return default_value
	}
	return value
}

// Gets value of given key as integer. If the key's value is
// not set returns the given default value.
func GetInt(key string, default_value int) int {
	value := Get(key, "")
	if value == "" {
		return default_value
	}
	int_value, err := strconv.Atoi(value)
	if err != nil {
		return default_value
	}
	return int_value
}
