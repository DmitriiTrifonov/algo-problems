// https://leetcode.com/problems/remove-element/

// This solution swaps elements for removal

func removeElement(nums []int, val int) int {
    var l int
    ln := len(nums) - 1
    for i := range nums {
        if nums[i] != val {
            l++
            continue
        }

        for nums[i] == nums[ln] && ln > 0{
            ln--
        }
        nums[i] = nums[ln]
        if ln > 0 {
            ln--
        }    
    }
    return l
}
