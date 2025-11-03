package main

import (
	"Third/todo"
)

func main() {
	manager := todo.NewTodoManager()
	manager.AddTask("Изучить питон", "2 часа программирования")
	manager.AddTask("Изучить питон", "2 часа программирования")

}
