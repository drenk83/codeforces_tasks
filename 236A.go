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

	hash := make(map[rune]struct{}, len(line))
	for _, char := range line {
		hash[char] = struct{}{}
	}

	if len(hash)%2 == 0 {
		w.WriteString("CHAT WITH HER!")
	} else {
		w.WriteString("IGNORE HIM!")
	}
	w.WriteByte('\n')
}
