package rodentia

import (
	"testing"
)

func TestGopherLike(t *testing.T) {
	expect := "ʕ ◔ϖ◔ʔ < We like tunnels!"
	actual := GopherLike()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
func TestBeaverLike(t *testing.T) {
	expect := "ʕ ˙ϖ˙ʔ < We like dams!"
	actual := BeaverLike()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

func TestHamsterLike(t *testing.T) {
	expect := "ʕ òﻌóʔ < We like wheels!"
	actual := HamsterLike()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}