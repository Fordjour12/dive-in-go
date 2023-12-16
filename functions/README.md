# Functions in Go

## Function Declaration

```go
// Function without parameters and return value
func greet() {
    fmt.Println("Hello, World!")
}

// Function with parameters and return value
func add(a, b int) int {
    return a + b
}
```

## Function Parameters

- Functions can take zero or more parameters.
- Parameters specify the types of values that the function expects when called.
- Parameters are enclosed in parentheses after the function name.

```go
func greet(name string) {
    fmt.Println("Hello, " + name + "!")
}
```

## Return Values

- Functions can return zero or more values.
- Return values are specified after the function parameters.
- If a function has a return type, it must include a `return` statement.

```go
func add(a, b int) int {
    return a + b
}
```

## Multiple Return Values

- Go allows functions to return multiple values.

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

## Named Return Values

- You can name return values, which creates named variables in the function body.
- Useful for documenting the intent of the returned values.

```go
func divide(a, b int) (result int, err error) {
    if b == 0 {
        err = errors.New("division by zero")
        return
    }
    result = a / b
    return
}
```

## Variadic Functions

- Functions can take a variable number of arguments using the `...` syntax.
- The arguments are treated as a slice inside the function.

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}
```

## Function as a First-Class Citizen

- In Go, functions are first-class citizens, meaning they can be assigned to variables, passed as arguments, and returned from other functions.

```go
func multiply(a, b int) int {
    return a * b
}

func main() {
    operation := multiply
    result := operation(3, 4)
    fmt.Println(result) // Output: 12
}
```

## Anonymous Functions (Closures)

- You can declare functions without a name, creating anonymous functions.
- Anonymous functions can form closures by capturing variables from the surrounding context.

```go
func main() {
    add := func(a, b int) int {
        return a + b
    }

    result := add(2, 3)
    fmt.Println(result) // Output: 5
}
```

## Defer, Panic, and Recover

- `defer` schedules a function call to be run after the function completes but before it returns.
- `panic` stops the normal execution of a function and begins panicking.
- `recover` stops the panic and returns the value passed to the `panic`.

```go
func example() {
    defer fmt.Println("This will be executed last")

    // Code that may panic
    panic("Something went wrong")

    fmt.Println("This will not be executed")
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    example()
}
```

## Function Composition

- Functions can be composed by calling one function from another.

```go
func double(x int) int {
    return x * 2
}

func square(x int) int {
    return x * x
}

func doubleAndSquare(x int) int {
    return square(double(x))
}
```

## Key Points

- Functions in Go are defined using the `func` keyword.
- They can take parameters, return values, and be assigned to variables.
- Variadic functions allow a variable number of arguments.
- Anonymous functions and closures provide flexibility.
- `defer`, `panic`, and `recover` handle exceptional cases.

Understanding these concepts will help you effectively use functions to structure your Go programs.