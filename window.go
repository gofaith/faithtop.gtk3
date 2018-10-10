package faithtop

import (
	"github.com/gotk3/gotk3/gtk"
)

var (
	windowCounter int
)

type FWindow struct {
	v         *gtk.Window
	showAfter bool
}

func Win() *FWindow {
	w, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	setupWindow(w)
	fw := &FWindow{}
	fw.v = w
	return fw
}
func PopupWin() *FWindow {
	w, _ := gtk.WindowNew(gtk.WINDOW_POPUP)
	setupWindow(w)
	fw := &FWindow{}
	fw.v = w
	return fw
}
func TopWin() *FWindow {
	return Win().Top()
}
func TopPopupWin() *FWindow {
	return PopupWin().Top()
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
func (v *FWindow) SetId(id string) *FWindow {
	idMap[id] = v
	return v
}
func GetWinById(id string) *FWindow {
	v, ok := idMap[id]
	if ok {
		if w, is := v.(*FWindow); is {
			return w
		}
	}
	return nil
}
func (v *FWindow) Size(width, height int) *FWindow {
	v.v.SetDefaultSize(width, height)
	return v
}
func (v *FWindow) Append(i IView) *FWindow {
	v.v.Add(i.getBaseView().view)
	if v.showAfter {
		v.Show()
	}
	return v
}
func (v *FWindow) Show() *FWindow {
	v.v.ShowAll()
	if windowCounter == 1 {
		gtk.Main()
	}
	return v
}
func (v *FWindow) Close() *FWindow {
	v.v.Close()
	return v
}
func (v *FWindow) Top() *FWindow {
	v.v.SetModal(true)
	v.v.SetKeepAbove(true)
	return v
}
func (v *FWindow) DeferShow() *FWindow {
	v.showAfter = true
	return v
}
func (v *FWindow) VBox(is ...IView) *FWindow {
	v.Append(VBox().Size(-2, -2).Append(is...))
	return v
}
func (v *FWindow) HBox(is ...IView) *FWindow {
	v.Append(HBox().Size(-2, -2).Append(is...))
	return v
}

// func (v *FWindow) Modal() *FWindow {
// 	v.v.SetModal(true)
// 	return v
// }
// func (v *FWindow) UnModal() *FWindow {
// 	v.v.SetModal(false)
// 	return v
// }
// func (v *FWindow) IsModal() bool {
// 	return v.v.GetModal()
// }
// func (v *FWindow) KeepAbove() *FWindow {
// 	v.v.SetKeepAbove(true)
// 	return v
// }
// func (v *FWindow) UnKeepAbove() *FWindow {
// 	v.v.SetKeepAbove(false)
// 	return v
// }
