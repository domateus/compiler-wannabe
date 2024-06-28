package main

import (
	"flag"
	"fmt"
	"interpreter/repl"
	"os"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	flag.Parse()
	fmt.Printf("Hi %s! This is the REPL for our language\n", usr.Username)
	repl.Start(os.Stdin, os.Stdout)
}
