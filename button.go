package faithtop

import (
	"github.com/gotk3/gotk3/gtk"
)

type FButton struct {
	FBaseView
	v *gtk.Button
}

func (v *FButton) GetBaseView() *FBaseView {
	return &v.FBaseView
}
func Button() *FButton {
	v, _ := gtk.ButtonNew()
	setupWidget(&v.Widget)
	fb := &FButton{}
	fb.v = v
	fb.view = v
	return fb
}
func (v *FButton) Size(width, height int) *FButton {
	v.v.SetSizeRequest(width, height)
	return v
}
func (v *FButton) Text(t string) *FButton {
	v.v.SetLabel(t)
	return v
}
func (v *FButton) OnClick(f func()) *FButton {
	v.v.Connect("clicked", f)
	return v
}
