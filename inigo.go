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

// Type for sections. Sections is a blocks of INI file with unic name and Params structure
//  as body.
type Sections struct {

	// Map [section_name]parameter_block.
	SectionsMap map[string]*Params `json:"sectionsmap"`
}

// Inifile is an main data type. It is also API methods reciever. Contains Sections field.
type Inifile struct {

	// Sections
	Sections *Sections `json:"sections"`
}

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

// GetSectionsNames gets Inifile's sections names. Returns []string.
func (i *Inifile) GetSectionsNames() []string {
	var secnames []string

	for key := range i.Sections.SectionsMap {
		secnames = append(secnames, key)
		fmt.Println(key)
	}

	return secnames
}

// Gets value by section name. Returns *Params
func (i *Inifile) getSectionParams(secname string) *Params {
	return i.Sections.SectionsMap[secname]
}

// Gets string representation of enabled parameter's value by given parameter name.
func (i *Inifile) GetEnableByName(secname, paramname string) string {
	params := i.getSectionParams(secname)
	return params.Enabled[paramname]
}

// Gets string representation of disabled parameter's value by given parameter name.
func (i *Inifile) GetDisableByName(secname, paramname string) string {
	params := i.getSectionParams(secname)
	return params.Disabled[paramname]
}

// Gets enabled parameters for given section.
func (i *Inifile) GetParamsEnabled(secname string) []string {
	var enabled []string

	for _, str := range i.Sections.SectionsMap[secname].Enabled {
		enabled = append(enabled, str)
	}

	return enabled
}

// Gets disabled parameters for given section.
func (i *Inifile) GetParamsDisabled(secname string) []string {
	var disabled []string

	for _, str := range i.Sections.SectionsMap[secname].Disabled {
		disabled = append(disabled, str)
	}

	return disabled
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

func (i *Inifile) GetAllErrors() []string {
	var errors []string
	secnames := i.GetSectionsNames()

	for _, val := range secnames {
		params := i.getSectionParams(val)

		for _, e := range params.Errors {
			errors = append(errors, e)
		}
	}

	return errors
}

// Parses parameter value. Method calls IniParser's constructor and then calls parser's Parse() method.
// Trys to get underlying value of string value. In case of error return string with error message.
func (i *Inifile) GetValue(secname, paramname string) interface{} {
	parser := NewParser()
	//params := i.GetSectionParams(secname)
	value := i.GetEnableByName(secname, paramname)

	parsed := parser.ParseValue(value)

	return parsed
}
