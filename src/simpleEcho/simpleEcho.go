package simpleEcho

import (
	"flag"
	"os"
)

var omitNewLine = flag.Bool("n", false, "don't print final newline")

const (
	space   = ""
	newLine = "\n"
)

func Echo() {
	flag.Parse()
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += space
		}
		s += flag.Arg(i)
	}
	if !*omitNewLine {
		s += newLine
	}
	os.Stdout.WriteString(s)
}
