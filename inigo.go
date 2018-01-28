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

// Params is a parameters strings. Here's different fields
// for enabled and disabled (commented) params and field to store syntax errors.
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

// Inifile is a main data type. It is also API methods receiver. Contains Sections field.
type Inifile struct {

	// Sections is a Map [section_name]parameter_block.
	Sections *Sections
}

// NewIniFile parses ini file by given file name,
// returns pointer to Inifile structure.
func NewIniFile(filename string) *Inifile {
	inifile := &Inifile{}

	newIni := newFromFile(filename)
	clearedIni := removeString(newIni, EMPTY)
	sectionsBodys := getSectionsBodys(clearedIni)
	inifile.Sections = sectionsConstruct(sectionsBodys)

	return inifile

}

// getSectionParams returns *Params of given section
func (i *Inifile) getSectionParams(secname string) *Params {
	return i.Sections.SectionsMap[secname]
}

// GetSectionsNames gets Inifile's sections names. Returns []string.
func (i *Inifile) GetSectionsNames() []string {
	var secnames []string

	for key := range i.Sections.SectionsMap {
		secnames = append(secnames, key)
	}

	return secnames
}

// GetParamString gets stored string representation of enabled parameter's by given parameter name.
func (i *Inifile) GetParamString(secname, paramname string) string {
	params := i.getSectionParams(secname)
	return params.Enabled[paramname]
}

// GetParamsEnabled gets enabled parameters of given section in form "paramName=paramValue"
func (i *Inifile) GetParamsEnabled(secname string) []string {
	var enabled []string

	for key, val := range i.Sections.SectionsMap[secname].Enabled {
		paramStr := key + "=" + val
		enabled = append(enabled, paramStr)
	}

	return enabled
}

// GetParamsDisabled gets disabled parameters of given section in form "paramName=paramValue"
func (i *Inifile) GetParamsDisabled(secname string) []string {
	var disabled []string

	for key, val := range i.Sections.SectionsMap[secname].Disabled {
		paramStr := key + "=" + val
		disabled = append(disabled, paramStr)
	}

	return disabled
}

// GetAllParams gets all parameter's stirngs of all sections in form "paramName=paramValue"
func (i *Inifile) GetAllParams() []string {
	var params []string
	secnames := i.GetSectionsNames()

	for _, val := range secnames {
		for ename, eval := range i.Sections.SectionsMap[val].Enabled {
			paramStr := ename + "=" + eval
			params = append(params, paramStr)
		}

		for ename, eval := range i.Sections.SectionsMap[val].Disabled {
			paramStr := ename + "=" + eval
			params = append(params, paramStr)
		}
	}

	return params
}

// GetParamValue parses parameter value. Method calls IniParser's constructor and then calls parser's
// ParseValue() method. Trys to get underlying value of string. In case of error
// returns string with error message.
func (i *Inifile) GetParamValue(secname, paramname string) interface{} {
	parser := NewParser()
	//params := i.getSectionParams(secname)
	value := i.GetParamString(secname, paramname)

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
				value = i.GetParamValue(val, paramname)
				break
			}
		}
	}

	return secname, value
}

// GetErrors gets all syntax errors of file
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
