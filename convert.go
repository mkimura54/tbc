package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func convert(val string) (string, error) {
	if strings.Contains(val, ":") {
		v, err := convertToSeconds(val)
		if err != nil {
			return "", err
		}

		f := float64(v) / 3600
		return fmt.Sprintf("%d (%.5f)", v, f), nil
	} else {
		i, err := strconv.Atoi(val)
		if err != nil {
			return "", err
		}
		if i < 0 {
			return "", errors.New("param error.")
		}

		f := float64(i) / 3600
		return fmt.Sprintf("%s (%.5f)", convertToString(i), f), nil
	}
}
