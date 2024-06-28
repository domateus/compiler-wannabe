package parser

import (
	"flag"
	"fmt"
	"strings"
)

// serves the purpose of tracing the parser, showing method by method how it is working
var traceLevel int = 0

var tracing bool = false

const traceIdentPlaceholder string = "\t"

func init() {
	traceFlag := flag.Bool("trace", false, "imprimir estado de parser")
	flag.Parse()
	tracing = *traceFlag
}

func identLevel() string {
	return strings.Repeat(traceIdentPlaceholder, traceLevel-1)
}

func tracePrint(fs string) {
	if tracing {
		fmt.Printf("%s%s\n", identLevel(), fs)
	}
}

func incIdent() {
	traceLevel = traceLevel + 1
}

func decIdent() {
	traceLevel = traceLevel - 1
}

func trace(msg string) string {
	incIdent()
	tracePrint("BEGIN " + msg)
	return msg
}

func untrace(msg string) {
	tracePrint("END " + msg)
	decIdent()
}
