package main

import (
	"testing"
)

func BenchmarkPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// generate first 1000 primes
		p := primeSeries()
		for j := 0; j < 1000; j++ {
			<-p
		}
	}
}
