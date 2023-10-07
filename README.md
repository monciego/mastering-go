# Mastering Go

Welcome to my GoLang learning journey! Below are my notes to help you understand the fundamentals of GoLang.

## Getting Started

```bash
go mod init mastering-go
```

- Initializes a new Go module.
- The module path can correspond to a repository where you plan to publish your module (e.g., github.com/monciego/mastering-go).
- All Go code must belong to a package.
- The main function serves as the entry point of a Go program.
- A program can only have one main function as it represents the single entry point.

## Go Packages

- Go programs are organized into packages.
- Go's standard library provides various core packages for developers to use.
- For instance, the "fmt" package is one of these core packages, which you can utilize by importing it.

## Variables and Constants

- Variables serve as containers for storing values, allowing you to reference the stored values throughout your application.
- Constants are similar to variables, but their values cannot be changed once set.
