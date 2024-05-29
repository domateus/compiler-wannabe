package main

import (
	"fmt"
	"interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hi %s! This is the REPL for our language\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
