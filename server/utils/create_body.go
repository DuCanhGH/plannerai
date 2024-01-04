package utils

import "log"

func CreateBody(message string) map[string]interface{} {
	log.Printf("creating body with message %s", message)

	return map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []interface{}{map[string]interface{}{
			"role":    "system",
			"content": message,
		}},
	}
}
