package calculator

import (
	"fmt"
	"os"
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
	value := v.Visit(ctx.Expr()).(Result).Value
	v.memory[id] = value
	fmt.Println(id, value)
	return Result{value, nil}
}

// Visit Print Expression Context
func (v *CalculatorVisitor) VisitPrintExpr(ctx *parsing.PrintExprContext) interface{} {
	value := v.Visit(ctx.Expr()).(Result).Value
	fmt.Println(value)
	return Result{value, nil}
}

// Visit Multiplication and Division Context
func (v *CalculatorVisitor) VisitMulDiv(ctx *parsing.MulDivContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(Result).Value
	right := v.Visit(ctx.Expr(1)).(Result).Value
	if ctx.GetOp().GetTokenType() == parsing.CalculatorLexerMUL {
		return Result{left * right, nil}
	}
	return Result{left / right, nil}
}

// Visit Addition and Subtraction Context
func (v *CalculatorVisitor) VisitAddSub(ctx *parsing.AddSubContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(Result).Value
	right := v.Visit(ctx.Expr(1)).(Result).Value
	if ctx.GetOp().GetTokenType() == parsing.CalculatorLexerADD {
		return Result{left + right, nil}
	}
	return Result{left - right, nil}
}

// Visit Integer Context
func (v *CalculatorVisitor) VisitNumber(ctx *parsing.NumberContext) interface{} {
	value, _ := strconv.ParseFloat(ctx.NUMBER().GetText(), 64)
	return Result{value, nil}
}

// Visit Identifier Context
func (v *CalculatorVisitor) VisitId(ctx *parsing.IdContext) interface{} {
	id := ctx.ID().GetText()
	if value, exists := v.memory[id]; exists {
		return Result{value, nil}
	}
	return Result{0, nil}
}

// Visit Parenthesis Context
func (v *CalculatorVisitor) VisitParens(ctx *parsing.ParensContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *CalculatorVisitor) Visit(tree antlr.ParseTree) interface{} {
	// fmt.Printf("visit input type: %v\n", reflect.TypeOf(tree))

	switch t := tree.(type) {
	case *antlr.ErrorNodeImpl:
		return fmt.Errorf("syntax error near '%s'", t.GetText())
	default:
		if cr, ok := tree.Accept(v).(Result); ok {
			return cr
		}
	}

	return Result{0, fmt.Errorf("visit result not of type CalcReturn")}
}

func (v *CalculatorVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, n := range node.GetChildren() {
		if false {
			fmt.Printf("> %s", n.(antlr.ParseTree).GetText())
		}
		cr := v.Visit(n.(antlr.ParseTree)).(Result)
		if cr.Error != nil {
			// if errors.Is(cr.Error, ErrorBlank) {
			// 	continue
			// }
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", cr.Error)
			continue
		}
		fmt.Printf("  %f\n", cr.Value)
	}
	return Result{0, nil}
}

func (v *CalculatorVisitor) VisitProg(ctx *parsing.ProgContext) interface{} {
	fmt.Printf("Calculating Programm: %s", ctx.GetText())
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
