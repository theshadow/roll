package dice

import (
	"fmt"
	"testing"

	"dice/formula"
)


func TestRollWithFormulaPrivate(t *testing.T) {
	// https://play.golang.org/p/kjHvdPfhHp-
	tests := []string {
		"d10",
		"2d10",
		"2d10+5",
		"2d10-5",
		"2d10+10",
		"2d10+10 advantage()",
		"2d10+10 critical(20)",
		"2d10+10 multiparam(20,\"bunny\",1)",
		"2d10+10 critical(20) advantage()",
	}

	for _, f := range tests {
		// rollF depends on RollN, we only need to be sure we can produce a valid Roll1.
		_, _, err := rollF(formula.Formula(f))
		if err != nil {
			t.Logf("unexpected error while rolling %s: %s", f, err)
			t.Fail()
		}
	}
}

func TestRollWithFormula(t *testing.T) {
	f := formula.Formula("d20+4")
	roll, err := Roll(f)
	if err != nil {
		t.Logf("unexpected error while rolling %s: %s", f, err)
	}

	var mod, total int

	_, err = fmt.Sscanf(roll.Extensions["Sum"], "With the modifier %d, the total is %d", &mod, &total)
	if err != nil {
		t.Logf("unexpected error while parsing Sum extension: %s", err)
		t.Fail()
	}

	// verify that the total was modified and within our expected range.
	if total < 5 || total > 24 {
		t.Logf("expected the total to be greater than 5 and less than 24 with a Roll1 of %s, instead received %d",
			string(f), total)
		t.Fail()
	}

	if t.Failed() {
		t.Logf("Roll1: %#v", roll)
	}
}

func TestRollN(t *testing.T) {
	tests := [][]int{
		{1, 10},
		{2, 10},
		{10, 10},
	}

	for _, tc := range tests {
		act := RollN(tc[0], tc[1])

		// validate that we rolled the correct number of dice.
		if len(act) != tc[0] {
			t.Logf("expected %d rolls, instead found %d rolls", tc[0], len(act))
			t.Fail()
		}

		// validate that each die was rolled within the correct range.
		for _, roll := range act {
			if roll < 1 || roll > tc[1] {
				t.Logf("expected Roll1 to be greater than 1 and less than %d, found %d", tc[1], roll)
				t.Fail()
			}
		}
	}
}