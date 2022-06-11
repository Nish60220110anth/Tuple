package main

import (
	"fmt"
	"tuple/src/tuple1"
)

func main() {
	tuple1 := tuple1.NewTuple1("Nis")
	fmt.Println((*tuple1).String())

	(*tuple1).Set(0, 3)
	fmt.Println((*tuple1).String())
}
