package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	r := bufio.NewReaderSize(os.Stdin, 1<<20)
	w := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer w.Flush()

	line, _ := r.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	for range n {
		line, _ = r.ReadString('\n')
		line = strings.TrimSpace(line)
		if count := utf8.RuneCountInString(line); count > 10 {
			first, _ := utf8.DecodeRuneInString(line)
			last, _ := utf8.DecodeLastRuneInString(line)
			w.WriteString(fmt.Sprintf("%c%d%c", first, count-2, last))
			w.WriteByte('\n')
		} else {
			w.WriteString(line)
			w.WriteByte('\n')
		}
	}
}
