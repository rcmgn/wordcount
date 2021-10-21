package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	cnt, err := readInput()
	if err != nil {
		fail(err)
	}
	fmt.Println(cnt)
}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() (res int, err error) {
	flag.Parse()
	// fmt.Println(flag.Args())
	str := strings.Join(flag.Args(), " ")
	// fmt.Println(str)
	res = len(strings.Split(str, " "))
	return res, nil
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("wordcount:", err)
	os.Exit(1)
}
