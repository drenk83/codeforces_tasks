package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 1<<20)
	w := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer w.Flush()

	for i := 0; i < 5; i++ {
		line, _ := r.ReadString('\n')
		parts := strings.Fields(line)

	LOOP:
		for j, part := range parts {
			num, _ := strconv.Atoi(part)
			if num == 1 {
				out := abs(j-2) + abs(i-2)
				w.WriteString(strconv.Itoa(out))
				w.WriteByte('\n')
				break LOOP
			}
		}
	}
}
