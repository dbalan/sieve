# Sieve of Eratosthe In go.

Experiments in go concurrency. :)

http://www.mathbugsme.com/sieveoferatosthe.html

## example output
```bash
# generates an infinte stream
go run sieve.go |head -n 20
2
3
5
7
9
11
13
15
17
19
21
23
25
27
29
31
33
35
37
39
signal: broken pipe
```

## Benchmark
```bash
go test -bench=.
testing: warning: no tests to run
PASS
BenchmarkPrime-4	      10	 224442312 ns/op
ok  	_/Users/dhananjay/Code/sieve	2.409s
```
