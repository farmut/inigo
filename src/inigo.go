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
	DELIM   string = ","
	COMM           = ";" // Ordinary comment
	UCOMM          = "#" // Unix-style comment
	EMPTY          = ""
	LSQUARE        = "["
	RSQUARE        = "]"
	EQUAL          = "="
)

type Params struct {
	Enabled  map[string]string
	Disabled map[string]string
}

type Sections map[string]Params

type Inifile struct {
	Sections    Sections
	Params      Params
	Comment     []string
	Unixcomment []string
	Empty       string
}

// Read contain of ini file, return it as slice of strings
func NewFromFile(filename string) []string {
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

// Remove string by given parameter from slice. E.g, empty string
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
		if (string(str[0]) != COMM) && (string(str[0]) != UCOMM) && (strings.Contains(str, EQUAL) != true) {
			uncommentIni = append(uncommentIni, str)
		}
	}

	return uncommentIni
}

/*
func Parse(newIni []string) *Inifile {
	iniFile = new(Inifile)
	params = new(Params)
	sections = new(Sections)
	var sectionNames []string
	var sectionSlices [][]string

	clearedIni := RemoveEmpty(newIni)

	for _, str := range clearedIni {
		if (str[0] == LSQUARE) && (str[len(str)-1:] == RSQUARE) {
			sectionNames = append(sectionNames, str)

			for _, strSec := range sectionNames {
				strIndex := strings.Index(clearedIni, sectionNames)
			}
		}

		if (str[0] != COMM) && (str[0] != UCOMM) && (str[0] != LSQUARE) {
			strSplit := string.Split(str, EQUAL)
			strKey := strSplit[0]
			strValue := strSplit[1]
			params[strKey] = strValue

			iniFile.Params = params

			return
		}
	}

}
*/

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

// Get sections names, e.g, such as [Name]
func getSections(clearedIni []string) []string {
	var sectionNames []string

	for _, str := range clearedIni {
		if (string(str[0]) == LSQUARE) && (str[len(str)-1:] == RSQUARE) {
			sectionNames = append(sectionNames, str)
		}
	}

	return sectionNames
}

/*
func getSectionsBodys(uncommentIni []string) map[string]string {
	sectionsMap := make(map[string]string)
	sectionNames := getSections(uncommentIni)
	stringBuf := strings.Join(uncommentIni, " ")
	stringIni := strings.FieldsFunc(stringBuf, func(r rune) bool {
		return r == '[' || r == ']'
	})

	for i, str := range stringIni {
		for _, name := range sectionNames {
			if str == clearSectionName(name) {
				sectionsMap[clearSectionName(name)] = stringIni[i+1]
			}
		}
	}

	return sectionsMap
}
*/

// Constructor fo Params
func paramsConstruct(body []string) *Params {
	params := new(Params)
	for _, str := range body {
		if (strings.Contains(str, EQUAL) == true) && (string(str[0]) != (COMM || UCOMM)) {
			splitted := strings.Split(str, EQUAL)
			if splitted[1] != EMPTY {
				params.Enabled[splitted[0]] = splitted[1]
			} else {
				params[splitted[0]] = nil
			}
		}
	}

	return &params
}

// Get unparsed []string as body of each Section
func getSectionsBodys(clearedIni []string) map[string][]string {
	sectionsMap := make(map[string][]string)
	sectionNames := getSections(clearedIni)

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

/*
func Parse(newIni []string) *Inifile {
	iniFile  = new(Inifile)
	params   = new(Params)
	sections = new(Sections)

	var sectionNames []string
	var sectionIndexes []int

	clearedIni := RemoveEmpty(newIni)

	for _, str := range clearedIni {
		if (str[0] == LSQUARE) && (str[len(str)-1:] == RSQUARE) {
			sectionNames = append(sectionNames, str)

}
*/
func main() {
	filename := "../ex.ini"
	//filename := "../ex_php.ini"
	iniIni := NewFromFile(filename)
	/*for i := 0; i < len(iniIni); i++ {
		fmt.Println(iniIni[i])
	}*/

	iniPini := removeString(iniIni, EMPTY)
	//	uiniPini := removeComments(iniPini)

	/*for i := 0; i < len(iniPini); i++ {
		fmt.Println(iniPini[i])
	}*/

	//cleared := clearSectionName("[Animals]")

	//fmt.Println(cleared)

	sectionsBodys := getSectionsBodys(iniPini)

	for key, val := range sectionsBodys {
		for _, valVal := range val {
			fmt.Println(key, "->", valVal)
		}
	}

	/*
		pind := GetSectionIndexes(iniPini, ind)

		for i := 0; i < len(pind); i++ {
			fmt.Println(pind[i])
		}*/
}
