package dice

import (
	"fmt"
	"roll/formula"
	"testing"
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
		"2d10+10 multiparam(20,bunny,1)",
		"2d10+10 critical(20) advantage()",
	}

	for _, f := range tests {
		// rollWithFormula depends on RollN, we only need to be sure we can produce a valid Roll.
		_, _, err := rollWithFormula(formula.Formula(f))
		if err != nil {
			t.Logf("unexpected error while rolling %s: %s", f, err)
			t.Fail()
		}
	}
}

func TestRollWithFormula(t *testing.T) {
	f := formula.Formula("d20+4")
	roll, err := RollWithFormula(f)
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
		t.Logf("expected the total to be greater than 5 and less than 24 with a Roll of %s, instead received %d",
			string(f), total)
		t.Fail()
	}

	if t.Failed() {
		t.Logf("Roll: %#v", roll)
	}
}
