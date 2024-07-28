// Code generated from Calculator.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // Calculator
import "github.com/antlr4-go/antlr/v4"

// BaseCalculatorListener is a complete listener for a parse tree produced by CalculatorParser.
type BaseCalculatorListener struct{}

var _ CalculatorListener = &BaseCalculatorListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCalculatorListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCalculatorListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCalculatorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCalculatorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseCalculatorListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseCalculatorListener) ExitProg(ctx *ProgContext) {}

// EnterPrintExpr is called when production printExpr is entered.
func (s *BaseCalculatorListener) EnterPrintExpr(ctx *PrintExprContext) {}

// ExitPrintExpr is called when production printExpr is exited.
func (s *BaseCalculatorListener) ExitPrintExpr(ctx *PrintExprContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseCalculatorListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseCalculatorListener) ExitAssign(ctx *AssignContext) {}

// EnterBlank is called when production blank is entered.
func (s *BaseCalculatorListener) EnterBlank(ctx *BlankContext) {}

// ExitBlank is called when production blank is exited.
func (s *BaseCalculatorListener) ExitBlank(ctx *BlankContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseCalculatorListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseCalculatorListener) ExitNumber(ctx *NumberContext) {}

// EnterParens is called when production parens is entered.
func (s *BaseCalculatorListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production parens is exited.
func (s *BaseCalculatorListener) ExitParens(ctx *ParensContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseCalculatorListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseCalculatorListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseCalculatorListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseCalculatorListener) ExitAddSub(ctx *AddSubContext) {}

// EnterId is called when production id is entered.
func (s *BaseCalculatorListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseCalculatorListener) ExitId(ctx *IdContext) {}
