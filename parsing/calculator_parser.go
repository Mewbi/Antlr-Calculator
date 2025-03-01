// Code generated from Calculator.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // Calculator
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CalculatorParser struct {
	*antlr.BaseParser
}

var CalculatorParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func calculatorParserInit() {
	staticData := &CalculatorParserStaticData
	staticData.LiteralNames = []string{
		"", "'='", "'('", "')'", "'*'", "'/'", "'+'", "'-'", "'=='", "'!='",
		"'<'", "'>'", "'<='", "'>='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "MUL", "DIV", "ADD", "SUB", "EQUAL", "DIFF", "LESS",
		"GREATER", "LESSEQ", "GREATEREQ", "ID", "NUMBER", "NEWLINE", "WS",
	}
	staticData.RuleNames = []string{
		"prog", "stat", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 17, 46, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 4, 0, 8, 8, 0,
		11, 0, 12, 0, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		3, 1, 21, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 30, 8,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 41, 8, 2,
		10, 2, 12, 2, 44, 9, 2, 1, 2, 0, 1, 4, 3, 0, 2, 4, 0, 3, 1, 0, 4, 5, 1,
		0, 6, 7, 1, 0, 8, 13, 50, 0, 7, 1, 0, 0, 0, 2, 20, 1, 0, 0, 0, 4, 29, 1,
		0, 0, 0, 6, 8, 3, 2, 1, 0, 7, 6, 1, 0, 0, 0, 8, 9, 1, 0, 0, 0, 9, 7, 1,
		0, 0, 0, 9, 10, 1, 0, 0, 0, 10, 1, 1, 0, 0, 0, 11, 12, 3, 4, 2, 0, 12,
		13, 5, 16, 0, 0, 13, 21, 1, 0, 0, 0, 14, 15, 5, 14, 0, 0, 15, 16, 5, 1,
		0, 0, 16, 17, 3, 4, 2, 0, 17, 18, 5, 16, 0, 0, 18, 21, 1, 0, 0, 0, 19,
		21, 5, 16, 0, 0, 20, 11, 1, 0, 0, 0, 20, 14, 1, 0, 0, 0, 20, 19, 1, 0,
		0, 0, 21, 3, 1, 0, 0, 0, 22, 23, 6, 2, -1, 0, 23, 30, 5, 15, 0, 0, 24,
		30, 5, 14, 0, 0, 25, 26, 5, 2, 0, 0, 26, 27, 3, 4, 2, 0, 27, 28, 5, 3,
		0, 0, 28, 30, 1, 0, 0, 0, 29, 22, 1, 0, 0, 0, 29, 24, 1, 0, 0, 0, 29, 25,
		1, 0, 0, 0, 30, 42, 1, 0, 0, 0, 31, 32, 10, 6, 0, 0, 32, 33, 7, 0, 0, 0,
		33, 41, 3, 4, 2, 7, 34, 35, 10, 5, 0, 0, 35, 36, 7, 1, 0, 0, 36, 41, 3,
		4, 2, 6, 37, 38, 10, 4, 0, 0, 38, 39, 7, 2, 0, 0, 39, 41, 3, 4, 2, 5, 40,
		31, 1, 0, 0, 0, 40, 34, 1, 0, 0, 0, 40, 37, 1, 0, 0, 0, 41, 44, 1, 0, 0,
		0, 42, 40, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 5, 1, 0, 0, 0, 44, 42, 1,
		0, 0, 0, 5, 9, 20, 29, 40, 42,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CalculatorParserInit initializes any static state used to implement CalculatorParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCalculatorParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CalculatorParserInit() {
	staticData := &CalculatorParserStaticData
	staticData.once.Do(calculatorParserInit)
}

// NewCalculatorParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCalculatorParser(input antlr.TokenStream) *CalculatorParser {
	CalculatorParserInit()
	this := new(CalculatorParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CalculatorParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Calculator.g4"

	return this
}

// CalculatorParser tokens.
const (
	CalculatorParserEOF       = antlr.TokenEOF
	CalculatorParserT__0      = 1
	CalculatorParserT__1      = 2
	CalculatorParserT__2      = 3
	CalculatorParserMUL       = 4
	CalculatorParserDIV       = 5
	CalculatorParserADD       = 6
	CalculatorParserSUB       = 7
	CalculatorParserEQUAL     = 8
	CalculatorParserDIFF      = 9
	CalculatorParserLESS      = 10
	CalculatorParserGREATER   = 11
	CalculatorParserLESSEQ    = 12
	CalculatorParserGREATEREQ = 13
	CalculatorParserID        = 14
	CalculatorParserNUMBER    = 15
	CalculatorParserNEWLINE   = 16
	CalculatorParserWS        = 17
)

// CalculatorParser rules.
const (
	CalculatorParserRULE_prog = 0
	CalculatorParserRULE_stat = 1
	CalculatorParserRULE_expr = 2
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStat() []IStatContext
	Stat(i int) IStatContext

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalculatorParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalculatorParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllStat() []IStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatContext); ok {
			len++
		}
	}

	tst := make([]IStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatContext); ok {
			tst[i] = t.(IStatContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Stat(i int) IStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalculatorParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CalculatorParserRULE_prog)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(7)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&114692) != 0) {
		{
			p.SetState(6)
			p.Stat()
		}

		p.SetState(9)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalculatorParserRULE_stat
	return p
}

