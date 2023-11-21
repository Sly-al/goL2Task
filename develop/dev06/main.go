package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func colsToDelete(f string) []int {
	fmt.Println(f)
	var (
		numSplit []string
		res      []int
	)
	if f == "" {
		return nil
	}
	if strings.Contains(f, "-") {
		numSplit = strings.Split(f, "-")
		start, _ := strconv.Atoi(numSplit[0])
		end, _ := strconv.Atoi(numSplit[1])
		for i := start - 1; i < end; i++ {
			res = append(res, i)
		}
		return res
	}
	if strings.Contains(f, ",") {
		numSplit = strings.Split(f, ",")
	} else {
		numSplit = strings.Split(f, " ")
	}
	for _, numStr := range numSplit {
		num, _ := strconv.Atoi(numStr)
		res = append(res, num-1)
	}
	return res
}

func cut(text []string, d string, f []int, s bool) {
	for _, line := range text {
		if s && !strings.Contains(line, d) {
			continue
		}
		if f != nil {
			lineSpl := strings.Split(line, d)
			var lineAfterDel []string
			for _, idxToGet := range f {
				if idxToGet < len(lineSpl) {
					lineAfterDel = append(lineAfterDel, lineSpl[idxToGet])
				}
			}
			fmt.Println(strings.Join(lineAfterDel, d))
		} else {
			fmt.Println(line)
		}
	}
}

func main() {
	var slStr []string
	f := flag.String("f", "", "\"fields\" - выбрать поля (колонки)")
	d := flag.String("d", "\t", "\"delimiter\" - использовать другой разделитель")
	s := flag.Bool("s", false, "\"separated\" - только строки с разделителем\n")
	flag.Parse()
	cols := colsToDelete(*f)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	for text != "" {
		slStr = append(slStr, text)
		fmt.Print("Enter text: ")
		text, _ = reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
	}
	fmt.Println(cols)
	cut(slStr, *d, cols, *s)
}
