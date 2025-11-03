package main

import (
	"Third/todo"
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	manager := todo.NewTodoManager()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("=== TODO LIST ===")
		fmt.Println("Доступные команды: add, list, delete, done, events, exit")
		fmt.Print("> ")
		scanner.Scan()
		commandText := scanner.Text()

		switch commandText {
		case "add":
			fmt.Print("Введите название задачи: ")
			scanner.Scan()
			nameTask := scanner.Text()
			fmt.Print("\nВведите текст задачи")
			scanner.Scan()
			taskDescription := scanner.Text()
			manager.AddTask(nameTask, taskDescription)

		case "list":
			manager.PullTasks()

		case "delete":
			fmt.Print("Введите название задачи которую хотите удалить")
			scanner.Scan()
			nameTask := scanner.Text()
			manager.DeleteTask(nameTask)

		case "exit":
			fmt.Println("Выход из программы...")
			return
		case "events":
			manager.PullEvents()
		}

	}

}
