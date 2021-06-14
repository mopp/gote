package titles

import (
	"github.com/gcla/gowid"
	"github.com/gcla/gowid/widgets/isselected"
	"github.com/gcla/gowid/widgets/list"
	"github.com/gcla/gowid/widgets/styled"
	"github.com/gcla/gowid/widgets/text"
	"github.com/gdamore/tcell"
	"github.com/mopp/gote/app"
)

type OnNoteSelected func(*app.Note, gowid.IApp)

type Widget struct {
	*list.Widget
	notes          []app.Note
	onNoteSelected OnNoteSelected
}

func New(notes []app.Note, f OnNoteSelected) *Widget {
	ws := make([]gowid.IWidget, len(notes))

	for i, n := range notes {
		ws[i] = createTitleText(n.String())
	}

	walker := list.NewSimpleListWalker(ws)

	return &Widget{
		Widget:         list.New(walker),
		notes:          notes,
		onNoteSelected: f,
	}
}

func (w *Widget) UserInput(ev interface{}, size gowid.IRenderSize, focus gowid.Selector, app gowid.IApp) bool {
	evk, ok := ev.(*tcell.EventKey)
	if !ok {
		return false
	}

	r := evk.Rune()
	walker := w.Walker().(*list.SimpleListWalker)
	current := walker.Focus()
	if evk.Key() == tcell.KeyUp || r == 'k' {
		if current == walker.First() {
			walker.SetFocus(walker.Last(), app)
		} else {
			walker.SetFocus(walker.Previous(current), app)
		}

		return true
	} else if evk.Key() == tcell.KeyDown || r == 'j' {
		if current == walker.Last() {
			walker.SetFocus(walker.First(), app)
		} else {
			walker.SetFocus(walker.Next(current), app)
		}

		return true
	} else if evk.Key() == tcell.KeyEnter {
		note := &w.notes[current.(list.ListPos).ToInt()]
		w.onNoteSelected(note, app)

		return true
	}

	return false
}

func createTitleText(title string) *isselected.Widget {
	t := text.New(title)
	focused := styled.New(t, gowid.MakePaletteRef("selected"))

	return isselected.New(t, nil, focused)
}
