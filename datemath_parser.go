// Code generated from Datemath.g4 by ANTLR 4.7.2. DO NOT EDIT.

package datemath // Datemath
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 32, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 5, 3, 17, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 25,
	10, 3, 12, 3, 14, 3, 28, 11, 3, 3, 4, 3, 4, 3, 4, 2, 3, 4, 5, 2, 4, 6,
	2, 4, 3, 2, 6, 7, 3, 2, 3, 4, 2, 33, 2, 8, 3, 2, 2, 2, 4, 16, 3, 2, 2,
	2, 6, 29, 3, 2, 2, 2, 8, 9, 5, 4, 3, 2, 9, 10, 7, 2, 2, 3, 10, 3, 3, 2,
	2, 2, 11, 12, 8, 3, 1, 2, 12, 17, 7, 9, 2, 2, 13, 17, 7, 12, 2, 2, 14,
	17, 7, 5, 2, 2, 15, 17, 7, 14, 2, 2, 16, 11, 3, 2, 2, 2, 16, 13, 3, 2,
	2, 2, 16, 14, 3, 2, 2, 2, 16, 15, 3, 2, 2, 2, 17, 26, 3, 2, 2, 2, 18, 19,
	12, 7, 2, 2, 19, 20, 9, 2, 2, 2, 20, 25, 5, 6, 4, 2, 21, 22, 12, 6, 2,
	2, 22, 23, 7, 8, 2, 2, 23, 25, 5, 6, 4, 2, 24, 18, 3, 2, 2, 2, 24, 21,
	3, 2, 2, 2, 25, 28, 3, 2, 2, 2, 26, 24, 3, 2, 2, 2, 26, 27, 3, 2, 2, 2,
	27, 5, 3, 2, 2, 2, 28, 26, 3, 2, 2, 2, 29, 30, 9, 3, 2, 2, 30, 7, 3, 2,
	2, 2, 5, 16, 24, 26,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "'+'", "'-'", "'/'",
}
var symbolicNames = []string{
	"", "EsDuration", "GoDuration", "Literal", "Plus", "Minus", "Round", "Builtins",
	"IntegerLiteral", "FloatLiteral", "Identifier", "WhiteSpaces", "DateLiteral",
}

var ruleNames = []string{
	"start", "timestamp", "duration",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type DatemathParser struct {
	*antlr.BaseParser
}

func NewDatemathParser(input antlr.TokenStream) *DatemathParser {
	this := new(DatemathParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Datemath.g4"

	return this
}

// DatemathParser tokens.
const (
	DatemathParserEOF            = antlr.TokenEOF
	DatemathParserEsDuration     = 1
	DatemathParserGoDuration     = 2
	DatemathParserLiteral        = 3
	DatemathParserPlus           = 4
	DatemathParserMinus          = 5
	DatemathParserRound          = 6
	DatemathParserBuiltins       = 7
	DatemathParserIntegerLiteral = 8
	DatemathParserFloatLiteral   = 9
	DatemathParserIdentifier     = 10
	DatemathParserWhiteSpaces    = 11
	DatemathParserDateLiteral    = 12
)

// DatemathParser rules.
const (
	DatemathParserRULE_start     = 0
	DatemathParserRULE_timestamp = 1
	DatemathParserRULE_duration  = 2
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE returns the e rule contexts.
	GetE() ITimestampContext

	// SetE sets the e rule contexts.
	SetE(ITimestampContext)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	e      ITimestampContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DatemathParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DatemathParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) GetE() ITimestampContext { return s.e }

func (s *StartContext) SetE(v ITimestampContext) { s.e = v }

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(DatemathParserEOF, 0)
}

func (s *StartContext) Timestamp() ITimestampContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimestampContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimestampContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DatemathListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DatemathListener); ok {
		listenerT.ExitStart(s)
	}
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DatemathVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DatemathParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DatemathParserRULE_start)

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
		p.SetState(6)

		var _x = p.timestamp(0)

		localctx.(*StartContext).e = _x
	}
	{
		p.SetState(7)
		p.Match(DatemathParserEOF)
	}

	return localctx
}

// ITimestampContext is an interface to support dynamic dispatch.
type ITimestampContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimestampContext differentiates from other interfaces.
	IsTimestampContext()
}

type TimestampContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimestampContext() *TimestampContext {
	var p = new(TimestampContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DatemathParserRULE_timestamp
	return p
}

func (*TimestampContext) IsTimestampContext() {}

func NewTimestampContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimestampContext {
	var p = new(TimestampContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DatemathParserRULE_timestamp

	return p
}

func (s *TimestampContext) GetParser() antlr.Parser { return s.parser }

func (s *TimestampContext) CopyFrom(ctx *TimestampContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TimestampContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimestampContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IdentifierContext struct {
	*TimestampContext
}

func NewIdentifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierContext {
	var p = new(IdentifierContext)

	p.TimestampContext = NewEmptyTimestampContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TimestampContext))

	return p
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) Identifier() antlr.TerminalNode {
	return s.GetToken(DatemathParserIdentifier, 0)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DatemathListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DatemathListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DatemathVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

type DateLiteralContext struct {
	*TimestampContext
}

func NewDateLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateLiteralContext {
	var p = new(DateLiteralContext)

	p.TimestampContext = NewEmptyTimestampContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TimestampContext))

	return p
}

func (s *DateLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateLiteralContext) DateLiteral() antlr.TerminalNode {
	return s.GetToken(DatemathParserDateLiteral, 0)
}

func (s *DateLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DatemathListener); ok {
		listenerT.EnterDateLiteral(s)
	}
}

func (s *DateLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DatemathListener); ok {
		listenerT.ExitDateLiteral(s)
	}
}

