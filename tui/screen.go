package tui

import (
	"time"

	"github.com/gdamore/tcell"
)

type KeyHandler func(tcell.EventKey)
type MainLoopHandler func(tcell.Screen)

func Screen(speed int, keyHandler KeyHandler, mainLoop MainLoopHandler, quit <-chan struct{}) error {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, e := tcell.NewScreen()
	if e != nil {
		return e
	}
	defer s.Fini()

	if e = s.Init(); e != nil {
		return e
	}

	s.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorBlack).
		Background(tcell.ColorWhite))
	s.Clear()

	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				keyHandler(*ev)
			case *tcell.EventResize:
				s.Sync()
			}
		}
	}()

loop:
	for {
		select {
		case <-quit:
			break loop
		case <-time.After(time.Duration(speed) * time.Millisecond):
			mainLoop(s)
		}
	}

	return nil
}
