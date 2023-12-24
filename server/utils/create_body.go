package utils

func CreateBody(data string) map[string]interface{} {
	return map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []interface{}{map[string]interface{}{
			"role":    "system",
			"content": data,
		}},
	}
}
