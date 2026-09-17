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
	line, _ = r.ReadString('\n')
	line = strings.TrimSpace(line)

	runes := []rune(line)
	count := 0

	for i, _ := range runes {
		if i+1 < len(runes) && runes[i] == runes[i+1] {
			count++
		}
	}

	w.WriteString(strconv.Itoa(count))
	w.WriteByte('\n')
}
