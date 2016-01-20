package main

import "fmt"

func main() {
    /*
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, " - ", j)
		}
	}
    */
    
    fmt.Println("\nMultiplication Table")
    for i := 0; i <= 5; i++ { 
        for j := 0; j <= 5; j++ {
            fmt.Printf("%v\t", (i*j))
        }
        fmt.Println()
    }
    fmt.Println()
}
