// Code generated by "stringer -type=SignatureKind -output=stringer_generated.go"; DO NOT EDIT.

package compiler

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SignatureKindCall-0]
	_ = x[SignatureKindConstruct-1]
}

const _SignatureKind_name = "SignatureKindCallSignatureKindConstruct"

var _SignatureKind_index = [...]uint8{0, 17, 39}

func (i SignatureKind) String() string {
	if i < 0 || i >= SignatureKind(len(_SignatureKind_index)-1) {
		return "SignatureKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SignatureKind_name[_SignatureKind_index[i]:_SignatureKind_index[i+1]]
}
