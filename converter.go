package converter

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Convert converts an interface{} to a target struct
func Convert(input interface{}, output interface{}) error {

	// Ensure output is a pointer to a struct
	val := reflect.ValueOf(output)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("output must be a pointer to a struct")
	}

	// Marshal the input to JSON, then unmarshal into the output
	data, err := json.Marshal(input)
	if err != nil {
		return fmt.Errorf("failed to marshal input: %v", err)
	}

	err = json.Unmarshal(data, output)
	if err != nil {
		return fmt.Errorf("failed to unmarshal to output struct: %v", err)
	}

	return nil
}
