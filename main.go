package main

import (
	"embed"
	"log"
	"os"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	if len(os.Args) > 2 && os.Args[1] == "flash" {
		if err := flash(); err != nil {
			println(err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}

	app := application.New(application.Options{
		Name:        "moonshot",
		Description: "A simple and intuitive tool for flashing OS images to drives",
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	app.RegisterService(application.NewService(NewAppService(app)))

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:          "Moonshot",
		Width:          712,
		MinWidth:       712,
		Height:         500,
		MinHeight:      500,
		EnableFileDrop: true,
	})

	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
