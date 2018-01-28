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

/* File globs.go contains constants and global variables */

package inigo

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

// Constants for parsing errors
const (
	ERROR   string = "; PARSING ERROR"
	NOERROR        = "NOERROR"
)

// Useful string constants
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

// Constants for comments parsing
var COMMS = []string{
	COMM,
	UCOMM,
	DUCOMM,
}

// Constants for octal, hexademical and negative values parsing
var PREFFS = []string{
	OCTAL,
	HEX,
	MINUS,
}

// Constants for boolean values parsing
var BOOLEANS = []string{
	UPPERTRUE,
	CTRUE,
	LOWERTRUE,
	UPPERFALSE,
	CFALSE,
	LOWERFALSE,
}

// Constants for logical expressions parsing
var LOGICS = []string{
	AND,
	OR,
	NOT,
}
