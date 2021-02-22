package main

import (
	"errors"
	"fmt"
	"os"

	"./calc"
)

func main() {

	if len(os.Args) != 2 {
		panic(errors.New("wrong number of commandline args"))
	}

	fmt.Println(calc.Calculate(os.Args[1]))
}
