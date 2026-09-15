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
	k, _ := strconv.Atoi(parts[1])
	count := 0

	line, _ = r.ReadString('\n')
	parts = strings.Fields(line)
	point, _ := strconv.Atoi(parts[k-1])

	for _, part := range parts {
		num, _ := strconv.Atoi(part)
		if num >= point && num != 0 {
			count++
		}
	}
	w.WriteString(strconv.Itoa(count))
	w.WriteByte('\n')
}
