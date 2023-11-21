package main

import (
	"fmt"
	sortlib "goL2Task/develop/dev03/lib"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		k        int
		n, r, u  bool
		fileName string
	)
	flags := os.Args[1]
	if strings.Contains(flags, "-") {
		if strings.Contains(flags, "k") {
			k, _ = strconv.Atoi(os.Args[2])
			fileName = os.Args[3]
		} else {
			k = 0
			fileName = os.Args[2]
		}
		if strings.Contains(flags, "n") {
			n = true
		}
		if strings.Contains(flags, "r") {
			r = true
		}
	} else {
		fileName = os.Args[1]
	}
	fileText := sortlib.InitFileSlice()
	err := fileText.LoadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fileText.Sort(k, n, r, u)
}
