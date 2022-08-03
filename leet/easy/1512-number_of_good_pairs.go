// https://leetcode.com/problems/number-of-good-pairs

// Two cycles and one condition

func numIdenticalPairs(nums []int) int {
    goodPairs := 0
    for i := range nums {
        for j := range nums {
            if i < j && nums[i] == nums[j] {
                goodPairs++
            }
        }
    }
    return goodPairs
}
