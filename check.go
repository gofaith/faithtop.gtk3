package faithtop

import (
	"github.com/gotk3/gotk3/gtk"
)

type FCheck struct {
	FBaseView
	v *gtk.CheckButton
}

func Check() *FCheck {
	v, _ := gtk.CheckButtonNew()
	setupWidget(&v.Widget)
	fb := &FCheck{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	return fb
}

func GetCheckById(id string) *FCheck {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FCheck); ok {
			return b
		}
	}
	return nil
}

func (vh *ViewHolder) GetCheckByItemId(id string) *FCheck {
	if v, ok := vh.vlist[id]; ok {
		if bt, ok := v.(*FCheck); ok {
			return bt
		}
	}
	return nil
}

// ----------------------------------------------------------
func (v *FCheck) getBaseView() *FBaseView {
	return &v.FBaseView
}

func (v *FCheck) SetId(id string) *FCheck {
	idMap[id] = v
	return v
}
func (v *FCheck) SetItemId(parent *FListView, id string) *FCheck {
	if parent.vhs[parent.currentCreation].vlist == nil {
		parent.vhs[parent.currentCreation].vlist = make(map[string]IView)
	}
	parent.vhs[parent.currentCreation].vlist[id] = v
	return v
}
func (v *FCheck) Size(width, height int) *FCheck {
	parseSize(v, &v.v.Widget, width, height)
	return v
}
func (v *FCheck) OnClick(f func()) *FCheck {
	v.v.Connect("clicked", f)
	return v
}

func (v *FCheck) GravityCenter() *FCheck {
	v.FBaseView.GravityCenter()
	return v
}
func (v *FCheck) GravityEnd() *FCheck {
	v.FBaseView.GravityEnd()
	return v
}
func (v *FCheck) Disable() *FCheck {
	v.v.SetSensitive(false)
	return v
}
func (v *FCheck) Enable() *FCheck {
	v.v.SetSensitive(true)
	return v
}
func (v *FCheck) Visible() *FCheck {
	v.v.SetVisible(true)
	return v
}
func (v *FCheck) Invisible() *FCheck {
	v.v.SetVisible(false)
	return v
}
func (v *FCheck) MarginAll(i int) *FCheck {
	v.v.SetMarginBottom(i)
	v.v.SetMarginEnd(i)
	v.v.SetMarginStart(i)
	v.v.SetMarginTop(i)
	return v
}
func (v *FCheck) MarginLeft(i int) *FCheck {
	v.v.SetMarginStart(i)
	return v
}
func (v *FCheck) MarginTop(i int) *FCheck {
	v.v.SetMarginTop(i)
	return v
}
func (v *FCheck) MarginBottom(i int) *FCheck {
	v.v.SetMarginBottom(i)
	return v
}
func (v *FCheck) MarginRight(i int) *FCheck {
	v.v.SetMarginEnd(i)
	return v
}
func (v *FCheck) Tooltips(s string) *FCheck {
	v.v.SetTooltipText(s)
	return v
}
func (v *FCheck) Focus() *FCheck {
	currentFocus = v.widget
	if currentWin != nil {
		currentWin.SetFocus(v.widget)
	}
	return v
}

//====================================================================

func (f *FCheck) Text(s string) *FCheck {
	f.v.SetLabel(s)
	return f
}

func (f *FCheck) GetText() string {
	s, _ := f.v.GetLabel()
	return s
}

func (f *FCheck) Checked(b bool) *FCheck {
	f.v.SetActive(b)
	return f
}

func (f *FCheck) OnChange(fn func(bool)) *FCheck {
	if fn == nil {
		return f
	}
	f.v.Connect("toggled", fn)
	return f
}

func (f *FCheck) GetChecked() bool {
	return f.v.GetActive()
}

