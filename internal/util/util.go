package util

import (
	"encoding/json"
	"log"
)

func ToJSON(src interface{}) string {
	str, err := json.Marshal(src) 
	if err != nil {
		log.Default().Printf("Error formatting JSON: %v", err)
		return ""
	}
	return string(str)
}