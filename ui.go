package faithtop

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type IView interface {
	getBaseView() *FBaseView
}

func setupWidget(w *gtk.Widget) {
	w.SetHAlign(gtk.ALIGN_START)
	w.SetVAlign(gtk.ALIGN_START)
	w.SetHExpand(false)
	w.SetVExpand(false)
}

type FBaseView struct {
	view         gtk.IWidget
	widget       *gtk.Widget
	gravity      int // default:0 , 1,2,3, 4,5,6, 7,8,9
	afterAppend  func()
	alreadyAdded bool
}

func init() {
	gtk.Init(nil)
}

var (
	idMap = make(map[string]interface{})
)

func RunOnUIThread(f func()) {
	glib.IdleAdd(f)
}

func parseSize(iv IView, widget *gtk.Widget, width, height int) {
	v := iv.getBaseView()
	align := v.gravity % 3
	if width < 0 {
		if width == -1 {
			widget.SetHExpand(false)
			if align == 2 {
				widget.SetHAlign(gtk.ALIGN_END)
			} else if align == 1 {
				widget.SetHAlign(gtk.ALIGN_CENTER)
			} else {
				widget.SetHAlign(gtk.ALIGN_START)
			}
		} else if width == -2 {
			widget.SetHExpand(true)
			widget.SetHAlign(gtk.ALIGN_FILL)
		}
		width = 0
	}
	if height < 0 {
		if height == -1 {
			widget.SetVExpand(false)
			if align == 2 {
				widget.SetVAlign(gtk.ALIGN_END)
			} else if align == 1 {
				widget.SetVAlign(gtk.ALIGN_CENTER)
			} else {
				widget.SetVAlign(gtk.ALIGN_START)
			}
		} else if height == -2 {
			widget.SetVExpand(true)
			widget.SetVAlign(gtk.ALIGN_FILL)
		}
		height = 0
	}
	widget.SetSizeRequest(width, height)
}
func (v *FBaseView) GravityStart() {
	v.gravity = 3
	v.widget.SetHAlign(gtk.ALIGN_START)
	v.widget.SetVAlign(gtk.ALIGN_START)
}
func (v *FBaseView) GravityCenter() {
	v.gravity = 4
	v.widget.SetHAlign(gtk.ALIGN_CENTER)
	v.widget.SetVAlign(gtk.ALIGN_CENTER)
}
func (v *FBaseView) GravityEnd() {
	v.gravity = 5
	v.widget.SetHAlign(gtk.ALIGN_END)
	v.widget.SetVAlign(gtk.ALIGN_END)
}
func (v *FBaseView) IsEnabled() bool {
	return v.widget.GetSensitive()
}
func (v *FBaseView) IsVisible() bool {
	return v.widget.GetVisible()
}
func (v *FBaseView) GetWidth() int {
	al := v.widget.GetAllocation()
	return al.GetWidth()
}
func (v *FBaseView) GetHeight() int {
	return v.widget.GetAllocation().GetHeight()
}
