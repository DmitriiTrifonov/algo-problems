// https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/

// Simply count product and sum in a loop

func subtractProductAndSum(n int) int {
    product := 1
    sum := 0
    for n > 0 {
        num := n % 10 
        product *= num
        sum += num
        n /= 10
    }
    return product - sum
}
