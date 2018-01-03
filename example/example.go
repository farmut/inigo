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

	sections := ini.GetSectionsNames()

	fmt.Println("Sections:")

	for _, sec := range sections {
		fmt.Println(sec)

		/*		e := ini.GetParamsEnabled(sec)

				fmt.Println("Enabled params:")

				for _, p := range e {
					fmt.Println(p)
					fmt.Println("Value:")
					value := ini.GetValue(sec, p)
					fmt.Println(value)
				}

				d := ini.GetParamsDisabled(sec)

				fmt.Println("Disabled params:")

				for _, p := range d {
					fmt.Println(p)
				}*/
	}
	/*
		all := ini.GetAllParams()

		fmt.Println("All Params:")

		for _, a := range all {
			fmt.Println(a)
		}*/

}
