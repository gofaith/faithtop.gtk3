package faithtop

import (
	"github.com/gotk3/gotk3/gtk"
)

type FScroll struct {
	FBaseView
	v *gtk.ScrolledWindow
}

func Scroll() *FScroll {
	v, _ := gtk.ScrolledWindowNew(nil, nil)
	setupWidget(&v.Widget)
	fb := &FScroll{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	return fb
}

func VScroll() *FScroll {
	v, _ := gtk.ScrolledWindowNew(nil, nil)
	setupWidget(&v.Widget)
	fb := &FScroll{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	fb.v.SetPolicy(gtk.POLICY_NEVER, gtk.POLICY_AUTOMATIC)
	return fb
}

func HScroll() *FScroll {
	v, _ := gtk.ScrolledWindowNew(nil, nil)
	setupWidget(&v.Widget)
	fb := &FScroll{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	fb.v.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_NEVER)
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

// ----------------------------------------------------------
func (v *FScroll) getBaseView() *FBaseView {
	return &v.FBaseView
}

func (v *FScroll) SetId(id string) *FScroll {
	idMap[id] = v
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
	for _, i := range is {
		v.v.Add(i.getBaseView().view)
	}
	return v
}
