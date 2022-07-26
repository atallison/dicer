package dice

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
)

type RollSegment interface {
	Value() int
	ToString() string
}

func Roll(diceString string) (*RollGroup, error) {
	group := &RollGroup{Equation: diceString}
	group, err := parse(group)
	if err != nil {
		return nil, err
	}
	return group, nil
}

var ErrInvalidDiceString = fmt.Errorf("invalid dice string")

func parse(r *RollGroup) (*RollGroup, error) {
	diceString := r.Equation

	group := regexp.MustCompile("(?P<count>\\d*)d(?P<value>\\d+)(?:(?P<highLow>k[hl])(?P<subCount>\\d*))?")
	countIndex := group.SubexpIndex("count")
	valueIndex := group.SubexpIndex("value")
	highLowIndex := group.SubexpIndex("highLow")
	subCountIndex := group.SubexpIndex("subCount")

	matches := group.FindStringSubmatch(diceString)
	if matches == nil {
		return r, ErrInvalidDiceString
	}

	count, err := strconv.Atoi(matches[countIndex])
	if err != nil {
		return r, ErrInvalidDiceString
	}

	value, err := strconv.Atoi(matches[valueIndex])
	if err != nil {
		return r, ErrInvalidDiceString
	}

	if len(matches) > highLowIndex+1 {
		highLow := matches[highLowIndex]
		if highLow != "" {
			subCountString := matches[subCountIndex]
			subCount := 1
			if subCountString != "" {
				subCount, err = strconv.Atoi(subCountString)
				if err != nil {
					return r, ErrInvalidDiceString
				}
			}
			if subCount >= count {
				return r, ErrInvalidDiceString
			}
			if highLow == "kh" {
				r.KeepHighest = subCount
			} else {
				r.KeepLowest = subCount
			}
		}
	}

	r.Rolls = make([]RollSegment, 0)
	for i := 0; i < count; i++ {
		r.Rolls = append(r.Rolls, RollDie{
			Equation: diceString,
			Sides:    value,
			Rolled:   random(value) + 1,
		})
	}

	return r, nil
}

func random(max int) int {
	return rand.Intn(max)
}
