/*
 * example.go
 *
 * @Author: hIMEI
 * @Date:   2017-12-13 17:49:49
 * @Last Modified by:   hIMEI
 * @Last Modified time: 2017-12-16 12:11:12
 *
 */

package main

import (
	"fmt"

	inigo "bitbucket.org/hIMEI/inigo"
)

func main() {
	//filename := "example.ini"
	filename := "short.ini"
	//filename := "example_php.ini"

	ini := inigo.NewIniFile(filename)

	ini.PrintSectionsNames()
	fmt.Println("___")

	//  params := ini.GetAllParams()
	ini.PrintAllParams()

	fmt.Println("___")

	ini.PrintSectionsNames()
	fmt.Println("___")

	//param := ini.GetParamByName("[SECTION1]", "angryByrds")

	fmt.Println("___")
	value := ini.GetValue("[SECTION1]", "param7")
	fmt.Println(value)

	/*
		for _, str := range sections {
			enParams := ini.GetParamsEnabled(str)

			for _, param := range enParams {
				value := ini.GetValue(str, param)
				fmt.Println(value)
			}
		}
	*/

}
