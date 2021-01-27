package models

type Task struct {
	Id    string `json:"id"`
	Desc  string `json:"desc" binding:"required"`
}

type List struct {
	Id    	string 	`json:"id"`
	Title  	string 	`json:"title" binding:"required"`
	Tasks   []Task	`json:"tasks" binding:"required"`
}