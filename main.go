package main

import (
	"fmt"
	"os"
)

func main() {
	arg := getArgs()

	switch arg[1] {
	case "clone":
		fmt.Println("Clone")
	default:
		fmt.Printf("%s argument does not exist.\n", arg[1])
	}
}

func getArgs() []string {
	if len(os.Args) == 1 {
		fmt.Println("No argument exists.")
		os.Exit(1)
	}

	return os.Args
}
