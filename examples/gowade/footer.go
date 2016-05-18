package main

import (
	"strconv"

	"honnef.co/go/js/dom"

	"github.com/gowade/vdom"
)

type Footer struct {
	Path  string
	Todos []Todo
}

func (footer Footer) changeFilter(event dom.Event) {
	router.NavigateEvent(event)
}

func (footer Footer) clearCompleted(dom.Event) {
	store.Dispatch(ClearCompleted{})
}

func (footer Footer) getFilterClasses(path string) string {
	if footer.Path == path {
		return "selected"
	}
	return ""
}

// This footer should hidden by default and shown when there are todos
// <footer class="footer">
// 	This should be `0 items left` by default
// 	<span class="todo-count"><strong>0</strong> item left</span>
// 	Remove this if you don't implement routing
// 	<ul class="filters">
// 		<li>
// 			<a class="selected" href="#/">All</a>
// 		</li>
// 		<li>
// 			<a href="#/active">Active</a>
// 		</li>
// 		<li>
// 			<a href="#/completed">Completed</a>
// 		</li>
// 	</ul>
// 	Hidden if no completed items are left
// 	<button class="clear-completed">Clear completed</button>
// </footer>

func (footer Footer) Render() vdom.VNode {
	if len(footer.Todos) < 1 {
		return nil
	}
	itemsLeft := 0
	for _, todo := range footer.Todos {
		if !todo.Completed {
			itemsLeft++
		}
	}
	itemsLeftText := " items left"
	if itemsLeft == 1 {
		itemsLeftText = " item left"
	}
	return &vdom.VElement{
		TagName: "footer",
		Props: map[string]interface{}{
			"className": "footer",
		},
		Children: []vdom.VNode{
			&vdom.VElement{
				TagName: "span",
				Props: map[string]interface{}{
					"className": "todo-count",
				},
				Children: []vdom.VNode{
					&vdom.VElement{
						TagName: "strong",
						Children: []vdom.VNode{
							vdom.VText(strconv.Itoa(itemsLeft)),
						},
					},
					vdom.VText(itemsLeftText),
				},
			},
			&vdom.VElement{
				TagName: "ul",
				Props: map[string]interface{}{
					"className": "filters",
				},
				Children: []vdom.VNode{
					&vdom.VElement{
						TagName: "li",
						Children: []vdom.VNode{
							&vdom.VElement{
								TagName: "a",
								Props: map[string]interface{}{
									"href":      "",
									"className": footer.getFilterClasses(""),
									"onclick":   footer.changeFilter,
								},
								Children: []vdom.VNode{
									vdom.VText("All"),
								},
							},
						},
					},
					&vdom.VElement{
						TagName: "li",
						Children: []vdom.VNode{
							&vdom.VElement{
								TagName: "a",
								Props: map[string]interface{}{
									"href":      "active",
									"className": footer.getFilterClasses("active"),
									"onclick":   footer.changeFilter,
								},
								Children: []vdom.VNode{
									vdom.VText("Active"),
								},
							},
						},
					},
					&vdom.VElement{
						TagName: "li",
						Children: []vdom.VNode{
							&vdom.VElement{
								TagName: "a",
								Props: map[string]interface{}{
									"href":      "completed",
									"className": footer.getFilterClasses("completed"),
									"onclick":   footer.changeFilter,
								},
								Children: []vdom.VNode{
									vdom.VText("Completed"),
								},
							},
						},
					},
				},
			},
			&vdom.VElement{
				TagName: "button",
				Props: map[string]interface{}{
					"className": "clear-completed",
					"onclick":   footer.clearCompleted,
				},
				Children: []vdom.VNode{
					vdom.VText("Clear completed"),
				},
			},
		},
	}
}
