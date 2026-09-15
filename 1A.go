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
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])
	a, _ := strconv.Atoi(parts[2])

	rows := (n + a - 1) / a
	cols := (m + a - 1) / a

	w.WriteString(strconv.Itoa(rows * cols))
	w.WriteByte('\n')
}
