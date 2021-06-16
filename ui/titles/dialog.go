package titles

import (
	"fmt"

	"github.com/gcla/gowid"
	"github.com/gcla/gowid/widgets/dialog"
	"github.com/gcla/gowid/widgets/edit"
	"github.com/gcla/gowid/widgets/holder"
	"github.com/gcla/gowid/widgets/pile"
	"github.com/gcla/gowid/widgets/text"
	log "github.com/sirupsen/logrus"
)

type createDialogWidget struct {
	*dialog.Widget
}

func newCreateDialogWidget() *createDialogWidget {
	e := edit.New()
	p := pile.New([]gowid.IContainerWidget{
		&gowid.ContainerWidget{IWidget: text.New("Create new note:"), D: gowid.RenderWithWeight{W: 1}},
		&gowid.ContainerWidget{IWidget: e, D: gowid.RenderFlow{}},
	})

	var d *dialog.Widget
	d = dialog.New(
		holder.New(p),
		dialog.Options{
			FocusOnWidget: true,
			Buttons: []dialog.Button{
				{
					Msg: "Create",
					Action: func(app gowid.IApp, widget gowid.IWidget) {
						log.Info(fmt.Sprintf("%+v", e.Text()))
						// TODO: Add create new note.
						// service.createNote()
						d.Close(app)
					},
				},
				dialog.Cancel,
			},
		},
	)

	return &createDialogWidget{
		Widget: d,
	}
}
