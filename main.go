package main

import (
	"context"
	"fmt"
	"math/rand"
)

type Age struct {
	age int
}

type Comparable interface {
	AgeBefore(n int) int
	AgeAfter(m int) int
	AgeCurrent() int
	Compareto(another Comparable, n int, m int) bool
}

func (this *Age) Compareto(ref Age, n int, m int) bool {
	if ref.AgeBefore(n) == (this.AgeCurrent()+ref.AgeCurrent())/2 {
		if this.AgeBefore(n)*2 == ref.AgeAfter(m) {
			if this.AgeAfter(m) == ref.AgeCurrent() {
				return true
			}
		}
	}
	return false
}

func (this *Age) AgeBefore(n int) int {
	return this.age - n
}

func (this *Age) AgeAfter(m int) int {
	return this.age + m
}

func (this *Age) AgeCurrent() int {
	return this.age
}

type Years struct {
	n int
	m int
}
type PrinceAndPrincess struct {
	Prince   Age
	Princess Age
}

func randGen(ctx context.Context) <-chan Years {
	ch := make(chan Years, 1)
	go func() {
		flag := false
		for !flag {
			select {
			case <-ctx.Done():
				close(ch)
				flag = true
				break
			default:
				var n int = rand.Intn(100)
				var m int = rand.Intn(100)
				if n < m {
					ch <- Years{n, m}
				}
			}
		}
	}()
	return ch
}

// Barduls'gate II chapter 3
func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	randgen := randGen(ctx)

	var found bool = false
	for !found {
		select {
		case year := <-randgen:
			go func() {
				for i := 1; i <= 100; i++ {
					for j := 100; j > i; j-- {
						legend := PrinceAndPrincess{Prince: Age{age: i}, Princess: Age{age: j}}
						if legend.Prince.Compareto(legend.Princess, year.n, year.m) {
							fmt.Printf("Prince age: %v, Princess age: %v, YearN: %d, YearM: %d\n",
								legend.Prince.age,
								legend.Princess.age,
								year.n,
								year.m)
							cancel()
						}
					}
				}
			}()
		case <-ctx.Done():
			found = true
		}
	}
}
