package faithtop

import (
	"github.com/gotk3/gotk3/gtk"
)

type FStatusBar struct {
	FBaseView
	v *gtk.Statusbar
}

func StatusBar() *FStatusBar {
	v, _ := gtk.StatusbarNew()
	setupWidget(&v.Widget)
	fb := &FStatusBar{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	return fb
}
func GetStatusBarById(id string) *FStatusBar {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FStatusBar); ok {
			return b
		}
	}
	return nil
}

func (vh *ViewHolder) GetStatusBarByItemId(id string) *FStatusBar {
	if v, ok := vh.vlist[id]; ok {
		if bt, ok := v.(*FStatusBar); ok {
			return bt
		}
	}
	return nil
}

// ----------------------------------------------------------
func (v *FStatusBar) getBaseView() *FBaseView {
	return &v.FBaseView
}

func (v *FStatusBar) SetId(id string) *FStatusBar {
	idMap[id] = v
	return v
}
func (v *FStatusBar) SetItemId(parent *FListView, id string) *FStatusBar {
	if parent.vhs[parent.currentCreation].vlist == nil {
		parent.vhs[parent.currentCreation].vlist = make(map[string]IView)
	}
	parent.vhs[parent.currentCreation].vlist[id] = v
	return v
}
func (v *FStatusBar) Size(width, height int) *FStatusBar {
	parseSize(v, &v.v.Widget, width, height)
	return v
}
func (v *FStatusBar) OnClick(f func()) *FStatusBar {
	v.v.Connect("clicked", f)
	return v
}

func (v *FStatusBar) GravityCenter() *FStatusBar {
	v.FBaseView.GravityCenter()
	return v
}
func (v *FStatusBar) GravityEnd() *FStatusBar {
	v.FBaseView.GravityEnd()
	return v
}
func (v *FStatusBar) Disable() *FStatusBar {
	v.v.SetSensitive(false)
	return v
}
func (v *FStatusBar) Enable() *FStatusBar {
	v.v.SetSensitive(true)
	return v
}
func (v *FStatusBar) Visible() *FStatusBar {
	v.v.SetVisible(true)
	return v
}
func (v *FStatusBar) Invisible() *FStatusBar {
	v.v.SetVisible(false)
	return v
}
func (v *FStatusBar) MarginAll(i int) *FStatusBar {
	v.v.SetMarginBottom(i)
	v.v.SetMarginEnd(i)
	v.v.SetMarginStart(i)
	v.v.SetMarginTop(i)
	return v
}
func (v *FStatusBar) MarginLeft(i int) *FStatusBar {
	v.v.SetMarginStart(i)
	return v
}
func (v *FStatusBar) MarginTop(i int) *FStatusBar {
	v.v.SetMarginTop(i)
	return v
}
func (v *FStatusBar) MarginBottom(i int) *FStatusBar {
	v.v.SetMarginBottom(i)
	return v
}
func (v *FStatusBar) MarginRight(i int) *FStatusBar {
	v.v.SetMarginEnd(i)
	return v
}
func (v *FStatusBar) Tooltips(s string) *FStatusBar {
	v.v.SetTooltipText(s)
	return v
}
func (v *FStatusBar) Focus() *FStatusBar {
	currentFocus = v.widget
	if currentWin != nil {
		currentWin.SetFocus(v.widget)
	}
	return v
}

//====================================================================

func (v *FStatusBar) Text(s string) *FStatusBar {
	contextID := v.v.GetContextId("context_description")
	v.v.Push(contextID, s)
	return v
}
