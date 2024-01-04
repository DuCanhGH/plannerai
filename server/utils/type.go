package utils

type Querry struct {
	Question string
	Answer   string
}
type Mydata struct {
	Usergoal string `json:"usergoal"`
	Querries []Querry
}
