package models

type Animal struct {
	Name   string `json:"name"`
	Breed  string `json:"breed"`
	LegNum int    `json:"legNum"`
}