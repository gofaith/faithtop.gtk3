package faithtop

import (
	"github.com/gotk3/gotk3/gtk"
)

type FBox struct {
	FBaseView
	v *gtk.Box
}

func (v *FBox) GetBaseView() *FBaseView {
	return &v.FBaseView
}
func VBox() *FBox {
	v, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	setupWidget(&v.Widget)
	fb := &FBox{}
	fb.v = v
	fb.view = v
	return fb
}
func HBox() *FBox {
	v, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	setupWidget(&v.Widget)
	fb := &FBox{}
	fb.v = v
	fb.view = v
	return fb
}
func (v *FBox) Append(is ...IView) *FBox {
	for _, i := range is {
		v.v.Add(i.GetBaseView().view)
	}
	return v
}

func (v *FBox) Size(width, height int) *FBox {
	parseSize(v, &v.v.Widget, width, height)
	return v
}
