package faithtop

import (
	"github.com/gotk3/gotk3/gtk"
)

type FButton struct {
	FBaseView
	v *gtk.Button
}

func (v *FButton) getBaseView() *FBaseView {
	return &v.FBaseView
}

func GetButtonById(id string) *FButton {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FButton); ok {
			return b
		}
	}
	return nil
}
func Button() *FButton {
	v, _ := gtk.ButtonNew()
	setupWidget(&v.Widget)
	fb := &FButton{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	return fb
}
func (v *FButton) SetId(id string) *FButton {
	idMap[id] = v
	return v
}
func (v *FButton) Size(width, height int) *FButton {
	parseSize(v, &v.v.Widget, width, height)
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

func (v *FButton) GravityCenter() *FButton {
	v.FBaseView.GravityCenter()
	return v
}
func (v *FButton) GravityEnd() *FButton {
	v.FBaseView.GravityEnd()
	return v
}
