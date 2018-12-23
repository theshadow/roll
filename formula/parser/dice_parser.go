// Code generated from /Users/xguzman/Projects/dice/formula/Dice.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Dice
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa


var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 12, 65, 4, 
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4, 
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 5, 2, 22, 10, 2, 3, 2, 3, 2, 5, 
	2, 26, 10, 2, 3, 2, 3, 2, 3, 3, 3, 3, 7, 3, 32, 10, 3, 12, 3, 14, 3, 35, 
	11, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 
	3, 8, 7, 8, 49, 10, 8, 12, 8, 14, 8, 52, 11, 8, 3, 8, 3, 8, 3, 9, 3, 9, 
	3, 10, 3, 10, 3, 10, 5, 10, 61, 10, 10, 3, 10, 3, 10, 3, 10, 2, 2, 11, 
	2, 4, 6, 8, 10, 12, 14, 16, 18, 2, 3, 4, 2, 10, 10, 12, 12, 2, 60, 2, 21, 
	3, 2, 2, 2, 4, 33, 3, 2, 2, 2, 6, 36, 3, 2, 2, 2, 8, 38, 3, 2, 2, 2, 10, 
	40, 3, 2, 2, 2, 12, 43, 3, 2, 2, 2, 14, 50, 3, 2, 2, 2, 16, 55, 3, 2, 2, 
	2, 18, 57, 3, 2, 2, 2, 20, 22, 5, 6, 4, 2, 21, 20, 3, 2, 2, 2, 21, 22, 
	3, 2, 2, 2, 22, 23, 3, 2, 2, 2, 23, 25, 5, 8, 5, 2, 24, 26, 5, 10, 6, 2, 
	25, 24, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 27, 3, 2, 2, 2, 27, 28, 5, 
	4, 3, 2, 28, 3, 3, 2, 2, 2, 29, 30, 7, 8, 2, 2, 30, 32, 5, 18, 10, 2, 31, 
	29, 3, 2, 2, 2, 32, 35, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2, 33, 34, 3, 2, 2, 
	2, 34, 5, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 36, 37, 7, 10, 2, 2, 37, 7, 3, 
	2, 2, 2, 38, 39, 7, 11, 2, 2, 39, 9, 3, 2, 2, 2, 40, 41, 7, 4, 2, 2, 41, 
	42, 7, 10, 2, 2, 42, 11, 3, 2, 2, 2, 43, 44, 9, 2, 2, 2, 44, 13, 3, 2, 
	2, 2, 45, 46, 5, 12, 7, 2, 46, 47, 7, 7, 2, 2, 47, 49, 3, 2, 2, 2, 48, 
	45, 3, 2, 2, 2, 49, 52, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 50, 51, 3, 2, 2, 
	2, 51, 53, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 53, 54, 5, 12, 7, 2, 54, 15, 
	3, 2, 2, 2, 55, 56, 7, 11, 2, 2, 56, 17, 3, 2, 2, 2, 57, 58, 5, 16, 9, 
	2, 58, 60, 7, 5, 2, 2, 59, 61, 5, 14, 8, 2, 60, 59, 3, 2, 2, 2, 60, 61, 
	3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 63, 7, 6, 2, 2, 63, 19, 3, 2, 2, 2, 
	7, 21, 25, 33, 50, 60,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "'('", "')'", "','", "' '",
}
var symbolicNames = []string{
	"", "D", "SIGN", "LPAREN", "RPAREN", "COMMA", "SPACE", "WS", "Integer", 
	"Id", "StringLiteral",
}

var ruleNames = []string{
	"formula", "extensions", "count", "sides", "modifier", "parameter", "parameters", 
	"funcname", "funccall",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type DiceParser struct {
	*antlr.BaseParser
}

func NewDiceParser(input antlr.TokenStream) *DiceParser {
	this := new(DiceParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Dice.g4"

	return this
}

// DiceParser tokens.
const (
	DiceParserEOF = antlr.TokenEOF
	DiceParserD = 1
	DiceParserSIGN = 2
	DiceParserLPAREN = 3
	DiceParserRPAREN = 4
	DiceParserCOMMA = 5
	DiceParserSPACE = 6
	DiceParserWS = 7
	DiceParserInteger = 8
	DiceParserId = 9
	DiceParserStringLiteral = 10
)

// DiceParser rules.
const (
	DiceParserRULE_formula = 0
	DiceParserRULE_extensions = 1
	DiceParserRULE_count = 2
	DiceParserRULE_sides = 3
	DiceParserRULE_modifier = 4
	DiceParserRULE_parameter = 5
	DiceParserRULE_parameters = 6
	DiceParserRULE_funcname = 7
	DiceParserRULE_funccall = 8
)

// IFormulaContext is an interface to support dynamic dispatch.
type IFormulaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormulaContext differentiates from other interfaces.
	IsFormulaContext()
}

type FormulaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormulaContext() *FormulaContext {
	var p = new(FormulaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DiceParserRULE_formula
	return p
}

func (*FormulaContext) IsFormulaContext() {}

func NewFormulaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormulaContext {
	var p = new(FormulaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceParserRULE_formula

	return p
}

func (s *FormulaContext) GetParser() antlr.Parser { return s.parser }

func (s *FormulaContext) Sides() ISidesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISidesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISidesContext)
}

func (s *FormulaContext) Extensions() IExtensionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtensionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExtensionsContext)
}

