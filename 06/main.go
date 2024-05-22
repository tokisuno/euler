package main

import (
    "fmt"
    "math"
)

func sumOfSquares(n int) int {
    var sum float64

    for i := 1; i < n + 1; i++ {
        sum += math.Pow(float64(i), 2)
    }
    return int(sum)
}

func squareOfSum(n int) int {
    var sum float64
    i := 1
    for i = range n + 1{
        sum = sum + float64(i)
    }

    return int(math.Pow(sum, 2))
}

func main() {
    x := sumOfSquares(100)
    y := squareOfSum(100)

    fmt.Println(y-x)
}
