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

	number, _ := strconv.Atoi(line)
	happy := 0

	for number > 0 {
		tmp := number % 10
		if tmp == 4 || tmp == 7 {
			happy++
		}

		number /= 10
	}

	if happy == 4 || happy == 7 {
		w.WriteString("YES\n")
	} else {
		w.WriteString("NO\n")
	}
}
