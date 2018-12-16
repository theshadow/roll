package formula

import "testing"

func TestParser_Parse(t *testing.T) {
	tests := map[string]Roll {
		"d2":     Roll{1, 2, 0},
		"d10":    Roll{1, 10, 0},
		"2d10":   Roll{2, 10, 0},
		"d10+3":  Roll{1, 10, 3},
		"d10-3":  Roll{1, 10, -3},
		"2d10+3": Roll{2, 10, 3},
		"2d10-3": Roll{2, 10, -3},
	}

	p := parser{}

	for f, exp := range tests {
		act, err := p.Parse(Formula(f))

		if err != nil {
			t.Logf("unexpected error while parsing formula %s: %s", f, err)
			t.Fail()
		}

		if compareRolls(exp, act) != true {
			t.Logf("formula %s expected %#v and instead received %#v", f, exp, act)
		}
	}
}

// compareRolls returns true when all properties of a Roll are equal
func compareRolls(a, b Roll) bool {
	return a.Count == b.Count && a.Sides == b.Sides && a.Modifier == b.Modifier
}