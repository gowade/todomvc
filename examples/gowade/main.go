package main

import (
	"github.com/go-humble/locstor"
	"github.com/gowade/wade"
	"honnef.co/go/js/dom"
)

var store *wade.Store
var router *wade.Router

func main() {
	localstore := locstor.NewDataStore(locstor.JSONEncoding)

	var initialState TodoApp
	if err := localstore.Find("state", &initialState); err != nil {
		initialState = TodoApp{
			Todos: []Todo{},
		}
	}

	store = wade.NewStore(rootReducer, initialState)

	store.Subscribe(func() {
		if err := localstore.Save("state", store.GetState().(TodoApp)); err != nil {
			panic("unable to save state in localStorage.")
		}
	})

	router = wade.NewRouter(store, wade.RouterOptions{
		Prefix:      "/examples/gowade",
		ForceHashes: true,
	})

	// go store.Run()

	element := dom.GetWindow().Document().GetElementByID("root")
	wade.Render(store, element, wade.RenderOptions{
		DisableRAF: true,
	})
}
