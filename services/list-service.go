package services

import (
	"context"
	"log"
	"time"

	"github.com/guilherme-brandao/to-go-list/database"
	"github.com/guilherme-brandao/to-go-list/models"
	"go.mongodb.org/mongo-driver/bson"
)

type ListService interface {
	NewTask(string, models.Task) models.Task
	NewList(models.List) models.List
	FindAll() []models.List
	GetList(string) models.List
	DeleteList(string) error
	DeleteTask(string, string) error
}

type listService struct {
	lists []models.List
}

func New() ListService {
	return &listService{}
}

func (service *listService) NewList(list models.List) models.List {

	collection := database.GetCollection("task-manager")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, list)
	if err != nil {
		log.Fatal(err)
	}

	return list
}

func (service *listService) NewTask(idList string, task models.Task) models.Task {

	collection := database.GetCollection("task-manager")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var list models.List
	collection.FindOne(ctx, bson.M{"id": idList}).Decode(&list)

	list.Tasks = append(list.Tasks, task)

	_, err := collection.UpdateOne(ctx, bson.M{"id": idList}, bson.M{"$set": list})
	if err != nil {
		log.Fatal(err)
	}

	return task
}

func (service *listService) GetList(id string) models.List {

	collection := database.GetCollection("task-manager")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var list models.List
	collection.FindOne(ctx, bson.M{"id": id}).Decode(&list)

	return list
}

func (service *listService) FindAll() []models.List {

	collection := database.GetCollection("task-manager")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var lists []models.List

	items, err := collection.Find(ctx, bson.M{})
	items.All(ctx, &lists)

	if err != nil {
		log.Fatal(err)
	}

	return lists

}

func (service *listService) DeleteTask(idList string, idTask string) error {

	collection := database.GetCollection("task-manager")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var list models.List
	collection.FindOne(ctx, bson.M{"id": idList}).Decode(&list)

	var index = 0

	for i, n := range list.Tasks {
		if n.Id == idList {
			index = i
		}
	}

	list.Tasks = append(list.Tasks[:index], list.Tasks[index+1:]...)


	_, err := collection.UpdateOne(ctx, bson.M{"id": idList}, bson.M{"$set": list})
	if err != nil {
		return err
	}

	return nil
}

func (service *listService) DeleteList(idList string) error {

	collection := database.GetCollection("task-manager")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"id": idList})
	if err != nil {
		return err
	}

	return nil
}

