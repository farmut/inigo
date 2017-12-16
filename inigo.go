/////////////////////////////////////////////////////
// @Author: hIMEI
// @Date:   2017-12-06 07:15:25
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

//package main

import (
	"encoding/json"
	"fmt"
)

// String type for commented strings in INI file.
type Comments struct {
	// Slice of commented strings
	Comms []string `json:"comms"`
}

// Type for parameters containing strings in INI file. Here's different fields
// for enabled and disabled (commented) params, for comments and field for store
// strings with syntax errors.
type Params struct {

	// Slice of commented strings
	Comments Comments `json:"comments"`

	// Not commented parameter strings
	Enabled map[string]string `json:"Enabled"`

	// Commented parameter strings
	Disabled map[string]string `json:"Disabled"`

	// Field to strore strings with syntx errors. []string
	Errors []string `json:"Errors"`
}

// Type for sections. Sections is a blocks of INI file with unic name and structure
// Params as body.
type Sections struct {

	// Map [section_name]parameter_block.
	SectionsMap map[string]*Params `json:"sectionsmap"`
}

// Main data type. It is also API methods reciever. Contains Sections field.
type Inifile struct {

	// Map [section_name]parameter_block.
	Sections *Sections `json:"sections"`
}

///////////////////////////////////////////////////////////////////////////////
//
//                           **Library's API is here:**
//

//
// * Functions
// 		- NewIniFile(filename string) *Inifile

// * Methods
// 		- IniToJson() ([]byte, error)
// 		- GetSectionsNames() []string
// 		- PrintSectionsNames()
// 		- GetSectionParams(secname string) *Params
// 		- GetParamByName(params *Params, paramname string) string
// 		- GetParamsEnabled(secname string) []string
// 		- PrintParamsEnabled(secname string)
// 		- GetParamsDisabled(secname string) []string
// 		- PrintParamsDisabled(secname string)
// 		- GetAllParams() []string
// 		- PrintAllParams()
// 		- GetValue(secname, paramname string) interface{}
//
//

// Parses ini file by given file name,
// returns pointer to app main type - Inifile structure.
func NewIniFile(filename string) *Inifile {
	inifile := &Inifile{}

	newIni := newFromFile(filename)
	clearedIni := removeString(newIni, EMPTY)
	sectionsBodys := getSectionsBodys(clearedIni)
	inifile.Sections = sectionsConstruct(sectionsBodys)

	return inifile

}

// Json-marshalizer
func (i *Inifile) IniToJson() ([]byte, error) {
	nosj, _ := json.Marshal(i)

	return nosj, nil
}

// Gets Inifile's sections names. Returns []string.
func (i *Inifile) GetSectionsNames() []string {
	var secnames []string

	for key := range i.Sections.SectionsMap {
		secnames = append(secnames, key)
	}

	return secnames
}

// Prints section names.
func (i *Inifile) PrintSectionsNames() {
	secnames := i.GetSectionsNames()

	for _, str := range secnames {
		fmt.Println(str)
	}
}

// Gets value by section name. Returns *Params
func (i *Inifile) GetSectionParams(secname string) *Params {
	return i.Sections.SectionsMap[secname]
}

// Gets stored string representation of enabled parameter's value by given parameter name.
func (i *Inifile) GetParamByName(secname, paramname string) string {
	params := i.GetSectionParams(secname)
	return params.Enabled[paramname]
}

// Gets enabled parameters for given section.
func (i *Inifile) GetParamsEnabled(secname string) []string {
	var enabled []string

	for _, str := range i.Sections.SectionsMap[secname].Enabled {
		enabled = append(enabled, str)
	}

	return enabled
}

// Prints enabled parameters for given section.
func (i *Inifile) PrintParamsEnabled(secname string) {
	paramsenabled := i.GetParamsEnabled(secname)

	for _, str := range paramsenabled {
		fmt.Println(str)
	}
}

// Gets disabled parameters for given section.
func (i *Inifile) GetParamsDisabled(secname string) []string {
	var disabled []string

	for _, str := range i.Sections.SectionsMap[secname].Disabled {
		disabled = append(disabled, str)
	}

	return disabled
}

// Prints disabled parameters for given section.
func (i *Inifile) PrintParamsDisabled(secname string) {
	paramsdisabled := i.GetParamsEnabled(secname)

	for _, str := range paramsdisabled {
		fmt.Println(str)
	}
}

// Gets parameters of all sections.
func (i *Inifile) GetAllParams() []string {
	var params []string
	secnames := i.GetSectionsNames()

	for _, val := range secnames {
		for ename := range i.Sections.SectionsMap[val].Enabled {
			params = append(params, ename)
		}

		for dname := range i.Sections.SectionsMap[val].Disabled {
			params = append(params, dname)
		}
	}

	return params
}

// Parses parameter value. Method calls IniParser's constructor end then calls parser's Parse() method.
// Trys to get underlying value of string value. In case of error return string with error message.
func (i *Inifile) GetValue(secname, paramname string) interface{} {
	parser := NewParser()
	//params := i.GetSectionParams(secname)
	value := i.GetParamByName(secname, paramname)

	parsed := parser.ParseValue(value)

	return parsed
}

// Prints parameters of all sections.
func (i *Inifile) PrintAllParams() {
	params := i.GetAllParams()

	for _, str := range params {
		fmt.Println(str)
	}
}
