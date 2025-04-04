package main

import (
	"embed"
	"flag"
	"log"
	"train_of_thought/backend/config"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	path := "./backend/env"

	// catch with '-appargs "args"' appended to 'wails dev
	envPtr := flag.String("env", "", "the current environment")
	flag.Parse()

	// TODO: remove once dev is complete
	if !config.IsFlagPassed("env") {
		flag.Set("env", "local")
	}

	if err := config.SetEnvironmentVariables(*envPtr, path); err != nil {
		// TODO: have an error show in the UI
		// panic due to inability to call APIs
		log.Panicln(err)
	}

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Train of Thought",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
