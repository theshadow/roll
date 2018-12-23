// Code generated from /Users/xguzman/Projects/dice/formula/Dice.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Dice
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseDiceVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDiceVisitor) VisitFormula(ctx *FormulaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDiceVisitor) VisitExtensions(ctx *ExtensionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDiceVisitor) VisitCount(ctx *CountContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDiceVisitor) VisitSides(ctx *SidesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDiceVisitor) VisitModifier(ctx *ModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDiceVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDiceVisitor) VisitParameters(ctx *ParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDiceVisitor) VisitFuncname(ctx *FuncnameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDiceVisitor) VisitFunccall(ctx *FunccallContext) interface{} {
	return v.VisitChildren(ctx)
}
