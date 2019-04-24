package faithtop

import "github.com/gotk3/gotk3/gtk"

type IMenuItem interface {
	getMenuItem() *gtk.MenuItem
}

type FMenuItem struct {
	FBaseView
	v       *gtk.MenuItem
	subMenu *gtk.Menu
}

func (f *FMenuItem) getMenuItem() *gtk.MenuItem {
	return f.v
}

func MenuItem(s string) *FMenuItem {
	v, _ := gtk.MenuItemNewWithMnemonic(s)
	setupWidget(&v.Widget)
	fb := &FMenuItem{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	return fb
}

func (v *FMenuItem) OnClick(f func()) *FMenuItem {
	v.v.Connect("activate", f)
	return v
}
func (v *FMenuItem) getBaseView() *FBaseView {
	return &v.FBaseView
}

func (f *FMenuItem) SubMenu(ms ...IMenuItem) *FMenuItem {
	if f.subMenu == nil {
		f.subMenu, _ = gtk.MenuNew()
	}
	for _, m := range ms {
		f.subMenu.Append(m.getMenuItem())
	}
	f.v.SetSubmenu(f.subMenu)
	return f
}

func (v *FMenuItem) SetId(s string) *FMenuItem {
	idMap[s] = v
	return v
}
func GetMenuItemById(id string) *FMenuItem {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FMenuItem); ok {
			return b
		}
	}
	return nil
}

//===================================

type FCheckMenuItem struct {
	FBaseView
	v       *gtk.CheckMenuItem
	subMenu *gtk.Menu
}

func (v *FCheckMenuItem) getMenuItem() *gtk.MenuItem {
	return &v.v.MenuItem
}

func CheckMenuItem(s string) *FCheckMenuItem {
	v, _ := gtk.CheckMenuItemNewWithMnemonic(s)
	setupWidget(&v.Widget)
	fb := &FCheckMenuItem{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	return fb
}

func (v *FCheckMenuItem) OnChange(f func(bool)) *FCheckMenuItem {
	v.v.Connect("activate", func() {
		f(v.v.GetActive())
	})
	return v
}
func (v *FCheckMenuItem) getBaseView() *FBaseView {
	return &v.FBaseView
}

func (v *FCheckMenuItem) SubMenu(ms ...IMenuItem) *FCheckMenuItem {
	if v.subMenu == nil {
		v.subMenu,_ = gtk.MenuNew()
	}
	for _, m := range ms {
		v.subMenu.Append(m.getMenuItem())
	}
	v.v.SetSubmenu(v.subMenu)
	return v
}

func (f *FCheckMenuItem) Checked(b bool) *FCheckMenuItem {
	f.v.SetActive(b)
	return f
}

func (f *FCheckMenuItem) GetChecked() bool {
	return f.v.GetActive()
}


func GetCheckMenuItemById(id string) *FCheckMenuItem {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FCheckMenuItem); ok {
			return b
		}
	}
	return nil
}