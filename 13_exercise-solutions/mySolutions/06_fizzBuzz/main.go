package main

import (
    "fmt"
)

func main() {
    var num int
    fmt.Print("Please enter a number: ")
    fmt.Scan(&num)
    
    if num % 15 == 0 {
        fmt.Println("FizzBuzz")
    } else if num % 3 == 0 {
        fmt.Println("Fizz")
    } else if num % 5 == 0 {
        fmt.Println("Buzz")
    } else {
        fmt.Println("Try Again")
    }
}