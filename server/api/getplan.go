package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"planner/utils"
)

func GetPlanHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			bodyBytes, _ := io.ReadAll(r.Body)
			var body utils.UserInfo
			err := json.Unmarshal(bodyBytes, &body)
			log.Println("Received user's information successfully")
			if err != nil {
				fmt.Fprintf(w, "Error: %s", err.Error())
			}
			plan, err := utils.GetPlanGpt(body.Usergoal, body.Time, body.Queries)
			if err != nil {
				fmt.Fprintf(w, "Error: %s", err.Error())
			}
			log.Println("Generate the plan successfully!")
			w.Write([]byte(plan))
		}
	})
}
