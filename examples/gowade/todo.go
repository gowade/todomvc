package main

import (
	"strings"

	"github.com/gowade/vdom"
	"honnef.co/go/js/dom"
)

type Todo struct {
	ID        string
	Text      string
	Editing   bool
	Completed bool
}

func (todo Todo) toggleCompleted(event dom.Event) {
	store.Dispatch(ToggleCompleted{
		ID: todo.ID,
	})
}

func (todo Todo) startEditing(event dom.Event) {
	store.Dispatch(StartEditing{
		ID: todo.ID,
	})
}

func (todo Todo) handleKeyup(event dom.Event) {
	if event, ok := event.(*dom.KeyboardEvent); ok {
		switch event.KeyCode {
		case 13:
			todo.finishEditing(event)
		case 27:
			event.Target().(*dom.HTMLInputElement).Value = todo.Text
			todo.finishEditing(event)
		}
	}
}

func (todo Todo) finishEditing(event dom.Event) {
	text := event.Target().(*dom.HTMLInputElement).Value
	if text == "" {
		todo.destroy(event)
		return
	}
	store.Dispatch(FinishEditing{
		ID:   todo.ID,
		Text: text,
	})
}

func (todo Todo) destroy(event dom.Event) {
	store.Dispatch(Destroy{
		ID: todo.ID,
	})
}

func (todo Todo) focusHook(element dom.Element) {
	element.(dom.HTMLElement).Focus()
}

// List items should get the class `editing` when editing and `completed` when marked as completed
// <li class="completed">
// 	<div class="view">
// 		<input class="toggle" type="checkbox" checked>
// 		<label>Taste JavaScript</label>
// 		<button class="destroy"></button>
// 	</div>
// 	<input class="edit" value="Create a TodoMVC template">
// </li>

func (todo Todo) Render() vdom.VNode {
	classes := []string{}
	hooks := []func(dom.Element){}
	if todo.Completed {
		classes = append(classes, "completed")
	}
	if todo.Editing {
		classes = append(classes, "editing")
		hooks = append(hooks, todo.focusHook)
	}
	return &vdom.VElement{
		TagName: "li",
		Props: map[string]interface{}{
			// "key":   todo.ID,
			"className": strings.Join(classes, " "),
		},
		Children: []vdom.VNode{
			&vdom.VElement{
				TagName: "div",
				Props: map[string]interface{}{
					"className": "view",
				},
				Children: []vdom.VNode{
					&vdom.VElement{
						TagName: "input",
						Props: map[string]interface{}{
							"className": "toggle",
							"type":      "checkbox",
							"checked":   todo.Completed,
							"onchange":  todo.toggleCompleted,
						},
					},
					&vdom.VElement{
						TagName: "label",
						Props: map[string]interface{}{
							"ondblclick": todo.startEditing,
						},
						Children: []vdom.VNode{
							vdom.VText(todo.Text),
						},
					},
					&vdom.VElement{
						TagName: "button",
						Props: map[string]interface{}{
							"className": "destroy",
							"onclick":   todo.destroy,
						},
					},
				},
			},
			&vdom.VElement{
				TagName: "input",
				Props: map[string]interface{}{
					"className": "edit",
					"onkeyup":   todo.handleKeyup,
					"onblur":    todo.finishEditing,
					"hooks":     hooks,
					"value":     todo.Text,
				},
			},
		},
	}
}
