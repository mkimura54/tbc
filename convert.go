package main

import (
	"fmt"
	"strings"
	"strconv"
	"errors"
)

func convert(val string) (string, error) {
	if strings.Contains(val, ":") {
		v, err := convertToSeconds(val)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%d", v), nil
	} else {
		i, err := strconv.Atoi(val)
		if err != nil {
			return "", err
		}
		if i < 0 {
			return "", errors.New("param error.")
		}

		return convertToString(i), nil
	}
}
