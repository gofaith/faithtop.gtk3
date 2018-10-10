package faithtop

import (
	"github.com/gotk3/gotk3/gtk"
)

type FBox struct {
	v *gtk.Box
}

func (v *FBox) GetView() gtk.IWidget {
	return v.v
}
func VBox() *FBox {
	v, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	setupWidget(&v.Widget)
	fb := &FBox{}
	fb.v = v
	return fb
}
func (v *FBox) Append(is ...IView) *FBox {
	for _, i := range is {
		v.v.PackStart(i.GetView(), false, false, 0)
	}
	return v
}
