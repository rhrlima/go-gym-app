package model

type Exercise struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Tags []string `json:"tags"`
}
