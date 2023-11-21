package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type LineWithIdx struct {
	idx  int
	line string
}

type FileLines struct {
	data []LineWithIdx
}

func InitFileLines() *FileLines {
	return &FileLines{
		data: make([]LineWithIdx, 0),
	}
}

func (fl *FileLines) ReadFile(filePath string, toLower bool) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		if toLower {
			fl.data = append(fl.data, LineWithIdx{
				idx:  i,
				line: strings.ToLower(scanner.Text()),
			})
		} else {
			fl.data = append(fl.data, LineWithIdx{
				idx:  i,
				line: scanner.Text(),
			})
		}

		i++
	}
	return nil
}

func (fl *FileLines) Compute(A, B, C int, c, V, F, n bool, pattern string) {
	if V {
		fl.invertFind(c, n, pattern)
	} else if F {
		fl.simpleFind(c, n, pattern)
	} else {
		fl.grep(A, B, C, c, n, pattern)
	}
}

func (fl *FileLines) simpleFind(c, n bool, pattern string) {
	result := InitFileLines()
	for _, curLine := range fl.data {
		if strings.Contains(curLine.line, pattern) {
			result.data = append(result.data, curLine)
		}
	}
	result.Print(c, n)
}

func (fl *FileLines) invertFind(c, n bool, pattern string) {
	result := InitFileLines()
	compRegExp, err := regexp.Compile(pattern)
	if err != nil {
		log.Println(err)
	}
	for _, curLine := range fl.data {
		if !compRegExp.MatchString(curLine.line) {
			result.data = append(result.data, curLine)
		}
	}
	result.Print(c, n)
}

func (fl *FileLines) grep(A, B, C int, c, n bool, pattern string) {
	result := InitFileLines()
	compRegExp, err := regexp.Compile(pattern)
	if err != nil {
		log.Println(err)
	}
	if A > 0 { // вывод A строк до вхождения
		next := 0
		for idx, curLine := range fl.data {
			if compRegExp.MatchString(curLine.line) {
				if idx >= next {
					result.data = append(result.data, fl.data[idx:Min(idx+A+1, len(fl.data))]...)
				} else {
					result.data = append(result.data, fl.data[next:Min(idx+A+1, len(fl.data))]...)
				}
				next = idx + A + 1
			}
		}
	} else if B > 0 { // вывод B строк до вхождения
		last := 0
		for idx, curLine := range fl.data {
			if compRegExp.MatchString(curLine.line) {
				result.data = append(result.data, fl.data[Max(last, idx-B):idx+1]...)
				last = idx + 1
			}
		}
	} else if C > 0 { // вывод C строк до и после вхождения
		next := 0
		for idx, curLine := range fl.data {
			if compRegExp.MatchString(curLine.line) {
				if idx > next {
					result.data = append(result.data, fl.data[Max(next, idx-C):idx+C+1]...)
				} else {
					result.data = append(result.data, fl.data[next:Min(idx+C+1, len(fl.data))]...)
				}
				next = idx + C + 1
			}
		}
	}
	//fmt.Println(result.data)
	result.Print(c, n)
}

func (fl *FileLines) Print(c, n bool) {
	if c { // вывод количества строк
		fmt.Println(len(fl.data))
	} else {
		if n { // вывод строки с её номером
			for _, lineWithIdx := range fl.data {
				fmt.Printf("%d %v \n", lineWithIdx.idx, lineWithIdx.line)
			}
		} else {
			for _, lineSimple := range fl.data {
				fmt.Printf("%v \n", lineSimple.line)
			}
		}
	}
}

func MatchRegExp(regExp string, curLine LineWithIdx) bool {
	matched, err := regexp.MatchString(regExp, curLine.line)
	if err != nil {
		log.Fatal(err)
	}
	return matched
}

func Max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func Min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}
