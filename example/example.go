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

	fmt.Println("==========")
	fmt.Println("Sections")
	fmt.Println("==========")

	secs := ini.GetSectionsNames()
	for _, sec := range secs {
		fmt.Println(sec)
	}
	fmt.Println("==========")
	fmt.Println("All params")
	fmt.Println("==========")

	params := ini.GetAllParams()
	for _, param := range params {
		fmt.Println(param)
	}

	fmt.Println("==========")
	fmt.Println("Disabled params")
	fmt.Println("==========")

	for _, sec := range secs {
		disparam := ini.GetParamsDisabled(sec)
		for _, dis := range disparam {
			fmt.Println(dis)
		}
	}

	fmt.Println("==========")
	fmt.Println("Param by name")
	fmt.Println("==========")

	namedParam := ini.GetParamString("insects", "spiderLegs")
	fmt.Println(namedParam)

	fmt.Println("==========")
	fmt.Println("Value")
	fmt.Println("==========")

	//param := ini.GetParamByName("[SECTION1]", "angryByrds")

	value := ini.GetParamValue("insects", "spiderLegs")
	fmt.Println(value)

	fmt.Println("==========")

	sec, val := ini.FindParam("bird3")

	fmt.Println(sec, val)

	fmt.Println("==========")

	errors := ini.GetErrors()
	for _, d := range errors {
		fmt.Println(d)
	}

}
