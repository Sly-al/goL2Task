package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type FileText struct {
	data [][]string
}

func InitFileSlice() *FileText {
	return &FileText{make([][]string, 0)}
}

func (ft *FileText) LoadFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ft.data = append(ft.data, strings.Fields(scanner.Text()))
	}
	return nil
}

func (ft *FileText) RemoveDuplicates() {
	setOfLines := make(map[string]bool)
	for i, line := range ft.data {
		lineCat := strings.Join(line, " ")
		if _, ok := setOfLines[lineCat]; ok {
			ft.data = append(ft.data[:i], ft.data[i+1:]...)
		}
		setOfLines[lineCat] = true
	}
}

func (ft *FileText) Sort(k int, n, r, u bool) {
	if u {
		ft.RemoveDuplicates()
	}
	k--
	if n {
		sort.Slice(ft.data, func(i, j int) bool {
			first, err := strconv.Atoi(ft.data[i][k])
			if err != nil {
				log.Fatal(err)
			}
			second, err := strconv.Atoi(ft.data[j][k])
			if err != nil {
				log.Fatal(nil)
			}
			return first < second
		})
	} else if k > -1 {
		sort.Slice(ft.data, func(i, j int) bool {
			return ft.data[i][k] < ft.data[j][k]
		})
	} else {
		sort.Slice(ft.data, func(i, j int) bool {
			var ansLen int
			if len(ft.data[i]) < len(ft.data[j]) {
				ansLen = len(ft.data[i])
			} else {
				ansLen = len(ft.data[j])
			}
			for x := 0; x < ansLen; x++ {
				if ft.data[i][x] == ft.data[j][x] {
					continue
				}
				return ft.data[i][x] < ft.data[j][x]
			}
			return false
		})
	}
	if r {
		ft.PrintReverse()
	} else {
		ft.Print()
	}
}

func (ft *FileText) Print() {
	for _, line := range ft.data {
		fmt.Println(strings.Join(line, " "))
	}
}

func (ft *FileText) PrintReverse() {
	for i := len(ft.data) - 1; i > -1; i-- {
		fmt.Println(strings.Join(ft.data[i], " "))
	}
}
