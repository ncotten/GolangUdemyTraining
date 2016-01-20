package main

import (
    "fmt"
)

func main() {
    //var small, large int
    var x, y, small, large int
    
    // Scan in the two numbers
    fmt.Print("Please enter two numbers: ")
    fmt.Scan(&x, &y)
    
    // Determine Largest and Smallest Value
    if x >= y {
        large = x
        small = y
    } else {
        large = y
        small = x
    }
    
    // Print out the Result
    fmt.Printf("Remainder of %v %% %v = %v\n", 
        large, small, (large % small))
}