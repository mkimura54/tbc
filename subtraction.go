package main

import (
	"fmt"
	"strings"
	"errors"
	"strconv"
)

func subtraction(val1, val2 string) (string, error) {
	var tm1, tm2 int
	if strings.Contains(val1, ":") {
		v, err := convertToSeconds(val1)
		if err != nil {
			return "", err
		}
		tm1 = v
	} else {
		v, err := strconv.Atoi(val1)
		if err != nil {
			return "", err
		}
		if tm1 < 0 {
			return "", errors.New("param error.")
		}
		tm1 = v
	}
	if strings.Contains(val2, ":") {
		v, err := convertToSeconds(val2)
		if err != nil {
			return "", err
		}
		tm2 = v
	} else {
		v, err := strconv.Atoi(val2)
		if err != nil {
			return "", err
		}
		if tm2 < 0 {
			return "", errors.New("param error.")
		}
		tm2 = v
	}

	var s int
	if tm1 > tm2 {
		s = tm1 - tm2
	} else {
		s = tm2 - tm1
	}
	return fmt.Sprintf("%s (%d)", convertToString(s), s), nil
}
