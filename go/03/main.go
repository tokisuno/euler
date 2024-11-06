package main

import (
	// "math"
	"fmt"
	"math"
)

func findFactors(n int) []int {
    var dom int = int(math.Sqrt(float64(n)))
    var result []int
    for i := 1; i <= dom; i++ {
        if n % i == 0 {
            // fmt.Printf("Factor: %v\n", i )
            result = append(result, i)
        }
    }
    return result
}

func isNotPrime(n int) bool {
    for i := 0; i < n; i++ {
        if len(findFactors(n)) == 2 {
            return true
        } else {
            return false
        }
    }
    return false
}

// I got my test cases backwards but hey, it works LMFAO
func main() {
    // ex := 13195
    ex := 600851475143
    var arr []int = findFactors(ex)
    // fmt.Println(arr)
    for i := 1; i < len(arr); i++ {
        if isNotPrime(arr[i]) == false {
            fmt.Println(arr[i])
        }
    }
}
