package formula

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	fp "dice/formula/parser"
)

func parse(f Formula) (_ Roll, err error) {
	defer func () {
		if r := recover(); r != nil {
			err = errors.Wrap(r.(error), fmt.Sprintf("failed to parse formula \"%s\"", f))
		}
	}()

	input := antlr.NewInputStream(string(f))
	lexer := fp.NewDiceLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	listener := new(DiceParserListener)

	dp := fp.NewDiceParser(stream)
	dp.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	dp.BuildParseTrees = true

	antlr.ParseTreeWalkerDefault.Walk(listener, dp.Formula())

	return listener.Roll(), err
}

// ListenerError Captures errors
type ListenerError struct {
	antlr.ErrorListener
}

// DiceParserListener implements the BaseDiceListener
type DiceParserListener struct {
	roll Roll

	extName string

	fp.BaseDiceListener
}

// Roll returns the parsed dice
func (s *DiceParserListener) Roll() Roll { return s.roll }

// EnterFormula is called when production formula is entered.
func (s *DiceParserListener) EnterFormula(ctx *fp.FormulaContext) {
	// default the Count to one, you can't have less than that.
	s.roll = Roll{
		Count: 1,
	}
}

// EnterCount is called when production count is entered.
func (s *DiceParserListener) EnterCount(ctx *fp.CountContext) {
	i, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		panic(fmt.Sprintf("unable to cast Count to integer: %s", err))
	}
	s.roll.Count = i
}

// EnterSides is called when production sides is entered.
func (s *DiceParserListener) EnterSides(ctx *fp.SidesContext) {
	i, err := strconv.Atoi(ctx.GetText()[1:])
	if err != nil {
		panic(fmt.Sprintf("unable to cast Sides to integer: %s", err))
	}
	s.roll.Sides = i
}

// EnterModifier is called when production modifier is entered.
func (s *DiceParserListener) EnterModifier(ctx *fp.ModifierContext) {
	n := 1
	if ctx.SIGN().GetText() == "-" {
		n = -1
	}

	str := ctx.Integer().GetText()
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("unable to cast Modifier to integer: %s", err))
	}
	s.roll.Modifier = i * n
}

// EnterParameter is called when production parameter is entered.
func (s *DiceParserListener) EnterParameter(ctx *fp.ParameterContext) {
	var str string
	t := ctx.GetToken(fp.DiceParserStringLiteral, 0)
	if t == nil {
		str = ctx.Integer().GetText()
	} else {
		v := ctx.StringLiteral().GetText()
		str = v[1:len(v)-1]
	}
	s.roll.Extensions[s.extName] = append(s.roll.Extensions[s.extName], str)
}

// EnterFunc_name is called when production func_name is entered.
func (s *DiceParserListener) EnterFuncname(ctx *fp.FuncnameContext) {
	s.extName = ctx.GetText()
	s.roll.Extensions[s.extName] = []string{}
}

// EnterFunc_call is called when production func_call is entered.
func (s *DiceParserListener) EnterFunccall(ctx *fp.FunccallContext) {
	if len(s.roll.Extensions) == 0 {
		s.roll.Extensions = make(map[string][]string)
	}
}

// ExitFunc_call is called when production func_call is exited.
func (s *DiceParserListener) ExitFunccall(ctx *fp.FunccallContext) {
	s.extName = ""
}
