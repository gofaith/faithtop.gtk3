package faithtop

import (
	"github.com/gotk3/gotk3/gtk"
)

type FBox struct {
	FBaseView
	v *gtk.Box
}

func VBox() *FBox {
	v, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	setupWidget(&v.Widget)
	fb := &FBox{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	return fb
}
func HBox() *FBox {
	v, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	setupWidget(&v.Widget)
	fb := &FBox{}
	fb.v = v
	fb.view = v
	return fb
}

func GetBoxById(id string) *FBox {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FBox); ok {
			return b
		}
	}
	return nil
}
func (vh *ViewHolder) GetBoxByItemId(id string) *FBox {
	if v, ok := vh.vlist[id]; ok {
		if bt, ok := v.(*FBox); ok {
			return bt
		}
	}
	return nil
}

// ---------------------------------------------------------

func (v *FBox) getBaseView() *FBaseView {
	return &v.FBaseView
}

func (v *FBox) SetId(id string) *FBox {
	idMap[id] = v
	return v
}
func (v *FBox) SetItemId(parent *FListView, id string) *FBox {
	if parent.vhs[parent.currentCreation].vlist == nil {
		parent.vhs[parent.currentCreation].vlist = make(map[string]IView)
	}
	parent.vhs[parent.currentCreation].vlist[id] = v
	return v
}
func (v *FBox) Size(width, height int) *FBox {
	parseSize(v, &v.v.Widget, width, height)
	return v
}
func (v *FBox) OnClick(f func()) *FBox {
	v.v.Connect("clicked", f)
	return v
}

func (v *FBox) GravityCenter() *FBox {
	v.FBaseView.GravityCenter()
	return v
}
func (v *FBox) GravityEnd() *FBox {
	v.FBaseView.GravityEnd()
	return v
}
func (v *FBox) Disable() *FBox {
	v.v.SetSensitive(false)
	return v
}
func (v *FBox) Enable() *FBox {
	v.v.SetSensitive(true)
	return v
}
func (v *FBox) Visible() *FBox {
	v.v.SetVisible(true)
	return v
}
func (v *FBox) Invisible() *FBox {
	v.v.SetVisible(false)
	return v
}
func (v *FBox) MarginAll(i int) *FBox {
	v.v.SetMarginBottom(i)
	v.v.SetMarginEnd(i)
	v.v.SetMarginStart(i)
	v.v.SetMarginTop(i)
	return v
}
func (v *FBox) MarginLeft(i int) *FBox {
	v.v.SetMarginStart(i)
	return v
}
func (v *FBox) MarginTop(i int) *FBox {
	v.v.SetMarginTop(i)
	return v
}
func (v *FBox) MarginBottom(i int) *FBox {
	v.v.SetMarginBottom(i)
	return v
}
func (v *FBox) MarginRight(i int) *FBox {
	v.v.SetMarginEnd(i)
	return v
}
func (v *FBox) Tooltips(s string) *FBox {
	v.v.SetTooltipText(s)
	return v
}
func (v *FBox) Focus() *FBox {
	currentFocus = v.widget
	if currentWin != nil {
		currentWin.SetFocus(v.widget)
	}
	return v
}

// ---------------------------------------------------------

func (v *FBox) Append(is ...IView) *FBox {
	var fs []func()
	for _, i := range is {
		v.v.Add(i.getBaseView().view)
		i.getBaseView().alreadyAdded = true
		if i.getBaseView().afterAppend != nil {
			fs = append(fs, i.getBaseView().afterAppend)
			i.getBaseView().afterAppend = nil
		}
	}
	for _, f := range fs {
		f()
	}
	return v
}
