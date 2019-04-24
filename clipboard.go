package faithtop

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type FClipboard struct {
	v *gtk.Clipboard
}

func Clipboard() *FClipboard {
	f := &FClipboard{}
	display, _ := gdk.DisplayGetDefault()
	f.v, _ = gtk.ClipboardGetForDisplay(display, gdk.SELECTION_CLIPBOARD)
	return f
}

func (f *FClipboard) GetText() string {
	t, _ := f.v.WaitForText()
	return t
}

func (f *FClipboard) SetText(t string) *FClipboard {
	f.v.SetText(t)
	return f
}
