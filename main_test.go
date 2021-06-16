package main

import (
	"testing"
)

////////////////////
// xxxx -> x:xx:xx
////////////////////

func TestConvertA01(t *testing.T) {
	result, err := convert("5025")
	if err != nil {
		t.Fatalf("failed %#v", err)
	}
	if result != "1:23:45" {
		t.Fatalf("failed %s", result)
	}
}

func TestConvertA02(t *testing.T) {
	result, err := convert("4025")
	if err != nil {
		t.Fatalf("failed %#v", err)
	}
	if result != "1:07:05" {
		t.Fatalf("failed %s", result)
	}
}

func TestConvertA03(t *testing.T) {
	result, err := convert("2025")
	if err != nil {
		t.Fatalf("failed %#v", err)
	}
	if result != "0:33:45" {
		t.Fatalf("failed %s", result)
	}
}

func TestConvertA04(t *testing.T) {
	result, err := convert("0")
	if err != nil {
		t.Fatalf("failed %#v", err)
	}
	if result != "0:00:00" {
		t.Fatalf("failed %s", result)
	}
}

func TestConvertA05(t *testing.T) {
	result, err := convert("")
	if err == nil || result != "" {
		t.Fatalf("failed")
	}
}

func TestConvertA06(t *testing.T) {
	result, err := convert("-100")
	if err == nil || result != "" {
		t.Fatalf("failed")
	}
}

////////////////////
// x:xx:xx -> xxxx
////////////////////

func TestConvertB01(t *testing.T) {
	result, err := convert("1:23:45")
	if err != nil {
		t.Fatalf("failed %#v", err)
	}
	if result != "5025" {
		t.Fatalf("failed %s", result)
	}
}

func TestConvertB02(t *testing.T) {
	result, err := convert("1:07:05")
	if err != nil {
		t.Fatalf("failed %#v", err)
	}
	if result != "4025" {
		t.Fatalf("failed %s", result)
	}
}

func TestConvertB03(t *testing.T) {
	result, err := convert("0:33:45")
	if err != nil {
		t.Fatalf("failed %#v", err)
	}
	if result != "2025" {
		t.Fatalf("failed %s", result)
	}
}

func TestConvertB04(t *testing.T) {
	result, err := convert("1:23:45:45")
	if err == nil || result != "" {
		t.Fatalf("failed")
	}
}

func TestConvertB05(t *testing.T) {
	result, err := convert("1:23")
	if err == nil || result != "" {
		t.Fatalf("failed")
	}
}

func TestConvertB06(t *testing.T) {
	result, err := convert(":")
	if err == nil || result != "" {
		t.Fatalf("failed")
	}
}

func TestConvertB07(t *testing.T) {
	result, err := convert("1:24:68")
	if err == nil || result != "" {
		t.Fatalf("failed")
	}
}

func TestConvertB08(t *testing.T) {
	result, err := convert("-1:23:45")
	if err == nil || result != "" {
		t.Fatalf("failed")
	}
}

func TestSubtraction01(t *testing.T) {
	result, err := subtraction("9296", "5025")
	if err != nil {
		t.Fatalf("failed %#v", err)
	}
	if result != "1:11:11 (4271)" {
		t.Fatalf("failed %s", result)
	}
}

func TestSubtraction02(t *testing.T) {
	result, err := subtraction("2:34:56", "1:23:45")
	if err != nil {
		t.Fatalf("failed %#v", err)
	}
	if result != "1:11:11 (4271)" {
		t.Fatalf("failed %s", result)
	}
}

func TestSubtraction03(t *testing.T) {
	result, err := subtraction("9296", "1:23:45")
	if err != nil {
		t.Fatalf("failed %#v", err)
	}
	if result != "1:11:11 (4271)" {
		t.Fatalf("failed %s", result)
	}
}

func TestSubtraction04(t *testing.T) {
	result, err := subtraction("9296", "")
	if err == nil || result != "" {
		t.Fatalf("failed")
	}
}

func TestSubtraction05(t *testing.T) {
	result, err := subtraction("5025", "9296")
	if err != nil {
		t.Fatalf("failed %#v", err)
	}
	if result != "1:11:11 (4271)" {
		t.Fatalf("failed %s", result)
	}
}
