// Code generated from Calculator.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type CalculatorLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var CalculatorLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func calculatorlexerLexerInit() {
	staticData := &CalculatorLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'='", "'('", "')'", "'*'", "'/'", "'+'", "'-'", "'=='", "'!='",
		"'<'", "'>'", "'<='", "'>='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "MUL", "DIV", "ADD", "SUB", "EQUAL", "DIFF", "LESS",
		"GREATER", "LESSEQ", "GREATEREQ", "ID", "NUMBER", "NEWLINE", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "MUL", "DIV", "ADD", "SUB", "EQUAL", "DIFF",
		"LESS", "GREATER", "LESSEQ", "GREATEREQ", "ID", "NUMBER", "NEWLINE",
		"WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 17, 97, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13,
		1, 13, 5, 13, 68, 8, 13, 10, 13, 12, 13, 71, 9, 13, 1, 14, 4, 14, 74, 8,
		14, 11, 14, 12, 14, 75, 1, 14, 1, 14, 4, 14, 80, 8, 14, 11, 14, 12, 14,
		81, 3, 14, 84, 8, 14, 1, 15, 3, 15, 87, 8, 15, 1, 15, 1, 15, 1, 16, 4,
		16, 92, 8, 16, 11, 16, 12, 16, 93, 1, 16, 1, 16, 0, 0, 17, 1, 1, 3, 2,
		5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 1, 0, 4, 3, 0, 65, 90, 95, 95, 97,
		122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 2, 0, 9, 9, 32,
		32, 102, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 1, 35, 1, 0, 0, 0, 3, 37, 1, 0, 0,
		0, 5, 39, 1, 0, 0, 0, 7, 41, 1, 0, 0, 0, 9, 43, 1, 0, 0, 0, 11, 45, 1,
		0, 0, 0, 13, 47, 1, 0, 0, 0, 15, 49, 1, 0, 0, 0, 17, 52, 1, 0, 0, 0, 19,
		55, 1, 0, 0, 0, 21, 57, 1, 0, 0, 0, 23, 59, 1, 0, 0, 0, 25, 62, 1, 0, 0,
		0, 27, 65, 1, 0, 0, 0, 29, 73, 1, 0, 0, 0, 31, 86, 1, 0, 0, 0, 33, 91,
		1, 0, 0, 0, 35, 36, 5, 61, 0, 0, 36, 2, 1, 0, 0, 0, 37, 38, 5, 40, 0, 0,
		38, 4, 1, 0, 0, 0, 39, 40, 5, 41, 0, 0, 40, 6, 1, 0, 0, 0, 41, 42, 5, 42,
		0, 0, 42, 8, 1, 0, 0, 0, 43, 44, 5, 47, 0, 0, 44, 10, 1, 0, 0, 0, 45, 46,
		5, 43, 0, 0, 46, 12, 1, 0, 0, 0, 47, 48, 5, 45, 0, 0, 48, 14, 1, 0, 0,
		0, 49, 50, 5, 61, 0, 0, 50, 51, 5, 61, 0, 0, 51, 16, 1, 0, 0, 0, 52, 53,
		5, 33, 0, 0, 53, 54, 5, 61, 0, 0, 54, 18, 1, 0, 0, 0, 55, 56, 5, 60, 0,
		0, 56, 20, 1, 0, 0, 0, 57, 58, 5, 62, 0, 0, 58, 22, 1, 0, 0, 0, 59, 60,
		5, 60, 0, 0, 60, 61, 5, 61, 0, 0, 61, 24, 1, 0, 0, 0, 62, 63, 5, 62, 0,
		0, 63, 64, 5, 61, 0, 0, 64, 26, 1, 0, 0, 0, 65, 69, 7, 0, 0, 0, 66, 68,
		7, 1, 0, 0, 67, 66, 1, 0, 0, 0, 68, 71, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0,
		69, 70, 1, 0, 0, 0, 70, 28, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 72, 74, 7,
		2, 0, 0, 73, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75,
		76, 1, 0, 0, 0, 76, 83, 1, 0, 0, 0, 77, 79, 5, 46, 0, 0, 78, 80, 7, 2,
		0, 0, 79, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 81, 82,
		1, 0, 0, 0, 82, 84, 1, 0, 0, 0, 83, 77, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0,
		84, 30, 1, 0, 0, 0, 85, 87, 5, 13, 0, 0, 86, 85, 1, 0, 0, 0, 86, 87, 1,
		0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 89, 5, 10, 0, 0, 89, 32, 1, 0, 0, 0, 90,
		92, 7, 3, 0, 0, 91, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 91, 1, 0, 0,
		0, 93, 94, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 96, 6, 16, 0, 0, 96, 34,
		1, 0, 0, 0, 7, 0, 69, 75, 81, 83, 86, 93, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CalculatorLexerInit initializes any static state used to implement CalculatorLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCalculatorLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CalculatorLexerInit() {
	staticData := &CalculatorLexerLexerStaticData
	staticData.once.Do(calculatorlexerLexerInit)
}

// NewCalculatorLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCalculatorLexer(input antlr.CharStream) *CalculatorLexer {
	CalculatorLexerInit()
	l := new(CalculatorLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &CalculatorLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Calculator.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CalculatorLexer tokens.
const (
	CalculatorLexerT__0      = 1
	CalculatorLexerT__1      = 2
	CalculatorLexerT__2      = 3
	CalculatorLexerMUL       = 4
	CalculatorLexerDIV       = 5
	CalculatorLexerADD       = 6
	CalculatorLexerSUB       = 7
	CalculatorLexerEQUAL     = 8
	CalculatorLexerDIFF      = 9
	CalculatorLexerLESS      = 10
	CalculatorLexerGREATER   = 11
	CalculatorLexerLESSEQ    = 12
	CalculatorLexerGREATEREQ = 13
	CalculatorLexerID        = 14
	CalculatorLexerNUMBER    = 15
	CalculatorLexerNEWLINE   = 16
	CalculatorLexerWS        = 17
)
