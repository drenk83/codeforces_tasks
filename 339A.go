package main

import (
	"bufio"
	"os"
	"slices"
	"strings"
)

func main() {
	r := bufio.NewReaderSize(os.Stdin, 1<<20)
	w := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer w.Flush()

	line, _ := r.ReadString('\n')
	line = strings.TrimSpace(line)

	arr := make([]int, 0, len(line))

	for _, char := range line {
		if char != '+' {
			num := int(char - '0')
			arr = append(arr, num)
		}
	}

	slices.Sort(arr)
	for i, num := range arr {
		if i > 0 {
			w.WriteByte('+')
		}
		w.WriteString(string(rune(num + '0')))
	}
	w.WriteByte('\n')
}
