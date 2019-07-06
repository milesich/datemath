// Code generated from Datemath.g4 by ANTLR 4.7.2. DO NOT EDIT.

package datemath // Datemath
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseDatemathVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDatemathVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDatemathVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDatemathVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDatemathVisitor) VisitRound(ctx *RoundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDatemathVisitor) VisitBinary(ctx *BinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDatemathVisitor) VisitBuiltin(ctx *BuiltinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDatemathVisitor) VisitDuration(ctx *DurationContext) interface{} {
	return v.VisitChildren(ctx)
}
