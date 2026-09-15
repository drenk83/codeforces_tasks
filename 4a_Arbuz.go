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
	weight, _ := strconv.Atoi(strings.TrimSpace(line))

	if weight%2 == 0 && weight > 2 {
		w.WriteString("YES\n")
	} else {
		w.WriteString("NO\n")
	}
}
