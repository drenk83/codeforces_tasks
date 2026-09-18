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

	n, _ := strconv.Atoi(line)

	x, y, z := 0, 0, 0

	for range n {
		line, _ = r.ReadString('\n')
		parts := strings.Fields(line)

		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		c, _ := strconv.Atoi(parts[2])

		x += a
		y += b
		z += c
	}

	if x == 0 && y == 0 && z == 0 {
		w.WriteString("YES\n")
	} else {
		w.WriteString("NO\n")
	}
}
