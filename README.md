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

## Formatting Output

1. Printing Variables:

```go
	var userName = "monciego"
    fmt.Printf("Hello %s\n", userName)
```

- Use the %s verb to print strings.
- %s is a placeholder for string variables.

1. Formatting Strings:

```go
	var userName = "monciego"
    fmt.Printf("Hello %s\n", userName) // Hello monciego
```

- Use the %s verb to print strings.
- %s is a placeholder for string variables.

2. Formatting Numbers:

```go
    num := 42
    fmt.Printf("Decimal: %d, Binary: %b, Hexadecimal: %x\n", num, num, num) // Decimal: 42, Binary: 101010, Hexadecimal: 2a
```

- Use verbs like %d (decimal), %b (binary), and %x (hexadecimal) to format and print integers.
- You can use multiple verbs in a single Printf statement.

3. Floating Point Numbers:

```go
    pi := 3.14159
    fmt.Printf("Pi: %.2f\n", pi) // Pi: 3.14
```

- Use %f to format and print floating-point numbers.
- You can specify the precision with %.2f to round to two decimal places.

4. Boolean Formatting:

```go
    isTrue := true
    fmt.Printf("Is true? %t\n", isTrue) // Is true? true
```

- Use %t to format and print boolean values.

5. Padding Numbers:

```go
    num := 7
    fmt.Printf("Padded: %05d\n", num)  // Output: Padded: 00007

```

6. Width and Precision:

```go
    num := 123.456789
    fmt.Printf("Width: %10.2f, Precision: %.2f\n", num, num) // Width:     123.46, Precision: 123.46
```

- Specify the width and precision of floating-point numbers with %w.pf.
- %10.2f means a width of 10 characters with 2 decimal places of precision.

7. Printing Structs:

```go
    type Person struct {
        Name string
        Age  int
    }
    person := Person{Name: "John", Age: 30}
    fmt.Printf("Person: %+v\n", person)  // Person: {Name:John Age:30}
```

- You can print struct values using %+v to include field names.
- Useful for debugging and displaying structured data.
