package overlay

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/relate"
)

// LineMerge returns a Geometry containing the LineMerges.
//	or an empty atomic geometry, or an empty GEOMETRYCOLLECTION
func LineMerge(ml matrix.Collection) matrix.Collection {
	for i := 0; i < len(ml)-1; i++ {
		for j := 1; j < len(ml); j++ {
			if mlMerge, ok := MergeLine(ml, i, j); ok {
				ml = mlMerge
				return LineMerge(ml)
			}
		}
	}
	return ml
}

// MergeLine  Computes the Merge of two geometries,either or both of which may be null.
func MergeLine(ml matrix.Collection, i, j int) (matrix.Collection, bool) {

	if ml[i] == nil && ml[j] == nil {
		return ml, false
	}
	if ml[i] == nil {
		return ml, false
	}

	if ml[j] == nil {
		return ml, false
	}
	var result matrix.Collection
	mark, ips := relate.IntersectionEdge(ml[i].(matrix.LineMatrix), ml[j].(matrix.LineMatrix))
	if mark {
		for _, v := range ips {
			if _, inVer := relate.InLineVertex(v.Matrix, ml[i].(matrix.LineMatrix)); inVer {
				r1 := mergeCheck(ml[i].(matrix.LineMatrix), ml[j].(matrix.LineMatrix))
				if r1 != nil {
					if i < j {
						temp1, temp2 := ml[i+1:j], ml[j+1:]
						if i > 0 {
							result = append(result, ml[:i])
						}
						if temp1 != nil && len(temp1) > 0 {
							result = append(result, temp1)
						}
						if temp2 != nil && len(temp2) > 0 {
							result = append(result, temp2)
						}
						result = append(result, r1)
					}
					if i > j {
						temp1, temp2 := ml[j+1:i], ml[i+1:]
						if j > 0 {
							result = append(result, ml[:j])
						}
						if temp1 != nil && len(temp1) > 0 {
							result = append(result, temp1)
						}
						if temp2 != nil && len(temp2) > 0 {
							result = append(result, temp2)
						}
						result = append(result, r1)
					}
					return result, true
				}
			}
			if _, inVer := relate.InLineVertex(v.Matrix, ml[j].(matrix.LineMatrix)); inVer {
				r1 := mergeCheck(ml[j].(matrix.LineMatrix), ml[i].(matrix.LineMatrix))
				if r1 != nil {
					if i < j {
						temp1, temp2 := ml[i+1:j], ml[j+1:]
						if i > 0 {
							result = append(result, ml[:i])
						}
						if temp1 != nil && len(temp1) > 0 {
							result = append(result, temp1)
						}
						if temp2 != nil && len(temp2) > 0 {
							result = append(result, temp2)
						}
						result = append(result, r1)
					}
					if i > j {
						temp1, temp2 := ml[j+1:i], ml[i+1:]
						if j > 0 {
							result = append(result, ml[:j])
						}
						if temp1 != nil && len(temp1) > 0 {
							result = append(result, temp1)
						}
						if temp2 != nil && len(temp2) > 0 {
							result = append(result, temp2)
						}
						result = append(result, r1)
					}
					return result, true
				}
			}
		}
	}
	return ml, false
}

func mergeCheck(m0, m1 matrix.LineMatrix) matrix.LineMatrix {
	for i, mv := range m1 {
		if matrix.Matrix(mv).Equals(matrix.Matrix(m0[len(m0)-1])) {
			var result matrix.LineMatrix
			j := 1
			for ; len(m1)-1-j >= 0 && i-j >= 0; j++ {
				if matrix.Matrix(m1[i-j]).Equals(matrix.Matrix(m0[len(m0)-1-j])) {
					if i-j == 0 {
						result = m0
						result = append(result, m1[1:]...)
						return result
					}
					continue
				}
				return nil
			}
			result = m0
			result = append(result, m1[j:]...)
			return result
		}
	}
	return nil
}
