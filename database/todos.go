package db

import (
	"context"
	"time"
	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"todo-crud-api-server/models"
)

func GetTodos() ([]bson.M, error) {
	var todosBson []bson.M

	todoCollection := openCollection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	cursor, err := todoCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &todosBson); err != nil { // All() function auto closes curor
		return nil, err
	}
	return todosBson, nil
}

func GetTodo(id int) (bson.M, error) {
	var todosBson bson.M

	todoCollection := openCollection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	cursor, err := todoCollection.Find(ctx, bson.M{"_id": id})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.Decode(todosBson); err != nil { 
		return nil, err
	}
	return todosBson, nil
}

func AddTodo(todo todo.Todo) error {
	todoCollection := openCollection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	_, err := todoCollection.InsertOne(ctx, todo)
	return err
}

func RemoveTodo(id string) error {
	docID, _ := primitive.ObjectIDFromHex(id)
	todoCollection := openCollection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	_, err := todoCollection.DeleteOne(ctx, bson.M{"_id": docID})
	return err
}

func UpdateTodo(id string, todo todo.Todo) error {
	docID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": docID,
	}
	replacement := bson.M{
		"_id": todo.Id,
		"content": todo.Content,
		"status": todo.Status,
	}

	todoCollection := openCollection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	_, err := todoCollection.ReplaceOne(ctx, filter, replacement)
	return err
}
