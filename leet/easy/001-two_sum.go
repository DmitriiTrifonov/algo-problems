// https://leetcode.com/problems/two-sum/

// This solving uses two for cycles and skips if indexes are the same
// If nums are making the needed sum they goes to output slice


func twoSum(nums []int, target int) []int {
    out := make([]int, 2)
    for i, num := range nums {
        for j, n := range nums {
            if i == j {
                continue
            }
            if num + n == target {
                out[0] = i
                out[1] = j
                return out
            }
        }
    }
    return out
}
