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

package main

import (
	"fmt"

	inigo "github.com/hIMEI29A/inigo"
)

func main() {
	filename := "example.ini"
	//filename := "short.ini"
	//filename := "example_php.ini"

	ini := inigo.NewIniFile(filename)

	secs := ini.GetSectionsNames()
	for _, sec := range secs {
		fmt.Println(sec)
	}

	fmt.Println("___")

	params := ini.GetAllParams()
	for _, param := range params {
		fmt.Println(param)
	}

	fmt.Println("___")

	//param := ini.GetParamByName("[SECTION1]", "angryByrds")

	value := ini.GetValue("insects", "ColonyPopulation")
	fmt.Println(value)

	fmt.Println("___")

	sec, val := ini.FindParam("bird3")

	fmt.Println(sec, val)

	fmt.Println("==========")

	errors := ini.GetErrors()
	for _, d := range errors {
		fmt.Println(d)
	}

}
