# DECISIONS.md

# TinyColor-Go Architectural Decisions

This document explains the intentional architectural decisions made while porting the original JavaScript TinyColor library to Go.

---

## 1. Package Structure

The original TinyColor project is written as a JavaScript module.

The Go implementation is organized as a reusable Go package, following standard Go project conventions. Public APIs are exposed through the `tinycolor` package.

---

## 2. Examples Organization

Originally, all example programs were located in a single directory.

Since Go cannot compile multiple `main()` functions within the same package, each example was moved into its own subdirectory under `examples/`.

This preserves functionality while following Go's package rules.

---

## 3. Chainable Manipulation API

Methods such as:

- Lighten
- Darken
- Saturate
- Desaturate
- Spin
- Brighten

modify the receiver in place and return the same pointer.

This matches the behavior of the original JavaScript implementation while remaining idiomatic in Go.

---

## 4. Behavioral Compatibility

The primary design goal was to preserve the observable behavior of the original TinyColor library while implementing it using idiomatic Go.

Parsing behavior, conversion algorithms, formatting rules, rounding behavior, readability calculations, and manipulation semantics were implemented to closely match the behavior of the original library.

---

## 5. Testing Strategy

The project includes:

- Unit tests
- Compatibility regression tests
- Differential testing against the JavaScript implementation
- GitHub Actions continuous integration

These validation methods were used throughout development to compare the Go implementation with the original JavaScript library and detect behavioral differences.

---

## 6. Error Handling

Where appropriate, Go conventions are followed while preserving TinyColor's permissive parsing behavior.

Invalid inputs are handled in a manner consistent with the original implementation whenever practical.

---

## 7. Continuous Integration

GitHub Actions is used to automatically run:

- gofmt verification
- go vet
- go test

on every push and pull request.

---

## Conclusion

This project adapts TinyColor to Go while preserving its behavior where practical and following standard Go project conventions for packaging, testing, and tooling.
