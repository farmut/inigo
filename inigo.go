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

/* File inigo.go contains package's API methods */

package inigo

import (
//"fmt"
)

// Params is a parameters containing strings in INI file. Here's different fields
// for enabled and disabled (commented) params and field to store
// strings with syntax errors.
type Params struct {

	// Enabled is a not commented parameters
	Enabled map[string]string

	// Disabled is a commented parameters
	Disabled map[string]string

	// Errors is a field to strore strings with syntx errors. []string
	Errors []string
}

// Sections structure for sections. Sections is a blocks of INI file with unic name and
// Params structure as body.
type Sections struct {

	// SectionsMap is a Map [section_name]parameter_block.
	SectionsMap map[string]*Params
}

// Inifile is a main data type. It is also API methods reciever. Contains Sections field.
type Inifile struct {

	// Sections is a Map [section_name]parameter_block.
	Sections *Sections
}

// NewIniFile parses ini file by given file name,
// returns pointer to app main type - Inifile structure.
func NewIniFile(filename string) *Inifile {
	inifile := &Inifile{}

	newIni := newFromFile(filename)
	clearedIni := removeString(newIni, EMPTY)
	sectionsBodys := getSectionsBodys(clearedIni)
	inifile.Sections = sectionsConstruct(sectionsBodys)

	return inifile

}

// GetSectionsNames gets Inifile's sections names. Returns []string.
func (i *Inifile) GetSectionsNames() []string {
	var secnames []string

	for key := range i.Sections.SectionsMap {
		secnames = append(secnames, key)
	}

	return secnames
}

// GetSectionParams gets value by section name. Returns *Params
func (i *Inifile) GetSectionParams(secname string) *Params {
	return i.Sections.SectionsMap[secname]
}

// GetParamByName gets stored string representation of enabled parameter's by given parameter name.
func (i *Inifile) GetParamByName(secname, paramname string) string {
	params := i.GetSectionParams(secname)
	return params.Enabled[paramname]
}

// GetParamsEnabled gets enabled parameters of given section.
func (i *Inifile) GetParamsEnabled(secname string) []string {
	var enabled []string

	for _, str := range i.Sections.SectionsMap[secname].Enabled {
		enabled = append(enabled, str)
	}

	return enabled
}

// GetParamsDisabled gets disabled parameters of given section.
func (i *Inifile) GetParamsDisabled(secname string) []string {
	var disabled []string

	for _, str := range i.Sections.SectionsMap[secname].Disabled {
		disabled = append(disabled, str)
	}

	return disabled
}

// GetAllParams gets all enabled parameters of all sections.
func (i *Inifile) GetAllParams() []string {
	var params []string
	secnames := i.GetSectionsNames()

	for _, val := range secnames {
		for _, ename := range i.Sections.SectionsMap[val].Enabled {
			params = append(params, ename)
		}
	}

	return params
}

// GetValue parses parameter value. Method calls IniParser's constructor and then calls parser's
// ParseValue() method. Trys to get underlying value of string. In case of error
// returns string with error message.
func (i *Inifile) GetValue(secname, paramname string) interface{} {
	parser := NewParser()
	//params := i.GetSectionParams(secname)
	value := i.GetParamByName(secname, paramname)

	parsed := parser.ParseValue(value)

	return parsed
}

// FindParam searches given enabled parameter in all sections
func (i *Inifile) FindParam(paramname string) (string, interface{}) {
	var secname string
	var value interface{}
	secnames := i.GetSectionsNames()

	for _, val := range secnames {
		for ename := range i.Sections.SectionsMap[val].Enabled {
			if ename == paramname {
				secname = val
				value = i.GetValue(val, paramname)
				break
			}
		}
	}

	return secname, value
}

// GetErrors gets all collected errors of file
func (i *Inifile) GetErrors() []string {
	var errors []string
	secnames := i.GetSectionsNames()

	for _, val := range secnames {
		for _, errs := range i.Sections.SectionsMap[val].Errors {
			errors = append(errors, errs)
		}
	}

	return errors
}
