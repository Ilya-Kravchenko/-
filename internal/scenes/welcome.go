package scenes

import (
	"strings"

	"github.com/devllart/backtrace/internal/page"
	"github.com/devllart/backtrace/internal/state"
	"github.com/gopherjs/gopherjs/js"
)

func Welcome() {
	// page.SetText(page.MainEl, "BackTrace")
	var pl = state.Player
	pl.El = page.NewElement(page.MainEl, "div", map[string]any{
		"className": "player",
	})

	page.AddEventAction("keydown", func(e *js.Object) {
		var key = strings.ToLower(e.Get("key").String())

		if key == "d" && pl.X+pl.Speed < page.InnerWidth()-30 {
			pl.X += pl.Speed
		} else if key == "a" && pl.X-pl.Speed > 0 {
			pl.X -= pl.Speed
		} else if key == "w" && pl.Y-pl.Speed > 0 {
			pl.Y -= pl.Speed
		} else if key == "s" && pl.Y+pl.Speed < page.InnerHeight()-30 {
			pl.Y += pl.Speed
		}

		pl.Draw()
	})
	pl.Draw()
}
