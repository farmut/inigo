////////////////////////////////////////////////////
// @Author: hIMEI
// @Date:   2017-12-11 22:02:59
// @Last Modified by:   hIMEI
// @Last Modified time: 2017-12-11 22:02:59
/////////////////////////////////////////////////////
//		 o8o               o8o
//		 `"'               `"'
//		oooo  ooo. .oo.   oooo   .oooooooo  .ooooo.
//		`888  `888P"Y88b  `888  888' `88b  d88' `88b
//		 888   888   888   888  888   888  888   888
//		 888   888   888   888  `88bod8P'  888   888
//		o888o o888o o888o o888o `8oooooo.  `Y8bod8P'
//		___.ini files parser___ d"     YD
//		                        "Y88888P'
/////////////////////////////////////////////////////

package inigo

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

const (
	BITSIZE0  int = 0
	BITSIZE8      = 8
	BITSIZE16     = 16
	BITSIZE32     = 32
	BITSIZE64     = 64
)

const (
	BASE0  int = 0 // We need only this base for all case
	BASE2      = 2
	BASE8      = 8
	BASE10     = 10
	BASE16     = 16
)

const (
	ERROR   string = "; PARSING ERROR"
	NOERROR        = "NOERROR"
)

const (
	DELIM      string = ","
	ONE               = "1"
	COMM              = ";" // Ordinary comment
	UCOMM             = "#" // Unix-style comment
	DUCOMM            = "##"
	EMPTY             = ""
	SPACE             = " "
	LSQUARE           = "["
	RSQUARE           = "]"
	LFIG              = "{"
	RFIG              = "}"
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
	HEX               = "0x"
	MINUS             = "-"
	DOT               = "."
)

var COMMS = []string{
	COMM,
	UCOMM,
	DUCOMM,
}

var PREFFS = []string{
	OCTAL,
	HEX,
	MINUS,
}

var BOOLEANS = []string{
	UPPERTRUE,
	CTRUE,
	LOWERTRUE,
	UPPERFALSE,
	CFALSE,
	LOWERFALSE,
}

var LOGICS = []string{
	AND,
	OR,
	NOT,
}

var BITS = []string{
	ORBIT,
	XORBIT,
	NOTBIT,
	ANDBIT,
}

type IniParser struct {
	Functions map[int]func(string) interface{}
}

// Parsing functions from strconv package (such as ParseBool(string), ParseUint(string))
// return parsed value end error value in case of error. But if second err is omitted
// with blank identifyer, functions return 0. To handle errors and (in same time0 avoiding program
// evaluation breaks in cases of errors we may match returned value and 0, and also check if given
// string not realy "0". If 0 (integer) is returned and string != "0", there is error.
func ErrCheck(value string) string {
	var errorString string
	if value != "0" {
		errorString += value + ERROR
		//errorString +=
	} else {
		errorString = NOERROR
	}

	return errorString
}

func NewParser() *IniParser {
	parser := &IniParser{}
	parser.Functions = make(map[int]func(string) interface{})

	// Parses boolean
	parser.Functions[0] = func(value string) interface{} {
		check, _ := strconv.ParseInt(value, BASE0, BITSIZE64)

		if check == 0 && ErrCheck(value) != NOERROR {
			return ErrCheck(value)
		} else {
			return check
		}

		return
	}

	// Parses int
	parser.Functions[1] = func(value string) interface{} {
		check, _ := strconv.ParseInt(value, BASE0, BITSIZE64)

		if check == 0 && ErrCheck(value) != NOERROR {
			return ErrCheck(value)
		} else {
			return check
		}

		return
	}

	// Parses uint
	parser.Functions[2] = func(value string) interface{} {
		check, _ := strconv.ParseUint(value, BASE0, BITSIZE64)

		if check == 0 && ErrCheck(value) != NOERROR {
			return ErrCheck(value)
		} else {
			return check
		}

		return
	}

	// Parses float
	parser.Functions[3] = func(value string) interface{} {
		check, _ := strconv.ParseFloat(value, BITSIZE64)

		if check == 0 && ErrCheck(value) != NOERROR {
			return ErrCheck(value)
		} else {
			return check
		}

		return
	}

	// Parses int in binary
	parser.Functions[4] = func(value string) interface{} {
		check, _ := strconv.ParseInt(value, BASE2, BITSIZE64)

		if check == 0 && ErrCheck(value) != NOERROR {
			return ErrCheck(value)
		} else {
			return check
		}

		return
	}

	// Parses int in octal
	parser.Functions[5] = func(value string) interface{} {
		check, _ := strconv.ParseInt(value, BASE8, BITSIZE64)

		if check == 0 && ErrCheck(value) != NOERROR {
			return ErrCheck(value)
		} else {
			return check
		}

		return
	}

	// Parses int in hexademical
	parser.Functions[6] = func(value string) interface{} {
		check, _ := strconv.ParseInt(value, BASE0, BITSIZE64)

		if check == 0 && ErrCheck(value) != NOERROR {
			return ErrCheck(value)
		} else {
			return check
		}

		return
	}
	////////////////////////////////////////////////////////////////////////////////////////
	parser.Functions[7] = func(string) interface{} {
		check, _ := strconv.ParseInt(string, BASE10, BITSIZE64)

		ErrCheck(err)

		return check
	}

	parser.Functions[8] = func(string) interface{} {
		check, _ := strconv.ParseInt(string, BASE10, BITSIZE64)

		ErrCheck(err)

		return check
	}

	parser.Functions[9] = func(string) interface{} {
		check, _ := strconv.ParseInt(string, BASE10, BITSIZE64)

		ErrCheck(err)

		return check
	}

	parser.Functions[10] = func(string) interface{} {
		check, _ := strconv.ParseInt(string, BASE10, BITSIZE64)

		ErrCheck(err)

		return check
	}

	return parser
}

