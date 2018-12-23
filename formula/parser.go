package formula

import (
	"fmt"
	"regexp"
	"strconv"
)

// Roll represents the component parts of a dice formula that can be used to actually perform a rolling of dice.
type Roll struct {
	Count int
	Sides int
	Modifier int
	Extensions map[string][]string
}

/**
 * Formula
 * A dice formula is a string that breaks down the concept of rolling a number of dice
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

// FromRoll accepts a Roll and attempts to transform it into the formula version.
func FromRoll(r Roll) (f Formula){
	count := ""
	if r.Count > 1 {
		count = strconv.Itoa(r.Count)
	}

	modifier := ""
	if r.Modifier != 0 {
		if r.Modifier > 0 {
			modifier = fmt.Sprintf("+%d", r.Modifier)
		} else {
			modifier = strconv.Itoa(r.Modifier)
		}
	}

	// TODO Build something better populate this. The major issue is detecting integers vs strings.
	extensions := ""

	return Formula(fmt.Sprintf("%sd%d%s%s", count, r.Sides, modifier, extensions))
}

// Parser provides a function for consuming a Formula and transforming it into a Roll
type Parser interface {
	// Parse accepts a string that represents a dice rolling formula
	Parse(f Formula) (Roll, error)
}

func New() Parser {
	return parser{}
}

// https://regexr.com/451jd
const rexpr = `(?P<formula>(?P<count>\d+)?d(?P<sides>\d+)((?P<sign>[+|\-])(?P<modifier>\d+))?(\ ((?P<extname>[[:alpha:]][[:word:]]+)\((?P<extparams>([[:word:]]+)(,\ ?([[:word:]]+))*)*\)))*)`
// re defines a regular expression that can match and extract the parts of a dice formula
var re = regexp.MustCompile(rexpr)

type parser struct {}

func (p parser) Parse(f Formula) (_ Roll, err error) {
	return parse(f)
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