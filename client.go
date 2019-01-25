package main

import (
	"boardbotclient/view"
	"github.com/nsf/termbox-go"
	"boardbotclient/api"
	"boardbotclient/startup"
)

func main() {
	server := api.NewServer("http://localhost:8080")
	startup.StartUp(server)
//	err := termbox.Init()
//	if err != nil {
//		panic(err)
//	}
//	defer termbox.Close()
//	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
//	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
//	drawables := make([]view.Drawable, 0)
//	board := view.NewBoard(0, 0)
//	hud := view.NewHud(0,board.Height() + 1)
//	drawables = append(drawables, board)
//	drawables = append(drawables, hud)
//	// dbg := view.NewDebugBox(0, board.Y + board.Height() + 1)
//	render(drawables)
//loop:
//	for {
//		switch ev := termbox.PollEvent(); ev.Type {
//		case termbox.EventKey:
//			if ev.Key == termbox.KeyCtrlQ {
//				break loop
//			}
//		case termbox.EventResize:
//			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
//		case termbox.EventMouse:
//			board.ResetSquares()
//			if collide, square := board.GetSquare(ev.MouseX, ev.MouseY); collide {
//				square.Highlight()
//			}
//		}
//		render(drawables)
//	}
}

func render(drawables []view.Drawable) {
	for _, drawable := range drawables {
		drawable.Draw()
	}
	termbox.Flush()
}
