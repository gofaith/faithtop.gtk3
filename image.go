package faithtop

import (
	"github.com/gotk3/gotk3/gtk"
)

type FImage struct {
	FBaseView
	v *gtk.Image
}

func Image() *FImage {
	v, _ := gtk.ImageNew()
	setupWidget(&v.Widget)
	fb := &FImage{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	return fb
}

func GetImageById(id string) *FImage {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FImage); ok {
			return b
		}
	}
	return nil
}

// ----------------------------------------------------------
func (v *FImage) getBaseView() *FBaseView {
	return &v.FBaseView
}

func (v *FImage) SetId(id string) *FImage {
	idMap[id] = v
	return v
}
func (v *FImage) Size(width, height int) *FImage {
	parseSize(v, &v.v.Widget, width, height)
	return v
}
func (v *FImage) OnClick(f func()) *FImage {
	v.v.Connect("clicked", f)
	return v
}

func (v *FImage) GravityCenter() *FImage {
	v.FBaseView.GravityCenter()
	return v
}
func (v *FImage) GravityEnd() *FImage {
	v.FBaseView.GravityEnd()
	return v
}
func (v *FImage) Disable() *FImage {
	v.v.SetSensitive(false)
	return v
}
func (v *FImage) Enable() *FImage {
	v.v.SetSensitive(true)
	return v
}
func (v *FImage) Visible() *FImage {
	v.v.SetVisible(true)
	return v
}
func (v *FImage) Invisible() *FImage {
	v.v.SetVisible(false)
	return v
}
func (v *FImage) MarginAll(i int) *FImage {
	v.v.SetMarginBottom(i)
	v.v.SetMarginEnd(i)
	v.v.SetMarginStart(i)
	v.v.SetMarginTop(i)
	return v
}
func (v *FImage) MarginLeft(i int) *FImage {
	v.v.SetMarginStart(i)
	return v
}
func (v *FImage) MarginTop(i int) *FImage {
	v.v.SetMarginTop(i)
	return v
}
func (v *FImage) MarginBottom(i int) *FImage {
	v.v.SetMarginBottom(i)
	return v
}
func (v *FImage) MarginRight(i int) *FImage {
	v.v.SetMarginEnd(i)
	return v
}
func (v *FImage) Tooltips(s string) *FImage {
	v.v.SetTooltipText(s)
	return v
}
func (v *FImage) Focus() *FImage {
	currentFocus = v.widget
	if currentWin != nil {
		currentWin.SetFocus(v.widget)
	}
	return v
}

//====================================================================
func (v *FImage) Src(url string) *FImage {
	if StartsWidth(url, "/") {
		v.v.SetFromFile(url)
	} else if StartsWidth(url, "file://") {
		v.v.SetFromFile(url[len("file://"):])
	} else if StartsWidth(url, "http") {
		CacheNetFile(url, GetCacheDir(), func(fpath string) {
			RunOnUIThread(func() {
				v.v.SetFromFile(fpath)
			})
		})
	}
	return v
}
