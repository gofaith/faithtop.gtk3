package faithtop

import (
	"github.com/gotk3/gotk3/gtk"
)

type FListView struct {
	FBaseView
	v               *gtk.ScrolledWindow
	lb              *gtk.ListBox
	vhs             []ViewHolder
	currentCreation int
	createView      func(*FListView) IView
	bindData        func(*ViewHolder, int)
	getCount        func() int
}
type ViewHolder struct {
	vlist map[string]IView
}

func VlistView(createView func(*FListView) IView, bindData func(*ViewHolder, int), getCount func() int) *FListView {
	fb := &FListView{}
	v, _ := gtk.ScrolledWindowNew(nil, nil)
	setupWidget(&v.Widget)
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	fb.createView = createView
	fb.bindData = bindData
	fb.getCount = getCount
	fb.v.SetPolicy(gtk.POLICY_NEVER, gtk.POLICY_AUTOMATIC)
	fb.lb, _ = gtk.ListBoxNew()
	fb.v.Add(fb.lb)
	for i := 0; i < getCount(); i++ {
		fb.currentCreation = i
		fb.vhs = append(fb.vhs, ViewHolder{vlist: make(map[string]IView)})
		fb.lb.Add(createView(fb).getBaseView().view)
	}
	fb.execBindData()
	return fb.Size(-2, -2)
}

func GetListViewById(id string) *FListView {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FListView); ok {
			return b
		}
	}
	return nil
}

func (vh *ViewHolder) GetListViewByItemId(id string) *FListView {
	if v, ok := vh.vlist[id]; ok {
		if bt, ok := v.(*FListView); ok {
			return bt
		}
	}
	return nil
}

// ---------------------------------------------------------

func (v *FListView) getBaseView() *FBaseView {
	return &v.FBaseView
}

func (v *FListView) SetId(id string) *FListView {
	idMap[id] = v
	return v
}
func (v *FListView) SetItemId(parent *FListView, id string) *FListView {
	if parent.vhs[parent.currentCreation].vlist == nil {
		parent.vhs[parent.currentCreation].vlist = make(map[string]IView)
	}
	parent.vhs[parent.currentCreation].vlist[id] = v
	return v
}
func (v *FListView) Size(width, height int) *FListView {
	parseSize(v, &v.v.Widget, width, height)
	return v
}
func (v *FListView) OnClick(f func()) *FListView {
	v.v.Connect("clicked", f)
	return v
}

func (v *FListView) GravityCenter() *FListView {
	v.FBaseView.GravityCenter()
	return v
}
func (v *FListView) GravityEnd() *FListView {
	v.FBaseView.GravityEnd()
	return v
}
func (v *FListView) Disable() *FListView {
	v.v.SetSensitive(false)
	return v
}
func (v *FListView) Enable() *FListView {
	v.v.SetSensitive(true)
	return v
}
func (v *FListView) Visible() *FListView {
	v.v.SetVisible(true)
	return v
}
func (v *FListView) Invisible() *FListView {
	v.v.SetVisible(false)
	return v
}
func (v *FListView) MarginAll(i int) *FListView {
	v.v.SetMarginBottom(i)
	v.v.SetMarginEnd(i)
	v.v.SetMarginStart(i)
	v.v.SetMarginTop(i)
	return v
}
func (v *FListView) MarginLeft(i int) *FListView {
	v.v.SetMarginStart(i)
	return v
}
func (v *FListView) MarginTop(i int) *FListView {
	v.v.SetMarginTop(i)
	return v
}
func (v *FListView) MarginBottom(i int) *FListView {
	v.v.SetMarginBottom(i)
	return v
}
func (v *FListView) MarginRight(i int) *FListView {
	v.v.SetMarginEnd(i)
	return v
}
func (v *FListView) Tooltips(s string) *FListView {
	v.v.SetTooltipText(s)
	return v
}
func (v *FListView) Focus() *FListView {
	currentFocus = v.widget
	if currentWin != nil {
		currentWin.SetFocus(v.widget)
	}
	return v
}

// ---------------------------------------------------------
func (v *FListView) execBindData() *FListView {
	for i := 0; i < v.getCount(); i++ {
		v.bindData(&v.vhs[i], i)
	}
	return v
}
