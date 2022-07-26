package dice

import (
	"fmt"
	"sort"
	"strings"
)

type RollGroup struct {
	Equation    string
	Rolls       []RollSegment
	KeepHighest int
	KeepLowest  int
}

func (r RollGroup) Value() int {
	sum := 0
	for _, roll := range r.getKeptRolls() {
		sum += roll.Value()
	}
	return sum
}

func (r RollGroup) getSortedRolls(desc bool) []RollSegment {
	rolls := make([]RollSegment, len(r.Rolls))
	copy(rolls, r.Rolls)
	sort.Slice(rolls, func(i, j int) bool {
		return rolls[i].Value() > rolls[j].Value() && desc || rolls[i].Value() < rolls[j].Value() && !desc
	})
	return rolls
}

func (r RollGroup) getKeptRolls() []RollSegment {
	if r.KeepHighest == 0 && r.KeepLowest == 0 {
		return r.Rolls
	}

	keptRolls := r.getSortedRolls(false)
	if r.KeepHighest > 0 {
		keptRolls = keptRolls[len(keptRolls)-r.KeepHighest:]
	} else {
		keptRolls = keptRolls[:r.KeepLowest]
	}
	return keptRolls
}

func (r RollGroup) ToString() string {
	if len(r.Rolls) < 2 {
		return fmt.Sprintf("%d", r.Value())
	}
	var rolls []RollSegment
	if r.KeepLowest != 0 {
		rolls = r.getSortedRolls(false)
	} else {
		rolls = r.getSortedRolls(true)
	}
	sb := strings.Builder{}
	sb.WriteString("(")
	for i, die := range rolls {
		if (r.KeepHighest != 0 || r.KeepLowest != 0) && i == 0 {
			sb.WriteString("[")
		}
		sb.WriteString(die.ToString())
		if i == r.KeepLowest-1 && r.KeepLowest != 0 || i == r.KeepHighest-1 && r.KeepHighest != 0 {
			sb.WriteString("]")
		}
		if i < len(rolls)-1 {
			sb.WriteString(" ")
		}
	}
	sb.WriteString(")")
	return sb.String()
}
