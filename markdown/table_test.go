package markdown

import "testing"

type LaunchServices struct {
	Executables *[]interface{}
	Arguments   *[]interface{}
	Delays      *[]interface{}
	EnabledGrp  *[]interface{}
	EnabledGrp1 *[]interface{}
}

func TestMDTable(t *testing.T) {
	ls := LaunchServices{
		Executables: &[]interface{}{"PATH_OF_SERVICE", "L", "exe", true, 1, 2, 3},
		Arguments:   &[]interface{}{"ARGUMENTS", "L", "flag", "f"},
		Delays:      &[]interface{}{"DELAY", "M", 1, 2, 3, 4, 5, 6, 7, 100},
		EnabledGrp:  &[]interface{}{"ENABLED", "", true, false, true, false, true, true},
		EnabledGrp1: &[]interface{}{"", "M", true, false, true, false, true},
	}
	MDTable(&ls, "test.md")
}
