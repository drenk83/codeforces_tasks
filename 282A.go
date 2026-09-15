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
	n, _ := strconv.Atoi(strings.TrimSpace(line))
	count := 0

	for range n {
		line, _ = r.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "X++" || line == "++X" {
			count++
		} else {
			count--
		}
	}
	w.WriteString(strconv.Itoa(count))
	w.WriteByte('\n')
}
