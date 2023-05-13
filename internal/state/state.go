package state

import (
	"github.com/devllart/backtrace/internal/page"
	"github.com/devllart/backtrace/internal/structs"
)

var Player = structs.Player{
	X:     0,
	Y:     page.InnerHeight() - 30,
	Speed: 5,
	El:    nil,
}
