package faithtop

import (
	"github.com/gotk3/gotk3/gtk"
)

type FBox struct {
	FBaseView
	v *gtk.Box
}

func (v *FBox) getBaseView() *FBaseView {
	return &v.FBaseView
}
func GetBoxById(id string) *FBox {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FBox); ok {
			return b
		}
	}
	return nil
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

func (v *FBox) Size(width, height int) *FBox {
	parseSize(v, &v.v.Widget, width, height)
	return v
}
func (v *FBox) GravityStart() *FBox {
	v.FBaseView.GravityStart()
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

// ---------------------------------------------------------

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
