//build-array-from-permutation

func buildArray(nums []int) []int {
    out := make([]int, len(nums))

    for i := range out {
        out[i] = nums[nums[i]]
    }

    return out
}
