package faithtop

import (
	"github.com/gotk3/gotk3/gtk"
)

type FButton struct {
	FBaseView
	v *gtk.Button
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
func GetButtonById(id string) *FButton {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FButton); ok {
			return b
		}
	}
	return nil
}

// ----------------------------------------------------------
func (v *FButton) getBaseView() *FBaseView {
	return &v.FBaseView
}

func (v *FButton) SetId(id string) *FButton {
	idMap[id] = v
	return v
}
func (v *FButton) Size(width, height int) *FButton {
	parseSize(v, &v.v.Widget, width, height)
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
func (v *FButton) Disable() *FButton {
	v.v.SetSensitive(false)
	return v
}
func (v *FButton) Enable() *FButton {
	v.v.SetSensitive(true)
	return v
}
func (v *FButton) Visible() *FButton {
	v.v.SetVisible(true)
	return v
}
func (v *FButton) Invisible() *FButton {
	v.v.SetVisible(false)
	return v
}
func (v *FButton) MarginAll(i int) *FButton {
	v.v.SetMarginBottom(i)
	v.v.SetMarginEnd(i)
	v.v.SetMarginStart(i)
	v.v.SetMarginTop(i)
	return v
}
func (v *FButton) MarginLeft(i int) *FButton {
	v.v.SetMarginStart(i)
	return v
}
func (v *FButton) MarginTop(i int) *FButton {
	v.v.SetMarginTop(i)
	return v
}
func (v *FButton) MarginBottom(i int) *FButton {
	v.v.SetMarginBottom(i)
	return v
}
func (v *FButton) MarginRight(i int) *FButton {
	v.v.SetMarginEnd(i)
	return v
}
func (v *FButton) Tooltips(s string) *FButton {
	v.v.SetTooltipText(s)
	return v
}
func (v *FButton) Focus() *FButton {
	currentFocus = v.widget
	if currentWin != nil {
		currentWin.SetFocus(v.widget)
	}
	return v
}

//====================================================================

func (v *FButton) Text(t string) *FButton {
	v.v.SetLabel(t)
	return v
}
func (v *FButton) GetText() string {
	s, _ := v.v.GetLabel()
	return s
}
