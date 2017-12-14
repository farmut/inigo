/*
 * example.go
 *
 * @Author: hIMEI
 * @Date:   2017-12-13 17:49:49
 * @Last Modified by:   hIMEI
 * @Last Modified time: 2017-12-14 02:09:31
 *
 */

package main

import (
	"fmt"

	inigo "bitbucket.org/hIMEI/inigo"
)

func main() {
	filename := "../bin/new.ini"
	//filename := "example_php.ini"

	ini := inigo.NewIniFile(filename)

	ini.PrintSectionsNames()
	fmt.Println("___")

	//  params := ini.GetAllParams()
	ini.PrintAllParams()

	sections := ini.GetSectionsNames()

	for _, str := range sections {
		enParams := ini.GetParamsEnabled(str)

		for _, param := range enParams {
			value := ini.GetValue(str, param)
			fmt.Println(value)
		}
	}
}
