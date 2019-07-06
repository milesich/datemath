// Code generated from Datemath.g4 by ANTLR 4.7.2. DO NOT EDIT.

package datemath // Datemath
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDatemathListener is a complete listener for a parse tree produced by DatemathParser.
type BaseDatemathListener struct{}

var _ DatemathListener = &BaseDatemathListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDatemathListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDatemathListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDatemathListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDatemathListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseDatemathListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseDatemathListener) ExitStart(ctx *StartContext) {}

// EnterIdentifier is called when production Identifier is entered.
func (s *BaseDatemathListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production Identifier is exited.
func (s *BaseDatemathListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterLiteral is called when production Literal is entered.
func (s *BaseDatemathListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production Literal is exited.
func (s *BaseDatemathListener) ExitLiteral(ctx *LiteralContext) {}

// EnterRound is called when production Round is entered.
func (s *BaseDatemathListener) EnterRound(ctx *RoundContext) {}

// ExitRound is called when production Round is exited.
func (s *BaseDatemathListener) ExitRound(ctx *RoundContext) {}

// EnterBinary is called when production Binary is entered.
func (s *BaseDatemathListener) EnterBinary(ctx *BinaryContext) {}

// ExitBinary is called when production Binary is exited.
func (s *BaseDatemathListener) ExitBinary(ctx *BinaryContext) {}

// EnterBuiltin is called when production Builtin is entered.
func (s *BaseDatemathListener) EnterBuiltin(ctx *BuiltinContext) {}

// ExitBuiltin is called when production Builtin is exited.
func (s *BaseDatemathListener) ExitBuiltin(ctx *BuiltinContext) {}

// EnterDuration is called when production duration is entered.
func (s *BaseDatemathListener) EnterDuration(ctx *DurationContext) {}

// ExitDuration is called when production duration is exited.
func (s *BaseDatemathListener) ExitDuration(ctx *DurationContext) {}
