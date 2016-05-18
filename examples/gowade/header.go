package main

import (
	"github.com/gowade/vdom"
	"honnef.co/go/js/dom"
)

type Header struct{}

func (header Header) CreateToDo(event dom.Event) {
	if event, ok := event.(*dom.KeyboardEvent); ok && event.KeyCode == 13 {
		element := event.Target().(*dom.HTMLInputElement)
		store.Dispatch(NewTodo{
			Text: element.Value,
		})
		element.Value = ""
		return
	}
}

// <header class="header">
// 	<h1>todos</h1>
// 	<input class="new-todo" placeholder="What needs to be done?" autofocus>
// </header>

func (header Header) Render() vdom.VNode {
	return &vdom.VElement{
		TagName: "header",
		Props: map[string]interface{}{
			"className": "header",
		},
		Children: []vdom.VNode{
			&vdom.VElement{
				TagName: "h1",
				Children: []vdom.VNode{
					vdom.VText("todos"),
				},
			},
			&vdom.VElement{
				TagName: "input",
				Props: map[string]interface{}{
					"className":   "new-todo",
					"placeholder": "What needs to be done?",
					"autofocus":   true,
					"onkeyup":     header.CreateToDo,
				},
			},
		},
	}
}
