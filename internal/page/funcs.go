package page

import "github.com/gopherjs/gopherjs/js"

func SetText(idEl string, value string) {
	js.Global.Get(idEl).Set("innerText", value)
}
