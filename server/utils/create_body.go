package utils

func CreateBody(message string, temperature float32) map[string]interface{} {

	return map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []interface{}{map[string]interface{}{
			"role":    "system",
			"content": message,
		}},
		"temperature": temperature,
	}
}
