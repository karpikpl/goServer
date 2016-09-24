package main

import "time"

// Todo object for handlign todos
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos is a collection of Todo objects
type Todos []Todo
