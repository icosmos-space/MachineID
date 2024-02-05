package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/atotto/clipboard"
	"github.com/denisbrodbeck/machineid"
)

func main() {
	go func() {
		w := app.NewWindow()
		err := run(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(w *app.Window) error {
	th := material.NewTheme()
	var ops op.Ops
	for {
		switch e := w.NextEvent().(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := layout.NewContext(&ops, e)

			// Define an large label with an appropriate text:
			id, _ := machineid.ProtectedID("myAppName")
			clipboard.WriteAll(id)
			title := material.H4(th, "机器ID已复制到剪切板,可以直接粘贴使用")

			// Change the color of the label.
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			title.Color = maroon

			// Change the position of the label.
			title.Alignment = text.Middle

			// Draw the label to the graphics context.
			title.Layout(gtx)

			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
		}
	}
}
