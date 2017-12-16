package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

// Constants for parsing errors handling
const (
	ERROR   string = "; PARSING ERROR"
	NOERROR        = "NOERROR"
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

const ENABLED Nbld = true

// Just a Enabled
type Nbld bool

type Counter struct {
	ENABLED Nbld
}

func count() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

var COUNTER = count()

type IniParser struct {
	// Parse functions selector
	Functions map[int]func(string) interface{}
}

func (p *IniParser) funcSelect(num int) func(string) interface{} {
	var function func(string) interface{}

	if num < len(p.Functions)-1 {
		function = p.Functions[num]
	}

	return function
}

func prval(val int) {
	fmt.Println(val)
}

func (p *IniParser) MarseValue() {
	for i := 0; i < 10; i++ {
		prval(COUNTER())
	}
}

func (p *IniParser) ParseValue(value string) interface{} {
	var parsed interface{}

	for i := 0; i < len(p.Functions); i++ {
		parseFunc := p.funcSelect(COUNTER())
		parsed = parseFunc(value)

		if p.CheckError(value, parsed) != false {
			break
		}
	}

	return parsed
}

func (p *IniParser) CheckError(value string, parsedValue interface{}) bool {
	//fmt.Println(value)
	fmt.Println(parsedValue)
	var check bool
	if parsedValue != 0 {
		//fmt.Println("pi")
		if value != ZERO {
			check = true
			//fmt.Println(check)
			return check
		}
	}

	return check
}

func NewParser() *IniParser {
	parser := &IniParser{}
	parser.Functions = make(map[int]func(string) interface{})

	// Parses boolean
	parser.Functions[0] = func(value string) interface{} {
		var check bool // E.g. error
		check, _ = strconv.ParseBool(value)

		return check
	}

	// Parses int
	parser.Functions[1] = func(value string) interface{} {
		check, _ := strconv.ParseInt(value, BASE0, BITSIZE64)
		return check
	}

	// Parses uint
	parser.Functions[2] = func(value string) interface{} {
		check, _ := strconv.ParseUint(value, BASE0, BITSIZE64)
		return check

	}

	// Parses float
	parser.Functions[3] = func(value string) interface{} {
		check, _ := strconv.ParseFloat(value, BITSIZE64)
		return check

	}

	// Parses int in binary
	parser.Functions[4] = func(value string) interface{} {
		check, _ := strconv.ParseInt(value, BASE2, BITSIZE64)
		return check
	}

	// Parses int in octal
	parser.Functions[5] = func(value string) interface{} {
		check, _ := strconv.ParseInt(value, BASE8, BITSIZE64)

		return check
	}

	// Parses int in hexademical
	parser.Functions[6] = func(value string) interface{} {
		check, _ := strconv.ParseInt(value, BASE0, BITSIZE64)

		return check
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

	return parser
}

func main() {
	parser := NewParser()

	value := "23.56"

	parsed := parser.ParseValue(value)

	fmt.Println(parsed)
}
