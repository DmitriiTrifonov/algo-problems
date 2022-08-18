// https://stepik.org/lesson/13229/step/5?unit=3415

package main

import "fmt"

func main() {
    var a, b int
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Println(euclidean(a,b))
}

func euclidean(a, b int) int {
    if a == 0 {
        return b
    }
    if b == 0 {
        return a
    }
    if a >= b {
        return euclidean(a % b, b)
    }
    return euclidean(a, b % a)
}
