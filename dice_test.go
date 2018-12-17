package dice

import (
	"testing"
)

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
				t.Logf("expected Roll to be greater than 1 and less than %d, found %d", tc[1], roll)
				t.Fail()
			}
		}
	}
}