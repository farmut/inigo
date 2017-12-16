//
// @Author: hIMEI
// @Date:   2017-12-14 12:15:21
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

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

type Editor struct {

	// Boolean const
	ENABLED Nbld

	// All file's content. Empty strings included
	Content []string
}

func (e Editor) EditFile(filename string) []string {
	//e.Content = make([]string)
	file, err := os.OpenFile(filename, os.O_RDWR, 0755)

	ErrFatal(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		e.Content = append(e.Content, scanner.Text())
	}

	return e.Content
}

func (e Editor) enableParam(paramstring string) {
	for i, str := range e.Content {
		if strings.Contains(str, paramstring) {
			e.Content[i] = strings.TrimPrefix(e.Content[i], COMM)
		}
	}

	newContent := strings.Join(e.Content, "\n")
	err := ioutil.WriteFile(FILENAME, []byte(newContent), 0644)

	ErrFatal(err)
}
