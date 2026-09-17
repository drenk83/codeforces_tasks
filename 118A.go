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
	line = strings.TrimSpace(line)
	line = strings.ToLower(line)

	glasnye := map[rune]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'y': {},
	}

	out := make([]rune, 0, len(line)*2)
	for _, r := range line {
		if _, ok := glasnye[r]; !ok {
			out = append(out, '.')
			out = append(out, r)
		}
	}

	w.WriteString(string(out))
	w.WriteByte('\n')

}
