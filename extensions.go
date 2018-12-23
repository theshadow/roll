package dice

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/theshadow/dice/formula"
)

const (
	SumExtensionName          = "sum"
	LogRollExtensionName      = "log"
	CriticalRollExtensionName = "critical"
	ExplodingDiceName         = "exploding"
	AdvantageExtensionName    = "advantage"
	DisadvantageExtensionName = "disadvantage"
)

// Extension is a Roll1 add-on. They work by providing post-Roll1 information and meta-information such as the sum of all
// the rolls, noting if any of the dice rolls are critical, or if any of the dice exploded and what the result was.
// They DO NOT and SHOULD NOT modify the original dice.
type Extension interface {
	Exec(fr Results, r formula.Roll) (string, error)
	Name() string
}

func New(extensions map[string][]string) (map[string]Extension, error) {
	exts := make(map[string]Extension)
	for ext, params := range extensions {
		e, err := NewExtension(ext, params)
		if err != nil {
			return exts, err
		}
		exts[ext] = e
	}
	return exts, nil
}

func NewExtension(name string, params []string) (Extension, error) {
	switch name {
	case SumExtensionName:
		return newSumExtension(params), nil
	case LogRollExtensionName:
		return newLogRollExtension(params), nil
	case CriticalRollExtensionName:
		return newCriticalRollExtension(params)
	case ExplodingDiceName:
		return newExplodingDiceExtension(params)
	case AdvantageExtensionName:
		return newAdvantageExtension(params), nil
	case DisadvantageExtensionName:
		return newDisadvantageExtension(params), nil
	}
	return nil, fmt.Errorf("unknown extension name \"%s\"", name)
}

func newLogRollExtension(params []string) LogRollExtension {
	_ = params
	return LogRollExtension{}
}

// LogRollExtension will record the specified Roll1 formula and the results of each Roll1.
type LogRollExtension struct{
	extension
}
func (ext LogRollExtension) Name() string { return LogRollExtensionName }
func (ext LogRollExtension) Exec(fr Results, r formula.Roll) (string, error) {
	var msg = fmt.Sprintf("Rolled: \"%s\" Rolls: %v", string(fr.Formula), fr.Rolls)
	ext.log.Printf(msg)
	return msg, nil
}

func newCriticalRollExtension(params []string) (CriticalRollExtension, error) {
	var ext CriticalRollExtension
	if len(params) != 1 {
		return ext, fmt.Errorf("invalid input, expected %d received %d", 1, len(params))
	}

	i, err := strconv.Atoi(params[0])
	if err != nil {
		return ext,
			fmt.Errorf("invalid input, LowerLimit is expected to be a signed integer received \"%s\"", params[0])
	}

	ext.LowerLimit = i

	return ext, nil
}

// CriticalRollExtension when any of the rolls are equal to or greater than the LowerLimit the entire Roll1 is considered
// a critical Roll1
type CriticalRollExtension struct{
	// LowerLimit any Roll1 equal to or greater than this value is considered a critical.
	LowerLimit int
}
func (ext CriticalRollExtension) Name() string { return CriticalRollExtensionName }
func (ext CriticalRollExtension) Exec(fr Results, r formula.Roll) (string, error) {
	if r.Count > 1 {
		return "", errors.New("only single die rolls may use the critical dice extension")
	}

	for _, r := range fr.Rolls {
		if r >= ext.LowerLimit {
			return "Yes", nil
		}
	}

	return "No", nil
}

func newExplodingDiceExtension(params []string) (ExplodingDiceExtension, error) {
	var ext ExplodingDiceExtension
	if len(params) != 1 {
		return ext, fmt.Errorf("invalid input, expected %d received %d", 1, len(params))
	}

	i, err := strconv.Atoi(params[0])
	if err != nil {
		return ext, fmt.Errorf("invalid input, LowerLimit is expected to be a signed integer received \"%s\"",
			params[0])
	}

	ext.LowerLimit = i

	return ext, nil
}

// ExplodingDiceExtension any die that rolls a value equal to or greater than the LowerLimit will be considered
// an exploding die and an additional die will be rolls. Only those dice in the original set may explode, and each die
// may only explode once.
type ExplodingDiceExtension struct{
	// LowerLimit any Roll1 equal to or greater than this value is considered a critical.
	LowerLimit int
}
func (ext ExplodingDiceExtension) Name() string { return ExplodingDiceName }
func (ext ExplodingDiceExtension) Exec(fr Results, r formula.Roll) (string, error) {
	var rolls []int
	for _, roll := range fr.Rolls {
		if roll >= ext.LowerLimit {
			rolls = append(rolls, Roll1(r.Sides))
		}
	}

	return fmt.Sprintf("%d dice exploded adding the rolls %v, their sum is %d", len(rolls), rolls,
		sum(rolls)), nil
}

func newAdvantageExtension(param []string) AdvantageExtension {
	_ = param
	return AdvantageExtension{}
}

// AdvantageExtension any single die dice will have an additional die rolled and the higher of the two reported.
type AdvantageExtension struct{}
func (ext AdvantageExtension) Name() string { return AdvantageExtensionName }
func (ext AdvantageExtension) Exec(fr Results, r formula.Roll) (string, error) {
	if r.Sides > 1 {
		return "", errors.New("only single die rolls may use the advantage extension")
	}

	sec := Roll1(r.Sides) + r.Modifier

	if sec > fr.Rolls[0] {
		return fmt.Sprintf("An advantage was had, you rolled a higher dice with %d", sec), nil
	}

	return "", nil
}

func newDisadvantageExtension(param []string) DisadvantageExtension {
	_ = param
	return DisadvantageExtension{}
}

// DisadvantageExtension any single die dice will have an additional die rolled and the lower of the two reported.
type DisadvantageExtension struct{}
func (ext DisadvantageExtension) Name() string { return "Disadvantage" }
func (ext DisadvantageExtension) Exec(fr Results, r formula.Roll) (string, error) {
	if r.Sides > 1 {
		return "", errors.New("only single die rolls may use the disadvantage extension")
	}

	sec := Roll1(r.Sides) + r.Modifier

	if sec < fr.Rolls[0] {
		return fmt.Sprintf("an disadvantage was had, you rolled a lower dice with %d", sec), nil
	}

	return "", nil
}

func newSumExtension(params []string) SumExtension {
	_ = params
	return SumExtension{}
}

// SumExtension calculates the sum of all the rolls and adds in the modifier if one is specified
type SumExtension struct {}
func (ext SumExtension) Name() string { return SumExtensionName }
func (ext SumExtension) Exec(fr Results, r formula.Roll) (string, error) {
	if r.Modifier != 0 {
		return fmt.Sprintf("With the modifier %d, the total is %d", r.Modifier,
			sumAddModifier(fr.Rolls, r.Modifier)), nil
	}

	return fmt.Sprintf("The total is %d", sum(fr.Rolls)), nil
}

func sumAddModifier(rolls []int, modifier int) int {
	return sum(rolls) + modifier
}

func sum(rolls []int) (n int) {
	for _, i := range rolls { n += i }
	return n
}

type extension struct {
	log log.Logger
}

