package faithtop

import (
	"github.com/gotk3/gotk3/gtk"
)

type FScroll struct {
	FBaseView
	v   *gtk.ScrolledWindow
	box *FBox
}

func newScroll() *FScroll {
	v, _ := gtk.ScrolledWindowNew(nil, nil)
	setupWidget(&v.Widget)
	fb := &FScroll{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	return fb
}
func Scroll(child IView) *FScroll {
	fb := newScroll()
	fb.v.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	fb.v.Add(child.getBaseView().widget)
	return fb
}
func VScroll() *FScroll {
	fb := newScroll()
	fb.v.SetPolicy(gtk.POLICY_NEVER, gtk.POLICY_AUTOMATIC)
	fb.box = VBox().Size(-2, -2)
	fb.v.Add(fb.box.widget)
	return fb
}

func HScroll() *FScroll {
	fb := newScroll()
	fb.v.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_NEVER)
	fb.box = HBox().Size(-2, -2)
	fb.v.Add(fb.box.widget)
	return fb
}

func GetScrollById(id string) *FScroll {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FScroll); ok {
			return b
		}
	}
	return nil
}

func (vh *ViewHolder) GetScrollByItemId(id string) *FScroll {
	if v, ok := vh.vlist[id]; ok {
		if bt, ok := v.(*FScroll); ok {
			return bt
		}
	}
	return nil
}

// ----------------------------------------------------------
func (v *FScroll) getBaseView() *FBaseView {
	return &v.FBaseView
}

func (v *FScroll) SetId(id string) *FScroll {
	idMap[id] = v
	return v
}
func (v *FScroll) SetItemId(parent *FListView, id string) *FScroll {
	if parent.vhs[parent.currentCreation].vlist == nil {
		parent.vhs[parent.currentCreation].vlist = make(map[string]IView)
	}
	parent.vhs[parent.currentCreation].vlist[id] = v
	return v
}
func (v *FScroll) Size(width, height int) *FScroll {
	parseSize(v, &v.v.Widget, width, height)
	return v
}
func (v *FScroll) OnClick(f func()) *FScroll {
	v.v.Connect("clicked", f)
	return v
}

func (v *FScroll) GravityCenter() *FScroll {
	v.FBaseView.GravityCenter()
	return v
}
func (v *FScroll) GravityEnd() *FScroll {
	v.FBaseView.GravityEnd()
	return v
}
func (v *FScroll) Disable() *FScroll {
	v.v.SetSensitive(false)
	return v
}
func (v *FScroll) Enable() *FScroll {
	v.v.SetSensitive(true)
	return v
}
func (v *FScroll) Visible() *FScroll {
	v.v.SetVisible(true)
	return v
}
func (v *FScroll) Invisible() *FScroll {
	v.v.SetVisible(false)
	return v
}
func (v *FScroll) MarginAll(i int) *FScroll {
	v.v.SetMarginBottom(i)
	v.v.SetMarginEnd(i)
	v.v.SetMarginStart(i)
	v.v.SetMarginTop(i)
	return v
}
func (v *FScroll) MarginLeft(i int) *FScroll {
	v.v.SetMarginStart(i)
	return v
}
func (v *FScroll) MarginTop(i int) *FScroll {
	v.v.SetMarginTop(i)
	return v
}
func (v *FScroll) MarginBottom(i int) *FScroll {
	v.v.SetMarginBottom(i)
	return v
}
func (v *FScroll) MarginRight(i int) *FScroll {
	v.v.SetMarginEnd(i)
	return v
}
func (v *FScroll) Tooltips(s string) *FScroll {
	v.v.SetTooltipText(s)
	return v
}
func (v *FScroll) Focus() *FScroll {
	currentFocus = v.widget
	if currentWin != nil {
		currentWin.SetFocus(v.widget)
	}
	return v
}

//====================================================================
func (v *FScroll) Append(is ...IView) *FScroll {
	if v.box == nil {
		return v
	}
	v.box.Append(is...)
	return v
}
