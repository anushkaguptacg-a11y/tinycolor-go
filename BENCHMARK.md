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

Each benchmark was executed multiple times by Go's benchmark runner until statistically stable results were obtained.

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

The implementation is designed to minimize allocations while preserving behavioral compatibility with the original JavaScript implementation.

---

## Comparison with Original TinyColor

The original TinyColor library is implemented in JavaScript and runs inside a JavaScript runtime.

Direct performance comparisons are not meaningful because:

- Different programming languages
- Different runtimes
- Different garbage collectors
- Different optimization strategies

Instead, this project focuses on preserving behavioral compatibility while providing efficient native Go performance.

---

## Conclusion

TinyColor-Go provides a fast, idiomatic Go implementation with comprehensive testing, low allocation overhead, and behavior closely matching the original TinyColor JavaScript library.
