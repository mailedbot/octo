package parser

import "encoding/json"

// ParseDataToType parses data to a target type using JSON marshalling and unmarshalling.
// The data is first marshalled to JSON and then unmarshalled to the target type.
// If an error occurs during marshalling or unmarshalling, it is returned.
func ParseDataToType(data interface{}, targetType interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, targetType)
	if err != nil {
		return err
	}

	return nil
}
