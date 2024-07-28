// Code generated from Calculator.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // Calculator
import "github.com/antlr4-go/antlr/v4"

// CalculatorListener is a complete listener for a parse tree produced by CalculatorParser.
type CalculatorListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterPrintExpr is called when entering the printExpr production.
	EnterPrintExpr(c *PrintExprContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterBlank is called when entering the blank production.
	EnterBlank(c *BlankContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterParens is called when entering the parens production.
	EnterParens(c *ParensContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitPrintExpr is called when exiting the printExpr production.
	ExitPrintExpr(c *PrintExprContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitBlank is called when exiting the blank production.
	ExitBlank(c *BlankContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitParens is called when exiting the parens production.
	ExitParens(c *ParensContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)
}
