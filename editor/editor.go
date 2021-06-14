package editor

import (
	"github.com/gcla/gowid"
	"github.com/gcla/gowid/widgets/divider"
	"github.com/gcla/gowid/widgets/edit"
	"github.com/gcla/gowid/widgets/pile"
	"github.com/gcla/gowid/widgets/text"
	"github.com/gdamore/tcell"
	log "github.com/sirupsen/logrus"
)

type Widget struct {
	*pile.Widget
}

func New() *Widget {
	e := edit.New()
	statusLine := text.New("Started.")

	return &Widget{
		Widget: pile.New([]gowid.IContainerWidget{
			&gowid.ContainerWidget{IWidget: e, D: gowid.RenderWithWeight{W: 1}},
			&gowid.ContainerWidget{IWidget: divider.NewAscii(), D: gowid.RenderFlow{}},
			&gowid.ContainerWidget{IWidget: statusLine, D: gowid.RenderWithUnits{U: 1}},
		}),
	}
}

func (w *Widget) edit() *edit.Widget {
	return w.SubWidgets()[0].(*gowid.ContainerWidget).SubWidget().(*edit.Widget)
}

func (w *Widget) statusLine() *text.Widget {
	return w.SubWidgets()[2].(*gowid.ContainerWidget).SubWidget().(*text.Widget)
}

func (w *Widget) UserInput(ev interface{}, size gowid.IRenderSize, focus gowid.Selector, app gowid.IApp) bool {
	evk, ok := ev.(*tcell.EventKey)
	if !ok {
		return false
	}

	if evk.Key() == tcell.KeyCtrlS {
		w.statusLine().SetText("Saved.", app)
		// TODO: Implement callback.
		log.Info("Saved.")
		log.Info(w.edit().Text())

		return true
	}

	// TODO: Clear status line using chanel.
	w.statusLine().SetText("", app)

	return w.edit().UserInput(ev, size, focus, app)
}

func (w *Widget) SetText(text string, app gowid.IApp) {
	w.edit().SetText(text, app)
}
