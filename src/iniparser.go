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
	"unicode"
	"strings"
)

const (
	BITSIZE0  int = 0
	BITSIZE8      = 8
	BITSIZE16     = 16
	BITSIZE32     = 32
	BITSIZE64     = 64
)

const (
	BASE2  int = 2
	BASE8      = 8
	BASE10     = 10
	BASE16     = 16
)

const (
	BOOL ParseFlag = 1 + iota
	INT64
	UINT64
	FLOAT64
	BIN64
	OCT64
	HEX64
	ARRAY
	MAP
	LIST
	STR
)

var Flags = [...]string{
	"BOOL",
	"INT64",
	"UINT64",
	"FLOAT64",
	"BIN64",
	"OCT64",
	"HEX64",
	"ARRAY",
	"MAP",
	"LIST",
	"STR",
}

type ParseFlag int

type IniParser struct {
	ParseFlag
	Functions map[int]func(string) interface{}
	Flags     [...]string
}

func (flag ParseFlag) String() string {
	return Flags[flag-1]
}

func ErrCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func NewParser() *IniParser {
	parser := &IniParser{}
	parser.Functions = make(map[int]func(string) interface{})

	// Parses boolean
	parser.Functions[0] = func(string) interface{} {
		check, err := strings.ParseBool(string)

		ErrCheck(err)

		return check
	}

	// Parses int
	parser.Functions[1] = func(string) interface{} {
		check, err := strings.ParseInt(string, BASE10, BITSIZE64)

		ErrCheck(err)

		return check
	}

	// Parses uint
	parser.Functions[2] = func(string) interface{} {
		check, err := strings.ParseUint(string, BASE10, BITSIZE64)

		ErrCheck(err)

		return check
	}

	// Parses float
	parser.Functions[3] = func(string) interface{} {
		check, err := strings.ParseFloat(string, BITSIZE64)

		ErrCheck(err)

		return check
	}

	// Parses int in binary
	parser.Functions[4] = func(string) interface{} {
		check, err := strings.ParseInt(string, BASE2, BITSIZE64)

		ErrCheck(err)

		return check
	}

	// Parses int in octal
	parser.Functions[5] = func(string) interface{} {
		check, err := strings.ParseInt(string, BASE8, BITSIZE64)

		ErrCheck(err)

		return check
	}

	// Parses int in hexademical
	parser.Functions[6] = func(string) interface{} {
		check, err := strings.ParseInt(string, BASE16, BITSIZE64)

		ErrCheck(err)

		return check
	}

	parser.Functions[7] = func(string) interface{} {
		check, err := ParseInt(string, BASE10, BITSIZE64)

		ErrCheck(err)

		return check
	}

	parser.Functions[8] = func(string) interface{} {
		check, err := ParseInt(string, BASE10, BITSIZE64)

		ErrCheck(err)

		return check
	}

	parser.Functions[9] = func(string) interface{} {
		check, err := ParseInt(string, BASE10, BITSIZE64)

		ErrCheck(err)

		return check
	}

	parser.Functions[10] = func(string) interface{} {
		check, err := ParseInt(string, BASE10, BITSIZE64)

		ErrCheck(err)

		return check
	}
}

func (parser *IniParser) checkDigs(value string) bool {
	var check bool

	for _, char := range value {
		if unicode.IsDigit(rune(char)) != true || (value[0] != "-" || char != ".") {
			check = false
		} else {
			check = true
		}
	}

	return check
}

func (parser *IniParser) checkChars(value string) bool {
	var check bool

	for _, char := range value {
		if (unicode.IsPunct(rune(char)) == true && (value[0] != "-" || char != ".")) ||
			unicode.IsSymbol(rune(char)) == true {
			check = true
		} else {
			check = false
		}
	}

	return check
}

func (parser *IniParser) parseValue(value string) interface{} {
	if checkDigs(value) == true {
		if strings.Contains(value, ".") == true {
			parsed := strconv.ParseFloat(value)
		}
}

//func (parser *IniParser) ReflectType(value string)
