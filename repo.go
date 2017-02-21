package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var db, err = sql.Open("mysql", "root:qwerty@/todo-list")

func RepoAllTodo() (Todos, error) {

	rows, err := db.Query("SELECT * FROM todo")
	defer rows.Close()
	if err != nil {
		return Todos{}, fmt.Errorf(err.Error())
	} else {
		var tempTodos Todos
		for rows.Next() {
			var tempTodo Todo
			err = rows.Scan(&tempTodo.Id, &tempTodo.Name, &tempTodo.Completed, &tempTodo.Due)
			if err != nil {
				return Todos{}, fmt.Errorf(err.Error())
			} else {
				tempTodos = append(tempTodos, tempTodo)
			}
		}
		return tempTodos, nil
	}
	return Todos{}, fmt.Errorf("Error")
}

func RepoFindTodo(id string) (Todo, error) {

	rows, err := db.Query("SELECT * FROM todo WHERE Id='" + id + "'")
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		var tempTodo Todo
		var rowCount = 0
		for rows.Next() {
			rowCount++
			err = rows.Scan(&tempTodo.Id, &tempTodo.Name, &tempTodo.Completed, &tempTodo.Due)
			if err != nil {
				return Todo{}, fmt.Errorf(err.Error())
			} else {
				return tempTodo, nil
			}
		}
	}
	return Todo{}, fmt.Errorf("error")
}

func RepoCreateTodo(t Todo) (string, error) {

	res, err := db.Exec("INSERT INTO todo (Name, Completed, Due) VALUES ('" + t.Name + "','" + strconv.Itoa(t.Completed) + "','" + t.Due + "')")
	if err != nil {
		return "", fmt.Errorf(err.Error())
	} else {
		id, err := res.LastInsertId()
		if err != nil {
			return "", fmt.Errorf(err.Error())
		} else {
			return strconv.FormatInt(id, 10), nil
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	return "", fmt.Errorf("error")
}

func RepoDeleteTodo(id string) error {

	res, err := db.Exec("DELETE FROM todo WHERE Id='" + id + "'")
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	deleted, _ := res.RowsAffected()
	if deleted == 0 {
		return fmt.Errorf("Nothing for delete")
	}
	return nil
}
