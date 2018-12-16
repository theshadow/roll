package dice

import (
	"testing"

	"roll/formula"
)

func TestRoll(t *testing.T) {
	tests := [][]int{
		{1, 10},
		{2, 10},
		{10, 10},
	}

	for _, tc := range tests {
		act := Roll(tc[0], tc[1])

		if len(act) != tc[0] {
			t.Logf("expected %d rolls, instead found %d rolls", tc[0], len(act))
			t.Fail()
		}

		for _, roll := range act {
			if roll < 1 || roll > tc[1] {
				t.Logf("expected roll to be greater than 1 and less than %d, found %d", tc[1], roll)
			}
		}
	}
}

func TestRollWithFormula(t *testing.T) {
	tests := []string {
		"d10",
		"3d10",
	}

	for _, f := range tests {
		// RollWithFormula depends on Roll, we only need to be sure we can produce a valid roll.
		_, err := RollWithFormula(formula.Formula(f))
		if err != nil {
			t.Logf("unexpected error while rolling %s: %s", f, err)
		}
	}
}
