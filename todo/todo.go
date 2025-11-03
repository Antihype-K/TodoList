package todo

type Task struct {
	NameTask         string
	TextTask         string
	TimeTaskCreate   string
	StatusTask       bool
	TimeTaskComplete string
}

type Event struct {
	timeMessage       string
	DescriptionEvents string
}
