package sieve

// start from 2.
func series() chan int {
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

func multiples(n int) chan int {
	out := make(chan int)
	go func() {
		s := series()
		for {
			i := <-s
			out <- i * n
		}
	}()
	return out
}

func sieve(n int, in chan int) chan int {
	out := make(chan int)
	s := multiples(n)

	go func() {
		for {
			next := <-s
			for {
				current := <-in
				if next != current {
					out <- current
				} else {
					break
				}
			}
		}
	}()
	return out
}

// Returns a stream of primes numbers
func PrimeSeries() chan int {
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
