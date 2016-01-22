# Sieve of Eratosthenes In go.

Experiments in go concurrency. :)

http://www.mathbugsme.com/sieveoferatosthe.html

## Tests and Benchmark
```bash
go test -bench=. -v
=== RUN   TestPrimeSeries
--- PASS: TestPrimeSeries (0.00s)
=== RUN   TestPrimeSeriesTwo
--- PASS: TestPrimeSeriesTwo (0.00s)
PASS
186209088 ns/op
2.042s_/Users/dhananjay/Code/sieve ok  10 BenchmarkPrimeSeries-4
```
