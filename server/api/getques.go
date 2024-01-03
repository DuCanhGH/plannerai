package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"planner/utils"
)

func GetQuesHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			bodyBytes, _ := io.ReadAll(r.Body)
			var body utils.Mydata
			err := json.Unmarshal(bodyBytes, &body)
			if err != nil {
				fmt.Fprintf(w, "Error: %s", err.Error())
			}
			questions, err := utils.GetQuesGpt(body.Usergoal)
			if err != nil {
				fmt.Fprintf(w, "Error: %s", err.Error())
			}
			resp, _ := json.Marshal(questions)
			w.Write(resp)
		}
	})
}
