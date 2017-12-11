/////////////////////////////////////////////////////
// @Author: hIMEI
// @Date:   2017-12-06 07:15:25
// @Last Modified by:   hIMEI
// @Last Modified time: 2017-12-06 07:15:25
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

//package inigo

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//"sort"
	"strings"
)

const (
	DELIM    string = ","
	COMM            = ";" // Ordinary comment
	UCOMM           = "#" // Unix-style comment
	EMPTY           = ""
	LSQUARE         = "["
	RSQUARE         = "]"
	EQUAL           = "="
	HEADLESS        = "Headless"
	NULL            = "0"
)

type Comments []string

type Params struct {
	Comments Comments
	Enabled  map[string]string
	Disabled map[string]string
}

type Sections map[string]*Params

type Inifile struct {
	Sections *Sections
}

// Parse ini file by given file name,
// return pointer to app main type - Inifile structure.
func NewIniFile(filename string) *Inifile {
	inifile := new(Inifile)

	newIni := newFromFile(filename)
	clearedIni := removeString(newIni, EMPTY)
	sectionsBodys := getSectionsBodys(clearedIni)
	sections := sectionsConstruct(sectionsBodys)

	return inifile

}

// Get Inifile's sections names. Return []string.
func (i *Inifile) GetSectionsNames() []string {
	var secnames []string

	for key := range i.Sections {
		secnames = append(secnames, key)
	}

	return secnames
}

// Print section names.
func (i *Inifile) PrintSectionsNames() {
	secnames := i.GetSectionsNames()

	for _, str := range secnames {
		fmt.Println(str)
	}
}

// Get enabled parameters for given section.
func (i *Inifile) GetParamsEnabled(secname string) []string {
	var enabled []string

	for _, str := range i.Sections[secname] {
		enabled = append(enabled, str)
	}

	return enabled
}

// Print enabled parameters for given section.
func (i *Inifile) PrintParamsEnabled(secname string) {
	paramsenabled := i.GetParamsEnabled(secname)

	for _, str := range paramsenabled {
		fmt.Println(str)
	}
}

// Get disabled parameters for given section.
func (i *Inifile) GetParamsDisabled(secname string) []string {
	var disabled []string

	for _, str := range i.Sections[secname] {
		disabled = append(disabled, str)
	}

	return disabled
}

func (i *Inifile) GetPVal() {

}

// Print disabled parameters for given section.
func (i *Inifile) PrintParamsDisabled(secname string) {
	paramsdisabled := i.GetParamsEnabled(secname)

	for _, str := range paramsdisabled {
		fmt.Println(str)
	}
}

// Get parameters of all sections.
func (i *Inifile) GetAllParams() []string {
	var params []string
	secnames := i.GetSectionsNames()

	for _, val := range secnames {
		for ename := range val.Enabled {
			params = append(params, ename)
		}

		for dname := range val.Disabled {
			params = append(params, dname)
		}
	}

	return params
}

// Print parameters of all sections.
func (i *Inifile) PrintAllParams() {
	params := i.GetAllParams()

	for _, str := range params {
		fmt.Println(str)
	}
}

// Read contain of ini file, return slice of strings
func newFromFile(filename string) []string {
	var newIni []string
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		newIni = append(newIni, scanner.Text())
	}

	return newIni
}

// Remove string from slice by given parameter. Empty string, for example.
func removeString(newIni []string, junk string) []string {
	for i, str := range newIni {
		if str == junk {
			newIni = append(newIni[:i], newIni[i+1:]...)
		}
	}

	return newIni
}

// Remove commented strings excepts EQUAL (e.g, "=") containing ones
func removeComments(clearedIni []string) []string {
	var uncommentIni []string

	for _, str := range clearedIni {
		if (string(str[0]) != COMM) && (string(str[0]) != UCOMM) &&
			(strings.Contains(str, EQUAL) != true) {
			uncommentIni = append(uncommentIni, str)
		}
	}

	return uncommentIni
}

// Find given string's index in slice
func findIndexByName(names []string, name string) int {
	var index int
	for i, str := range names {
		if str == name {
			index = i
		}
	}

	return index
}

// Get sections names, e.g, [Name]
func getSections(clearedIni []string) []string {
	var sectionNames []string

	for _, str := range clearedIni {
		if (string(str[0]) == LSQUARE) && (str[len(str)-1:] == RSQUARE) {
			sectionNames = append(sectionNames, str)
		}
	}

	return sectionNames
}

// Split given paramString with EQUAL
func splitParamString(paramString string) []string {
	splitted := strings.Split(paramString, EQUAL)

	return splitted
}

// Constructor for Params
func paramsConstruct(body []string) *Params {
	params := new(Params)
	for _, str := range body {
		if strings.Contains(str, EQUAL) == true {
			if (string(str[0]) != COMM) || (string(str[0]) != UCOMM) {
				splitted := splitParamString(str)
				if splitted[1] != EMPTY {
					params.Enabled[splitted[0]] = splitted[1]
				} else {
					params.Enabled[splitted[0]] = NULL
				}
				// string(str[0]) == (COMM || UCOMM))
			} else {
				splitted := splitParamString(str)
				if splitted[1] != EMPTY {
					params.Disabled[splitted[0]] = splitted[1]
				} else {
					params.Disabled[splitted[0]] = NULL
				}
			}
		}

		if (strings.Contains(str, EQUAL) != true) &&
			((string(str[0]) == COMM) || (string(str[0]) == UCOMM)) {
			params.Comments = append(params.Comments, str)
		}
	}

	return *params
}

// Constructor for Sections
func sectionsConstruct(sectionsMap map[string][]string) *Sections {
	sections := new(Sections)
	for key, value := range sectionsMap {
		sections[key] = paramsConstruct(value)
	}

	return *sections
}

// Get unparsed []string as body of each Section
func getSectionsBodys(clearedIni []string) map[string][]string {
	sectionsMap := make(map[string][]string)
	sectionNames := getSections(clearedIni)
	sectionHeadless := clearedIni[:sectionNames[0]]

	// For parameters wich havenot sections
	sectionsMap[HEADLESS] = sectionHeadless

	for i, str := range clearedIni {
		for j, name := range sectionNames {
			if str == name {

				// Index always in range of slice
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

// Remove "[" and "]" from section name
func clearSectionName(sectionName string) string {
	sectionName = strings.TrimPrefix(sectionName, LSQUARE)
	sectionName = strings.TrimSuffix(sectionName, RSQUARE)

	return sectionName
}

func main() {
	filename := "../example/ex.ini"
	//filename := "../example/ex_php.ini"

	ini := NewIniFile(filename)

	ini.PrintSectionsNames()
	fmt.Println(EMPTY)

	ini.PrintAllParams()
}
