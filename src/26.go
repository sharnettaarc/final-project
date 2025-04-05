package main

import (
    "fmt"
    "math"
)

func main() {
    var x float64 = 2.0
    if math.Mod(x, 1) == 0 {
        fmt.Println("x is divisible by 1")
    } else {
        fmt.Println("x is not divisible by 1")
    }
}
