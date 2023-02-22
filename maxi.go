package mini

// we might want value to be a string,
// then each 'set' function converts the respective type to string?
func set(values *map[string]interface{}, key string, value interface{}) {
	if len(key) != 0 && values != nil {
		(*values)[key] = value
	}
}

/*
SetStringForSection attempts to set the specified key in sectionName to the value string.
If the section name is not found a new section is created.
*/
func (config *Config) SetStringForSection(sectionName string, key string, value string) {
	section := config.sectionForName(sectionName)

	if section == nil {
		// need to handle populating a potentially new section
		// pretty sure there's a better way
		section = new(configSection)
		section.name = sectionName
		section.values = make(map[string]interface{})
		config.sections[section.name] = section
	}
	set(&section.values, key, value)
}

/*
SetBooleanForSection attempts to set the specified key in sectionName to the value boolean.
If the section name is not found a new section is created.
*/
func (config *Config) SetBooleanForSection(sectionName string, key string, value bool) {
	section := config.sectionForName(sectionName)

	if section == nil {
		// need to handle populating a potentially new section
		section = new(configSection)
		section.name = sectionName
		section.values = make(map[string]interface{})
		config.sections[section.name] = section
	}
	set(&section.values, key, value)
}
