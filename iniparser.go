////////////////////////////////////////////////////
// @Author: hIMEI
// @Date:   2017-12-11 22:02:59
// @Copyright © 2017 hIMEI <himei@tuta.io>
// @license MIT
/////////////////////////////////////////////////////
//
//    ██╗███╗   ██╗██╗ ██████╗  ██████╗
//    ██║████╗  ██║██║██╔════╝ ██╔═══██╗
//    ██║██╔██╗ ██║██║██║  ███╗██║   ██║
//    ██║██║╚██╗██║██║██║   ██║██║   ██║
//    ██║██║ ╚████║██║╚██████╔╝╚██████╔╝
//    ╚═╝╚═╝  ╚═══╝╚═╝ ╚═════╝  ╚═════╝
//    ___.ini files parser___
/////////////////////////////////////////////////////

package inigo

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

const ENABLED Nbld = true

// Just a Enabled
type Nbld bool

// Non-instantiated type for value-parsing logic separation.
type Checker struct {

	// Boolean const
	ENABLED Nbld
}

// Type representing value-parsing logic. In default case, Inigo stores parameter's
// value as string. For recognize it true type and parse it true
// value IniParser's methods used.
type IniParser struct {

	// Embedded type Checker gives the methods for pre-parsing value check.
	Checker

	// Parse functions selector
	Functions map[int]func(string) interface{}
}

// Parsing functions from strconv package (such as ParseBool(string), ParseUint(string))
// return parsed value end error value in case of error. But if err is omitted
// with blank identifyer, functions return 0. To handle errors and (in same time) avoiding program
// evaluation breaks in cases of errors we may match returned value with 0, and also check if given
// string not realy "0". If 0 (integer) is returned and string != "0", there is error.
func ErrCheck(value string) string {
	var errorString string
	if value != ZERO {
		errorString += value + ERROR
		//errorString +=
	} else {
		errorString = NOERROR
	}

	return errorString
}

// Common error handling
func ErrFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// IniParser's constructor
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

///////////////////////////////////////////////
// IniParser's methods
//////////////////////////////////////////////

// Parser main
func (p *IniParser) ParseValue(value string) interface{} {
	//var parsed interface{}

	function := p.funcSelect(p.typeChecker(value))

	parsed := function(value)

	return parsed
}

func (p *IniParser) funcSelect(num int) func(string) interface{} {
	var function func(value string) interface{}
	if num < len(p.Functions) {
		function = p.Functions[num]
	}

	return function
}

func (p *IniParser) typeChecker(value string) int {
	var counted int
	switch {
	case p.Booleans(value) == true:
		counted = 0

	case p.Dot(value) == false && p.Minus(value) == true && p.Digs(value) == true:
		counted = 1

	case p.Dot(value) == false && p.Minus(value) == false && p.Digs(value) == true:
		counted = 2

	case p.Dot(value) == true && p.Digs(value) == true:
		counted = 3

	case p.Hexs(value) == true:
		counted = 4

	default:
		counted = 5
	}

	return counted
}

////////////////////////////////////////////////////
// Checker's methods
////////////////////////////////////////////////////

func (c Checker) Hexs(value string) bool {
	var check bool

	if strings.HasPrefix(value, HEX) ||
		(c.Minus(value) == true && strings.Index(value, HEX) == 1) {
		check = true
		return check
	}

	return check
}

func (c Checker) Digs(value string) bool {
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

func (c Checker) Booleans(value string) bool {
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

func (c Checker) Dot(value string) bool {
	var check bool

	if strings.Contains(value, DOT) == true &&
		strings.Count(value, DOT) < 2 {
		check = true
		return check
	}

	return check
}

func (c Checker) Minus(value string) bool {
	var check bool

	if strings.Contains(value, MINUS) == true &&
		strings.Index(value, MINUS) == 0 &&
		strings.Count(value, MINUS) < 2 {
		check = true
		return check
	}

	return check
}

//                                                              |⛀ ⛂ ⛀|
// TODO: Logical and boolean expressions parse, default string  |⛀ ⛀ ⛂|
//                                                              |⛂ ⛂ ⛂|
