package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gcla/gowid"
	"github.com/gcla/gowid/widgets/columns"
	"github.com/gcla/gowid/widgets/divider"
	"github.com/gcla/gowid/widgets/fill"
	"github.com/gcla/gowid/widgets/framed"
	"github.com/gcla/gowid/widgets/pile"
	"github.com/gcla/gowid/widgets/text"
	"github.com/gcla/gowid/widgets/vpadding"
	"github.com/mopp/gote/app"
	"github.com/mopp/gote/ui/editor"
	"github.com/mopp/gote/ui/titles"
	log "github.com/sirupsen/logrus"
)

func main() {
	{
		f := redirectLogger("gote.log")
		defer f.Close()
	}

	config := newConfig()

	notes := loadNotes(&config)

	// TODO: pass struct title { display_name string, path string } instead of string.
	editor := editor.New()

	var content *columns.Widget

	titles := titles.New(
		notes,
		func(n *app.Note, app gowid.IApp) {
			text, err := n.Read()
			if err != nil {
				log.Fatal(err)
			}

			editor.SetText(text, app)

			// TODO: define interface.
			content.SetFocus(app, 2)
		},
	)

	// TODO
	keywords := pile.New([]gowid.IContainerWidget{
		&gowid.ContainerWidget{IWidget: text.New("Keywords"), D: gowid.RenderWithWeight{W: 1}},
		&gowid.ContainerWidget{IWidget: divider.NewAscii(), D: gowid.RenderFlow{}},
		&gowid.ContainerWidget{IWidget: text.New("Relation"), D: gowid.RenderWithWeight{W: 1}},
	})

	vline := fill.New('|')
	content = columns.New([]gowid.IContainerWidget{
		&gowid.ContainerWidget{IWidget: vpadding.New(titles, gowid.VAlignTop{}, gowid.RenderFlow{}), D: gowid.RenderWithWeight{W: 1}},
		&gowid.ContainerWidget{IWidget: vline, D: gowid.RenderWithUnits{U: 1}},
		&gowid.ContainerWidget{IWidget: editor, D: gowid.RenderWithWeight{W: 2}},
		&gowid.ContainerWidget{IWidget: vline, D: gowid.RenderWithUnits{U: 1}},
		&gowid.ContainerWidget{IWidget: keywords, D: gowid.RenderWithWeight{W: 1}},
	})

	view := framed.New(content, framed.Options{
		Frame:       framed.AsciiFrame,
		TitleWidget: text.New("Gote"),
	})

	palette := gowid.Palette{
		"red":      gowid.MakePaletteEntry(gowid.ColorRed, gowid.ColorDarkBlue),
		"selected": gowid.MakePaletteEntry(gowid.ColorBlack, gowid.ColorDarkGray),
	}
	app, err := gowid.NewApp(gowid.AppArgs{
		View:    view,
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

func loadNotes(config *Config) []app.Note {
	files, err := ioutil.ReadDir(config.note_dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		log.Fatal(err)
	}

	notes := make([]app.Note, len(files))
	for i, file := range files {
		notes[i] = app.NewNote(config.note_dir, file)
	}

	return notes
}
