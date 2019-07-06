package datemath

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/araddon/dateparse"
)

//go:generate antlr -Dlanguage=Go -listener -visitor -o . -package datemath Datemath.g4

type parser struct {
	*BaseDatemathListener

	env map[string]time.Time

	errors []error
	ts     time.Time

	// duration
	d             time.Duration
	ddy, ddm, ddd int
}

func Eval(input string) (time.Time, error) {
	return EvalWithEnv(input, map[string]time.Time{})
}

func EvalWithEnv(input string, env map[string]time.Time) (time.Time, error) {
	is := antlr.NewInputStream(input)

	lexer := NewDatemathLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	datemath := NewDatemathParser(stream)

	p := &parser{
		env: env,
	}

	lexer.RemoveErrorListeners()
	datemath.RemoveErrorListeners()
	lexer.AddErrorListener(p)
	datemath.AddErrorListener(p)

	antlr.ParseTreeWalkerDefault.Walk(p, datemath.Start())

	if len(p.errors) > 0 {
		return time.Time{}, p.errors[0]
	}
	return p.ts, nil
}

func (p *parser) ExitDuration(c *DurationContext) {
	var err error
	if terminalnode, ok := c.GetChild(0).(*antlr.TerminalNodeImpl); ok {
		switch terminalnode.GetSymbol().GetTokenType() {
		case DatemathLexerEsDuration:
			txt := c.GetText()
			num := 1
			if len(txt) > 1 {
				num, err = strconv.Atoi(txt[:len(txt)-1])
			}
			if err == nil {
				switch txt[len(txt)-1] {
				case 'y':
					p.ddy = num
				case 'M':
					p.ddm = num
				case 'd':
					p.ddd = num
				case 'w':
					p.ddd = 7 * num
				case 'H':
					p.d = time.Duration(num) * time.Hour
				case 'h':
					p.d = time.Duration(num) * time.Hour
				case 'm':
					p.d = time.Duration(num) * time.Minute
				case 's':
					p.d = time.Duration(num) * time.Second
				}
			}
		case DatemathLexerGoDuration:
			p.d, err = time.ParseDuration(c.GetText())
		}

	}
	if err != nil {
		p.errors = append(p.errors, err)
	}
}

func (p *parser) ExitBinary(c *BinaryContext) {
	op := c.GetOp().GetText()
	switch op {
	case "-":
		p.ts = p.ts.AddDate(-p.ddy, -p.ddm, -p.ddd)
		p.ts = p.ts.Add(-1 * p.d)
	case "+":
		p.ts = p.ts.AddDate(p.ddy, p.ddm, p.ddd)
		p.ts = p.ts.Add(p.d)
	}

	p.d, p.ddy, p.ddm, p.ddd = 0, 0, 0, 0
}

func (p *parser) ExitRound(c *RoundContext) {
	p.ts = p.ts.Truncate(p.d)
	if p.ddy != 0 {
		p.ts = time.Date(p.ts.Year()/p.ddy*p.ddy, 1, 1, 0, 0, 0, 0, time.UTC)
	} else if p.ddm != 0 {
		p.ts = time.Date(p.ts.Year(), time.Month((int(p.ts.Month())-1)/p.ddm*p.ddm+1), 1, 0, 0, 0, 0, time.UTC)
	} else if p.ddd != 0 {
		if p.ddd%7 == 0 {
			// Special treatment for weeks
			weekday := int(p.ts.Weekday())
			if weekday == 0 {
				weekday = 7
			}
			_, w := p.ts.ISOWeek()
			p.ts = p.ts.AddDate(0, 0, -weekday+1-w%(p.ddd/7)*7).Truncate(86400 * time.Second)
		} else {
			p.ts = time.Date(p.ts.Year(), p.ts.Month(), p.ts.Day()/p.ddd*p.ddd, 0, 0, 0, 0, time.UTC)
		}
	}
	p.d, p.ddy, p.ddm, p.ddd = 0, 0, 0, 0
}

func (p *parser) ExitBuiltin(c *BuiltinContext) {
	name := c.GetText()
	switch name {
	case "now":
		if ts, ok := p.env["now"]; ok {
			p.ts = ts
			return
		}
		p.ts = time.Now().UTC()
	}
}

func (p *parser) ExitIdentifier(c *IdentifierContext) {
	name := c.GetText()
	if ts, ok := p.env[name]; ok {
		p.ts = ts
		return
	}
	p.errors = append(p.errors, fmt.Errorf("Unknown identifier: '%s'", name))
}

func (p *parser) ExitLiteral(c *LiteralContext) {
	var err error
	p.ts, err = dateparse.ParseAny(strings.Trim(c.GetText(), "'"))
	if err != nil {
		p.errors = append(p.errors, err)
	}
}

func (p *parser) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	p.errors = append(p.errors, fmt.Errorf(":%d:%d:%s", line, column, msg))
}

func (p *parser) ReportAmbiguity(_ antlr.Parser, _ *antlr.DFA, _, _ int, _ bool, _ *antlr.BitSet, _ antlr.ATNConfigSet) {
}

func (p *parser) ReportAttemptingFullContext(_ antlr.Parser, _ *antlr.DFA, _, _ int, _ *antlr.BitSet, _ antlr.ATNConfigSet) {
}

func (p *parser) ReportContextSensitivity(_ antlr.Parser, _ *antlr.DFA, _, _, _ int, _ antlr.ATNConfigSet) {
}
