// Code generated by "stringer -type=Op"; DO NOT EDIT.

package suite

import "strconv"

const _Op_name = "AddSubMulDivFMASqrtRemRFICFFCFICIFCFDCDFQuietCmpSigCmpCopyNegAbsCopySignScalbLogbNextAfterClassIsSignedIsNormalIsInfIsZeroIsSubNormalIsNaNIsSignalingIsFiniteMinNumMaxNumMinNumMagMaxNumMagSameQuantumQuantizeNextUpNextDownEquivSetRatSignSignbitExpLogLog10PowIntDivNormalizeRoundToIntShift"

var _Op_index = [...]uint16{0, 3, 6, 9, 12, 15, 19, 22, 25, 28, 31, 34, 37, 40, 48, 54, 58, 61, 64, 72, 77, 81, 90, 95, 103, 111, 116, 122, 133, 138, 149, 157, 163, 169, 178, 187, 198, 206, 212, 220, 225, 231, 235, 242, 245, 248, 253, 256, 262, 271, 281, 286}

func (i Op) String() string {
	if i >= Op(len(_Op_index)-1) {
		return "Op(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Op_name[_Op_index[i]:_Op_index[i+1]]
}