func (parser *IniParser) checkPreff(value string) int {
	var preff int
	for i, str := range PREFFS {
		if value == str {
			preff = i
			break
		} else {
			preff = -1
		}
	}

	return preff
}

func (parser *IniParser) checkDigs(value string) bool {
	check := true

	for _, char := range value {
		if unicode.IsDigit(rune(char)) != true && char != DOT {
			check = false
			break
		}
	}

	return check
}

func (parser *IniParser) checkDot(value string) bool {
	check := false

	if strings.Contains(value, DOT) == true && strings.Count(value, DOT) < 2 {
		check = true
	}

	return check
}

func (p *IniParser) checkBin(value string) bool {
	check := true

	for _, char := range value {
		if char != OCTAL && char != ONE {
			check = false
			break
		}
	}

	return check
}

func (parser *IniParser) checkBool(value string) bool {
	check := false

	for _, str := range BOOLEAN {
		if value == str {
			check = true
			break
		}
	}

	return check
}

func (parser *IniParser) checkChar(value string) bool {
	check := false

	for _, char := range value {
		if unicode.IsPunct(rune(char)) == true ||
			unicode.IsSymbol(rune(char)) == true {
			check = true
			break
		}
	}

	return check
}
func (parser *IniParser) checkString(value string) bool {
	check := false

	for _, char := range value {
		if unicode.IsPunct(rune(char)) == true ||
			unicode.IsSymbol(rune(char)) == true ||
			unicode.IsLetter(rune(char)) == true {
			check = true
			break
		}
	}

	return check
}

/*
	0	BOOL
	1	INT64
	2	UINT64
	3	FLOAT64
	4	BIN64
	5	OCT64
	6	HEX64
	7	ARRAY
	8	MAP
	9	LIST
	10	STR
*/
func (parser *IniParser) getParserFunc(value string) interface{} {
	var parseFunc func(value string) interface{}

	if parser.checkBool(value) == true {
		parseFunc = parser.Functions[0]
	}

	if parser.checkDot(value) == false &&
		parser.checkDigs(value) == true &&
		parser.checkPreff(value) == -1 {
		parseFunc = parser.Functions[1]
	}

	if parser.checkPreff(value) == 2 &&
		parser.checkDigs(strings.TrimPrefix(value, MINUS)) == true {
		parseFunc = parser.Functions[2]
	}

	if parser.checkDigs(value) == true &&
		parser.checkDot(value) == true {
		parseFunc = parser.Functions[3]
	}

	if parser.checkBool(value) == true {
		parseFunc = parser.Functions[4]
	}

	if parser.checkDots(value) == false &&
		parser.checkDigs(value) == true &&
		parser.checkPreff(value) == 0 {
		parseFunc = parser.Functions[5]
	}

	if parser.checkDots(value) == false &&
		parser.checkDigs(value) == true &&
		parser.checkChar(value) == false &&
		parser.checkPreff(value) == 1 {
		parseFunc = parser.Functions[6]
	}

}

//func (parser *IniParser) ReflectType(value string)
