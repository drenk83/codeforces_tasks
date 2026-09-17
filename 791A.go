package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReaderSize(os.Stdin, 1<<20)
	w := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer w.Flush()

	line, _ := r.ReadString('\n')
	parts := strings.Fields(line)

	a, _ := strconv.Atoi(parts[0])
	b, _ := strconv.Atoi(parts[1])
	year := 0

	for a <= b {
		a *= 3
		b *= 2
		year++
	}

	w.WriteString(strconv.Itoa(year))
	w.WriteByte('\n')
}
