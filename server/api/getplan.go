package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"planner/utils"
)

func GetPlanHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			bodyBytes, _ := io.ReadAll(r.Body)
			body := make(map[string]string)
			err := json.Unmarshal(bodyBytes, &body)
			if err != nil {
				fmt.Fprintf(w, "Error: %s", err.Error())
			}
			plan, err := utils.GetPlanGpt(body["usergoal"], body["answers"])
			if err != nil {
				fmt.Fprintf(w, "Error: %s", err.Error())
			}
			resp, _ := json.Marshal(plan) 
			w.Write(resp)
		}
	})
}
