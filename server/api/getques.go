package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"planner/utils"
)

func GetQuesHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			bodyBytes, _ := io.ReadAll(r.Body)
			var body utils.UserInfo
			err := json.Unmarshal(bodyBytes, &body)
			log.Print("Received user's goal: ", body.Usergoal)
			if err != nil {
				fmt.Fprintf(w, "Error: %s", err.Error())
			}
			questions, err := utils.GetQuesGpt(body.Usergoal, body.Time)
			if err != nil {
				fmt.Fprintf(w, "Error: %s", err.Error())
			}
			log.Printf("Send %d questions successfully", len(questions))
			resp, _ := json.Marshal(questions)
			w.Write(resp)
		}
	})
}
