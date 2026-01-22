package main

import (
	"embed"
	"moonshot/lib"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	if len(os.Args) > 2 && os.Args[1] == "flash" {
		ch := make(chan int)
		go func() {
			select {
			case n := <-ch:
				println("Flashed", n, "bytes")
			}
		}()
		if err := lib.Flash(os.Args[2], os.Args[3], ch); err != nil {
			println("Error:", err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Moonshot",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		DragAndDrop: &options.DragAndDrop{
			EnableFileDrop: true,
		},
		// Mac: &mac.Options{
		// 	TitleBar: mac.TitleBarHidden(),
		// },
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
