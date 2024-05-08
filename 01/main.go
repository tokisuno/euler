package main

import "fmt"

func arraySum(arr []int) int {
    var sum int;
    for i := 0; i < len(arr); i++ {
        sum += arr[i]
    }

    return sum
}

func main() {
    of3 := []int{}
    of5 := []int{}

    for i := range 1000 {
        if i % 3 == 0 {
            of3 = append(of3, i)
        }
    }
    sumof3 := arraySum(of3)

    for j := range 1000 {
        if j % 5 == 0 {
            if j % 3 == 0 {
                continue
            }else{
                of5 = append(of5, j)
            }
        }
    }
    sumof5 := arraySum(of5)

    fmt.Printf("Answer: %v", sumof3 + sumof5)

}
