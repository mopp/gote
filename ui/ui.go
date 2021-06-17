package ui

import (
	"fmt"

	"github.com/gcla/gowid"
	"github.com/gcla/gowid/widgets/columns"
	"github.com/gcla/gowid/widgets/divider"
	"github.com/gcla/gowid/widgets/fill"
	"github.com/gcla/gowid/widgets/framed"
	"github.com/gcla/gowid/widgets/pile"
	"github.com/gcla/gowid/widgets/text"
	"github.com/gcla/gowid/widgets/vpadding"
	"github.com/gdamore/tcell"
	"github.com/mopp/gote/app"
)

type MainWidget struct {
	*framed.Widget
	service *app.Service
	titles  *titlesWidget
	editor  *editorWidget
	content *columns.Widget
}

func NewMainWidget(service *app.Service, config *app.Config) (*MainWidget, error) {
	var titles *titlesWidget
	var editor *editorWidget
	var content *columns.Widget

	editor = newEditorWidget()

	notes, err := service.FetchAllNotes()
	if err != nil {
		return nil, err
	}

	titles = newTitlesWidget(
		notes,
		func(note *app.Note, app gowid.IApp) {
			editor.OpenNote(note, app)
			// TODO: Define struct and method to change focus.
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

	return &MainWidget{
		Widget: framed.New(
			content,
			framed.Options{
				Frame:       framed.AsciiFrame,
				TitleWidget: text.New("Gote"),
			},
		),
		service: service,
		titles:  titles,
		editor:  editor,
		content: content,
	}, nil
}

func (w *MainWidget) Created(n *app.Note) {

}

func (w *MainWidget) Updated(n *app.Note) {

}

func (w *MainWidget) UserInput(ev interface{}, size gowid.IRenderSize, focus gowid.Selector, app gowid.IApp) bool {
	evk, ok := ev.(*tcell.EventKey)
	r := evk.Rune()

	if !ok {
		goto skip
	}

	if evk.Key() == tcell.KeyCtrlN || r == 'N' {
		d := newCreateDialogWidget(func(app gowid.IApp, widget gowid.IWidget, name string) {
			n := w.service.CreateNote(name)
			w.titles.AddNote(n, app)
		})
		d.Open(w, gowid.RenderWithRatio{R: 0.5}, app)

		return true
	} else if evk.Key() == tcell.KeyCtrlD {
		n, err := w.service.FindDailyNoteToday()
		if err != nil {
			msg := fmt.Sprintf("could not find daily note: %v", err)
			w.editor.SetStatusLine(msg, app)
			return true
		}

		if n != nil {
			// Already exist.
			w.editor.OpenNote(n, app)
			// TODO: Define struct and method to change focus.
			w.content.SetFocus(app, 2)
			return true
		}

		n, err = w.service.CreateDailyNoteTody()

		if err != nil {
			msg := fmt.Sprintf("could not create daily note: %v", err)
			w.editor.SetStatusLine(msg, app)
			return true
		}

		w.titles.AddNote(n, app)
		w.editor.OpenNote(n, app)
		// TODO: Define struct and method to change focus.
		w.content.SetFocus(app, 2)

		return true
	}

skip:
	return w.Widget.UserInput(ev, size, focus, app)
}
