package models

type Task struct {
	Id    string `json:"id"`
	Title string `json:"title" binding:"required"`
	Desc  string `json:"desc" binding:"required"`
}

