package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	PROGRAM_NAME = "tbc"
	VERSION      = "v0.4"
)

func main() {
	l := len(os.Args)
	if l == 1 {
		fmt.Println(PROGRAM_NAME + " " + VERSION)
		for {
			sc := bufio.NewScanner(os.Stdin)
			for sc.Scan() {
				input := sc.Text()
				if input == "quit" || input == "exit" || input == "q" || input == "" {
					return
				} else {
					param := strings.Split(input, " ")
					count := len(param)
					if count == 1 {
						execConvert(param[0])
					} else if count == 2 {
						execSubtraction(param[0], param[1])
					}
				}
			}
		}
	} else if l == 2 {
		execConvert(os.Args[1])
	} else if l == 3 {
		execSubtraction(os.Args[1], os.Args[2])
	} else {
		fmt.Println("param count error.")
	}
}

func execConvert(param string) {
	result, err := convert(param)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Println(result)
	}
}

func execSubtraction(param1, param2 string) {
	result, err := subtraction(param1, param2)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Println(result)
	}
}
