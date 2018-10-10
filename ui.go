package faithtop

import (
	"github.com/gotk3/gotk3/gtk"
)

type IView interface {
	GetView() gtk.IWidget
}

func setupWidget(w *gtk.Widget) {
	w.SetHAlign(gtk.ALIGN_START)
	w.SetVAlign(gtk.ALIGN_START)
	w.SetHExpand(false)
	w.SetVExpand(false)
}