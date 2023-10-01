package main

import (
    "fmt"
)

func addNumbers(a, b int) int {
    return a + b
}

func subtractNumbers(a, b int) int {
    return a - b
}

func multiplyNumbers(a, b int) int {
    return a * b
}

func divideNumbers(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("Division by zero is not allowed")
    }
    return a / b, nil
}

func main() {
    result := addNumbers(5, 3)
    fmt.Printf("Addition result: %d\n", result)

    result = subtractNumbers(8, 2)
    fmt.Printf("Subtraction result: %d\n", result)

    result = multiplyNumbers(4, 7)
    fmt.Printf("Multiplication result: %d\n", result)

    result, err := divideNumbers(6, 0)
    if err != nil {
        fmt.Printf("Error: %s\n", err)
    } else {
        fmt.Printf("Division result: %d\n", result)
    }
}
