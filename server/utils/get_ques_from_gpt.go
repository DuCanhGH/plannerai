package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
)

var endpoint = "https://api.openai.com/v1/chat/completions"

func GetQuesGpt(userGoal string) ([]string, error) {
	api_key := os.Getenv("OPENAI_API_KEY")

	body_req := CreateBody(
		"You are planner assistant and user has a goal is : \" " + userGoal + "\". Can you give some questions for user to get more information about his/her goal (Just send list of questions)",
	)

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(body_req)

	client := &http.Client{}
	req, err := http.NewRequest("POST", endpoint, buffer)

	if err != nil {
		return []string{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api_key)

	resp, err := client.Do(req)

	if err != nil {
		return []string{}, err
	}

	bodyBytes, _ := io.ReadAll(resp.Body)
	body_resp := make(map[string]interface{})
	_ = json.Unmarshal(bodyBytes, &body_resp)

	result := body_resp["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	return strings.Split(result, "\n"), nil
}
