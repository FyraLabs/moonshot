package main

import (
	"embed"
	"log"
	"os"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
	"github.com/wailsapp/wails/v3/pkg/services/notifications"
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

	ns := notifications.New()
	authorized, err := ns.CheckNotificationAuthorization()
	if err != nil {
		println("Error checking notification authorization:", err.Error())
	}
	if !authorized {
		_, err := ns.RequestNotificationAuthorization()
		if err != nil {
			println("Error requesting notification authorization:", err.Error())
		}
	}

	app := application.New(application.Options{
		Name:        "moonshot",
		Description: "A simple and intuitive tool for flashing OS images to drives",
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Services: []application.Service{
			application.NewService(ns),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	app.RegisterService(application.NewService(NewAppService(app)))

	win := app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:          "Moonshot",
		Width:          712,
		MinWidth:       712,
		Height:         500,
		MinHeight:      500,
		EnableFileDrop: true,
	})

	win.OnWindowEvent(events.Common.WindowFilesDropped, func(event *application.WindowEvent) {
		files := event.Context().DroppedFiles()
		details := event.Context().DropTargetDetails()

		application.Get().Event.Emit("files-dropped", map[string]any{
			"files":   files,
			"details": details,
		})
	})

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
