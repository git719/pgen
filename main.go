// main.go

package main

import (
	"fmt"
	"github.com/sethvargo/go-diceware/diceware"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	prgname = "pgen"
	prgver  = "1.0.0"
)

func printUsage() {
	fmt.Printf(prgname + " Memorable password generator v" + prgver + "\n" +
		"               With no arguments, it generates a 4-word memorable password phrase\n" +
		"    NUMBER     Generate a NUMBER-word memorable password phrase\n" +
		"                 For example, if NUMBER is '6' it generates a 6-word phrase\n" +
		"                 Mininum is 1, maximum is 99\n" +
		"    -v         Print this usage page\n")
	os.Exit(0)
}

func GeneratePassphrase(words int) {
	// Generate 6 words using the diceware algorithm.
	list, err := diceware.Generate(words)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.Join(list, "-"))
}

func main() {
	args := len(os.Args[1:]) // Not including the program itself
	switch args {
	case 0: // Process 0-argument requests
		GeneratePassphrase(4)
	case 1: // Process 1-argument requests
		arg1 := os.Args[1]
		if arg1 == "-v" {
			printUsage()
		} else {
			numberOfWords, err := strconv.Atoi(arg1)
			if err == nil && numberOfWords > 0 && numberOfWords < 100 {
				GeneratePassphrase(numberOfWords)
			} else {
				fmt.Println("NUMBER must be 1 thru 99.")
			}
		}
	}
	os.Exit(0)
}
