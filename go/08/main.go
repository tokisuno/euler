// I didn't make this solution on my own.
// https://stackoverflow.com/questions/26893381/porting-project-euler-8-to-go
// I wrote comments in here in order to figure out the logic of the code.
// --------------------------------------------------------------------------
// I finally understand how slices work, and how they're just low:up bounds.
// Indexing into slices has been hell for me so far.
// Seeing them as slices instead of "go arrays" has helped a lot!

package main

import (
	"fmt"
	"strconv"
	"strings"
)

var numbers string = "7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450"

func string2arr(data string) []int64 {
    var grid []int64
    stringGrid := strings.Split(data, "")
    for i := 0; i < 1000; i++ {
        cell, _ := strconv.Atoi(stringGrid[i])
        grid = append(grid, int64(cell))
    }
    return grid
}

func largestSum(amount int) int64 {
    var grid []int64 = string2arr(numbers)
    var largest int64
    for i := 0; i < len(grid)-(amount-1); i++ {
        // *slices calculate lower:upper bound
        // therefore 0: 0 + 13 gives the range of ints within 0-13
        //          when looping, it becomes 1:14, 2:15, and so on
        // a is a slice containing this range
        a := grid[i : i+amount]
        var total int64 = 1
        // b calculates the total of the grid slice
        for b := 0; b < len(a); b++ {
            total *= a[b]
        }
        // save the next total to largest
        // if it isn't larger than the last sum, then skip and move on
        //      to the next 13 digits in the sequence.
        if total > largest {
            largest = total
        }
    }
    return largest
}

func main() {
    fmt.Println(largestSum(13))
}
