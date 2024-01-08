package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// func format_plan(plan string) string {

// }
func GetPlanGpt(usergoal string, time string, information []Query) (string, error) {
	api_key := os.Getenv("OPENAI_API_KEY")
	information_json, _ := json.Marshal(information)

	body_req := CreateBody(
		"You are a planner. User has a goal is: \"" + Escape(usergoal) + "\" and you have just got information about user's goal with the conversation in json format: " + string(information_json) + ".Divide " + time + " into maximum 10 interval time and tell user what they should do in each time to achieve their goal (just send a json format like this example [{\"time\": \"interval time...\", \"plan\": \"plan in that interval time...\"}...]",
	)
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(body_req)

	client := &http.Client{}
	req, err := http.NewRequest("POST", endpoint, buffer)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api_key)

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	bodyBytes, _ := io.ReadAll(resp.Body)
	body_resp := make(map[string]interface{})
	_ = json.Unmarshal(bodyBytes, &body_resp)

	plan_result := body_resp["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	log.Println(plan_result)

	return plan_result, nil
}
