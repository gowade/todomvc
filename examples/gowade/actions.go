package main

type SetVisibilityFilter struct {
	Filter int
}

type NewTodo struct {
	Text string
}

type UpdateAll struct {
	Completed bool
}

type ToggleCompleted struct {
	ID string
}

type StartEditing struct {
	ID string
}

type FinishEditing struct {
	ID   string
	Text string
}

type Destroy struct {
	ID string
}

type ClearCompleted struct{}
