package dice

import (
	"math/rand"

	"roll/formula"
)

// RollWithFormula takes a roll Formula and makes the roll by parsing the formula and then rolling the dice.
func RollWithFormula(f formula.Formula) (rolls []int, err error) {
	r, err := parser.Parse(f)
	if err != nil {
		return []int{}, err
	}
	return Roll(r.Count, r.Sides), nil
}

func Roll(count, sides int) (rolls []int) {
	for i := 0; i < count; i++ {
		rolls = append(rolls, int(rand.Int31n(int32(sides))+1))
	}
	return rolls
}

var parser = formula.New()