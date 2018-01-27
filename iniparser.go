// Copyright 2018 hIMEI
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//////////////////////////////////////////////////////////////////////////

/* File iniparser.go contains types and methods for parameters values parsing */

package inigo

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

// Checker is a data type for logic separation only.
type Checker struct {
}

// Type representing value parsing logic. In default case, Inigo stores parameter's
// value as string. For recognize it true type and parse it true
// value IniParser's methods used.
type IniParser struct {

	// Checker gives the methods for pre-parsing value check.
	Checker

	// Functions is a parse functions selector.
	Functions map[int]func(string) interface{}
}

// ErrFatal is a common error handling
func ErrFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// NewParser is an IniParser's constructor
func NewParser() *IniParser {
	parser := &IniParser{}
	parser.Functions = make(map[int]func(string) interface{})

	// Parses boolean
	parser.Functions[0] = func(value string) interface{} {
		var parsed bool
		parsed, _ = strconv.ParseBool(value)
		return parsed
	}

	// Parses int64
	parser.Functions[1] = func(value string) interface{} {
		parsed, _ := strconv.ParseInt(value, BASE0, BITSIZE64)
		return parsed
	}

	// Parses uint64
	parser.Functions[2] = func(value string) interface{} {
		parsed, _ := strconv.ParseUint(value, BASE0, BITSIZE64)
		return parsed

	}

	// Parses float64
	parser.Functions[3] = func(value string) interface{} {
		parsed, _ := strconv.ParseFloat(value, BITSIZE64)
		return parsed

	}

	// Parses int64 in hexadecimal
	parser.Functions[4] = func(value string) interface{} {
		parsed, _ := strconv.ParseInt(value, BASE0, BITSIZE64)
		return parsed
	}

	// Returns default string
	parser.Functions[5] = func(value string) interface{} {
		parsed := value
		return parsed
	}

	// Parses slice of strings
	parser.Functions[6] = func(value string) interface{} {
		parsed := strings.Split(value, COMMA)

		for _, p := range parsed {
			if string(p[0]) == SPACE {
				p = strings.TrimPrefix(p, SPACE)
			}
		}

		return parsed
	}

	// Parses map[strings]strings
	parser.Functions[7] = func(value string) interface{} {
		parsed := make(map[string]string)

		bufslice := strings.Split(value, COMMA)

		for _, str := range bufslice {
			splitted := strings.Split(str, COLON)
			parsed[splitted[0]] = splitted[1]
		}

		return parsed
	}

	return parser
}

///////////////////////////////////////////////
// IniParser's methods
//////////////////////////////////////////////

// Parser main
func (p *IniParser) ParseValue(value string) interface{} {
	function := p.funcSelect(p.typeChecker(value))

	parsed := function(value)

	return parsed
}

// funcSelect chooses IniParser's function by given number. Number it is IniParser.Functions
// key returned by typeChecker().
func (p *IniParser) funcSelect(num int) func(string) interface{} {
	var function func(value string) interface{}
	if num < len(p.Functions) {
		function = p.Functions[num]
	}

	return function
}

// typeChecker makes pre-check of the value's type.
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

	case p.Comma(value) == true:
		counted = 6

	case p.Mapped(value) == true:
		counted = 7

	default:
		counted = 5
	}

	return counted
}

////////////////////////////////////////////////////
// Checker's methods
////////////////////////////////////////////////////

// Hexs checks if string contains hexadecimal number
func (c Checker) Hexs(value string) bool {
	var check bool

	if strings.HasPrefix(value, HEX) ||
		(c.Minus(value) == true && strings.Index(value, HEX) == 1) {
		check = true
		return check
	}

	return check
}

// Digs checks if string contains digits
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

// Booleans checks if string contains boolean value
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

// Dot checks if string contains dot
func (c Checker) Dot(value string) bool {
	var check bool

	if strings.Contains(value, DOT) == true &&
		strings.Count(value, DOT) < 2 {
		check = true
		return check
	}

	return check
}

// Minus checks if string contains the minus
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

// Comma checks if string contains the slice of strings
func (c Checker) Comma(value string) bool {
	var check bool

	if strings.Contains(value, COMMA) == true {
		check = true
	}

	return check
}

// Mapped checks if string contains map[string]string
func (c Checker) Mapped(value string) bool {
	var check bool

	if strings.Contains(value, COMMA) == true {
		splitted := strings.Split(value, COMMA)

		for _, s := range splitted {
			if strings.Contains(s, COLON) == true {
				check = true
			}
		}
	}

	return check
}
