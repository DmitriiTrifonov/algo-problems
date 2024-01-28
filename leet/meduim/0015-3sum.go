// WIP https://leetcode.com/problems/3sum/

func threeSum(nums []int) [][]int {
    out := make([][]int, 0)
    for i := range nums {
        num1 := nums[i]
        for j := range nums {
            if i == j {
                continue 
            }
            num2 := nums[j]
            for k := range nums {
                if i == k || j == k {
                    continue 
                }
                num3 := nums[k]
                if num1 + num2 + num3 == 0 {
                a, b, c := sortNums(num1, num2, num3)
                if checkDupls(out, a, b, c) {
                    out = append(out, []int{a, b, c})
                }
            }
            }
        }
    }
    return out
}

func sortNums(a, b, c int) (int, int, int) {
    if a > c {
        a, c = c, a
    }
    if b > c {
        b, c = c, b
    }
    if a > b {
        a, b = b, a
    }

    return a, b, c
}

func checkDupls(table [][]int, a, b, c int) bool {
    if len(table) == 0 {
        return true
    }
    duplList := make([]bool, len(table))
    for i, sl := range table {
                    dupl := sl[0] == a && sl[1] == b && sl[2] == c
                    duplList[i] = dupl
                }
    
    for _, val := range duplList {
        if val {
            return false
        }
    }
    return true
}
