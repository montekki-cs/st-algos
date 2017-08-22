package main

import (
	"./karatsuba"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	fmt.Println(karatsuba.KaratsubaMul(args[0], args[1]))
}
