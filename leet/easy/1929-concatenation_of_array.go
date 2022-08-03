// https://leetcode.com/problems/concatenation-of-array/

// O(n) with shifting and alocation

func getConcatenation(nums []int) []int {
    shift := len(nums)
    out := make([]int, shift * 2)
    for i := 0; i < shift; i++ {
        out[i] = nums[i]
        out[shift + i] = nums[i]
    }
    return out
}
