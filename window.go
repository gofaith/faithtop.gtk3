package faithtop

import (
	"github.com/gotk3/gotk3/gtk"
)

var (
	windowCounter int
)

type FWindow struct {
	v *gtk.Window
}

func Window() *FWindow {
	w, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	setupWindow(w)
	fw := &FWindow{}
	fw.v = w
	return fw
}
func setupWindow(w *gtk.Window) {
	windowCounter++
	w.SetDefaultSize(600, 400)
	w.SetPosition(gtk.WIN_POS_CENTER)
	w.Connect("destroy", func() {
		windowCounter--
		if windowCounter < 1 {
			gtk.MainQuit()
		}
	})
}
func (v *FWindow) Size(width, height int) *FWindow {
	v.v.SetDefaultSize(width, height)
	return v
}
func (v *FWindow) Append(i IView) *FWindow {
	v.v.Add(i.GetView())
	return v
}
func (v *FWindow) Show() *FWindow {
	v.v.ShowAll()
	if windowCounter == 1 {
		gtk.Main()
	}
	return v
}
