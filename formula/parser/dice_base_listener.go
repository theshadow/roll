// Code generated from /Users/xguzman/Projects/dice/formula/Dice.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Dice
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDiceListener is a complete listener for a parse tree produced by DiceParser.
type BaseDiceListener struct{}

var _ DiceListener = &BaseDiceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDiceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDiceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDiceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDiceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFormula is called when production formula is entered.
func (s *BaseDiceListener) EnterFormula(ctx *FormulaContext) {}

// ExitFormula is called when production formula is exited.
func (s *BaseDiceListener) ExitFormula(ctx *FormulaContext) {}

// EnterExtensions is called when production extensions is entered.
func (s *BaseDiceListener) EnterExtensions(ctx *ExtensionsContext) {}

// ExitExtensions is called when production extensions is exited.
func (s *BaseDiceListener) ExitExtensions(ctx *ExtensionsContext) {}

// EnterCount is called when production count is entered.
func (s *BaseDiceListener) EnterCount(ctx *CountContext) {}

// ExitCount is called when production count is exited.
func (s *BaseDiceListener) ExitCount(ctx *CountContext) {}

// EnterSides is called when production sides is entered.
func (s *BaseDiceListener) EnterSides(ctx *SidesContext) {}

// ExitSides is called when production sides is exited.
func (s *BaseDiceListener) ExitSides(ctx *SidesContext) {}

// EnterModifier is called when production modifier is entered.
func (s *BaseDiceListener) EnterModifier(ctx *ModifierContext) {}

// ExitModifier is called when production modifier is exited.
func (s *BaseDiceListener) ExitModifier(ctx *ModifierContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseDiceListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseDiceListener) ExitParameter(ctx *ParameterContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BaseDiceListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BaseDiceListener) ExitParameters(ctx *ParametersContext) {}

// EnterFuncname is called when production funcname is entered.
func (s *BaseDiceListener) EnterFuncname(ctx *FuncnameContext) {}

// ExitFuncname is called when production funcname is exited.
func (s *BaseDiceListener) ExitFuncname(ctx *FuncnameContext) {}

// EnterFunccall is called when production funccall is entered.
func (s *BaseDiceListener) EnterFunccall(ctx *FunccallContext) {}

// ExitFunccall is called when production funccall is exited.
func (s *BaseDiceListener) ExitFunccall(ctx *FunccallContext) {}
