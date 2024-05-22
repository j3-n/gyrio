package main

import (
	"log"
	"os"

	ui "github.com/gizak/termui/v3"
	"github.com/j3-n/gyrio/internal/cli"
	"github.com/j3-n/gyrio/internal/controller"
	"github.com/j3-n/gyrio/internal/view"
)

func main() {
	app := cli.New()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	if err := ui.Init(); err != nil {
		log.Fatalf("Error initialising termui: %v", err)
	}
	defer ui.Close()

	view.Init()
	controller.Start()
}