func (s *FormulaContext) Count() ICountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICountContext)
}

func (s *FormulaContext) Modifier() IModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModifierContext)
}

func (s *FormulaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormulaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FormulaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.EnterFormula(s)
	}
}

func (s *FormulaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.ExitFormula(s)
	}
}

func (s *FormulaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DiceVisitor:
		return t.VisitFormula(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *DiceParser) Formula() (localctx IFormulaContext) {
	localctx = NewFormulaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DiceParserRULE_formula)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == DiceParserInteger {
		{
			p.SetState(18)
			p.Count()
		}

	}
	{
		p.SetState(21)
		p.Sides()
	}
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == DiceParserSIGN {
		{
			p.SetState(22)
			p.Modifier()
		}

	}
	{
		p.SetState(25)
		p.Extensions()
	}



	return localctx
}


// IExtensionsContext is an interface to support dynamic dispatch.
type IExtensionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtensionsContext differentiates from other interfaces.
	IsExtensionsContext()
}

type ExtensionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtensionsContext() *ExtensionsContext {
	var p = new(ExtensionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DiceParserRULE_extensions
	return p
}

func (*ExtensionsContext) IsExtensionsContext() {}

func NewExtensionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtensionsContext {
	var p = new(ExtensionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceParserRULE_extensions

	return p
}

func (s *ExtensionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtensionsContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(DiceParserSPACE)
}

func (s *ExtensionsContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(DiceParserSPACE, i)
}

func (s *ExtensionsContext) AllFunccall() []IFunccallContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunccallContext)(nil)).Elem())
	var tst = make([]IFunccallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunccallContext)
		}
	}

	return tst
}

func (s *ExtensionsContext) Funccall(i int) IFunccallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunccallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunccallContext)
}

func (s *ExtensionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtensionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExtensionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.EnterExtensions(s)
	}
}

func (s *ExtensionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.ExitExtensions(s)
	}
}

func (s *ExtensionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DiceVisitor:
		return t.VisitExtensions(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *DiceParser) Extensions() (localctx IExtensionsContext) {
	localctx = NewExtensionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DiceParserRULE_extensions)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == DiceParserSPACE {
		{
			p.SetState(27)
			p.Match(DiceParserSPACE)
		}
		{
			p.SetState(28)
			p.Funccall()
		}


		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// ICountContext is an interface to support dynamic dispatch.
type ICountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCountContext differentiates from other interfaces.
	IsCountContext()
}

type CountContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCountContext() *CountContext {
	var p = new(CountContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DiceParserRULE_count
	return p
}

func (*CountContext) IsCountContext() {}

func NewCountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CountContext {
	var p = new(CountContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceParserRULE_count

	return p
}

func (s *CountContext) GetParser() antlr.Parser { return s.parser }

func (s *CountContext) Integer() antlr.TerminalNode {
	return s.GetToken(DiceParserInteger, 0)
}

func (s *CountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *CountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.EnterCount(s)
	}
}

func (s *CountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.ExitCount(s)
	}
}

func (s *CountContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DiceVisitor:
		return t.VisitCount(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *DiceParser) Count() (localctx ICountContext) {
	localctx = NewCountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DiceParserRULE_count)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		p.Match(DiceParserInteger)
	}



	return localctx
}


// ISidesContext is an interface to support dynamic dispatch.
type ISidesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSidesContext differentiates from other interfaces.
	IsSidesContext()
}

type SidesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySidesContext() *SidesContext {
	var p = new(SidesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DiceParserRULE_sides
	return p
}

func (*SidesContext) IsSidesContext() {}

func NewSidesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SidesContext {
	var p = new(SidesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceParserRULE_sides

	return p
}

func (s *SidesContext) GetParser() antlr.Parser { return s.parser }

func (s *SidesContext) Id() antlr.TerminalNode {
	return s.GetToken(DiceParserId, 0)
}

func (s *SidesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SidesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SidesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.EnterSides(s)
	}
}

func (s *SidesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.ExitSides(s)
	}
}

func (s *SidesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DiceVisitor:
		return t.VisitSides(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *DiceParser) Sides() (localctx ISidesContext) {
	localctx = NewSidesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DiceParserRULE_sides)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)
		p.Match(DiceParserId)
	}



	return localctx
}


