package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/guilherme-brandao/to-go-list/database"
	"github.com/guilherme-brandao/to-go-list/models"
	"go.mongodb.org/mongo-driver/bson"
)

type TaskService interface {
	Save(models.Task) models.Task
	FindAll() []models.Task
	GetTask(string) models.Task
	Update(string, models.Task) models.Task
	Delete(string) error
}

type taskService struct {
	tasks []models.Task
}

func New() TaskService {
	return &taskService{}
}

func (service *taskService) Save(task models.Task) models.Task {

	collection := database.GetCollection("task-manager")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, task)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted task with ID:", res.InsertedID)

	return task
}

func (service *taskService) GetTask(id string) models.Task {

	collection := database.GetCollection("task-manager")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var task models.Task
	collection.FindOne(ctx, bson.M{"id": id}).Decode(&task)

	return task
}

func (service *taskService) FindAll() []models.Task {

	collection := database.GetCollection("task-manager")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var tasks []models.Task

	items, err := collection.Find(ctx, bson.M{})
	items.All(ctx, &tasks)

	if err != nil {
		log.Fatal(err)
	}

	return tasks

}

func (service *taskService) Update(id string, task models.Task) models.Task {

	collection := database.GetCollection("task-manager")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updateResult, err := collection.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": task})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	return task
}

func (service *taskService) Delete(id string) error {

	collection := database.GetCollection("task-manager")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}

	return nil
}
