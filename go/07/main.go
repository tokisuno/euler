package main

import "fmt"

func isPrime(num int) bool {
    if num <= 1 {
        return false
    }

    if num == 2 || num == 3 {
        return true
    }

    if num % 2 == 0 || num % 3 == 0 {
        return false
    }

    // SLOW
    // for i := 2; i < num; i++ {
    //     if num % i == 0 {
    //         return false
    //     }
    // }

    // FAST
    //  6k-1                  6k+1
    for i := 5; i * i <= num; i += 6 {
        if num % i == 0 || num % (i + 2) == 0 {
            return false
        }
    }

    return true
}

func nThPrime(n int) int {
    i := 2

    for n>0 {
        if isPrime(i) {
            n--
        }
        i++
    }
    i-=1

    return i
}

func main() {
    input := 10001
    fmt.Printf("Prime at %v is %v", input, nThPrime(input))
}
