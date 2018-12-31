package view

import tb "github.com/nsf/termbox-go"
const HUDWidth = BoundBoxWidth
const HUDHeight = 5
type (
	HUD struct {
		View
	}
)

func NewHud(x, y int) HUD {
	return HUD {
		View{x, y},
	}
}

func (h HUD) Draw() {
	drawBorder(h)
}

func drawBorder(h HUD) {
	tb.SetCell(h.X+0, h.Y+0, DoubleTopLeftCorner, tb.ColorWhite, tb.ColorBlack)
	tb.SetCell(h.X+HUDWidth, h.Y, DoubleTopRightCorner, tb.ColorWhite, tb.ColorBlack)
	tb.SetCell(h.X, h.Y+HUDHeight, DoubleBotLeftCorner, tb.ColorWhite, tb.ColorBlack)
	tb.SetCell(h.X+HUDWidth, h.Y+HUDHeight, DoubleBotRightCorner, tb.ColorWhite, tb.ColorBlack)
	for col := 1; col < HUDWidth; col++ {
		tb.SetCell(h.X+col, h.Y+0, DoubleHorizontalBar, tb.ColorWhite, tb.ColorBlack)
		tb.SetCell(h.X+col, h.Y+HUDHeight, DoubleHorizontalBar, tb.ColorWhite, tb.ColorBlack)
	}
	for row := 1; row < HUDHeight; row++ {
		tb.SetCell(h.X, h.Y+row, DoubleVerticalBar, tb.ColorWhite, tb.ColorBlack)
		tb.SetCell(h.X+HUDWidth, h.Y+row, DoubleVerticalBar, tb.ColorWhite, tb.ColorBlack)
	}
}
