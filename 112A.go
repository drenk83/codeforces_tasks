package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReaderSize(os.Stdin, 1<<20)
	w := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer w.Flush()

	str1, _ := r.ReadString('\n')
	str1 = strings.TrimSpace(str1)
	str1 = strings.ToLower(str1)

	str2, _ := r.ReadString('\n')
	str2 = strings.TrimSpace(str2)
	str2 = strings.ToLower(str2)

	if str1 < str2 {
		w.WriteString("-1")
	} else if str1 > str2 {
		w.WriteString("1")
	} else {
		w.WriteString("0")
	}
	w.WriteByte('\n')
}
