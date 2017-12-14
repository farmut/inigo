# inigo
--
    import "bitbucket.org/hIMEI/inigo"


## Usage

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


    var BITS = []string{
    	ORBIT,
    	XORBIT,
    	NOTBIT,
    	ANDBIT,
    }


    var BOOLEANS = []string{
    	UPPERTRUE,
    	CTRUE,
    	LOWERTRUE,
    	UPPERFALSE,
    	CFALSE,
    	LOWERFALSE,
    }


    var COMMS = []string{
    	COMM,
    	UCOMM,
    	DUCOMM,
    }


    var LOGICS = []string{
    	AND,
    	OR,
    	NOT,
    }


    var PREFFS = []string{
    	OCTAL,
    	HEX,
    	MINUS,
    }


#### func  ErrCheck

    func ErrCheck(value string) string

Parsing functions from strconv package (such as ParseBool(string),
ParseUint(string)) return parsed value end error value in case of error. But if
err is omitted with blank identifyer, functions return 0. To handle errors and
(in same time) avoiding program evaluation breaks in cases of errors we may
match returned value with 0, and also check if given string not realy "0". If 0
(integer) is returned and string != "0", there is error.

#### func  ErrFatal

    func ErrFatal(err error)

Common error handling

#### type Checker

    type Checker struct {
    	ENABLED Nbld
    }


Non-instantiated type for logic separation only

#### type Comments

    type Comments struct {
    	Comms []string `json:"comms"`
    }



#### type IniParser

    type IniParser struct {
    	Checker
    	Functions map[int]func(string) interface{}
    }


Type representing value-parsing logic. In default case, Inigo stores parameter's
value as string. For recognize it true type and parse it true value IniParser's
methods used.

#### func  NewParser

    func NewParser() *IniParser

IniParser's constructor

#### func (*IniParser) Parse

    func (parser *IniParser) Parse(value string) interface{}

Parser main

#### type Inifile

    type Inifile struct {
    	Sections *Sections `json:"sections"`
    }



#### func  NewIniFile

    func NewIniFile(filename string) *Inifile

Parses ini file by given file name, returns pointer to app main type - Inifile
structure.

#### func (*Inifile) GetAllParams

    func (i *Inifile) GetAllParams() []string

Gets parameters of all sections.

#### func (*Inifile) GetParamByName

    func (i *Inifile) GetParamByName(params *Params, paramname string) string

Gets stored string representation of enabled parameter's value by given
parameter name.

#### func (*Inifile) GetParamsDisabled

    func (i *Inifile) GetParamsDisabled(secname string) []string

Gets disabled parameters for given section.

#### func (*Inifile) GetParamsEnabled

    func (i *Inifile) GetParamsEnabled(secname string) []string

Gets enabled parameters for given section.

#### func (*Inifile) GetSectionByName

    func (i *Inifile) GetSectionByName(secname string) *Params

Gets value by section name. Returns *Params

#### func (*Inifile) GetSectionsNames

    func (i *Inifile) GetSectionsNames() []string

Gets Inifile's sections names. Returns []string.

#### func (*Inifile) GetValue

    func (i *Inifile) GetValue(secname, paramname string) interface{}

Parses parameter value. Method calls IniParser's constructor end then calls
parser's Parse() method. Trys to get underlying value of string value. In case
of error return string with error message.

#### func (*Inifile) IniToJson

    func (i *Inifile) IniToJson() ([]byte, error)


#### func (*Inifile) PrintAllParams

    func (i *Inifile) PrintAllParams()

Prints parameters of all sections.

#### func (*Inifile) PrintParamsDisabled

    func (i *Inifile) PrintParamsDisabled(secname string)

Prints disabled parameters for given section.

#### func (*Inifile) PrintParamsEnabled

    func (i *Inifile) PrintParamsEnabled(secname string)

Prints enabled parameters for given section.

#### func (*Inifile) PrintSectionsNames

    func (i *Inifile) PrintSectionsNames()

Prints section names.

#### type Nbld

    type Nbld bool


Just a Enabled

    const ENABLED Nbld = true


#### type Params

    type Params struct {
    	Comments Comments          `json:"comments"`
    	Enabled  map[string]string `json:"Enabled"`
    	Disabled map[string]string `json:"Disabled"`
    	Errors   []string          `json:"Errors"`
    }



#### type Sections

    type Sections struct {
    	SectionsMap map[string]*Params `json:"sectionsmap"`
    }
