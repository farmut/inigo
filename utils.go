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

/* File utils.go contains unexported functions and local helpers */

package inigo

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// newFromFile reads contains of ini file, returns slice of strings
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

// removeString removes string from slice by given parameter. Empty string, for example.
func removeString(newIni []string, junk string) []string {
	for i, str := range newIni {
		if str == junk {
			newIni = append(newIni[:i], newIni[i+1:]...)
		}
	}

	return newIni
}

// findIndexByName finds given string's index in slice
func findIndexByName(names []string, name string) int {
	var index int
	for i, str := range names {
		if str == name {
			index = i
		}
	}

	return index
}

// parseParamName parses given parameter name, returns false in case of syntax errors.
// If syntax is correct, returns true
func parseParamName(paramname string) bool {
	check := true

	if string(paramname[0]) != COMM && string(paramname[0]) != UCOMM {
		if strings.Contains(paramname, SPACE) == true &&
			strings.Index(paramname, SPACE) != (len(paramname)-1) {
			check = false
		} else if (strings.Contains(paramname, COMM) == true && strings.Index(paramname, COMM) != 0) ||
			(strings.Contains(paramname, UCOMM) == true && strings.Index(paramname, COMM) != 0) {
			check = false
		} else {
			check = true
		}
	}

	return check
}

// getSections gets sections names, e.g, [Name]
func getSections(clearedIni []string) []string {
	var sectionNames []string

	for _, str := range clearedIni {
		if (string(str[0]) == LSQUARE) && (str[len(str)-1:] == RSQUARE) {
			sectionNames = append(sectionNames, str)
		}
	}

	return sectionNames
}

// splitParamString splits given paramString with EQUAL
func splitParamString(paramString string) []string {
	splitted := strings.Split(paramString, EQUAL)

	return splitted
}

// joinParamStrings joins []string to string with "=" as delimiter
func joinParamStrings(creds []string) string {
	paramString := strings.Join(creds, EQUAL)

	return paramString
}

// removeLastSpace removes trailing space from parameter's name. Returns trimmed paramname.
func removeLastSpace(paramname string) string {
	if strings.Contains(paramname, SPACE) == true &&
		strings.Index(paramname, SPACE) == (len(paramname)-1) {
		paramname = strings.TrimSuffix(paramname, SPACE)
	}

	return paramname
}

// removeInline removes inline comments form parameter's string value
func removeInline(stringvalue string) string {
	var newValue string

	switch {
	case strings.Contains(stringvalue, COMM) == true:
		splitted := strings.Split(stringvalue, COMM)
		newValue = splitted[0]

	case strings.Contains(stringvalue, UCOMM) == true:
		splitted := strings.Split(stringvalue, UCOMM)
		newValue = splitted[0]

	default:
		newValue = stringvalue
	}

	return newValue
}

// removeFirstSpace removes leading space from parameter's string value.
// Returns trimmed value.
func removeFirstSpace(stringvalue string) string {
	if string(stringvalue[0]) == SPACE {
		stringvalue = strings.TrimPrefix(stringvalue, SPACE)
	}

	return stringvalue
}

// checkDubls checks for duplicated paramnames.
func checkDubls(paramnames []string, paramname string) bool {
	check := false

	for i, str := range paramnames {
		if paramname == str && findIndexByName(paramnames, paramname) != i {
			check = true
			break
		}
	}

	return check
}

// renameDubls if duplicated paramnames found, renames thats with incremented siffix.
func renameDubls(paramnames []string) []string {
	for i, name := range paramnames {
		for j, checkedname := range paramnames {
			if (name == checkedname) && (i != j) {
				itoa := strconv.Itoa(i)
				name += itoa
			}
		}
	}

	return paramnames
}

// checkForEqual checks parameter name for EQUAL
func checkForEqual(paramstring string) bool {
	check := false
	if strings.Contains(paramstring, EQUAL) == true {
		check = true
	}

	return check
}

// checkSpaceComma checks parameter's string value for SPACE and COMMA.
// If string contains SPACE and if SPACE
// not follows the COMMA (e.g, syntax error), returns false
func checkSpaceComma(stringvalue string) bool {
	check := true

	if strings.Contains(stringvalue, SPACE) == true &&
		strings.Contains(stringvalue, CSPACE) == false {
		check = false
	}

	return check
}

// checkErrors is a syntax error's checker for parameter's string value
func checkErrors(stringvalue string) bool {
	check := false

	if stringvalue[0] != DQUOTE && stringvalue[len(stringvalue)-1] != DQUOTE {
		if checkSpaceComma(stringvalue) == false {
			check = true
		}
	}

	return check
}

// prepareValue cuts trailing and leading spaces and inline comments if found
func prepareValue(stringvalue string) string {
	stringvalue = removeLastSpace(removeFirstSpace(removeInline(stringvalue)))

	return stringvalue
}

// errorConstruct constructor for Params.Error
func errorConstruct(body []string) []string {
	var errors []string
	var value string

	for _, param := range body {
		if checkForEqual(param) == true {
			splitted := splitParamString(param)
			if splitted[1] != EMPTY {
				value = prepareValue(splitted[1])
			}

			if parseParamName(splitted[0]) == false ||
				checkErrors(value) == true {
				errors = append(errors, param)
			}
		}
	}

	return errors
}

// paramsConstruct constructor for Params
func paramsConstruct(body []string) *Params {
	params := &Params{}

	params.Errors = errorConstruct(body)
	params.Enabled = make(map[string]string)
	params.Disabled = make(map[string]string)

	for _, str := range body {
		if checkForEqual(str) == true {
			if (string(str[0]) != COMM) && (string(str[0]) != UCOMM) {
				splitted := splitParamString(str)
				if splitted[1] == EMPTY {
					params.Enabled[splitted[0]] = NONE
				} else {
					// Non empty params
					value := prepareValue(splitted[1])

					if checkErrors(value) != true {
						params.Enabled[splitted[0]] = value
					}
				}

			} else {
				// Commented params
				splitted := splitParamString(str)
				if splitted[1] == EMPTY {
					params.Disabled[splitted[0]] = NONE
				} else {
					// Non emty params
					value := prepareValue(splitted[1])

					if checkErrors(value) != true {
						params.Disabled[splitted[0]] = value
					}
				}
			}
		}
	}

	return params
}

// sectionsConstruct constructor for Sections
func sectionsConstruct(sectionsMap map[string][]string) *Sections {
	sections := &Sections{}
	sections.SectionsMap = make(map[string]*Params)

	for key, value := range sectionsMap {
		sections.SectionsMap[clearSectionName(key)] = paramsConstruct(value)
	}

	return sections
}

// getSectionsBodys gets unparsed []string as body of each Section
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

// clearSectionName removes "[" and "]" from section name
func clearSectionName(sectionName string) string {
	sectionName = strings.TrimPrefix(sectionName, LSQUARE)
	sectionName = strings.TrimSuffix(sectionName, RSQUARE)

	return sectionName
}
