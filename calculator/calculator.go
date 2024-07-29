package calculator

import (
	"fmt"
	"strconv"

	"calculator/parsing"

	"github.com/antlr4-go/antlr/v4"
)

type Result struct {
	Value float64
	Error error
}

// Define the CalculatorVisitor struct
type CalculatorVisitor struct {
	*parsing.BaseCalculatorVisitor
	memory map[string]float64
}

// Create a new CalculatorVisitor
func NewCalculatorVisitor() *CalculatorVisitor {
	return &CalculatorVisitor{
		BaseCalculatorVisitor: &parsing.BaseCalculatorVisitor{},
		memory:                make(map[string]float64),
	}
}

// Visit Assign Context
func (v *CalculatorVisitor) VisitAssign(ctx *parsing.AssignContext) interface{} {
	id := ctx.ID().GetText()
	result := v.Visit(ctx.Expr()).(Result)
	if result.Error != nil {
		return result
	}

	v.memory[id] = result.Value
	return Result{result.Value, nil}
}

// Visit Print Expression Context
func (v *CalculatorVisitor) VisitPrintExpr(ctx *parsing.PrintExprContext) interface{} {
	result := v.Visit(ctx.Expr()).(Result)
	if result.Error != nil {
		return result
	}
	fmt.Println(result.Value)
	return Result{result.Value, nil}
}

// Visit Multiplication and Division Context
func (v *CalculatorVisitor) VisitMulDiv(ctx *parsing.MulDivContext) interface{} {
	resultLeft := v.Visit(ctx.Expr(0)).(Result)
	resultRight := v.Visit(ctx.Expr(1)).(Result)

	if resultLeft.Error != nil {
		return Result{0, fmt.Errorf("error getting left value: %s", resultLeft.Error.Error())}
	}

	if resultRight.Error != nil {
		return Result{0, fmt.Errorf("error getting right value: %s", resultRight.Error.Error())}
	}

	if ctx.GetOp().GetTokenType() == parsing.CalculatorLexerMUL {
		return Result{resultLeft.Value * resultRight.Value, nil}
	}

	if ctx.GetOp().GetTokenType() == parsing.CalculatorLexerDIV {
		return Result{resultLeft.Value / resultRight.Value, nil}
	}

	return Result{0, fmt.Errorf("invalid operator: %s", ctx.GetOp().GetText())}
}

// Visit Addition and Subtraction Context
func (v *CalculatorVisitor) VisitAddSub(ctx *parsing.AddSubContext) interface{} {
	resultLeft := v.Visit(ctx.Expr(0)).(Result)
	resultRight := v.Visit(ctx.Expr(1)).(Result)

	if resultLeft.Error != nil {
		return Result{0, fmt.Errorf("error getting left value: %s", resultLeft.Error.Error())}
	}

	if resultRight.Error != nil {
		return Result{0, fmt.Errorf("error getting right value: %s", resultRight.Error.Error())}
	}

	if ctx.GetOp().GetTokenType() == parsing.CalculatorLexerADD {
		return Result{resultLeft.Value + resultRight.Value, nil}
	}

	if ctx.GetOp().GetTokenType() == parsing.CalculatorLexerSUB {
		return Result{resultLeft.Value - resultRight.Value, nil}
	}

	return Result{0, fmt.Errorf("invalid operator: %s", ctx.GetOp().GetText())}
}

// Visit Integer Context
func (v *CalculatorVisitor) VisitNumber(ctx *parsing.NumberContext) interface{} {
	value, err := strconv.ParseFloat(ctx.NUMBER().GetText(), 64)
	if err != nil {
		return Result{0, fmt.Errorf("error converting number to float: %s", err.Error())}
	}
	return Result{value, nil}
}

// Visit Identifier Context
func (v *CalculatorVisitor) VisitId(ctx *parsing.IdContext) interface{} {
	id := ctx.ID().GetText()
	if value, exists := v.memory[id]; exists {
		return Result{value, nil}
	}
	return Result{0, fmt.Errorf("cannot access variable %s", id)}
}

// Visit Parenthesis Context
func (v *CalculatorVisitor) VisitParens(ctx *parsing.ParensContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *CalculatorVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch t := tree.(type) {
	case *antlr.ErrorNodeImpl:
		return Result{0, fmt.Errorf("syntax error near '%s'", t.GetText())}
	default:
		if cr, ok := tree.Accept(v).(Result); ok {
			return cr
		}
	}

	return Result{0, fmt.Errorf("visit result not of type CalcReturn")}
}

func (v *CalculatorVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, n := range node.GetChildren() {
		cr := v.Visit(n.(antlr.ParseTree)).(Result)
		if cr.Error != nil {
			fmt.Printf("ERROR: %s\n", cr.Error)
			continue
		}
	}
	return Result{0, nil}
}

func (v *CalculatorVisitor) VisitProg(ctx *parsing.ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CalculatorVisitor) performExpression(expr string) {
	input := antlr.NewInputStream(expr)
	lexer := parsing.NewCalculatorLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parsing.NewCalculatorParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	tree := parser.Prog()
	v.Visit(tree)
}
