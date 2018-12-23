package dice

import (
	"math/rand"
	"roll/formula"
)

// RollN takes a count of dice and the number of sides per dice and rolls them.
func RollN(n, sides int) (rolls []int) {
	for i := 0; i < n; i++ {
		rolls = append(rolls, Roll1(sides))
	}
	return rolls
}

// Roll1 takes the number of sides for a die and returns the result of that dice.
func Roll1(sides int) int {
	return int(rand.Int31n(int32(sides))+1)
}

// Results is a Roll1 that results from rolling with a Formula
type Results struct {
	Formula formula.Formula
	Rolls []int
	Extensions map[string]string
}

// Roll accepts a roll formula, parses it, and then performs the roll
func Roll(f formula.Formula) (Results, error) {
	// parse the formula and perform the dice
	rolls, roll, err := rollF(f)
	if err != nil {
		return Results{}, err
	}

	// convert the extensions into their respective objects
	exts := make(map[string]Extension)
	for k, ps := range roll.Extensions {
		exts[k], err = NewExtension(k, ps)
		if err != nil {
			return Results{}, err
		}
	}

	// if we have a modifier and the Sum extension isn't found, we implicitly add it.
	_, ok := exts[SumExtensionName]
	if roll.Modifier != 0 && !ok {
		exts[SumExtensionName], err = NewExtension(SumExtensionName, []string{})
		if err != nil {
			return Results{}, err
		}
	}

	// Loop over each of the extensions and call the Exec() method.
	res := Results{
		Formula: f,
		Rolls: rolls,
		Extensions: make(map[string]string),
	}
	for name, ext := range exts {
		s, err := ext.Exec(res, roll)
		if err != nil {
			return Results{}, err
		}
		res.Extensions[name] = s
	}

	return res, nil
}

// rollF takes a dice Formula and makes the Roll1 by parsing the formula and then rolling the dice.
func rollF(f formula.Formula) (rolls []int, roll formula.Roll, err error) {
	roll, err = parser.Parse(f)
	if err != nil {
		return []int{}, roll, err
	}

	rolls = RollN(roll.Count, roll.Sides)

	return rolls, roll, err
}

var parser = formula.New()