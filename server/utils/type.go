package utils

type Query struct {
	Question string `json:"planner"`
	Answer   string `json:"user"`
}
type UserInfo struct {
	Usergoal string `json:"usergoal"`
	Time     string `json:"time"`
	Queries  []Query
}
