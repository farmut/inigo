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

// Non-instantiated type fo Tick() method
type Ticker struct {

	// Boolean const
	ENABLED Nbld
}

func (t Ticker) Tick() int {

}

// Type representing value-parsing logic. In default case, Inigo stores parameter's
// value as string. For recognize it true type and parse it true
// value IniParser's methods used.
type IniParser struct {

	// Embedded type Checker gives the methods for pre-parsing value check.
	Checker

	// Embedd type Ticker to provide Tick() method
	Ticker

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
		var check string // E.g. error
		for _, str := range BOOLEANS {
			if value == str {
				check = ZERO // No error
				break
			} else {
				check = ONE
			}
		}

		newcheck = ErrCheck(check)

		if check != NOERROR {
			return check
		} else {
			return value
		}
	}

	// Parses int
	parser.Functions[1] = func(value string) interface{} {
		check, _ := strconv.ParseInt(value, BASE0, BITSIZE64)

		if check == 0 && ErrCheck(value) != NOERROR {
			return ErrCheck(value)
		} else {
			return check
		}
	}

	// Parses uint
	parser.Functions[2] = func(value string) interface{} {
		check, _ := strconv.ParseUint(value, BASE0, BITSIZE64)

		if check == 0 && ErrCheck(value) != NOERROR {
			return ErrCheck(value)
		} else {
			return check
		}
	}

	// Parses float
	parser.Functions[3] = func(value string) interface{} {
		check, _ := strconv.ParseFloat(value, BITSIZE64)

		if check == 0 && ErrCheck(value) != NOERROR {
			return ErrCheck(value)
		} else {
			return check
		}
	}

	// Parses int in binary
	parser.Functions[4] = func(value string) interface{} {
		check, _ := strconv.ParseInt(value, BASE2, BITSIZE64)

		if check == 0 && ErrCheck(value) != NOERROR {
			return ErrCheck(value)
		} else {
			return check
		}
	}

	// Parses int in octal
	parser.Functions[5] = func(value string) interface{} {
		check, _ := strconv.ParseInt(value, BASE8, BITSIZE64)

		if check == 0 && ErrCheck(value) != NOERROR {
			return ErrCheck(value)
		} else {
			return check
		}
	}

	// Parses int in hexademical
	parser.Functions[6] = func(value string) interface{} {
		check, _ := strconv.ParseInt(value, BASE0, BITSIZE64)

		if check == 0 && ErrCheck(value) != NOERROR {
			return ErrCheck(value)
		} else {
			return check
		}
	}

	// Parses slice of strings
	parser.Functions[7] = func(value string) interface{} {
		check := strings.Split(value, COMMA)

		return check
	}

	// Parses map[strings]strings
	parser.Functions[8] = func(value string) interface{} {
		check := make(map[string]string)

		bufslice := strings.Split(value, COMMA)

		for _, str := range bufslice {
			splitted := strings.Split(str, COLON)
			check[splitted[0]] = splitted[1]
		}

		return check
	}
	/* {{{Template}}}

	parser.Functions[10] = func(string) interface{} {
		check, _ := strconv.ParseInt(string, BASE10, BITSIZE64)

		ErrCheck(err)

		return check
	}
	*/
	return parser
}

///////////////////////////////////////////////
// IniParser's methods
//////////////////////////////////////////////

// Primary checker and functions selector.
func (parser *IniParser) getParserFunc(value string) func(value string) interface{} {
	var parseFunc func(value string) interface{}

	// Bool
	if parser.checkBool(value) == true {
		parseFunc = parser.Functions[0]
	}

	// Int64
	if parser.checkPreff(value) == 2 &&
		parser.checkDigs(strings.TrimPrefix(value, MINUS)) == true {
		parseFunc = parser.Functions[1]
	}

	// Uint64
	if parser.checkDot(value) == false &&
		parser.checkDigs(value) == true &&
		parser.checkPreff(value) == -1 {
		parseFunc = parser.Functions[2]
	}

	// Float64
	if parser.checkDigs(value) == true &&
		parser.checkDot(value) == true {
		parseFunc = parser.Functions[3]
	}

	// Int64 in bin representation
	if parser.checkBool(value) == true {
		parseFunc = parser.Functions[4]
	}

	// Int64 in oct representation
	if parser.checkDot(value) == false &&
		parser.checkDigs(value) == true &&
		parser.checkPreff(value) == 0 &&
		parser.checkBin(value) == false {
		parseFunc = parser.Functions[5]
	}

	// Int64 in hex representation
	if parser.checkDot(value) == false &&
		parser.checkDigs(value) == true &&
		parser.checkChar(value) == false &&
		parser.checkPreff(value) == 1 {
		parseFunc = parser.Functions[6]
	}

	// Slice of string
	if strings.Contains(value, COMMA) == true {
		parseFunc = parser.Functions[7]
	}

	// map[strings]strings
	if parser.checkColonsComma(value) == true {
		parseFunc = parser.Functions[8]
	}

	return parseFunc
}

// Parser main
func (parser *IniParser) Parse(value string) interface{} {
	parseFunc := parser.getParserFunc(value)

	res := parseFunc(value)
	return res
}

////////////////////////////////////////////////////
// Checker's methods
////////////////////////////////////////////////////

// Hex, octal or "-" check
func (checker Checker) checkPreff(value string) int {
	var preff int
	for i, str := range PREFFS {
		if strings.HasPrefix(value, str) == true {
			preff = i
			break
		} else {
			preff = -1
		}
	}

	return preff
}

// Checks for digits and dot.
func (checker Checker) checkDigs(value string) bool {
	check := true

	for _, char := range value {
		if unicode.IsDigit(rune(char)) != true && string(char) != DOT {
			check = false
			break
		}
	}

	return check
}

// Checks for dot. Only one dot in the string allowed.
func (checker Checker) checkDot(value string) bool {
	check := false

	if strings.Contains(value, DOT) == true && strings.Count(value, DOT) < 2 {
		check = true
	}

	return check
}

// Checks for binary representation of int.`
func (checker Checker) checkBin(value string) bool {
	check := true

	for _, char := range value {
		if string(char) != OCTAL && string(char) != ONE {
			check = false
			break
		}
	}

	return check
}

// Checks for bool.
func (checker Checker) checkBool(value string) bool {
	check := false

	for _, str := range BOOLEANS {
		if value == str {
			check = true
			break
		}
	}

	return check
}

// Checks for spesial characters.
func (checker Checker) checkChar(value string) bool {
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

// Checks for quotes.
func (checker Checker) checkQuotes(value string) bool {
	check := false

	if value[0] == DQUOTE && value[len(value)-1] == DQUOTE {
		check = true
	}

	return check
}

// Checks for colons and commas for map.
func (checker Checker) checkColonsComma(value string) bool {
	check := false

	if strings.Contains(value, COLON) == true &&
		strings.Contains(value, COMMA) == true &&
		strings.Count(value, COLON) == strings.Count(value, COMMA) {
		check = true
	}

	return check
}

// Common check for string.
func (checker Checker) checkString(value string) bool {
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

//                                                            |⛀ ⛂ ⛀|
// TODO: Logic and boolean expressions parse, default string  |⛀ ⛀ ⛂|
//                                                            |⛂ ⛂ ⛂|
