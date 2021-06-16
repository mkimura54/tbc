package main

import (
	"os"
	"fmt"
)

func main() {
	l := len(os.Args)
	if l == 2 {
		result, _ := convert(os.Args[1])
		fmt.Println("=> " + result)
	} else if l == 3 {
		subtraction(os.Args[1], os.Args[2])
	} else {
		fmt.Println("param count error.")
	}
}

func subtraction(val1, val2 string) (string, error) {
	return val1 + "," + val2, nil
}
