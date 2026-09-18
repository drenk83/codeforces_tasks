package main

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

func main() {
	r := bufio.NewReaderSize(os.Stdin, 1<<20)
	w := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer w.Flush()

	line, _ := r.ReadString('\n')
	line = strings.TrimSpace(line)
	upper := 0
	lower := 0

	for _, r := range line {
		if unicode.IsUpper(r) {
			upper++
		} else {
			lower++
		}
	}

	if upper > lower {
		line = strings.ToUpper(line)
	} else {
		line = strings.ToLower(line)
	}
	w.WriteString(line)
	w.WriteByte('\n')
}
