package formula

import "testing"

func TestParser_Parse(t *testing.T) {
	tests := map[string]Roll {
		"d2":     Roll{1, 2, 0, map[string][]string{}},
		"d10":    Roll{1, 10, 0, map[string][]string{}},
		"2d10":   Roll{2, 10, 0,map[string][]string{}},
		"d10+3":  Roll{1, 10, 3, map[string][]string{}},
		"d10-3":  Roll{1, 10, -3, map[string][]string{}},
		"2d10+3": Roll{2, 10, 3, map[string][]string{}},
		"2d10-3": Roll{2, 10, -3, map[string][]string{}},
		"2d10+3 critical(19)": Roll{2, 10, -3, map[string][]string{}},
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
	return a.Count == b.Count && a.Sides == b.Sides && a.Modifier == b.Modifier && compareParams(a.Extensions, b.Extensions)
}

func compareParams(a, b map[string][]string) bool {
	if len(a) != len(b) { return false }

	for k, av := range a {
		bv, ok := b[k];
		if !ok {
			return false
		}

		if testEq(av, bv) == false {
			return false
		}
	}

	return true
}

func testEq(a, b []string) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false;
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}