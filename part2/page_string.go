// Code generated by "stringer -type=Page"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Ya-0]
	_ = x[RBT-1]
	_ = x[Eldorado-2]
	_ = x[MVideo-3]
}

const _Page_name = "YaRBTEldoradoMVideo"

var _Page_index = [...]uint8{0, 2, 5, 13, 19}

func (i Page) String() string {
	if i < 0 || i >= Page(len(_Page_index)-1) {
		return "Page(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Page_name[_Page_index[i]:_Page_index[i+1]]
}
