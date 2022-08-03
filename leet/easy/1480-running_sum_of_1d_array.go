// https://leetcode.com/problems/running-sum-of-1d-array/

// Simply save previous number in the cycle

func runningSum(nums []int) []int {
    prev := 0
    for i, n := range nums {
        num := n + prev
        nums[i] = num
        prev = num
    }
    return nums
}
