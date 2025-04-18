package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/queone/utl"
	"github.com/sethvargo/go-diceware/diceware"
)

const (
	program_name    = "pgen"
	program_version = "1.2.2"
	delimiter       = "_" // Use underscore for pwd word delimiter
)

func printUsage() {
	n := utl.Whi2(program_name)
	v := program_version
	usageHeader := fmt.Sprintf("%s v%s\n"+
		"Memorable password generator - github.com/git719/pgen\n"+
		"%s\n"+
		"  %s [option]\n\n"+
		"%s\n"+
		"                     Without arguments it generates a 3-word memorable password phrase\n"+
		"  NUMBER             Generates a NUMBER-word memorable password phrase\n"+
		"                     For example, if NUMBER is '6' it generates a 6-word phrase\n"+
		"                     Mininum is 1, maximum is 9\n"+
		"  -?, -h, --help     Print this usage page\n",
		n, v, utl.Whi2("Usage"), n, utl.Whi2("Options"))
	fmt.Print(usageHeader)
	os.Exit(0)
}

func GeneratePassphrase(words int) {
	// Generate 6 words using the diceware algorithm.
	list, err := diceware.Generate(words)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.Join(list, delimiter))
}

func main() {
	args := len(os.Args[1:]) // Not including the program itself
	switch args {
	case 0: // Process 0-argument requests
		GeneratePassphrase(3)
	case 1: // Process 1-argument requests
		arg1 := os.Args[1]
		switch arg1 {
		case "-?", "-h", "--help":
			printUsage()
		default:
			numberOfWords, err := strconv.Atoi(arg1)
			if err == nil && numberOfWords > 0 && numberOfWords < 10 {
				GeneratePassphrase(numberOfWords)
			} else {
				fmt.Println("NUMBER must be 1 thru 9.")
			}
		}
	}
	os.Exit(0)
}
