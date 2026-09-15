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
	m, _ := strconv.Atoi(parts[0])
	n, _ := strconv.Atoi(parts[1])

	area := m * n

	w.WriteString(strconv.Itoa(area / 2))
	w.WriteByte('\n')
}
