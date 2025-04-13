package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/azdanov/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Fprintln(os.Stdout, "Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
