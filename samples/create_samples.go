//go:build ignore

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const contents = `
package samples

var Files = []string{%s}
`

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err)
	}
	fileNames := make([]string, 0, len(files))
	for _, file := range files {
		if !strings.Contains(file.Name(), ".txt") {
			continue
		}
		fileNames = append(fileNames, "\""+file.Name()+"\"")
	}
	samples, err := os.Create("samples.go")
	if err != nil {
		panic(err)
	}
	samples.WriteString(
		fmt.Sprintf(contents, strings.Join(fileNames, ",")),
	)
}
