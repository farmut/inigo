/////////////////////////////////////////////////////
// @Author: hIMEI
// @Date:   2017-12-13 00:42:43
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

const (
	BITSIZE0  int = 0x0
	BITSIZE8      = 0x8
	BITSIZE16     = 0x10
	BITSIZE32     = 0x20
	BITSIZE64     = 0x40 // We use only this bitsize for all case.
)

const (
	BASE0  int = 0x0
	BASE2      = 0x2
	BASE8      = 0x8
	BASE10     = 0xA
	BASE16     = 0x10
)

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
