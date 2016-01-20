package main

import (
    "fmt"
)

func main() {
    
    // Print out all even numbers from 0 - 100
    for i := 0; i <= 100; i++ {
        if i % 2 == 0 && i != 0 {
            fmt.Printf("%v \t", i)
            
            if i % 10 == 0 {
                fmt.Println()
            }
        }
    }
    
}