package view

import "github.com/nsf/termbox-go"

type (
	DebugBox struct {
		View
		message string
		prevLength int
	}
)

func NewDebugBox(x, y int) DebugBox {
	return DebugBox{
		View{x, y},
		"",
		0}
}

func (db *DebugBox) Draw() {
	off := 0
	for _, c := range db.message {
		termbox.SetCell(db.X + off, db.Y, c, termbox.ColorWhite, termbox.ColorBlack)
		off++
	}
	for ; off < db.prevLength; off++ {
		termbox.SetCell(db.X + off, db.Y, ' ', termbox.ColorWhite, termbox.ColorBlack)
	}
	db.prevLength = len(db.message)
}

func (db *DebugBox) SetMessage(message string) {
	db.message = message
}
