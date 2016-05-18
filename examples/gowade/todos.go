package main

import (
	"github.com/gowade/vdom"
	"honnef.co/go/js/dom"
)

type Todos struct {
	Todos []Todo
}

func (todos Todos) allCompleted(event dom.Event) {
	store.Dispatch(UpdateAll{
		Completed: event.Target().(*dom.HTMLInputElement).Checked,
	})
}

// This section should be hidden by default and shown when there are todos.
// <section class="main">
// 	<input class="toggle-all" type="checkbox">
// 	<label for="toggle-all">Mark all as complete</label>
// 	<ul class="todo-list">
// 	</ul>
// </section>

func (todos Todos) Render() vdom.VNode {
	if len(todos.Todos) == 0 {
		return nil
	}

	allCompleted := true
	for _, todo := range todos.Todos {
		if !todo.Completed {
			allCompleted = false
			break
		}
	}

	items := []vdom.VNode{}
	for _, todo := range todos.Todos {
		items = append(items, todo.Render())
	}

	return &vdom.VElement{
		TagName: "section",
		Props: map[string]interface{}{
			"className": "main",
		},
		Children: []vdom.VNode{
			&vdom.VElement{
				TagName: "input",
				Props: map[string]interface{}{
					"className": "toggle-all",
					"type":      "checkbox",
					"checked":   allCompleted,
					"onclick":   todos.allCompleted,
				},
			},
			&vdom.VElement{
				TagName: "label",
				Props: map[string]interface{}{
					"htmlFor": "toggle-all",
				},
				Children: []vdom.VNode{
					vdom.VText("Mark all as complete"),
				},
			},
			&vdom.VElement{
				TagName: "ul",
				Props: map[string]interface{}{
					"className": "todo-list",
				},
				Children: items,
			},
		},
	}
}
