package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"planner/utils"
)

func GetQuesHandler() {
	http.HandleFunc("/getques", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			bodyBytes, _ := io.ReadAll(r.Body)
			body := make(map[string]string)
			err := json.Unmarshal(bodyBytes, &body)
			if err != nil {
				fmt.Fprintf(w, "Error: %s", err.Error())
			}
			questions, err := utils.GetQuesGpt(body["usergoal"])
			if err != nil {
				fmt.Fprintf(w, "Error: %s", err.Error())
			}
			resp, _ := json.Marshal(questions)
			w.Write(resp)
		}
	})
}
