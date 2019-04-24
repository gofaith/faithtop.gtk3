package faithtop

import (
	"github.com/gotk3/gotk3/gtk"
)

type FMenuBar struct {
	FBaseView
	v *gtk.MenuBar
}

func MenuBar() *FMenuBar {
	v, _ := gtk.MenuBarNew()
	setupWidget(&v.Widget)
	fb := &FMenuBar{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	return fb
}
func GetMenuBarById(id string) *FMenuBar {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FMenuBar); ok {
			return b
		}
	}
	return nil
}

func (vh *ViewHolder) GetMenuBarByItemId(id string) *FMenuBar {
	if v, ok := vh.vlist[id]; ok {
		if bt, ok := v.(*FMenuBar); ok {
			return bt
		}
	}
	return nil
}

// ----------------------------------------------------------
func (v *FMenuBar) getBaseView() *FBaseView {
	return &v.FBaseView
}

func (v *FMenuBar) SetId(id string) *FMenuBar {
	idMap[id] = v
	return v
}
func (v *FMenuBar) SetItemId(parent *FListView, id string) *FMenuBar {
	if parent.vhs[parent.currentCreation].vlist == nil {
		parent.vhs[parent.currentCreation].vlist = make(map[string]IView)
	}
	parent.vhs[parent.currentCreation].vlist[id] = v
	return v
}
func (v *FMenuBar) Size(width, height int) *FMenuBar {
	parseSize(v, &v.v.Widget, width, height)
	return v
}
func (v *FMenuBar) OnClick(f func()) *FMenuBar {
	v.v.Connect("clicked", f)
	return v
}

func (v *FMenuBar) GravityCenter() *FMenuBar {
	v.FBaseView.GravityCenter()
	return v
}
func (v *FMenuBar) GravityEnd() *FMenuBar {
	v.FBaseView.GravityEnd()
	return v
}
func (v *FMenuBar) Disable() *FMenuBar {
	v.v.SetSensitive(false)
	return v
}
func (v *FMenuBar) Enable() *FMenuBar {
	v.v.SetSensitive(true)
	return v
}
func (v *FMenuBar) Visible() *FMenuBar {
	v.v.SetVisible(true)
	return v
}
func (v *FMenuBar) Invisible() *FMenuBar {
	v.v.SetVisible(false)
	return v
}
func (v *FMenuBar) MarginAll(i int) *FMenuBar {
	v.v.SetMarginBottom(i)
	v.v.SetMarginEnd(i)
	v.v.SetMarginStart(i)
	v.v.SetMarginTop(i)
	return v
}
func (v *FMenuBar) MarginLeft(i int) *FMenuBar {
	v.v.SetMarginStart(i)
	return v
}
func (v *FMenuBar) MarginTop(i int) *FMenuBar {
	v.v.SetMarginTop(i)
	return v
}
func (v *FMenuBar) MarginBottom(i int) *FMenuBar {
	v.v.SetMarginBottom(i)
	return v
}
func (v *FMenuBar) MarginRight(i int) *FMenuBar {
	v.v.SetMarginEnd(i)
	return v
}
func (v *FMenuBar) Tooltips(s string) *FMenuBar {
	v.v.SetTooltipText(s)
	return v
}
func (v *FMenuBar) Focus() *FMenuBar {
	currentFocus = v.widget
	if currentWin != nil {
		currentWin.SetFocus(v.widget)
	}
	return v
}

//====================================================================

func (f *FMenuBar) Menus(is ...IMenuItem) *FMenuBar {
	for _, i := range is {
		f.v.Append(i.getMenuItem())
	}
	return f
}
