// Code generated from Calculator.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // Calculator
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by CalculatorParser.
type CalculatorVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CalculatorParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by CalculatorParser#printExpr.
	VisitPrintExpr(ctx *PrintExprContext) interface{}

	// Visit a parse tree produced by CalculatorParser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by CalculatorParser#blank.
	VisitBlank(ctx *BlankContext) interface{}

	// Visit a parse tree produced by CalculatorParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by CalculatorParser#parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by CalculatorParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by CalculatorParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by CalculatorParser#OpRel.
	VisitOpRel(ctx *OpRelContext) interface{}

	// Visit a parse tree produced by CalculatorParser#id.
	VisitId(ctx *IdContext) interface{}
}
