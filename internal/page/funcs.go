package page

import (
	"github.com/devllart/backtrace/internal/names"
	"github.com/gopherjs/gopherjs/js"
)

// aliases
var global = js.Global
var document = global.Get("document")
var window = global.Get("window")

var MainEl = global.Get(names.MainElID)

func Screen() *js.Object {
	return window.Get("Screen")
}

func InnerHeight() int {
	// fmt.Println(window.Get("innerHeight"))
	return window.Get("innerHeight").Int()
}

func InnerWidth() int {
	return window.Get("innerWidth").Int()
}
func Alert(msg string) {
	global.Call("alert", msg)
}

func GetById(idEl string) *js.Object {
	return global.Get(idEl)
}

func SetText(el *js.Object, value string) {
	el.Set("innerText", value)
}

func NewElement(parentEl *js.Object, typeEl string, opts map[string]any) *js.Object {
	// var parentEl = global.Get(parentElID)
	var el = document.Call("createElement", typeEl)

	for key, opt := range opts {
		if key == "text" {
			var text = document.Call("createTextNode", opt)
			el.Call("appendChild", text)
		} else if key == "content" {
			el.Call("appendChild", opt)
		} else {
			el.Set(key, opt)
		}
	}

	parentEl.Call("appendChild", el)

	return el
}

func AddEventAction(action string, fn func(e *js.Object)) {
	document.Call("addEventListener", action, func(e *js.Object) {
		go func(e *js.Object) {
			fn(e)
		}(e)
	})
}
