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
	line = strings.TrimSpace(line)

	x, _ := strconv.Atoi(line)

	out := (x + 4) / 5

	w.WriteString(strconv.Itoa(out))
	w.WriteByte('\n')
}
