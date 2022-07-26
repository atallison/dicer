package dice

import "strconv"

type RollConstant struct {
	ConstantValue int
}

func (r RollConstant) Value() int {
	return r.ConstantValue
}

func (r RollConstant) ToString() string {
	return strconv.Itoa(r.ConstantValue)
}
