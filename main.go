package main

import (
	"os"
	"fmt"
)

func main() {
	l := len(os.Args)
	if l == 2 {
		result, err := convert(os.Args[1])
		if err != nil {
			fmt.Printf("%s\n", err)
		} else {
			fmt.Println("=> " + result)
		}
	} else if l == 3 {
		result, err := subtraction(os.Args[1], os.Args[2])
		if err != nil {
			fmt.Printf("%s\n", err)
		} else {
			fmt.Println("=> " + result)
		}
	} else {
		fmt.Println("param count error.")
	}
}
