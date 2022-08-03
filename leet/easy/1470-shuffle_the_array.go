// https://leetcode.com/problems/shuffle-the-array/

// Take the data with shift and append it to slice

func shuffle(nums []int, n int) []int {
    out := make([]int, 0, len(nums))
    for i := 0; i < n; i++ {
        x := nums[i]
        y := nums[i + n]
        out = append(out, x)
        out = append(out, y)
    }
    return out
}
