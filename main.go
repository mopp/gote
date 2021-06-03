package main

import (
	"fmt"
	"os"

	"github.com/gcla/gowid"
	"github.com/gcla/gowid/widgets/columns"
	"github.com/gcla/gowid/widgets/divider"
	"github.com/gcla/gowid/widgets/edit"
	"github.com/gcla/gowid/widgets/fill"
	"github.com/gcla/gowid/widgets/framed"
	"github.com/gcla/gowid/widgets/pile"
	"github.com/gcla/gowid/widgets/text"
	"github.com/gcla/gowid/widgets/vpadding"
	log "github.com/sirupsen/logrus"
)

func main() {
	text1 := text.New("hello world")

	titles := pile.New([]gowid.IContainerWidget{
		&gowid.ContainerWidget{IWidget: text1, D: gowid.RenderWithWeight{W: 1}},
	})

	editor := edit.New(edit.Options{Text: "abcde"})
	statusLine := text.New("status line")

	main := pile.New([]gowid.IContainerWidget{
		&gowid.ContainerWidget{IWidget: editor, D: gowid.RenderWithWeight{W: 1}},
		&gowid.ContainerWidget{IWidget: divider.NewAscii(), D: gowid.RenderFlow{}},
		&gowid.ContainerWidget{IWidget: statusLine, D: gowid.RenderWithUnits{U: 1}},
	})

	keywords := pile.New([]gowid.IContainerWidget{
		&gowid.ContainerWidget{IWidget: text.New("Keywords"), D: gowid.RenderWithWeight{W: 1}},
		&gowid.ContainerWidget{IWidget: divider.NewAscii(), D: gowid.RenderFlow{}},
		&gowid.ContainerWidget{IWidget: text.New("Relation"), D: gowid.RenderWithWeight{W: 1}},
	})

	vline := fill.New('|')
	content := columns.New([]gowid.IContainerWidget{
		&gowid.ContainerWidget{IWidget: vpadding.New(titles, gowid.VAlignTop{}, gowid.RenderFlow{}), D: gowid.RenderWithWeight{W: 1}},
		&gowid.ContainerWidget{IWidget: vline, D: gowid.RenderWithUnits{U: 1}},
		&gowid.ContainerWidget{IWidget: main, D: gowid.RenderWithWeight{W: 2}},
		&gowid.ContainerWidget{IWidget: vline, D: gowid.RenderWithUnits{U: 1}},
		&gowid.ContainerWidget{IWidget: keywords, D: gowid.RenderWithWeight{W: 1}},
	})

	view := framed.New(content, framed.Options{
		Frame:       framed.AsciiFrame,
		TitleWidget: text.New("Gote"),
	})

	app, err := gowid.NewApp(gowid.AppArgs{
		View: view,
		Log:  log.StandardLogger(),
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	app.SimpleMainLoop()
}