func InitEmptyStatContext(p *StatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalculatorParserRULE_stat
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) CopyAll(ctx *StatContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BlankContext struct {
	StatContext
}

func NewBlankContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlankContext {
	var p = new(BlankContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *BlankContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlankContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(CalculatorParserNEWLINE, 0)
}

func (s *BlankContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.EnterBlank(s)
	}
}

func (s *BlankContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.ExitBlank(s)
	}
}

func (s *BlankContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitBlank(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintExprContext struct {
	StatContext
}

func NewPrintExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintExprContext {
	var p = new(PrintExprContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *PrintExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintExprContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(CalculatorParserNEWLINE, 0)
}

func (s *PrintExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.EnterPrintExpr(s)
	}
}

func (s *PrintExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.ExitPrintExpr(s)
	}
}

func (s *PrintExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitPrintExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignContext struct {
	StatContext
}

func NewAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignContext {
	var p = new(AssignContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ID() antlr.TerminalNode {
	return s.GetToken(CalculatorParserID, 0)
}

func (s *AssignContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(CalculatorParserNEWLINE, 0)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (s *AssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalculatorParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CalculatorParserRULE_stat)
	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPrintExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(11)
			p.expr(0)
		}
		{
			p.SetState(12)
			p.Match(CalculatorParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(14)
			p.Match(CalculatorParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(15)
			p.Match(CalculatorParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(16)
			p.expr(0)
		}
		{
			p.SetState(17)
			p.Match(CalculatorParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewBlankContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(19)
			p.Match(CalculatorParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalculatorParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalculatorParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NumberContext struct {
	ExprContext
}

func NewNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberContext {
	var p = new(NumberContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CalculatorParserNUMBER, 0)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParensContext struct {
	ExprContext
}

func NewParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensContext {
	var p = new(ParensContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.EnterParens(s)
	}
}

func (s *ParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.ExitParens(s)
	}
}

func (s *ParensContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitParens(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulDivContext struct {
	ExprContext
	op antlr.Token
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MulDivContext) GetOp() antlr.Token { return s.op }

func (s *MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MulDivContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MulDivContext) MUL() antlr.TerminalNode {
	return s.GetToken(CalculatorParserMUL, 0)
}

func (s *MulDivContext) DIV() antlr.TerminalNode {
	return s.GetToken(CalculatorParserDIV, 0)
}

func (s *MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.EnterMulDiv(s)
	}
}

func (s *MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.ExitMulDiv(s)
	}
}

func (s *MulDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitMulDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubContext struct {
	ExprContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AddSubContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddSubContext) ADD() antlr.TerminalNode {
	return s.GetToken(CalculatorParserADD, 0)
}

func (s *AddSubContext) SUB() antlr.TerminalNode {
	return s.GetToken(CalculatorParserSUB, 0)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.ExitAddSub(s)
	}
}

func (s *AddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type OpRelContext struct {
	ExprContext
	op antlr.Token
}

func NewOpRelContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpRelContext {
	var p = new(OpRelContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *OpRelContext) GetOp() antlr.Token { return s.op }

func (s *OpRelContext) SetOp(v antlr.Token) { s.op = v }

func (s *OpRelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpRelContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *OpRelContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OpRelContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(CalculatorParserEQUAL, 0)
}

func (s *OpRelContext) DIFF() antlr.TerminalNode {
	return s.GetToken(CalculatorParserDIFF, 0)
}

func (s *OpRelContext) LESS() antlr.TerminalNode {
	return s.GetToken(CalculatorParserLESS, 0)
}

func (s *OpRelContext) GREATER() antlr.TerminalNode {
	return s.GetToken(CalculatorParserGREATER, 0)
}

func (s *OpRelContext) LESSEQ() antlr.TerminalNode {
	return s.GetToken(CalculatorParserLESSEQ, 0)
}

func (s *OpRelContext) GREATEREQ() antlr.TerminalNode {
	return s.GetToken(CalculatorParserGREATEREQ, 0)
}

func (s *OpRelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.EnterOpRel(s)
	}
}

func (s *OpRelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.ExitOpRel(s)
	}
}

func (s *OpRelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitOpRel(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdContext struct {
	ExprContext
}

func NewIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdContext {
	var p = new(IdContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ID() antlr.TerminalNode {
	return s.GetToken(CalculatorParserID, 0)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.ExitId(s)
	}
}

func (s *IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitId(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalculatorParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *CalculatorParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, CalculatorParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CalculatorParserNUMBER:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(23)
			p.Match(CalculatorParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CalculatorParserID:
		localctx = NewIdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(24)
			p.Match(CalculatorParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CalculatorParserT__1:
		localctx = NewParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(25)
			p.Match(CalculatorParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(26)
			p.expr(0)
		}
		{
			p.SetState(27)
			p.Match(CalculatorParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(40)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalculatorParserRULE_expr)
				p.SetState(31)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(32)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CalculatorParserMUL || _la == CalculatorParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(33)
					p.expr(7)
				}

			case 2:
				localctx = NewAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalculatorParserRULE_expr)
				p.SetState(34)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(35)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CalculatorParserADD || _la == CalculatorParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(36)
					p.expr(6)
				}

			case 3:
				localctx = NewOpRelContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalculatorParserRULE_expr)
				p.SetState(37)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(38)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpRelContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16128) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpRelContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(39)
					p.expr(5)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *CalculatorParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CalculatorParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
