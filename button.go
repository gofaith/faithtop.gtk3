package faithtop

import (
	"github.com/gotk3/gotk3/gtk"
)

type FButton struct {
	v *gtk.Button
}

func (v *FButton) GetView() gtk.IWidget {
	return v.v
}
func Button() *FButton {
	v, _ := gtk.ButtonNew()
	setupWidget(&v.Widget)
	fb := &FButton{}
	fb.v = v
	return fb
}
func (v *FButton) Text(t string) *FButton {
	v.v.SetLabel(t)
	return v
}
func (v *FButton) OnClick(f func()) *FButton {
	v.v.Connect("clicked", f)
	return v
}
