package main

import "testing"

func TestCheckEven(t *testing.T) {
	expect := true
	actual := checkEven(10)
	if expect != actual {
		t.Errorf(`expect="%t actual="%t"`, expect, actual)
	}
}
