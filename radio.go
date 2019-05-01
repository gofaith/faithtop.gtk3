package faithtop

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type FRadio struct {
	FBaseView
	v               *gtk.RadioButton
	groupID         string
	onGroupSelected func(string)
}

var groupIDMap = make(map[string]*FRadio)

func Radio(grouID string) *FRadio {
	var group *glib.SList
	if radio, ok := groupIDMap[grouID]; ok {
		group, _ = radio.v.GetGroup()
	}
	v, _ := gtk.RadioButtonNew(group)
	setupWidget(&v.Widget)
	fb := &FRadio{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	fb.groupID = grouID
	if group == nil {
		groupIDMap[grouID] = fb
	}
	fb.v.Connect("toggled", func() {
		if fb.v.GetActive() && groupIDMap[grouID].onGroupSelected != nil {
			label, _ := fb.v.GetLabel()
			groupIDMap[grouID].onGroupSelected(label)
		}
	})
	return fb
}
func GetRadioById(id string) *FRadio {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FRadio); ok {
			return b
		}
	}
	return nil
}

func (vh *ViewHolder) GetRadioByItemId(id string) *FRadio {
	if v, ok := vh.vlist[id]; ok {
		if bt, ok := v.(*FRadio); ok {
			return bt
		}
	}
	return nil
}

// ----------------------------------------------------------
func (v *FRadio) getBaseView() *FBaseView {
	return &v.FBaseView
}

func (v *FRadio) SetId(id string) *FRadio {
	idMap[id] = v
	return v
}
func (v *FRadio) SetItemId(parent *FListView, id string) *FRadio {
	if parent.vhs[parent.currentCreation].vlist == nil {
		parent.vhs[parent.currentCreation].vlist = make(map[string]IView)
	}
	parent.vhs[parent.currentCreation].vlist[id] = v
	return v
}
func (v *FRadio) Size(width, height int) *FRadio {
	parseSize(v, &v.v.Widget, width, height)
	return v
}
func (v *FRadio) OnClick(f func()) *FRadio {
	v.v.Connect("clicked", f)
	return v
}

func (v *FRadio) GravityCenter() *FRadio {
	v.FBaseView.GravityCenter()
	return v
}
func (v *FRadio) GravityEnd() *FRadio {
	v.FBaseView.GravityEnd()
	return v
}
func (v *FRadio) Disable() *FRadio {
	v.v.SetSensitive(false)
	return v
}
func (v *FRadio) Enable() *FRadio {
	v.v.SetSensitive(true)
	return v
}
func (v *FRadio) Visible() *FRadio {
	v.v.SetVisible(true)
	return v
}
func (v *FRadio) Invisible() *FRadio {
	v.v.SetVisible(false)
	return v
}
func (v *FRadio) MarginAll(i int) *FRadio {
	v.v.SetMarginBottom(i)
	v.v.SetMarginEnd(i)
	v.v.SetMarginStart(i)
	v.v.SetMarginTop(i)
	return v
}
func (v *FRadio) MarginLeft(i int) *FRadio {
	v.v.SetMarginStart(i)
	return v
}
func (v *FRadio) MarginTop(i int) *FRadio {
	v.v.SetMarginTop(i)
	return v
}
func (v *FRadio) MarginBottom(i int) *FRadio {
	v.v.SetMarginBottom(i)
	return v
}
func (v *FRadio) MarginRight(i int) *FRadio {
	v.v.SetMarginEnd(i)
	return v
}
func (v *FRadio) Tooltips(s string) *FRadio {
	v.v.SetTooltipText(s)
	return v
}
func (v *FRadio) Focus() *FRadio {
	currentFocus = v.widget
	if currentWin != nil {
		currentWin.SetFocus(v.widget)
	}
	return v
}

//====================================================================

func (v *FRadio) Text(t string) *FRadio {
	v.v.SetLabel(t)
	return v
}
func (v *FRadio) GetText() string {
	s, _ := v.v.GetLabel()
	return s
}
func (f *FRadio) Selected() *FRadio {
	f.v.SetActive(true)
	return f
}
func (f *FRadio) OnGroupSelected(fn func(string)) *FRadio {
	groupIDMap[f.groupID].onGroupSelected = fn
	return f
}
