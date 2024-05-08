package main

import "fmt"

func arraySum(arr []int32) int32 {
    var sum int32
    for i := 0; i < len(arr); i++ {
        sum += arr[i]
    }
    return sum
}

func genFib() []int32 {
    var arr []int32
    var a int32 = 0
    var b int32 = 1
    var c int32

    for i := 2; i <= 1000; i++ {
        c = a + b
        a = b
        b = c
        if b >= 4e6 {
            break
        }
        if b % 2 == 0 {
            arr = append(arr, b)
        }
    }
    return arr
}

func main() {
    fmt.Println(arraySum(genFib()))
}
