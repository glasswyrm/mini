package mini

import (
	"errors"
	"strings"
)

// we might want value to be a string,
// then each 'set' function converts the respective type to string
func set(values *map[string]interface{}, key string, value interface{}) error {
	if len(key) == 0 || values == nil {
		return nil
	}

	key = strings.ToLower(key)
	val, ok := (*values)[key]

	if ok {
		switch val.(type) {
		case []interface{}:
			return nil
		default:
			(*values)[key] = value
		}
	}

	return errors.New("key not found")
}

/*
SetStringForSection attempts to set the specified key in sectionName to the value string.
If the section name is not found an error is returned.
*/
func (config *Config) SetStringForSection(sectionName string, key string, value string) error {
	section := config.sectionForName(sectionName)

	if section != nil {
		return set(&section.values, key, value)
	}
	return errors.New("section not found")
}

/*
SetBooleanForSection attempts to set the specified key in sectionName to the value boolean.
If the section name is not found an error is returned.
*/
func (config *Config) SetBooleanForSection(sectionName string, key string, value bool) error {
	section := config.sectionForName(sectionName)

	if section != nil {
		return set(&section.values, key, value)
	}
	return errors.New("section not found")
}
