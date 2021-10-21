package main

/*
**	release 0.1
 */

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	cnt := readInput()
	fmt.Println(cnt)
}

func readInput() (res int) {
	flag.Parse()
	str := strings.Join(flag.Args(), " ")
	if str == "" {
		return 0
	}

	res = len(strings.Split(str, " "))
	return res
}
