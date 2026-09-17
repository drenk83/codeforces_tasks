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

	k, _ := strconv.Atoi(parts[0])
	n, _ := strconv.Atoi(parts[1])
	ww, _ := strconv.Atoi(parts[2])

	totalCost := ww * (k + (k * ww)) / 2
	amountToBorrow := totalCost - n

	if amountToBorrow < 0 {
		amountToBorrow = 0
	}

	w.WriteString(strconv.Itoa(amountToBorrow))
	w.WriteByte('\n')
}
