// https://leetcode.com/problems/powx-n/description/

// Uses recursion

func myPow(x float64, n int) float64 {
    if n == 0 {
        return 1
    }
    if n < 0 {
        n = -n
        x = 1/x
    }

    if n % 2 == 0 {
        return myPow(x*x, n/2)
    }
    return x*myPow(x*x, n/2)
}
