package main

import (
	"github.com/metonimie/monkeyInterpreter/repl"
	"os"
)

func main() {
	print("Welcome to the MonkeyScript repl.\n")
	repl.Start(os.Stdin, os.Stdout)
}
