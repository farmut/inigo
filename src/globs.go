//
// @Author: hIMEI
// @Date:   2017-12-13 00:42:43
// @Last Modified by:   hIMEI
// @Last Modified time: 2017-12-13 00:42:43
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
	COMMA      string = ","
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
