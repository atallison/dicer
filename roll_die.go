package dice

import "fmt"

type RollDie struct {
	Equation string
	Sides    int
	Rolled   int
}

func (r RollDie) Value() int {
	return r.Rolled
}

func (r RollDie) ToString() string {
	return fmt.Sprintf("%d", r.Value())
}
