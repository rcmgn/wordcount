package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	cnt := readInput()
	fmt.Println(cnt)
}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() (res int) {
	flag.Parse()
	// fmt.Println(flag.Args())
	str := strings.Join(flag.Args(), " ")
	// fmt.Println(str)
	res = len(strings.Split(str, " "))
	return res
}
