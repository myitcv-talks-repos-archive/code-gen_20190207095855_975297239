package main

import "fmt"

type GoSheffield int

//go:generate typename GoSheffield

func main() {
	var gopher GoSheffield = 5
	fmt.Printf("%v (%v)\n", gopher, gopher.TypeName())
}
