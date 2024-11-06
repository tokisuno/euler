// * 2520 is the smallest number that can
//  be divided by each of the numbers from
//    1-10 without any remainder

// * What's the smallest positive number
//  that is "evenly divisible" by all of the
//    numbers from 1-20?

package main

import "fmt"


func main() {
    amnt := lcm(20)

    fmt.Println(amnt)
}

// find greatest common divisor
func gcd(ans int, i int) int {
    if ans%i != 0 {
        return gcd(i, ans%i)
    } else {
        return i
    }
}
// find lowest common multiple
func lcm(n int) int {
    ans := 1
    var i int

    for i = 1; i <= n; i++ {
        ans = (ans * i)/(gcd(ans, i))
    }
    return ans
}

