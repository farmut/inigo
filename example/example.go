/*
 * example.go
 *
 * @Author: hIMEI
 * @Date:   2017-12-13 17:49:49
 * @Last Modified by:   hIMEI
 * @Last Modified time: 2017-12-13 17:54:48
 *
 */

package main

import (
	"fmt"

	inigo "bitbucket.org/hIMEI/inigo"
)

func main() {
	filename := "example.ini"
	//filename := "example_php.ini"

	ini := inigo.NewIniFile(filename)

	ini.PrintSectionsNames()
	fmt.Println(EMPTY)

	ini.PrintAllParams()

}
