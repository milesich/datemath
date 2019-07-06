// Code generated from Datemath.g4 by ANTLR 4.7.2. DO NOT EDIT.

package datemath // Datemath
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by DatemathParser.
type DatemathVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DatemathParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by DatemathParser#Identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by DatemathParser#Literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by DatemathParser#Round.
	VisitRound(ctx *RoundContext) interface{}

	// Visit a parse tree produced by DatemathParser#Binary.
	VisitBinary(ctx *BinaryContext) interface{}

	// Visit a parse tree produced by DatemathParser#Builtin.
	VisitBuiltin(ctx *BuiltinContext) interface{}

	// Visit a parse tree produced by DatemathParser#duration.
	VisitDuration(ctx *DurationContext) interface{}
}
