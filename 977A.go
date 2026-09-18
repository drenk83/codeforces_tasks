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

	n, _ := strconv.Atoi(parts[0])
	k, _ := strconv.Atoi(parts[1])

	for range k {
		if n%10 == 0 {
			n /= 10
		} else {
			n--
		}
	}

	w.WriteString(strconv.Itoa(n))
	w.WriteByte('\n')
}
