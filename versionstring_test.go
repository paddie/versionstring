package versionstring

import (
	"testing"
)

type TestComp struct {
	in  string
	exp int
}

type TestParse struct {
	in  string
	exp []int
}

func TestParseVersionString(t *testing.T) {
	T := []TestParse{
		{"1.02.3.", []int{1, 2, 3}},
		{"1.2..3", []int{1, 2, 3}},
		{"1.00.2", []int{1, 0, 2}},
	}

	for i := 0; i < len(T); i++ {
		exp := T[i].exp
		act := ParseVersionString(T[i].in)

		if len(act) != len(exp) {
			t.Errorf("Exp: %d != %d Act: (err: %v)", exp, act)
		}
		for j := 0; j < len(act); j++ {
			if exp[j] != act[j] {
				t.Errorf("Exp %d != %d Act", exp, act)
			}

		}
	}
}

func TestCompareToString(t *testing.T) {

	lvs := "1.2.3"

	T := []TestComp{
		{"1.2.3", 0},
		{"1.1", -1},
		{"1.3", 1},
		{"1.2.4", 1},
		{"1.2.2", -1},
		{"1.2.2.2", -1},
		{"1.2.3.0.0", 0},
		{"1.2.3.2", 1},
		{"1.2.3.2.4", 1},
		//_____|___
	}

	for i := 0; i < len(T); i++ {
		exp := T[i].exp
		if act := CompareStrings(lvs, T[i].in); act != exp {
			t.Errorf("Exp %d != %d Act", exp, act)
		}
	}
}