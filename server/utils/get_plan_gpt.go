package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
)

func GetPlanGpt(usergoal string, time string, information []Query) (string, error) {
	api_key := os.Getenv("OPENAI_API_KEY")
	information_json, _ := json.Marshal(information)

	body_req := CreateBody(
		"You are a planner assistant. User has a goal is: \""+Escape(usergoal)+"\". They have exact "+time+" to achieve the goal no more no less and there are some extra information about user is provided in this json format: "+string(information_json)+".Divide "+time+" into maximum 10 periods. All the period do not need to have the same time interval. (Just send a json format the exact same as this example [{\"time\": \"Year a, Month b,... (if time is year you should divide like year a, month b; else if time is months you should divide like month a, week b; else if time is week you should divide like week a, day b. Remember a year only has 12 months , a month only has 4 weeks, a week only has 7 days. Please don't across these limits)\", \"plan\": \"plan in that period...\",...] and between 2 elements of the array mustn't have new line character).",
		0.1,
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
	// log.Println(plan_result)

	return strings.ReplaceAll(plan_result, "\n", ""), nil
}
