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
		todo.LogCommand(commandText)

		switch commandText {
		case "add":
			fmt.Print("Введите название задачи: ")
			scanner.Scan()
			nameTask := scanner.Text()
			todo.LogCommand(nameTask)
			fmt.Print("\nВведите текст задачи")
			scanner.Scan()
			taskDescription := scanner.Text()
			todo.LogCommand(taskDescription)
			manager.AddTask(nameTask, taskDescription)

		case "list":
			manager.PullTasks()

		case "delete":
			fmt.Print("Введите название задачи которую хотите удалить")
			scanner.Scan()

			nameTask := scanner.Text()
			todo.LogCommand(nameTask)
			manager.DeleteTask(nameTask)

		case "done":
			fmt.Println("Введите название задачи которую хотите выполнить?")
			scanner.Scan()
			nameTask := scanner.Text()
			todo.LogCommand(nameTask)
			manager.CompleteTask(nameTask)
		case "exit":
			fmt.Println("Выход из программы...")
			return
		case "events":
			manager.PullEvents()
		case "help":
			todo.LogCommand("help")
			fmt.Println(`💡 Команды: 
    add <заголовок> <текст> - добавить задачу
    list - показать задачи
    delete <заголовок> - удалить задачу
    done <заголовок> - отметить выполненной
    events - история событий
    exit - выход`)
		}
		fmt.Println("")

	}

}
