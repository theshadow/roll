package dice

import (
	"math/rand"

	"dice/formula"
)


// Roll accepts a roll formula, parses it, and then performs the roll
func Roll(f formula.Formula) (Results, error) {
	// parse the formula and perform the dice
	rolls, roll, err := rollF(f)
	if err != nil {
		return Results{}, err
	}

	// convert the extensions into their respective objects
	exts, err := buildExts(roll)

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

// RollR works like Roll except it takes a formula.Roll instead of a formula.Formula.
func RollR(r formula.Roll) (res Results, err error) {
	rolls := RollN(r.Count, r.Sides)

	// convert the extensions into their respective objects
	exts, err := buildExts(r)

	// Loop over each of the extensions and call the Exec() method.
	res = Results{
		Formula: formula.FromRoll(r),
		Rolls: rolls,
		Extensions: make(map[string]string),
	}
	for name, ext := range exts {
		s, err := ext.Exec(res, r)
		if err != nil {
			return Results{}, err
		}
		res.Extensions[name] = s
	}

	return Results{}, nil
}

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

// rollF takes a dice Formula and makes the Roll1 by parsing the formula and then rolling the dice.
func rollF(f formula.Formula) (rolls []int, roll formula.Roll, err error) {
	roll, err = parser.Parse(f)
	if err != nil {
		return []int{}, roll, err
	}

	rolls = RollN(roll.Count, roll.Sides)

	return rolls, roll, err
}

func buildExts(roll formula.Roll) (map[string]Extension, error) {
	// convert the extensions into their respective objects
	exts, err := New(roll.Extensions)
	if err != nil {
		return map[string]Extension{}, err
	}

	// if we have a modifier and the Sum extension isn't found, we implicitly add it.
	_, ok := exts[SumExtensionName]
	if roll.Modifier != 0 && !ok {
		exts[SumExtensionName], err = NewExtension(SumExtensionName, []string{})
		if err != nil {
			return map[string]Extension{}, err
		}
	}

	return exts, err
}

var parser = formula.New()