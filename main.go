package main

import (
    "fmt"
)

func reverseString(s string) string {
    reversed := ""
    for _, char := range s {
        reversed = string(char) + reversed
    }
    return reversed
}

func main() {
    input := "hello"
    fmt.Println("Original:", input)
    fmt.Println("Reversed:", reverseString(input))
}
