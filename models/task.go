package models

type Task struct {
	Id    string `json:"id"`
	Title string `json:"title" binding:"required"`
	Desc  string `json:"desc" binding:"required"`
}

// type ToDoList struct {
// 	ListId		string	`json:"listId"`
// 	ListTitle 	string	`json:"listTitle" binding:"required"`
// 	ListTasks	[]Task 	`json:"tasks"`
// }
