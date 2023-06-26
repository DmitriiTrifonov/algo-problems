// https://leetcode.com/problems/left-and-right-sum-differences/

// Straight-forward solution with additional functions

func leftRightDifference(nums []int) []int {
    out := make([]int, len(nums))

    for i := range out {
        ls := leftSumFor(i, nums)
        rs := rightSumFor(i, nums)
        sum := mod(ls - rs)
        out[i] = sum
    }

    return out
}

func mod(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func leftSumFor(i int, arr []int) int {
    sum := 0
    if i == 0 {
        return sum
    }

    for j := i - 1; j > -1; j-- {
        sum += arr[j]
    }
    return sum
}

func rightSumFor(i int, arr []int) int {
    sum := 0
    if i == len(arr) - 1 {
        return sum
    }

    for j := i + 1; j < len(arr); j++ {
        sum += arr[j]
    }
    return sum
}
