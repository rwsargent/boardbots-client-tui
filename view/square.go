package view

import (
	tb "github.com/nsf/termbox-go"
	"boardbotclient/model"
)

type (
	Square struct {
		View
		model.Square
		borderColor tb.Attribute
	}
)

func DefaultSquare(view View, square model.Square) Square {
	return Square{
		view,
		square,
		tb.ColorWhite,
	}
}

func (s *Square) Reset() {
	s.Highlighted = false
	s.borderColor = tb.ColorWhite
}

func (s *Square) Highlight() {
	s.Highlighted = true
	s.borderColor = tb.ColorYellow
}
func (s *Square) BorderColor() tb.Attribute{
	return s.borderColor
}

func (s *Square) Update(model model.Square) {

}
func (s *Square) DisplayCoordinates() {
	tb.SetCell(s.X + 1, s.Y+1, rune('0' + (s.Row)), tb.ColorWhite, tb.ColorBlack)
	tb.SetCell(s.X + 2, s.Y+1, rune('0' + (s.Col)), tb.ColorWhite, tb.ColorBlack)
}

func (s *Square) Draw() {
	tb.SetCell(s.X + 0, s.Y + 0, TopLeftCorner,  s.borderColor, tb.ColorBlack)
	tb.SetCell(s.X + 1, s.Y + 0, HorizontalBar,  s.borderColor, tb.ColorBlack)
	tb.SetCell(s.X + 2, s.Y + 0, HorizontalBar,  s.borderColor, tb.ColorBlack)
	tb.SetCell(s.X + 3, s.Y + 0, TopRightCorner, s.borderColor, tb.ColorBlack)
	tb.SetCell(s.X + 3, s.Y + 1, VerticalBar,    s.borderColor, tb.ColorBlack)
	tb.SetCell(s.X + 3, s.Y + 2, BotRightCorner, s.borderColor, tb.ColorBlack)
	tb.SetCell(s.X + 2, s.Y + 2, HorizontalBar,  s.borderColor, tb.ColorBlack)
	tb.SetCell(s.X + 1, s.Y + 2, HorizontalBar,  s.borderColor, tb.ColorBlack)
	tb.SetCell(s.X + 0, s.Y + 2, BotLeftCorner,  s.borderColor, tb.ColorBlack)
	tb.SetCell(s.X + 0, s.Y + 1, VerticalBar,    s.borderColor, tb.ColorBlack)
}