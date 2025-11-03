package todo

import (
	"fmt"
	"time"
	// "github.com/k0kubun/pp"
)

type Task struct {
	NameTask         string
	TextTask         string
	TimeTaskCreate   string
	StatusTask       bool
	TimeTaskComplete string
}

type Event struct {
	TimeMessage       string
	Command           string
	DescriptionEvents string
}

type TodoManager struct {
	Tasks  []Task
	Events []Event
}

func NewTodoManager() *TodoManager {
	return &TodoManager{
		Tasks:  []Task{},
		Events: []Event{},
	}
}

func (t *TodoManager) AddTask(taskName string, taskDescription string) {
	task := Task{
		NameTask:         taskName,
		TextTask:         taskDescription,
		TimeTaskCreate:   time.Now().Format("02.01.2006 15:04"),
		StatusTask:       false,
		TimeTaskComplete: "",
	}
	t.Tasks = append(t.Tasks, task)

	event := Event{
		TimeMessage:       time.Now().Format("02.01.2006 15:04"),
		Command:           "add",
		DescriptionEvents: fmt.Sprintf("Задача %v была добавлена", taskName),
	}
	t.Events = append(t.Events, event)
}

func (t *TodoManager) DeleteTask(taskName string) bool {
	for i, task := range t.Tasks {
		if task.NameTask == taskName {
			t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)

			event := Event{
				TimeMessage:       time.Now().Format("02.01.2006 15:04"),
				Command:           "dell",
				DescriptionEvents: fmt.Sprintf("Задача %v была удалена", taskName),
			}
			t.Events = append(t.Events, event)
			return true
		} else {
			fmt.Println("Такой задачи не существует, обратитесь к команде list или проверьте корректность написания задачи")
		}

	}
	return false
}

func (t *TodoManager) CompleteTask(taskName string) bool {
	for i := range t.Tasks {
		if t.Tasks[i].NameTask == taskName {
			t.Tasks[i].StatusTask = true
			t.Tasks[i].TimeTaskComplete = time.Now().Format("02.01.2006 15:04")

			event := Event{
				TimeMessage:       time.Now().Format("02.01.2006 15:04"),
				Command:           "done",
				DescriptionEvents: fmt.Sprintf("Задача %v успешно выполнена!", taskName),
			}
			t.Events = append(t.Events, event)
			return true
		}

	}
	fmt.Println("Задача с таким названием не найдена")
	return false
}

func (t *TodoManager) PullTasks() {
	for i, task := range t.Tasks {
		fmt.Println(i+1, "-", task)
	}
	event := Event{
		TimeMessage:       time.Now().Format("02.01.2006 15:04"),
		Command:           "list",
		DescriptionEvents: fmt.Sprintln("Запрошен список задач"),
	}
	t.Events = append(t.Events, event)
}

func (t *TodoManager) PullEvents() {
	for i, event := range t.Events {
		fmt.Println(i+1, "-", event)
	}
	event := Event{
		TimeMessage:       time.Now().Format("02.01.2006 15:04"),
		Command:           "events",
		DescriptionEvents: fmt.Sprintln("Запрошен список событий"),
	}
	t.Events = append(t.Events, event)

}
