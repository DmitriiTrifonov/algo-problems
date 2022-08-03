// https://leetcode.com/problems/jewels-and-stones/

// Two cycles and condition

func numJewelsInStones(jewels string, stones string) int {
    num := 0
    for s := range stones {
        for j := range jewels {
            if stones[s] == jewels[j] {
                num++
                break
            }
        }
    }
    return num
}
