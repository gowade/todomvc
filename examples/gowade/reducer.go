package main

import "github.com/gowade/wade"

func rootReducer(state wade.State, action wade.Action) wade.State {
	todos := state.(TodoApp)
	switch action := action.(type) {
	case wade.PathChanged:
		todos.Path = string(action.Path)
	case NewTodo:
		todos.Todos = append(todos.Todos, Todo{
			ID:        uuid(),
			Text:      action.Text,
			Completed: false,
		})
	case UpdateAll:
		for index := range todos.Todos {
			todos.Todos[index].Completed = action.Completed
		}
	case ToggleCompleted:
		for index := range todos.Todos {
			if todos.Todos[index].ID == action.ID {
				todos.Todos[index].Completed = !todos.Todos[index].Completed
			}
		}
	case StartEditing:
		for index := range todos.Todos {
			if todos.Todos[index].ID == action.ID {
				todos.Todos[index].Editing = true
			}
		}
	case FinishEditing:
		for index := range todos.Todos {
			if todos.Todos[index].ID == action.ID {
				todos.Todos[index].Text = action.Text
				todos.Todos[index].Editing = false
			}
		}
	case Destroy:
		for i := len(todos.Todos) - 1; i >= 0; i-- {
			if todos.Todos[i].ID == action.ID {
				todos.Todos = append(todos.Todos[:i], todos.Todos[i+1:]...)
			}
		}
	case ClearCompleted:
		for i := len(todos.Todos) - 1; i >= 0; i-- {
			if todos.Todos[i].Completed {
				todos.Todos = append(todos.Todos[:i], todos.Todos[i+1:]...)
			}
		}
	}
	return todos
}