func (s *DateLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DatemathVisitor:
		return t.VisitDateLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralContext struct {
	*TimestampContext
}

func NewLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralContext {
	var p = new(LiteralContext)

	p.TimestampContext = NewEmptyTimestampContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TimestampContext))

	return p
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) Literal() antlr.TerminalNode {
	return s.GetToken(DatemathParserLiteral, 0)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DatemathListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DatemathListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DatemathVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type RoundContext struct {
	*TimestampContext
}

func NewRoundContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RoundContext {
	var p = new(RoundContext)

	p.TimestampContext = NewEmptyTimestampContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TimestampContext))

	return p
}

func (s *RoundContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoundContext) Timestamp() ITimestampContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimestampContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimestampContext)
}

func (s *RoundContext) Round() antlr.TerminalNode {
	return s.GetToken(DatemathParserRound, 0)
}

func (s *RoundContext) Duration() IDurationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDurationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDurationContext)
}

func (s *RoundContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DatemathListener); ok {
		listenerT.EnterRound(s)
	}
}

func (s *RoundContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DatemathListener); ok {
		listenerT.ExitRound(s)
	}
}

func (s *RoundContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DatemathVisitor:
		return t.VisitRound(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinaryContext struct {
	*TimestampContext
	op antlr.Token
}

func NewBinaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryContext {
	var p = new(BinaryContext)

	p.TimestampContext = NewEmptyTimestampContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TimestampContext))

	return p
}

func (s *BinaryContext) GetOp() antlr.Token { return s.op }

func (s *BinaryContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryContext) Timestamp() ITimestampContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimestampContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimestampContext)
}

func (s *BinaryContext) Duration() IDurationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDurationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDurationContext)
}

func (s *BinaryContext) Plus() antlr.TerminalNode {
	return s.GetToken(DatemathParserPlus, 0)
}

func (s *BinaryContext) Minus() antlr.TerminalNode {
	return s.GetToken(DatemathParserMinus, 0)
}

func (s *BinaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DatemathListener); ok {
		listenerT.EnterBinary(s)
	}
}

func (s *BinaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DatemathListener); ok {
		listenerT.ExitBinary(s)
	}
}

func (s *BinaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DatemathVisitor:
		return t.VisitBinary(s)

	default:
		return t.VisitChildren(s)
	}
}

type BuiltinContext struct {
	*TimestampContext
}

func NewBuiltinContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BuiltinContext {
	var p = new(BuiltinContext)

	p.TimestampContext = NewEmptyTimestampContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TimestampContext))

	return p
}

func (s *BuiltinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuiltinContext) Builtins() antlr.TerminalNode {
	return s.GetToken(DatemathParserBuiltins, 0)
}

func (s *BuiltinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DatemathListener); ok {
		listenerT.EnterBuiltin(s)
	}
}

func (s *BuiltinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DatemathListener); ok {
		listenerT.ExitBuiltin(s)
	}
}

func (s *BuiltinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DatemathVisitor:
		return t.VisitBuiltin(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DatemathParser) Timestamp() (localctx ITimestampContext) {
	return p.timestamp(0)
}

func (p *DatemathParser) timestamp(_p int) (localctx ITimestampContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewTimestampContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITimestampContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, DatemathParserRULE_timestamp, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	p.SetState(14)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DatemathParserBuiltins:
		localctx = NewBuiltinContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(10)
			p.Match(DatemathParserBuiltins)
		}

	case DatemathParserIdentifier:
		localctx = NewIdentifierContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(11)
			p.Match(DatemathParserIdentifier)
		}

	case DatemathParserLiteral:
		localctx = NewLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(12)
			p.Match(DatemathParserLiteral)
		}

	case DatemathParserDateLiteral:
		localctx = NewDateLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(13)
			p.Match(DatemathParserDateLiteral)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(22)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryContext(p, NewTimestampContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DatemathParserRULE_timestamp)
				p.SetState(16)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(17)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == DatemathParserPlus || _la == DatemathParserMinus) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(18)
					p.Duration()
				}

			case 2:
				localctx = NewRoundContext(p, NewTimestampContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DatemathParserRULE_timestamp)
				p.SetState(19)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(20)
					p.Match(DatemathParserRound)
				}
				{
					p.SetState(21)
					p.Duration()
				}

			}

		}
		p.SetState(26)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IDurationContext is an interface to support dynamic dispatch.
type IDurationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDurationContext differentiates from other interfaces.
	IsDurationContext()
}

type DurationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDurationContext() *DurationContext {
	var p = new(DurationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DatemathParserRULE_duration
	return p
}

func (*DurationContext) IsDurationContext() {}

func NewDurationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DurationContext {
	var p = new(DurationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DatemathParserRULE_duration

	return p
}

func (s *DurationContext) GetParser() antlr.Parser { return s.parser }

func (s *DurationContext) EsDuration() antlr.TerminalNode {
	return s.GetToken(DatemathParserEsDuration, 0)
}

func (s *DurationContext) GoDuration() antlr.TerminalNode {
	return s.GetToken(DatemathParserGoDuration, 0)
}

func (s *DurationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DurationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DurationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DatemathListener); ok {
		listenerT.EnterDuration(s)
	}
}

func (s *DurationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DatemathListener); ok {
		listenerT.ExitDuration(s)
	}
}

func (s *DurationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DatemathVisitor:
		return t.VisitDuration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DatemathParser) Duration() (localctx IDurationContext) {
	localctx = NewDurationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DatemathParserRULE_duration)
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
		p.SetState(27)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DatemathParserEsDuration || _la == DatemathParserGoDuration) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *DatemathParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *TimestampContext = nil
		if localctx != nil {
			t = localctx.(*TimestampContext)
		}
		return p.Timestamp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *DatemathParser) Timestamp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
