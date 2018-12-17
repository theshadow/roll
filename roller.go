package dice

import "roll/formula"

// FormulaRoll is a Roll that results from rolling with a Formula
type FormulaRoll struct {
	Formula formula.Formula
	Rolls []int
	Extensions map[string]string
}

// RollWithFormula accepts a Roll formula, parses it, and then performs the Roll
func RollWithFormula(f formula.Formula) (FormulaRoll, error) {
	// parse the formula and perform the roll
	rolls, roll, err := rollWithFormula(f)
	if err != nil {
		return FormulaRoll{}, err
	}

	// convert the extensions into their respective objects
	exts := make(map[string]Extension)
	for k, ps := range roll.Extensions {
		exts[k], err = New(k, ps)
		if err != nil {
			return FormulaRoll{}, err
		}
	}

	// if we have a modifier and the Sum extension isn't found, we implicitly add it.
	_, ok := exts[SumExtensionName]
	if roll.Modifier != 0 && !ok {
		exts[SumExtensionName], err = New(SumExtensionName, []string{})
		if err != nil {
			return FormulaRoll{}, err
		}
	}

	// Loop over each of the extensions and call the Exec() method.
	fr := FormulaRoll{
		Formula: f,
		Rolls: rolls,
		Extensions: make(map[string]string),
	}
	for name, ext := range exts {
		s, err := ext.Exec(fr, roll)
		if err != nil {
			return FormulaRoll{}, err
		}
		fr.Extensions[name] = s
	}

	return fr, nil
}

// rollWithFormula takes a Roll Formula and makes the Roll by parsing the formula and then rolling the dice.
func rollWithFormula(f formula.Formula) (rolls []int, roll formula.Roll, err error) {
	roll, err = parser.Parse(f)
	if err != nil {
		return []int{}, roll, err
	}

	rolls = RollN(roll.Count, roll.Sides)

	return rolls, roll, err
}

var parser = formula.New()



