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

Shorthand syntax
`:=` is a shorthand syntax for declaring and initializing variables. It's a part of the short variable declaration syntax. With `:=`, you can declare a variable and assign a value to it, and Go will automatically infer the data type of the variable based on the assigned value.

```go
  name := "John"      // inferred as string
  age := 30           // inferred as int
  pi := 3.14159       // inferred as float64
  isActive := true    // inferred as bool
```

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

## Data Types

GoLang is statically typed, requiring explicit data type declaration for variables:

```golang
  var age int // Explicitly declaring an 'int' variable 'age'
```

BUT GoLang also supports type inference, allowing the compiler to deduce types from values:

```golang
  name := "John" // Type inferred as 'string' based on the assigned value
```

%T is a formatting verb used with the fmt.Printf function to print the type of a variable.

```golang
    var num int = 42
    var name string = "John"
    var pi float64 = 3.14159

    fmt.Printf("Type of num: %T\n", num)   // Output: Type of num: int
    fmt.Printf("Type of name: %T\n", name) // Output: Type of name: string
    fmt.Printf("Type of pi: %T\n", pi)     // Output: Type of pi: float64

```

Basic Data Types:

- Integers: int, int8, int16, int32, int64
  Integers are used to store whole numbers without any decimal points. They can be either positive or negative.

  ```go
    var age int = 25
  ```

- Unsigned Integers:
  `uint` is a built-in numeric data type representing unsigned integers. "Unsigned" means that these integers can only hold non-negative values (zero and positive integers) and cannot store negative values.

  ```go
     var positiveNumber uint = 42
    var anotherNumber uint = 100

    fmt.Println("Positive Number:", positiveNumber) // Output: Positive Number: 42
    fmt.Println("Another Number:", anotherNumber)   // Output: Another Number: 100

    // Unsigned integers cannot hold negative values.
    // var negativeNumber uint = -10 // This will result in a compilation error.
  ```

- Floating-Point Numbers: float32, float64
  Floating-point numbers are used to store numbers with decimal points. GoLang provides two types: float32 and float64 for single and double-precision floating-point numbers, respectively.

  ```go
    var price float64 = 29.99
  ```

- Floating-Point Numbers: float32, float64
  Floating-point numbers are used to store numbers with decimal points. GoLang provides two types: float32 and float64 for single and double-precision floating-point numbers, respectively.

  ```go
    var price float64 = 29.99
  ```

- Booleans: bool
  Booleans represent true or false values. They are often used in conditions and control structures to make decisions in the program.

  ```go
    var isAdult bool = true
  ```

- Characters: byte (alias for uint8), rune (alias for int32)
  Characters, represented by the rune type, are used to store single Unicode characters. In GoLang, characters are essentially aliases for int32 values representing Unicode code points.
  ```go
    var grade rune = 'A'
  ```
- Strings:
  Strings are sequences of characters. They are used to represent textual data. Strings in GoLang are immutable, meaning once created, their values cannot be changed.
  ```go
    var message string = "Hello, World!"
  ```
