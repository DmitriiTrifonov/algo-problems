// https://stepik.org/lesson/13228/step/6

package main

import "fmt"

func main() {
    var input int
    fmt.Scan(&input)
    fmt.Println(fibonacciTable(input))
}

func fibonacciTable(n int) int {
    // Corner case
    if n == 1 {
         n = 2
    }
    table := make([]int, n)
    table[0] = 1
    table[1] = 1
    for i := 2; i < n; i++ {
        table[i] = table[i - 1] + table[i - 2]
    }
    return table[n - 1]
}
