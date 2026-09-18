package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReaderSize(os.Stdin, 1<<20)
	w := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer w.Flush()

	line, _ := r.ReadString('\n')
	line, _ = r.ReadString('\n')
	line = strings.TrimSpace(line)

	anton, danik := 0, 0

	for _, r := range line {
		if r == 'A' {
			anton++
		} else {
			danik++
		}
	}

	if anton > danik {
		w.WriteString("Anton\n")
	} else if anton < danik {
		w.WriteString("Danik\n")
	} else {
		w.WriteString("Friendship\n")
	}
}
