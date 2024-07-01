package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func convertToString(seconds int) string {
	h := seconds / 3600
	m := (seconds - 3600*h) / 60
	s := seconds - 3600*h - 60*m
	return fmt.Sprintf("%01d:%02d:%02d", h, m, s)
}

func convertToSeconds(val string) (int, error) {
	tm := strings.Split(val, ":")
	if len(tm) != 3 {
		return 0, errors.New("format error.")
	}

	h, err := strconv.Atoi(tm[0])
	if err != nil {
		return 0, err
	}

	m, err := strconv.Atoi(tm[1])
	if err != nil {
		return 0, err
	}

	s, err := strconv.Atoi(tm[2])
	if err != nil {
		return 0, err
	}

	if m > 60 || s > 60 || h < 0 || m < 0 || s < 0 {
		return 0, errors.New("param error.")
	}

	return 3600*h + 60*m + s, nil
}
