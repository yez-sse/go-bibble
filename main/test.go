package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int64
	n = 123457890678
	var s = strconv.FormatInt(n, 10)
	fmt.Printf(s)
}
