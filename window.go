package faithtop

import (
	"github.com/gotk3/gotk3/gtk"
)

var (
	currentWin   *gtk.Window
	currentFocus *gtk.Widget
)

type FWindow struct {
	v           *gtk.Window
	showAfter   bool
	ondestroyFn func()
}

func Win() *FWindow {
	w, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	fw := &FWindow{}
	fw.v = w
	setupWindow(fw)
	return fw
}
func PopupWin() *FWindow {
	w, _ := gtk.WindowNew(gtk.WINDOW_POPUP)
	fw := &FWindow{}
	fw.v = w
	setupWindow(fw)
	return fw
}
func TopWin() *FWindow {
	return Win().Top()
}
func TopPopupWin() *FWindow {
	return PopupWin().Top()
}
func setupWindow(fw *FWindow) {
	fw.v.SetDefaultSize(600, 400)
	fw.v.SetPosition(gtk.WIN_POS_CENTER)
	fw.v.Connect("destroy", func() {
		if fw.ondestroyFn != nil {
			fw.ondestroyFn()
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
func (v *FWindow) Add(i IView) *FWindow {
	v.v.Add(i.getBaseView().view)
	i.getBaseView().alreadyAdded = true
	if i.getBaseView().afterAppend != nil {
		i.getBaseView().afterAppend()
		i.getBaseView().afterAppend = nil
	}
	if v.showAfter {
		v.Show()
	}
	return v
}
func (v *FWindow) Show() *FWindow {
	currentWin = v.v
	v.v.ShowAll()
	if currentFocus != nil {
		v.v.SetFocus(currentFocus)
		currentFocus = nil
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
	v.Add(VBox().Size(-2, -2).Append(is...))
	return v
}
func (v *FWindow) HBox(is ...IView) *FWindow {
	v.Add(HBox().Size(-2, -2).Append(is...))
	return v
}
func (v *FWindow) Resizable(b bool) *FWindow {
	v.v.SetResizable(b)
	return v
}

func (f *FWindow) OnDestroy(fn func()) *FWindow {
	f.ondestroyFn = fn
	return f
}

func (f *FWindow) OnCloseClicked(fn func()) *FWindow {
	if fn == nil {
		return f
	}
	f.v.Connect("delete-event", fn)
	return f
}

func Main() {
	gtk.Main()
}
func MainQuit() {
	gtk.MainQuit()
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
