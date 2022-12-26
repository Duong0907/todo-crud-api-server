package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"todo-crud-api-server/model"
)

const (
	dbuser = "root"
	dbpass = "09072003Tpl@msql"
	dbname = "todo"
)

func GetTodos() []todo.Todo {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()
	results, err := db.Query("SELECT * FROM todo")
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	todos := []todo.Todo{}
	for results.Next() {
		var todo todo.Todo
		err := results.Scan(&todo.Id, &todo.Content, &todo.Status)
		if err != nil {
			panic(err.Error())
		}
		todos = append(todos, todo)
	}
	return todos
}

func GetTodo(id int) *todo.Todo {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()
	results, err := db.Query("SELECT * FROM todo WHERE id = ?", id)
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	var todo todo.Todo
	if results.Next() {
		err := results.Scan(&todo.Id, &todo.Content, &todo.Status)
		if err != nil {
			return nil
		}
	}  else {
		return nil
	}
	return &todo
}

func AddTodo(todo todo.Todo) {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	insert, err := db.Query("INSERT INTO todo VALUES (?, ?, ?)", todo.Id, todo.Content, todo.Status)
	if (err != nil) {
		panic(err.Error())
	}
	defer insert.Close()
}

func RemoveTodo(id int) {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	delete, err := db.Query("DELETE FROM todo WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	defer delete.Close()
}

func UpdateTodo(id int, todo todo.Todo) {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	update, err := db.Query("UPDATE todo SET content = ?, status = ? WHERE id = ?", todo.Content, todo.Status, id)
	if err != nil {
		panic(err.Error())
	}
	defer update.Close()
}