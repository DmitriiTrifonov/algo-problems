// https://stepik.org/lesson/13228/step/7

package main

import "fmt"

func main() {
    var input int
    fmt.Scan(&input)
    fmt.Println(fibonacciLastDigitTable(input))
}

func fibonacciLastDigitTable(n int) int {
    if n == 1 {
         n = 2
    }
    table := make([]int, n)
    table[0] = 1
    table[1] = 1
    for i := 2; i < n; i++ {
        table[i] = (table[i - 1] + table[i - 2]) % 10
    }
    return table[n - 1]
}