// IModifierContext is an interface to support dynamic dispatch.
type IModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModifierContext differentiates from other interfaces.
	IsModifierContext()
}

type ModifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModifierContext() *ModifierContext {
	var p = new(ModifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DiceParserRULE_modifier
	return p
}

func (*ModifierContext) IsModifierContext() {}

func NewModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModifierContext {
	var p = new(ModifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceParserRULE_modifier

	return p
}

func (s *ModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *ModifierContext) SIGN() antlr.TerminalNode {
	return s.GetToken(DiceParserSIGN, 0)
}

func (s *ModifierContext) Integer() antlr.TerminalNode {
	return s.GetToken(DiceParserInteger, 0)
}

func (s *ModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.EnterModifier(s)
	}
}

func (s *ModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.ExitModifier(s)
	}
}

func (s *ModifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DiceVisitor:
		return t.VisitModifier(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *DiceParser) Modifier() (localctx IModifierContext) {
	localctx = NewModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DiceParserRULE_modifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(DiceParserSIGN)
	}
	{
		p.SetState(39)
		p.Match(DiceParserInteger)
	}



	return localctx
}


// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DiceParserRULE_parameter
	return p
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) Integer() antlr.TerminalNode {
	return s.GetToken(DiceParserInteger, 0)
}

func (s *ParameterContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(DiceParserStringLiteral, 0)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DiceVisitor:
		return t.VisitParameter(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *DiceParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DiceParserRULE_parameter)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DiceParserInteger || _la == DiceParserStringLiteral) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DiceParserRULE_parameters
	return p
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) AllParameter() []IParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParameterContext)(nil)).Elem())
	var tst = make([]IParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParameterContext)
		}
	}

	return tst
}

func (s *ParametersContext) Parameter(i int) IParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *ParametersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(DiceParserCOMMA)
}

func (s *ParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(DiceParserCOMMA, i)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.EnterParameters(s)
	}
}

func (s *ParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.ExitParameters(s)
	}
}

func (s *ParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DiceVisitor:
		return t.VisitParameters(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *DiceParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DiceParserRULE_parameters)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(43)
				p.Parameter()
			}
			{
				p.SetState(44)
				p.Match(DiceParserCOMMA)
			}


		}
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}
	{
		p.SetState(51)
		p.Parameter()
	}



	return localctx
}


// IFuncnameContext is an interface to support dynamic dispatch.
type IFuncnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncnameContext differentiates from other interfaces.
	IsFuncnameContext()
}

type FuncnameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncnameContext() *FuncnameContext {
	var p = new(FuncnameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DiceParserRULE_funcname
	return p
}

func (*FuncnameContext) IsFuncnameContext() {}

func NewFuncnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncnameContext {
	var p = new(FuncnameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceParserRULE_funcname

	return p
}

func (s *FuncnameContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncnameContext) Id() antlr.TerminalNode {
	return s.GetToken(DiceParserId, 0)
}

func (s *FuncnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FuncnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.EnterFuncname(s)
	}
}

func (s *FuncnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.ExitFuncname(s)
	}
}

func (s *FuncnameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DiceVisitor:
		return t.VisitFuncname(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *DiceParser) Funcname() (localctx IFuncnameContext) {
	localctx = NewFuncnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DiceParserRULE_funcname)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(DiceParserId)
	}



	return localctx
}


// IFunccallContext is an interface to support dynamic dispatch.
type IFunccallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunccallContext differentiates from other interfaces.
	IsFunccallContext()
}

type FunccallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunccallContext() *FunccallContext {
	var p = new(FunccallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DiceParserRULE_funccall
	return p
}

func (*FunccallContext) IsFunccallContext() {}

func NewFunccallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunccallContext {
	var p = new(FunccallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceParserRULE_funccall

	return p
}

func (s *FunccallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunccallContext) Funcname() IFuncnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncnameContext)
}

func (s *FunccallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(DiceParserLPAREN, 0)
}

func (s *FunccallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(DiceParserRPAREN, 0)
}

func (s *FunccallContext) Parameters() IParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *FunccallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunccallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FunccallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.EnterFunccall(s)
	}
}

func (s *FunccallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.ExitFunccall(s)
	}
}

func (s *FunccallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DiceVisitor:
		return t.VisitFunccall(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *DiceParser) Funccall() (localctx IFunccallContext) {
	localctx = NewFunccallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DiceParserRULE_funccall)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
		p.Funcname()
	}
	{
		p.SetState(56)
		p.Match(DiceParserLPAREN)
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == DiceParserInteger || _la == DiceParserStringLiteral {
		{
			p.SetState(57)
			p.Parameters()
		}

	}
	{
		p.SetState(60)
		p.Match(DiceParserRPAREN)
	}



	return localctx
}


