package formula

import (
	"errors"
	"regexp"
	"strconv"
)

// Roll represents the component parts of a roll formula that can be used to actually perform a rolling of dice.
type Roll struct {
	Count int
	Sides int
	Modifier int
}

/**
 * Formula
 * A roll formula is a string that breaks down the concept of rolling a number of dice
 * with a number of sides into an algebraic like form. The idea is to quickly be able
 * to describe a role in a more human friendly way.
 *
 * count := uint
 * sides := uint
 * modifer := integer
 * expr  := count? 'd' sides modifier?
 *
 * EXAMPLES:
 * - d10     - Basic
 * - 3d10    - With count
 * - d10+10  - With modifier
 * - 3d10+10 - With count and modifier
 * - 3d10-2  - With negative modifier
 */
type Formula string

// Parser provides a function for consuming a Formula and transforming it into a Roll
type Parser interface {
	// Parse accepts a string that represents a dice rolling formula
	Parse(f Formula) (Roll, error)
}

func New() Parser {
	return parser{}
}

// re defines a regular expression that can match and extract the parts of a roll formula
var re = regexp.MustCompile(`(?P<count>\d+)?d(?P<sides>\d+)((?P<sign>[+|\-])(?P<modifier>\d+))?`)

type parser struct {}

func (p parser) Parse(f Formula) (Roll, error) {
	fs := string(f)
	if len(fs) < 2 {
		return Roll{}, errors.New("a valid formula must be at least two characters long, 'd' and a number of sides")
	}
	if !re.MatchString(fs) {
		return Roll{}, errors.New("invalid roll formula")
	}

	sm := re.FindStringSubmatch(fs)
	// "formula", "count", "sides", "", "sign", "modifier"
	count, sides, modifier := p.atoi(sm[1], 1), p.atoi(sm[2], 2), p.atoi(sm[5], 0)

	if sm[4] == "-" {
			modifier *= -1
	}

	return Roll{
		Count: count,
		Sides: sides,
		Modifier: modifier,
	}, nil
}

func (p parser) atoi(s string, def int) int {
	if len(s) == 0 {
		return def
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return def
	}

	return i
}