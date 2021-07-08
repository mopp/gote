package ui

import (
	"fmt"

	"github.com/gcla/gowid"
	"github.com/gcla/gowid/widgets/divider"
	"github.com/gcla/gowid/widgets/edit"
	"github.com/gcla/gowid/widgets/pile"
	"github.com/gcla/gowid/widgets/text"
	"github.com/gdamore/tcell"
	"github.com/mopp/gote/internal/gote"
	log "github.com/sirupsen/logrus"
	"golang.design/x/clipboard"
)

type editorWidget struct {
	*pile.Widget
	note *gote.Note
}

func newEditorWidget() *editorWidget {
	e := edit.New()
	statusLine := text.New("Started.")

	return &editorWidget{
		Widget: pile.New([]gowid.IContainerWidget{
			&gowid.ContainerWidget{IWidget: e, D: gowid.RenderWithWeight{W: 1}},
			&gowid.ContainerWidget{IWidget: divider.NewAscii(), D: gowid.RenderFlow{}},
			&gowid.ContainerWidget{IWidget: statusLine, D: gowid.RenderWithUnits{U: 1}},
		}),
	}
}

func (w *editorWidget) edit() *edit.Widget {
	return w.SubWidgets()[0].(*gowid.ContainerWidget).SubWidget().(*edit.Widget)
}

func (w *editorWidget) statusLine() *text.Widget {
	return w.SubWidgets()[2].(*gowid.ContainerWidget).SubWidget().(*text.Widget)
}

func (w *editorWidget) UserInput(ev interface{}, size gowid.IRenderSize, focus gowid.Selector, app gowid.IApp) bool {
	evk, ok := ev.(*tcell.EventKey)
	if !ok {
		return false
	}

	if evk.Key() == tcell.KeyCtrlS {
		err := w.note.WriteFrom(w.edit())
		if err != nil {
			w.statusLine().SetText(fmt.Sprintf("Error: %v", err), app)
		} else {
			w.statusLine().SetText("Saved.", app)
		}

		return true
	} else if evk.Key() == tcell.KeyCtrlC {
		clipboard.Write(clipboard.FmtText, []byte(w.edit().Text()))
		w.statusLine().SetText("Copied to clipboard.", app)

		return true
	}

	if w.note != nil {
		w.SetStatusLine(w.note.Title(), app)
	}

	return w.edit().UserInput(ev, size, focus, app)
}

func (w *editorWidget) OpenNote(note *gote.Note, app gowid.IApp) {
	text, err := note.Read()
	if err != nil {
		log.Fatal(err)
	}

	e := w.edit()
	e.SetText(text, app)
	e.SetCursorPos(0, app)

	w.note = note
	w.SetStatusLine(w.note.Title(), app)
}

func (w *editorWidget) SetStatusLine(text string, app gowid.IApp) {
	w.statusLine().SetText(w.note.Title(), app)
}
