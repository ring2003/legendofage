package main

import (
	"testing"
)

func TestCompare(t *testing.T) {
	legend := []PrinceAndPrincess{
		PrinceAndPrincess{Age{age: 20}, Age{age: 30}},
		PrinceAndPrincess{Age{age: 40}, Age{age: 30}},
		PrinceAndPrincess{Age{age: 30}, Age{age: 40}},
		PrinceAndPrincess{Age{age: 20}, Age{age: 30}},
	}
	right := false
	for _, l := range legend {
		for i := 1; i <= 100; i++ {
			for j := 100; j >= i; j-- {
				if l.Prince.Compareto(l.Princess, i, j) {
					right = true
					break
				}
			}
			if right {
				break
			}
		}
		if right {
			break
		}
	}
	if !right {
		t.Fail()
	}
}
func TestMain(t *testing.T) {
	TestCompare(t)
}
