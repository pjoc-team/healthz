// Code generated by "stringer -type Status"; DO NOT EDIT.

package api

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Up-1]
	_ = x[Loading-5]
	_ = x[Down-9]
}

const (
	_Status_name_0 = "Up"
	_Status_name_1 = "Loading"
	_Status_name_2 = "Down"
)

func (i Status) String() string {
	switch {
	case i == 1:
		return _Status_name_0
	case i == 5:
		return _Status_name_1
	case i == 9:
		return _Status_name_2
	default:
		return "Status(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
