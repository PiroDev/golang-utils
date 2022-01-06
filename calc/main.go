package main

import (
	"fmt"
	"log"
	"os"

	"./calc"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal("wrong number of commandline args")
	}

	calculated, err := calc.Calculate(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(calculated)
}
