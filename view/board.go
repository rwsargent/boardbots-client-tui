package view

import (
	"boardbotclient/model"
	tb "github.com/nsf/termbox-go"
)

const BoundBoxWidth = 47
const BoundBoxHeight = 28
const BoardXOffset = 1
const BoardYOffset = 1
const SquareXOffset = 5
const SquareYOffset = 3

var DisplayCoordinates = true

type Board struct {
	View
	Board   model.Board
	Squares []Square
}

func NewBoard(x, y int) *Board {
	b := model.Board{}
	b.Pieces = make(map[model.Pos]model.Piece)
	b.Pieces[model.Pos{4, 0}] = model.Piece{
		Type:           1,
		PlayerPosition: 0,
	}
	b.Pieces[model.Pos{4, 8}] = model.Piece{
		Type:           1,
		PlayerPosition: 1,
	}
	origin := View{x, y}
	squares := makeSquares(origin)
	return &Board{origin, b, squares}
}
func makeSquares(board View) []Square {
	squares := make([]Square, 0)
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			xOffset := col * SquareXOffset
			yOffset := row * SquareYOffset
			x := board.X + BoardXOffset + xOffset
			y := board.Y + BoardYOffset + yOffset
			squares = append(squares, DefaultSquare(View{x, y},
				model.Square{
					model.Pos{
						col, row,
					},
					false,
				},
			))
		}
	}
	return squares
}

func (b *Board) Draw() {
	b.drawBoundary()
	b.drawBoard()
	b.drawPieces()
}

func (b *Board) Width() int {
	return b.X + BoundBoxWidth
}

func (b *Board) Height() int {
	return b.Y + BoundBoxHeight
}

func (b *Board) drawBoundary() {
	b.drawBoundaryCorners()
	for col := 1; col < BoundBoxWidth; col++ {
		tb.SetCell(b.X+col, b.Y+0, DoubleHorizontalBar, tb.ColorWhite, tb.ColorBlack)
		tb.SetCell(b.X+col, b.Y+BoundBoxHeight, DoubleHorizontalBar, tb.ColorWhite, tb.ColorBlack)
	}
	for row := 1; row < BoundBoxHeight; row++ {
		tb.SetCell(b.X, b.Y+row, DoubleVerticalBar, tb.ColorWhite, tb.ColorBlack)
		tb.SetCell(b.X+BoundBoxWidth, b.Y+row, DoubleVerticalBar, tb.ColorWhite, tb.ColorBlack)
	}
}

func (b *Board) drawBoundaryCorners() {
	tb.SetCell(b.X+0, b.Y+0, DoubleTopLeftCorner, tb.ColorWhite, tb.ColorBlack)
	tb.SetCell(b.X+BoundBoxWidth, b.Y, DoubleTopRightCorner, tb.ColorWhite, tb.ColorBlack)
	tb.SetCell(b.X, b.Y+BoundBoxHeight, DoubleBotLeftCorner, tb.ColorWhite, tb.ColorBlack)
	tb.SetCell(b.X+BoundBoxWidth, b.Y+BoundBoxHeight, DoubleBotRightCorner, tb.ColorWhite, tb.ColorBlack)
}
func (b *Board) drawBoard() {
	for _, square := range b.Squares {
		if DisplayCoordinates {
			square.DisplayCoordinates()
		}
		square.Draw()
	}
}
func (b *Board) drawPieces() {
	for pos, piece := range b.Board.Pieces {
		if x, y := translate(pos); piece.Type == 0 {
			// drawn barrier
		} else {
			x = b.X + BoardXOffset + x*SquareXOffset
			y = b.Y + BoardYOffset + y*SquareYOffset
			tb.SetCell(x+1, y+1, 'P', tb.ColorRed, tb.ColorWhite)
			tb.SetCell(x+2, y+1, rune('0'+piece.PlayerPosition), tb.ColorRed, tb.ColorWhite)
		}
	}
}
func (b *Board) GetSquare(mouseX int, mouseY int) (bool, *Square) {
	if !(mouseX > b.X && mouseX < b.Width()  &&
		 mouseY > b.Y && mouseY < b.Height()) {
		 	return false, nil
	}
	squareX := (mouseX - b.X - BoardXOffset)/ SquareXOffset
	squareY := (mouseY - b.Y - BoardYOffset)/ SquareYOffset
	idx := (squareY * 9) + squareX
	if idx < 0 || idx >= len(b.Squares) {
		return false, nil
	}
	return true, &b.Squares[idx]
}
func (b *Board) HighlightSquare(mouseX int, mouseY int) {
	if !(mouseX > b.X && mouseX < b.Width()  &&
		mouseY > b.Y && mouseY < b.Height()) {
			return
	}
	squareX := (mouseX - b.X - BoardXOffset)/ SquareXOffset
	squareY := (mouseY - b.Y - BoardYOffset)/ SquareYOffset
	idx := (squareX * 9) + squareY
	if idx < 0 || idx >= len(b.Squares) {
		return
	}
	b.Squares[idx].Highlight()
}
func (b *Board) ResetSquares() {
	for i := 0; i < len(b.Squares); i++ {
		b.Squares[i].Reset()
	}
}

func translate(pos model.Pos) (int, int) {
	return pos.Col, pos.Row
}
