package faithtop

import (
	"github.com/gotk3/gotk3/gtk"
)

type FEdit struct {
	FBaseView
	v *gtk.Entry
}

func Edit() *FEdit {
	v, _ := gtk.EntryNew()
	setupWidget(&v.Widget)
	fb := &FEdit{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	return fb
}

func GetEditById(id string) *FEdit {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FEdit); ok {
			return b
		}
	}
	return nil
}

func (vh *ViewHolder) GetEditByItemId(id string) *FEdit {
	if v, ok := vh.vlist[id]; ok {
		if bt, ok := v.(*FEdit); ok {
			return bt
		}
	}
	return nil
}

// ----------------------------------------------------------
func (v *FEdit) getBaseView() *FBaseView {
	return &v.FBaseView
}

func (v *FEdit) SetId(id string) *FEdit {
	idMap[id] = v
	return v
}
func (v *FEdit) SetItemId(parent *FListView, id string) *FEdit {
	if parent.vhs[parent.currentCreation].vlist == nil {
		parent.vhs[parent.currentCreation].vlist = make(map[string]IView)
	}
	parent.vhs[parent.currentCreation].vlist[id] = v
	return v
}
func (v *FEdit) Size(width, height int) *FEdit {
	parseSize(v, &v.v.Widget, width, height)
	return v
}
func (v *FEdit) OnClick(f func()) *FEdit {
	v.v.Connect("clicked", f)
	return v
}

func (v *FEdit) GravityCenter() *FEdit {
	v.FBaseView.GravityCenter()
	return v
}
func (v *FEdit) GravityEnd() *FEdit {
	v.FBaseView.GravityEnd()
	return v
}
func (v *FEdit) Disable() *FEdit {
	v.v.SetSensitive(false)
	return v
}
func (v *FEdit) Enable() *FEdit {
	v.v.SetSensitive(true)
	return v
}
func (v *FEdit) Visible() *FEdit {
	v.v.SetVisible(true)
	return v
}
func (v *FEdit) Invisible() *FEdit {
	v.v.SetVisible(false)
	return v
}
func (v *FEdit) MarginAll(i int) *FEdit {
	v.v.SetMarginBottom(i)
	v.v.SetMarginEnd(i)
	v.v.SetMarginStart(i)
	v.v.SetMarginTop(i)
	return v
}
func (v *FEdit) MarginLeft(i int) *FEdit {
	v.v.SetMarginStart(i)
	return v
}
func (v *FEdit) MarginTop(i int) *FEdit {
	v.v.SetMarginTop(i)
	return v
}
func (v *FEdit) MarginBottom(i int) *FEdit {
	v.v.SetMarginBottom(i)
	return v
}
func (v *FEdit) MarginRight(i int) *FEdit {
	v.v.SetMarginEnd(i)
	return v
}
func (v *FEdit) Tooltips(s string) *FEdit {
	v.v.SetTooltipText(s)
	return v
}
func (v *FEdit) Focus() *FEdit {
	currentFocus = v.widget
	if currentWin != nil {
		currentWin.SetFocus(v.widget)
	}
	return v
}

//====================================================================

func (v *FEdit) Text(t string) *FEdit {
	v.v.SetText(t)
	return v
}
func (v *FEdit) GetText() string {
	s, _ := v.v.GetText()
	return s
}
func (v *FEdit) Editable(b bool) *FEdit {
	v.v.SetEditable(b)
	return v
}
func (v *FEdit) MaxLength(l int) *FEdit {
	v.v.SetMaxLength(l)
	return v
}
func (v *FEdit) TextVisible(b bool) *FEdit {
	v.v.SetVisibility(b)
	return v
}
func (v *FEdit) OnEnter(f func()) *FEdit {
	v.v.Connect("activate", f)
	return v
}
