package sieve

// start from 2.
func series() <-chan int {
	out := make(chan int)
	go func() {
		i := 1
		for {
			i++
			out <- i
		}
	}()
	return out
}

func sieve(n int, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for {
			current := <-in
			if current%n != 0 {
				out <- current
			}
		}
	}()
	return out
}

// Returns a stream of primes numbers
func PrimeSeries() <-chan int {
	out := series()
	p := make(chan int)
	go func() {
		for {
			prime := <-out
			p <- prime
			out = sieve(prime, out)
		}
	}()
	return p
}
