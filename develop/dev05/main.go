package main

import (
	"dev05/lib"
	"flag"
	"fmt"
)

func main() {
	A := flag.Int("A", 0, "\"after\" печатать +N строк после совпадения")
	B := flag.Int("B", 0, "\"before\" печатать +N строк до совпадения")
	C := flag.Int("C", 0, "\"context\" (A+B) печатать ±N строк вокруг совпадения")
	c := flag.Bool("c", false, "\"count\" (количество строк)")
	i := flag.Bool("i", false, "\"ignore-case\" (игнорировать регистр)")
	V := flag.Bool("V", false, "\"invert\" (вместо совпадения, исключать)")
	F := flag.Bool("F", false, "\"fixed\", точное совпадение со строкой, не паттерн")
	n := flag.Bool("n", false, "\"line num\", напечатать номер строки")
	flag.Parse()
	filePath := flag.Arg(len(flag.Args()) - 1)
	regExp := flag.Arg(len(flag.Args()) - 2)
	fmt.Println(*A, *B, *C, *c, *i, *V, *F, *n)
	_ = regExp
	File := lib.InitFileLines()
	err := File.ReadFile(filePath, *i)
	if err != nil {
		fmt.Println(err)
	}
	//File.Print(*c, *n)
	File.Compute(*A, *B, *C, *c, *V, *F, *n, regExp)
}
