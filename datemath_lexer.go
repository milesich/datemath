// Code generated from Datemath.g4 by ANTLR 4.7.2. DO NOT EDIT.

package datemath

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 138,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3,
	2, 5, 2, 37, 10, 2, 3, 2, 3, 2, 3, 3, 3, 3, 5, 3, 43, 10, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 54, 10, 3, 6, 3, 56, 10,
	3, 13, 3, 14, 3, 57, 3, 4, 3, 4, 7, 4, 62, 10, 4, 12, 4, 14, 4, 65, 11,
	4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 9, 7, 9, 82, 10, 9, 12, 9, 14, 9, 85, 11, 9, 5, 9, 87,
	10, 9, 3, 10, 3, 10, 3, 10, 6, 10, 92, 10, 10, 13, 10, 14, 10, 93, 3, 10,
	3, 10, 6, 10, 98, 10, 10, 13, 10, 14, 10, 99, 5, 10, 102, 10, 10, 3, 11,
	3, 11, 7, 11, 106, 10, 11, 12, 11, 14, 11, 109, 11, 11, 3, 12, 6, 12, 112,
	10, 12, 13, 12, 14, 12, 113, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14,
	3, 14, 5, 14, 123, 10, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3,
	17, 7, 17, 132, 10, 17, 12, 17, 14, 17, 135, 11, 17, 5, 17, 137, 10, 17,
	2, 2, 18, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11,
	21, 12, 23, 13, 25, 2, 27, 2, 29, 2, 31, 2, 33, 2, 3, 2, 10, 10, 2, 74,
	74, 79, 79, 102, 102, 106, 106, 111, 111, 117, 117, 121, 121, 123, 123,
	5, 2, 106, 106, 111, 111, 117, 117, 3, 2, 41, 41, 3, 2, 51, 59, 4, 2, 50,
	59, 97, 97, 6, 2, 11, 11, 13, 14, 34, 34, 162, 162, 3, 2, 97, 97, 4, 2,
	67, 92, 99, 124, 2, 151, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2,
	2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3,
	2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23,
	3, 2, 2, 2, 3, 36, 3, 2, 2, 2, 5, 55, 3, 2, 2, 2, 7, 59, 3, 2, 2, 2, 9,
	68, 3, 2, 2, 2, 11, 70, 3, 2, 2, 2, 13, 72, 3, 2, 2, 2, 15, 74, 3, 2, 2,
	2, 17, 86, 3, 2, 2, 2, 19, 101, 3, 2, 2, 2, 21, 103, 3, 2, 2, 2, 23, 111,
	3, 2, 2, 2, 25, 117, 3, 2, 2, 2, 27, 122, 3, 2, 2, 2, 29, 124, 3, 2, 2,
	2, 31, 126, 3, 2, 2, 2, 33, 136, 3, 2, 2, 2, 35, 37, 5, 17, 9, 2, 36, 35,
	3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 39, 9, 2, 2, 2,
	39, 4, 3, 2, 2, 2, 40, 43, 5, 17, 9, 2, 41, 43, 5, 19, 10, 2, 42, 40, 3,
	2, 2, 2, 42, 41, 3, 2, 2, 2, 43, 53, 3, 2, 2, 2, 44, 54, 9, 3, 2, 2, 45,
	46, 7, 111, 2, 2, 46, 54, 7, 117, 2, 2, 47, 48, 7, 119, 2, 2, 48, 54, 7,
	117, 2, 2, 49, 50, 7, 183, 2, 2, 50, 54, 7, 117, 2, 2, 51, 52, 7, 112,
	2, 2, 52, 54, 7, 117, 2, 2, 53, 44, 3, 2, 2, 2, 53, 45, 3, 2, 2, 2, 53,
	47, 3, 2, 2, 2, 53, 49, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 54, 56, 3, 2, 2,
	2, 55, 42, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 57, 58,
	3, 2, 2, 2, 58, 6, 3, 2, 2, 2, 59, 63, 7, 41, 2, 2, 60, 62, 10, 4, 2, 2,
	61, 60, 3, 2, 2, 2, 62, 65, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 63, 64, 3,
	2, 2, 2, 64, 66, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 66, 67, 7, 41, 2, 2, 67,
	8, 3, 2, 2, 2, 68, 69, 7, 45, 2, 2, 69, 10, 3, 2, 2, 2, 70, 71, 7, 47,
	2, 2, 71, 12, 3, 2, 2, 2, 72, 73, 7, 49, 2, 2, 73, 14, 3, 2, 2, 2, 74,
	75, 7, 112, 2, 2, 75, 76, 7, 113, 2, 2, 76, 77, 7, 121, 2, 2, 77, 16, 3,
	2, 2, 2, 78, 87, 7, 50, 2, 2, 79, 83, 9, 5, 2, 2, 80, 82, 9, 6, 2, 2, 81,
	80, 3, 2, 2, 2, 82, 85, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2,
	2, 84, 87, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 86, 78, 3, 2, 2, 2, 86, 79,
	3, 2, 2, 2, 87, 18, 3, 2, 2, 2, 88, 89, 5, 33, 17, 2, 89, 91, 7, 48, 2,
	2, 90, 92, 5, 31, 16, 2, 91, 90, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 91,
	3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 102, 3, 2, 2, 2, 95, 97, 7, 48, 2,
	2, 96, 98, 5, 31, 16, 2, 97, 96, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 97,
	3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 102, 3, 2, 2, 2, 101, 88, 3, 2, 2,
	2, 101, 95, 3, 2, 2, 2, 102, 20, 3, 2, 2, 2, 103, 107, 5, 25, 13, 2, 104,
	106, 5, 27, 14, 2, 105, 104, 3, 2, 2, 2, 106, 109, 3, 2, 2, 2, 107, 105,
	3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 22, 3, 2, 2, 2, 109, 107, 3, 2,
	2, 2, 110, 112, 9, 7, 2, 2, 111, 110, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2,
	113, 111, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115,
	116, 8, 12, 2, 2, 116, 24, 3, 2, 2, 2, 117, 118, 5, 29, 15, 2, 118, 26,
	3, 2, 2, 2, 119, 123, 5, 29, 15, 2, 120, 123, 5, 31, 16, 2, 121, 123, 9,
	8, 2, 2, 122, 119, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 122, 121, 3, 2, 2,
	2, 123, 28, 3, 2, 2, 2, 124, 125, 9, 9, 2, 2, 125, 30, 3, 2, 2, 2, 126,
	127, 4, 50, 59, 2, 127, 32, 3, 2, 2, 2, 128, 137, 7, 50, 2, 2, 129, 133,
	9, 5, 2, 2, 130, 132, 5, 31, 16, 2, 131, 130, 3, 2, 2, 2, 132, 135, 3,
	2, 2, 2, 133, 131, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 137, 3, 2, 2,
	2, 135, 133, 3, 2, 2, 2, 136, 128, 3, 2, 2, 2, 136, 129, 3, 2, 2, 2, 137,
	34, 3, 2, 2, 2, 18, 2, 36, 42, 53, 57, 63, 83, 86, 93, 99, 101, 107, 113,
	122, 133, 136, 3, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "", "", "'+'", "'-'", "'/'",
}

var lexerSymbolicNames = []string{
	"", "EsDuration", "GoDuration", "Literal", "Plus", "Minus", "Round", "Builtins",
	"IntegerLiteral", "FloatLiteral", "Identifier", "WhiteSpaces",
}

var lexerRuleNames = []string{
	"EsDuration", "GoDuration", "Literal", "Plus", "Minus", "Round", "Builtins",
	"IntegerLiteral", "FloatLiteral", "Identifier", "WhiteSpaces", "IdentifierStart",
	"IdentifierPart", "Letter", "Digit", "DecimalLiteral",
}

type DatemathLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewDatemathLexer(input antlr.CharStream) *DatemathLexer {

	l := new(DatemathLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Datemath.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// DatemathLexer tokens.
const (
	DatemathLexerEsDuration     = 1
	DatemathLexerGoDuration     = 2
	DatemathLexerLiteral        = 3
	DatemathLexerPlus           = 4
	DatemathLexerMinus          = 5
	DatemathLexerRound          = 6
	DatemathLexerBuiltins       = 7
	DatemathLexerIntegerLiteral = 8
	DatemathLexerFloatLiteral   = 9
	DatemathLexerIdentifier     = 10
	DatemathLexerWhiteSpaces    = 11
)
