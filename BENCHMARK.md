# BENCHMARK.md

# TinyColor-Go Benchmark Report

## Environment

- Operating System: Windows 11
- CPU: 13th Gen Intel Core i7-1360P
- Go Version: 1.26.5
- Architecture: amd64

---

## Methodology

Benchmarks were executed using Go's built-in benchmarking framework.

Command used:

```bash
go test -bench=. -benchmem
```

Go's benchmarking framework automatically executes each benchmark multiple times to produce stable timing and allocation measurements.

---

## Results

| Benchmark                  | Performance |
| -------------------------- | ----------: |
| Parse                      | ~18.9 µs/op |
| RGB ↔ HSL / HSV Conversion |  ~172 ns/op |
| Formatting                 |  ~1.2 µs/op |

---

## Memory Usage

Representative benchmark results include memory allocation statistics reported by `go test -benchmem`.

Memory allocation statistics are included as reported by go test -benchmem to provide additional context for the benchmark results.

---

## Comparison with Original TinyColor

The original TinyColor library is implemented in JavaScript and runs inside a JavaScript runtime.

Direct performance comparisons are not meaningful because:

- Different programming languages
- Different runtimes
- Different garbage collectors
- Different optimization strategies

The benchmarks are intended to document the performance of this Go implementation rather than compare execution speed across languages.

---

## Conclusion

This benchmark report provides baseline performance measurements for the current implementation. The results can be used to track future performance changes as the project evolves.
