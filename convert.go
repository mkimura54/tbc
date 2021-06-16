package main

import (
	"fmt"
	"strings"
	"strconv"
	"errors"
)

func convert(val string) (string, error) {
	if strings.Contains(val, ":") {
		tm := strings.Split(val, ":")
		if len(tm) != 3 {
			return "", errors.New("format error.")
		}

		h, err := strconv.Atoi(tm[0])
		if err != nil {
			return "", err
		}

		m, err := strconv.Atoi(tm[1])
		if err != nil {
			return "", err
		}

		s, err := strconv.Atoi(tm[2])
		if err != nil {
			return "", err
		}

		if m > 60 || s > 60 || h < 0 || m < 0 || s < 0 {
			return "", errors.New("param error.")
		}

		sec := 3600 * h + 60 * m + s
		return fmt.Sprintf("%d", sec), nil
	} else {
		i, err := strconv.Atoi(val)
		if err != nil {
			return "", err
		}
		if i < 0 {
			return "", errors.New("param error.")
		}


		h := i / 3600
		m := (i - 3600 * h) / 60
		s := i - 3600 * h - 60 * m
		return fmt.Sprintf("%01d:%02d:%02d", h, m, s), nil
	}
}
