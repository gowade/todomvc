package main

import "github.com/gowade/vdom"

type TodoApp struct {
	Path  string
	Todos []Todo
}

func (todos TodoApp) Render() vdom.VNode {
	remainingCount := 0
	for _, todo := range todos.Todos {
		if !todo.Completed {
			remainingCount++
		}
	}

	filteredItems := []Todo{}
	switch todos.Path {
	case "active":
		for _, todo := range todos.Todos {
			if !todo.Completed {
				filteredItems = append(filteredItems, todo)
			}
		}
	case "completed":
		for _, todo := range todos.Todos {
			if todo.Completed {
				filteredItems = append(filteredItems, todo)
			}
		}
	default:
		filteredItems = todos.Todos
	}
	return &vdom.VElement{
		TagName: "section",
		Props: map[string]interface{}{
			"className": "todoapp",
		},
		Children: []vdom.VNode{
			Header{}.Render(),
			Todos{Todos: filteredItems}.Render(),
			Footer{Path: todos.Path, Todos: todos.Todos}.Render(),
		},
	}
}
