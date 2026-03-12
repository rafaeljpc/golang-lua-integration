// Package util provides utility functions.
package util

import (
	"encoding/json"
	"log"
)

// ToJSON converts the given interface to a JSON string.
func ToJSON(src interface{}) string {
	str, err := json.Marshal(src)
	if err != nil {
		log.Default().Printf("Error formatting JSON: %v", err)

		return ""
	}

	return string(str)
}
