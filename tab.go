package faithtop

import (
	"github.com/gotk3/gotk3/gtk"
)

type FTabLayout struct {
	FBaseView
	v *gtk.Notebook
}

func TabLayout() *FTabLayout {
	v, _ := gtk.NotebookNew()
	setupWidget(&v.Widget)
	fb := &FTabLayout{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	return fb
}
func GetTabLayoutById(id string) *FTabLayout {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FTabLayout); ok {
			return b
		}
	}
	return nil
}

func (vh *ViewHolder) GetTabLayoutByItemId(id string) *FTabLayout {
	if v, ok := vh.vlist[id]; ok {
		if bt, ok := v.(*FTabLayout); ok {
			return bt
		}
	}
	return nil
}

// ----------------------------------------------------------
func (v *FTabLayout) getBaseView() *FBaseView {
	return &v.FBaseView
}

func (v *FTabLayout) SetId(id string) *FTabLayout {
	idMap[id] = v
	return v
}
func (v *FTabLayout) SetItemId(parent *FListView, id string) *FTabLayout {
	if parent.vhs[parent.currentCreation].vlist == nil {
		parent.vhs[parent.currentCreation].vlist = make(map[string]IView)
	}
	parent.vhs[parent.currentCreation].vlist[id] = v
	return v
}
func (v *FTabLayout) Size(width, height int) *FTabLayout {
	parseSize(v, &v.v.Widget, width, height)
	return v
}
func (v *FTabLayout) OnClick(f func()) *FTabLayout {
	v.v.Connect("clicked", f)
	return v
}

func (v *FTabLayout) GravityCenter() *FTabLayout {
	v.FBaseView.GravityCenter()
	return v
}
func (v *FTabLayout) GravityEnd() *FTabLayout {
	v.FBaseView.GravityEnd()
	return v
}
func (v *FTabLayout) Disable() *FTabLayout {
	v.v.SetSensitive(false)
	return v
}
func (v *FTabLayout) Enable() *FTabLayout {
	v.v.SetSensitive(true)
	return v
}
func (v *FTabLayout) Visible() *FTabLayout {
	v.v.SetVisible(true)
	return v
}
func (v *FTabLayout) Invisible() *FTabLayout {
	v.v.SetVisible(false)
	return v
}
func (v *FTabLayout) MarginAll(i int) *FTabLayout {
	v.v.SetMarginBottom(i)
	v.v.SetMarginEnd(i)
	v.v.SetMarginStart(i)
	v.v.SetMarginTop(i)
	return v
}
func (v *FTabLayout) MarginLeft(i int) *FTabLayout {
	v.v.SetMarginStart(i)
	return v
}
func (v *FTabLayout) MarginTop(i int) *FTabLayout {
	v.v.SetMarginTop(i)
	return v
}
func (v *FTabLayout) MarginBottom(i int) *FTabLayout {
	v.v.SetMarginBottom(i)
	return v
}
func (v *FTabLayout) MarginRight(i int) *FTabLayout {
	v.v.SetMarginEnd(i)
	return v
}
func (v *FTabLayout) Tooltips(s string) *FTabLayout {
	v.v.SetTooltipText(s)
	return v
}
func (v *FTabLayout) Focus() *FTabLayout {
	currentFocus = v.widget
	if currentWin != nil {
		currentWin.SetFocus(v.widget)
	}
	return v
}

//====================================================================
type FTab struct {
	title   *gtk.Label
	content IView
}

func Tab(title string, view IView) *FTab {
	fb := &FTab{}
	fb.title, _ = gtk.LabelNew(title)
	fb.content = view
	return fb
}

func (f *FTabLayout) Tabs(ps ...*FTab) *FTabLayout {
	for _, p := range ps {
		f.v.AppendPage(p.content.getBaseView().view, p.title)
	}
	return f
}

func (f *FTabLayout) OnSwitchPage(fn func()) *FTabLayout {
	f.v.Connect("switch-page", fn)
	return f
}
