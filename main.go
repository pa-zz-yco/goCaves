package main

import (
	"github.com/payco/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)

	defaultFG := termbox.ColorWhite
	defaultBG := termbox.ColorBlack

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	// convert termbox polling into a channel
	events := make(chan termbox.Event)
	go func() {
		for {
			events <- termbox.PollEvent()
		}
	}()
	// draw
	termbox.DrawString(
		0, 0,
		defaultFG, defaultBG,
		"Welcome to the Caves of Golang!")
	termbox.DrawString(
		0, 1,
		defaultFG, defaultBG,
		"Press any key to exit...")
	termbox.Flush()
	// keyboard for-loop
loop:
	for {
		select {
		case evt := <-events:
			if evt.Type == termbox.EventKey {
				break loop
			}
		}
	}
}
