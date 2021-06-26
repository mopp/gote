package main

import (
	"os"

	"github.com/gcla/gowid"
	"github.com/mopp/gote/internal/gote"
	"github.com/mopp/gote/internal/ui"
	log "github.com/sirupsen/logrus"
)

func main() {
	logFile := redirectLogger("gote.log")
	defer logFile.Close()

	config := gote.NewConfig()
	service := gote.NewService(config)
	mainWidget, err := ui.NewMainWidget(service, config)

	if err != nil {
		log.Fatal(err)
	}

	palette := gowid.Palette{
		"red":      gowid.MakePaletteEntry(gowid.ColorRed, gowid.ColorDarkBlue),
		"selected": gowid.MakePaletteEntry(gowid.ColorBlack, gowid.ColorDarkGray),
	}

	app, err := gowid.NewApp(gowid.AppArgs{
		View:    mainWidget,
		Palette: &palette,
		Log:     log.StandardLogger(),
	})

	if err != nil {
		log.Fatal(err)
	}

	app.SimpleMainLoop()
}

func redirectLogger(path string) *os.File {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}

	log.SetOutput(f)

	return f
}
