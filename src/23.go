package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    var data = []int{
        1, 2, 3,
        -5, 0, 4,
        8, -1, 7
    }

    fmt.Println("Random data:", data)
}
