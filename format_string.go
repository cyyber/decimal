// Code generated by "stringer -type=format"; DO NOT EDIT.

package decimal

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[normal-0]
	_ = x[plain-1]
	_ = x[sci-2]
}

const _format_name = "normalplainsci"

var _format_index = [...]uint8{0, 6, 11, 14}

func (i format) String() string {
	if i >= format(len(_format_index)-1) {
		return "format(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _format_name[_format_index[i]:_format_index[i+1]]
}
