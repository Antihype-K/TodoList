package todo

import (
	"fmt"
	"time"

	"github.com/k0kubun/pp"
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
		Command:           "Добавить",
		DescriptionEvents: taskDescription,
	}
	t.Events = append(t.Events, event)
}

func (t *TodoManager) DeleteTask(taskName string) bool {
	for i, task := range t.Tasks {
		if task.NameTask == taskName {
			t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)

			event := Event{
				TimeMessage:       time.Now().Format("02.01.2006 15:04"),
				Command:           "Удалить",
				DescriptionEvents: fmt.Sprintf("Задача %v успешно удалена", taskName),
			}
			t.Events = append(t.Events, event)
			return true
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
				Command:           "Выполнено",
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
	event := Event{
		TimeMessage:       time.Now().Format("02.01.2006 15:04"),
		Command:           "list",
		DescriptionEvents: pp.Sprintf("Список задач: %v\n", t.Tasks),
	}
}

func (t *TodoManager) PullEvents() {
	fmt.Println(t.Events)
}
