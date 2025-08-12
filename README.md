# Go Learning Projects

This repository contains small Go projects for practicing Go programming fundamentals. The projects are based on the official [Go Tutorial: Getting started](https://go.dev/doc/tutorial/getting-started).

## Projects

### 1. Hello (`hello/`)
A simple "Hello, World!" application that demonstrates:
- Basic Go module structure
- Importing and using local packages
- Error handling with logging
- Working with slices and maps
- Random selection from data

**Features:**
- Greets multiple people with randomized greeting messages
- Includes error handling for empty names
- Uses the local `greetings` package

**To run:**
```bash
cd hello
go run .
```

### 2. Greetings (`greetings/`)
A reusable greeting package that provides greeting functionality:
- `Hello(name string)` - Returns a personalized greeting for one person
- `Hellos(names []string)` - Returns personalized greetings for multiple people
- Random greeting message selection
- Input validation and error handling

**Features:**
- Multiple greeting message formats
- Error handling for empty names
- Support for both single and multiple greetings
- Comprehensive test coverage

**To test:**
```bash
cd greetings
go test
```

**To test with verbose output:**
```bash
cd greetings
go test -v
```

## What You'll Learn

These projects cover fundamental Go concepts including:
- **Modules and packages**: Creating and importing Go modules
- **Functions**: Writing functions with multiple return values
- **Error handling**: Go's explicit error handling approach
- **Variables and types**: Working with strings, slices, and maps
- **Control flow**: Using for loops and range
- **Testing**: Writing unit tests with the `testing` package
- **Random numbers**: Using the `math/rand` package

## Getting Started

1. Make sure you have Go installed (version 1.16 or later)
2. Clone this repository
3. Navigate to either project directory
4. Run the code using `go run .` or run tests using `go test`

## Project Structure

```
go-learning/
├── README.md
├── hello/
│   ├── go.mod
│   └── hello.go
└── greetings/
    ├── go.mod
    ├── greetings.go
    └── greetings_test.go
```

## Resources

- [Go Tutorial: Getting started](https://go.dev/doc/tutorial/getting-started)
- [Go Documentation](https://go.dev/doc/)
- [Go Language Tour](https://go.dev/tour)
- [Go by Example](https://gobyexample.com/)

---

*These projects follow the official Go tutorial and serve as a foundation for learning Go programming concepts.*
