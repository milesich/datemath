// Code generated from Datemath.g4 by ANTLR 4.7.2. DO NOT EDIT.

package datemath // Datemath
import "github.com/antlr/antlr4/runtime/Go/antlr"

// DatemathListener is a complete listener for a parse tree produced by DatemathParser.
type DatemathListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterIdentifier is called when entering the Identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterLiteral is called when entering the Literal production.
	EnterLiteral(c *LiteralContext)

	// EnterRound is called when entering the Round production.
	EnterRound(c *RoundContext)

	// EnterBinary is called when entering the Binary production.
	EnterBinary(c *BinaryContext)

	// EnterBuiltin is called when entering the Builtin production.
	EnterBuiltin(c *BuiltinContext)

	// EnterDuration is called when entering the duration production.
	EnterDuration(c *DurationContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitIdentifier is called when exiting the Identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitLiteral is called when exiting the Literal production.
	ExitLiteral(c *LiteralContext)

	// ExitRound is called when exiting the Round production.
	ExitRound(c *RoundContext)

	// ExitBinary is called when exiting the Binary production.
	ExitBinary(c *BinaryContext)

	// ExitBuiltin is called when exiting the Builtin production.
	ExitBuiltin(c *BuiltinContext)

	// ExitDuration is called when exiting the duration production.
	ExitDuration(c *DurationContext)
}
