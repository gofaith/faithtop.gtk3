package faithtop

import (
	"github.com/gotk3/gotk3/gtk"
)

func init() {
	gtk.Init(nil)
}

var (
	idMap = make(map[string]interface{})
)
