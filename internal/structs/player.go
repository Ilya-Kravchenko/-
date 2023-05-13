package structs

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
)

type Player struct {
	W, H, X, Y, Speed int
	El          *js.Object
}

func (pl *Player) Draw() {
	var translate = fmt.Sprintf("translate: %vpx %vpx", pl.X, pl.Y)
	pl.El.Set("style", translate)
}
