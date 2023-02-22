package mini

import (
	"errors"
	"fmt"
	"os"
	"sort"
)

// we might want value to be a string,
// then each 'set' function converts the respective type to string?
func set(values *map[string]interface{}, key string, value interface{}) error {
	if len(key) == 0 || values == nil {
		return nil
	}
	(*values)[key] = value
	return nil

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

func (config *Config) WriteSectionsToFile(fileName string) error {
	// two major considerations-- How to handle config outside a section
	// and do we want to preserve comments? we'll have to change the ingestion
	// right now both those things are lost

	//output := os.Stdout
	output, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	// sort section names before processing
	nameList := config.SectionNames()
	sort.Strings(nameList)
	for _, sectionName := range nameList {
		_, err := fmt.Fprint(output, "[", sectionName, "]\n")
		if err != nil {
			return err
		}
		keyList := config.KeysForSection(sectionName)
		sort.Strings(keyList)
		for _, keyName := range keyList {
			// again, sort section names
			_, err = fmt.Fprint(output, keyName, "=", config.StringFromSection(sectionName, keyName, ""), "\n")
			if err != nil {
				return err
			}
		}
		fmt.Fprintln(output)
	}
	return output.Close()
}
