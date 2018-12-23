// Code generated from /Users/xguzman/Projects/dice/formula/Dice.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Dice
import "github.com/antlr/antlr4/runtime/Go/antlr"

// DiceListener is a complete listener for a parse tree produced by DiceParser.
type DiceListener interface {
	antlr.ParseTreeListener

	// EnterFormula is called when entering the formula production.
	EnterFormula(c *FormulaContext)

	// EnterExtensions is called when entering the extensions production.
	EnterExtensions(c *ExtensionsContext)

	// EnterCount is called when entering the count production.
	EnterCount(c *CountContext)

	// EnterSides is called when entering the sides production.
	EnterSides(c *SidesContext)

	// EnterModifier is called when entering the modifier production.
	EnterModifier(c *ModifierContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterParameters is called when entering the parameters production.
	EnterParameters(c *ParametersContext)

	// EnterFuncname is called when entering the funcname production.
	EnterFuncname(c *FuncnameContext)

	// EnterFunccall is called when entering the funccall production.
	EnterFunccall(c *FunccallContext)

	// ExitFormula is called when exiting the formula production.
	ExitFormula(c *FormulaContext)

	// ExitExtensions is called when exiting the extensions production.
	ExitExtensions(c *ExtensionsContext)

	// ExitCount is called when exiting the count production.
	ExitCount(c *CountContext)

	// ExitSides is called when exiting the sides production.
	ExitSides(c *SidesContext)

	// ExitModifier is called when exiting the modifier production.
	ExitModifier(c *ModifierContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitParameters is called when exiting the parameters production.
	ExitParameters(c *ParametersContext)

	// ExitFuncname is called when exiting the funcname production.
	ExitFuncname(c *FuncnameContext)

	// ExitFunccall is called when exiting the funccall production.
	ExitFunccall(c *FunccallContext)
}
