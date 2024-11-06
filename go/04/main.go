/*************************************************************************
A palindromic number reads the same both ways.
The largest palindrome made from the product of two 2-digit numbers is:
   9009 = 91x99
Objective: Find the largest palindrome made from two 3-digit numbers

1. Find lower limit on n-digit numbers
2. Find upper limit on n-digit numbers

* This solution was taken from: https://www.youtube.com/watch?v=2ZGeZY7UZ6I
* This is modified only to separate one part into a function
* Studied the solution to learn it
*************************************************************************/

package main

import (
    "strconv"
    "fmt"
    "math"
)


func isPalindrome(x int, y int) bool {
    pal := false
    val := strconv.Itoa(x * y)
    for i := 0; i < len(val)/2; i++ {
        if val[i] == val[len(val) - i - 1] {
            pal = true
        } else {
            pal = false
            break
        }
    }
    return pal
}

func findPalindrome(n int) (int, int, int) {
    var largestPalindrome int

    var x1 int
    var x2 int

    var ulim int = int(math.Pow(10, float64(n)) - 1)
    var llim int = int(math.Pow(10, float64(n-1)) - 1)

    for x := ulim; x > llim; x-- {
        for y := ulim; y > llim; y-- {
            if(isPalindrome(x, y)) {
                palindrome := x * y
                if palindrome > largestPalindrome {
                    x1 = x
                    x2 = y
                    largestPalindrome = palindrome
                }
            }
        }
    }
    return x1, x2, largestPalindrome
}

func main() {
    fmt.Println(findPalindrome(3))

}
