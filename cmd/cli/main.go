package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/j3-n/gyrio/internal/controller"
	"github.com/j3-n/gyrio/internal/view"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("Error initialising termui: %v", err)
	}
	defer ui.Close()

	view.Init()
	controller.Start()
}
