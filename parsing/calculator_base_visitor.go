// Code generated from Calculator.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // Calculator
import "github.com/antlr4-go/antlr/v4"

type BaseCalculatorVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCalculatorVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitPrintExpr(ctx *PrintExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitAssign(ctx *AssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitBlank(ctx *BlankContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitParens(ctx *ParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitOpRel(ctx *OpRelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}
