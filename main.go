package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"

	"github.com/skeetcha/goplural/themes"
)

func main() {
	themes.SetupThemes()

	go func() {
		window := new(app.Window)
		err := run(window)

		if err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}()

	app.Main()
}

func run(window *app.Window) error {
	//theme := material.NewTheme()
	var ops op.Ops

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			e.Frame(gtx.Ops)
		}
	}
}
