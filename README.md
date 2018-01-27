# inigo

**Inigo** - INI files parsing library for Golang applications with wide range of data types parsed. 

    ██╗███╗   ██╗██╗ ██████╗  ██████╗ 
    ██║████╗  ██║██║██╔════╝ ██╔═══██╗
    ██║██╔██╗ ██║██║██║  ███╗██║   ██║
    ██║██║╚██╗██║██║██║   ██║██║   ██║
    ██║██║ ╚████║██║╚██████╔╝╚██████╔╝
    ╚═╝╚═╝  ╚═══╝╚═╝ ╚═════╝  ╚═════╝ 
        ___.ini files parser___

[![Go Report Card](https://goreportcard.com/badge/github.com/hIMEI29A/inigo)](https://goreportcard.com/report/github.com/hIMEI29A/inigo) [![GoDoc](https://godoc.org/github.com/hIMEI29A/inigo?status.svg)](http://godoc.org/github.com/hIMEI29A/inigo)



## TOC
- [Features](#features)
- [License](#license)
- [Version](#version)
- [Install](#install)
- [Usage](#usage)
- [Example](#example)
- [Dependencies](#dependencies)
- [Contributing](#contributing)

1. Features
2. License
3. Version
4. Install
5. Usage
6. Example
7. Dependencies
8. Contributing

## Features

* **Inigo** parses disabled (_commented_) parameters of INI file as well as enabled ones
* Unlike some others INI parsers, **Inigo** recognizes parameter's value not just as string, but tries to parse its native data type. Here is the types that **Inigo** understands: 

    - Int64
    - Uint64
    - Float64
    - Boolean
    - Octal representation of Int
    - Hexadecimal representation of Int
    - Binary representation of Int
    - String (default)
    - Array (as **Go**'s slice of strings)
    - Map (currently _[string]string_ only)
 
 * **Inigo** also parses ini file's parameters with syntax errors and stores its as special data type
 * **Inifo** it is INI parser, **not editor**. For INI files editing use your preferred text editor (Vim is great!)
 
For details of used INI syntax see package's documentation.

## License

[![Apache-2.0 License](http://img.shields.io/badge/License-Apache-2.0-yellow.svg)](LICENSE)

## Version

v0.1.0

## Install

Check the [release page](https://github.com/hIMEI29A/inigo/releases)!

**Inigo** project uses standart installation way with `go get`

To install **Inigo** to your $GOPATH, enter:

```sh
go get github.com/hIMEI29A/inigo
```

## Usage

To use **Inigo** in your application, import it with next snippet:

```go
import (
    inigo "github.com/hIMEI29A/inigo"
)
```

After it, you may call library's API methods such as

```go
func main() {
    filename := "example.ini"

    inifile := inigo.NewIniFile(filename)

    params := inifile.GetAllParams()

    inifile.PrintAllParams()

    value := inifile.GetValue(section1, parameter1)

    fmt.Println(value)
}
```

## Example

See package's `example` folder. There you find sample Go program which demonstrates **Inigo**'s features.

Here is the short example of INI files parsed by **Inigo**:

    ;;;;;;;;;;;;;;;;;;;;;
    ;; new.ini file begin
    ;;;;;;;;;;;;;;;;;;;;;

    ; empty strings will be ignored

    ; global section without name

    parameter="default string"
    parameterWithError = spaces found here
    ; disParam="disabled param"

    ; section1
    [SECTION1]

    ; bool values
    param0=true ; inline comments will be ignored
    param101 = TRUE
    # integer
    param1=-4 # UNIX style comments. Inline will be ignored too
    ; uint64
    param2=53436790665
    ; Float64
    param3=657.0076

    ; binary 
    ; it will be parsed as 25
    param4=011001

    ; disParam2="disabled param"

    ; octal
    ; it will be parsed as 100
    param5=0144

    ; hex
    ; it will be parsed as 2323
    param6=0x913

    ; slice of strings
    param7=nazguls, hobbits, strange_creatures
    ; map [string]string
    param8=first:string, second:443, third:0

## Dependencies

**Inigo** hasn't any external deps!

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

                           |⛀ ⛂ ⛀|
                           |⛀ ⛀ ⛂|
                           |⛂ ⛂ ⛂|
