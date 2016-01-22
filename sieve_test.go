package sieve

import (
	"testing"
)

func TestPrimeSeries(t *testing.T) {
	ps := []int{2, 3, 5, 7, 11, 13, 17, 19, 23}
	gen := PrimeSeries()
	for _, v := range ps {
		out := <-gen
		if out != v {
			t.Error("Expected: ", v, " Got: ", out)
		}
	}
}
func BenchmarkPrimeSeries(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// generate first 1000 primes
		p := PrimeSeries()
		for j := 0; j < 1000; j++ {
			<-p
		}
	}
}
