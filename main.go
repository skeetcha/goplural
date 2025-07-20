package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/skeetcha/goplural/themes"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

func main() {
	themes.SetupThemes()

	go func() {
		window := new(app.Window)
		err := mainWindow(window)

		if err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}()

	app.Main()
}

func settingsWindow(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,
				layout.Flexed(1, func(gtx C) D {
					return material.H3(theme, "General Settings").Layout(gtx)
				}),
				layout.Flexed(1, func(gtx C) D {
					return material.Body1(theme, "Select a Theme:").Layout(gtx)
				}),
				// select
				// current theme message
				// font settings header
				// horizontal flex
				// // font label
				// // font select
				// // size label
				// // size select
				// // apply font button
				// window size header
				// horizontal flex
				// // width input box
				// // x label
				// // height input box
				// apply size button
			)

			e.Frame(gtx.Ops)
		}
	}
}

func mainWindow(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops

	var settingsButton widget.Clickable
	settingsOpen := false

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			if settingsButton.Clicked(gtx) && !settingsOpen {
				settingsOpen = true

				go func() {
					window := new(app.Window)
					err := settingsWindow(window)

					if err != nil {
						log.Fatal(err)
					}

					settingsOpen = false
				}()
			}

			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,
				layout.Flexed(1, func(gtx C) D {
					return layout.Flex{
						Axis: layout.Horizontal,
					}.Layout(gtx,
						layout.Flexed(1, func(gtx C) D {
							return material.Button(theme, &settingsButton, "Settings").Layout(gtx)
						}),
					)
				}),
			)
			e.Frame(gtx.Ops)
		}
	}
}
