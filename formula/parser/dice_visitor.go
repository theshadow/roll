// Code generated from /Users/xguzman/Projects/dice/formula/Dice.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Dice
import "github.com/antlr/antlr4/runtime/Go/antlr"
// A complete Visitor for a parse tree produced by DiceParser.
type DiceVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DiceParser#formula.
	VisitFormula(ctx *FormulaContext) interface{}

	// Visit a parse tree produced by DiceParser#extensions.
	VisitExtensions(ctx *ExtensionsContext) interface{}

	// Visit a parse tree produced by DiceParser#count.
	VisitCount(ctx *CountContext) interface{}

	// Visit a parse tree produced by DiceParser#sides.
	VisitSides(ctx *SidesContext) interface{}

	// Visit a parse tree produced by DiceParser#modifier.
	VisitModifier(ctx *ModifierContext) interface{}

	// Visit a parse tree produced by DiceParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by DiceParser#parameters.
	VisitParameters(ctx *ParametersContext) interface{}

	// Visit a parse tree produced by DiceParser#funcname.
	VisitFuncname(ctx *FuncnameContext) interface{}

	// Visit a parse tree produced by DiceParser#funccall.
	VisitFunccall(ctx *FunccallContext) interface{}

}