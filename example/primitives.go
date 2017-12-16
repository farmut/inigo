package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Constants for numbers parsing options
const (
	BITSIZE0  int = 0x0
	BITSIZE8      = 0x8
	BITSIZE16     = 0x10
	BITSIZE32     = 0x20
	BITSIZE64     = 0x40 // We use only this bitsize for all case.
)

// Constants for numbers parsing options
const (
	BASE0  int = 0x0
	BASE2      = 0x2
	BASE8      = 0x8
	BASE10     = 0xA
	BASE16     = 0x10
)

const (
	COMMA      string = ","
	ONE               = "1"
	COLON             = ":"
	COMM              = ";" // Ordinary comment
	UCOMM             = "#" // Unix-style comment
	DUCOMM            = "##"
	EMPTY             = ""
	SPACE             = " "
	CSPACE            = ", "
	LSQUARE           = "["
	RSQUARE           = "]"
	LFIG              = "{"
	RFIG              = "}"
	QUOTE             = "'"
	DQUOTE            = '"'
	EQUAL             = "="
	UPPERTRUE         = "TRUE"
	CTRUE             = "True"
	LOWERTRUE         = "true"
	UPPERFALSE        = "FALSE"
	CFALSE            = "False"
	LOWERFALSE        = "false"
	GLOBAL            = "GLOBAL"
	NONE              = "NONE"
	AND               = "&&"
	ANDBIT            = "&" // Bitwise AND
	OR                = "||"
	ORBIT             = "|" // Bitwise OR
	XORBIT            = "^" // Bitwise XOR
	NOT               = "!"
	NOTBIT            = "~" // Bitwise NOT
	OCTAL             = "0"
	ZERO              = "0"
	HEX               = "0x"
	MINUS             = "-"
	DOT               = "."
)

var BOOLEANS = []string{
	UPPERTRUE,
	CTRUE,
	LOWERTRUE,
	UPPERFALSE,
	CFALSE,
	LOWERFALSE,
}

//type ParceFlag int
/*
const (
	0BOOL ParceFlag = 1 + iota
	1INT64
	2UINT64
	3FLOAT64
	4INT64HEX
	5STRING
)
*/

type IniParser struct {
	// Parse functions selector
	Functions map[int]func(string) interface{}
}

func NewParser() *IniParser {
	parser := &IniParser{}
	parser.Functions = make(map[int]func(string) interface{})

	// Parses boolean
	parser.Functions[0] = func(value string) interface{} {
		var parsed bool // E.g. error
		parsed, _ = strconv.ParseBool(value)
		return parsed
	}

	// Parses int
	parser.Functions[1] = func(value string) interface{} {
		parsed, _ := strconv.ParseInt(value, BASE0, BITSIZE64)
		return parsed
	}

	// Parses uint
	parser.Functions[2] = func(value string) interface{} {
		parsed, _ := strconv.ParseUint(value, BASE0, BITSIZE64)
		return parsed

	}

	// Parses float
	parser.Functions[3] = func(value string) interface{} {
		parsed, _ := strconv.ParseFloat(value, BITSIZE64)
		return parsed

	}

	// Parses int in hexademical
	parser.Functions[4] = func(value string) interface{} {
		parsed, _ := strconv.ParseInt(value, BASE0, BITSIZE64)
		return parsed
	}

	// Returns default string
	parser.Functions[5] = func(value string) interface{} {
		parsed := value
		return parsed
	}
	/*
		// Parses slice of strings
		parser.Functions[7] = func(value string) interface{} {
			parsed := strings.Split(value, COMMA)

			return parsed
		}

		// Parses map[strings]strings
		parser.Functions[8] = func(value string) interface{} {
			parsed := make(map[string]string)

			bufslice := strings.Split(value, COMMA)

			for _, str := range bufslice {
				splitted := strings.Split(str, COLON)
				check[splitted[0]] = splitted[1]
			}

			return parsed
		}
	*/

	return parser
}

func (p *IniParser) funcSelect(num int) func(string) interface{} {
	var function func(value string) interface{}
	if num < len(p.Functions) {
		function = p.Functions[num]
	}

	return function
}

func (p *IniParser) ParseValue(value string) interface{} {
	//var parsed interface{}

	function := p.funcSelect(TypeChecker(value))

	parsed := function(value)

	return parsed
}

func Hexs(value string) bool {
	var check bool

	if strings.HasPrefix(value, HEX) ||
		(Minus(value) == true && strings.Index(value, HEX) == 1) {
		check = true
		return check
	}

	return check
}

func Digs(value string) bool {
	for i, char := range value {
		if i != 0 || string(char) != MINUS {
			if unicode.IsDigit(char) == false && string(char) != DOT {
				return unicode.IsDigit(char)
				break
			}
		}
	}

	return true
}

func Booleans(value string) bool {
	var check bool

	for _, boo := range BOOLEANS {
		if value == boo {
			check = true
			return check
			break
		}
	}

	return check
}

func Dot(value string) bool {
	var check bool

	if strings.Contains(value, DOT) == true &&
		strings.Count(value, DOT) < 2 {
		check = true
		return check
	}

	return check
}

func Minus(value string) bool {
	var check bool

	if strings.Contains(value, MINUS) == true &&
		strings.Index(value, MINUS) == 0 &&
		strings.Count(value, MINUS) < 2 {
		check = true
		return check
	}

	return check
}

func TypeChecker(value string) int {
	var counted int
	switch {
	case Booleans(value) == true:
		counted = 0

	case Dot(value) == false && Minus(value) == true && Digs(value) == true:
		counted = 1

	case Dot(value) == false && Minus(value) == false && Digs(value) == true:
		counted = 2

	case Dot(value) == true && Digs(value) == true:
		counted = 3

	case Hexs(value) == true:
		counted = 4

	default:
		counted = 5
	}

	return counted
}

func main() {
	parser := NewParser()

	value := "True"

	res := parser.ParseValue(value)

	fmt.Println(res)

	//fmt.Println(TypeChecker(value))
	//fmt.Println(Hexs(value))
}
