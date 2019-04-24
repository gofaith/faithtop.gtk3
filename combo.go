package faithtop

import (
	"github.com/gotk3/gotk3/gtk"
)

type FCombo struct {
	FBaseView
	v                 *gtk.ComboBoxText
	strs              []string
	isActivatedByCode bool
}

func Combo() *FCombo {
	v, _ := gtk.ComboBoxTextNew()
	setupWidget(&v.Widget)
	fb := &FCombo{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	return fb
}

func GetComboById(id string) *FCombo {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FCombo); ok {
			return b
		}
	}
	return nil
}

func (vh *ViewHolder) GetComboByItemId(id string) *FCombo {
	if v, ok := vh.vlist[id]; ok {
		if bt, ok := v.(*FCombo); ok {
			return bt
		}
	}
	return nil
}

// ----------------------------------------------------------
func (v *FCombo) getBaseView() *FBaseView {
	return &v.FBaseView
}

func (v *FCombo) SetId(id string) *FCombo {
	idMap[id] = v
	return v
}
func (v *FCombo) SetItemId(parent *FListView, id string) *FCombo {
	if parent.vhs[parent.currentCreation].vlist == nil {
		parent.vhs[parent.currentCreation].vlist = make(map[string]IView)
	}
	parent.vhs[parent.currentCreation].vlist[id] = v
	return v
}
func (v *FCombo) Size(width, height int) *FCombo {
	parseSize(v, &v.v.Widget, width, height)
	return v
}
func (v *FCombo) OnClick(f func()) *FCombo {
	v.v.Connect("clicked", f)
	return v
}

func (v *FCombo) GravityCenter() *FCombo {
	v.FBaseView.GravityCenter()
	return v
}
func (v *FCombo) GravityEnd() *FCombo {
	v.FBaseView.GravityEnd()
	return v
}
func (v *FCombo) Disable() *FCombo {
	v.v.SetSensitive(false)
	return v
}
func (v *FCombo) Enable() *FCombo {
	v.v.SetSensitive(true)
	return v
}
func (v *FCombo) Visible() *FCombo {
	v.v.SetVisible(true)
	return v
}
func (v *FCombo) Invisible() *FCombo {
	v.v.SetVisible(false)
	return v
}
func (v *FCombo) MarginAll(i int) *FCombo {
	v.v.SetMarginBottom(i)
	v.v.SetMarginEnd(i)
	v.v.SetMarginStart(i)
	v.v.SetMarginTop(i)
	return v
}
func (v *FCombo) MarginLeft(i int) *FCombo {
	v.v.SetMarginStart(i)
	return v
}
func (v *FCombo) MarginTop(i int) *FCombo {
	v.v.SetMarginTop(i)
	return v
}
func (v *FCombo) MarginBottom(i int) *FCombo {
	v.v.SetMarginBottom(i)
	return v
}
func (v *FCombo) MarginRight(i int) *FCombo {
	v.v.SetMarginEnd(i)
	return v
}
func (v *FCombo) Tooltips(s string) *FCombo {
	v.v.SetTooltipText(s)
	return v
}
func (v *FCombo) Focus() *FCombo {
	currentFocus = v.widget
	if currentWin != nil {
		currentWin.SetFocus(v.widget)
	}
	return v
}

//====================================================================

func (v *FCombo) GetActiveText() string {
	return v.v.GetActiveText()
}
func (v *FCombo) GetAllText() []string {
	return v.strs
}
func (v *FCombo) TextList(ts ...string) *FCombo {
	for _, _ = range v.strs {
		v.v.Remove(0)
	}
	for _, t := range ts {
		v.v.AppendText(t)
	}
	v.strs = ts
	return v
}
func (f *FCombo) ActiveText(index int) *FCombo {
	f.isActivatedByCode = true
	f.v.SetActive(index)
	return f
}
func (v *FCombo) OnChange(f func(str string)) *FCombo {
	v.v.Connect("changed", func() {
		f(v.GetActiveText())
	})
	return v
}
