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
			var body utils.UserInfo
			err := json.Unmarshal(bodyBytes, &body)
			if err != nil {
				fmt.Fprintf(w, "Error: %s", err.Error())
			}
			questions, err := utils.GetPlanGpt(body.Usergoal, body.Time, body.Queries)
			if err != nil {
				fmt.Fprintf(w, "Error: %s", err.Error())
			}
			resp, _ := json.Marshal(questions)
			w.Write(resp)
		}
	})
}
