package main

import (
	"testing"
)

func TestAppName(t * testing.T) {
	expect := "This is the Rodentia Application!!"
	actual := AppName()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}