package ui

import (
	"fmt"

	"github.com/gcla/gowid"
	"github.com/gcla/gowid/widgets/columns"
	"github.com/gcla/gowid/widgets/dialog"
	"github.com/gcla/gowid/widgets/divider"
	"github.com/gcla/gowid/widgets/fill"
	"github.com/gcla/gowid/widgets/framed"
	"github.com/gcla/gowid/widgets/pile"
	"github.com/gcla/gowid/widgets/text"
	"github.com/gcla/gowid/widgets/vpadding"
	"github.com/gdamore/tcell"
	"github.com/mopp/gote/internal/gote"
)

type MainWidget struct {
	*framed.Widget
	service    *gote.Service
	titles     *titlesWidget
	editor     *editorWidget
	content    *columns.Widget
	quitDialog *dialog.Widget
}

func NewMainWidget(service *gote.Service, config *gote.Config) (*MainWidget, error) {
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
		func(note *gote.Note, app gowid.IApp) {
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
		&gowid.ContainerWidget{IWidget: editor, D: gowid.RenderWithWeight{W: 3}},
		&gowid.ContainerWidget{IWidget: vline, D: gowid.RenderWithUnits{U: 1}},
		&gowid.ContainerWidget{IWidget: keywords, D: gowid.RenderWithWeight{W: 2}},
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

func (w *MainWidget) UserInput(ev interface{}, size gowid.IRenderSize, focus gowid.Selector, app gowid.IApp) bool {
	evk, ok := ev.(*tcell.EventKey)

	if !ok {
		return w.Widget.UserInput(ev, size, focus, app)
	}

	k := evk.Key()
	if k == tcell.KeyCtrlN {
		w.createNewNote(app)
		return true
	} else if k == tcell.KeyCtrlD {
		w.findOrCreateDailyNoteToday(app)
		return true
	} else if k == tcell.KeyCtrlG {
		w.setFocusOnTitles(app)
		return true
	} else if k == tcell.KeyCtrlC || k == tcell.KeyESC {
		if w.quitDialog == nil {
			w.confirmQuit(app)
		}
		return true
	}

	return w.Widget.UserInput(ev, size, focus, app)
}

func (w *MainWidget) setFocusOnTitles(app gowid.IApp) {
	w.content.SetFocus(app, 0)
}

func (w *MainWidget) setFocusOnEditor(app gowid.IApp) {
	w.content.SetFocus(app, 0)
}

func (w *MainWidget) createNewNote(app gowid.IApp) {
	onCreate := func(app gowid.IApp, widget gowid.IWidget, name string) {
		n := w.service.CreateNote(name)
		w.titles.AddNote(n, app)
	}

	newCreateDialogWidget(onCreate).
		Open(w, gowid.RenderWithRatio{R: 0.5}, app)
}

func (w *MainWidget) findOrCreateDailyNoteToday(app gowid.IApp) {
	n, err := w.service.FindDailyNoteToday()
	if err != nil {
		msg := fmt.Sprintf("could not find daily note: %v", err)
		w.editor.SetStatusLine(msg, app)
		return
	}

	if n != nil {
		// Already exist.
		w.editor.OpenNote(n, app)
		w.setFocusOnEditor(app)
		return
	}

	n, err = w.service.CreateDailyNoteTody()

	if err != nil {
		msg := fmt.Sprintf("could not create daily note: %v", err)
		w.editor.SetStatusLine(msg, app)
		return
	}

	w.titles.AddNote(n, app)
	w.editor.OpenNote(n, app)
	w.setFocusOnEditor(app)

	return
}

func (w *MainWidget) confirmQuit(app gowid.IApp) {
	w.quitDialog = dialog.New(
		framed.NewSpace(text.New("Do you want to quit?")),
		dialog.Options{
			Buttons: []dialog.Button{
				{
					Msg: "Quit",
					Action: func(app gowid.IApp, widget gowid.IWidget) {
						app.Quit()
					},
				},
				{
					Msg: "Cancel",
					Action: func(app gowid.IApp, widget gowid.IWidget) {
						w.quitDialog.Close(app)
						w.quitDialog = nil
					},
				},
			},
		},
	)

	w.quitDialog.Open(w, gowid.RenderWithRatio{R: 0.5}, app)
}
