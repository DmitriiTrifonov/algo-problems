// https://leetcode.com/problems/remove-duplicates-from-sorted-array/


// This solution uses map to check for duplicates


func removeDuplicates(nums []int) int {
	m := make(map[int]struct{}, 0)
	idx := 0
	for _, num := range nums {
		if _, ok := m[num]; ok {
			continue
		}
		m[num] = struct{}{}
		nums[idx] = num
		idx++
	}
	return idx
}
