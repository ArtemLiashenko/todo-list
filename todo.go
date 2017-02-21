package main

//import "time"

type Todo struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Completed int    `json:"completed"`
	Due       string `json:"due"`
}

type Todos []Todo
