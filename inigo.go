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
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Comments struct {
	Comms []string `json:"comms"`
}

type Params struct {
	Comments Comments          `json:"comments"`
	Enabled  map[string]string `json:"Enabled"`
	Disabled map[string]string `json:"Disabled"`
	Errors   []string          `json:"Errors"`
}

type Sections struct {
	SectionsMap map[string]*Params `json:"sectionsmap"`
}

type Inifile struct {
	Sections *Sections `json:"sections"`
}

///////////////////////////////////////////////////////////////////////////////
//
//                           **Library's API is here:**
//
///////////////////////////////////////////////////////////////////////////////
//
// * Functions
// 		- NewIniFile(filename string) *Inifile
//
// * Methods
// 		- IniToJson() ([]byte, error)
// 		- GetSectionsNames() []string
// 		- PrintSectionsNames()
// 		- GetSectionByName(secname string) *Params
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
///////////////////////////////////////////////////////////////////////////////

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
func (i *Inifile) GetSectionByName(secname string) *Params {
	return i.Sections.SectionsMap[secname]
}

// Gets stored string representation of enabled parameter's value by given parameter name.
func (i *Inifile) GetParamByName(params *Params, paramname string) string {
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

// Parses parameter value.
// Method calls IniParser's constructor end then calls parser's Parse() method.
// Trys to get underlying value of string value.
// In case of error return string with error message.
func (i *Inifile) GetValue(secname, paramname string) interface{} {
	parser := NewParser()
	params := i.GetSectionByName(secname)
	value := i.GetParamByName(params, paramname)

	parsed := parser.Parse(value)

	return parsed
}

// Prints parameters of all sections.
func (i *Inifile) PrintAllParams() {
	params := i.GetAllParams()

	for _, str := range params {
		fmt.Println(str)
	}
}

//////////////////////////////////////
// Private functions and local helpers
//////////////////////////////////////

// Reads contain of ini file, returns slice of strings
func newFromFile(filename string) []string {
	var newIni []string
	file, _ := os.Open(filename)

	//	ErrCheck(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		newIni = append(newIni, scanner.Text())
	}

	return newIni
}

// Removes string from slice by given parameter. Empty string, for example.
func removeString(newIni []string, junk string) []string {
	for i, str := range newIni {
		if str == junk {
			newIni = append(newIni[:i], newIni[i+1:]...)
		}
	}

	return newIni
}

// Finds given string's index in slice
func findIndexByName(names []string, name string) int {
	var index int
	for i, str := range names {
		if str == name {
			index = i
		}
	}

	return index
}

// Parses given parameter name, returns false in case of syntax errors.
// If syntax is correct, returns true
func parseParamName(paramname string) bool {
	var check bool

	if strings.Contains(paramname, SPACE) == true &&
		strings.Index(paramname, SPACE) != (len(paramname)-1) {
		check = false
	} else if strings.Contains(paramname, COMM) == true ||
		strings.Contains(paramname, UCOMM) == true {
		check = false
	} else {
		check = true
	}

	return check
}

// Gets sections names, e.g, [Name]
func getSections(clearedIni []string) []string {
	var sectionNames []string

	for _, str := range clearedIni {
		if (string(str[0]) == LSQUARE) && (str[len(str)-1:] == RSQUARE) {
			sectionNames = append(sectionNames, str)
		}
	}

	return sectionNames
}

// Splits given paramString with EQUAL
func splitParamString(paramString string) []string {
	splitted := strings.Split(paramString, EQUAL)

	return splitted
}

// Join []string to string with "=" as delimiter
func joinParamStrings(creds []string) string {
	paramString := strings.Join(creds, EQUAL)

	return paramString
}

// Constructor for Params
func paramsConstruct(body []string) *Params {
	params := &Params{}
	params.Enabled = make(map[string]string)
	params.Disabled = make(map[string]string)

	for _, str := range body {
		if strings.Contains(str, EQUAL) == true {
			if (string(str[0]) != COMM) || (string(str[0]) != UCOMM) {
				splitted := splitParamString(str)
				if splitted[1] != EMPTY {
					params.Enabled[splitted[0]] = splitted[1]
				} else {
					params.Enabled[splitted[0]] = NONE
				}

			} else {
				splitted := splitParamString(str)
				if splitted[1] != EMPTY {
					params.Disabled[splitted[0]] = splitted[1]
				} else {
					params.Disabled[splitted[0]] = NONE
				}
			}
		}

		if (strings.Contains(str, EQUAL) != true) &&
			((string(str[0]) == COMM) || (string(str[0]) == UCOMM)) {
			params.Comments.Comms = append(params.Comments.Comms, str)
		}
	}

	return params
}

// Constructor for Sections
func sectionsConstruct(sectionsMap map[string][]string) *Sections {
	sections := &Sections{}
	sections.SectionsMap = make(map[string]*Params)

	for key, value := range sectionsMap {
		sections.SectionsMap[key] = paramsConstruct(value)
	}

	return sections
}

// Gets unparsed []string as body of each Section
func getSectionsBodys(clearedIni []string) map[string][]string {
	sectionsMap := make(map[string][]string)
	sectionNames := getSections(clearedIni)
	sectionHeadless := clearedIni[:findIndexByName(clearedIni, sectionNames[0])]

	// For parameters without section name e.g. GLOBAL parameters
	sectionsMap[GLOBAL] = sectionHeadless

	for i, str := range clearedIni {
		for j, name := range sectionNames {
			if str == name {

				// Index is always in range of slice
				if j != (len(sectionNames) - 1) {
					nextName := sectionNames[j+1]
					nextInd := findIndexByName(clearedIni, nextName)
					body := clearedIni[i:nextInd]
					sectionsMap[name] = body
				} else {
					body := clearedIni[i:]
					sectionsMap[name] = body
				}

			}
		}
	}

	return sectionsMap
}

// Removes "[" and "]" from section name
func clearSectionName(sectionName string) string {
	sectionName = strings.TrimPrefix(sectionName, LSQUARE)
	sectionName = strings.TrimSuffix(sectionName, RSQUARE)

	return sectionName
}
