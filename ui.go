package faithtop

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
)

type IView interface {
	GetBaseView() *FBaseView
}

func setupWidget(w *gtk.Widget) {
	w.SetHAlign(gtk.ALIGN_START)
	w.SetVAlign(gtk.ALIGN_START)
	w.SetHExpand(false)
	w.SetVExpand(false)
}

type FBaseView struct {
	view    gtk.IWidget
	gravity int // default:0 , 1,2,3, 4,5,6, 7,8,9
}

func parseSize(iv IView, widget *gtk.Widget, width, height int) {
	v := iv.GetBaseView()
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
				fmt.Println("HAlign start")
			}
		} else if width == -2 {
			widget.SetHExpand(true)
			widget.SetHAlign(gtk.ALIGN_FILL)
			fmt.Println("width -2")
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
				fmt.Println("VAlign start")
			}
		} else if height == -2 {
			widget.SetVExpand(true)
			widget.SetVAlign(gtk.ALIGN_FILL)
			fmt.Println("height -2")
		}
		height = 0
	}
	widget.SetSizeRequest(width, height)
}
